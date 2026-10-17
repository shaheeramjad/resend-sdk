# Resend

Resend is the email platform for developers.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 41 entities and 142 HTTP routes. There are 3 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### ApiKey

Results: OK.

SDK operations: `create`, `list`, `remove`, `update`.

Key fields to recognise:

- `created_at`: The date and time the API key was created.
- `domain_id`: Restrict an API key to send emails only from a specific domain.
- `id`: The ID of the API key.
- `last_used_at`: The date and time the API key was last used.
- `name`: The name of the API key.

### Attachment

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `content_disposition`: How the attachment should be displayed.
- `content_id`: The content ID for inline attachments.
- `content_type`: The MIME type of the attachment.
- `download_url`: Signed URL to download the attachment content.
- `expires_at`: Timestamp when the download URL expires.

### Audience

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the audience was created.
- `id`: The ID of the audience.
- `name`: The name of the audience.
- `object`: The object of the audience.

### Automation

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `connections`: The connections between steps in the active version of the automation.
- `created_at`: The date and time the automation was created.
- `id`: The ID of the duplicated automation.
- `name`: The name of the automation.
- `object`: Type of the response object.

### AutomationRun

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `completed_at`: The date and time the run completed.
- `created_at`: The date and time the run was created.
- `id`: The ID of the automation run.
- `object`: Type of the response object.
- `started_at`: The date and time the run started.

### Broadcast

A marketing email sent to a segment.

Results: OK; Created.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `audience_id`: Deprecated. Use `segment_id` instead.
- `created_at`: Timestamp indicating when the broadcast was created.
- `from`: The email address of the sender.
- `html`: The HTML version of the broadcast content.
- `id`: The ID of the broadcast.

### BroadcastClickedLink

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `clicks`: Total number of clicks on this URL.
- `id`: An opaque cursor for this row, used only for pagination. It does not identify any entity in Resend.
- `unique_clicks`: Number of unique clicks on this URL.
- `url`: The URL that was clicked.

### BroadcastRecipient

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `bounce_type`: The type of bounce. Only present when `type` is `bounced`.
- `clicked_links`: The links this recipient clicked. Only present when `type` is `clicked`.
- `contact_id`: The ID of the contact associated with this recipient, if one exists.
- `count`: The number of times this recipient triggered the event. Only present when `type` is `opened` or `clicked`.
- `email`: The recipient&#39;s email address.

### Contact

A person on your contact list.

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `audience_id`: Unique identifier of the audience to which the contact belongs.
- `created_at`: Timestamp indicating when the contact was created.
- `email`: Email address of the contact.
- `first_name`: First name of the contact.
- `id`: Unique identifier for the created contact.

### ContactImport

Results: OK.

SDK operations: `create`, `list`, `load`.

Key fields to recognise:

- `column_map`: JSON-encoded object mapping contact fields and custom property keys to CSV column names.
- `completed_at`: Timestamp indicating when the contact import completed.
- `created_at`: Timestamp indicating when the contact import was created.
- `file`: CSV file to import.
- `id`: Unique identifier for the created contact import.

### ContactProperty

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the contact property was created.
- `fallback_value`: The default value when the property is not set for a contact.
- `id`: The ID of the contact property.
- `key`: The property key.
- `object`: The object type of the response.

### ContactSegment

Results: OK.

SDK operations: `create`, `list`, `remove`.

Key fields to recognise:

- `contact_id`: The ID of the contact.
- `created_at`: Timestamp indicating when the contact was added to the segment.
- `id`: Unique identifier for the segment.
- `name`: Name of the segment.
- `object`: The object type.

### ContactTopic

Results: OK.

SDK operations: `list`, `update`.

Key fields to recognise:

- `contact_id`: The ID of the contact.
- `description`: Description of the topic.
- `id`: Unique identifier for the topic.
- `name`: Name of the topic.
- `object`: Type of the response object.

### Domain

A sending domain and its DNS verification state.

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `capabilities`: Configure the domain capabilities for sending and receiving emails. At least one capability must be enabled.
- `click_tracking`: Whether click tracking is enabled for this domain.
- `created_at`: The date and time the domain was created.
- `custom_return_path`: For advanced use cases, choose a subdomain for the Return-Path address.
- `id`: The ID of the domain.

### DomainClaim

Results: OK; An identical pending claim already existed and was returned unchanged.; Claim created.

SDK operations: `create`, `load`.

Key fields to recognise:

- `blocked_reason`: Why the claim is currently blocked, if applicable.
- `click_tracking`: Track clicks within the body of each HTML email.
- `created_at`: The date and time the claim was created.
- `custom_return_path`: For advanced use cases, choose a subdomain for the Return-Path address.
- `domain_id`: The ID of the placeholder domain created for the claim.

### Email

A transactional email sent through Resend.

Results: OK.

SDK operations: `create`, `list`, `load`, `update`.

Key fields to recognise:

- `bcc`: The email addresses of the blind carbon copy recipients.
- `cc`: The email addresses of the carbon copy recipients.
- `created_at`: The date and time the email was created.
- `from`: The email address of the sender.
- `headers`: Custom headers to add to the email.

### EmailBatch

Results: OK.

SDK operations: `create`.

### EmailsMetric

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `broadcast_id`: Present when `broadcast` is in `dimensions`.
- `broadcast_name`: Present when `broadcast` is in `dimensions`.
- `domain_id`: Present when `domain` is in `dimensions`.
- `domain_name`: Present when `domain` is in `dimensions`.
- `email_id`: Present when `email` is in `dimensions`.

### Event

Results: OK; Accepted.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `created_at`: The date and time the event was created.
- `id`: The ID of the created event.
- `name`: The event name.
- `object`: Type of the response object.
- `schema`: A flat key/type map defining the event payload schema. Supported types are `string`, `number`, `boolean`, and `date`.

### Inbox

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `created_at`: ISO 8601 timestamp when the inbox was created.
- `domain_id`: The ID of the domain that owns the address.
- `drafts`: The number of unsent drafts.
- `email_address`: The address of the inbox.
- `forwarding`: When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record.

### InboxAgent

Results: OK.

SDK operations: `load`, `update`.

Key fields to recognise:

- `actions`: The actions the agent can take, in the order of the enum. Empty when no actions are enabled.
- `id`: The ID of the inbox agent.
- `instructions`: Instructions the agent follows when it handles the threads of the inbox. Null when no instructions are set.
- `object`: The object type. Always `inbox_agent`.
- `tone`: The tone the agent writes in. Null when no tone is set.

### InboxDraft

Results: OK; The reply draft for this email already exists. The response holds the ID of the existing draft, and the request fields other than `attachments` are not applied.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `attachments`: The files stored on the draft.
- `bcc`: The BCC recipients.
- `cc`: The CC recipients.
- `created_at`: Timestamp when the draft was created.
- `email_id`: The ID of the outbound email queued when the draft was sent. `null` until the draft is sent.

### InboxLabel

Results: OK.

SDK operations: `create`, `list`, `remove`, `update`.

Key fields to recognise:

- `color`: The color of the label, as a `#RRGGBB` hex color in uppercase.
- `created_at`: Timestamp when the label was created.
- `id`: The ID of the label.
- `name`: The name of the label.
- `object`: Type of the response object.

### InboxThread

Results: OK.

SDK operations: `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `bcc`: The BCC recipients of the most recent email. Always empty for a received email.
- `cc`: The CC recipients of the most recent email.
- `email_count`: The number of emails in the thread.
- `folder`: The folder the thread is in. Every thread is in exactly one folder. `inbox` holds threads with at least one received email. `archive` and `spam` hold threads moved there. `trash` holds threads scheduled for deletion in 30 days. `sent` holds threads whose emails are all outbound, and is assigned automatically.
- `from`: The sender email address. Can include the display name.

### InboxThreadEmail

Results: OK.

SDK operations: `create`, `list`, `load`.

Key fields to recognise:

- `attachments`: The attachments of the message.
- `bcc`: The BCC recipients of the message. Always empty for a received message.
- `cc`: The CC recipients of the message.
- `direction`: Whether the inbox received the message (`inbound`) or sent it (`outbound`). A message the inbox sent to itself is `inbound`, so a reply to it goes to its sender.
- `from`: Sender email address. Includes the display name when the sender has one, in the format &quot;Name &lt;sender@domain.com&gt;&quot;.

### InboxThreadLabel

Results: OK.

SDK operations: `create`, `remove`, `update`.

Key fields to recognise:

- `data`: Every label on the thread after the change.
- `id`: The ID of the label.
- `labels`: The IDs of the labels to apply to the thread.
- `object`: Type of the response object.

### Log

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `created_at`: The date the log was created.
- `endpoint`: The API endpoint that was called.
- `id`: The log ID.
- `method`: The HTTP method used.
- `object`: Type of the response object.

### OauthGrant

Results: OK.

SDK operations: `list`, `remove`.

Key fields to recognise:

- `client`: The OAuth client the grant was issued to.
- `client_id`: The ID of the OAuth client the grant was issued to.
- `created_at`: The date and time the OAuth grant was created.
- `id`: The ID of the OAuth grant.
- `revoked_at`: The date and time the OAuth grant was revoked, or null if it is still active.

### ReceivedAttachment

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `content_disposition`: How the attachment should be displayed.
- `content_id`: The content ID for inline attachments.
- `content_type`: The MIME type of the attachment.
- `download_url`: Signed URL to download the attachment content.
- `expires_at`: Timestamp when the download URL expires.

### ReceivedEmail

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `attachments`: Array of attachments for this email.
- `bcc`: The BCC recipients.
- `cc`: The CC recipients.
- `created_at`: Timestamp when the email was received.
- `from`: The sender email address.

### Segment

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `audience_id`: The ID of the audience this segment belongs to.
- `created_at`: Timestamp indicating when the segment was created.
- `filter`: Filter conditions for the segment.
- `id`: The ID of the segment.
- `name`: Name of the segment.

### SegmentContact

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the contact was created.
- `email`: Email address of the contact.
- `first_name`: First name of the contact.
- `id`: Unique identifier for the contact.
- `last_name`: Last name of the contact.

### Suppression

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the suppression was created.
- `email`: Email address that is suppressed.
- `id`: Unique identifier for the created suppression.
- `object`: Type of the response object.
- `origin`: Origin of the suppression.

### SuppressionBatchAdd

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `emails`: Email addresses to suppress.

### SuppressionBatchRemove

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `data`: Array containing the removed suppressions.
- `emails`: Email addresses to remove from the suppression list.
- `ids`: Suppression IDs to remove from the suppression list.

### Template

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `alias`: The alias of the template.
- `created_at`: Timestamp indicating when the template was created.
- `current_version_id`: The ID of the current version of the template.
- `from`: Sender email address. To include a friendly name, use the format &quot;Your Name &lt;sender@domain.com&gt;&quot;.
- `has_unpublished_versions`: Indicates whether the template has unpublished versions.

### Topic

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the topic was created.
- `default_subscription`: The default subscription status for the topic.
- `description`: A description of the topic.
- `id`: The ID of the topic.
- `name`: Name of the topic.

### Usage

Results: OK.

SDK operations: `load`.

Key fields to recognise:

- `object`: The type of object.

### Webhook

An endpoint Resend posts delivery events to.

Results: OK; Created.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the webhook was created.
- `endpoint`: The URL where webhook events are sent.
- `events`: Array of event types subscribed to.
- `id`: The ID of the webhook.
- `object`: The type of object.

### WebhookEvent

Results: OK.

SDK operations: `create`, `list`, `load`.

Key fields to recognise:

- `created_at`: Timestamp indicating when the event was created.
- `id`: The ID of the replayed webhook event.
- `next_attempt_at`: Timestamp of the next scheduled delivery attempt, or null when none is scheduled. Always null once the event has succeeded or permanently failed.
- `object`: The type of object.
- `payload`: The event payload sent to the webhook endpoint.

### WebhookEventAttempt

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `http_status_code`: The HTTP status code returned by the webhook endpoint.
- `id`: The ID of the webhook event attempt.
- `response`: The response body returned by the webhook endpoint.
- `sent_at`: Timestamp indicating when the attempt was sent.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| ApiKey | `create` | `POST /api-keys` | Required |
| ApiKey | `list` | `GET /api-keys` | Required |
| ApiKey | `remove` | `DELETE /api-keys/{api_key_id}` | Required |
| ApiKey | `update` | `PATCH /api-keys/{api_key_id}` | Required |
| Attachment | `list` | `GET /emails/{email_id}/attachments` | Required |
| Attachment | `load` | `GET /emails/{email_id}/attachments/{attachment_id}` | Required |
| Audience | `create` | `POST /audiences` | Required |
| Audience | `list` | `GET /audiences` | Required |
| Audience | `load` | `GET /audiences/{id}` | Required |
| Audience | `remove` | `DELETE /audiences/{id}` | Required |
| Automation | `create` | `POST /automations/{automation_id}/duplicate` | Required |
| Automation | `create` | `POST /automations/{automation_id}/stop` | Required |
| Automation | `create` | `POST /automations` | Required |
| Automation | `list` | `GET /automations` | Required |
| Automation | `load` | `GET /automations/{automation_id}` | Required |
| Automation | `remove` | `DELETE /automations/{automation_id}` | Required |
| Automation | `update` | `PATCH /automations/{automation_id}` | Required |
| AutomationRun | `list` | `GET /automations/{automation_id}/runs` | Required |
| AutomationRun | `load` | `GET /automations/{automation_id}/runs/{run_id}` | Required |
| Broadcast | `create` | `POST /broadcasts/{id}/cancel` | Required |
| Broadcast | `create` | `POST /broadcasts/{id}/duplicate` | Required |
| Broadcast | `create` | `POST /broadcasts/{id}/send` | Required |
| Broadcast | `create` | `POST /broadcasts` | Required |
| Broadcast | `list` | `GET /broadcasts` | Required |
| Broadcast | `load` | `GET /broadcasts/{id}` | Required |
| Broadcast | `remove` | `DELETE /broadcasts/{id}` | Required |
| Broadcast | `update` | `PATCH /broadcasts/{id}` | Required |
| BroadcastClickedLink | `list` | `GET /broadcasts/{id}/clicked-links` | Required |
| BroadcastRecipient | `list` | `GET /broadcasts/{id}/recipients` | Required |
| Contact | `create` | `POST /contacts` | Required |
| Contact | `list` | `GET /contacts` | Required |
| Contact | `load` | `GET /contacts/{id}` | Required |
| Contact | `remove` | `DELETE /contacts/{id}` | Required |
| Contact | `update` | `PATCH /contacts/{id}` | Required |
| ContactImport | `create` | `POST /contacts/imports` | Required |
| ContactImport | `list` | `GET /contacts/imports` | Required |
| ContactImport | `load` | `GET /contacts/imports/{id}` | Required |
| ContactProperty | `create` | `POST /contact-properties` | Required |
| ContactProperty | `list` | `GET /contact-properties` | Required |
| ContactProperty | `load` | `GET /contact-properties/{id}` | Required |
| ContactProperty | `remove` | `DELETE /contact-properties/{id}` | Required |
| ContactProperty | `update` | `PATCH /contact-properties/{id}` | Required |
| ContactSegment | `create` | `POST /contacts/{contact_id}/segments/{segment_id}` | Required |
| ContactSegment | `list` | `GET /contacts/{contact_id}/segments` | Required |
| ContactSegment | `remove` | `DELETE /contacts/{contact_id}/segments/{segment_id}` | Required |
| ContactTopic | `list` | `GET /contacts/{contact_id}/topics` | Required |
| ContactTopic | `update` | `PATCH /contacts/{contact_id}/topics` | Required |
| Domain | `create` | `POST /domains/{domain_id}/verify` | Required |
| Domain | `create` | `POST /domains` | Required |
| Domain | `list` | `GET /domains` | Required |
| Domain | `load` | `GET /domains/{domain_id}` | Required |
| Domain | `remove` | `DELETE /domains/{domain_id}` | Required |
| Domain | `update` | `PATCH /domains/{domain_id}` | Required |
| DomainClaim | `create` | `POST /domains/{domain_id}/claim/verify` | Required |
| DomainClaim | `create` | `POST /domains/claim` | Required |
| DomainClaim | `load` | `GET /domains/{domain_id}/claim` | Required |
| Email | `create` | `POST /emails/{email_id}/cancel` | Required |
| Email | `create` | `POST /emails/{email_id}/share` | Required |
| Email | `create` | `POST /emails` | Required |
| Email | `list` | `GET /emails` | Required |
| Email | `load` | `GET /emails/{email_id}` | Required |
| Email | `update` | `PATCH /emails/{email_id}` | Required |
| EmailBatch | `create` | `POST /emails/batch` | Required |
| EmailsMetric | `list` | `GET /emails/metrics` | Required |
| Event | `create` | `POST /events` | Required |
| Event | `create` | `POST /events/send` | Required |
| Event | `list` | `GET /events` | Required |
| Event | `load` | `GET /events/{identifier}` | Required |
| Event | `remove` | `DELETE /events/{identifier}` | Required |
| Event | `update` | `PATCH /events/{identifier}` | Required |
| Inbox | `create` | `POST /inboxes` | Required |
| Inbox | `list` | `GET /inboxes` | Required |
| Inbox | `load` | `GET /inboxes/{inbox_id}` | Required |
| Inbox | `remove` | `DELETE /inboxes/{inbox_id}` | Required |
| Inbox | `update` | `PATCH /inboxes/{inbox_id}` | Required |
| InboxAgent | `load` | `GET /inboxes/{inbox_id}/agent` | Required |
| InboxAgent | `update` | `PATCH /inboxes/{inbox_id}/agent` | Required |
| InboxDraft | `create` | `POST /inboxes/{inbox_id}/drafts/{draft_id}/send` | Required |
| InboxDraft | `create` | `POST /inboxes/{inbox_id}/drafts` | Required |
| InboxDraft | `list` | `GET /inboxes/{inbox_id}/drafts` | Required |
| InboxDraft | `load` | `GET /inboxes/{inbox_id}/drafts/{draft_id}` | Required |
| InboxDraft | `remove` | `DELETE /inboxes/{inbox_id}/drafts/{draft_id}` | Required |
| InboxDraft | `update` | `PATCH /inboxes/{inbox_id}/drafts/{draft_id}` | Required |
| InboxLabel | `create` | `POST /inboxes/{inbox_id}/labels` | Required |
| InboxLabel | `list` | `GET /inboxes/{inbox_id}/labels` | Required |
| InboxLabel | `remove` | `DELETE /inboxes/{inbox_id}/labels/{label_id}` | Required |
| InboxLabel | `update` | `PATCH /inboxes/{inbox_id}/labels/{label_id}` | Required |
| InboxThread | `list` | `GET /inboxes/{inbox_id}/threads` | Required |
| InboxThread | `list` | `GET /inboxes/{inbox_id}/threads/search` | Required |
| InboxThread | `load` | `GET /inboxes/{inbox_id}/threads/{thread_id}` | Required |
| InboxThread | `remove` | `DELETE /inboxes/{inbox_id}/threads/{thread_id}` | Required |
| InboxThread | `update` | `PATCH /inboxes/{inbox_id}/threads/{thread_id}` | Required |
| InboxThreadEmail | `create` | `POST /inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/forward` | Required |
| InboxThreadEmail | `create` | `POST /inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/reply` | Required |
| InboxThreadEmail | `list` | `GET /inboxes/{inbox_id}/threads/{thread_id}/emails` | Required |
| InboxThreadEmail | `load` | `GET /inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}` | Required |
| InboxThreadLabel | `create` | `POST /inboxes/{inbox_id}/threads/{thread_id}/labels` | Required |
| InboxThreadLabel | `remove` | `DELETE /inboxes/{inbox_id}/threads/{thread_id}/labels/{label_id}` | Required |
| InboxThreadLabel | `update` | `PUT /inboxes/{inbox_id}/threads/{thread_id}/labels` | Required |
| Log | `list` | `GET /logs` | Required |
| Log | `load` | `GET /logs/{log_id}` | Required |
| OauthGrant | `list` | `GET /oauth/grants` | Required |
| OauthGrant | `remove` | `DELETE /oauth/grants/{oauth_grant_id}` | Required |
| ReceivedAttachment | `list` | `GET /emails/receiving/{email_id}/attachments` | Required |
| ReceivedAttachment | `load` | `GET /emails/receiving/{email_id}/attachments/{attachment_id}` | Required |
| ReceivedEmail | `list` | `GET /emails/receiving` | Required |
| ReceivedEmail | `load` | `GET /emails/receiving/{email_id}` | Required |
| Segment | `create` | `POST /segments` | Required |
| Segment | `list` | `GET /segments` | Required |
| Segment | `load` | `GET /segments/{id}` | Required |
| Segment | `remove` | `DELETE /segments/{id}` | Required |
| Segment | `update` | `PATCH /segments/{id}` | Required |
| SegmentContact | `list` | `GET /segments/{id}/contacts` | Required |
| Suppression | `create` | `POST /suppressions` | Required |
| Suppression | `list` | `GET /suppressions` | Required |
| Suppression | `load` | `GET /suppressions/{suppression}` | Required |
| Suppression | `remove` | `DELETE /suppressions/{suppression}` | Required |
| SuppressionBatchAdd | `create` | `POST /suppressions/batch/add` | Required |
| SuppressionBatchRemove | `create` | `POST /suppressions/batch/remove` | Required |
| Template | `create` | `POST /templates/{id}/duplicate` | Required |
| Template | `create` | `POST /templates/{id}/publish` | Required |
| Template | `create` | `POST /templates` | Required |
| Template | `list` | `GET /templates` | Required |
| Template | `load` | `GET /templates/{id}` | Required |
| Template | `remove` | `DELETE /templates/{id}` | Required |
| Template | `update` | `PATCH /templates/{id}` | Required |
| Topic | `create` | `POST /topics` | Required |
| Topic | `list` | `GET /topics` | Required |
| Topic | `load` | `GET /topics/{id}` | Required |
| Topic | `remove` | `DELETE /topics/{id}` | Required |
| Topic | `update` | `PATCH /topics/{id}` | Required |
| Usage | `load` | `GET /usage` | Required |
| Webhook | `create` | `POST /webhooks/{webhook_id}/signing-secret/rotate` | Required |
| Webhook | `create` | `POST /webhooks` | Required |
| Webhook | `list` | `GET /webhooks` | Required |
| Webhook | `load` | `GET /webhooks/{webhook_id}` | Required |
| Webhook | `remove` | `DELETE /webhooks/{webhook_id}` | Required |
| Webhook | `update` | `PATCH /webhooks/{webhook_id}` | Required |
| WebhookEvent | `create` | `POST /webhooks/{webhook_id}/events/{event_id}/replay` | Required |
| WebhookEvent | `list` | `GET /webhooks/{webhook_id}/events` | Required |
| WebhookEvent | `load` | `GET /webhooks/{webhook_id}/events/{event_id}` | Required |
| WebhookEventAttempt | `list` | `GET /webhooks/{webhook_id}/events/{event_id}/attempts` | Required |

## Connect to the API

- API server: `https://api.resend.com`

The default credential is sent in the `Authorization` header with the `Bearer` prefix.

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Python | `py/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### Go CLI

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### Go MCP server

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `resend_list`: List records for an entity. Supported entities: `api_key`, `attachment`, `audience`, `automation`, `automation_run`, `broadcast`, `broadcast_clicked_link`, `broadcast_recipient`, `contact`, `contact_import`, `contact_property`, `contact_segment`, `contact_topic`, `domain`, `email`, `emails_metric`, `event`, `inbox`, `inbox_draft`, `inbox_label`, `inbox_thread`, `inbox_thread_email`, `log`, `oauth_grant`, `received_attachment`, `received_email`, `segment`, `segment_contact`, `suppression`, `template`, `topic`, `webhook`, `webhook_event`, `webhook_event_attempt`.
- `resend_load`: Load one record for an entity. Supported entities: `attachment`, `audience`, `automation`, `automation_run`, `broadcast`, `contact`, `contact_import`, `contact_property`, `domain`, `domain_claim`, `email`, `event`, `inbox`, `inbox_agent`, `inbox_draft`, `inbox_thread`, `inbox_thread_email`, `log`, `received_attachment`, `received_email`, `segment`, `suppression`, `template`, `topic`, `usage`, `webhook`, `webhook_event`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `debug`: Request/response capture ring buffer for debugging
- `idempotency`: Idempotency keys for safe retries of mutating operations
- `metrics`: Statistics capture: per-operation counters and latency
- `netsim`: Network behaviour simulation for offline testing (latency, failures, outages)
- `paging`: Pagination signals for list operations
- `ratelimit`: Client-side rate limiting via a token bucket
- `retry`: Automatic retry of transient failures with exponential backoff
- `test`: In-memory mock transport for testing without a live server
- `timeout`: Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

