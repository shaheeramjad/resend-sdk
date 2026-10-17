# Resend TypeScript SDK

Send email, manage domains, contacts and broadcasts on Resend from TypeScript, Python and Go.

The TypeScript SDK for the Resend API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.ApiKey()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `py` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`, see [Tags](https://github.com/shaheeramjad/resend-sdk/tags)), or from a
clone, which carries the compiled `dist/`:

```bash
git clone https://github.com/shaheeramjad/resend-sdk
npm install ./resend-sdk/ts
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { ResendSDK } from '@shaheeramjad/resend-sdk'

const client = new ResendSDK({
  apikey: process.env.RESEND_APIKEY,
})
```

### 2. List apikey records

`list()` resolves to an array of ApiKey ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const apikeys = await client.ApiKey().list()

for (const apikey of apikeys) {
  console.log(apikey.data())
}
```

### 3. Load an attachment

Attachment is nested under email, so provide the `email_id`.
`load()` returns the entity and throws on failure; `.data()` reads its record:

```ts
try {
  const attachment = await client.Attachment().load({
    email_id: 'example_email_id',
    id: 'example_id',
  })
  console.log(attachment.data())
} catch (err) {
  console.error('load failed:', err)
}
```

### 4. Create, update, and remove

```ts
// Create — returns the created ApiKey ENTITY (.data() for the record)
const created = await client.ApiKey().create({
  name: 'example_name',
})

// Update — the id comes off the returned entity's data()
const updated = await client.ApiKey().update({
  id: created.data().id!,
  name: 'example_name',
})

// Remove
await client.ApiKey().remove({
  id: created.data().id!,
})
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const oauthgrants = await client.OauthGrant().list()
  console.log(oauthgrants.map((item) => item.data()))
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
result envelope. Branch on `ok`; on failure `status` holds the HTTP status
(for error responses) and `err` holds the error:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (!result.ok) {
  console.error('request failed:', result.status, result.err)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = ResendSDK.test()

const oauthgrants = await client.OauthGrant().list()
// oauthgrants is an array of OauthGrant entities, one per mock record
console.log(oauthgrants.map((oauthgrant) => oauthgrant.data()))
```

You can also use the instance method:

```ts
const client = new ResendSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.OauthGrant()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new ResendSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```

Live entity tests continue independent operations after errors and attempt
supported cleanup. Their final result reports failures and missing prerequisites
after the remaining work completes. The model and test inputs determine which
API operations the generated scenarios cover.


## Reference

### ResendSDK

#### Constructor

```ts
new ResendSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `ApiKey(data?)` | `ApiKeyEntity` | Create an ApiKey entity instance. |
| `Attachment(data?)` | `AttachmentEntity` | Create an Attachment entity instance. |
| `Audience(data?)` | `AudienceEntity` | Create an Audience entity instance. |
| `Automation(data?)` | `AutomationEntity` | Create an Automation entity instance. |
| `AutomationRun(data?)` | `AutomationRunEntity` | Create an AutomationRun entity instance. |
| `Broadcast(data?)` | `BroadcastEntity` | Create a Broadcast entity instance. |
| `BroadcastClickedLink(data?)` | `BroadcastClickedLinkEntity` | Create a BroadcastClickedLink entity instance. |
| `BroadcastRecipient(data?)` | `BroadcastRecipientEntity` | Create a BroadcastRecipient entity instance. |
| `Contact(data?)` | `ContactEntity` | Create a Contact entity instance. |
| `ContactImport(data?)` | `ContactImportEntity` | Create a ContactImport entity instance. |
| `ContactProperty(data?)` | `ContactPropertyEntity` | Create a ContactProperty entity instance. |
| `ContactSegment(data?)` | `ContactSegmentEntity` | Create a ContactSegment entity instance. |
| `ContactTopic(data?)` | `ContactTopicEntity` | Create a ContactTopic entity instance. |
| `Domain(data?)` | `DomainEntity` | Create a Domain entity instance. |
| `DomainClaim(data?)` | `DomainClaimEntity` | Create a DomainClaim entity instance. |
| `Email(data?)` | `EmailEntity` | Create an Email entity instance. |
| `EmailBatch(data?)` | `EmailBatchEntity` | Create an EmailBatch entity instance. |
| `EmailsMetric(data?)` | `EmailsMetricEntity` | Create an EmailsMetric entity instance. |
| `Event(data?)` | `EventEntity` | Create an Event entity instance. |
| `Inbox(data?)` | `InboxEntity` | Create an Inbox entity instance. |
| `InboxAgent(data?)` | `InboxAgentEntity` | Create an InboxAgent entity instance. |
| `InboxDraft(data?)` | `InboxDraftEntity` | Create an InboxDraft entity instance. |
| `InboxLabel(data?)` | `InboxLabelEntity` | Create an InboxLabel entity instance. |
| `InboxThread(data?)` | `InboxThreadEntity` | Create an InboxThread entity instance. |
| `InboxThreadEmail(data?)` | `InboxThreadEmailEntity` | Create an InboxThreadEmail entity instance. |
| `InboxThreadLabel(data?)` | `InboxThreadLabelEntity` | Create an InboxThreadLabel entity instance. |
| `Log(data?)` | `LogEntity` | Create a Log entity instance. |
| `OauthGrant(data?)` | `OauthGrantEntity` | Create an OauthGrant entity instance. |
| `ReceivedAttachment(data?)` | `ReceivedAttachmentEntity` | Create a ReceivedAttachment entity instance. |
| `ReceivedEmail(data?)` | `ReceivedEmailEntity` | Create a ReceivedEmail entity instance. |
| `Segment(data?)` | `SegmentEntity` | Create a Segment entity instance. |
| `SegmentContact(data?)` | `SegmentContactEntity` | Create a SegmentContact entity instance. |
| `Suppression(data?)` | `SuppressionEntity` | Create a Suppression entity instance. |
| `SuppressionBatchAdd(data?)` | `SuppressionBatchAddEntity` | Create a SuppressionBatchAdd entity instance. |
| `SuppressionBatchRemove(data?)` | `SuppressionBatchRemoveEntity` | Create a SuppressionBatchRemove entity instance. |
| `Template(data?)` | `TemplateEntity` | Create a Template entity instance. |
| `Topic(data?)` | `TopicEntity` | Create a Topic entity instance. |
| `Usage(data?)` | `UsageEntity` | Create an Usage entity instance. |
| `Webhook(data?)` | `WebhookEntity` | Create a Webhook entity instance. |
| `WebhookEvent(data?)` | `WebhookEventEntity` | Create a WebhookEvent entity instance. |
| `WebhookEventAttempt(data?)` | `WebhookEventAttemptEntity` | Create a WebhookEventAttempt entity instance. |
| `tester(testopts?, sdkopts?)` | `ResendSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `ResendSDK.test(testopts?, sdkopts?)` | `ResendSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria, and return it. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria, one per record. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity, and return it. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity, and return it. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<Entity>` | Remove an entity, and return it marked as deleted. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): ResendSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity itself — there is no result
envelope, and an entity's `data()` reads its record:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to the entity, marked as deleted.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: create, list, remove, update.

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

Operations: list, load.

API path: `/emails/{email_id}/attachments`

#### Audience

| Field | Description |
| --- | --- |
| `created_at` | The date that the object was created. |
| `id` | The ID of the audience. |
| `name` | The name of the audience. |
| `object` | The object of the audience. |

Operations: create, list, load, remove.

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

Operations: create, list, load, remove, update.

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

Operations: list, load.

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

Operations: create, list, load, remove, update.

API path: `/broadcasts/{id}/cancel`

#### BroadcastClickedLink

| Field | Description |
| --- | --- |
| `clicks` | Total number of clicks on this URL. |
| `id` | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | Number of unique clicks on this URL. |
| `url` | The URL that was clicked. |

Operations: list.

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

Operations: list.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, load.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, remove.

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

Operations: list, update.

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

Operations: create, list, load, remove, update.

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

Operations: create, load.

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

Operations: create, list, load, update.

API path: `/emails/{email_id}/cancel`

#### EmailBatch

| Field | Description |
| --- | --- |
| `data` |  |
| `send_emails` |  |

Operations: create.

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

Operations: list.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, load, remove, update.

API path: `/inboxes`

#### InboxAgent

| Field | Description |
| --- | --- |
| `actions` | The actions the agent can take, in the order of the enum. |
| `id` | The ID of the inbox agent. |
| `instructions` | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | The object type. |
| `tone` | The tone the agent writes in. |

Operations: load, update.

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

Operations: create, list, load, remove, update.

API path: `/inboxes/{inbox_id}/drafts/{draft_id}/send`

#### InboxLabel

| Field | Description |
| --- | --- |
| `color` | The color of the label, as a `#RRGGBB` hex color. |
| `created_at` | Timestamp when the label was created. |
| `id` | The ID of the updated label. |
| `name` | The name of the label. |
| `object` | Type of the response object. |

Operations: create, list, remove, update.

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

Operations: list, load, remove, update.

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

Operations: create, list, load.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/forward`

#### InboxThreadLabel

| Field | Description |
| --- | --- |
| `data` | Every label on the thread after the change. |
| `id` |  |
| `labels` | The IDs of the labels to apply to the thread. |
| `object` | Type of the response object. |

Operations: create, remove, update.

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

Operations: list, load.

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

Operations: list, remove.

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

Operations: list, load.

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

Operations: list, load.

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

Operations: create, list, load, remove, update.

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

Operations: list.

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

Operations: create, list, load, remove.

API path: `/suppressions`

#### SuppressionBatchAdd

| Field | Description |
| --- | --- |
| `emails` | Email addresses to suppress. |

Operations: create.

API path: `/suppressions/batch/add`

#### SuppressionBatchRemove

| Field | Description |
| --- | --- |
| `data` | Array containing the removed suppressions. |
| `emails` | Email addresses to remove from the suppression list. |
| `ids` | Suppression IDs to remove from the suppression list. |

Operations: create.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, load, remove, update.

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

Operations: load.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, load.

API path: `/webhooks/{webhook_id}/events/{event_id}/replay`

#### WebhookEventAttempt

| Field | Description |
| --- | --- |
| `http_status_code` | The HTTP status code returned by the webhook endpoint. |
| `id` | The ID of the webhook event attempt. |
| `response` | The response body returned by the webhook endpoint. |
| `sent_at` | Timestamp indicating when the attempt was sent. |

Operations: list.

API path: `/webhooks/{webhook_id}/events/{event_id}/attempts`



## Entities


### ApiKey

Create an instance: `const api_key = client.ApiKey()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date and time the API key was created. |
| `domain_id` | `string` | Restrict an API key to send emails only from a specific domain. |
| `id` | `string` | The ID of the API key. |
| `last_used_at` | `string | null` | The date and time the API key was last used. |
| `name` | `string` | The API key name. |
| `object` | `string` | The type of object. |
| `permission` | `string` | The API key can have full access to Resend’s API or be only restricted to send emails. |

#### Example: List

```ts
const api_keys = await client.ApiKey().list()
```

#### Example: Create

```ts
const api_key = await client.ApiKey().create({
  name: 'example_name',
})
```


### Attachment

Create an instance: `const attachment = client.Attachment()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_disposition` | `string | null` | How the attachment should be displayed. |
| `content_id` | `string` | The content ID for inline attachments. |
| `content_type` | `string` | The MIME type of the attachment. |
| `download_url` | `string` | Signed URL to download the attachment content. |
| `expires_at` | `string` | Timestamp when the download URL expires. |
| `filename` | `string | null` | The filename of the attachment. |
| `id` | `string` | The ID of the attachment. |
| `object` | `string` | The type of object. |
| `size` | `number` | Size of the attachment in bytes. |

#### Example: Load

```ts
const attachment = await client.Attachment().load({ id: 'attachment_id', email_id: 'email_id' })
```

#### Example: List

```ts
const attachments = await client.Attachment().list({ email_id: "example" })
```


### Audience

Create an instance: `const audience = client.Audience()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date that the object was created. |
| `id` | `string` | The ID of the audience. |
| `name` | `string` | The name of the audience. |
| `object` | `string` | The object of the audience. |

#### Example: Load

```ts
const audience = await client.Audience().load({ id: 'audience_id' })
```

#### Example: List

```ts
const audiences = await client.Audience().list()
```

#### Example: Create

```ts
const audience = await client.Audience().create({
  name: 'example_name',
})
```


### Automation

Create an instance: `const automation = client.Automation()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `connections` | `any[]` | The connections between steps in the active version of the automation. |
| `created_at` | `string` | The date and time the automation was created. |
| `id` | `string` | The ID of the automation. |
| `name` | `string` | The name of the automation. |
| `object` | `string` | Type of the response object. |
| `status` | `string` | The current status of the automation. |
| `steps` | `any[]` | The steps in the active version of the automation. |
| `updated_at` | `string` | The date and time the automation was last updated. |

#### Example: Load

```ts
const automation = await client.Automation().load({ id: 'automation_id' })
```

#### Example: List

```ts
const automations = await client.Automation().list()
```

#### Example: Create

```ts
const automation = await client.Automation().create({
  connections: [],
  name: 'example_name',
  steps: [],
})
```


### AutomationRun

Create an instance: `const automation_run = client.AutomationRun()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `completed_at` | `string | null` | The date and time the run completed. |
| `created_at` | `string` | The date and time the run was created. |
| `id` | `string` | The ID of the automation run. |
| `object` | `string` | Type of the response object. |
| `started_at` | `string | null` | The date and time the run started. |
| `status` | `string` | The current status of the automation run. |
| `steps` | `any[]` | The steps executed in this run, sorted in graph order. |

#### Example: Load

```ts
const automation_run = await client.AutomationRun().load({ id: 'automation_run_id', automation_id: 'automation_id' })
```

#### Example: List

```ts
const automation_runs = await client.AutomationRun().list({ automation_id: "example" })
```


### Broadcast

Create an instance: `const broadcast = client.Broadcast()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string | null` | Deprecated: use `segment_id` instead. |
| `created_at` | `string` | Timestamp indicating when the broadcast was created. |
| `from` | `string` | The email address of the sender. |
| `html` | `string | null` | The HTML version of the broadcast content. |
| `id` | `string` | Unique identifier for the broadcast. |
| `name` | `string` | Name of the broadcast. |
| `object` | `string` | The object type of the response. |
| `preview_text` | `string` | The preview text of the email. |
| `reply_to` | `any[]` | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `string | null` | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `boolean` | Whether to send the broadcast immediately or keep it as a draft. |
| `sent_at` | `string` | Timestamp indicating when the broadcast was sent. |
| `status` | `string` | The status of the broadcast. |
| `subject` | `string` | The subject line of the email. |
| `text` | `string | null` | The plain text version of the broadcast content. |
| `topic_id` | `string | null` | The topic ID that the broadcast is scoped to. |

#### Example: Load

```ts
const broadcast = await client.Broadcast().load({ id: 'broadcast_id' })
```

#### Example: List

```ts
const broadcasts = await client.Broadcast().list()
```

#### Example: Create

```ts
const broadcast = await client.Broadcast().create({
  from: 'example_from',
  segment_id: 'example_segment_id',
  subject: 'example_subject',
})
```


### BroadcastClickedLink

Create an instance: `const broadcast_clicked_link = client.BroadcastClickedLink()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clicks` | `number` | Total number of clicks on this URL. |
| `id` | `string` | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `number` | Number of unique clicks on this URL. |
| `url` | `string` | The URL that was clicked. |

#### Example: List

```ts
const broadcast_clicked_links = await client.BroadcastClickedLink().list({ broadcast_id: "example" })
```


### BroadcastRecipient

Create an instance: `const broadcast_recipient = client.BroadcastRecipient()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bounce_type` | `string` | The type of bounce. |
| `clicked_links` | `any[]` | The links this recipient clicked. |
| `contact_id` | `string | null` | The ID of the contact associated with this recipient, if one exists. |
| `count` | `number` | The number of times this recipient triggered the event. |
| `email` | `string` | The recipient's email address. |
| `id` | `string` | Opaque cursor identifying this row, used for pagination. |

#### Example: List

```ts
const broadcast_recipients = await client.BroadcastRecipient().list({ broadcast_id: "example", type: "example" })
```


### Contact

Create an instance: `const contact = client.Contact()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | Unique identifier of the audience to which the contact belongs. |
| `created_at` | `string` | Timestamp indicating when the contact was created. |
| `email` | `string` | Email address of the contact. |
| `first_name` | `string | null` | First name of the contact. |
| `id` | `string` | Unique identifier for the contact. |
| `last_name` | `string | null` | Last name of the contact. |
| `object` | `string` | Type of the response object. |
| `properties` | `Record<string, any>` | A map of custom property keys and values. |
| `segments` | `any[]` | Array of segment IDs to add the contact to. |
| `topics` | `any[]` | Array of topic subscriptions for the contact. |
| `unsubscribed` | `boolean` | Indicates if the contact is unsubscribed. |

#### Example: Load

```ts
const contact = await client.Contact().load({ id: 'contact_id' })
```

#### Example: List

```ts
const contacts = await client.Contact().list()
```

#### Example: Create

```ts
const contact = await client.Contact().create({
  email: 'example_email',
})
```


### ContactImport

Create an instance: `const contact_import = client.ContactImport()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `column_map` | `string` | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | `string | null` | Timestamp indicating when the contact import completed. |
| `counts` | `Record<string, any>` |  |
| `created_at` | `string` | Timestamp indicating when the contact import was created. |
| `file` | `string` | CSV file to import. |
| `id` | `string` | Unique identifier for the contact import. |
| `object` | `string` | Type of the response object. |
| `on_conflict` | `string` | Strategy to use when an imported contact already exists. |
| `segments` | `string` | JSON-encoded array of segments to add imported contacts to. |
| `status` | `string` | Current status of the contact import. |
| `topics` | `string` | JSON-encoded array of topic subscriptions to apply to imported contacts. |

#### Example: Load

```ts
const contact_import = await client.ContactImport().load({ id: 'contact_import_id' })
```

#### Example: List

```ts
const contact_imports = await client.ContactImport().list()
```

#### Example: Create

```ts
const contact_import = await client.ContactImport().create({
  file: 'example_file',
})
```


### ContactProperty

Create an instance: `const contact_property = client.ContactProperty()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the contact property was created. |
| `fallback_value` | `any` | The default value when the property is not set for a contact. |
| `id` | `string` | The ID of the contact property. |
| `key` | `string` | The property key. |
| `object` | `string` | The object type. |
| `type` | `string` | The property type. |

#### Example: Load

```ts
const contact_property = await client.ContactProperty().load({ id: 'contact_property_id' })
```

#### Example: List

```ts
const contact_propertys = await client.ContactProperty().list()
```

#### Example: Create

```ts
const contact_property = await client.ContactProperty().create({
  key: 'example_key',
  type: 'example_type',
})
```


### ContactSegment

Create an instance: `const contact_segment = client.ContactSegment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` | The ID of the contact. |
| `created_at` | `string` | Timestamp indicating when the contact was added to the segment. |
| `id` | `string` | Unique identifier for the segment. |
| `name` | `string` | Name of the segment. |
| `object` | `string` | The object type. |
| `segment_id` | `string` | The ID of the segment. |

#### Example: List

```ts
const contact_segments = await client.ContactSegment().list({ contact_id: "example" })
```

#### Example: Create

```ts
const contact_segment = await client.ContactSegment().create({
  contact_id: 'example_contact_id',
  id: 'example_id',
})
```


### ContactTopic

Create an instance: `const contact_topic = client.ContactTopic()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` | The ID of the contact. |
| `description` | `string` | Description of the topic. |
| `id` | `string` | Unique identifier for the topic. |
| `name` | `string` | Name of the topic. |
| `object` | `string` | The object type. |
| `subscription` | `string` | The subscription status for this topic. |
| `topics` | `any[]` | Array of updated topic subscriptions. |

#### Example: List

```ts
const contact_topics = await client.ContactTopic().list({ contact_id: "example" })
```


### Domain

Create an instance: `const domain = client.Domain()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `capabilities` | `Record<string, any>` | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `boolean` | Whether click tracking is enabled for this domain. |
| `created_at` | `string` | The date and time the domain was created. |
| `custom_return_path` | `string` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | `string` | The ID of the domain. |
| `name` | `string` | The name of the domain. |
| `object` | `string` | The type of object. |
| `open_tracking` | `boolean` | Whether open tracking is enabled for this domain. |
| `records` | `any[]` |  |
| `region` | `string` | The region where the domain is hosted. |
| `status` | `string` | The status of the domain. |
| `tls` | `string` | TLS mode. |
| `tracking_subdomain` | `string` | The subdomain used for click and open tracking. |

#### Example: Load

```ts
const domain = await client.Domain().load({ id: 'domain_id' })
```

#### Example: List

```ts
const domains = await client.Domain().list()
```

#### Example: Create

```ts
const domain = await client.Domain().create({
  name: 'example_name',
})
```


### DomainClaim

Create an instance: `const domain_claim = client.DomainClaim()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `blocked_reason` | `string | null` | Why the claim is currently blocked, if applicable. |
| `click_tracking` | `boolean` | Track clicks within the body of each HTML email. |
| `created_at` | `string` | The date and time the claim was created. |
| `custom_return_path` | `string` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | `string | null` | The ID of the placeholder domain created for the claim. |
| `expires_at` | `string` | The date and time the claim expires if not verified. |
| `failure_reason` | `string | null` | Why the claim failed, if applicable. |
| `id` | `string` | The ID of the claim. |
| `name` | `string` | The name of the domain being claimed. |
| `object` | `string` | The type of object. |
| `open_tracking` | `boolean` | Track the open rate of each email. |
| `record` | `Record<string, any>` | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | `string | null` | The region where the claimed domain will send from. |
| `status` | `string` | The status of the claim. |
| `tracking_subdomain` | `string` | The subdomain to use for click and open tracking. |

#### Example: Load

```ts
const domain_claim = await client.DomainClaim().load({ id: 'domain_claim_id' })
```

#### Example: Create

```ts
const domain_claim = await client.DomainClaim().create({
  name: 'example_name',
})
```


### Email

Create an instance: `const email = client.Email()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any[]` |  |
| `bcc` | `any[]` | The email addresses of the blind carbon copy recipients. |
| `cc` | `any[]` | The email addresses of the carbon copy recipients. |
| `created_at` | `string` | The date and time the email was created. |
| `from` | `string` | The email address of the sender. |
| `headers` | `Record<string, any>` | Custom headers to add to the email. |
| `html` | `string` | The HTML body of the email. |
| `id` | `string` | The ID of the email. |
| `last_event` | `string` | The status of the email. |
| `message_id` | `string` | The Message-ID header value of the email. |
| `object` | `string` | The type of object. |
| `reply_to` | `any[]` | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | Schedule email to be sent later. |
| `subject` | `string` | The subject line of the email. |
| `tags` | `any[]` |  |
| `template` | `any` |  |
| `text` | `string` | The plain text body of the email. |
| `to` | `any[]` | Recipient email address. |
| `topic_id` | `string` | The topic ID to scope the email to. |

#### Example: Load

```ts
const email = await client.Email().load({ id: 'email_id' })
```

#### Example: List

```ts
const emails = await client.Email().list()
```

#### Example: Create

```ts
const email = await client.Email().create({
  from: 'example_from',
  subject: 'example_subject',
  to: [],
})
```


### EmailBatch

Create an instance: `const email_batch = client.EmailBatch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `any[]` |  |
| `send_emails` | `any[]` |  |

#### Example: Create

```ts
const email_batch = await client.EmailBatch().create({
})
```


### EmailsMetric

Create an instance: `const emails_metric = client.EmailsMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `broadcast_id` | `string` | Present when `broadcast` is in `dimensions`. |
| `broadcast_name` | `string` | Present when `broadcast` is in `dimensions`. |
| `domain_id` | `string` | Present when `domain` is in `dimensions`. |
| `domain_name` | `string` | Present when `domain` is in `dimensions`. |
| `email_id` | `string` | Present when `email` is in `dimensions`. |
| `period` | `string` | Present when `period` is in `dimensions`. |

#### Example: List

```ts
const emails_metrics = await client.EmailsMetric().list()
```


### Event

Create an instance: `const event = client.Event()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date and time the event was created. |
| `id` | `string` | The event ID. |
| `name` | `string` | The event name. |
| `object` | `string` | Type of the response object. |
| `schema` | `Record<string, any> | null` | A flat key/type map defining the event payload schema. |
| `updated_at` | `string | null` | The date and time the event was last updated. |

#### Example: Load

```ts
const event = await client.Event().load({ id: 'event_id' })
```

#### Example: List

```ts
const events = await client.Event().list()
```

#### Example: Create

```ts
const event = await client.Event().create({
  name: 'example_name',
})
```


### Inbox

Create an instance: `const inbox = client.Inbox()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | `string` | The ID of the domain that owns the address. |
| `drafts` | `number` | The number of unsent drafts. |
| `email_address` | `string` | The address of the inbox. |
| `forwarding` | `boolean` | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | `string | null` | The name recipients see when mail is sent from this inbox. |
| `id` | `string` | The ID of the inbox. |
| `last_received` | `string | null` | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | `string` | Internal name for the inbox. |
| `object` | `string` | The object type. |
| `receiving_address` | `string | null` | The address to forward mail to when forwarding is enabled. |
| `unread` | `number` | The number of unread threads in the Inbox folder. |
| `updated_at` | `string` | ISO 8601 timestamp when the inbox was last updated. |

#### Example: Load

```ts
const inbox = await client.Inbox().load({ id: 'inbox_id' })
```

#### Example: List

```ts
const inboxs = await client.Inbox().list()
```

#### Example: Create

```ts
const inbox = await client.Inbox().create({
  email_address: 'example_email_address',
})
```


### InboxAgent

Create an instance: `const inbox_agent = client.InboxAgent()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `actions` | `any[]` | The actions the agent can take, in the order of the enum. |
| `id` | `string` | The ID of the inbox agent. |
| `instructions` | `string | null` | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | `string` | The object type. |
| `tone` | `string | null` | The tone the agent writes in. |

#### Example: Load

```ts
const inbox_agent = await client.InboxAgent().load({ inbox_id: 'inbox_id' })
```


### InboxDraft

Create an instance: `const inbox_draft = client.InboxDraft()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any[]` | The files stored on the draft. |
| `bcc` | `any[]` | The BCC recipients. |
| `cc` | `any[]` | The CC recipients. |
| `created_at` | `string` | Timestamp when the draft was created. |
| `email_id` | `string | null` | The ID of the outbound email queued when the draft was sent. |
| `html` | `string | null` | The HTML body. |
| `id` | `string` | The ID of the draft. |
| `object` | `string` | The object type. |
| `reply_to_email_id` | `string | null` | The Email ID the draft replies to. |
| `snippet` | `string | null` | A short preview of the body. |
| `subject` | `string | null` | The subject. |
| `text` | `string | null` | The plain text body. |
| `thread_id` | `string | null` | The Thread ID when the draft is a reply. |
| `to` | `any[] | null` | The recipients. |
| `type` | `string` | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `string` | Timestamp when the draft was last saved. |

#### Example: Load

```ts
const inbox_draft = await client.InboxDraft().load({ id: 'inbox_draft_id', inbox_id: 'inbox_id' })
```

#### Example: List

```ts
const inbox_drafts = await client.InboxDraft().list({ inbox_id: "example" })
```

#### Example: Create

```ts
const inbox_draft = await client.InboxDraft().create({
  inbox_id: 'example_inbox_id',
})
```


### InboxLabel

Create an instance: `const inbox_label = client.InboxLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `string` | The color of the label, as a `#RRGGBB` hex color. |
| `created_at` | `string` | Timestamp when the label was created. |
| `id` | `string` | The ID of the updated label. |
| `name` | `string` | The name of the label. |
| `object` | `string` | Type of the response object. |

#### Example: List

```ts
const inbox_labels = await client.InboxLabel().list({ inbox_id: "example" })
```

#### Example: Create

```ts
const inbox_label = await client.InboxLabel().create({
  inbox_id: 'example_inbox_id',
  name: 'example_name',
})
```


### InboxThread

Create an instance: `const inbox_thread = client.InboxThread()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bcc` | `any[]` | The BCC recipients of the most recent email. |
| `cc` | `any[]` | The CC recipients of the most recent email. |
| `email_count` | `number` | The number of emails in the thread. |
| `folder` | `string` | The folder the thread is in. |
| `from` | `string | null` | The sender email address. |
| `has_attachment` | `boolean` | Whether the most recent email in the thread has an attachment. |
| `has_draft` | `boolean` | Whether the thread has an unsent draft. |
| `id` | `string` | The ID of the thread. |
| `labels` | `any[]` | The labels attached to the thread. |
| `object` | `string` | The object type. |
| `read` | `boolean` | True only when every email in the thread is read. |
| `received_at` | `string` | Timestamp when the thread was last active. |
| `subject` | `string | null` | The subject of the thread. |
| `to` | `any[]` | The recipients of the most recent email. |

#### Example: Load

```ts
const inbox_thread = await client.InboxThread().load({ id: 'inbox_thread_id', inbox_id: 'inbox_id' })
```

#### Example: List

```ts
const inbox_threads = await client.InboxThread().list({ inbox_id: "example" })
```


### InboxThreadEmail

Create an instance: `const inbox_thread_email = client.InboxThreadEmail()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any[]` | The attachments of the message. |
| `bcc` | `any[]` | The BCC recipients of the message. |
| `cc` | `any[]` | The CC recipients of the message. |
| `direction` | `string` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `string` | Sender email address. |
| `html` | `string | null` | The HTML body. |
| `id` | `string` | The ID of the message. |
| `message_id` | `string | null` | The Message-ID header of the message. |
| `read` | `boolean` | Whether the message has been read. |
| `received_at` | `string` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `any[]` | The Reply-To addresses of the message. |
| `subject` | `string | null` | The subject of the message. |
| `text` | `string | null` | The plain-text body. |
| `to` | `any[]` | The recipients of the message. |

#### Example: Load

```ts
const inbox_thread_email = await client.InboxThreadEmail().load({ id: 'inbox_thread_email_id', inbox_id: 'inbox_id', thread_id: 'thread_id' })
```

#### Example: List

```ts
const inbox_thread_emails = await client.InboxThreadEmail().list({ inbox_id: "example", thread_id: "example" })
```


### InboxThreadLabel

Create an instance: `const inbox_thread_label = client.InboxThreadLabel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `any[]` | Every label on the thread after the change. |
| `id` | `string` |  |
| `labels` | `any[]` | The IDs of the labels to apply to the thread. |
| `object` | `string` | Type of the response object. |

#### Example: Create

```ts
const inbox_thread_label = await client.InboxThreadLabel().create({
  inbox_id: 'example_inbox_id',
  thread_id: 'example_thread_id',
  labels: [],
})
```


### Log

Create an instance: `const log = client.Log()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date the log was created. |
| `endpoint` | `string` | The API endpoint that was called. |
| `id` | `string` | The log ID. |
| `method` | `string` | The HTTP method used. |
| `object` | `string` | Type of the response object. |
| `request_body` | `Record<string, any> | null` | The request body sent to the API. |
| `response_body` | `Record<string, any> | null` | The response body returned by the API. |
| `response_status` | `number` | The HTTP status code of the response. |
| `user_agent` | `string | null` | The user agent of the request. |

#### Example: Load

```ts
const log = await client.Log().load({ id: 'log_id' })
```

#### Example: List

```ts
const logs = await client.Log().list()
```


### OauthGrant

Create an instance: `const oauth_grant = client.OauthGrant()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client` | `Record<string, any>` | The OAuth client the grant was issued to. |
| `client_id` | `string` | The ID of the OAuth client the grant was issued to. |
| `created_at` | `string` | The date and time the OAuth grant was created. |
| `id` | `string` | The ID of the OAuth grant. |
| `revoked_at` | `string | null` | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | `string | null` | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | `any[]` | The scopes granted to the OAuth client. |

#### Example: List

```ts
const oauth_grants = await client.OauthGrant().list()
```


### ReceivedAttachment

Create an instance: `const received_attachment = client.ReceivedAttachment()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_disposition` | `string | null` | How the attachment should be displayed. |
| `content_id` | `string` | The content ID for inline attachments. |
| `content_type` | `string` | The MIME type of the attachment. |
| `download_url` | `string` | Signed URL to download the attachment content. |
| `expires_at` | `string` | Timestamp when the download URL expires. |
| `filename` | `string | null` | The filename of the attachment. |
| `id` | `string` | The ID of the attachment. |
| `object` | `string` | The type of object. |
| `size` | `number` | Size of the attachment in bytes. |

#### Example: Load

```ts
const received_attachment = await client.ReceivedAttachment().load({ id: 'received_attachment_id', email_id: 'email_id' })
```

#### Example: List

```ts
const received_attachments = await client.ReceivedAttachment().list({ email_id: "example" })
```


### ReceivedEmail

Create an instance: `const received_email = client.ReceivedEmail()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any[]` | Array of attachments. |
| `bcc` | `any[] | null` | The BCC recipients. |
| `cc` | `any[] | null` | The CC recipients. |
| `created_at` | `string` | Timestamp when the email was received. |
| `from` | `string` | The sender email address. |
| `headers` | `Record<string, any> | null` | The email headers. |
| `html` | `string | null` | The HTML content of the email. |
| `id` | `string` | The ID of the received email. |
| `message_id` | `string` | The unique message ID from the email headers. |
| `object` | `string` | The type of object. |
| `received_for` | `any[]` | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | `any[] | null` | The reply-to addresses. |
| `subject` | `string` | The email subject. |
| `text` | `string | null` | The plain text content of the email. |
| `to` | `any[]` | The recipient email addresses. |

#### Example: Load

```ts
const received_email = await client.ReceivedEmail().load({ email_id: 'email_id' })
```

#### Example: List

```ts
const received_emails = await client.ReceivedEmail().list()
```


### Segment

Create an instance: `const segment = client.Segment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | The ID of the audience this segment belongs to. |
| `created_at` | `string` | Timestamp indicating when the segment was created. |
| `filter` | `Record<string, any>` | Filter conditions for the segment. |
| `id` | `string` | The ID of the segment. |
| `name` | `string` | The name of the segment. |
| `object` | `string` | The object type. |

#### Example: Load

```ts
const segment = await client.Segment().load({ id: 'segment_id' })
```

#### Example: List

```ts
const segments = await client.Segment().list()
```

#### Example: Create

```ts
const segment = await client.Segment().create({
  name: 'example_name',
})
```


### SegmentContact

Create an instance: `const segment_contact = client.SegmentContact()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the contact was created. |
| `email` | `string` | Email address of the contact. |
| `first_name` | `string | null` | First name of the contact. |
| `id` | `string` | Unique identifier for the contact. |
| `last_name` | `string | null` | Last name of the contact. |
| `unsubscribed` | `boolean` | Indicates if the contact is unsubscribed. |

#### Example: List

```ts
const segment_contacts = await client.SegmentContact().list({ segment_id: "example" })
```


### Suppression

Create an instance: `const suppression = client.Suppression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the suppression was created. |
| `email` | `string` | Email address that is suppressed. |
| `id` | `string` | Unique identifier for the suppression. |
| `object` | `string` | Type of the response object. |
| `origin` | `string` | Origin of the suppression. |
| `source_id` | `string | null` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

#### Example: Load

```ts
const suppression = await client.Suppression().load({ id: 'suppression_id' })
```

#### Example: List

```ts
const suppressions = await client.Suppression().list()
```

#### Example: Create

```ts
const suppression = await client.Suppression().create({
  email: 'example_email',
})
```


### SuppressionBatchAdd

Create an instance: `const suppression_batch_add = client.SuppressionBatchAdd()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `emails` | `any[]` | Email addresses to suppress. |

#### Example: Create

```ts
const suppression_batch_add = await client.SuppressionBatchAdd().create({
  emails: [],
})
```


### SuppressionBatchRemove

Create an instance: `const suppression_batch_remove = client.SuppressionBatchRemove()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `any[]` | Array containing the removed suppressions. |
| `emails` | `any[]` | Email addresses to remove from the suppression list. |
| `ids` | `any[]` | Suppression IDs to remove from the suppression list. |

#### Example: Create

```ts
const suppression_batch_remove = await client.SuppressionBatchRemove().create({
})
```


### Template

Create an instance: `const template = client.Template()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alias` | `string` | The alias of the template. |
| `created_at` | `string` | Timestamp indicating when the template was created. |
| `current_version_id` | `string` | The ID of the current version of the template. |
| `from` | `string` | Sender email address. |
| `has_unpublished_versions` | `boolean` | Indicates whether the template has unpublished versions. |
| `html` | `string` | The HTML version of the template. |
| `id` | `string` | The ID of the template. |
| `name` | `string` | The name of the template. |
| `object` | `string` | The type of object. |
| `published_at` | `string | null` | Timestamp indicating when the template was published. |
| `reply_to` | `any[] | null` | Reply-to email addresses. |
| `status` | `string` | The publication status of the template. |
| `subject` | `string` | Email subject. |
| `text` | `string` | The plain text version of the template. |
| `updated_at` | `string` | Timestamp indicating when the template was last updated. |
| `variables` | `any[]` |  |

#### Example: Load

```ts
const template = await client.Template().load({ id: 'template_id' })
```

#### Example: List

```ts
const templates = await client.Template().list()
```

#### Example: Create

```ts
const template = await client.Template().create({
  html: 'example_html',
  name: 'example_name',
})
```


### Topic

Create an instance: `const topic = client.Topic()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the topic was created. |
| `default_subscription` | `string` | The default subscription status for the topic. |
| `description` | `string` | A description of the topic. |
| `id` | `string` | The ID of the topic. |
| `name` | `string` | The name of the topic. |
| `object` | `string` | The object type. |
| `visibility` | `string` | The visibility of the topic. |

#### Example: Load

```ts
const topic = await client.Topic().load({ id: 'topic_id' })
```

#### Example: List

```ts
const topics = await client.Topic().list()
```

#### Example: Create

```ts
const topic = await client.Topic().create({
  default_subscription: 'example_default_subscription',
  name: 'example_name',
})
```


### Usage

Create an instance: `const usage = client.Usage()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ai_credits` | `Record<string, any>` |  |
| `automation_runs` | `Record<string, any>` |  |
| `broadcasts` | `Record<string, any>` |  |
| `contacts` | `Record<string, any>` |  |
| `domains` | `Record<string, any>` |  |
| `emails` | `Record<string, any>` |  |
| `object` | `string` | The type of object. |
| `rate_limit` | `Record<string, any>` |  |
| `segments` | `Record<string, any>` |  |

#### Example: Load

```ts
const usage = await client.Usage().load()
```


### Webhook

Create an instance: `const webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the webhook was created. |
| `endpoint` | `string` | The URL where webhook events are sent. |
| `events` | `any[] | null` | Array of event types subscribed to. |
| `id` | `string` | The ID of the webhook. |
| `object` | `string` | The type of object. |
| `signing_secret` | `string` | The secret key used to verify webhook payloads. |
| `status` | `string` | The status of the webhook. |

#### Example: Load

```ts
const webhook = await client.Webhook().load({ id: 'webhook_id' })
```

#### Example: List

```ts
const webhooks = await client.Webhook().list()
```

#### Example: Create

```ts
const webhook = await client.Webhook().create({
  endpoint: 'example_endpoint',
  events: [],
})
```


### WebhookEvent

Create an instance: `const webhook_event = client.WebhookEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the event was created. |
| `id` | `string` | The ID of the webhook event. |
| `next_attempt_at` | `string | null` | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | `string` | The type of object. |
| `payload` | `Record<string, any>` | The event payload sent to the webhook endpoint. |
| `status` | `string` | The delivery status of the event for this webhook. |
| `type` | `string` | The type of the event. |

#### Example: Load

```ts
const webhook_event = await client.WebhookEvent().load({ id: 'webhook_event_id', webhook_id: 'webhook_id' })
```

#### Example: List

```ts
const webhook_events = await client.WebhookEvent().list({ webhook_id: "example" })
```

#### Example: Create

```ts
const webhook_event = await client.WebhookEvent().create({
  event_id: 'example_event_id',
  webhook_id: 'example_webhook_id',
})
```


### WebhookEventAttempt

Create an instance: `const webhook_event_attempt = client.WebhookEventAttempt()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `http_status_code` | `number` | The HTTP status code returned by the webhook endpoint. |
| `id` | `string` | The ID of the webhook event attempt. |
| `response` | `string` | The response body returned by the webhook endpoint. |
| `sent_at` | `string` | Timestamp indicating when the attempt was sent. |

#### Example: List

```ts
const webhook_event_attempts = await client.WebhookEventAttempt().list({ event_id: "example", webhook_id: "example" })
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

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

### Module structure

```
resend/
├── src/
│   ├── ResendSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { ResendSDK } from '@shaheeramjad/resend-sdk'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const oauthgrant = client.OauthGrant()
await oauthgrant.list()

// oauthgrant.data() now returns the oauthgrant data from the last `list`
// oauthgrant.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
