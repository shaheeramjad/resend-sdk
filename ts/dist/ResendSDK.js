"use strict";
// Resend Ts SDK
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.ResendSDK = exports.ResendEntityBase = exports.BaseFeature = exports.config = exports.stdutil = void 0;
const ApiKeyEntity_1 = require("./entity/ApiKeyEntity");
const AttachmentEntity_1 = require("./entity/AttachmentEntity");
const AudienceEntity_1 = require("./entity/AudienceEntity");
const AutomationEntity_1 = require("./entity/AutomationEntity");
const AutomationRunEntity_1 = require("./entity/AutomationRunEntity");
const BroadcastEntity_1 = require("./entity/BroadcastEntity");
const BroadcastClickedLinkEntity_1 = require("./entity/BroadcastClickedLinkEntity");
const BroadcastRecipientEntity_1 = require("./entity/BroadcastRecipientEntity");
const ContactEntity_1 = require("./entity/ContactEntity");
const ContactImportEntity_1 = require("./entity/ContactImportEntity");
const ContactPropertyEntity_1 = require("./entity/ContactPropertyEntity");
const ContactSegmentEntity_1 = require("./entity/ContactSegmentEntity");
const ContactTopicEntity_1 = require("./entity/ContactTopicEntity");
const DomainEntity_1 = require("./entity/DomainEntity");
const DomainClaimEntity_1 = require("./entity/DomainClaimEntity");
const EmailEntity_1 = require("./entity/EmailEntity");
const EmailBatchEntity_1 = require("./entity/EmailBatchEntity");
const EmailsMetricEntity_1 = require("./entity/EmailsMetricEntity");
const EventEntity_1 = require("./entity/EventEntity");
const InboxEntity_1 = require("./entity/InboxEntity");
const InboxAgentEntity_1 = require("./entity/InboxAgentEntity");
const InboxDraftEntity_1 = require("./entity/InboxDraftEntity");
const InboxLabelEntity_1 = require("./entity/InboxLabelEntity");
const InboxThreadEntity_1 = require("./entity/InboxThreadEntity");
const InboxThreadEmailEntity_1 = require("./entity/InboxThreadEmailEntity");
const InboxThreadLabelEntity_1 = require("./entity/InboxThreadLabelEntity");
const LogEntity_1 = require("./entity/LogEntity");
const OauthGrantEntity_1 = require("./entity/OauthGrantEntity");
const ReceivedAttachmentEntity_1 = require("./entity/ReceivedAttachmentEntity");
const ReceivedEmailEntity_1 = require("./entity/ReceivedEmailEntity");
const SegmentEntity_1 = require("./entity/SegmentEntity");
const SegmentContactEntity_1 = require("./entity/SegmentContactEntity");
const SuppressionEntity_1 = require("./entity/SuppressionEntity");
const SuppressionBatchAddEntity_1 = require("./entity/SuppressionBatchAddEntity");
const SuppressionBatchRemoveEntity_1 = require("./entity/SuppressionBatchRemoveEntity");
const TemplateEntity_1 = require("./entity/TemplateEntity");
const TopicEntity_1 = require("./entity/TopicEntity");
const UsageEntity_1 = require("./entity/UsageEntity");
const WebhookEntity_1 = require("./entity/WebhookEntity");
const WebhookEventEntity_1 = require("./entity/WebhookEventEntity");
const WebhookEventAttemptEntity_1 = require("./entity/WebhookEventAttemptEntity");
const node_util_1 = require("node:util");
const Config_1 = require("./Config");
Object.defineProperty(exports, "config", { enumerable: true, get: function () { return Config_1.config; } });
const ResendEntityBase_1 = require("./ResendEntityBase");
Object.defineProperty(exports, "ResendEntityBase", { enumerable: true, get: function () { return ResendEntityBase_1.ResendEntityBase; } });
const Utility_1 = require("./utility/Utility");
const ResultBodyUtility_1 = require("./utility/ResultBodyUtility");
const MakeRequestUtility_1 = require("./utility/MakeRequestUtility");
const PrepareMethodUtility_1 = require("./utility/PrepareMethodUtility");
const BaseFeature_1 = require("./feature/base/BaseFeature");
Object.defineProperty(exports, "BaseFeature", { enumerable: true, get: function () { return BaseFeature_1.BaseFeature; } });
const stdutil = new Utility_1.Utility();
exports.stdutil = stdutil;
class ResendSDK {
    _mode = 'live';
    _options;
    _utility = new Utility_1.Utility();
    _features;
    _rootctx;
    constructor(options) {
        this._rootctx = this._utility.makeContext({
            client: this,
            utility: this._utility,
            config: Config_1.config,
            options,
            shared: new WeakMap()
        });
        this._options = this._utility.makeOptions(this._rootctx);
        for (const key of ['_options', '_rootctx', '_features']) {
            Object.defineProperty(this, key, {
                value: this[key], enumerable: false, writable: true, configurable: true
            });
        }
        const struct = this._utility.struct;
        const getpath = struct.getpath;
        if (true === getpath(this._options.feature, 'test.active')) {
            this._mode = 'test';
        }
        this._rootctx.options = this._options;
        this._features = [];
        const featureAdd = this._utility.featureAdd;
        const featureInit = this._utility.featureInit;
        // Add features in the resolved order (makeOptions puts an explicit
        // array order first, else defaults to test-first). Ordering matters:
        // the `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
        // so `test` must be added before them to sit at the base of the chain.
        const extend = this._options.extend || [];
        const featureorder = getpath(this._options, '__derived__.featureorder') || [];
        for (const fname of featureorder) {
            const fopts = this._options.feature[fname] || {};
            if (fopts.active) {
                // An active name with no generated class is legal when an
                // extend-supplied instance carries that name (station's adopt
                // path): the instance is added below, positioned by its own
                // __after__ entry, so skip it here rather than fail construction.
                if (!this._rootctx.config.hasFeature(fname) &&
                    extend.some((f) => fname === f.name)) {
                    continue;
                }
                featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname));
            }
        }
        for (let f of extend) {
            featureAdd(this._rootctx, f);
        }
        for (let f of this._features) {
            featureInit(this._rootctx, f);
        }
        const featureHook = this._utility.featureHook;
        featureHook(this._rootctx, 'PostConstruct');
    }
    options() {
        return this._utility.struct.clone(this._options);
    }
    utility() {
        return this._utility.struct.clone(this._utility);
    }
    async prepare(fetchargs) {
        const utility = this._utility;
        const struct = utility.struct;
        const clone = struct.clone;
        const { makeContext, makeFetchDef, prepareHeaders, prepareAuth, } = utility;
        fetchargs = fetchargs || {};
        let ctx = makeContext({
            opname: 'prepare',
            ctrl: fetchargs.ctrl || {},
        }, this._rootctx);
        const options = this._options;
        const method = String(fetchargs.method || 'GET').toUpperCase();
        if (!(0, PrepareMethodUtility_1.allowed)(options.allow.method, method)) {
            return ctx.error('spec_method_allow', 'Method "' + method +
                '" not allowed by SDK option allow.method value: "' + options.allow.method + '"');
        }
        const spec = {
            base: options.base,
            prefix: options.prefix,
            suffix: options.suffix,
            path: fetchargs.path || '',
            method,
            params: fetchargs.params || {},
            query: fetchargs.query || {},
            headers: prepareHeaders(ctx),
            body: fetchargs.body,
            step: 'start',
        };
        ctx.spec = spec;
        if (fetchargs.headers) {
            const uheaders = fetchargs.headers;
            for (let key in uheaders) {
                spec.headers[key] = uheaders[key];
            }
        }
        const authResult = prepareAuth(ctx);
        if (authResult instanceof Error) {
            return authResult;
        }
        return makeFetchDef(ctx);
    }
    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    // either one reaches the same endpoint.
    async direct(fetchargs) {
        if (!(0, PrepareMethodUtility_1.allowed)(this._options.allow.op, 'direct')) {
            return {
                ok: false,
                err: new Error('ResendSDK: direct: operation not allowed by' +
                    ' SDK option allow.op value: "' + this._options.allow.op + '"'),
            };
        }
        return this._rawRequest(fetchargs);
    }
    // Ungated request path shared by direct() and graphql(), each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    async _rawRequest(fetchargs) {
        const utility = this._utility;
        const fetcher = utility.fetcher;
        const makeContext = utility.makeContext;
        const fetchdef = await this.prepare(fetchargs);
        if (fetchdef instanceof Error) {
            return { ok: false, err: utility.clean(this._rootctx, fetchdef) };
        }
        let ctx = makeContext({
            opname: 'direct',
            ctrl: (fetchargs || {}).ctrl || {},
        }, this._rootctx);
        try {
            if (true === fetchdef.signal?.aborted) {
                throw fetchdef.signal.reason;
            }
            const fetched = await fetcher(ctx, fetchdef.url, fetchdef);
            if (null == fetched) {
                return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') };
            }
            else if (fetched instanceof Error) {
                return { ok: false, err: utility.clean(ctx, (0, MakeRequestUtility_1.abortError)(ctx, fetched)) };
            }
            const status = fetched.status;
            // No body responses (204 No Content, 304 Not Modified) and explicit
            // zero content-length must skip JSON parsing — fetched.json() would
            // throw `Unexpected end of JSON input` on an empty body.
            const headers = fetched.headers;
            const contentLength = headers && 'function' === typeof headers.get
                ? headers.get('content-length')
                : (headers || {})['content-length'];
            const noBody = 204 === status || 304 === status || '0' === String(contentLength);
            let json = undefined;
            let err = undefined;
            if (!noBody) {
                let text = undefined;
                try {
                    const raw = fetched;
                    if ('function' === typeof raw.text) {
                        text = await raw.text();
                        json = '' === text.trim() ? undefined : JSON.parse(text);
                    }
                    else {
                        json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json;
                    }
                }
                catch (parseErr) {
                    if ('SyntaxError' !== parseErr?.name) {
                        throw parseErr;
                    }
                    err = (0, ResultBodyUtility_1.unreadableBody)(ctx, {
                        status, headers, text: text ?? parseErr.text, sent: fetchdef.headers,
                        failed: 200 <= status && status < 300 ? undefined :
                            ctx.error('request_status', 'request: ' + status + ': ' + fetched.statusText),
                    });
                }
            }
            return {
                ok: null == err && status >= 200 && status < 300,
                status,
                headers: fetched.headers,
                data: json,
                ...(null == err ? {} : { err: utility.clean(ctx, err) }),
            };
        }
        catch (err) {
            return { ok: false, err: utility.clean(ctx, (0, MakeRequestUtility_1.abortError)(ctx, err)) };
        }
    }
    async graphql(query, variables, ctrl) {
        const options = this._options;
        if (!(0, PrepareMethodUtility_1.allowed)(options.allow.op, 'graphql')) {
            return {
                ok: false,
                err: new Error('ResendSDK: graphql: operation not allowed by' +
                    ' SDK option allow.op value: "' + options.allow.op + '"'),
            };
        }
        const res = await this._rawRequest({
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: { query, variables: variables || {} },
            ctrl,
        });
        // Errors are read BEFORE any status check: a GraphQL parse or validation
        // failure comes back as HTTP 400 carrying the standard { errors: [...] }
        // body, and the raw path represents a non-2xx as { ok: false } with no
        // err — so returning early on status would discard the server's own
        // diagnostics, which are the only useful part of that response.
        const errors = null == res.data ? undefined : res.data.errors;
        if (null != errors && Array.isArray(errors) && 0 < errors.length) {
            const first = errors[0] || {};
            const err = new Error('ResendSDK: graphql: ' +
                (first.message || 'graphql error'));
            err.graphql = errors;
            return { ok: false, status: res.status, headers: res.headers, err, data: res.data };
        }
        return res;
    }
    // Entity access: `client.ApiKey().list()` / `client.ApiKey().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ApiKey(entopts) {
        const self = this;
        return new ApiKeyEntity_1.ApiKeyEntity(self, entopts);
    }
    // Entity access: `client.Attachment().list()` / `client.Attachment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Attachment(entopts) {
        const self = this;
        return new AttachmentEntity_1.AttachmentEntity(self, entopts);
    }
    // Entity access: `client.Audience().list()` / `client.Audience().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Audience(entopts) {
        const self = this;
        return new AudienceEntity_1.AudienceEntity(self, entopts);
    }
    // Entity access: `client.Automation().list()` / `client.Automation().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Automation(entopts) {
        const self = this;
        return new AutomationEntity_1.AutomationEntity(self, entopts);
    }
    // Entity access: `client.AutomationRun().list()` / `client.AutomationRun().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    AutomationRun(entopts) {
        const self = this;
        return new AutomationRunEntity_1.AutomationRunEntity(self, entopts);
    }
    // Entity access: `client.Broadcast().list()` / `client.Broadcast().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Broadcast(entopts) {
        const self = this;
        return new BroadcastEntity_1.BroadcastEntity(self, entopts);
    }
    // Entity access: `client.BroadcastClickedLink().list()` / `client.BroadcastClickedLink().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    BroadcastClickedLink(entopts) {
        const self = this;
        return new BroadcastClickedLinkEntity_1.BroadcastClickedLinkEntity(self, entopts);
    }
    // Entity access: `client.BroadcastRecipient().list()` / `client.BroadcastRecipient().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    BroadcastRecipient(entopts) {
        const self = this;
        return new BroadcastRecipientEntity_1.BroadcastRecipientEntity(self, entopts);
    }
    // Entity access: `client.Contact().list()` / `client.Contact().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Contact(entopts) {
        const self = this;
        return new ContactEntity_1.ContactEntity(self, entopts);
    }
    // Entity access: `client.ContactImport().list()` / `client.ContactImport().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ContactImport(entopts) {
        const self = this;
        return new ContactImportEntity_1.ContactImportEntity(self, entopts);
    }
    // Entity access: `client.ContactProperty().list()` / `client.ContactProperty().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ContactProperty(entopts) {
        const self = this;
        return new ContactPropertyEntity_1.ContactPropertyEntity(self, entopts);
    }
    // Entity access: `client.ContactSegment().list()` / `client.ContactSegment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ContactSegment(entopts) {
        const self = this;
        return new ContactSegmentEntity_1.ContactSegmentEntity(self, entopts);
    }
    // Entity access: `client.ContactTopic().list()` / `client.ContactTopic().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ContactTopic(entopts) {
        const self = this;
        return new ContactTopicEntity_1.ContactTopicEntity(self, entopts);
    }
    // Entity access: `client.Domain().list()` / `client.Domain().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Domain(entopts) {
        const self = this;
        return new DomainEntity_1.DomainEntity(self, entopts);
    }
    // Entity access: `client.DomainClaim().list()` / `client.DomainClaim().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    DomainClaim(entopts) {
        const self = this;
        return new DomainClaimEntity_1.DomainClaimEntity(self, entopts);
    }
    // Entity access: `client.Email().list()` / `client.Email().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Email(entopts) {
        const self = this;
        return new EmailEntity_1.EmailEntity(self, entopts);
    }
    // Entity access: `client.EmailBatch().list()` / `client.EmailBatch().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailBatch(entopts) {
        const self = this;
        return new EmailBatchEntity_1.EmailBatchEntity(self, entopts);
    }
    // Entity access: `client.EmailsMetric().list()` / `client.EmailsMetric().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    EmailsMetric(entopts) {
        const self = this;
        return new EmailsMetricEntity_1.EmailsMetricEntity(self, entopts);
    }
    // Entity access: `client.Event().list()` / `client.Event().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Event(entopts) {
        const self = this;
        return new EventEntity_1.EventEntity(self, entopts);
    }
    // Entity access: `client.Inbox().list()` / `client.Inbox().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Inbox(entopts) {
        const self = this;
        return new InboxEntity_1.InboxEntity(self, entopts);
    }
    // Entity access: `client.InboxAgent().list()` / `client.InboxAgent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InboxAgent(entopts) {
        const self = this;
        return new InboxAgentEntity_1.InboxAgentEntity(self, entopts);
    }
    // Entity access: `client.InboxDraft().list()` / `client.InboxDraft().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InboxDraft(entopts) {
        const self = this;
        return new InboxDraftEntity_1.InboxDraftEntity(self, entopts);
    }
    // Entity access: `client.InboxLabel().list()` / `client.InboxLabel().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InboxLabel(entopts) {
        const self = this;
        return new InboxLabelEntity_1.InboxLabelEntity(self, entopts);
    }
    // Entity access: `client.InboxThread().list()` / `client.InboxThread().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InboxThread(entopts) {
        const self = this;
        return new InboxThreadEntity_1.InboxThreadEntity(self, entopts);
    }
    // Entity access: `client.InboxThreadEmail().list()` / `client.InboxThreadEmail().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InboxThreadEmail(entopts) {
        const self = this;
        return new InboxThreadEmailEntity_1.InboxThreadEmailEntity(self, entopts);
    }
    // Entity access: `client.InboxThreadLabel().list()` / `client.InboxThreadLabel().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InboxThreadLabel(entopts) {
        const self = this;
        return new InboxThreadLabelEntity_1.InboxThreadLabelEntity(self, entopts);
    }
    // Entity access: `client.Log().list()` / `client.Log().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Log(entopts) {
        const self = this;
        return new LogEntity_1.LogEntity(self, entopts);
    }
    // Entity access: `client.OauthGrant().list()` / `client.OauthGrant().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    OauthGrant(entopts) {
        const self = this;
        return new OauthGrantEntity_1.OauthGrantEntity(self, entopts);
    }
    // Entity access: `client.ReceivedAttachment().list()` / `client.ReceivedAttachment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ReceivedAttachment(entopts) {
        const self = this;
        return new ReceivedAttachmentEntity_1.ReceivedAttachmentEntity(self, entopts);
    }
    // Entity access: `client.ReceivedEmail().list()` / `client.ReceivedEmail().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ReceivedEmail(entopts) {
        const self = this;
        return new ReceivedEmailEntity_1.ReceivedEmailEntity(self, entopts);
    }
    // Entity access: `client.Segment().list()` / `client.Segment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Segment(entopts) {
        const self = this;
        return new SegmentEntity_1.SegmentEntity(self, entopts);
    }
    // Entity access: `client.SegmentContact().list()` / `client.SegmentContact().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SegmentContact(entopts) {
        const self = this;
        return new SegmentContactEntity_1.SegmentContactEntity(self, entopts);
    }
    // Entity access: `client.Suppression().list()` / `client.Suppression().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Suppression(entopts) {
        const self = this;
        return new SuppressionEntity_1.SuppressionEntity(self, entopts);
    }
    // Entity access: `client.SuppressionBatchAdd().list()` / `client.SuppressionBatchAdd().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SuppressionBatchAdd(entopts) {
        const self = this;
        return new SuppressionBatchAddEntity_1.SuppressionBatchAddEntity(self, entopts);
    }
    // Entity access: `client.SuppressionBatchRemove().list()` / `client.SuppressionBatchRemove().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SuppressionBatchRemove(entopts) {
        const self = this;
        return new SuppressionBatchRemoveEntity_1.SuppressionBatchRemoveEntity(self, entopts);
    }
    // Entity access: `client.Template().list()` / `client.Template().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Template(entopts) {
        const self = this;
        return new TemplateEntity_1.TemplateEntity(self, entopts);
    }
    // Entity access: `client.Topic().list()` / `client.Topic().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Topic(entopts) {
        const self = this;
        return new TopicEntity_1.TopicEntity(self, entopts);
    }
    // Entity access: `client.Usage().list()` / `client.Usage().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Usage(entopts) {
        const self = this;
        return new UsageEntity_1.UsageEntity(self, entopts);
    }
    // Entity access: `client.Webhook().list()` / `client.Webhook().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Webhook(entopts) {
        const self = this;
        return new WebhookEntity_1.WebhookEntity(self, entopts);
    }
    // Entity access: `client.WebhookEvent().list()` / `client.WebhookEvent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    WebhookEvent(entopts) {
        const self = this;
        return new WebhookEventEntity_1.WebhookEventEntity(self, entopts);
    }
    // Entity access: `client.WebhookEventAttempt().list()` / `client.WebhookEventAttempt().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    WebhookEventAttempt(entopts) {
        const self = this;
        return new WebhookEventAttemptEntity_1.WebhookEventAttemptEntity(self, entopts);
    }
    static test(testoptsarg, sdkoptsarg) {
        const struct = stdutil.struct;
        const setpath = struct.setpath;
        const getdef = struct.getdef;
        const clone = struct.clone;
        const setprop = struct.setprop;
        const sdkopts = getdef(clone(sdkoptsarg), {});
        const testopts = getdef(clone(testoptsarg), {});
        setprop(testopts, 'active', true);
        setpath(sdkopts, 'feature.test', testopts);
        const testsdk = new ResendSDK(sdkopts);
        testsdk._mode = 'test';
        return testsdk;
    }
    tester(testopts, sdkopts) {
        return ResendSDK.test(testopts, sdkopts);
    }
    toJSON() {
        return { name: 'Resend' };
    }
    toString() {
        return 'Resend ' + this._utility.struct.jsonify(this.toJSON());
    }
    [node_util_1.inspect.custom]() {
        return this.toString();
    }
}
exports.ResendSDK = ResendSDK;
const SDK = ResendSDK;
exports.SDK = SDK;
//# sourceMappingURL=ResendSDK.js.map