# Resend SDK

from resend_sdk.utility.voxgig_struct import voxgig_struct as vs
from resend_sdk.core.utility_type import ResendUtility
from resend_sdk.core.spec import ResendSpec
from resend_sdk.core import helpers
from resend_sdk.utility.prepare_method import allowed
from resend_sdk.utility.result_body import unreadable_body

# Load utility registration (populates Utility._registrar)
from resend_sdk.utility import register

# Load features
from resend_sdk.feature.base_feature import ResendBaseFeature
from resend_sdk.features import _has_feature, _make_feature


class ResendSDK:
    # The options hold the credential. A slot keeps them reachable as
    # `client.options` and out of `vars(client)` and every attribute dump;
    # the dict entry keeps the instance open for everything else.
    __slots__ = ("_options", "__dict__")

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self._options = None

        utility = ResendUtility()
        self._utility = utility

        from resend_sdk.config import shared_config
        config = shared_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features in the resolved order (make_options puts an explicit
        # list order first, else defaults to test-first). Ordering matters: the
        # `test` feature installs the base mock transport and the transport
        # features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        # current, so `test` must be added before them to sit at the base.
        # Extension feature INSTANCES come from the RAW construction
        # options - extend is consumed exactly once, here. make_options
        # strips the key before cloning (vs.clone flattens arbitrary
        # objects), so self.options never carries the instances.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        extend = options.get("extend") if isinstance(options, dict) else None
        if not isinstance(extend, list):
            extend = []
        if feature_opts is not None:
            featureorder = vs.getpath(self.options, "__derived__.featureorder")
            if isinstance(featureorder, list):
                for fname in featureorder:
                    fopts = helpers.to_map(feature_opts.get(fname))
                    if fopts is not None and fopts.get("active") is True:
                        # An active name with no generated feature class is
                        # legal when an extend-supplied instance carries that
                        # name (station's adopt path): the instance is added
                        # below, positioned by its own __after__ entry, so
                        # skip it here rather than add a BaseFeature stray
                        # that would silently shift feature positions.
                        if not _has_feature(fname) and any(
                            fname == (f.get("name") if isinstance(f, dict)
                                      else getattr(f, "name", None))
                            for f in extend
                        ):
                            continue
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        for f in extend:
            if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    @property
    def options(self):
        return self._options

    @options.setter
    def options(self, value):
        self._options = value

    def __repr__(self):
        return "ResendSDK(mode=" + repr(self.mode) + ")"

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return ResendUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"
        method = method.upper()

        allow_method = vs.getpath(options, "allow.method")
        if not allowed(allow_method, method):
            raise ctx.make_error("spec_method_allow",
                'Method "' + method +
                '" not allowed by SDK option allow.method value: "' + str(allow_method) + '"')

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = ResendSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    # Raw endpoint access is operator-controllable, like every entity op.
    # Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    # either one reaches the same endpoint.
    def direct(self, fetchargs=None):
        if not self._op_allowed("direct"):
            return self._op_denied("direct")

        return self._raw_request(fetchargs)

    # Is this raw-access op permitted by the SDK's allow.op option?
    def _op_allowed(self, op):
        return allowed(vs.getpath(self.options, "allow.op"), op)

    def _op_denied(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return {
            "ok": False,
            "err": Exception(
                "ResendSDK: " + op + ": operation not allowed by"
                ' SDK option allow.op value: "' + str(allow_op) + '"'),
        }

    # Ungated request path shared by direct and graphql, each of which checks
    # its own allow.op token first. Private, rather than a flag on fetchargs:
    # a caller-supplied marker would let anyone opt straight back out of the
    # gate by passing it.
    def _raw_request(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            # That error never passes through make_error, so it is cleaned.
            return {"ok": False, "err": utility.clean(self._rootctx, err)}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": utility.clean(ctx, fetch_err)}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            body_err = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None
                if vs.getprop(fetched, "unreadable") is True:
                    failed = None if 200 <= status < 300 else ctx.make_error(
                        "request_status",
                        "request: " + str(status) + ": " + str(vs.getprop(fetched, "statusText")))
                    body_err = unreadable_body(ctx, status, headers, vs.getprop(fetched, "body"),
                                               fetchdef.get("headers"), failed)

            out = {
                "ok": body_err is None and status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }
            if body_err is not None:
                out["err"] = utility.clean(ctx, body_err)
            return out

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }

    # Raw GraphQL access: the pressure valve that makes the generated
    # surface's deliberate omissions (per-call selection sets, typed filter
    # builders, batching, subscriptions) livable — the whole schema stays
    # reachable.
    #
    # Thin wrapper over the same prepare/fetch path direct uses, with the one
    # thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
    # as a top-level `errors` array, so status alone would report a failed
    # query as ok.
    #
    # NOTE: like direct, this bypasses the feature pipeline — no retry,
    # ratelimit or paging features apply.
    def graphql(self, query, variables=None, ctrl=None):
        if not self._op_allowed("graphql"):
            return self._op_denied("graphql")

        res = self._raw_request({
            "method": "POST",
            "headers": {"content-type": "application/json"},
            "body": {"query": query, "variables": variables or {}},
            "ctrl": ctrl or {},
        })

        # Errors are read BEFORE any status check: a GraphQL parse or
        # validation failure comes back as HTTP 400 carrying the standard
        # { errors: [...] } body, and the raw path represents a non-2xx as
        # ok:False with no err — so returning early on status would discard
        # the server's own diagnostics, which are the only useful part of
        # that response.
        errors = vs.getpath(res, "data.errors")

        if isinstance(errors, list) and 0 < len(errors):
            first = errors[0] if isinstance(errors[0], dict) else {}
            msg = first.get("message") or "graphql error"
            res["ok"] = False
            res["err"] = Exception("ResendSDK: graphql: " + str(msg))
            res["graphql"] = errors

        return res


    def ApiKey(self, data=None) -> "ApiKeyEntity":
        """Entity factory: client.ApiKey().list() / client.ApiKey().load({"id": ...})."""
        from resend_sdk.entity.api_key_entity import ApiKeyEntity
        return ApiKeyEntity(self, data)


    def Attachment(self, data=None) -> "AttachmentEntity":
        """Entity factory: client.Attachment().list() / client.Attachment().load({"id": ...})."""
        from resend_sdk.entity.attachment_entity import AttachmentEntity
        return AttachmentEntity(self, data)


    def Audience(self, data=None) -> "AudienceEntity":
        """Entity factory: client.Audience().list() / client.Audience().load({"id": ...})."""
        from resend_sdk.entity.audience_entity import AudienceEntity
        return AudienceEntity(self, data)


    def Automation(self, data=None) -> "AutomationEntity":
        """Entity factory: client.Automation().list() / client.Automation().load({"id": ...})."""
        from resend_sdk.entity.automation_entity import AutomationEntity
        return AutomationEntity(self, data)


    def AutomationRun(self, data=None) -> "AutomationRunEntity":
        """Entity factory: client.AutomationRun().list() / client.AutomationRun().load({"id": ...})."""
        from resend_sdk.entity.automation_run_entity import AutomationRunEntity
        return AutomationRunEntity(self, data)


    def Broadcast(self, data=None) -> "BroadcastEntity":
        """Entity factory: client.Broadcast().list() / client.Broadcast().load({"id": ...})."""
        from resend_sdk.entity.broadcast_entity import BroadcastEntity
        return BroadcastEntity(self, data)


    def BroadcastClickedLink(self, data=None) -> "BroadcastClickedLinkEntity":
        """Entity factory: client.BroadcastClickedLink().list() / client.BroadcastClickedLink().load({"id": ...})."""
        from resend_sdk.entity.broadcast_clicked_link_entity import BroadcastClickedLinkEntity
        return BroadcastClickedLinkEntity(self, data)


    def BroadcastRecipient(self, data=None) -> "BroadcastRecipientEntity":
        """Entity factory: client.BroadcastRecipient().list() / client.BroadcastRecipient().load({"id": ...})."""
        from resend_sdk.entity.broadcast_recipient_entity import BroadcastRecipientEntity
        return BroadcastRecipientEntity(self, data)


    def Contact(self, data=None) -> "ContactEntity":
        """Entity factory: client.Contact().list() / client.Contact().load({"id": ...})."""
        from resend_sdk.entity.contact_entity import ContactEntity
        return ContactEntity(self, data)


    def ContactImport(self, data=None) -> "ContactImportEntity":
        """Entity factory: client.ContactImport().list() / client.ContactImport().load({"id": ...})."""
        from resend_sdk.entity.contact_import_entity import ContactImportEntity
        return ContactImportEntity(self, data)


    def ContactProperty(self, data=None) -> "ContactPropertyEntity":
        """Entity factory: client.ContactProperty().list() / client.ContactProperty().load({"id": ...})."""
        from resend_sdk.entity.contact_property_entity import ContactPropertyEntity
        return ContactPropertyEntity(self, data)


    def ContactSegment(self, data=None) -> "ContactSegmentEntity":
        """Entity factory: client.ContactSegment().list() / client.ContactSegment().load({"id": ...})."""
        from resend_sdk.entity.contact_segment_entity import ContactSegmentEntity
        return ContactSegmentEntity(self, data)


    def ContactTopic(self, data=None) -> "ContactTopicEntity":
        """Entity factory: client.ContactTopic().list() / client.ContactTopic().load({"id": ...})."""
        from resend_sdk.entity.contact_topic_entity import ContactTopicEntity
        return ContactTopicEntity(self, data)


    def Domain(self, data=None) -> "DomainEntity":
        """Entity factory: client.Domain().list() / client.Domain().load({"id": ...})."""
        from resend_sdk.entity.domain_entity import DomainEntity
        return DomainEntity(self, data)


    def DomainClaim(self, data=None) -> "DomainClaimEntity":
        """Entity factory: client.DomainClaim().list() / client.DomainClaim().load({"id": ...})."""
        from resend_sdk.entity.domain_claim_entity import DomainClaimEntity
        return DomainClaimEntity(self, data)


    def Email(self, data=None) -> "EmailEntity":
        """Entity factory: client.Email().list() / client.Email().load({"id": ...})."""
        from resend_sdk.entity.email_entity import EmailEntity
        return EmailEntity(self, data)


    def EmailBatch(self, data=None) -> "EmailBatchEntity":
        """Entity factory: client.EmailBatch().list() / client.EmailBatch().load({"id": ...})."""
        from resend_sdk.entity.email_batch_entity import EmailBatchEntity
        return EmailBatchEntity(self, data)


    def EmailsMetric(self, data=None) -> "EmailsMetricEntity":
        """Entity factory: client.EmailsMetric().list() / client.EmailsMetric().load({"id": ...})."""
        from resend_sdk.entity.emails_metric_entity import EmailsMetricEntity
        return EmailsMetricEntity(self, data)


    def Event(self, data=None) -> "EventEntity":
        """Entity factory: client.Event().list() / client.Event().load({"id": ...})."""
        from resend_sdk.entity.event_entity import EventEntity
        return EventEntity(self, data)


    def Inbox(self, data=None) -> "InboxEntity":
        """Entity factory: client.Inbox().list() / client.Inbox().load({"id": ...})."""
        from resend_sdk.entity.inbox_entity import InboxEntity
        return InboxEntity(self, data)


    def InboxAgent(self, data=None) -> "InboxAgentEntity":
        """Entity factory: client.InboxAgent().list() / client.InboxAgent().load({"id": ...})."""
        from resend_sdk.entity.inbox_agent_entity import InboxAgentEntity
        return InboxAgentEntity(self, data)


    def InboxDraft(self, data=None) -> "InboxDraftEntity":
        """Entity factory: client.InboxDraft().list() / client.InboxDraft().load({"id": ...})."""
        from resend_sdk.entity.inbox_draft_entity import InboxDraftEntity
        return InboxDraftEntity(self, data)


    def InboxLabel(self, data=None) -> "InboxLabelEntity":
        """Entity factory: client.InboxLabel().list() / client.InboxLabel().load({"id": ...})."""
        from resend_sdk.entity.inbox_label_entity import InboxLabelEntity
        return InboxLabelEntity(self, data)


    def InboxThread(self, data=None) -> "InboxThreadEntity":
        """Entity factory: client.InboxThread().list() / client.InboxThread().load({"id": ...})."""
        from resend_sdk.entity.inbox_thread_entity import InboxThreadEntity
        return InboxThreadEntity(self, data)


    def InboxThreadEmail(self, data=None) -> "InboxThreadEmailEntity":
        """Entity factory: client.InboxThreadEmail().list() / client.InboxThreadEmail().load({"id": ...})."""
        from resend_sdk.entity.inbox_thread_email_entity import InboxThreadEmailEntity
        return InboxThreadEmailEntity(self, data)


    def InboxThreadLabel(self, data=None) -> "InboxThreadLabelEntity":
        """Entity factory: client.InboxThreadLabel().list() / client.InboxThreadLabel().load({"id": ...})."""
        from resend_sdk.entity.inbox_thread_label_entity import InboxThreadLabelEntity
        return InboxThreadLabelEntity(self, data)


    def Log(self, data=None) -> "LogEntity":
        """Entity factory: client.Log().list() / client.Log().load({"id": ...})."""
        from resend_sdk.entity.log_entity import LogEntity
        return LogEntity(self, data)


    def OauthGrant(self, data=None) -> "OauthGrantEntity":
        """Entity factory: client.OauthGrant().list() / client.OauthGrant().load({"id": ...})."""
        from resend_sdk.entity.oauth_grant_entity import OauthGrantEntity
        return OauthGrantEntity(self, data)


    def ReceivedAttachment(self, data=None) -> "ReceivedAttachmentEntity":
        """Entity factory: client.ReceivedAttachment().list() / client.ReceivedAttachment().load({"id": ...})."""
        from resend_sdk.entity.received_attachment_entity import ReceivedAttachmentEntity
        return ReceivedAttachmentEntity(self, data)


    def ReceivedEmail(self, data=None) -> "ReceivedEmailEntity":
        """Entity factory: client.ReceivedEmail().list() / client.ReceivedEmail().load({"id": ...})."""
        from resend_sdk.entity.received_email_entity import ReceivedEmailEntity
        return ReceivedEmailEntity(self, data)


    def Segment(self, data=None) -> "SegmentEntity":
        """Entity factory: client.Segment().list() / client.Segment().load({"id": ...})."""
        from resend_sdk.entity.segment_entity import SegmentEntity
        return SegmentEntity(self, data)


    def SegmentContact(self, data=None) -> "SegmentContactEntity":
        """Entity factory: client.SegmentContact().list() / client.SegmentContact().load({"id": ...})."""
        from resend_sdk.entity.segment_contact_entity import SegmentContactEntity
        return SegmentContactEntity(self, data)


    def Suppression(self, data=None) -> "SuppressionEntity":
        """Entity factory: client.Suppression().list() / client.Suppression().load({"id": ...})."""
        from resend_sdk.entity.suppression_entity import SuppressionEntity
        return SuppressionEntity(self, data)


    def SuppressionBatchAdd(self, data=None) -> "SuppressionBatchAddEntity":
        """Entity factory: client.SuppressionBatchAdd().list() / client.SuppressionBatchAdd().load({"id": ...})."""
        from resend_sdk.entity.suppression_batch_add_entity import SuppressionBatchAddEntity
        return SuppressionBatchAddEntity(self, data)


    def SuppressionBatchRemove(self, data=None) -> "SuppressionBatchRemoveEntity":
        """Entity factory: client.SuppressionBatchRemove().list() / client.SuppressionBatchRemove().load({"id": ...})."""
        from resend_sdk.entity.suppression_batch_remove_entity import SuppressionBatchRemoveEntity
        return SuppressionBatchRemoveEntity(self, data)


    def Template(self, data=None) -> "TemplateEntity":
        """Entity factory: client.Template().list() / client.Template().load({"id": ...})."""
        from resend_sdk.entity.template_entity import TemplateEntity
        return TemplateEntity(self, data)


    def Topic(self, data=None) -> "TopicEntity":
        """Entity factory: client.Topic().list() / client.Topic().load({"id": ...})."""
        from resend_sdk.entity.topic_entity import TopicEntity
        return TopicEntity(self, data)


    def Usage(self, data=None) -> "UsageEntity":
        """Entity factory: client.Usage().list() / client.Usage().load({"id": ...})."""
        from resend_sdk.entity.usage_entity import UsageEntity
        return UsageEntity(self, data)


    def Webhook(self, data=None) -> "WebhookEntity":
        """Entity factory: client.Webhook().list() / client.Webhook().load({"id": ...})."""
        from resend_sdk.entity.webhook_entity import WebhookEntity
        return WebhookEntity(self, data)


    def WebhookEvent(self, data=None) -> "WebhookEventEntity":
        """Entity factory: client.WebhookEvent().list() / client.WebhookEvent().load({"id": ...})."""
        from resend_sdk.entity.webhook_event_entity import WebhookEventEntity
        return WebhookEventEntity(self, data)


    def WebhookEventAttempt(self, data=None) -> "WebhookEventAttemptEntity":
        """Entity factory: client.WebhookEventAttempt().list() / client.WebhookEventAttempt().load({"id": ...})."""
        from resend_sdk.entity.webhook_event_attempt_entity import WebhookEventAttemptEntity
        return WebhookEventAttemptEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "ResendSDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from resend_sdk.entity.api_key_entity import ApiKeyEntity
    from resend_sdk.entity.attachment_entity import AttachmentEntity
    from resend_sdk.entity.audience_entity import AudienceEntity
    from resend_sdk.entity.automation_entity import AutomationEntity
    from resend_sdk.entity.automation_run_entity import AutomationRunEntity
    from resend_sdk.entity.broadcast_entity import BroadcastEntity
    from resend_sdk.entity.broadcast_clicked_link_entity import BroadcastClickedLinkEntity
    from resend_sdk.entity.broadcast_recipient_entity import BroadcastRecipientEntity
    from resend_sdk.entity.contact_entity import ContactEntity
    from resend_sdk.entity.contact_import_entity import ContactImportEntity
    from resend_sdk.entity.contact_property_entity import ContactPropertyEntity
    from resend_sdk.entity.contact_segment_entity import ContactSegmentEntity
    from resend_sdk.entity.contact_topic_entity import ContactTopicEntity
    from resend_sdk.entity.domain_entity import DomainEntity
    from resend_sdk.entity.domain_claim_entity import DomainClaimEntity
    from resend_sdk.entity.email_entity import EmailEntity
    from resend_sdk.entity.email_batch_entity import EmailBatchEntity
    from resend_sdk.entity.emails_metric_entity import EmailsMetricEntity
    from resend_sdk.entity.event_entity import EventEntity
    from resend_sdk.entity.inbox_entity import InboxEntity
    from resend_sdk.entity.inbox_agent_entity import InboxAgentEntity
    from resend_sdk.entity.inbox_draft_entity import InboxDraftEntity
    from resend_sdk.entity.inbox_label_entity import InboxLabelEntity
    from resend_sdk.entity.inbox_thread_entity import InboxThreadEntity
    from resend_sdk.entity.inbox_thread_email_entity import InboxThreadEmailEntity
    from resend_sdk.entity.inbox_thread_label_entity import InboxThreadLabelEntity
    from resend_sdk.entity.log_entity import LogEntity
    from resend_sdk.entity.oauth_grant_entity import OauthGrantEntity
    from resend_sdk.entity.received_attachment_entity import ReceivedAttachmentEntity
    from resend_sdk.entity.received_email_entity import ReceivedEmailEntity
    from resend_sdk.entity.segment_entity import SegmentEntity
    from resend_sdk.entity.segment_contact_entity import SegmentContactEntity
    from resend_sdk.entity.suppression_entity import SuppressionEntity
    from resend_sdk.entity.suppression_batch_add_entity import SuppressionBatchAddEntity
    from resend_sdk.entity.suppression_batch_remove_entity import SuppressionBatchRemoveEntity
    from resend_sdk.entity.template_entity import TemplateEntity
    from resend_sdk.entity.topic_entity import TopicEntity
    from resend_sdk.entity.usage_entity import UsageEntity
    from resend_sdk.entity.webhook_entity import WebhookEntity
    from resend_sdk.entity.webhook_event_entity import WebhookEventEntity
    from resend_sdk.entity.webhook_event_attempt_entity import WebhookEventAttemptEntity
