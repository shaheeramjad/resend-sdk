package core

import (
	"encoding/json"
	"fmt"
	"strings"

	vs "github.com/shaheeramjad/resend-sdk/go/utility/struct"
)

type ResendSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewResendSDK(options map[string]any) *ResendSDK {
	sdk := &ResendSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := SharedConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath(sdk.options, []any{"feature", "test", "active"}) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath(sdk.options, []any{"__derived__", "featureorder"}).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

// The client holds the credential in its options, so a print or a JSON dump
// carries the name alone. Value receivers: a dereferenced client prints the
// same way.
func (sdk ResendSDK) String() string {
	return "Resend " + vs.Jsonify(map[string]any{"name": "Resend"},
		map[string]any{"indent": 0})
}

func (sdk ResendSDK) GoString() string {
	return sdk.String()
}

func (sdk ResendSDK) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{"name": "Resend"})
}

func (sdk *ResendSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *ResendSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *ResendSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *ResendSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}
	method = strings.ToUpper(method)

	allowMethodVal := vs.GetPath(options, []any{"allow", "method"})
	if !Allowed(allowMethodVal, method) {
		allowMethod, _ := allowMethodVal.(string)
		return nil, ctx.MakeError("spec_method_allow",
			"Method \""+method+"\" not allowed by SDK option allow.method value: \""+allowMethod+"\"")
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

// Raw endpoint access is operator-controllable, like every entity op.
// Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
// either one reaches the same endpoint.
func (sdk *ResendSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	if !sdk.opAllowed("direct") {
		return sdk.opDenied("direct"), nil
	}

	return sdk.rawRequest(fetchargs)
}

// Is this raw-access op permitted by the SDK's allow.op option?
func (sdk *ResendSDK) opAllowed(op string) bool {
	return Allowed(vs.GetPath(sdk.options, []any{"allow", "op"}), op)
}

func (sdk *ResendSDK) opDenied(op string) map[string]any {
	allowOp, _ := vs.GetPath(sdk.options, []any{"allow", "op"}).(string)
	return map[string]any{
		"ok": false,
		"err": fmt.Errorf("ResendSDK: %s: operation not allowed by"+
			" SDK option allow.op value: \"%s\"", op, allowOp),
	}
}

// Ungated request path shared by Direct and Graphql, each of which checks
// its own allow.op token first. Unexported, rather than a flag on fetchargs:
// a caller-supplied marker would let anyone opt straight back out of the
// gate by passing it.
func (sdk *ResendSDK) rawRequest(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": sdk.cleanErr(sdk.rootctx, err)}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": sdk.cleanErr(ctx, fetchErr)}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		var bodyErr error
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					jsonData = f()
				}
			}
			if unreadable, _ := vs.GetProp(fm, "unreadable").(bool); unreadable {
				var failed error
				if status < 200 || status >= 300 {
					failed = ctx.MakeError("request_status",
						fmt.Sprintf("request: %d: %v", status, vs.GetProp(fm, "statusText")))
				}
				bodyErr = UnreadableBody(ctx, status, headers, vs.GetProp(fm, "body"),
					fetchdef["headers"], failed)
			}
		}

		out := map[string]any{
			"ok":      bodyErr == nil && status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}
		if bodyErr != nil {
			out["err"] = sdk.cleanErr(ctx, bodyErr)
		}
		return out, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}

// A raw request returns its error rather than passing it through MakeError.
func (sdk *ResendSDK) cleanErr(ctx *Context, err error) error {
	if cleaned, ok := sdk.utility.Clean(ctx, err).(error); ok {
		return cleaned
	}
	return err
}

func (sdk *ResendSDK) Graphql(
	query string, variables map[string]any, ctrl map[string]any,
) (map[string]any, error) {
	if !sdk.opAllowed("graphql") {
		return sdk.opDenied("graphql"), nil
	}

	if variables == nil {
		variables = map[string]any{}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	res, err := sdk.rawRequest(map[string]any{
		"method":  "POST",
		"headers": map[string]any{"content-type": "application/json"},
		"body":    map[string]any{"query": query, "variables": variables},
		"ctrl":    ctrl,
	})

	if err != nil {
		return res, err
	}

	// Errors are read BEFORE any status check: a GraphQL parse or validation
	// failure comes back as HTTP 400 carrying the standard { errors: [...] }
	// body, and the raw path represents a non-2xx as ok:false with no err —
	// so returning early on status would discard the server's own
	// diagnostics, which are the only useful part of that response.
	errors, _ := vs.GetPath(res, []any{"data", "errors"}).([]any)

	if 0 < len(errors) {
		msg, _ := vs.GetProp(errors[0], "message").(string)
		if msg == "" {
			msg = "graphql error"
		}
		res["ok"] = false
		res["err"] = fmt.Errorf("ResendSDK: graphql: %s", msg)
		res["graphql"] = errors
	}

	return res, nil
}


// ApiKey returns a ApiKey entity bound to this client.
// Idiomatic usage: client.ApiKey(nil).List(nil, nil) or
// client.ApiKey(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ApiKey(data map[string]any) ResendEntity {
	return NewApiKeyEntityFunc(sdk, data)
}


// Attachment returns a Attachment entity bound to this client.
// Idiomatic usage: client.Attachment(nil).List(nil, nil) or
// client.Attachment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Attachment(data map[string]any) ResendEntity {
	return NewAttachmentEntityFunc(sdk, data)
}


// Audience returns a Audience entity bound to this client.
// Idiomatic usage: client.Audience(nil).List(nil, nil) or
// client.Audience(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Audience(data map[string]any) ResendEntity {
	return NewAudienceEntityFunc(sdk, data)
}


// Automation returns a Automation entity bound to this client.
// Idiomatic usage: client.Automation(nil).List(nil, nil) or
// client.Automation(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Automation(data map[string]any) ResendEntity {
	return NewAutomationEntityFunc(sdk, data)
}


// AutomationRun returns a AutomationRun entity bound to this client.
// Idiomatic usage: client.AutomationRun(nil).List(nil, nil) or
// client.AutomationRun(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) AutomationRun(data map[string]any) ResendEntity {
	return NewAutomationRunEntityFunc(sdk, data)
}


// Broadcast returns a Broadcast entity bound to this client.
// Idiomatic usage: client.Broadcast(nil).List(nil, nil) or
// client.Broadcast(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Broadcast(data map[string]any) ResendEntity {
	return NewBroadcastEntityFunc(sdk, data)
}


// BroadcastClickedLink returns a BroadcastClickedLink entity bound to this client.
// Idiomatic usage: client.BroadcastClickedLink(nil).List(nil, nil) or
// client.BroadcastClickedLink(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) BroadcastClickedLink(data map[string]any) ResendEntity {
	return NewBroadcastClickedLinkEntityFunc(sdk, data)
}


// BroadcastRecipient returns a BroadcastRecipient entity bound to this client.
// Idiomatic usage: client.BroadcastRecipient(nil).List(nil, nil) or
// client.BroadcastRecipient(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) BroadcastRecipient(data map[string]any) ResendEntity {
	return NewBroadcastRecipientEntityFunc(sdk, data)
}


// Contact returns a Contact entity bound to this client.
// Idiomatic usage: client.Contact(nil).List(nil, nil) or
// client.Contact(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Contact(data map[string]any) ResendEntity {
	return NewContactEntityFunc(sdk, data)
}


// ContactImport returns a ContactImport entity bound to this client.
// Idiomatic usage: client.ContactImport(nil).List(nil, nil) or
// client.ContactImport(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactImport(data map[string]any) ResendEntity {
	return NewContactImportEntityFunc(sdk, data)
}


// ContactProperty returns a ContactProperty entity bound to this client.
// Idiomatic usage: client.ContactProperty(nil).List(nil, nil) or
// client.ContactProperty(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactProperty(data map[string]any) ResendEntity {
	return NewContactPropertyEntityFunc(sdk, data)
}


// ContactSegment returns a ContactSegment entity bound to this client.
// Idiomatic usage: client.ContactSegment(nil).List(nil, nil) or
// client.ContactSegment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactSegment(data map[string]any) ResendEntity {
	return NewContactSegmentEntityFunc(sdk, data)
}


// ContactTopic returns a ContactTopic entity bound to this client.
// Idiomatic usage: client.ContactTopic(nil).List(nil, nil) or
// client.ContactTopic(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactTopic(data map[string]any) ResendEntity {
	return NewContactTopicEntityFunc(sdk, data)
}


// Domain returns a Domain entity bound to this client.
// Idiomatic usage: client.Domain(nil).List(nil, nil) or
// client.Domain(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Domain(data map[string]any) ResendEntity {
	return NewDomainEntityFunc(sdk, data)
}


// DomainClaim returns a DomainClaim entity bound to this client.
// Idiomatic usage: client.DomainClaim(nil).List(nil, nil) or
// client.DomainClaim(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) DomainClaim(data map[string]any) ResendEntity {
	return NewDomainClaimEntityFunc(sdk, data)
}


// Email returns a Email entity bound to this client.
// Idiomatic usage: client.Email(nil).List(nil, nil) or
// client.Email(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Email(data map[string]any) ResendEntity {
	return NewEmailEntityFunc(sdk, data)
}


// EmailBatch returns a EmailBatch entity bound to this client.
// Idiomatic usage: client.EmailBatch(nil).List(nil, nil) or
// client.EmailBatch(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) EmailBatch(data map[string]any) ResendEntity {
	return NewEmailBatchEntityFunc(sdk, data)
}


// EmailsMetric returns a EmailsMetric entity bound to this client.
// Idiomatic usage: client.EmailsMetric(nil).List(nil, nil) or
// client.EmailsMetric(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) EmailsMetric(data map[string]any) ResendEntity {
	return NewEmailsMetricEntityFunc(sdk, data)
}


// Event returns a Event entity bound to this client.
// Idiomatic usage: client.Event(nil).List(nil, nil) or
// client.Event(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Event(data map[string]any) ResendEntity {
	return NewEventEntityFunc(sdk, data)
}


// Inbox returns a Inbox entity bound to this client.
// Idiomatic usage: client.Inbox(nil).List(nil, nil) or
// client.Inbox(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Inbox(data map[string]any) ResendEntity {
	return NewInboxEntityFunc(sdk, data)
}


// InboxAgent returns a InboxAgent entity bound to this client.
// Idiomatic usage: client.InboxAgent(nil).List(nil, nil) or
// client.InboxAgent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxAgent(data map[string]any) ResendEntity {
	return NewInboxAgentEntityFunc(sdk, data)
}


// InboxDraft returns a InboxDraft entity bound to this client.
// Idiomatic usage: client.InboxDraft(nil).List(nil, nil) or
// client.InboxDraft(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxDraft(data map[string]any) ResendEntity {
	return NewInboxDraftEntityFunc(sdk, data)
}


// InboxLabel returns a InboxLabel entity bound to this client.
// Idiomatic usage: client.InboxLabel(nil).List(nil, nil) or
// client.InboxLabel(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxLabel(data map[string]any) ResendEntity {
	return NewInboxLabelEntityFunc(sdk, data)
}


// InboxThread returns a InboxThread entity bound to this client.
// Idiomatic usage: client.InboxThread(nil).List(nil, nil) or
// client.InboxThread(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxThread(data map[string]any) ResendEntity {
	return NewInboxThreadEntityFunc(sdk, data)
}


// InboxThreadEmail returns a InboxThreadEmail entity bound to this client.
// Idiomatic usage: client.InboxThreadEmail(nil).List(nil, nil) or
// client.InboxThreadEmail(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxThreadEmail(data map[string]any) ResendEntity {
	return NewInboxThreadEmailEntityFunc(sdk, data)
}


// InboxThreadLabel returns a InboxThreadLabel entity bound to this client.
// Idiomatic usage: client.InboxThreadLabel(nil).List(nil, nil) or
// client.InboxThreadLabel(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxThreadLabel(data map[string]any) ResendEntity {
	return NewInboxThreadLabelEntityFunc(sdk, data)
}


// Log returns a Log entity bound to this client.
// Idiomatic usage: client.Log(nil).List(nil, nil) or
// client.Log(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Log(data map[string]any) ResendEntity {
	return NewLogEntityFunc(sdk, data)
}


// OauthGrant returns a OauthGrant entity bound to this client.
// Idiomatic usage: client.OauthGrant(nil).List(nil, nil) or
// client.OauthGrant(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) OauthGrant(data map[string]any) ResendEntity {
	return NewOauthGrantEntityFunc(sdk, data)
}


// ReceivedAttachment returns a ReceivedAttachment entity bound to this client.
// Idiomatic usage: client.ReceivedAttachment(nil).List(nil, nil) or
// client.ReceivedAttachment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ReceivedAttachment(data map[string]any) ResendEntity {
	return NewReceivedAttachmentEntityFunc(sdk, data)
}


// ReceivedEmail returns a ReceivedEmail entity bound to this client.
// Idiomatic usage: client.ReceivedEmail(nil).List(nil, nil) or
// client.ReceivedEmail(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ReceivedEmail(data map[string]any) ResendEntity {
	return NewReceivedEmailEntityFunc(sdk, data)
}


// Segment returns a Segment entity bound to this client.
// Idiomatic usage: client.Segment(nil).List(nil, nil) or
// client.Segment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Segment(data map[string]any) ResendEntity {
	return NewSegmentEntityFunc(sdk, data)
}


// SegmentContact returns a SegmentContact entity bound to this client.
// Idiomatic usage: client.SegmentContact(nil).List(nil, nil) or
// client.SegmentContact(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) SegmentContact(data map[string]any) ResendEntity {
	return NewSegmentContactEntityFunc(sdk, data)
}


// Suppression returns a Suppression entity bound to this client.
// Idiomatic usage: client.Suppression(nil).List(nil, nil) or
// client.Suppression(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Suppression(data map[string]any) ResendEntity {
	return NewSuppressionEntityFunc(sdk, data)
}


// SuppressionBatchAdd returns a SuppressionBatchAdd entity bound to this client.
// Idiomatic usage: client.SuppressionBatchAdd(nil).List(nil, nil) or
// client.SuppressionBatchAdd(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) SuppressionBatchAdd(data map[string]any) ResendEntity {
	return NewSuppressionBatchAddEntityFunc(sdk, data)
}


// SuppressionBatchRemove returns a SuppressionBatchRemove entity bound to this client.
// Idiomatic usage: client.SuppressionBatchRemove(nil).List(nil, nil) or
// client.SuppressionBatchRemove(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) SuppressionBatchRemove(data map[string]any) ResendEntity {
	return NewSuppressionBatchRemoveEntityFunc(sdk, data)
}


// Template returns a Template entity bound to this client.
// Idiomatic usage: client.Template(nil).List(nil, nil) or
// client.Template(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Template(data map[string]any) ResendEntity {
	return NewTemplateEntityFunc(sdk, data)
}


// Topic returns a Topic entity bound to this client.
// Idiomatic usage: client.Topic(nil).List(nil, nil) or
// client.Topic(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Topic(data map[string]any) ResendEntity {
	return NewTopicEntityFunc(sdk, data)
}


// Usage returns a Usage entity bound to this client.
// Idiomatic usage: client.Usage(nil).List(nil, nil) or
// client.Usage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Usage(data map[string]any) ResendEntity {
	return NewUsageEntityFunc(sdk, data)
}


// Webhook returns a Webhook entity bound to this client.
// Idiomatic usage: client.Webhook(nil).List(nil, nil) or
// client.Webhook(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Webhook(data map[string]any) ResendEntity {
	return NewWebhookEntityFunc(sdk, data)
}


// WebhookEvent returns a WebhookEvent entity bound to this client.
// Idiomatic usage: client.WebhookEvent(nil).List(nil, nil) or
// client.WebhookEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) WebhookEvent(data map[string]any) ResendEntity {
	return NewWebhookEventEntityFunc(sdk, data)
}


// WebhookEventAttempt returns a WebhookEventAttempt entity bound to this client.
// Idiomatic usage: client.WebhookEventAttempt(nil).List(nil, nil) or
// client.WebhookEventAttempt(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) WebhookEventAttempt(data map[string]any) ResendEntity {
	return NewWebhookEventAttemptEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *ResendSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewResendSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
