package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	sdk "github.com/voxgig-sdk/resend-sdk/go"
)

// ListArgs is what an agent sends to resend_list.
type ListArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: api_key | audience | automation | automation_run_list_item | broadcast | contact | contact_import_response_success | contact_property | contact_topics_response_success | domain | email | emails_metric | event | inbox | inbox_agent_response_success | inbox_draft_list_item | inbox_label | inbox_thread | inbox_thread_email | list_attachment | list_broadcast_clicked_links_response_success | list_broadcast_recipients_response_success | list_contact_segments_response_success | list_contacts_response_success | list_webhook_event | list_webhook_event_attempt | log | received_email | remove_segment_response_success | remove_suppression_response_success | remove_template_response_success | remove_topic_response_success | revoke_o_auth_grant | update_webhook"`
	Query  map[string]any `json:"query,omitempty" jsonschema:"optional filter map; omit it for the first page"`
}

// LoadArgs is what an agent sends to resend_load.
type LoadArgs struct {
	Entity string         `json:"entity" jsonschema:"one of: audience | automation | automation_run | broadcast | contact | contact_import_response_success | contact_property | domain | domain_claim | email | event | inbox | inbox_draft_response_success | inbox_thread_email_response_success | inbox_thread_response_success | log | received_email | retrieved_attachment | segment | suppression | template | topic | usage | webhook | webhook_event"`
	Query  map[string]any `json:"query" jsonschema:"match map naming the record, such as {\"id\":1}"`
}

func registerTools(server *mcp.Server, client *sdk.ResendSDK) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "resend_list",
		Description: "List records from Resend. Args: entity, query (optional filter map; omit it for the first page). Returns the first page of records as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[ListArgs]("api_key", "audience", "automation", "automation_run_list_item", "broadcast", "contact", "contact_import_response_success", "contact_property", "contact_topics_response_success", "domain", "email", "emails_metric", "event", "inbox", "inbox_agent_response_success", "inbox_draft_list_item", "inbox_label", "inbox_thread", "inbox_thread_email", "list_attachment", "list_broadcast_clicked_links_response_success", "list_broadcast_recipients_response_success", "list_contact_segments_response_success", "list_contacts_response_success", "list_webhook_event", "list_webhook_event_attempt", "log", "received_email", "remove_segment_response_success", "remove_suppression_response_success", "remove_template_response_success", "remove_topic_response_success", "revoke_o_auth_grant", "update_webhook"),
	}, func(ctx context.Context, req *mcp.CallToolRequest, args ListArgs) (*mcp.CallToolResult, any, error) {
		return runOp(ctx, client, "list", args.Entity, args.Query)
	})
	mcp.AddTool(server, &mcp.Tool{
		Name:        "resend_load",
		Description: "Load one record from Resend. Args: entity, query (match map naming the record, such as {\"id\":1}). Returns the record as JSON.",
		Annotations: &mcp.ToolAnnotations{ReadOnlyHint: true},
		InputSchema: entitySchema[LoadArgs]("audience", "automation", "automation_run", "broadcast", "contact", "contact_import_response_success", "contact_property", "domain", "domain_claim", "email", "event", "inbox", "inbox_draft_response_success", "inbox_thread_email_response_success", "inbox_thread_response_success", "log", "received_email", "retrieved_attachment", "segment", "suppression", "template", "topic", "usage", "webhook", "webhook_event"),
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
	case "add_contact_to_segment_response_success":
		return client.AddContactToSegmentResponseSuccess(nil), nil
	case "add_inbox_thread_labels_response_success":
		return client.AddInboxThreadLabelsResponseSuccess(nil), nil
	case "api_key":
		return client.ApiKey(nil), nil
	case "audience":
		return client.Audience(nil), nil
	case "automation":
		return client.Automation(nil), nil
	case "automation_run":
		return client.AutomationRun(nil), nil
	case "automation_run_list_item":
		return client.AutomationRunListItem(nil), nil
	case "batch_add_suppressions_response_success":
		return client.BatchAddSuppressionsResponseSuccess(nil), nil
	case "batch_remove_suppressions_response_success":
		return client.BatchRemoveSuppressionsResponseSuccess(nil), nil
	case "broadcast":
		return client.Broadcast(nil), nil
	case "contact":
		return client.Contact(nil), nil
	case "contact_import_response_success":
		return client.ContactImportResponseSuccess(nil), nil
	case "contact_property":
		return client.ContactProperty(nil), nil
	case "contact_topics_response_success":
		return client.ContactTopicsResponseSuccess(nil), nil
	case "create_api_key":
		return client.CreateApiKey(nil), nil
	case "create_batch_email":
		return client.CreateBatchEmail(nil), nil
	case "create_broadcast_response_success":
		return client.CreateBroadcastResponseSuccess(nil), nil
	case "create_inbox_draft_response_success":
		return client.CreateInboxDraftResponseSuccess(nil), nil
	case "create_inbox_label_response_success":
		return client.CreateInboxLabelResponseSuccess(nil), nil
	case "domain":
		return client.Domain(nil), nil
	case "domain_claim":
		return client.DomainClaim(nil), nil
	case "email":
		return client.Email(nil), nil
	case "emails_metric":
		return client.EmailsMetric(nil), nil
	case "event":
		return client.Event(nil), nil
	case "inbox":
		return client.Inbox(nil), nil
	case "inbox_agent_response_success":
		return client.InboxAgentResponseSuccess(nil), nil
	case "inbox_draft_list_item":
		return client.InboxDraftListItem(nil), nil
	case "inbox_draft_response_success":
		return client.InboxDraftResponseSuccess(nil), nil
	case "inbox_label":
		return client.InboxLabel(nil), nil
	case "inbox_thread":
		return client.InboxThread(nil), nil
	case "inbox_thread_email":
		return client.InboxThreadEmail(nil), nil
	case "inbox_thread_email_response_success":
		return client.InboxThreadEmailResponseSuccess(nil), nil
	case "inbox_thread_response_success":
		return client.InboxThreadResponseSuccess(nil), nil
	case "list_attachment":
		return client.ListAttachment(nil), nil
	case "list_broadcast_clicked_links_response_success":
		return client.ListBroadcastClickedLinksResponseSuccess(nil), nil
	case "list_broadcast_recipients_response_success":
		return client.ListBroadcastRecipientsResponseSuccess(nil), nil
	case "list_contact_segments_response_success":
		return client.ListContactSegmentsResponseSuccess(nil), nil
	case "list_contacts_response_success":
		return client.ListContactsResponseSuccess(nil), nil
	case "list_webhook_event":
		return client.ListWebhookEvent(nil), nil
	case "list_webhook_event_attempt":
		return client.ListWebhookEventAttempt(nil), nil
	case "log":
		return client.Log(nil), nil
	case "received_email":
		return client.ReceivedEmail(nil), nil
	case "remove_audience_response_success":
		return client.RemoveAudienceResponseSuccess(nil), nil
	case "remove_broadcast_response_success":
		return client.RemoveBroadcastResponseSuccess(nil), nil
	case "remove_contact_from_segment_response_success":
		return client.RemoveContactFromSegmentResponseSuccess(nil), nil
	case "remove_contact_property_response_success":
		return client.RemoveContactPropertyResponseSuccess(nil), nil
	case "remove_contact_response_success":
		return client.RemoveContactResponseSuccess(nil), nil
	case "remove_event":
		return client.RemoveEvent(nil), nil
	case "remove_inbox_draft_response_success":
		return client.RemoveInboxDraftResponseSuccess(nil), nil
	case "remove_inbox_label_response_success":
		return client.RemoveInboxLabelResponseSuccess(nil), nil
	case "remove_inbox_response_success":
		return client.RemoveInboxResponseSuccess(nil), nil
	case "remove_inbox_thread_label_response_success":
		return client.RemoveInboxThreadLabelResponseSuccess(nil), nil
	case "remove_inbox_thread_response_success":
		return client.RemoveInboxThreadResponseSuccess(nil), nil
	case "remove_segment_response_success":
		return client.RemoveSegmentResponseSuccess(nil), nil
	case "remove_suppression_response_success":
		return client.RemoveSuppressionResponseSuccess(nil), nil
	case "remove_template_response_success":
		return client.RemoveTemplateResponseSuccess(nil), nil
	case "remove_topic_response_success":
		return client.RemoveTopicResponseSuccess(nil), nil
	case "retrieved_attachment":
		return client.RetrievedAttachment(nil), nil
	case "revoke_o_auth_grant":
		return client.RevokeOAuthGrant(nil), nil
	case "rotate":
		return client.Rotate(nil), nil
	case "segment":
		return client.Segment(nil), nil
	case "set_inbox_thread_labels_response_success":
		return client.SetInboxThreadLabelsResponseSuccess(nil), nil
	case "suppression":
		return client.Suppression(nil), nil
	case "template":
		return client.Template(nil), nil
	case "topic":
		return client.Topic(nil), nil
	case "update_api_key":
		return client.UpdateApiKey(nil), nil
	case "update_broadcast_response_success":
		return client.UpdateBroadcastResponseSuccess(nil), nil
	case "update_contact_property_response_success":
		return client.UpdateContactPropertyResponseSuccess(nil), nil
	case "update_contact_response_success":
		return client.UpdateContactResponseSuccess(nil), nil
	case "update_contact_topics_response_success":
		return client.UpdateContactTopicsResponseSuccess(nil), nil
	case "update_domain_response_success":
		return client.UpdateDomainResponseSuccess(nil), nil
	case "update_email_option":
		return client.UpdateEmailOption(nil), nil
	case "update_event":
		return client.UpdateEvent(nil), nil
	case "update_inbox_draft_response_success":
		return client.UpdateInboxDraftResponseSuccess(nil), nil
	case "update_inbox_label_response_success":
		return client.UpdateInboxLabelResponseSuccess(nil), nil
	case "update_inbox_response_success":
		return client.UpdateInboxResponseSuccess(nil), nil
	case "update_inbox_thread_response_success":
		return client.UpdateInboxThreadResponseSuccess(nil), nil
	case "update_segment_response_success":
		return client.UpdateSegmentResponseSuccess(nil), nil
	case "update_template_response_success":
		return client.UpdateTemplateResponseSuccess(nil), nil
	case "update_topic_response_success":
		return client.UpdateTopicResponseSuccess(nil), nil
	case "update_webhook":
		return client.UpdateWebhook(nil), nil
	case "usage":
		return client.Usage(nil), nil
	case "webhook":
		return client.Webhook(nil), nil
	case "webhook_event":
		return client.WebhookEvent(nil), nil
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
