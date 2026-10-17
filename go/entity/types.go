// Typed models for the Resend SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/resend-sdk/go/core"
)

// AddContactToSegmentResponseSuccess is the typed data model for the add_contact_to_segment_response_success entity.
type AddContactToSegmentResponseSuccess struct {
}

// AddContactToSegmentResponseSuccessCreateData is the typed request payload for AddContactToSegmentResponseSuccess.CreateTyped.
type AddContactToSegmentResponseSuccessCreateData struct {
	ContactId string `json:"contact_id"`
	SegmentId string `json:"segment_id"`
	Object *string `json:"object,omitempty"`
}

// AddInboxThreadLabelsResponseSuccess is the typed data model for the add_inbox_thread_labels_response_success entity.
type AddInboxThreadLabelsResponseSuccess struct {
}

// AddInboxThreadLabelsResponseSuccessCreateData is the typed request payload for AddInboxThreadLabelsResponseSuccess.CreateTyped.
type AddInboxThreadLabelsResponseSuccessCreateData struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	Labels []any `json:"labels"`
}

// ApiKey is the typed data model for the api_key entity.
type ApiKey struct {
}

// ApiKeyListMatch is the typed request payload for ApiKey.ListTyped.
type ApiKeyListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ApiKeyRemoveMatch is the typed request payload for ApiKey.RemoveTyped.
type ApiKeyRemoveMatch struct {
	Id string `json:"id"`
}

// Audience is the typed data model for the audience entity.
type Audience struct {
}

// AudienceLoadMatch is the typed request payload for Audience.LoadTyped.
type AudienceLoadMatch struct {
	Id string `json:"id"`
}

// AudienceListMatch is the typed request payload for Audience.ListTyped.
type AudienceListMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
}

// AudienceCreateData is the typed request payload for Audience.CreateTyped.
type AudienceCreateData struct {
	Name string `json:"name"`
}

// Automation is the typed data model for the automation entity.
type Automation struct {
}

// AutomationLoadMatch is the typed request payload for Automation.LoadTyped.
type AutomationLoadMatch struct {
	Id string `json:"id"`
}

// AutomationListMatch is the typed request payload for Automation.ListTyped.
type AutomationListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// AutomationCreateData is the typed request payload for Automation.CreateTyped.
type AutomationCreateData struct {
	Connections []any `json:"connections"`
	Name string `json:"name"`
	Status *string `json:"status,omitempty"`
	Steps []any `json:"steps"`
}

// AutomationUpdateData is the typed request payload for Automation.UpdateTyped.
type AutomationUpdateData struct {
	Id string `json:"id"`
	Connections *[]any `json:"connections,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Steps *[]any `json:"steps,omitempty"`
}

// AutomationRemoveMatch is the typed request payload for Automation.RemoveTyped.
type AutomationRemoveMatch struct {
	Id string `json:"id"`
}

// AutomationRun is the typed data model for the automation_run entity.
type AutomationRun struct {
}

// AutomationRunLoadMatch is the typed request payload for AutomationRun.LoadTyped.
type AutomationRunLoadMatch struct {
	AutomationId string `json:"automation_id"`
	Id string `json:"id"`
}

// AutomationRunListItem is the typed data model for the automation_run_list_item entity.
type AutomationRunListItem struct {
}

// AutomationRunListItemListMatch is the typed request payload for AutomationRunListItem.ListTyped.
type AutomationRunListItemListMatch struct {
	Id string `json:"id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// BatchAddSuppressionsResponseSuccess is the typed data model for the batch_add_suppressions_response_success entity.
type BatchAddSuppressionsResponseSuccess struct {
}

// BatchAddSuppressionsResponseSuccessCreateData is the typed request payload for BatchAddSuppressionsResponseSuccess.CreateTyped.
type BatchAddSuppressionsResponseSuccessCreateData struct {
	Emails []any `json:"emails"`
}

// BatchRemoveSuppressionsResponseSuccess is the typed data model for the batch_remove_suppressions_response_success entity.
type BatchRemoveSuppressionsResponseSuccess struct {
}

// BatchRemoveSuppressionsResponseSuccessCreateData is the typed request payload for BatchRemoveSuppressionsResponseSuccess.CreateTyped.
type BatchRemoveSuppressionsResponseSuccessCreateData struct {
	Emails *[]any `json:"emails,omitempty"`
	Ids *[]any `json:"ids,omitempty"`
}

// Broadcast is the typed data model for the broadcast entity.
type Broadcast struct {
}

// BroadcastLoadMatch is the typed request payload for Broadcast.LoadTyped.
type BroadcastLoadMatch struct {
	Id string `json:"id"`
}

// BroadcastListMatch is the typed request payload for Broadcast.ListTyped.
type BroadcastListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// BroadcastCreateData is the typed request payload for Broadcast.CreateTyped.
type BroadcastCreateData struct {
	Id string `json:"id"`
	AudienceId *any `json:"audience_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	From *string `json:"from,omitempty"`
	Html *any `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	PreviewText *string `json:"preview_text,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	SegmentId *any `json:"segment_id,omitempty"`
	SentAt *string `json:"sent_at,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *any `json:"text,omitempty"`
	TopicId *any `json:"topic_id,omitempty"`
}

// Contact is the typed data model for the contact entity.
type Contact struct {
}

// ContactLoadMatch is the typed request payload for Contact.LoadTyped.
type ContactLoadMatch struct {
	Id string `json:"id"`
}

// ContactListMatch is the typed request payload for Contact.ListTyped.
type ContactListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ContactCreateData is the typed request payload for Contact.CreateTyped.
type ContactCreateData struct {
	AudienceId *string `json:"audience_id,omitempty"`
	Email string `json:"email"`
	FirstName *any `json:"first_name,omitempty"`
	LastName *any `json:"last_name,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
	Segments *[]any `json:"segments,omitempty"`
	Topics *[]any `json:"topics,omitempty"`
	Unsubscribed *bool `json:"unsubscribed,omitempty"`
}

// ContactImportResponseSuccess is the typed data model for the contact_import_response_success entity.
type ContactImportResponseSuccess struct {
}

// ContactImportResponseSuccessLoadMatch is the typed request payload for ContactImportResponseSuccess.LoadTyped.
type ContactImportResponseSuccessLoadMatch struct {
	Id string `json:"id"`
}

// ContactImportResponseSuccessListMatch is the typed request payload for ContactImportResponseSuccess.ListTyped.
type ContactImportResponseSuccessListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// ContactImportResponseSuccessCreateData is the typed request payload for ContactImportResponseSuccess.CreateTyped.
type ContactImportResponseSuccessCreateData struct {
	ColumnMap *string `json:"column_map,omitempty"`
	CompletedAt *any `json:"completed_at,omitempty"`
	Counts *map[string]any `json:"counts,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	File string `json:"file"`
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	OnConflict *string `json:"on_conflict,omitempty"`
	Segments *string `json:"segments,omitempty"`
	Status *string `json:"status,omitempty"`
	Topics *string `json:"topics,omitempty"`
}

// ContactProperty is the typed data model for the contact_property entity.
type ContactProperty struct {
}

// ContactPropertyLoadMatch is the typed request payload for ContactProperty.LoadTyped.
type ContactPropertyLoadMatch struct {
	Id string `json:"id"`
}

// ContactPropertyListMatch is the typed request payload for ContactProperty.ListTyped.
type ContactPropertyListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ContactPropertyCreateData is the typed request payload for ContactProperty.CreateTyped.
type ContactPropertyCreateData struct {
	FallbackValue *any `json:"fallback_value,omitempty"`
	Key string `json:"key"`
	Type string `json:"type"`
}

// ContactTopicsResponseSuccess is the typed data model for the contact_topics_response_success entity.
type ContactTopicsResponseSuccess struct {
}

// ContactTopicsResponseSuccessListMatch is the typed request payload for ContactTopicsResponseSuccess.ListTyped.
type ContactTopicsResponseSuccessListMatch struct {
	Id string `json:"id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// CreateApiKey is the typed data model for the create_api_key entity.
type CreateApiKey struct {
}

// CreateApiKeyCreateData is the typed request payload for CreateApiKey.CreateTyped.
type CreateApiKeyCreateData struct {
	DomainId *string `json:"domain_id,omitempty"`
	Name string `json:"name"`
	Permission *string `json:"permission,omitempty"`
}

// CreateBatchEmail is the typed data model for the create_batch_email entity.
type CreateBatchEmail struct {
}

// CreateBatchEmailCreateData is the typed request payload for CreateBatchEmail.CreateTyped.
type CreateBatchEmailCreateData struct {
	SendEmails *[]any `json:"send_emails,omitempty"`
}

// CreateBroadcastResponseSuccess is the typed data model for the create_broadcast_response_success entity.
type CreateBroadcastResponseSuccess struct {
}

// CreateBroadcastResponseSuccessCreateData is the typed request payload for CreateBroadcastResponseSuccess.CreateTyped.
type CreateBroadcastResponseSuccessCreateData struct {
	AudienceId *string `json:"audience_id,omitempty"`
	From string `json:"from"`
	Html *string `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	PreviewText *string `json:"preview_text,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	SegmentId string `json:"segment_id"`
	Send *bool `json:"send,omitempty"`
	Subject string `json:"subject"`
	Text *string `json:"text,omitempty"`
	TopicId *string `json:"topic_id,omitempty"`
}

// CreateInboxDraftResponseSuccess is the typed data model for the create_inbox_draft_response_success entity.
type CreateInboxDraftResponseSuccess struct {
}

// CreateInboxDraftResponseSuccessCreateData is the typed request payload for CreateInboxDraftResponseSuccess.CreateTyped.
type CreateInboxDraftResponseSuccessCreateData struct {
	InboxId string `json:"inbox_id"`
	Attachments *any `json:"attachments,omitempty"`
	Bcc *any `json:"bcc,omitempty"`
	Cc *any `json:"cc,omitempty"`
	Html *any `json:"html,omitempty"`
	ReplyToEmailId *string `json:"reply_to_email_id,omitempty"`
	Subject *any `json:"subject,omitempty"`
	Text *any `json:"text,omitempty"`
	ThreadId *string `json:"thread_id,omitempty"`
	To *any `json:"to,omitempty"`
}

// CreateInboxLabelResponseSuccess is the typed data model for the create_inbox_label_response_success entity.
type CreateInboxLabelResponseSuccess struct {
}

// CreateInboxLabelResponseSuccessCreateData is the typed request payload for CreateInboxLabelResponseSuccess.CreateTyped.
type CreateInboxLabelResponseSuccessCreateData struct {
	InboxId string `json:"inbox_id"`
	Color *string `json:"color,omitempty"`
	Name string `json:"name"`
}

// Domain is the typed data model for the domain entity.
type Domain struct {
}

// DomainLoadMatch is the typed request payload for Domain.LoadTyped.
type DomainLoadMatch struct {
	Id string `json:"id"`
}

// DomainListMatch is the typed request payload for Domain.ListTyped.
type DomainListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// DomainCreateData is the typed request payload for Domain.CreateTyped.
type DomainCreateData struct {
	Capabilities *map[string]any `json:"capabilities,omitempty"`
	ClickTracking *bool `json:"click_tracking,omitempty"`
	CustomReturnPath *string `json:"custom_return_path,omitempty"`
	Name string `json:"name"`
	OpenTracking *bool `json:"open_tracking,omitempty"`
	Region *string `json:"region,omitempty"`
	Tls *string `json:"tls,omitempty"`
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
}

// DomainRemoveMatch is the typed request payload for Domain.RemoveTyped.
type DomainRemoveMatch struct {
	Id string `json:"id"`
}

// DomainClaim is the typed data model for the domain_claim entity.
type DomainClaim struct {
}

// DomainClaimLoadMatch is the typed request payload for DomainClaim.LoadTyped.
type DomainClaimLoadMatch struct {
	Id string `json:"id"`
}

// DomainClaimCreateData is the typed request payload for DomainClaim.CreateTyped.
type DomainClaimCreateData struct {
	BlockedReason *any `json:"blocked_reason,omitempty"`
	ClickTracking *bool `json:"click_tracking,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomReturnPath *string `json:"custom_return_path,omitempty"`
	DomainId *any `json:"domain_id,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	FailureReason *any `json:"failure_reason,omitempty"`
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Object *string `json:"object,omitempty"`
	OpenTracking *bool `json:"open_tracking,omitempty"`
	Record *map[string]any `json:"record,omitempty"`
	Region *any `json:"region,omitempty"`
	Status *string `json:"status,omitempty"`
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
}

// Email is the typed data model for the email entity.
type Email struct {
}

// EmailLoadMatch is the typed request payload for Email.LoadTyped.
type EmailLoadMatch struct {
	Id string `json:"id"`
}

// EmailListMatch is the typed request payload for Email.ListTyped.
type EmailListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// EmailCreateData is the typed request payload for Email.CreateTyped.
type EmailCreateData struct {
	Attachments *[]any `json:"attachments,omitempty"`
	Bcc *[]any `json:"bcc,omitempty"`
	Cc *[]any `json:"cc,omitempty"`
	From string `json:"from"`
	Headers *map[string]any `json:"headers,omitempty"`
	Html *string `json:"html,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	Subject string `json:"subject"`
	Tags *[]any `json:"tags,omitempty"`
	Template *any `json:"template,omitempty"`
	Text *string `json:"text,omitempty"`
	To []any `json:"to"`
	TopicId *string `json:"topic_id,omitempty"`
}

// EmailsMetric is the typed data model for the emails_metric entity.
type EmailsMetric struct {
}

// EmailsMetricListMatch is the typed request payload for EmailsMetric.ListTyped.
type EmailsMetricListMatch struct {
	BroadcastId *[]any `json:"broadcast_id,omitempty"`
	Dimension *[]any `json:"dimension,omitempty"`
	DomainId *[]any `json:"domain_id,omitempty"`
	EmailId *[]any `json:"email_id,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	Granularity *string `json:"granularity,omitempty"`
	Metric *[]any `json:"metric,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// Event is the typed data model for the event entity.
type Event struct {
}

// EventLoadMatch is the typed request payload for Event.LoadTyped.
type EventLoadMatch struct {
	Id string `json:"id"`
}

// EventListMatch is the typed request payload for Event.ListTyped.
type EventListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// EventCreateData is the typed request payload for Event.CreateTyped.
type EventCreateData struct {
	Name string `json:"name"`
	Schema *any `json:"schema,omitempty"`
}

// Inbox is the typed data model for the inbox entity.
type Inbox struct {
}

// InboxLoadMatch is the typed request payload for Inbox.LoadTyped.
type InboxLoadMatch struct {
	Id string `json:"id"`
}

// InboxListMatch is the typed request payload for Inbox.ListTyped.
type InboxListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// InboxCreateData is the typed request payload for Inbox.CreateTyped.
type InboxCreateData struct {
	EmailAddress string `json:"email_address"`
	Forwarding *bool `json:"forwarding,omitempty"`
	FromName *any `json:"from_name,omitempty"`
	Name *string `json:"name,omitempty"`
}

// InboxUpdateData is the typed request payload for Inbox.UpdateTyped.
type InboxUpdateData struct {
	Id string `json:"id"`
	Actions *any `json:"actions,omitempty"`
	Instructions *any `json:"instructions,omitempty"`
	Tone *any `json:"tone,omitempty"`
}

// InboxAgentResponseSuccess is the typed data model for the inbox_agent_response_success entity.
type InboxAgentResponseSuccess struct {
}

// InboxAgentResponseSuccessListMatch is the typed request payload for InboxAgentResponseSuccess.ListTyped.
type InboxAgentResponseSuccessListMatch struct {
	Id string `json:"id"`
}

// InboxDraftListItem is the typed data model for the inbox_draft_list_item entity.
type InboxDraftListItem struct {
}

// InboxDraftListItemListMatch is the typed request payload for InboxDraftListItem.ListTyped.
type InboxDraftListItemListMatch struct {
	Id string `json:"id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// InboxDraftResponseSuccess is the typed data model for the inbox_draft_response_success entity.
type InboxDraftResponseSuccess struct {
}

// InboxDraftResponseSuccessLoadMatch is the typed request payload for InboxDraftResponseSuccess.LoadTyped.
type InboxDraftResponseSuccessLoadMatch struct {
	DraftId string `json:"draft_id"`
	InboxId string `json:"inbox_id"`
}

// InboxDraftResponseSuccessCreateData is the typed request payload for InboxDraftResponseSuccess.CreateTyped.
type InboxDraftResponseSuccessCreateData struct {
	DraftId string `json:"draft_id"`
	InboxId string `json:"inbox_id"`
	Headers *any `json:"headers,omitempty"`
}

// InboxLabel is the typed data model for the inbox_label entity.
type InboxLabel struct {
}

// InboxLabelListMatch is the typed request payload for InboxLabel.ListTyped.
type InboxLabelListMatch struct {
	Id string `json:"id"`
}

// InboxThread is the typed data model for the inbox_thread entity.
type InboxThread struct {
}

// InboxThreadListMatch is the typed request payload for InboxThread.ListTyped.
type InboxThreadListMatch struct {
	Id string `json:"id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Folder *[]any `json:"folder,omitempty"`
	Label *[]any `json:"label,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Read *bool `json:"read,omitempty"`
}

// InboxThreadEmail is the typed data model for the inbox_thread_email entity.
type InboxThreadEmail struct {
}

// InboxThreadEmailListMatch is the typed request payload for InboxThreadEmail.ListTyped.
type InboxThreadEmailListMatch struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// InboxThreadEmailResponseSuccess is the typed data model for the inbox_thread_email_response_success entity.
type InboxThreadEmailResponseSuccess struct {
}

// InboxThreadEmailResponseSuccessLoadMatch is the typed request payload for InboxThreadEmailResponseSuccess.LoadTyped.
type InboxThreadEmailResponseSuccessLoadMatch struct {
	EmailId string `json:"email_id"`
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
}

// InboxThreadEmailResponseSuccessCreateData is the typed request payload for InboxThreadEmailResponseSuccess.CreateTyped.
type InboxThreadEmailResponseSuccessCreateData struct {
	EmailId string `json:"email_id"`
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	Attachments []any `json:"attachments"`
	Bcc []any `json:"bcc"`
	Cc []any `json:"cc"`
	Html any `json:"html"`
	Subject any `json:"subject"`
	Text any `json:"text"`
	To []any `json:"to"`
	Headers *any `json:"headers,omitempty"`
	ReplyAll *any `json:"reply_all,omitempty"`
}

// InboxThreadResponseSuccess is the typed data model for the inbox_thread_response_success entity.
type InboxThreadResponseSuccess struct {
}

// InboxThreadResponseSuccessLoadMatch is the typed request payload for InboxThreadResponseSuccess.LoadTyped.
type InboxThreadResponseSuccessLoadMatch struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
}

// ListAttachment is the typed data model for the list_attachment entity.
type ListAttachment struct {
}

// ListAttachmentListMatch is the typed request payload for ListAttachment.ListTyped.
type ListAttachmentListMatch struct {
	EmailId string `json:"email_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ListBroadcastClickedLinksResponseSuccess is the typed data model for the list_broadcast_clicked_links_response_success entity.
type ListBroadcastClickedLinksResponseSuccess struct {
}

// ListBroadcastClickedLinksResponseSuccessListMatch is the typed request payload for ListBroadcastClickedLinksResponseSuccess.ListTyped.
type ListBroadcastClickedLinksResponseSuccessListMatch struct {
	BroadcastId string `json:"broadcast_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ListBroadcastRecipientsResponseSuccess is the typed data model for the list_broadcast_recipients_response_success entity.
type ListBroadcastRecipientsResponseSuccess struct {
}

// ListBroadcastRecipientsResponseSuccessListMatch is the typed request payload for ListBroadcastRecipientsResponseSuccess.ListTyped.
type ListBroadcastRecipientsResponseSuccessListMatch struct {
	BroadcastId string `json:"broadcast_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	BounceType *string `json:"bounce_type,omitempty"`
	Email *string `json:"email,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Type string `json:"type"`
}

// ListContactSegmentsResponseSuccess is the typed data model for the list_contact_segments_response_success entity.
type ListContactSegmentsResponseSuccess struct {
}

// ListContactSegmentsResponseSuccessListMatch is the typed request payload for ListContactSegmentsResponseSuccess.ListTyped.
type ListContactSegmentsResponseSuccessListMatch struct {
	ContactId string `json:"contact_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ListContactsResponseSuccess is the typed data model for the list_contacts_response_success entity.
type ListContactsResponseSuccess struct {
}

// ListContactsResponseSuccessListMatch is the typed request payload for ListContactsResponseSuccess.ListTyped.
type ListContactsResponseSuccessListMatch struct {
	SegmentId string `json:"segment_id"`
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ListWebhookEvent is the typed data model for the list_webhook_event entity.
type ListWebhookEvent struct {
}

// ListWebhookEventListMatch is the typed request payload for ListWebhookEvent.ListTyped.
type ListWebhookEventListMatch struct {
	WebhookId string `json:"webhook_id"`
	After *string `json:"after,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ListWebhookEventAttempt is the typed data model for the list_webhook_event_attempt entity.
type ListWebhookEventAttempt struct {
}

// ListWebhookEventAttemptListMatch is the typed request payload for ListWebhookEventAttempt.ListTyped.
type ListWebhookEventAttemptListMatch struct {
	EventId string `json:"event_id"`
	WebhookId string `json:"webhook_id"`
	After *string `json:"after,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// Log is the typed data model for the log entity.
type Log struct {
}

// LogLoadMatch is the typed request payload for Log.LoadTyped.
type LogLoadMatch struct {
	Id string `json:"id"`
}

// LogListMatch is the typed request payload for Log.ListTyped.
type LogListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// ReceivedEmail is the typed data model for the received_email entity.
type ReceivedEmail struct {
}

// ReceivedEmailLoadMatch is the typed request payload for ReceivedEmail.LoadTyped.
type ReceivedEmailLoadMatch struct {
	EmailId string `json:"email_id"`
}

// ReceivedEmailListMatch is the typed request payload for ReceivedEmail.ListTyped.
type ReceivedEmailListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// RemoveAudienceResponseSuccess is the typed data model for the remove_audience_response_success entity.
type RemoveAudienceResponseSuccess struct {
}

// RemoveAudienceResponseSuccessRemoveMatch is the typed request payload for RemoveAudienceResponseSuccess.RemoveTyped.
type RemoveAudienceResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveBroadcastResponseSuccess is the typed data model for the remove_broadcast_response_success entity.
type RemoveBroadcastResponseSuccess struct {
}

// RemoveBroadcastResponseSuccessRemoveMatch is the typed request payload for RemoveBroadcastResponseSuccess.RemoveTyped.
type RemoveBroadcastResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveContactFromSegmentResponseSuccess is the typed data model for the remove_contact_from_segment_response_success entity.
type RemoveContactFromSegmentResponseSuccess struct {
}

// RemoveContactFromSegmentResponseSuccessRemoveMatch is the typed request payload for RemoveContactFromSegmentResponseSuccess.RemoveTyped.
type RemoveContactFromSegmentResponseSuccessRemoveMatch struct {
	ContactId string `json:"contact_id"`
	SegmentId string `json:"segment_id"`
}

// RemoveContactPropertyResponseSuccess is the typed data model for the remove_contact_property_response_success entity.
type RemoveContactPropertyResponseSuccess struct {
}

// RemoveContactPropertyResponseSuccessRemoveMatch is the typed request payload for RemoveContactPropertyResponseSuccess.RemoveTyped.
type RemoveContactPropertyResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveContactResponseSuccess is the typed data model for the remove_contact_response_success entity.
type RemoveContactResponseSuccess struct {
}

// RemoveContactResponseSuccessRemoveMatch is the typed request payload for RemoveContactResponseSuccess.RemoveTyped.
type RemoveContactResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveEvent is the typed data model for the remove_event entity.
type RemoveEvent struct {
}

// RemoveEventRemoveMatch is the typed request payload for RemoveEvent.RemoveTyped.
type RemoveEventRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveInboxDraftResponseSuccess is the typed data model for the remove_inbox_draft_response_success entity.
type RemoveInboxDraftResponseSuccess struct {
}

// RemoveInboxDraftResponseSuccessRemoveMatch is the typed request payload for RemoveInboxDraftResponseSuccess.RemoveTyped.
type RemoveInboxDraftResponseSuccessRemoveMatch struct {
	DraftId string `json:"draft_id"`
	InboxId string `json:"inbox_id"`
}

// RemoveInboxLabelResponseSuccess is the typed data model for the remove_inbox_label_response_success entity.
type RemoveInboxLabelResponseSuccess struct {
}

// RemoveInboxLabelResponseSuccessRemoveMatch is the typed request payload for RemoveInboxLabelResponseSuccess.RemoveTyped.
type RemoveInboxLabelResponseSuccessRemoveMatch struct {
	InboxId string `json:"inbox_id"`
	LabelId string `json:"label_id"`
}

// RemoveInboxResponseSuccess is the typed data model for the remove_inbox_response_success entity.
type RemoveInboxResponseSuccess struct {
}

// RemoveInboxResponseSuccessRemoveMatch is the typed request payload for RemoveInboxResponseSuccess.RemoveTyped.
type RemoveInboxResponseSuccessRemoveMatch struct {
	InboxId string `json:"inbox_id"`
}

// RemoveInboxThreadLabelResponseSuccess is the typed data model for the remove_inbox_thread_label_response_success entity.
type RemoveInboxThreadLabelResponseSuccess struct {
}

// RemoveInboxThreadLabelResponseSuccessRemoveMatch is the typed request payload for RemoveInboxThreadLabelResponseSuccess.RemoveTyped.
type RemoveInboxThreadLabelResponseSuccessRemoveMatch struct {
	InboxId string `json:"inbox_id"`
	LabelId string `json:"label_id"`
	ThreadId string `json:"thread_id"`
}

// RemoveInboxThreadResponseSuccess is the typed data model for the remove_inbox_thread_response_success entity.
type RemoveInboxThreadResponseSuccess struct {
}

// RemoveInboxThreadResponseSuccessRemoveMatch is the typed request payload for RemoveInboxThreadResponseSuccess.RemoveTyped.
type RemoveInboxThreadResponseSuccessRemoveMatch struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
}

// RemoveSegmentResponseSuccess is the typed data model for the remove_segment_response_success entity.
type RemoveSegmentResponseSuccess struct {
}

// RemoveSegmentResponseSuccessListMatch is the typed request payload for RemoveSegmentResponseSuccess.ListTyped.
type RemoveSegmentResponseSuccessListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// RemoveSegmentResponseSuccessCreateData is the typed request payload for RemoveSegmentResponseSuccess.CreateTyped.
type RemoveSegmentResponseSuccessCreateData struct {
	AudienceId *string `json:"audience_id,omitempty"`
	Filter *map[string]any `json:"filter,omitempty"`
	Name string `json:"name"`
}

// RemoveSegmentResponseSuccessRemoveMatch is the typed request payload for RemoveSegmentResponseSuccess.RemoveTyped.
type RemoveSegmentResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveSuppressionResponseSuccess is the typed data model for the remove_suppression_response_success entity.
type RemoveSuppressionResponseSuccess struct {
}

// RemoveSuppressionResponseSuccessListMatch is the typed request payload for RemoveSuppressionResponseSuccess.ListTyped.
type RemoveSuppressionResponseSuccessListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Origin *string `json:"origin,omitempty"`
}

// RemoveSuppressionResponseSuccessCreateData is the typed request payload for RemoveSuppressionResponseSuccess.CreateTyped.
type RemoveSuppressionResponseSuccessCreateData struct {
	Email string `json:"email"`
}

// RemoveSuppressionResponseSuccessRemoveMatch is the typed request payload for RemoveSuppressionResponseSuccess.RemoveTyped.
type RemoveSuppressionResponseSuccessRemoveMatch struct {
	Suppression string `json:"suppression"`
}

// RemoveTemplateResponseSuccess is the typed data model for the remove_template_response_success entity.
type RemoveTemplateResponseSuccess struct {
}

// RemoveTemplateResponseSuccessListMatch is the typed request payload for RemoveTemplateResponseSuccess.ListTyped.
type RemoveTemplateResponseSuccessListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// RemoveTemplateResponseSuccessCreateData is the typed request payload for RemoveTemplateResponseSuccess.CreateTyped.
type RemoveTemplateResponseSuccessCreateData struct {
	Alias *string `json:"alias,omitempty"`
	From *string `json:"from,omitempty"`
	Html string `json:"html"`
	Name string `json:"name"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
	Variables *[]any `json:"variables,omitempty"`
}

// RemoveTemplateResponseSuccessRemoveMatch is the typed request payload for RemoveTemplateResponseSuccess.RemoveTyped.
type RemoveTemplateResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RemoveTopicResponseSuccess is the typed data model for the remove_topic_response_success entity.
type RemoveTopicResponseSuccess struct {
}

// RemoveTopicResponseSuccessListMatch is the typed request payload for RemoveTopicResponseSuccess.ListTyped.
type RemoveTopicResponseSuccessListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// RemoveTopicResponseSuccessCreateData is the typed request payload for RemoveTopicResponseSuccess.CreateTyped.
type RemoveTopicResponseSuccessCreateData struct {
	DefaultSubscription string `json:"default_subscription"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
	Visibility *string `json:"visibility,omitempty"`
}

// RemoveTopicResponseSuccessRemoveMatch is the typed request payload for RemoveTopicResponseSuccess.RemoveTyped.
type RemoveTopicResponseSuccessRemoveMatch struct {
	Id string `json:"id"`
}

// RetrievedAttachment is the typed data model for the retrieved_attachment entity.
type RetrievedAttachment struct {
}

// RetrievedAttachmentLoadMatch is the typed request payload for RetrievedAttachment.LoadTyped.
type RetrievedAttachmentLoadMatch struct {
	EmailId *string `json:"email_id,omitempty"`
	Id string `json:"id"`
	ReceivingId *string `json:"receiving_id,omitempty"`
}

// RevokeOAuthGrant is the typed data model for the revoke_o_auth_grant entity.
type RevokeOAuthGrant struct {
}

// RevokeOAuthGrantListMatch is the typed request payload for RevokeOAuthGrant.ListTyped.
type RevokeOAuthGrantListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// RevokeOAuthGrantRemoveMatch is the typed request payload for RevokeOAuthGrant.RemoveTyped.
type RevokeOAuthGrantRemoveMatch struct {
	Id string `json:"id"`
}

// Rotate is the typed data model for the rotate entity.
type Rotate struct {
}

// RotateCreateData is the typed request payload for Rotate.CreateTyped.
type RotateCreateData struct {
	WebhookId string `json:"webhook_id"`
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	SigningSecret *string `json:"signing_secret,omitempty"`
}

// Segment is the typed data model for the segment entity.
type Segment struct {
}

// SegmentLoadMatch is the typed request payload for Segment.LoadTyped.
type SegmentLoadMatch struct {
	Id string `json:"id"`
}

// SetInboxThreadLabelsResponseSuccess is the typed data model for the set_inbox_thread_labels_response_success entity.
type SetInboxThreadLabelsResponseSuccess struct {
}

// SetInboxThreadLabelsResponseSuccessUpdateData is the typed request payload for SetInboxThreadLabelsResponseSuccess.UpdateTyped.
type SetInboxThreadLabelsResponseSuccessUpdateData struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	Labels *[]any `json:"labels,omitempty"`
}

// Suppression is the typed data model for the suppression entity.
type Suppression struct {
}

// SuppressionLoadMatch is the typed request payload for Suppression.LoadTyped.
type SuppressionLoadMatch struct {
	Id string `json:"id"`
}

// Template is the typed data model for the template entity.
type Template struct {
}

// TemplateLoadMatch is the typed request payload for Template.LoadTyped.
type TemplateLoadMatch struct {
	Id string `json:"id"`
}

// TemplateCreateData is the typed request payload for Template.CreateTyped.
type TemplateCreateData struct {
	Id string `json:"id"`
	Alias *string `json:"alias,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CurrentVersionId *string `json:"current_version_id,omitempty"`
	From *string `json:"from,omitempty"`
	HasUnpublishedVersions *bool `json:"has_unpublished_versions,omitempty"`
	Html *string `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	PublishedAt *any `json:"published_at,omitempty"`
	ReplyTo *any `json:"reply_to,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Variables *[]any `json:"variables,omitempty"`
}

// Topic is the typed data model for the topic entity.
type Topic struct {
}

// TopicLoadMatch is the typed request payload for Topic.LoadTyped.
type TopicLoadMatch struct {
	Id string `json:"id"`
}

// UpdateApiKey is the typed data model for the update_api_key entity.
type UpdateApiKey struct {
}

// UpdateApiKeyUpdateData is the typed request payload for UpdateApiKey.UpdateTyped.
type UpdateApiKeyUpdateData struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateBroadcastResponseSuccess is the typed data model for the update_broadcast_response_success entity.
type UpdateBroadcastResponseSuccess struct {
}

// UpdateBroadcastResponseSuccessUpdateData is the typed request payload for UpdateBroadcastResponseSuccess.UpdateTyped.
type UpdateBroadcastResponseSuccessUpdateData struct {
	Id string `json:"id"`
	AudienceId *string `json:"audience_id,omitempty"`
	From *string `json:"from,omitempty"`
	Html *string `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	PreviewText *string `json:"preview_text,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	SegmentId *string `json:"segment_id,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
	TopicId *string `json:"topic_id,omitempty"`
}

// UpdateContactPropertyResponseSuccess is the typed data model for the update_contact_property_response_success entity.
type UpdateContactPropertyResponseSuccess struct {
}

// UpdateContactPropertyResponseSuccessUpdateData is the typed request payload for UpdateContactPropertyResponseSuccess.UpdateTyped.
type UpdateContactPropertyResponseSuccessUpdateData struct {
	Id string `json:"id"`
	FallbackValue *any `json:"fallback_value,omitempty"`
}

// UpdateContactResponseSuccess is the typed data model for the update_contact_response_success entity.
type UpdateContactResponseSuccess struct {
}

// UpdateContactResponseSuccessUpdateData is the typed request payload for UpdateContactResponseSuccess.UpdateTyped.
type UpdateContactResponseSuccessUpdateData struct {
	Id string `json:"id"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
	Unsubscribed *bool `json:"unsubscribed,omitempty"`
}

// UpdateContactTopicsResponseSuccess is the typed data model for the update_contact_topics_response_success entity.
type UpdateContactTopicsResponseSuccess struct {
}

// UpdateContactTopicsResponseSuccessUpdateData is the typed request payload for UpdateContactTopicsResponseSuccess.UpdateTyped.
type UpdateContactTopicsResponseSuccessUpdateData struct {
	ContactId string `json:"contact_id"`
	Topics *[]any `json:"topics,omitempty"`
}

// UpdateDomainResponseSuccess is the typed data model for the update_domain_response_success entity.
type UpdateDomainResponseSuccess struct {
}

// UpdateDomainResponseSuccessUpdateData is the typed request payload for UpdateDomainResponseSuccess.UpdateTyped.
type UpdateDomainResponseSuccessUpdateData struct {
	DomainId string `json:"domain_id"`
	Capabilities *map[string]any `json:"capabilities,omitempty"`
	ClickTracking *bool `json:"click_tracking,omitempty"`
	OpenTracking *bool `json:"open_tracking,omitempty"`
	Tls *string `json:"tls,omitempty"`
	TrackingSubdomain *string `json:"tracking_subdomain,omitempty"`
}

// UpdateEmailOption is the typed data model for the update_email_option entity.
type UpdateEmailOption struct {
}

// UpdateEmailOptionUpdateData is the typed request payload for UpdateEmailOption.UpdateTyped.
type UpdateEmailOptionUpdateData struct {
	EmailId string `json:"email_id"`
	ScheduledAt *string `json:"scheduled_at,omitempty"`
}

// UpdateEvent is the typed data model for the update_event entity.
type UpdateEvent struct {
}

// UpdateEventUpdateData is the typed request payload for UpdateEvent.UpdateTyped.
type UpdateEventUpdateData struct {
	Id string `json:"id"`
	Schema *any `json:"schema,omitempty"`
}

// UpdateInboxDraftResponseSuccess is the typed data model for the update_inbox_draft_response_success entity.
type UpdateInboxDraftResponseSuccess struct {
}

// UpdateInboxDraftResponseSuccessUpdateData is the typed request payload for UpdateInboxDraftResponseSuccess.UpdateTyped.
type UpdateInboxDraftResponseSuccessUpdateData struct {
	DraftId string `json:"draft_id"`
	InboxId string `json:"inbox_id"`
	Attachments *any `json:"attachments,omitempty"`
	Bcc *any `json:"bcc,omitempty"`
	Cc *any `json:"cc,omitempty"`
	Html *any `json:"html,omitempty"`
	Subject *any `json:"subject,omitempty"`
	Text *any `json:"text,omitempty"`
	To *any `json:"to,omitempty"`
}

// UpdateInboxLabelResponseSuccess is the typed data model for the update_inbox_label_response_success entity.
type UpdateInboxLabelResponseSuccess struct {
}

// UpdateInboxLabelResponseSuccessUpdateData is the typed request payload for UpdateInboxLabelResponseSuccess.UpdateTyped.
type UpdateInboxLabelResponseSuccessUpdateData struct {
	InboxId string `json:"inbox_id"`
	LabelId string `json:"label_id"`
	Color *string `json:"color,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateInboxResponseSuccess is the typed data model for the update_inbox_response_success entity.
type UpdateInboxResponseSuccess struct {
}

// UpdateInboxResponseSuccessUpdateData is the typed request payload for UpdateInboxResponseSuccess.UpdateTyped.
type UpdateInboxResponseSuccessUpdateData struct {
	InboxId string `json:"inbox_id"`
	FromName *string `json:"from_name,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdateInboxThreadResponseSuccess is the typed data model for the update_inbox_thread_response_success entity.
type UpdateInboxThreadResponseSuccess struct {
}

// UpdateInboxThreadResponseSuccessUpdateData is the typed request payload for UpdateInboxThreadResponseSuccess.UpdateTyped.
type UpdateInboxThreadResponseSuccessUpdateData struct {
	InboxId string `json:"inbox_id"`
	ThreadId string `json:"thread_id"`
	Folder *string `json:"folder,omitempty"`
	Read *bool `json:"read,omitempty"`
}

// UpdateSegmentResponseSuccess is the typed data model for the update_segment_response_success entity.
type UpdateSegmentResponseSuccess struct {
}

// UpdateSegmentResponseSuccessUpdateData is the typed request payload for UpdateSegmentResponseSuccess.UpdateTyped.
type UpdateSegmentResponseSuccessUpdateData struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdateTemplateResponseSuccess is the typed data model for the update_template_response_success entity.
type UpdateTemplateResponseSuccess struct {
}

// UpdateTemplateResponseSuccessUpdateData is the typed request payload for UpdateTemplateResponseSuccess.UpdateTyped.
type UpdateTemplateResponseSuccessUpdateData struct {
	Id string `json:"id"`
	Alias *string `json:"alias,omitempty"`
	From *string `json:"from,omitempty"`
	Html *string `json:"html,omitempty"`
	Name *string `json:"name,omitempty"`
	ReplyTo *[]any `json:"reply_to,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Text *string `json:"text,omitempty"`
	Variables *[]any `json:"variables,omitempty"`
}

// UpdateTopicResponseSuccess is the typed data model for the update_topic_response_success entity.
type UpdateTopicResponseSuccess struct {
}

// UpdateTopicResponseSuccessUpdateData is the typed request payload for UpdateTopicResponseSuccess.UpdateTyped.
type UpdateTopicResponseSuccessUpdateData struct {
	Id string `json:"id"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// UpdateWebhook is the typed data model for the update_webhook entity.
type UpdateWebhook struct {
}

// UpdateWebhookListMatch is the typed request payload for UpdateWebhook.ListTyped.
type UpdateWebhookListMatch struct {
	After *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

// UpdateWebhookCreateData is the typed request payload for UpdateWebhook.CreateTyped.
type UpdateWebhookCreateData struct {
	Endpoint string `json:"endpoint"`
	Events []any `json:"events"`
}

// UpdateWebhookUpdateData is the typed request payload for UpdateWebhook.UpdateTyped.
type UpdateWebhookUpdateData struct {
	Id string `json:"id"`
	Endpoint *string `json:"endpoint,omitempty"`
	Events *[]any `json:"events,omitempty"`
	Status *string `json:"status,omitempty"`
}

// Usage is the typed data model for the usage entity.
type Usage struct {
}

// UsageLoadMatch is the typed request payload for Usage.LoadTyped.
type UsageLoadMatch struct {
	AiCredits *map[string]any `json:"ai_credits,omitempty"`
	AutomationRuns *map[string]any `json:"automation_runs,omitempty"`
	Broadcasts *map[string]any `json:"broadcasts,omitempty"`
	Contacts *map[string]any `json:"contacts,omitempty"`
	Domains *map[string]any `json:"domains,omitempty"`
	Emails *map[string]any `json:"emails,omitempty"`
	Object *string `json:"object,omitempty"`
	RateLimit *map[string]any `json:"rate_limit,omitempty"`
	Segments *map[string]any `json:"segments,omitempty"`
}

// Webhook is the typed data model for the webhook entity.
type Webhook struct {
}

// WebhookLoadMatch is the typed request payload for Webhook.LoadTyped.
type WebhookLoadMatch struct {
	Id string `json:"id"`
}

// WebhookRemoveMatch is the typed request payload for Webhook.RemoveTyped.
type WebhookRemoveMatch struct {
	Id string `json:"id"`
}

// WebhookEvent is the typed data model for the webhook_event entity.
type WebhookEvent struct {
}

// WebhookEventLoadMatch is the typed request payload for WebhookEvent.LoadTyped.
type WebhookEventLoadMatch struct {
	Id string `json:"id"`
	WebhookId string `json:"webhook_id"`
}

// WebhookEventCreateData is the typed request payload for WebhookEvent.CreateTyped.
type WebhookEventCreateData struct {
	EventId string `json:"event_id"`
	WebhookId string `json:"webhook_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	NextAttemptAt *any `json:"next_attempt_at,omitempty"`
	Object *string `json:"object,omitempty"`
	Payload *map[string]any `json:"payload,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
