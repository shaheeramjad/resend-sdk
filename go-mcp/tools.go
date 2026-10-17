package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	sdk "github.com/shaheeramjad/resend-sdk/go"
)

// ListArgs is what an agent sends to resend_list.
type ListArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: api_key | attachment | audience | automation | automation_run | broadcast | broadcast_clicked_link | broadcast_recipient | contact | contact_import | contact_property | contact_segment | contact_topic | domain | email | emails_metric | event | inbox | inbox_draft | inbox_label | inbox_thread | inbox_thread_email | log | oauth_grant | received_attachment | received_email | segment | segment_contact | suppression | template | topic | webhook | webhook_event | webhook_event_attempt"`
	Query  map[string]any `json:"query,omitempty" jsonschema:"optional filter map; omit it for the first page"`
}

// LoadArgs is what an agent sends to resend_load.
type LoadArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: attachment | audience | automation | automation_run | broadcast | contact | contact_import | contact_property | domain | domain_claim | email | event | inbox | inbox_agent | inbox_draft | inbox_thread | inbox_thread_email | log | received_attachment | received_email | segment | suppression | template | topic | usage | webhook | webhook_event"`
	Query  map[string]any `json:"query" jsonschema:"match map naming the record, such as {\"id\":1}"`
}

func registerTools(server *mcp.Server, client *sdk.ResendSDK) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "resend_list",
		Description: "List records from Resend. Args: entity, query (optional filter map; omit it for the first page). Returns the first page of records as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[ListArgs]("api_key", "attachment", "audience", "automation", "automation_run", "broadcast", "broadcast_clicked_link", "broadcast_recipient", "contact", "contact_import", "contact_property", "contact_segment", "contact_topic", "domain", "email", "emails_metric", "event", "inbox", "inbox_draft", "inbox_label", "inbox_thread", "inbox_thread_email", "log", "oauth_grant", "received_attachment", "received_email", "segment", "segment_contact", "suppression", "template", "topic", "webhook", "webhook_event", "webhook_event_attempt"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args ListArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "list", args.Entity, args.Query)
	})
	mcp.AddTool(server, &mcp.Tool{
		Name:        "resend_load",
		Description: "Load one record from Resend. Args: entity, query (match map naming the record, such as {\"id\":1}). Returns the record as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[LoadArgs]("attachment", "audience", "automation", "automation_run", "broadcast", "contact", "contact_import", "contact_property", "domain", "domain_claim", "email", "event", "inbox", "inbox_agent", "inbox_draft", "inbox_thread", "inbox_thread_email", "log", "received_attachment", "received_email", "segment", "suppression", "template", "topic", "usage", "webhook", "webhook_event"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args LoadArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "load", args.Entity, args.Query)
	})
}

// entitySchema is the schema inferred from In, its entity limited to the
// entities the tool serves.
func entitySchema[In any](names ...string) *jsonschema.Schema {
	schema, err := jsonschema.For[In](nil)
	if err != nil {
		panic(err)
	}
	enum := make([]any, len(names))
	for i, name := range names {
		enum[i] = name
	}
	schema.Properties["entity"].Enum = enum
	return schema
}

func runOp(_ context.Context, client *sdk.ResendSDK, op string, entity string, input map[string]any) (*mcp.CallToolResult, any, error) {
	ent, err := entityFor(client, entity)
	if err != nil {
		return toolError(err.Error())
	}

	var result any
	switch op {
	case "list":
		result, err = ent.List(input, nil)
	case "load":
		result, err = ent.Load(input, nil)
	case "create":
		result, err = ent.Create(input, nil)
	case "update":
		result, err = ent.Update(input, nil)
	case "patch":
		result, err = ent.Patch(input, nil)
	case "remove":
		result, err = ent.Remove(input, nil)
	default:
		return toolError(fmt.Sprintf("unknown op %q", op))
	}
	if err != nil {
		return toolError(err.Error())
	}

	// SDK returns *Entity wrappers; unwrap each via .Data() to get a
	// plain map[string]any (or []any of maps for list) suitable for
	// JSON marshalling.
	data := extractData(result)
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return toolError(fmt.Sprintf("marshal: %v", err))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(body)},
		},
	}, data, nil
}

// entityFor dispatches on the lowercase entity name. The generator
// emits one `case "<name>":` per entity defined in the SDK model.
func entityFor(client *sdk.ResendSDK, name string) (sdk.ResendEntity, error) {
	switch strings.ToLower(name) {
	case "api_key":
		return client.ApiKey(nil), nil
	case "attachment":
		return client.Attachment(nil), nil
	case "audience":
		return client.Audience(nil), nil
	case "automation":
		return client.Automation(nil), nil
	case "automation_run":
		return client.AutomationRun(nil), nil
	case "broadcast":
		return client.Broadcast(nil), nil
	case "broadcast_clicked_link":
		return client.BroadcastClickedLink(nil), nil
	case "broadcast_recipient":
		return client.BroadcastRecipient(nil), nil
	case "contact":
		return client.Contact(nil), nil
	case "contact_import":
		return client.ContactImport(nil), nil
	case "contact_property":
		return client.ContactProperty(nil), nil
	case "contact_segment":
		return client.ContactSegment(nil), nil
	case "contact_topic":
		return client.ContactTopic(nil), nil
	case "domain":
		return client.Domain(nil), nil
	case "domain_claim":
		return client.DomainClaim(nil), nil
	case "email":
		return client.Email(nil), nil
	case "email_batch":
		return client.EmailBatch(nil), nil
	case "emails_metric":
		return client.EmailsMetric(nil), nil
	case "event":
		return client.Event(nil), nil
	case "inbox":
		return client.Inbox(nil), nil
	case "inbox_agent":
		return client.InboxAgent(nil), nil
	case "inbox_draft":
		return client.InboxDraft(nil), nil
	case "inbox_label":
		return client.InboxLabel(nil), nil
	case "inbox_thread":
		return client.InboxThread(nil), nil
	case "inbox_thread_email":
		return client.InboxThreadEmail(nil), nil
	case "inbox_thread_label":
		return client.InboxThreadLabel(nil), nil
	case "log":
		return client.Log(nil), nil
	case "oauth_grant":
		return client.OauthGrant(nil), nil
	case "received_attachment":
		return client.ReceivedAttachment(nil), nil
	case "received_email":
		return client.ReceivedEmail(nil), nil
	case "segment":
		return client.Segment(nil), nil
	case "segment_contact":
		return client.SegmentContact(nil), nil
	case "suppression":
		return client.Suppression(nil), nil
	case "suppression_batch_add":
		return client.SuppressionBatchAdd(nil), nil
	case "suppression_batch_remove":
		return client.SuppressionBatchRemove(nil), nil
	case "template":
		return client.Template(nil), nil
	case "topic":
		return client.Topic(nil), nil
	case "usage":
		return client.Usage(nil), nil
	case "webhook":
		return client.Webhook(nil), nil
	case "webhook_event":
		return client.WebhookEvent(nil), nil
	case "webhook_event_attempt":
		return client.WebhookEventAttempt(nil), nil
	}
	return nil, fmt.Errorf("unknown entity %q", name)
}

func extractData(x any) any {
	switch v := x.(type) {
	case sdk.Entity:
		return extractData(v.Data())
	case []any:
		out := make([]any, len(v))
		for i, e := range v {
			out[i] = extractData(e)
		}
		return out
	case map[string]any:
		out := make(map[string]any, len(v))
		for k, vv := range v {
			out[k] = extractData(vv)
		}
		return out
	}
	return x
}

func toolError(msg string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}, nil, nil
}

// hint is an MCP annotation that defaults to true unless stated.
func hint(b bool) *bool {
	return &b
}
