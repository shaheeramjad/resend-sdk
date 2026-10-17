# Resend Python SDK

Send email, manage domains, contacts and broadcasts on Resend from TypeScript, Python and Go.

The Python SDK for the Resend API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.ApiKey()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Tags](https://github.com/shaheeramjad/resend-sdk/tags)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from resend_sdk import ResendSDK

client = ResendSDK({
    "apikey": os.environ.get("RESEND_APIKEY"),
})
```

### 2. List apikey records

`list()` returns a `list` of entities, one per record, and raises on
error; an entity's `data_get()` reads its record (a `dict`).

```python
try:
    apikeys = client.ApiKey().list()
    for apikey in apikeys:
        print(apikey.data_get())
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load an attachment

Attachment is nested under email, so provide the `email_id`.
`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    attachment = client.Attachment().load({"email_id": "example_email_id", "id": "example_id"})
    print(attachment.data_get())
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.ApiKey().create({"name": "example_name"})

# Update — the created entity's record holds its id
client.ApiKey().update({"id": created.data_get()["id"], "name": "example_name"})

# Remove
client.ApiKey().remove({"id": created.data_get()["id"]})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    oauthgrants = client.OauthGrant().list()
    print([item.data_get() for item in oauthgrants])
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = ResendSDK.test()

# Entity ops return the entity, and list one per record; they raise on error.
oauthgrant = client.OauthGrant().list()
# data_get() on an entity reads its mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = ResendSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
RESEND_TEST_LIVE=TRUE
RESEND_APIKEY=<your-key>
```

Then run:

```bash
cd py && pytest test/
```


## Reference

### ResendSDK

```python
from resend_sdk import ResendSDK

client = ResendSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = ResendSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### ResendSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `ApiKey` | `(data) -> ApiKeyEntity` | Create an ApiKey entity instance. |
| `Attachment` | `(data) -> AttachmentEntity` | Create an Attachment entity instance. |
| `Audience` | `(data) -> AudienceEntity` | Create an Audience entity instance. |
| `Automation` | `(data) -> AutomationEntity` | Create an Automation entity instance. |
| `AutomationRun` | `(data) -> AutomationRunEntity` | Create an AutomationRun entity instance. |
| `Broadcast` | `(data) -> BroadcastEntity` | Create a Broadcast entity instance. |
| `BroadcastClickedLink` | `(data) -> BroadcastClickedLinkEntity` | Create a BroadcastClickedLink entity instance. |
| `BroadcastRecipient` | `(data) -> BroadcastRecipientEntity` | Create a BroadcastRecipient entity instance. |
| `Contact` | `(data) -> ContactEntity` | Create a Contact entity instance. |
| `ContactImport` | `(data) -> ContactImportEntity` | Create a ContactImport entity instance. |
| `ContactProperty` | `(data) -> ContactPropertyEntity` | Create a ContactProperty entity instance. |
| `ContactSegment` | `(data) -> ContactSegmentEntity` | Create a ContactSegment entity instance. |
| `ContactTopic` | `(data) -> ContactTopicEntity` | Create a ContactTopic entity instance. |
| `Domain` | `(data) -> DomainEntity` | Create a Domain entity instance. |
| `DomainClaim` | `(data) -> DomainClaimEntity` | Create a DomainClaim entity instance. |
| `Email` | `(data) -> EmailEntity` | Create an Email entity instance. |
| `EmailBatch` | `(data) -> EmailBatchEntity` | Create an EmailBatch entity instance. |
| `EmailsMetric` | `(data) -> EmailsMetricEntity` | Create an EmailsMetric entity instance. |
| `Event` | `(data) -> EventEntity` | Create an Event entity instance. |
| `Inbox` | `(data) -> InboxEntity` | Create an Inbox entity instance. |
| `InboxAgent` | `(data) -> InboxAgentEntity` | Create an InboxAgent entity instance. |
| `InboxDraft` | `(data) -> InboxDraftEntity` | Create an InboxDraft entity instance. |
| `InboxLabel` | `(data) -> InboxLabelEntity` | Create an InboxLabel entity instance. |
| `InboxThread` | `(data) -> InboxThreadEntity` | Create an InboxThread entity instance. |
| `InboxThreadEmail` | `(data) -> InboxThreadEmailEntity` | Create an InboxThreadEmail entity instance. |
| `InboxThreadLabel` | `(data) -> InboxThreadLabelEntity` | Create an InboxThreadLabel entity instance. |
| `Log` | `(data) -> LogEntity` | Create a Log entity instance. |
| `OauthGrant` | `(data) -> OauthGrantEntity` | Create an OauthGrant entity instance. |
| `ReceivedAttachment` | `(data) -> ReceivedAttachmentEntity` | Create a ReceivedAttachment entity instance. |
| `ReceivedEmail` | `(data) -> ReceivedEmailEntity` | Create a ReceivedEmail entity instance. |
| `Segment` | `(data) -> SegmentEntity` | Create a Segment entity instance. |
| `SegmentContact` | `(data) -> SegmentContactEntity` | Create a SegmentContact entity instance. |
| `Suppression` | `(data) -> SuppressionEntity` | Create a Suppression entity instance. |
| `SuppressionBatchAdd` | `(data) -> SuppressionBatchAddEntity` | Create a SuppressionBatchAdd entity instance. |
| `SuppressionBatchRemove` | `(data) -> SuppressionBatchRemoveEntity` | Create a SuppressionBatchRemove entity instance. |
| `Template` | `(data) -> TemplateEntity` | Create a Template entity instance. |
| `Topic` | `(data) -> TopicEntity` | Create a Topic entity instance. |
| `Usage` | `(data) -> UsageEntity` | Create an Usage entity instance. |
| `Webhook` | `(data) -> WebhookEntity` | Create a Webhook entity instance. |
| `WebhookEvent` | `(data) -> WebhookEventEntity` | Create a WebhookEvent entity instance. |
| `WebhookEventAttempt` | `(data) -> WebhookEventAttemptEntity` | Create a WebhookEventAttempt entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria, and return it. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria, one per record. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity, and return it. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity, and return it. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity, and return it marked as deleted. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the entity, and `list` a `list` of entities, one
per record; an entity's `data_get()` reads its record (a `dict`). They raise
on error, so wrap calls in `try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### ApiKey

| Field | Description |
| --- | --- |
| `created_at` | The date and time the API key was created. |
| `domain_id` | Restrict an API key to send emails only from a specific domain. |
| `id` | The ID of the API key. |
| `last_used_at` | The date and time the API key was last used. |
| `name` | The API key name. |
| `object` | The type of object. |
| `permission` | The API key can have full access to Resend’s API or be only restricted to send emails. |

Operations: Create, List, Remove, Update.

API path: `/api-keys`

#### Attachment

| Field | Description |
| --- | --- |
| `content_disposition` | How the attachment should be displayed. |
| `content_id` | The content ID for inline attachments. |
| `content_type` | The MIME type of the attachment. |
| `download_url` | Signed URL to download the attachment content. |
| `expires_at` | Timestamp when the download URL expires. |
| `filename` | The filename of the attachment. |
| `id` | The ID of the attachment. |
| `object` | The type of object. |
| `size` | Size of the attachment in bytes. |

Operations: List, Load.

API path: `/emails/{email_id}/attachments`

#### Audience

| Field | Description |
| --- | --- |
| `created_at` | The date that the object was created. |
| `id` | The ID of the audience. |
| `name` | The name of the audience. |
| `object` | The object of the audience. |

Operations: Create, List, Load, Remove.

API path: `/audiences`

#### Automation

| Field | Description |
| --- | --- |
| `connections` | The connections between steps in the active version of the automation. |
| `created_at` | The date and time the automation was created. |
| `id` | The ID of the automation. |
| `name` | The name of the automation. |
| `object` | Type of the response object. |
| `status` | The current status of the automation. |
| `steps` | The steps in the active version of the automation. |
| `updated_at` | The date and time the automation was last updated. |

Operations: Create, List, Load, Remove, Update.

API path: `/automations/{automation_id}/duplicate`

#### AutomationRun

| Field | Description |
| --- | --- |
| `completed_at` | The date and time the run completed. |
| `created_at` | The date and time the run was created. |
| `id` | The ID of the automation run. |
| `object` | Type of the response object. |
| `started_at` | The date and time the run started. |
| `status` | The current status of the automation run. |
| `steps` | The steps executed in this run, sorted in graph order. |

Operations: List, Load.

API path: `/automations/{automation_id}/runs`

#### Broadcast

| Field | Description |
| --- | --- |
| `audience_id` | Deprecated: use `segment_id` instead. |
| `created_at` | Timestamp indicating when the broadcast was created. |
| `from` | The email address of the sender. |
| `html` | The HTML version of the broadcast content. |
| `id` | Unique identifier for the broadcast. |
| `name` | Name of the broadcast. |
| `object` | The object type of the response. |
| `preview_text` | The preview text of the email. |
| `reply_to` | The email addresses to which replies should be sent. |
| `scheduled_at` | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | Unique identifier of the segment this broadcast will be sent to. |
| `send` | Whether to send the broadcast immediately or keep it as a draft. |
| `sent_at` | Timestamp indicating when the broadcast was sent. |
| `status` | The status of the broadcast. |
| `subject` | The subject line of the email. |
| `text` | The plain text version of the broadcast content. |
| `topic_id` | The topic ID that the broadcast is scoped to. |

Operations: Create, List, Load, Remove, Update.

API path: `/broadcasts/{id}/cancel`

#### BroadcastClickedLink

| Field | Description |
| --- | --- |
| `clicks` | Total number of clicks on this URL. |
| `id` | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | Number of unique clicks on this URL. |
| `url` | The URL that was clicked. |

Operations: List.

API path: `/broadcasts/{id}/clicked-links`

#### BroadcastRecipient

| Field | Description |
| --- | --- |
| `bounce_type` | The type of bounce. |
| `clicked_links` | The links this recipient clicked. |
| `contact_id` | The ID of the contact associated with this recipient, if one exists. |
| `count` | The number of times this recipient triggered the event. |
| `email` | The recipient's email address. |
| `id` | Opaque cursor identifying this row, used for pagination. |

Operations: List.

API path: `/broadcasts/{id}/recipients`

#### Contact

| Field | Description |
| --- | --- |
| `audience_id` | Unique identifier of the audience to which the contact belongs. |
| `created_at` | Timestamp indicating when the contact was created. |
| `email` | Email address of the contact. |
| `first_name` | First name of the contact. |
| `id` | Unique identifier for the contact. |
| `last_name` | Last name of the contact. |
| `object` | Type of the response object. |
| `properties` | A map of custom property keys and values. |
| `segments` | Array of segment IDs to add the contact to. |
| `topics` | Array of topic subscriptions for the contact. |
| `unsubscribed` | Indicates if the contact is unsubscribed. |

Operations: Create, List, Load, Remove, Update.

API path: `/contacts`

#### ContactImport

| Field | Description |
| --- | --- |
| `column_map` | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | Timestamp indicating when the contact import completed. |
| `counts` |  |
| `created_at` | Timestamp indicating when the contact import was created. |
| `file` | CSV file to import. |
| `id` | Unique identifier for the contact import. |
| `object` | Type of the response object. |
| `on_conflict` | Strategy to use when an imported contact already exists. |
| `segments` | JSON-encoded array of segments to add imported contacts to. |
| `status` | Current status of the contact import. |
| `topics` | JSON-encoded array of topic subscriptions to apply to imported contacts. |

Operations: Create, List, Load.

API path: `/contacts/imports`

#### ContactProperty

| Field | Description |
| --- | --- |
| `created_at` | Timestamp indicating when the contact property was created. |
| `fallback_value` | The default value when the property is not set for a contact. |
| `id` | The ID of the contact property. |
| `key` | The property key. |
| `object` | The object type. |
| `type` | The property type. |

Operations: Create, List, Load, Remove, Update.

API path: `/contact-properties`

#### ContactSegment

| Field | Description |
| --- | --- |
| `contact_id` | The ID of the contact. |
| `created_at` | Timestamp indicating when the contact was added to the segment. |
| `id` | Unique identifier for the segment. |
| `name` | Name of the segment. |
| `object` | The object type. |
| `segment_id` | The ID of the segment. |

Operations: Create, List, Remove.

API path: `/contacts/{contact_id}/segments/{segment_id}`

#### ContactTopic

| Field | Description |
| --- | --- |
| `contact_id` | The ID of the contact. |
| `description` | Description of the topic. |
| `id` | Unique identifier for the topic. |
| `name` | Name of the topic. |
| `object` | The object type. |
| `subscription` | The subscription status for this topic. |
| `topics` | Array of updated topic subscriptions. |

Operations: List, Update.

API path: `/contacts/{contact_id}/topics`

#### Domain

| Field | Description |
| --- | --- |
| `capabilities` | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | Whether click tracking is enabled for this domain. |
| `created_at` | The date and time the domain was created. |
| `custom_return_path` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | The ID of the domain. |
| `name` | The name of the domain. |
| `object` | The type of object. |
| `open_tracking` | Whether open tracking is enabled for this domain. |
| `records` |  |
| `region` | The region where the domain is hosted. |
| `status` | The status of the domain. |
| `tls` | TLS mode. |
| `tracking_subdomain` | The subdomain used for click and open tracking. |

Operations: Create, List, Load, Remove, Update.

API path: `/domains/{domain_id}/verify`

#### DomainClaim

| Field | Description |
| --- | --- |
| `blocked_reason` | Why the claim is currently blocked, if applicable. |
| `click_tracking` | Track clicks within the body of each HTML email. |
| `created_at` | The date and time the claim was created. |
| `custom_return_path` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | The ID of the placeholder domain created for the claim. |
| `expires_at` | The date and time the claim expires if not verified. |
| `failure_reason` | Why the claim failed, if applicable. |
| `id` | The ID of the claim. |
| `name` | The name of the domain being claimed. |
| `object` | The type of object. |
| `open_tracking` | Track the open rate of each email. |
| `record` | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | The region where the claimed domain will send from. |
| `status` | The status of the claim. |
| `tracking_subdomain` | The subdomain to use for click and open tracking. |

Operations: Create, Load.

API path: `/domains/{domain_id}/claim/verify`

#### Email

| Field | Description |
| --- | --- |
| `attachments` |  |
| `bcc` | The email addresses of the blind carbon copy recipients. |
| `cc` | The email addresses of the carbon copy recipients. |
| `created_at` | The date and time the email was created. |
| `from` | The email address of the sender. |
| `headers` | Custom headers to add to the email. |
| `html` | The HTML body of the email. |
| `id` | The ID of the email. |
| `last_event` | The status of the email. |
| `message_id` | The Message-ID header value of the email. |
| `object` | The type of object. |
| `reply_to` | The email addresses to which replies should be sent. |
| `scheduled_at` | Schedule email to be sent later. |
| `subject` | The subject line of the email. |
| `tags` |  |
| `template` |  |
| `text` | The plain text body of the email. |
| `to` | Recipient email address. |
| `topic_id` | The topic ID to scope the email to. |

Operations: Create, List, Load, Update.

API path: `/emails/{email_id}/cancel`

#### EmailBatch

| Field | Description |
| --- | --- |
| `data` |  |
| `send_emails` |  |

Operations: Create.

API path: `/emails/batch`

#### EmailsMetric

| Field | Description |
| --- | --- |
| `broadcast_id` | Present when `broadcast` is in `dimensions`. |
| `broadcast_name` | Present when `broadcast` is in `dimensions`. |
| `domain_id` | Present when `domain` is in `dimensions`. |
| `domain_name` | Present when `domain` is in `dimensions`. |
| `email_id` | Present when `email` is in `dimensions`. |
| `period` | Present when `period` is in `dimensions`. |

Operations: List.

API path: `/emails/metrics`

#### Event

| Field | Description |
| --- | --- |
| `created_at` | The date and time the event was created. |
| `id` | The event ID. |
| `name` | The event name. |
| `object` | Type of the response object. |
| `schema` | A flat key/type map defining the event payload schema. |
| `updated_at` | The date and time the event was last updated. |

Operations: Create, List, Load, Remove, Update.

API path: `/events`

#### Inbox

| Field | Description |
| --- | --- |
| `created_at` | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | The ID of the domain that owns the address. |
| `drafts` | The number of unsent drafts. |
| `email_address` | The address of the inbox. |
| `forwarding` | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | The name recipients see when mail is sent from this inbox. |
| `id` | The ID of the inbox. |
| `last_received` | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | Internal name for the inbox. |
| `object` | The object type. |
| `receiving_address` | The address to forward mail to when forwarding is enabled. |
| `unread` | The number of unread threads in the Inbox folder. |
| `updated_at` | ISO 8601 timestamp when the inbox was last updated. |

Operations: Create, List, Load, Remove, Update.

API path: `/inboxes`

#### InboxAgent

| Field | Description |
| --- | --- |
| `actions` | The actions the agent can take, in the order of the enum. |
| `id` | The ID of the inbox agent. |
| `instructions` | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | The object type. |
| `tone` | The tone the agent writes in. |

Operations: Load, Update.

API path: `/inboxes/{inbox_id}/agent`

#### InboxDraft

| Field | Description |
| --- | --- |
| `attachments` | The files stored on the draft. |
| `bcc` | The BCC recipients. |
| `cc` | The CC recipients. |
| `created_at` | Timestamp when the draft was created. |
| `email_id` | The ID of the outbound email queued when the draft was sent. |
| `html` | The HTML body. |
| `id` | The ID of the draft. |
| `object` | The object type. |
| `reply_to_email_id` | The Email ID the draft replies to. |
| `snippet` | A short preview of the body. |
| `subject` | The subject. |
| `text` | The plain text body. |
| `thread_id` | The Thread ID when the draft is a reply. |
| `to` | The recipients. |
| `type` | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | Timestamp when the draft was last saved. |

Operations: Create, List, Load, Remove, Update.

API path: `/inboxes/{inbox_id}/drafts/{draft_id}/send`

#### InboxLabel

| Field | Description |
| --- | --- |
| `color` | The color of the label, as a `#RRGGBB` hex color. |
| `created_at` | Timestamp when the label was created. |
| `id` | The ID of the updated label. |
| `name` | The name of the label. |
| `object` | Type of the response object. |

Operations: Create, List, Remove, Update.

API path: `/inboxes/{inbox_id}/labels`

#### InboxThread

| Field | Description |
| --- | --- |
| `bcc` | The BCC recipients of the most recent email. |
| `cc` | The CC recipients of the most recent email. |
| `email_count` | The number of emails in the thread. |
| `folder` | The folder the thread is in. |
| `from` | The sender email address. |
| `has_attachment` | Whether the most recent email in the thread has an attachment. |
| `has_draft` | Whether the thread has an unsent draft. |
| `id` | The ID of the thread. |
| `labels` | The labels attached to the thread. |
| `object` | The object type. |
| `read` | True only when every email in the thread is read. |
| `received_at` | Timestamp when the thread was last active. |
| `subject` | The subject of the thread. |
| `to` | The recipients of the most recent email. |

Operations: List, Load, Remove, Update.

API path: `/inboxes/{inbox_id}/threads`

#### InboxThreadEmail

| Field | Description |
| --- | --- |
| `attachments` | The attachments of the message. |
| `bcc` | The BCC recipients of the message. |
| `cc` | The CC recipients of the message. |
| `direction` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | Sender email address. |
| `html` | The HTML body. |
| `id` | The ID of the message. |
| `message_id` | The Message-ID header of the message. |
| `read` | Whether the message has been read. |
| `received_at` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | The Reply-To addresses of the message. |
| `subject` | The subject of the message. |
| `text` | The plain-text body. |
| `to` | The recipients of the message. |

Operations: Create, List, Load.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/forward`

#### InboxThreadLabel

| Field | Description |
| --- | --- |
| `data` | Every label on the thread after the change. |
| `id` |  |
| `labels` | The IDs of the labels to apply to the thread. |
| `object` | Type of the response object. |

Operations: Create, Remove, Update.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/labels`

#### Log

| Field | Description |
| --- | --- |
| `created_at` | The date the log was created. |
| `endpoint` | The API endpoint that was called. |
| `id` | The log ID. |
| `method` | The HTTP method used. |
| `object` | Type of the response object. |
| `request_body` | The request body sent to the API. |
| `response_body` | The response body returned by the API. |
| `response_status` | The HTTP status code of the response. |
| `user_agent` | The user agent of the request. |

Operations: List, Load.

API path: `/logs`

#### OauthGrant

| Field | Description |
| --- | --- |
| `client` | The OAuth client the grant was issued to. |
| `client_id` | The ID of the OAuth client the grant was issued to. |
| `created_at` | The date and time the OAuth grant was created. |
| `id` | The ID of the OAuth grant. |
| `revoked_at` | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | The scopes granted to the OAuth client. |

Operations: List, Remove.

API path: `/oauth/grants`

#### ReceivedAttachment

| Field | Description |
| --- | --- |
| `content_disposition` | How the attachment should be displayed. |
| `content_id` | The content ID for inline attachments. |
| `content_type` | The MIME type of the attachment. |
| `download_url` | Signed URL to download the attachment content. |
| `expires_at` | Timestamp when the download URL expires. |
| `filename` | The filename of the attachment. |
| `id` | The ID of the attachment. |
| `object` | The type of object. |
| `size` | Size of the attachment in bytes. |

Operations: List, Load.

API path: `/emails/receiving/{email_id}/attachments`

#### ReceivedEmail

| Field | Description |
| --- | --- |
| `attachments` | Array of attachments. |
| `bcc` | The BCC recipients. |
| `cc` | The CC recipients. |
| `created_at` | Timestamp when the email was received. |
| `from` | The sender email address. |
| `headers` | The email headers. |
| `html` | The HTML content of the email. |
| `id` | The ID of the received email. |
| `message_id` | The unique message ID from the email headers. |
| `object` | The type of object. |
| `received_for` | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | The reply-to addresses. |
| `subject` | The email subject. |
| `text` | The plain text content of the email. |
| `to` | The recipient email addresses. |

Operations: List, Load.

API path: `/emails/receiving`

#### Segment

| Field | Description |
| --- | --- |
| `audience_id` | The ID of the audience this segment belongs to. |
| `created_at` | Timestamp indicating when the segment was created. |
| `filter` | Filter conditions for the segment. |
| `id` | The ID of the segment. |
| `name` | The name of the segment. |
| `object` | The object type. |

Operations: Create, List, Load, Remove, Update.

API path: `/segments`

#### SegmentContact

| Field | Description |
| --- | --- |
| `created_at` | Timestamp indicating when the contact was created. |
| `email` | Email address of the contact. |
| `first_name` | First name of the contact. |
| `id` | Unique identifier for the contact. |
| `last_name` | Last name of the contact. |
| `unsubscribed` | Indicates if the contact is unsubscribed. |

Operations: List.

API path: `/segments/{id}/contacts`

#### Suppression

| Field | Description |
| --- | --- |
| `created_at` | Timestamp indicating when the suppression was created. |
| `email` | Email address that is suppressed. |
| `id` | Unique identifier for the suppression. |
| `object` | Type of the response object. |
| `origin` | Origin of the suppression. |
| `source_id` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

Operations: Create, List, Load, Remove.

API path: `/suppressions`

#### SuppressionBatchAdd

| Field | Description |
| --- | --- |
| `emails` | Email addresses to suppress. |

Operations: Create.

API path: `/suppressions/batch/add`

#### SuppressionBatchRemove

| Field | Description |
| --- | --- |
| `data` | Array containing the removed suppressions. |
| `emails` | Email addresses to remove from the suppression list. |
| `ids` | Suppression IDs to remove from the suppression list. |

Operations: Create.

API path: `/suppressions/batch/remove`

#### Template

| Field | Description |
| --- | --- |
| `alias` | The alias of the template. |
| `created_at` | Timestamp indicating when the template was created. |
| `current_version_id` | The ID of the current version of the template. |
| `from` | Sender email address. |
| `has_unpublished_versions` | Indicates whether the template has unpublished versions. |
| `html` | The HTML version of the template. |
| `id` | The ID of the template. |
| `name` | The name of the template. |
| `object` | The type of object. |
| `published_at` | Timestamp indicating when the template was published. |
| `reply_to` | Reply-to email addresses. |
| `status` | The publication status of the template. |
| `subject` | Email subject. |
| `text` | The plain text version of the template. |
| `updated_at` | Timestamp indicating when the template was last updated. |
| `variables` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/templates/{id}/duplicate`

#### Topic

| Field | Description |
| --- | --- |
| `created_at` | Timestamp indicating when the topic was created. |
| `default_subscription` | The default subscription status for the topic. |
| `description` | A description of the topic. |
| `id` | The ID of the topic. |
| `name` | The name of the topic. |
| `object` | The object type. |
| `visibility` | The visibility of the topic. |

Operations: Create, List, Load, Remove, Update.

API path: `/topics`

#### Usage

| Field | Description |
| --- | --- |
| `ai_credits` |  |
| `automation_runs` |  |
| `broadcasts` |  |
| `contacts` |  |
| `domains` |  |
| `emails` |  |
| `object` | The type of object. |
| `rate_limit` |  |
| `segments` |  |

Operations: Load.

API path: `/usage`

#### Webhook

| Field | Description |
| --- | --- |
| `created_at` | Timestamp indicating when the webhook was created. |
| `endpoint` | The URL where webhook events are sent. |
| `events` | Array of event types subscribed to. |
| `id` | The ID of the webhook. |
| `object` | The type of object. |
| `signing_secret` | The secret key used to verify webhook payloads. |
| `status` | The status of the webhook. |

Operations: Create, List, Load, Remove, Update.

API path: `/webhooks/{webhook_id}/signing-secret/rotate`

#### WebhookEvent

| Field | Description |
| --- | --- |
| `created_at` | Timestamp indicating when the event was created. |
| `id` | The ID of the webhook event. |
| `next_attempt_at` | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | The type of object. |
| `payload` | The event payload sent to the webhook endpoint. |
| `status` | The delivery status of the event for this webhook. |
| `type` | The type of the event. |

Operations: Create, List, Load.

API path: `/webhooks/{webhook_id}/events/{event_id}/replay`

#### WebhookEventAttempt

| Field | Description |
| --- | --- |
| `http_status_code` | The HTTP status code returned by the webhook endpoint. |
| `id` | The ID of the webhook event attempt. |
| `response` | The response body returned by the webhook endpoint. |
| `sent_at` | Timestamp indicating when the attempt was sent. |

Operations: List.

API path: `/webhooks/{webhook_id}/events/{event_id}/attempts`



## Entities


### ApiKey

Create an instance: `api_key = client.ApiKey()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | The date and time the API key was created. |
| `domain_id` | `str` | Restrict an API key to send emails only from a specific domain. |
| `id` | `str` | The ID of the API key. |
| `last_used_at` | `str | None` | The date and time the API key was last used. |
| `name` | `str` | The API key name. |
| `object` | `str` | The type of object. |
| `permission` | `str` | The API key can have full access to Resend’s API or be only restricted to send emails. |

#### Example: List

```python
api_keys = client.ApiKey().list()
```

#### Example: Create

```python
api_key = client.ApiKey().create({
    "name": "example_name",  # str
})
```


### Attachment

Create an instance: `attachment = client.Attachment()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_disposition` | `str | None` | How the attachment should be displayed. |
| `content_id` | `str` | The content ID for inline attachments. |
| `content_type` | `str` | The MIME type of the attachment. |
| `download_url` | `str` | Signed URL to download the attachment content. |
| `expires_at` | `str` | Timestamp when the download URL expires. |
| `filename` | `str | None` | The filename of the attachment. |
| `id` | `str` | The ID of the attachment. |
| `object` | `str` | The type of object. |
| `size` | `int` | Size of the attachment in bytes. |

#### Example: Load

```python
attachment = client.Attachment().load({"id": "attachment_id", "email_id": "email_id"})
```

#### Example: List

```python
attachments = client.Attachment().list({"email_id": "example"})
```


### Audience

Create an instance: `audience = client.Audience()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | The date that the object was created. |
| `id` | `str` | The ID of the audience. |
| `name` | `str` | The name of the audience. |
| `object` | `str` | The object of the audience. |

#### Example: Load

```python
audience = client.Audience().load({"id": "audience_id"})
```

#### Example: List

```python
audiences = client.Audience().list()
```

#### Example: Create

```python
audience = client.Audience().create({
    "name": "example_name",  # str
})
```


### Automation

Create an instance: `automation = client.Automation()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `connections` | `list` | The connections between steps in the active version of the automation. |
| `created_at` | `str` | The date and time the automation was created. |
| `id` | `str` | The ID of the automation. |
| `name` | `str` | The name of the automation. |
| `object` | `str` | Type of the response object. |
| `status` | `str` | The current status of the automation. |
| `steps` | `list` | The steps in the active version of the automation. |
| `updated_at` | `str` | The date and time the automation was last updated. |

#### Example: Load

```python
automation = client.Automation().load({"id": "automation_id"})
```

#### Example: List

```python
automations = client.Automation().list()
```

#### Example: Create

```python
automation = client.Automation().create({
    "connections": [],  # list
    "name": "example_name",  # str
    "steps": [],  # list
})
```


### AutomationRun

Create an instance: `automation_run = client.AutomationRun()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `completed_at` | `str | None` | The date and time the run completed. |
| `created_at` | `str` | The date and time the run was created. |
| `id` | `str` | The ID of the automation run. |
| `object` | `str` | Type of the response object. |
| `started_at` | `str | None` | The date and time the run started. |
| `status` | `str` | The current status of the automation run. |
| `steps` | `list` | The steps executed in this run, sorted in graph order. |

#### Example: Load

```python
automation_run = client.AutomationRun().load({"id": "automation_run_id", "automation_id": "automation_id"})
```

#### Example: List

```python
automation_runs = client.AutomationRun().list({"automation_id": "example"})
```


### Broadcast

Create an instance: `broadcast = client.Broadcast()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `str | None` | Deprecated: use `segment_id` instead. |
| `created_at` | `str` | Timestamp indicating when the broadcast was created. |
| `from` | `str` | The email address of the sender. |
| `html` | `str | None` | The HTML version of the broadcast content. |
| `id` | `str` | Unique identifier for the broadcast. |
| `name` | `str` | Name of the broadcast. |
| `object` | `str` | The object type of the response. |
| `preview_text` | `str` | The preview text of the email. |
| `reply_to` | `list` | The email addresses to which replies should be sent. |
| `scheduled_at` | `str` | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `str | None` | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `bool` | Whether to send the broadcast immediately or keep it as a draft. |
| `sent_at` | `str` | Timestamp indicating when the broadcast was sent. |
| `status` | `str` | The status of the broadcast. |
| `subject` | `str` | The subject line of the email. |
| `text` | `str | None` | The plain text version of the broadcast content. |
| `topic_id` | `str | None` | The topic ID that the broadcast is scoped to. |

#### Example: Load

```python
broadcast = client.Broadcast().load({"id": "broadcast_id"})
```

#### Example: List

```python
broadcasts = client.Broadcast().list()
```

#### Example: Create

```python
broadcast = client.Broadcast().create({
    "from": "example_from",  # str
    "segment_id": "example_segment_id",  # str | None
    "subject": "example_subject",  # str
})
```


### BroadcastClickedLink

Create an instance: `broadcast_clicked_link = client.BroadcastClickedLink()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clicks` | `int` | Total number of clicks on this URL. |
| `id` | `str` | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `int` | Number of unique clicks on this URL. |
| `url` | `str` | The URL that was clicked. |

#### Example: List

```python
broadcast_clicked_links = client.BroadcastClickedLink().list({"broadcast_id": "example"})
```


### BroadcastRecipient

Create an instance: `broadcast_recipient = client.BroadcastRecipient()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bounce_type` | `str` | The type of bounce. |
| `clicked_links` | `list` | The links this recipient clicked. |
| `contact_id` | `str | None` | The ID of the contact associated with this recipient, if one exists. |
| `count` | `int` | The number of times this recipient triggered the event. |
| `email` | `str` | The recipient's email address. |
| `id` | `str` | Opaque cursor identifying this row, used for pagination. |

#### Example: List

```python
broadcast_recipients = client.BroadcastRecipient().list({"broadcast_id": "example", "type": "example"})
```


### Contact

Create an instance: `contact = client.Contact()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `str` | Unique identifier of the audience to which the contact belongs. |
| `created_at` | `str` | Timestamp indicating when the contact was created. |
| `email` | `str` | Email address of the contact. |
| `first_name` | `str | None` | First name of the contact. |
| `id` | `str` | Unique identifier for the contact. |
| `last_name` | `str | None` | Last name of the contact. |
| `object` | `str` | Type of the response object. |
| `properties` | `dict` | A map of custom property keys and values. |
| `segments` | `list` | Array of segment IDs to add the contact to. |
| `topics` | `list` | Array of topic subscriptions for the contact. |
| `unsubscribed` | `bool` | Indicates if the contact is unsubscribed. |

#### Example: Load

```python
contact = client.Contact().load({"id": "contact_id"})
```

#### Example: List

```python
contacts = client.Contact().list()
```

#### Example: Create

```python
contact = client.Contact().create({
    "email": "example_email",  # str
})
```


### ContactImport

Create an instance: `contact_import = client.ContactImport()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `column_map` | `str` | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | `str | None` | Timestamp indicating when the contact import completed. |
| `counts` | `dict` |  |
| `created_at` | `str` | Timestamp indicating when the contact import was created. |
| `file` | `str` | CSV file to import. |
| `id` | `str` | Unique identifier for the contact import. |
| `object` | `str` | Type of the response object. |
| `on_conflict` | `str` | Strategy to use when an imported contact already exists. |
| `segments` | `str` | JSON-encoded array of segments to add imported contacts to. |
| `status` | `str` | Current status of the contact import. |
| `topics` | `str` | JSON-encoded array of topic subscriptions to apply to imported contacts. |

#### Example: Load

```python
contact_import = client.ContactImport().load({"id": "contact_import_id"})
```

#### Example: List

```python
contact_imports = client.ContactImport().list()
```

#### Example: Create

```python
contact_import = client.ContactImport().create({
    "file": "example_file",  # str
})
```


### ContactProperty

Create an instance: `contact_property = client.ContactProperty()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Timestamp indicating when the contact property was created. |
| `fallback_value` | `Any` | The default value when the property is not set for a contact. |
| `id` | `str` | The ID of the contact property. |
| `key` | `str` | The property key. |
| `object` | `str` | The object type. |
| `type` | `str` | The property type. |

#### Example: Load

```python
contact_property = client.ContactProperty().load({"id": "contact_property_id"})
```

#### Example: List

```python
contact_propertys = client.ContactProperty().list()
```

#### Example: Create

```python
contact_property = client.ContactProperty().create({
    "key": "example_key",  # str
    "type": "example_type",  # str
})
```


### ContactSegment

Create an instance: `contact_segment = client.ContactSegment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `str` | The ID of the contact. |
| `created_at` | `str` | Timestamp indicating when the contact was added to the segment. |
| `id` | `str` | Unique identifier for the segment. |
| `name` | `str` | Name of the segment. |
| `object` | `str` | The object type. |
| `segment_id` | `str` | The ID of the segment. |

#### Example: List

```python
contact_segments = client.ContactSegment().list({"contact_id": "example"})
```

#### Example: Create

```python
contact_segment = client.ContactSegment().create({
    "contact_id": "example_contact_id",  # str
    "id": "example_id",  # str
})
```


### ContactTopic

Create an instance: `contact_topic = client.ContactTopic()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `str` | The ID of the contact. |
| `description` | `str` | Description of the topic. |
| `id` | `str` | Unique identifier for the topic. |
| `name` | `str` | Name of the topic. |
| `object` | `str` | The object type. |
| `subscription` | `str` | The subscription status for this topic. |
| `topics` | `list` | Array of updated topic subscriptions. |

#### Example: List

```python
contact_topics = client.ContactTopic().list({"contact_id": "example"})
```


### Domain

Create an instance: `domain = client.Domain()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `capabilities` | `dict` | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | Whether click tracking is enabled for this domain. |
| `created_at` | `str` | The date and time the domain was created. |
| `custom_return_path` | `str` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | `str` | The ID of the domain. |
| `name` | `str` | The name of the domain. |
| `object` | `str` | The type of object. |
| `open_tracking` | `bool` | Whether open tracking is enabled for this domain. |
| `records` | `list` |  |
| `region` | `str` | The region where the domain is hosted. |
| `status` | `str` | The status of the domain. |
| `tls` | `str` | TLS mode. |
| `tracking_subdomain` | `str` | The subdomain used for click and open tracking. |

#### Example: Load

```python
domain = client.Domain().load({"id": "domain_id"})
```

#### Example: List

```python
domains = client.Domain().list()
```

#### Example: Create

```python
domain = client.Domain().create({
    "name": "example_name",  # str
})
```


### DomainClaim

Create an instance: `domain_claim = client.DomainClaim()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `blocked_reason` | `str | None` | Why the claim is currently blocked, if applicable. |
| `click_tracking` | `bool` | Track clicks within the body of each HTML email. |
| `created_at` | `str` | The date and time the claim was created. |
| `custom_return_path` | `str` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | `str | None` | The ID of the placeholder domain created for the claim. |
| `expires_at` | `str` | The date and time the claim expires if not verified. |
| `failure_reason` | `str | None` | Why the claim failed, if applicable. |
| `id` | `str` | The ID of the claim. |
| `name` | `str` | The name of the domain being claimed. |
| `object` | `str` | The type of object. |
| `open_tracking` | `bool` | Track the open rate of each email. |
| `record` | `dict` | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | `str | None` | The region where the claimed domain will send from. |
| `status` | `str` | The status of the claim. |
| `tracking_subdomain` | `str` | The subdomain to use for click and open tracking. |

#### Example: Load

```python
domain_claim = client.DomainClaim().load({"id": "domain_claim_id"})
```

#### Example: Create

```python
domain_claim = client.DomainClaim().create({
    "name": "example_name",  # str
})
```


### Email

Create an instance: `email = client.Email()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `list` |  |
| `bcc` | `list` | The email addresses of the blind carbon copy recipients. |
| `cc` | `list` | The email addresses of the carbon copy recipients. |
| `created_at` | `str` | The date and time the email was created. |
| `from` | `str` | The email address of the sender. |
| `headers` | `dict` | Custom headers to add to the email. |
| `html` | `str` | The HTML body of the email. |
| `id` | `str` | The ID of the email. |
| `last_event` | `str` | The status of the email. |
| `message_id` | `str` | The Message-ID header value of the email. |
| `object` | `str` | The type of object. |
| `reply_to` | `list` | The email addresses to which replies should be sent. |
| `scheduled_at` | `str` | Schedule email to be sent later. |
| `subject` | `str` | The subject line of the email. |
| `tags` | `list` |  |
| `template` | `Any` |  |
| `text` | `str` | The plain text body of the email. |
| `to` | `list` | Recipient email address. |
| `topic_id` | `str` | The topic ID to scope the email to. |

#### Example: Load

```python
email = client.Email().load({"id": "email_id"})
```

#### Example: List

```python
emails = client.Email().list()
```

#### Example: Create

```python
email = client.Email().create({
    "from": "example_from",  # str
    "subject": "example_subject",  # str
    "to": [],  # list
})
```


### EmailBatch

Create an instance: `email_batch = client.EmailBatch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `list` |  |
| `send_emails` | `list` |  |

#### Example: Create

```python
email_batch = client.EmailBatch().create({
})
```


### EmailsMetric

Create an instance: `emails_metric = client.EmailsMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `broadcast_id` | `str` | Present when `broadcast` is in `dimensions`. |
| `broadcast_name` | `str` | Present when `broadcast` is in `dimensions`. |
| `domain_id` | `str` | Present when `domain` is in `dimensions`. |
| `domain_name` | `str` | Present when `domain` is in `dimensions`. |
| `email_id` | `str` | Present when `email` is in `dimensions`. |
| `period` | `str` | Present when `period` is in `dimensions`. |

#### Example: List

```python
emails_metrics = client.EmailsMetric().list()
```


### Event

Create an instance: `event = client.Event()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | The date and time the event was created. |
| `id` | `str` | The event ID. |
| `name` | `str` | The event name. |
| `object` | `str` | Type of the response object. |
| `schema` | `dict | None` | A flat key/type map defining the event payload schema. |
| `updated_at` | `str | None` | The date and time the event was last updated. |

#### Example: Load

```python
event = client.Event().load({"id": "event_id"})
```

#### Example: List

```python
events = client.Event().list()
```

#### Example: Create

```python
event = client.Event().create({
    "name": "example_name",  # str
})
```


### Inbox

Create an instance: `inbox = client.Inbox()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | `str` | The ID of the domain that owns the address. |
| `drafts` | `int` | The number of unsent drafts. |
| `email_address` | `str` | The address of the inbox. |
| `forwarding` | `bool` | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | `str | None` | The name recipients see when mail is sent from this inbox. |
| `id` | `str` | The ID of the inbox. |
| `last_received` | `str | None` | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | `str` | Internal name for the inbox. |
| `object` | `str` | The object type. |
| `receiving_address` | `str | None` | The address to forward mail to when forwarding is enabled. |
| `unread` | `int` | The number of unread threads in the Inbox folder. |
| `updated_at` | `str` | ISO 8601 timestamp when the inbox was last updated. |

#### Example: Load

```python
inbox = client.Inbox().load({"id": "inbox_id"})
```

#### Example: List

```python
inboxs = client.Inbox().list()
```

#### Example: Create

```python
inbox = client.Inbox().create({
    "email_address": "example_email_address",  # str
})
```


### InboxAgent

Create an instance: `inbox_agent = client.InboxAgent()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `list` | The actions the agent can take, in the order of the enum. |
| `id` | `str` | The ID of the inbox agent. |
| `instructions` | `str | None` | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | `str` | The object type. |
| `tone` | `str | None` | The tone the agent writes in. |

#### Example: Load

```python
inbox_agent = client.InboxAgent().load({"inbox_id": "inbox_id"})
```


### InboxDraft

Create an instance: `inbox_draft = client.InboxDraft()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `list` | The files stored on the draft. |
| `bcc` | `list` | The BCC recipients. |
| `cc` | `list` | The CC recipients. |
| `created_at` | `str` | Timestamp when the draft was created. |
| `email_id` | `str | None` | The ID of the outbound email queued when the draft was sent. |
| `html` | `str | None` | The HTML body. |
| `id` | `str` | The ID of the draft. |
| `object` | `str` | The object type. |
| `reply_to_email_id` | `str | None` | The Email ID the draft replies to. |
| `snippet` | `str | None` | A short preview of the body. |
| `subject` | `str | None` | The subject. |
| `text` | `str | None` | The plain text body. |
| `thread_id` | `str | None` | The Thread ID when the draft is a reply. |
| `to` | `list | None` | The recipients. |
| `type` | `str` | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `str` | Timestamp when the draft was last saved. |

#### Example: Load

```python
inbox_draft = client.InboxDraft().load({"id": "inbox_draft_id", "inbox_id": "inbox_id"})
```

#### Example: List

```python
inbox_drafts = client.InboxDraft().list({"inbox_id": "example"})
```

#### Example: Create

```python
inbox_draft = client.InboxDraft().create({
    "inbox_id": "example_inbox_id",  # str
})
```


### InboxLabel

Create an instance: `inbox_label = client.InboxLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `str` | The color of the label, as a `#RRGGBB` hex color. |
| `created_at` | `str` | Timestamp when the label was created. |
| `id` | `str` | The ID of the updated label. |
| `name` | `str` | The name of the label. |
| `object` | `str` | Type of the response object. |

#### Example: List

```python
inbox_labels = client.InboxLabel().list({"inbox_id": "example"})
```

#### Example: Create

```python
inbox_label = client.InboxLabel().create({
    "inbox_id": "example_inbox_id",  # str
    "name": "example_name",  # str
})
```


### InboxThread

Create an instance: `inbox_thread = client.InboxThread()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bcc` | `list` | The BCC recipients of the most recent email. |
| `cc` | `list` | The CC recipients of the most recent email. |
| `email_count` | `int` | The number of emails in the thread. |
| `folder` | `str` | The folder the thread is in. |
| `from` | `str | None` | The sender email address. |
| `has_attachment` | `bool` | Whether the most recent email in the thread has an attachment. |
| `has_draft` | `bool` | Whether the thread has an unsent draft. |
| `id` | `str` | The ID of the thread. |
| `labels` | `list` | The labels attached to the thread. |
| `object` | `str` | The object type. |
| `read` | `bool` | True only when every email in the thread is read. |
| `received_at` | `str` | Timestamp when the thread was last active. |
| `subject` | `str | None` | The subject of the thread. |
| `to` | `list` | The recipients of the most recent email. |

#### Example: Load

```python
inbox_thread = client.InboxThread().load({"id": "inbox_thread_id", "inbox_id": "inbox_id"})
```

#### Example: List

```python
inbox_threads = client.InboxThread().list({"inbox_id": "example"})
```


### InboxThreadEmail

Create an instance: `inbox_thread_email = client.InboxThreadEmail()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `list` | The attachments of the message. |
| `bcc` | `list` | The BCC recipients of the message. |
| `cc` | `list` | The CC recipients of the message. |
| `direction` | `str` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `str` | Sender email address. |
| `html` | `str | None` | The HTML body. |
| `id` | `str` | The ID of the message. |
| `message_id` | `str | None` | The Message-ID header of the message. |
| `read` | `bool` | Whether the message has been read. |
| `received_at` | `str` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `list` | The Reply-To addresses of the message. |
| `subject` | `str | None` | The subject of the message. |
| `text` | `str | None` | The plain-text body. |
| `to` | `list` | The recipients of the message. |

#### Example: Load

```python
inbox_thread_email = client.InboxThreadEmail().load({"id": "inbox_thread_email_id", "inbox_id": "inbox_id", "thread_id": "thread_id"})
```

#### Example: List

```python
inbox_thread_emails = client.InboxThreadEmail().list({"inbox_id": "example", "thread_id": "example"})
```


### InboxThreadLabel

Create an instance: `inbox_thread_label = client.InboxThreadLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `list` | Every label on the thread after the change. |
| `id` | `str` |  |
| `labels` | `list` | The IDs of the labels to apply to the thread. |
| `object` | `str` | Type of the response object. |

#### Example: Create

```python
inbox_thread_label = client.InboxThreadLabel().create({
    "inbox_id": "example_inbox_id",  # str
    "thread_id": "example_thread_id",  # str
    "labels": [],  # list
})
```


### Log

Create an instance: `log = client.Log()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | The date the log was created. |
| `endpoint` | `str` | The API endpoint that was called. |
| `id` | `str` | The log ID. |
| `method` | `str` | The HTTP method used. |
| `object` | `str` | Type of the response object. |
| `request_body` | `dict | None` | The request body sent to the API. |
| `response_body` | `dict | None` | The response body returned by the API. |
| `response_status` | `int` | The HTTP status code of the response. |
| `user_agent` | `str | None` | The user agent of the request. |

#### Example: Load

```python
log = client.Log().load({"id": "log_id"})
```

#### Example: List

```python
logs = client.Log().list()
```


### OauthGrant

Create an instance: `oauth_grant = client.OauthGrant()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client` | `dict` | The OAuth client the grant was issued to. |
| `client_id` | `str` | The ID of the OAuth client the grant was issued to. |
| `created_at` | `str` | The date and time the OAuth grant was created. |
| `id` | `str` | The ID of the OAuth grant. |
| `revoked_at` | `str | None` | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | `str | None` | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | `list` | The scopes granted to the OAuth client. |

#### Example: List

```python
oauth_grants = client.OauthGrant().list()
```


### ReceivedAttachment

Create an instance: `received_attachment = client.ReceivedAttachment()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_disposition` | `str | None` | How the attachment should be displayed. |
| `content_id` | `str` | The content ID for inline attachments. |
| `content_type` | `str` | The MIME type of the attachment. |
| `download_url` | `str` | Signed URL to download the attachment content. |
| `expires_at` | `str` | Timestamp when the download URL expires. |
| `filename` | `str | None` | The filename of the attachment. |
| `id` | `str` | The ID of the attachment. |
| `object` | `str` | The type of object. |
| `size` | `int` | Size of the attachment in bytes. |

#### Example: Load

```python
received_attachment = client.ReceivedAttachment().load({"id": "received_attachment_id", "email_id": "email_id"})
```

#### Example: List

```python
received_attachments = client.ReceivedAttachment().list({"email_id": "example"})
```


### ReceivedEmail

Create an instance: `received_email = client.ReceivedEmail()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `list` | Array of attachments. |
| `bcc` | `list | None` | The BCC recipients. |
| `cc` | `list | None` | The CC recipients. |
| `created_at` | `str` | Timestamp when the email was received. |
| `from` | `str` | The sender email address. |
| `headers` | `dict | None` | The email headers. |
| `html` | `str | None` | The HTML content of the email. |
| `id` | `str` | The ID of the received email. |
| `message_id` | `str` | The unique message ID from the email headers. |
| `object` | `str` | The type of object. |
| `received_for` | `list` | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | `list | None` | The reply-to addresses. |
| `subject` | `str` | The email subject. |
| `text` | `str | None` | The plain text content of the email. |
| `to` | `list` | The recipient email addresses. |

#### Example: Load

```python
received_email = client.ReceivedEmail().load({"email_id": "email_id"})
```

#### Example: List

```python
received_emails = client.ReceivedEmail().list()
```


### Segment

Create an instance: `segment = client.Segment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `str` | The ID of the audience this segment belongs to. |
| `created_at` | `str` | Timestamp indicating when the segment was created. |
| `filter` | `dict` | Filter conditions for the segment. |
| `id` | `str` | The ID of the segment. |
| `name` | `str` | The name of the segment. |
| `object` | `str` | The object type. |

#### Example: Load

```python
segment = client.Segment().load({"id": "segment_id"})
```

#### Example: List

```python
segments = client.Segment().list()
```

#### Example: Create

```python
segment = client.Segment().create({
    "name": "example_name",  # str
})
```


### SegmentContact

Create an instance: `segment_contact = client.SegmentContact()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Timestamp indicating when the contact was created. |
| `email` | `str` | Email address of the contact. |
| `first_name` | `str | None` | First name of the contact. |
| `id` | `str` | Unique identifier for the contact. |
| `last_name` | `str | None` | Last name of the contact. |
| `unsubscribed` | `bool` | Indicates if the contact is unsubscribed. |

#### Example: List

```python
segment_contacts = client.SegmentContact().list({"segment_id": "example"})
```


### Suppression

Create an instance: `suppression = client.Suppression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Timestamp indicating when the suppression was created. |
| `email` | `str` | Email address that is suppressed. |
| `id` | `str` | Unique identifier for the suppression. |
| `object` | `str` | Type of the response object. |
| `origin` | `str` | Origin of the suppression. |
| `source_id` | `str | None` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

#### Example: Load

```python
suppression = client.Suppression().load({"id": "suppression_id"})
```

#### Example: List

```python
suppressions = client.Suppression().list()
```

#### Example: Create

```python
suppression = client.Suppression().create({
    "email": "example_email",  # str
})
```


### SuppressionBatchAdd

Create an instance: `suppression_batch_add = client.SuppressionBatchAdd()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `emails` | `list` | Email addresses to suppress. |

#### Example: Create

```python
suppression_batch_add = client.SuppressionBatchAdd().create({
    "emails": [],  # list
})
```


### SuppressionBatchRemove

Create an instance: `suppression_batch_remove = client.SuppressionBatchRemove()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `list` | Array containing the removed suppressions. |
| `emails` | `list` | Email addresses to remove from the suppression list. |
| `ids` | `list` | Suppression IDs to remove from the suppression list. |

#### Example: Create

```python
suppression_batch_remove = client.SuppressionBatchRemove().create({
})
```


### Template

Create an instance: `template = client.Template()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alias` | `str` | The alias of the template. |
| `created_at` | `str` | Timestamp indicating when the template was created. |
| `current_version_id` | `str` | The ID of the current version of the template. |
| `from` | `str` | Sender email address. |
| `has_unpublished_versions` | `bool` | Indicates whether the template has unpublished versions. |
| `html` | `str` | The HTML version of the template. |
| `id` | `str` | The ID of the template. |
| `name` | `str` | The name of the template. |
| `object` | `str` | The type of object. |
| `published_at` | `str | None` | Timestamp indicating when the template was published. |
| `reply_to` | `list | None` | Reply-to email addresses. |
| `status` | `str` | The publication status of the template. |
| `subject` | `str` | Email subject. |
| `text` | `str` | The plain text version of the template. |
| `updated_at` | `str` | Timestamp indicating when the template was last updated. |
| `variables` | `list` |  |

#### Example: Load

```python
template = client.Template().load({"id": "template_id"})
```

#### Example: List

```python
templates = client.Template().list()
```

#### Example: Create

```python
template = client.Template().create({
    "html": "example_html",  # str
    "name": "example_name",  # str
})
```


### Topic

Create an instance: `topic = client.Topic()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Timestamp indicating when the topic was created. |
| `default_subscription` | `str` | The default subscription status for the topic. |
| `description` | `str` | A description of the topic. |
| `id` | `str` | The ID of the topic. |
| `name` | `str` | The name of the topic. |
| `object` | `str` | The object type. |
| `visibility` | `str` | The visibility of the topic. |

#### Example: Load

```python
topic = client.Topic().load({"id": "topic_id"})
```

#### Example: List

```python
topics = client.Topic().list()
```

#### Example: Create

```python
topic = client.Topic().create({
    "default_subscription": "example_default_subscription",  # str
    "name": "example_name",  # str
})
```


### Usage

Create an instance: `usage = client.Usage()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ai_credits` | `dict` |  |
| `automation_runs` | `dict` |  |
| `broadcasts` | `dict` |  |
| `contacts` | `dict` |  |
| `domains` | `dict` |  |
| `emails` | `dict` |  |
| `object` | `str` | The type of object. |
| `rate_limit` | `dict` |  |
| `segments` | `dict` |  |

#### Example: Load

```python
usage = client.Usage().load()
```


### Webhook

Create an instance: `webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Timestamp indicating when the webhook was created. |
| `endpoint` | `str` | The URL where webhook events are sent. |
| `events` | `list | None` | Array of event types subscribed to. |
| `id` | `str` | The ID of the webhook. |
| `object` | `str` | The type of object. |
| `signing_secret` | `str` | The secret key used to verify webhook payloads. |
| `status` | `str` | The status of the webhook. |

#### Example: Load

```python
webhook = client.Webhook().load({"id": "webhook_id"})
```

#### Example: List

```python
webhooks = client.Webhook().list()
```

#### Example: Create

```python
webhook = client.Webhook().create({
    "endpoint": "example_endpoint",  # str
    "events": [],  # list | None
})
```


### WebhookEvent

Create an instance: `webhook_event = client.WebhookEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` | Timestamp indicating when the event was created. |
| `id` | `str` | The ID of the webhook event. |
| `next_attempt_at` | `str | None` | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | `str` | The type of object. |
| `payload` | `dict` | The event payload sent to the webhook endpoint. |
| `status` | `str` | The delivery status of the event for this webhook. |
| `type` | `str` | The type of the event. |

#### Example: Load

```python
webhook_event = client.WebhookEvent().load({"id": "webhook_event_id", "webhook_id": "webhook_id"})
```

#### Example: List

```python
webhook_events = client.WebhookEvent().list({"webhook_id": "example"})
```

#### Example: Create

```python
webhook_event = client.WebhookEvent().create({
    "event_id": "example_event_id",  # str
    "webhook_id": "example_webhook_id",  # str
})
```


### WebhookEventAttempt

Create an instance: `webhook_event_attempt = client.WebhookEventAttempt()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `http_status_code` | `int` | The HTTP status code returned by the webhook endpoint. |
| `id` | `str` | The ID of the webhook event attempt. |
| `response` | `str` | The response body returned by the webhook endpoint. |
| `sent_at` | `str` | Timestamp indicating when the attempt was sent. |

#### Example: List

```python
webhook_event_attempts = client.WebhookEventAttempt().list({"event_id": "example", "webhook_id": "example"})
```

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Debug capture |
| [`idempotency`](#idempotency) | Idempotency |
| [`metrics`](#metrics) | Metrics |
| [`paging`](#paging) | Paging |
| [`ratelimit`](#ratelimit) | Rate limiting |
| [`retry`](#retry) | Retry |
| [`test`](#test) | Test transport |
| [`timeout`](#timeout) | Timeout |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Debug capture.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Metrics.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Paging.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Rate limiting.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Retry.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Timeout.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Open types

1 field is carried as open values rather than typed structures.
This follows from the API definition, not from a gap in this SDK: the
definition describes it with untagged unions —
`oneOf`/`anyOf` branches with no `discriminator` — so it never states which
variant a given value is. Nothing can select a branch reliably, so the SDK
passes the value through unchanged rather than assert a shape the API does not
guarantee.

| Entity | Field | Variants | Nesting |
| --- | --- | --- | --- |
| `template` | `variables` | 5 | 3 levels |

These values round-trip unchanged — read them, modify them, send them back. If
the API adds a `discriminator` to the definition, regenerating will type them.
Every other field is typed normally.

## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **DebugFeature**: Debug capture
- **IdempotencyFeature**: Idempotency
- **MetricsFeature**: Metrics
- **PagingFeature**: Paging
- **RatelimitFeature**: Rate limiting
- **RetryFeature**: Retry
- **TestFeature**: Test transport
- **TimeoutFeature**: Timeout

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── resend_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── schema.py                    -- Generated option + entity specs
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`resend_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
oauthgrant = client.OauthGrant()
oauthgrant.list()

# oauthgrant.data_get() now returns the oauthgrant data from the last list
# oauthgrant.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
