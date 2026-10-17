package core

import (
	"encoding/json"
	"fmt"
	"strings"

	vs "github.com/voxgig-sdk/resend-sdk/go/utility/struct"
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


// AddContactToSegmentResponseSuccess returns a AddContactToSegmentResponseSuccess entity bound to this client.
// Idiomatic usage: client.AddContactToSegmentResponseSuccess(nil).List(nil, nil) or
// client.AddContactToSegmentResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) AddContactToSegmentResponseSuccess(data map[string]any) ResendEntity {
	return NewAddContactToSegmentResponseSuccessEntityFunc(sdk, data)
}


// AddInboxThreadLabelsResponseSuccess returns a AddInboxThreadLabelsResponseSuccess entity bound to this client.
// Idiomatic usage: client.AddInboxThreadLabelsResponseSuccess(nil).List(nil, nil) or
// client.AddInboxThreadLabelsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) AddInboxThreadLabelsResponseSuccess(data map[string]any) ResendEntity {
	return NewAddInboxThreadLabelsResponseSuccessEntityFunc(sdk, data)
}


// ApiKey returns a ApiKey entity bound to this client.
// Idiomatic usage: client.ApiKey(nil).List(nil, nil) or
// client.ApiKey(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ApiKey(data map[string]any) ResendEntity {
	return NewApiKeyEntityFunc(sdk, data)
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


// AutomationRunListItem returns a AutomationRunListItem entity bound to this client.
// Idiomatic usage: client.AutomationRunListItem(nil).List(nil, nil) or
// client.AutomationRunListItem(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) AutomationRunListItem(data map[string]any) ResendEntity {
	return NewAutomationRunListItemEntityFunc(sdk, data)
}


// BatchAddSuppressionsResponseSuccess returns a BatchAddSuppressionsResponseSuccess entity bound to this client.
// Idiomatic usage: client.BatchAddSuppressionsResponseSuccess(nil).List(nil, nil) or
// client.BatchAddSuppressionsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) BatchAddSuppressionsResponseSuccess(data map[string]any) ResendEntity {
	return NewBatchAddSuppressionsResponseSuccessEntityFunc(sdk, data)
}


// BatchRemoveSuppressionsResponseSuccess returns a BatchRemoveSuppressionsResponseSuccess entity bound to this client.
// Idiomatic usage: client.BatchRemoveSuppressionsResponseSuccess(nil).List(nil, nil) or
// client.BatchRemoveSuppressionsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) BatchRemoveSuppressionsResponseSuccess(data map[string]any) ResendEntity {
	return NewBatchRemoveSuppressionsResponseSuccessEntityFunc(sdk, data)
}


// Broadcast returns a Broadcast entity bound to this client.
// Idiomatic usage: client.Broadcast(nil).List(nil, nil) or
// client.Broadcast(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Broadcast(data map[string]any) ResendEntity {
	return NewBroadcastEntityFunc(sdk, data)
}


// Contact returns a Contact entity bound to this client.
// Idiomatic usage: client.Contact(nil).List(nil, nil) or
// client.Contact(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Contact(data map[string]any) ResendEntity {
	return NewContactEntityFunc(sdk, data)
}


// ContactImportResponseSuccess returns a ContactImportResponseSuccess entity bound to this client.
// Idiomatic usage: client.ContactImportResponseSuccess(nil).List(nil, nil) or
// client.ContactImportResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactImportResponseSuccess(data map[string]any) ResendEntity {
	return NewContactImportResponseSuccessEntityFunc(sdk, data)
}


// ContactProperty returns a ContactProperty entity bound to this client.
// Idiomatic usage: client.ContactProperty(nil).List(nil, nil) or
// client.ContactProperty(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactProperty(data map[string]any) ResendEntity {
	return NewContactPropertyEntityFunc(sdk, data)
}


// ContactTopicsResponseSuccess returns a ContactTopicsResponseSuccess entity bound to this client.
// Idiomatic usage: client.ContactTopicsResponseSuccess(nil).List(nil, nil) or
// client.ContactTopicsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ContactTopicsResponseSuccess(data map[string]any) ResendEntity {
	return NewContactTopicsResponseSuccessEntityFunc(sdk, data)
}


// CreateApiKey returns a CreateApiKey entity bound to this client.
// Idiomatic usage: client.CreateApiKey(nil).List(nil, nil) or
// client.CreateApiKey(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) CreateApiKey(data map[string]any) ResendEntity {
	return NewCreateApiKeyEntityFunc(sdk, data)
}


// CreateBatchEmail returns a CreateBatchEmail entity bound to this client.
// Idiomatic usage: client.CreateBatchEmail(nil).List(nil, nil) or
// client.CreateBatchEmail(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) CreateBatchEmail(data map[string]any) ResendEntity {
	return NewCreateBatchEmailEntityFunc(sdk, data)
}


// CreateBroadcastResponseSuccess returns a CreateBroadcastResponseSuccess entity bound to this client.
// Idiomatic usage: client.CreateBroadcastResponseSuccess(nil).List(nil, nil) or
// client.CreateBroadcastResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) CreateBroadcastResponseSuccess(data map[string]any) ResendEntity {
	return NewCreateBroadcastResponseSuccessEntityFunc(sdk, data)
}


// CreateInboxDraftResponseSuccess returns a CreateInboxDraftResponseSuccess entity bound to this client.
// Idiomatic usage: client.CreateInboxDraftResponseSuccess(nil).List(nil, nil) or
// client.CreateInboxDraftResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) CreateInboxDraftResponseSuccess(data map[string]any) ResendEntity {
	return NewCreateInboxDraftResponseSuccessEntityFunc(sdk, data)
}


// CreateInboxLabelResponseSuccess returns a CreateInboxLabelResponseSuccess entity bound to this client.
// Idiomatic usage: client.CreateInboxLabelResponseSuccess(nil).List(nil, nil) or
// client.CreateInboxLabelResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) CreateInboxLabelResponseSuccess(data map[string]any) ResendEntity {
	return NewCreateInboxLabelResponseSuccessEntityFunc(sdk, data)
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


// InboxAgentResponseSuccess returns a InboxAgentResponseSuccess entity bound to this client.
// Idiomatic usage: client.InboxAgentResponseSuccess(nil).List(nil, nil) or
// client.InboxAgentResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxAgentResponseSuccess(data map[string]any) ResendEntity {
	return NewInboxAgentResponseSuccessEntityFunc(sdk, data)
}


// InboxDraftListItem returns a InboxDraftListItem entity bound to this client.
// Idiomatic usage: client.InboxDraftListItem(nil).List(nil, nil) or
// client.InboxDraftListItem(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxDraftListItem(data map[string]any) ResendEntity {
	return NewInboxDraftListItemEntityFunc(sdk, data)
}


// InboxDraftResponseSuccess returns a InboxDraftResponseSuccess entity bound to this client.
// Idiomatic usage: client.InboxDraftResponseSuccess(nil).List(nil, nil) or
// client.InboxDraftResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxDraftResponseSuccess(data map[string]any) ResendEntity {
	return NewInboxDraftResponseSuccessEntityFunc(sdk, data)
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


// InboxThreadEmailResponseSuccess returns a InboxThreadEmailResponseSuccess entity bound to this client.
// Idiomatic usage: client.InboxThreadEmailResponseSuccess(nil).List(nil, nil) or
// client.InboxThreadEmailResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxThreadEmailResponseSuccess(data map[string]any) ResendEntity {
	return NewInboxThreadEmailResponseSuccessEntityFunc(sdk, data)
}


// InboxThreadResponseSuccess returns a InboxThreadResponseSuccess entity bound to this client.
// Idiomatic usage: client.InboxThreadResponseSuccess(nil).List(nil, nil) or
// client.InboxThreadResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) InboxThreadResponseSuccess(data map[string]any) ResendEntity {
	return NewInboxThreadResponseSuccessEntityFunc(sdk, data)
}


// ListAttachment returns a ListAttachment entity bound to this client.
// Idiomatic usage: client.ListAttachment(nil).List(nil, nil) or
// client.ListAttachment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListAttachment(data map[string]any) ResendEntity {
	return NewListAttachmentEntityFunc(sdk, data)
}


// ListBroadcastClickedLinksResponseSuccess returns a ListBroadcastClickedLinksResponseSuccess entity bound to this client.
// Idiomatic usage: client.ListBroadcastClickedLinksResponseSuccess(nil).List(nil, nil) or
// client.ListBroadcastClickedLinksResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListBroadcastClickedLinksResponseSuccess(data map[string]any) ResendEntity {
	return NewListBroadcastClickedLinksResponseSuccessEntityFunc(sdk, data)
}


// ListBroadcastRecipientsResponseSuccess returns a ListBroadcastRecipientsResponseSuccess entity bound to this client.
// Idiomatic usage: client.ListBroadcastRecipientsResponseSuccess(nil).List(nil, nil) or
// client.ListBroadcastRecipientsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListBroadcastRecipientsResponseSuccess(data map[string]any) ResendEntity {
	return NewListBroadcastRecipientsResponseSuccessEntityFunc(sdk, data)
}


// ListContactSegmentsResponseSuccess returns a ListContactSegmentsResponseSuccess entity bound to this client.
// Idiomatic usage: client.ListContactSegmentsResponseSuccess(nil).List(nil, nil) or
// client.ListContactSegmentsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListContactSegmentsResponseSuccess(data map[string]any) ResendEntity {
	return NewListContactSegmentsResponseSuccessEntityFunc(sdk, data)
}


// ListContactsResponseSuccess returns a ListContactsResponseSuccess entity bound to this client.
// Idiomatic usage: client.ListContactsResponseSuccess(nil).List(nil, nil) or
// client.ListContactsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListContactsResponseSuccess(data map[string]any) ResendEntity {
	return NewListContactsResponseSuccessEntityFunc(sdk, data)
}


// ListWebhookEvent returns a ListWebhookEvent entity bound to this client.
// Idiomatic usage: client.ListWebhookEvent(nil).List(nil, nil) or
// client.ListWebhookEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListWebhookEvent(data map[string]any) ResendEntity {
	return NewListWebhookEventEntityFunc(sdk, data)
}


// ListWebhookEventAttempt returns a ListWebhookEventAttempt entity bound to this client.
// Idiomatic usage: client.ListWebhookEventAttempt(nil).List(nil, nil) or
// client.ListWebhookEventAttempt(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ListWebhookEventAttempt(data map[string]any) ResendEntity {
	return NewListWebhookEventAttemptEntityFunc(sdk, data)
}


// Log returns a Log entity bound to this client.
// Idiomatic usage: client.Log(nil).List(nil, nil) or
// client.Log(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Log(data map[string]any) ResendEntity {
	return NewLogEntityFunc(sdk, data)
}


// ReceivedEmail returns a ReceivedEmail entity bound to this client.
// Idiomatic usage: client.ReceivedEmail(nil).List(nil, nil) or
// client.ReceivedEmail(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) ReceivedEmail(data map[string]any) ResendEntity {
	return NewReceivedEmailEntityFunc(sdk, data)
}


// RemoveAudienceResponseSuccess returns a RemoveAudienceResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveAudienceResponseSuccess(nil).List(nil, nil) or
// client.RemoveAudienceResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveAudienceResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveAudienceResponseSuccessEntityFunc(sdk, data)
}


// RemoveBroadcastResponseSuccess returns a RemoveBroadcastResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveBroadcastResponseSuccess(nil).List(nil, nil) or
// client.RemoveBroadcastResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveBroadcastResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveBroadcastResponseSuccessEntityFunc(sdk, data)
}


// RemoveContactFromSegmentResponseSuccess returns a RemoveContactFromSegmentResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveContactFromSegmentResponseSuccess(nil).List(nil, nil) or
// client.RemoveContactFromSegmentResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveContactFromSegmentResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveContactFromSegmentResponseSuccessEntityFunc(sdk, data)
}


// RemoveContactPropertyResponseSuccess returns a RemoveContactPropertyResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveContactPropertyResponseSuccess(nil).List(nil, nil) or
// client.RemoveContactPropertyResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveContactPropertyResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveContactPropertyResponseSuccessEntityFunc(sdk, data)
}


// RemoveContactResponseSuccess returns a RemoveContactResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveContactResponseSuccess(nil).List(nil, nil) or
// client.RemoveContactResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveContactResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveContactResponseSuccessEntityFunc(sdk, data)
}


// RemoveEvent returns a RemoveEvent entity bound to this client.
// Idiomatic usage: client.RemoveEvent(nil).List(nil, nil) or
// client.RemoveEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveEvent(data map[string]any) ResendEntity {
	return NewRemoveEventEntityFunc(sdk, data)
}


// RemoveInboxDraftResponseSuccess returns a RemoveInboxDraftResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveInboxDraftResponseSuccess(nil).List(nil, nil) or
// client.RemoveInboxDraftResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveInboxDraftResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveInboxDraftResponseSuccessEntityFunc(sdk, data)
}


// RemoveInboxLabelResponseSuccess returns a RemoveInboxLabelResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveInboxLabelResponseSuccess(nil).List(nil, nil) or
// client.RemoveInboxLabelResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveInboxLabelResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveInboxLabelResponseSuccessEntityFunc(sdk, data)
}


// RemoveInboxResponseSuccess returns a RemoveInboxResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveInboxResponseSuccess(nil).List(nil, nil) or
// client.RemoveInboxResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveInboxResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveInboxResponseSuccessEntityFunc(sdk, data)
}


// RemoveInboxThreadLabelResponseSuccess returns a RemoveInboxThreadLabelResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveInboxThreadLabelResponseSuccess(nil).List(nil, nil) or
// client.RemoveInboxThreadLabelResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveInboxThreadLabelResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveInboxThreadLabelResponseSuccessEntityFunc(sdk, data)
}


// RemoveInboxThreadResponseSuccess returns a RemoveInboxThreadResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveInboxThreadResponseSuccess(nil).List(nil, nil) or
// client.RemoveInboxThreadResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveInboxThreadResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveInboxThreadResponseSuccessEntityFunc(sdk, data)
}


// RemoveSegmentResponseSuccess returns a RemoveSegmentResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveSegmentResponseSuccess(nil).List(nil, nil) or
// client.RemoveSegmentResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveSegmentResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveSegmentResponseSuccessEntityFunc(sdk, data)
}


// RemoveSuppressionResponseSuccess returns a RemoveSuppressionResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveSuppressionResponseSuccess(nil).List(nil, nil) or
// client.RemoveSuppressionResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveSuppressionResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveSuppressionResponseSuccessEntityFunc(sdk, data)
}


// RemoveTemplateResponseSuccess returns a RemoveTemplateResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveTemplateResponseSuccess(nil).List(nil, nil) or
// client.RemoveTemplateResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveTemplateResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveTemplateResponseSuccessEntityFunc(sdk, data)
}


// RemoveTopicResponseSuccess returns a RemoveTopicResponseSuccess entity bound to this client.
// Idiomatic usage: client.RemoveTopicResponseSuccess(nil).List(nil, nil) or
// client.RemoveTopicResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RemoveTopicResponseSuccess(data map[string]any) ResendEntity {
	return NewRemoveTopicResponseSuccessEntityFunc(sdk, data)
}


// RetrievedAttachment returns a RetrievedAttachment entity bound to this client.
// Idiomatic usage: client.RetrievedAttachment(nil).List(nil, nil) or
// client.RetrievedAttachment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RetrievedAttachment(data map[string]any) ResendEntity {
	return NewRetrievedAttachmentEntityFunc(sdk, data)
}


// RevokeOAuthGrant returns a RevokeOAuthGrant entity bound to this client.
// Idiomatic usage: client.RevokeOAuthGrant(nil).List(nil, nil) or
// client.RevokeOAuthGrant(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) RevokeOAuthGrant(data map[string]any) ResendEntity {
	return NewRevokeOAuthGrantEntityFunc(sdk, data)
}


// Rotate returns a Rotate entity bound to this client.
// Idiomatic usage: client.Rotate(nil).List(nil, nil) or
// client.Rotate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Rotate(data map[string]any) ResendEntity {
	return NewRotateEntityFunc(sdk, data)
}


// Segment returns a Segment entity bound to this client.
// Idiomatic usage: client.Segment(nil).List(nil, nil) or
// client.Segment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Segment(data map[string]any) ResendEntity {
	return NewSegmentEntityFunc(sdk, data)
}


// SetInboxThreadLabelsResponseSuccess returns a SetInboxThreadLabelsResponseSuccess entity bound to this client.
// Idiomatic usage: client.SetInboxThreadLabelsResponseSuccess(nil).List(nil, nil) or
// client.SetInboxThreadLabelsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) SetInboxThreadLabelsResponseSuccess(data map[string]any) ResendEntity {
	return NewSetInboxThreadLabelsResponseSuccessEntityFunc(sdk, data)
}


// Suppression returns a Suppression entity bound to this client.
// Idiomatic usage: client.Suppression(nil).List(nil, nil) or
// client.Suppression(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) Suppression(data map[string]any) ResendEntity {
	return NewSuppressionEntityFunc(sdk, data)
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


// UpdateApiKey returns a UpdateApiKey entity bound to this client.
// Idiomatic usage: client.UpdateApiKey(nil).List(nil, nil) or
// client.UpdateApiKey(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateApiKey(data map[string]any) ResendEntity {
	return NewUpdateApiKeyEntityFunc(sdk, data)
}


// UpdateBroadcastResponseSuccess returns a UpdateBroadcastResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateBroadcastResponseSuccess(nil).List(nil, nil) or
// client.UpdateBroadcastResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateBroadcastResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateBroadcastResponseSuccessEntityFunc(sdk, data)
}


// UpdateContactPropertyResponseSuccess returns a UpdateContactPropertyResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateContactPropertyResponseSuccess(nil).List(nil, nil) or
// client.UpdateContactPropertyResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateContactPropertyResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateContactPropertyResponseSuccessEntityFunc(sdk, data)
}


// UpdateContactResponseSuccess returns a UpdateContactResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateContactResponseSuccess(nil).List(nil, nil) or
// client.UpdateContactResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateContactResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateContactResponseSuccessEntityFunc(sdk, data)
}


// UpdateContactTopicsResponseSuccess returns a UpdateContactTopicsResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateContactTopicsResponseSuccess(nil).List(nil, nil) or
// client.UpdateContactTopicsResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateContactTopicsResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateContactTopicsResponseSuccessEntityFunc(sdk, data)
}


// UpdateDomainResponseSuccess returns a UpdateDomainResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateDomainResponseSuccess(nil).List(nil, nil) or
// client.UpdateDomainResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateDomainResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateDomainResponseSuccessEntityFunc(sdk, data)
}


// UpdateEmailOption returns a UpdateEmailOption entity bound to this client.
// Idiomatic usage: client.UpdateEmailOption(nil).List(nil, nil) or
// client.UpdateEmailOption(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateEmailOption(data map[string]any) ResendEntity {
	return NewUpdateEmailOptionEntityFunc(sdk, data)
}


// UpdateEvent returns a UpdateEvent entity bound to this client.
// Idiomatic usage: client.UpdateEvent(nil).List(nil, nil) or
// client.UpdateEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateEvent(data map[string]any) ResendEntity {
	return NewUpdateEventEntityFunc(sdk, data)
}


// UpdateInboxDraftResponseSuccess returns a UpdateInboxDraftResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateInboxDraftResponseSuccess(nil).List(nil, nil) or
// client.UpdateInboxDraftResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateInboxDraftResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateInboxDraftResponseSuccessEntityFunc(sdk, data)
}


// UpdateInboxLabelResponseSuccess returns a UpdateInboxLabelResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateInboxLabelResponseSuccess(nil).List(nil, nil) or
// client.UpdateInboxLabelResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateInboxLabelResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateInboxLabelResponseSuccessEntityFunc(sdk, data)
}


// UpdateInboxResponseSuccess returns a UpdateInboxResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateInboxResponseSuccess(nil).List(nil, nil) or
// client.UpdateInboxResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateInboxResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateInboxResponseSuccessEntityFunc(sdk, data)
}


// UpdateInboxThreadResponseSuccess returns a UpdateInboxThreadResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateInboxThreadResponseSuccess(nil).List(nil, nil) or
// client.UpdateInboxThreadResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateInboxThreadResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateInboxThreadResponseSuccessEntityFunc(sdk, data)
}


// UpdateSegmentResponseSuccess returns a UpdateSegmentResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateSegmentResponseSuccess(nil).List(nil, nil) or
// client.UpdateSegmentResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateSegmentResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateSegmentResponseSuccessEntityFunc(sdk, data)
}


// UpdateTemplateResponseSuccess returns a UpdateTemplateResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateTemplateResponseSuccess(nil).List(nil, nil) or
// client.UpdateTemplateResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateTemplateResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateTemplateResponseSuccessEntityFunc(sdk, data)
}


// UpdateTopicResponseSuccess returns a UpdateTopicResponseSuccess entity bound to this client.
// Idiomatic usage: client.UpdateTopicResponseSuccess(nil).List(nil, nil) or
// client.UpdateTopicResponseSuccess(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateTopicResponseSuccess(data map[string]any) ResendEntity {
	return NewUpdateTopicResponseSuccessEntityFunc(sdk, data)
}


// UpdateWebhook returns a UpdateWebhook entity bound to this client.
// Idiomatic usage: client.UpdateWebhook(nil).List(nil, nil) or
// client.UpdateWebhook(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *ResendSDK) UpdateWebhook(data map[string]any) ResendEntity {
	return NewUpdateWebhookEntityFunc(sdk, data)
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
