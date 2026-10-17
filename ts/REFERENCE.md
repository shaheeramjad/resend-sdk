# Resend TypeScript SDK Reference

Complete API reference for the Resend TypeScript SDK.


## ResendSDK

### Constructor

```ts
new ResendSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `ResendSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = ResendSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `ResendSDK` instance in test mode.


### Instance Methods

#### `ApiKey(data?: object)`

Create a new `ApiKey` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ApiKeyEntity` instance.

#### `Attachment(data?: object)`

Create a new `Attachment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AttachmentEntity` instance.

#### `Audience(data?: object)`

Create a new `Audience` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AudienceEntity` instance.

#### `Automation(data?: object)`

Create a new `Automation` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AutomationEntity` instance.

#### `AutomationRun(data?: object)`

Create a new `AutomationRun` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AutomationRunEntity` instance.

#### `Broadcast(data?: object)`

Create a new `Broadcast` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BroadcastEntity` instance.

#### `BroadcastClickedLink(data?: object)`

Create a new `BroadcastClickedLink` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BroadcastClickedLinkEntity` instance.

#### `BroadcastRecipient(data?: object)`

Create a new `BroadcastRecipient` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BroadcastRecipientEntity` instance.

#### `Contact(data?: object)`

Create a new `Contact` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactEntity` instance.

#### `ContactImport(data?: object)`

Create a new `ContactImport` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactImportEntity` instance.

#### `ContactProperty(data?: object)`

Create a new `ContactProperty` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactPropertyEntity` instance.

#### `ContactSegment(data?: object)`

Create a new `ContactSegment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactSegmentEntity` instance.

#### `ContactTopic(data?: object)`

Create a new `ContactTopic` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactTopicEntity` instance.

#### `Domain(data?: object)`

Create a new `Domain` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DomainEntity` instance.

#### `DomainClaim(data?: object)`

Create a new `DomainClaim` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DomainClaimEntity` instance.

#### `Email(data?: object)`

Create a new `Email` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailEntity` instance.

#### `EmailBatch(data?: object)`

Create a new `EmailBatch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailBatchEntity` instance.

#### `EmailsMetric(data?: object)`

Create a new `EmailsMetric` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EmailsMetricEntity` instance.

#### `Event(data?: object)`

Create a new `Event` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventEntity` instance.

#### `Inbox(data?: object)`

Create a new `Inbox` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxEntity` instance.

#### `InboxAgent(data?: object)`

Create a new `InboxAgent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxAgentEntity` instance.

#### `InboxDraft(data?: object)`

Create a new `InboxDraft` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxDraftEntity` instance.

#### `InboxLabel(data?: object)`

Create a new `InboxLabel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxLabelEntity` instance.

#### `InboxThread(data?: object)`

Create a new `InboxThread` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxThreadEntity` instance.

#### `InboxThreadEmail(data?: object)`

Create a new `InboxThreadEmail` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxThreadEmailEntity` instance.

#### `InboxThreadLabel(data?: object)`

Create a new `InboxThreadLabel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InboxThreadLabelEntity` instance.

#### `Log(data?: object)`

Create a new `Log` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `LogEntity` instance.

#### `OauthGrant(data?: object)`

Create a new `OauthGrant` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OauthGrantEntity` instance.

#### `ReceivedAttachment(data?: object)`

Create a new `ReceivedAttachment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReceivedAttachmentEntity` instance.

#### `ReceivedEmail(data?: object)`

Create a new `ReceivedEmail` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReceivedEmailEntity` instance.

#### `Segment(data?: object)`

Create a new `Segment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SegmentEntity` instance.

#### `SegmentContact(data?: object)`

Create a new `SegmentContact` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SegmentContactEntity` instance.

#### `Suppression(data?: object)`

Create a new `Suppression` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SuppressionEntity` instance.

#### `SuppressionBatchAdd(data?: object)`

Create a new `SuppressionBatchAdd` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SuppressionBatchAddEntity` instance.

#### `SuppressionBatchRemove(data?: object)`

Create a new `SuppressionBatchRemove` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SuppressionBatchRemoveEntity` instance.

#### `Template(data?: object)`

Create a new `Template` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TemplateEntity` instance.

#### `Topic(data?: object)`

Create a new `Topic` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TopicEntity` instance.

#### `Usage(data?: object)`

Create a new `Usage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UsageEntity` instance.

#### `Webhook(data?: object)`

Create a new `Webhook` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebhookEntity` instance.

#### `WebhookEvent(data?: object)`

Create a new `WebhookEvent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebhookEventEntity` instance.

#### `WebhookEventAttempt(data?: object)`

Create a new `WebhookEventAttempt` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebhookEventAttemptEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |
| `fetchargs.ctrl.signal` | `AbortSignal` | Aborts the request in flight: `ok` is then `false` and `err.code` is `request_aborted`. |

**Returns:** `Promise<{ ok, status, headers, data }>`. On a failure
`ok` is `false` and `err` holds the error.

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `ResendSDK.test()`.

**Returns:** `ResendSDK` instance in test mode.

#### Cancelling a call

Every entity operation takes an optional `ctrl` object after its match or
data, and an `AbortSignal` in `ctrl.signal` cancels the request in flight.
The operation then rejects with an error whose `code` is
`request_aborted` and whose `cause` is the signal's reason. A request
whose signal has already aborted is not sent. `stream()` takes the signal
as `callopts.signal`, and ends when it aborts.


---

## ApiKeyEntity

```ts
const api_key = client.ApiKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date and time the API key was created. |
| `domain_id` | `string` | No | Restrict an API key to send emails only from a specific domain. |
| `id` | `string` | No | The ID of the API key. |
| `last_used_at` | `string | null` | No | The date and time the API key was last used. |
| `name` | `string` | Yes | The API key name. |
| `object` | `string` | No | The type of object. |
| `permission` | `string` | No | The API key can have full access to Resend’s API or be only restricted to send emails. |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `domain_id` | - | - | - | - |
| `id` | - | - | - | - |
| `last_used_at` | - | - | - | - |
| `name` | Yes | - | - | - |
| `object` | - | - | - | - |
| `permission` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.ApiKey().create({
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ApiKey().list()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.ApiKey().remove({ id: 'api_key_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.ApiKey().update({
  id: 'api_key_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ApiKeyEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AttachmentEntity

```ts
const attachment = client.Attachment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_disposition` | `string | null` | No | How the attachment should be displayed. |
| `content_id` | `string` | No | The content ID for inline attachments. |
| `content_type` | `string` | No | The MIME type of the attachment. |
| `download_url` | `string` | No | Signed URL to download the attachment content. |
| `expires_at` | `string` | No | Timestamp when the download URL expires. |
| `filename` | `string | null` | No | The filename of the attachment. |
| `id` | `string` | No | The ID of the attachment. |
| `object` | `string` | No | The type of object. |
| `size` | `number` | No | Size of the attachment in bytes. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Attachment().list({ email_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Attachment().load({ id: 'attachment_id', email_id: 'email_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AttachmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AudienceEntity

```ts
const audience = client.Audience()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date that the object was created. |
| `id` | `string` | No | The ID of the audience. |
| `name` | `string` | No | The name of the audience. |
| `object` | `string` | No | The object of the audience. |

### Field Usage by Operation

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `object` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Audience().create({
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Audience().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Audience().load({ id: 'audience_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Audience().remove({ id: 'audience_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AudienceEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AutomationEntity

```ts
const automation = client.Automation()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `connections` | `any[]` | No | The connections between steps in the active version of the automation. |
| `created_at` | `string` | No | The date and time the automation was created. |
| `id` | `string` | No | The ID of the automation. |
| `name` | `string` | No | The name of the automation. |
| `object` | `string` | No | Type of the response object. |
| `status` | `string` | No | The current status of the automation. |
| `steps` | `any[]` | No | The steps in the active version of the automation. |
| `updated_at` | `string` | No | The date and time the automation was last updated. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `connections` | - | - | Yes | - | - |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `steps` | - | - | Yes | - | - |
| `updated_at` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `duplicate` | `/automations/{automation_id}/duplicate` | `client.Automation().create({ $action: 'duplicate', ... })` |
| `stop` | `/automations/{automation_id}/stop` | `client.Automation().create({ $action: 'stop', ... })` |

An action returns that action's OWN response, which is not necessarily a
Automation record — check the API definition for its shape.

```ts
const result = await client.Automation().create({
  $action: 'duplicate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Automation().create({
  connections: [],
  name: 'example_name',
  steps: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Automation().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Automation().load({ id: 'automation_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Automation().remove({ id: 'automation_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Automation().update({
  id: 'automation_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AutomationEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AutomationRunEntity

```ts
const automation_run = client.AutomationRun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `completed_at` | `string | null` | No | The date and time the run completed. |
| `created_at` | `string` | No | The date and time the run was created. |
| `id` | `string` | No | The ID of the automation run. |
| `object` | `string` | No | Type of the response object. |
| `started_at` | `string | null` | No | The date and time the run started. |
| `status` | `string` | No | The current status of the automation run. |
| `steps` | `any[]` | No | The steps executed in this run, sorted in graph order. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.AutomationRun().list({ automation_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.AutomationRun().load({ id: 'automation_run_id', automation_id: 'automation_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AutomationRunEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BroadcastEntity

```ts
const broadcast = client.Broadcast()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string | null` | No | Deprecated: use `segment_id` instead. |
| `created_at` | `string` | No | Timestamp indicating when the broadcast was created. |
| `from` | `string` | No | The email address of the sender. |
| `html` | `string | null` | No | The HTML version of the broadcast content. |
| `id` | `string` | No | Unique identifier for the broadcast. |
| `name` | `string` | No | Name of the broadcast. |
| `object` | `string` | No | The object type of the response. |
| `preview_text` | `string` | No | The preview text of the email. |
| `reply_to` | `any[]` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | No | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `string | null` | No | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `boolean` | No | Whether to send the broadcast immediately or keep it as a draft. |
| `sent_at` | `string` | No | Timestamp indicating when the broadcast was sent. |
| `status` | `string` | No | The status of the broadcast. |
| `subject` | `string` | No | The subject line of the email. |
| `text` | `string | null` | No | The plain text version of the broadcast content. |
| `topic_id` | `string | null` | No | The topic ID that the broadcast is scoped to. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `audience_id` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `from` | - | - | Yes | - | - |
| `html` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `object` | - | - | - | - | - |
| `preview_text` | - | - | - | - | - |
| `reply_to` | - | - | - | - | - |
| `scheduled_at` | - | - | - | - | - |
| `segment_id` | - | - | Yes | - | - |
| `send` | - | - | - | - | - |
| `sent_at` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `subject` | - | - | Yes | - | - |
| `text` | - | - | - | - | - |
| `topic_id` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `cancel` | `/broadcasts/{id}/cancel` | `client.Broadcast().create({ $action: 'cancel', ... })` |
| `duplicate` | `/broadcasts/{id}/duplicate` | `client.Broadcast().create({ $action: 'duplicate', ... })` |
| `send` | `/broadcasts/{id}/send` | `client.Broadcast().create({ $action: 'send', ... })` |

An action returns that action's OWN response, which is not necessarily a
Broadcast record — check the API definition for its shape.

```ts
const result = await client.Broadcast().create({
  $action: 'cancel',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Broadcast().create({
  from: 'example_from',
  segment_id: 'example_segment_id',
  subject: 'example_subject',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Broadcast().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Broadcast().load({ id: 'broadcast_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Broadcast().remove({ id: 'broadcast_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Broadcast().update({
  id: 'broadcast_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BroadcastEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BroadcastClickedLinkEntity

```ts
const broadcast_clicked_link = client.BroadcastClickedLink()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clicks` | `number` | No | Total number of clicks on this URL. |
| `id` | `string` | No | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `number` | No | Number of unique clicks on this URL. |
| `url` | `string` | No | The URL that was clicked. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.BroadcastClickedLink().list({ broadcast_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BroadcastClickedLinkEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BroadcastRecipientEntity

```ts
const broadcast_recipient = client.BroadcastRecipient()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bounce_type` | `string` | No | The type of bounce. |
| `clicked_links` | `any[]` | No | The links this recipient clicked. |
| `contact_id` | `string | null` | No | The ID of the contact associated with this recipient, if one exists. |
| `count` | `number` | No | The number of times this recipient triggered the event. |
| `email` | `string` | No | The recipient's email address. |
| `id` | `string` | No | Opaque cursor identifying this row, used for pagination. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.BroadcastRecipient().list({ broadcast_id: "example", type: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BroadcastRecipientEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactEntity

```ts
const contact = client.Contact()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | Unique identifier of the audience to which the contact belongs. |
| `created_at` | `string` | No | Timestamp indicating when the contact was created. |
| `email` | `string` | No | Email address of the contact. |
| `first_name` | `string | null` | No | First name of the contact. |
| `id` | `string` | No | Unique identifier for the contact. |
| `last_name` | `string | null` | No | Last name of the contact. |
| `object` | `string` | No | Type of the response object. |
| `properties` | `Record<string, any>` | No | A map of custom property keys and values. |
| `segments` | `any[]` | No | Array of segment IDs to add the contact to. |
| `topics` | `any[]` | No | Array of topic subscriptions for the contact. |
| `unsubscribed` | `boolean` | No | Indicates if the contact is unsubscribed. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `audience_id` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `email` | - | - | Yes | - | - |
| `first_name` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `last_name` | - | - | - | - | - |
| `object` | - | - | - | - | - |
| `properties` | - | - | - | - | - |
| `segments` | - | - | - | - | - |
| `topics` | - | - | - | - | - |
| `unsubscribed` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Contact().create({
  email: 'example_email',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Contact().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Contact().load({ id: 'contact_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Contact().remove({ id: 'contact_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Contact().update({
  id: 'contact_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactImportEntity

```ts
const contact_import = client.ContactImport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `column_map` | `string` | No | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | `string | null` | No | Timestamp indicating when the contact import completed. |
| `counts` | `Record<string, any>` | No |  |
| `created_at` | `string` | No | Timestamp indicating when the contact import was created. |
| `file` | `string` | Yes | CSV file to import. |
| `id` | `string` | No | Unique identifier for the contact import. |
| `object` | `string` | No | Type of the response object. |
| `on_conflict` | `string` | No | Strategy to use when an imported contact already exists. |
| `segments` | `string` | No | JSON-encoded array of segments to add imported contacts to. |
| `status` | `string` | No | Current status of the contact import. |
| `topics` | `string` | No | JSON-encoded array of topic subscriptions to apply to imported contacts. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.ContactImport().create({
  file: 'example_file',
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ContactImport().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.ContactImport().load({ id: 'contact_import_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactImportEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactPropertyEntity

```ts
const contact_property = client.ContactProperty()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the contact property was created. |
| `fallback_value` | `any` | No | The default value when the property is not set for a contact. |
| `id` | `string` | No | The ID of the contact property. |
| `key` | `string` | No | The property key. |
| `object` | `string` | No | The object type. |
| `type` | `string` | No | The property type. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `fallback_value` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `key` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `type` | - | - | Yes | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.ContactProperty().create({
  key: 'example_key',
  type: 'example_type',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ContactProperty().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.ContactProperty().load({ id: 'contact_property_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.ContactProperty().remove({ id: 'contact_property_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.ContactProperty().update({
  id: 'contact_property_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactPropertyEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactSegmentEntity

```ts
const contact_segment = client.ContactSegment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No | The ID of the contact. |
| `created_at` | `string` | No | Timestamp indicating when the contact was added to the segment. |
| `id` | `string` | No | Unique identifier for the segment. |
| `name` | `string` | No | Name of the segment. |
| `object` | `string` | No | The object type. |
| `segment_id` | `string` | No | The ID of the segment. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.ContactSegment().create({
  contact_id: 'example_contact_id',
  id: 'example_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ContactSegment().list({ contact_id: "example" })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.ContactSegment().remove({ id: 'contact_segment_id', contact_id: 'contact_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactSegmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactTopicEntity

```ts
const contact_topic = client.ContactTopic()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No | The ID of the contact. |
| `description` | `string` | No | Description of the topic. |
| `id` | `string` | No | Unique identifier for the topic. |
| `name` | `string` | No | Name of the topic. |
| `object` | `string` | No | The object type. |
| `subscription` | `string` | No | The subscription status for this topic. |
| `topics` | `any[]` | No | Array of updated topic subscriptions. |

### Field Usage by Operation

| Field | list | update |
| --- | --- | --- |
| `contact_id` | - | - |
| `description` | - | - |
| `id` | - | - |
| `name` | - | - |
| `object` | - | - |
| `subscription` | - | - |
| `topics` | - | Yes |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ContactTopic().list({ contact_id: "example" })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.ContactTopic().update({
  contact_id: 'contact_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactTopicEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DomainEntity

```ts
const domain = client.Domain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capabilities` | `Record<string, any>` | No | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `boolean` | No | Whether click tracking is enabled for this domain. |
| `created_at` | `string` | No | The date and time the domain was created. |
| `custom_return_path` | `string` | No | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | `string` | No | The ID of the domain. |
| `name` | `string` | No | The name of the domain. |
| `object` | `string` | No | The type of object. |
| `open_tracking` | `boolean` | No | Whether open tracking is enabled for this domain. |
| `records` | `any[]` | No |  |
| `region` | `string` | No | The region where the domain is hosted. |
| `status` | `string` | No | The status of the domain. |
| `tls` | `string` | No | TLS mode. |
| `tracking_subdomain` | `string` | No | The subdomain used for click and open tracking. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `capabilities` | - | - | - | - | - |
| `click_tracking` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `custom_return_path` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `open_tracking` | - | - | - | - | - |
| `records` | - | - | - | - | - |
| `region` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `tls` | - | - | - | - | - |
| `tracking_subdomain` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `verify` | `/domains/{domain_id}/verify` | `client.Domain().create({ $action: 'verify', ... })` |

An action returns that action's OWN response, which is not necessarily a
Domain record — check the API definition for its shape.

```ts
const result = await client.Domain().create({
  $action: 'verify',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Domain().create({
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Domain().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Domain().load({ id: 'domain_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Domain().remove({ id: 'domain_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Domain().update({
  id: 'domain_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DomainEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DomainClaimEntity

```ts
const domain_claim = client.DomainClaim()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `blocked_reason` | `string | null` | No | Why the claim is currently blocked, if applicable. |
| `click_tracking` | `boolean` | No | Track clicks within the body of each HTML email. |
| `created_at` | `string` | No | The date and time the claim was created. |
| `custom_return_path` | `string` | No | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | `string | null` | No | The ID of the placeholder domain created for the claim. |
| `expires_at` | `string` | No | The date and time the claim expires if not verified. |
| `failure_reason` | `string | null` | No | Why the claim failed, if applicable. |
| `id` | `string` | No | The ID of the claim. |
| `name` | `string` | No | The name of the domain being claimed. |
| `object` | `string` | No | The type of object. |
| `open_tracking` | `boolean` | No | Track the open rate of each email. |
| `record` | `Record<string, any>` | No | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | `string | null` | No | The region where the claimed domain will send from. |
| `status` | `string` | No | The status of the claim. |
| `tracking_subdomain` | `string` | No | The subdomain to use for click and open tracking. |

### Field Usage by Operation

| Field | load | create |
| --- | --- | --- |
| `blocked_reason` | - | - |
| `click_tracking` | - | - |
| `created_at` | - | - |
| `custom_return_path` | - | - |
| `domain_id` | - | - |
| `expires_at` | - | - |
| `failure_reason` | - | - |
| `id` | - | - |
| `name` | - | Yes |
| `object` | - | - |
| `open_tracking` | - | - |
| `record` | - | - |
| `region` | - | - |
| `status` | - | - |
| `tracking_subdomain` | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.DomainClaim().create({
  name: 'example_name',
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.DomainClaim().load({ id: 'domain_claim_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DomainClaimEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailEntity

```ts
const email = client.Email()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any[]` | No |  |
| `bcc` | `any[]` | No | The email addresses of the blind carbon copy recipients. |
| `cc` | `any[]` | No | The email addresses of the carbon copy recipients. |
| `created_at` | `string` | No | The date and time the email was created. |
| `from` | `string` | No | The email address of the sender. |
| `headers` | `Record<string, any>` | No | Custom headers to add to the email. |
| `html` | `string` | No | The HTML body of the email. |
| `id` | `string` | No | The ID of the email. |
| `last_event` | `string` | No | The status of the email. |
| `message_id` | `string` | No | The Message-ID header value of the email. |
| `object` | `string` | No | The type of object. |
| `reply_to` | `any[]` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | No | Schedule email to be sent later. |
| `subject` | `string` | No | The subject line of the email. |
| `tags` | `any[]` | No |  |
| `template` | `any` | No |  |
| `text` | `string` | No | The plain text body of the email. |
| `to` | `any[]` | No | Recipient email address. |
| `topic_id` | `string` | No | The topic ID to scope the email to. |

### Field Usage by Operation

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `attachments` | - | - | - | - |
| `bcc` | - | - | - | - |
| `cc` | - | - | - | - |
| `created_at` | - | - | - | - |
| `from` | - | - | Yes | - |
| `headers` | - | - | - | - |
| `html` | - | - | - | - |
| `id` | - | - | - | - |
| `last_event` | - | - | - | - |
| `message_id` | - | - | - | - |
| `object` | - | - | - | - |
| `reply_to` | - | - | - | - |
| `scheduled_at` | - | - | - | - |
| `subject` | - | - | Yes | - |
| `tags` | - | - | - | - |
| `template` | - | - | - | - |
| `text` | - | - | - | - |
| `to` | - | - | Yes | - |
| `topic_id` | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `cancel` | `/emails/{email_id}/cancel` | `client.Email().create({ $action: 'cancel', ... })` |
| `share` | `/emails/{email_id}/share` | `client.Email().create({ $action: 'share', ... })` |

An action returns that action's OWN response, which is not necessarily a
Email record — check the API definition for its shape.

```ts
const result = await client.Email().create({
  $action: 'cancel',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Email().create({
  from: 'example_from',
  subject: 'example_subject',
  to: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Email().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Email().load({ id: 'email_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Email().update({
  id: 'email_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailBatchEntity

```ts
const email_batch = client.EmailBatch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any[]` | No |  |
| `send_emails` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.EmailBatch().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailBatchEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EmailsMetricEntity

```ts
const emails_metric = client.EmailsMetric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `broadcast_id` | `string` | No | Present when `broadcast` is in `dimensions`. |
| `broadcast_name` | `string` | No | Present when `broadcast` is in `dimensions`. |
| `domain_id` | `string` | No | Present when `domain` is in `dimensions`. |
| `domain_name` | `string` | No | Present when `domain` is in `dimensions`. |
| `email_id` | `string` | No | Present when `email` is in `dimensions`. |
| `period` | `string` | No | Present when `period` is in `dimensions`. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.EmailsMetric().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EmailsMetricEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventEntity

```ts
const event = client.Event()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date and time the event was created. |
| `id` | `string` | No | The event ID. |
| `name` | `string` | No | The event name. |
| `object` | `string` | No | Type of the response object. |
| `schema` | `Record<string, any> | null` | No | A flat key/type map defining the event payload schema. |
| `updated_at` | `string | null` | No | The date and time the event was last updated. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `schema` | - | - | - | Yes | - |
| `updated_at` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `send` | `/events/send` | `client.Event().create({ $action: 'send', ... })` |

An action returns that action's OWN response, which is not necessarily a
Event record — check the API definition for its shape.

```ts
const result = await client.Event().create({
  $action: 'send',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Event().create({
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Event().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Event().load({ id: 'event_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Event().remove({ id: 'event_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Event().update({
  id: 'event_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxEntity

```ts
const inbox = client.Inbox()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | `string` | Yes | The ID of the domain that owns the address. |
| `drafts` | `number` | Yes | The number of unsent drafts. |
| `email_address` | `string` | Yes | The address of the inbox. |
| `forwarding` | `boolean` | No | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | `string | null` | Yes | The name recipients see when mail is sent from this inbox. |
| `id` | `string` | Yes | The ID of the inbox. |
| `last_received` | `string | null` | Yes | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | `string` | Yes | Internal name for the inbox. |
| `object` | `string` | Yes | The object type. |
| `receiving_address` | `string | null` | Yes | The address to forward mail to when forwarding is enabled. |
| `unread` | `number` | Yes | The number of unread threads in the Inbox folder. |
| `updated_at` | `string` | Yes | ISO 8601 timestamp when the inbox was last updated. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `domain_id` | - | - | - | - | - |
| `drafts` | - | - | - | - | - |
| `email_address` | - | - | - | - | - |
| `forwarding` | - | - | - | - | - |
| `from_name` | - | - | Yes | Yes | - |
| `id` | - | - | - | - | - |
| `last_received` | - | - | - | - | - |
| `name` | - | - | Yes | Yes | - |
| `object` | - | - | - | - | - |
| `receiving_address` | - | - | - | - | - |
| `unread` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Inbox().create({
  email_address: 'example_email_address',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Inbox().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Inbox().load({ id: 'inbox_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Inbox().remove({ id: 'inbox_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Inbox().update({
  id: 'inbox_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxAgentEntity

```ts
const inbox_agent = client.InboxAgent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `any[]` | Yes | The actions the agent can take, in the order of the enum. |
| `id` | `string` | Yes | The ID of the inbox agent. |
| `instructions` | `string | null` | Yes | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | `string` | Yes | The object type. |
| `tone` | `string | null` | Yes | The tone the agent writes in. |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `actions` | - | Yes |
| `id` | - | - |
| `instructions` | - | Yes |
| `object` | - | - |
| `tone` | - | Yes |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.InboxAgent().load({ inbox_id: 'inbox_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.InboxAgent().update({
  inbox_id: 'inbox_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxAgentEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxDraftEntity

```ts
const inbox_draft = client.InboxDraft()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any[]` | Yes | The files stored on the draft. |
| `bcc` | `any[]` | Yes | The BCC recipients. |
| `cc` | `any[]` | Yes | The CC recipients. |
| `created_at` | `string` | Yes | Timestamp when the draft was created. |
| `email_id` | `string | null` | Yes | The ID of the outbound email queued when the draft was sent. |
| `html` | `string | null` | Yes | The HTML body. |
| `id` | `string` | Yes | The ID of the draft. |
| `object` | `string` | Yes | The object type. |
| `reply_to_email_id` | `string | null` | Yes | The Email ID the draft replies to. |
| `snippet` | `string | null` | Yes | A short preview of the body. |
| `subject` | `string | null` | Yes | The subject. |
| `text` | `string | null` | Yes | The plain text body. |
| `thread_id` | `string | null` | Yes | The Thread ID when the draft is a reply. |
| `to` | `any[] | null` | Yes | The recipients. |
| `type` | `string` | Yes | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `string` | Yes | Timestamp when the draft was last saved. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `attachments` | - | - | Yes | Yes | - |
| `bcc` | - | - | Yes | Yes | - |
| `cc` | - | - | Yes | Yes | - |
| `created_at` | - | - | - | - | - |
| `email_id` | - | - | - | - | - |
| `html` | - | - | Yes | Yes | - |
| `id` | - | - | - | - | - |
| `object` | - | - | - | - | - |
| `reply_to_email_id` | - | - | Yes | - | - |
| `snippet` | - | - | - | - | - |
| `subject` | - | - | Yes | Yes | - |
| `text` | - | - | Yes | Yes | - |
| `thread_id` | - | - | Yes | - | - |
| `to` | - | - | Yes | Yes | - |
| `type` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `send` | `/inboxes/{inbox_id}/drafts/{draft_id}/send` | `client.InboxDraft().create({ $action: 'send', ... })` |

An action returns that action's OWN response, which is not necessarily a
InboxDraft record — check the API definition for its shape.

```ts
const result = await client.InboxDraft().create({
  $action: 'send',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.InboxDraft().create({
  inbox_id: 'example_inbox_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.InboxDraft().list({ inbox_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.InboxDraft().load({ id: 'inbox_draft_id', inbox_id: 'inbox_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.InboxDraft().remove({ id: 'inbox_draft_id', inbox_id: 'inbox_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.InboxDraft().update({
  id: 'inbox_draft_id',
  inbox_id: 'inbox_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxDraftEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxLabelEntity

```ts
const inbox_label = client.InboxLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `string` | No | The color of the label, as a `#RRGGBB` hex color. |
| `created_at` | `string` | Yes | Timestamp when the label was created. |
| `id` | `string` | Yes | The ID of the updated label. |
| `name` | `string` | Yes | The name of the label. |
| `object` | `string` | Yes | Type of the response object. |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `color` | Yes | - | - | - |
| `created_at` | - | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `object` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.InboxLabel().create({
  inbox_id: 'example_inbox_id',
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.InboxLabel().list({ inbox_id: "example" })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.InboxLabel().remove({ id: 'inbox_label_id', inbox_id: 'inbox_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.InboxLabel().update({
  id: 'inbox_label_id',
  inbox_id: 'inbox_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxLabelEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxThreadEntity

```ts
const inbox_thread = client.InboxThread()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bcc` | `any[]` | Yes | The BCC recipients of the most recent email. |
| `cc` | `any[]` | Yes | The CC recipients of the most recent email. |
| `email_count` | `number` | Yes | The number of emails in the thread. |
| `folder` | `string` | Yes | The folder the thread is in. |
| `from` | `string | null` | Yes | The sender email address. |
| `has_attachment` | `boolean` | Yes | Whether the most recent email in the thread has an attachment. |
| `has_draft` | `boolean` | Yes | Whether the thread has an unsent draft. |
| `id` | `string` | Yes | The ID of the thread. |
| `labels` | `any[]` | Yes | The labels attached to the thread. |
| `object` | `string` | Yes | The object type. |
| `read` | `boolean` | Yes | True only when every email in the thread is read. |
| `received_at` | `string` | Yes | Timestamp when the thread was last active. |
| `subject` | `string | null` | Yes | The subject of the thread. |
| `to` | `any[]` | Yes | The recipients of the most recent email. |

### Field Usage by Operation

| Field | load | list | update | remove |
| --- | --- | --- | --- | --- |
| `bcc` | - | - | - | - |
| `cc` | - | - | - | - |
| `email_count` | - | - | - | - |
| `folder` | - | - | Yes | - |
| `from` | - | - | - | - |
| `has_attachment` | - | - | - | - |
| `has_draft` | - | - | - | - |
| `id` | - | - | - | - |
| `labels` | - | - | - | - |
| `object` | - | - | - | - |
| `read` | - | - | Yes | - |
| `received_at` | - | - | - | - |
| `subject` | - | - | - | - |
| `to` | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `search` | `/inboxes/{inbox_id}/threads/search` | `client.InboxThread().list({ $action: 'search', ... })` |

An action returns that action's OWN response, which is not necessarily a
InboxThread record — check the API definition for its shape.

```ts
const result = await client.InboxThread().list({
  $action: 'search',
  /* ...the action's own arguments */
})
```

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.InboxThread().list({ inbox_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.InboxThread().load({ id: 'inbox_thread_id', inbox_id: 'inbox_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.InboxThread().remove({ id: 'inbox_thread_id', inbox_id: 'inbox_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.InboxThread().update({
  id: 'inbox_thread_id',
  inbox_id: 'inbox_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxThreadEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxThreadEmailEntity

```ts
const inbox_thread_email = client.InboxThreadEmail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any[]` | Yes | The attachments of the message. |
| `bcc` | `any[]` | Yes | The BCC recipients of the message. |
| `cc` | `any[]` | Yes | The CC recipients of the message. |
| `direction` | `string` | Yes | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `string` | Yes | Sender email address. |
| `html` | `string | null` | Yes | The HTML body. |
| `id` | `string` | Yes | The ID of the message. |
| `message_id` | `string | null` | Yes | The Message-ID header of the message. |
| `read` | `boolean` | Yes | Whether the message has been read. |
| `received_at` | `string` | Yes | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `any[]` | Yes | The Reply-To addresses of the message. |
| `subject` | `string | null` | Yes | The subject of the message. |
| `text` | `string | null` | Yes | The plain-text body. |
| `to` | `any[]` | Yes | The recipients of the message. |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `forward` | `/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/forward` | `client.InboxThreadEmail().create({ $action: 'forward', ... })` |
| `reply` | `/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/reply` | `client.InboxThreadEmail().create({ $action: 'reply', ... })` |

An action returns that action's OWN response, which is not necessarily a
InboxThreadEmail record — check the API definition for its shape.

```ts
const result = await client.InboxThreadEmail().create({
  $action: 'forward',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.InboxThreadEmail().list({ inbox_id: "example", thread_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.InboxThreadEmail().load({ id: 'inbox_thread_email_id', inbox_id: 'inbox_id', thread_id: 'thread_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxThreadEmailEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InboxThreadLabelEntity

```ts
const inbox_thread_label = client.InboxThreadLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any[]` | Yes | Every label on the thread after the change. |
| `id` | `string` | No |  |
| `labels` | `any[]` | Yes | The IDs of the labels to apply to the thread. |
| `object` | `string` | Yes | Type of the response object. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.InboxThreadLabel().create({
  inbox_id: 'example_inbox_id',
  thread_id: 'example_thread_id',
  labels: [],
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.InboxThreadLabel().remove({ id: 'inbox_thread_label_id', inbox_id: 'inbox_id', thread_id: 'thread_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.InboxThreadLabel().update({
  inbox_id: 'inbox_id',
  thread_id: 'thread_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InboxThreadLabelEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## LogEntity

```ts
const log = client.Log()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date the log was created. |
| `endpoint` | `string` | No | The API endpoint that was called. |
| `id` | `string` | No | The log ID. |
| `method` | `string` | No | The HTTP method used. |
| `object` | `string` | No | Type of the response object. |
| `request_body` | `Record<string, any> | null` | No | The request body sent to the API. |
| `response_body` | `Record<string, any> | null` | No | The response body returned by the API. |
| `response_status` | `number` | No | The HTTP status code of the response. |
| `user_agent` | `string | null` | No | The user agent of the request. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Log().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Log().load({ id: 'log_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `LogEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OauthGrantEntity

```ts
const oauth_grant = client.OauthGrant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client` | `Record<string, any>` | No | The OAuth client the grant was issued to. |
| `client_id` | `string` | No | The ID of the OAuth client the grant was issued to. |
| `created_at` | `string` | No | The date and time the OAuth grant was created. |
| `id` | `string` | No | The ID of the OAuth grant. |
| `revoked_at` | `string | null` | No | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | `string | null` | No | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | `any[]` | No | The scopes granted to the OAuth client. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.OauthGrant().list()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.OauthGrant().remove({ id: 'oauth_grant_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OauthGrantEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReceivedAttachmentEntity

```ts
const received_attachment = client.ReceivedAttachment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_disposition` | `string | null` | No | How the attachment should be displayed. |
| `content_id` | `string` | No | The content ID for inline attachments. |
| `content_type` | `string` | No | The MIME type of the attachment. |
| `download_url` | `string` | No | Signed URL to download the attachment content. |
| `expires_at` | `string` | No | Timestamp when the download URL expires. |
| `filename` | `string | null` | No | The filename of the attachment. |
| `id` | `string` | No | The ID of the attachment. |
| `object` | `string` | No | The type of object. |
| `size` | `number` | No | Size of the attachment in bytes. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ReceivedAttachment().list({ email_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.ReceivedAttachment().load({ id: 'received_attachment_id', email_id: 'email_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReceivedAttachmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReceivedEmailEntity

```ts
const received_email = client.ReceivedEmail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any[]` | No | Array of attachments. |
| `bcc` | `any[] | null` | No | The BCC recipients. |
| `cc` | `any[] | null` | No | The CC recipients. |
| `created_at` | `string` | No | Timestamp when the email was received. |
| `from` | `string` | No | The sender email address. |
| `headers` | `Record<string, any> | null` | No | The email headers. |
| `html` | `string | null` | No | The HTML content of the email. |
| `id` | `string` | No | The ID of the received email. |
| `message_id` | `string` | No | The unique message ID from the email headers. |
| `object` | `string` | No | The type of object. |
| `received_for` | `any[]` | No | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | `any[] | null` | No | The reply-to addresses. |
| `subject` | `string` | No | The email subject. |
| `text` | `string | null` | No | The plain text content of the email. |
| `to` | `any[]` | No | The recipient email addresses. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.ReceivedEmail().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.ReceivedEmail().load({ email_id: 'email_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReceivedEmailEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SegmentEntity

```ts
const segment = client.Segment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | The ID of the audience this segment belongs to. |
| `created_at` | `string` | No | Timestamp indicating when the segment was created. |
| `filter` | `Record<string, any>` | No | Filter conditions for the segment. |
| `id` | `string` | No | The ID of the segment. |
| `name` | `string` | No | The name of the segment. |
| `object` | `string` | No | The object type. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `audience_id` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `filter` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | Yes | - |
| `object` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Segment().create({
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Segment().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Segment().load({ id: 'segment_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Segment().remove({ id: 'segment_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Segment().update({
  id: 'segment_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SegmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SegmentContactEntity

```ts
const segment_contact = client.SegmentContact()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the contact was created. |
| `email` | `string` | No | Email address of the contact. |
| `first_name` | `string | null` | No | First name of the contact. |
| `id` | `string` | No | Unique identifier for the contact. |
| `last_name` | `string | null` | No | Last name of the contact. |
| `unsubscribed` | `boolean` | No | Indicates if the contact is unsubscribed. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.SegmentContact().list({ segment_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SegmentContactEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SuppressionEntity

```ts
const suppression = client.Suppression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the suppression was created. |
| `email` | `string` | No | Email address that is suppressed. |
| `id` | `string` | No | Unique identifier for the suppression. |
| `object` | `string` | No | Type of the response object. |
| `origin` | `string` | No | Origin of the suppression. |
| `source_id` | `string | null` | No | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

### Field Usage by Operation

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `email` | - | - | Yes | - |
| `id` | - | - | - | - |
| `object` | - | - | - | - |
| `origin` | - | - | - | - |
| `source_id` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Suppression().create({
  email: 'example_email',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Suppression().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Suppression().load({ id: 'suppression_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Suppression().remove({ id: 'suppression_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SuppressionEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SuppressionBatchAddEntity

```ts
const suppression_batch_add = client.SuppressionBatchAdd()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `emails` | `any[]` | Yes | Email addresses to suppress. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.SuppressionBatchAdd().create({
  emails: [],
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SuppressionBatchAddEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SuppressionBatchRemoveEntity

```ts
const suppression_batch_remove = client.SuppressionBatchRemove()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `any[]` | No | Array containing the removed suppressions. |
| `emails` | `any[]` | No | Email addresses to remove from the suppression list. |
| `ids` | `any[]` | No | Suppression IDs to remove from the suppression list. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.SuppressionBatchRemove().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SuppressionBatchRemoveEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TemplateEntity

```ts
const template = client.Template()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `string` | No | The alias of the template. |
| `created_at` | `string` | No | Timestamp indicating when the template was created. |
| `current_version_id` | `string` | No | The ID of the current version of the template. |
| `from` | `string` | No | Sender email address. |
| `has_unpublished_versions` | `boolean` | No | Indicates whether the template has unpublished versions. |
| `html` | `string` | No | The HTML version of the template. |
| `id` | `string` | No | The ID of the template. |
| `name` | `string` | No | The name of the template. |
| `object` | `string` | No | The type of object. |
| `published_at` | `string | null` | No | Timestamp indicating when the template was published. |
| `reply_to` | `any[] | null` | No | Reply-to email addresses. |
| `status` | `string` | No | The publication status of the template. |
| `subject` | `string` | No | Email subject. |
| `text` | `string` | No | The plain text version of the template. |
| `updated_at` | `string` | No | Timestamp indicating when the template was last updated. |
| `variables` | `any[]` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `alias` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `current_version_id` | - | - | - | - | - |
| `from` | - | - | - | - | - |
| `has_unpublished_versions` | - | - | - | - | - |
| `html` | - | - | Yes | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `published_at` | - | - | - | - | - |
| `reply_to` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `subject` | - | - | - | - | - |
| `text` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |
| `variables` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `duplicate` | `/templates/{id}/duplicate` | `client.Template().create({ $action: 'duplicate', ... })` |
| `publish` | `/templates/{id}/publish` | `client.Template().create({ $action: 'publish', ... })` |

An action returns that action's OWN response, which is not necessarily a
Template record — check the API definition for its shape.

```ts
const result = await client.Template().create({
  $action: 'duplicate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Template().create({
  html: 'example_html',
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Template().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Template().load({ id: 'template_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Template().remove({ id: 'template_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Template().update({
  id: 'template_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TemplateEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TopicEntity

```ts
const topic = client.Topic()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the topic was created. |
| `default_subscription` | `string` | No | The default subscription status for the topic. |
| `description` | `string` | No | A description of the topic. |
| `id` | `string` | No | The ID of the topic. |
| `name` | `string` | No | The name of the topic. |
| `object` | `string` | No | The object type. |
| `visibility` | `string` | No | The visibility of the topic. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `default_subscription` | - | - | Yes | - | - |
| `description` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `visibility` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Topic().create({
  default_subscription: 'example_default_subscription',
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Topic().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Topic().load({ id: 'topic_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Topic().remove({ id: 'topic_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Topic().update({
  id: 'topic_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TopicEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UsageEntity

```ts
const usage = client.Usage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ai_credits` | `Record<string, any>` | No |  |
| `automation_runs` | `Record<string, any>` | No |  |
| `broadcasts` | `Record<string, any>` | No |  |
| `contacts` | `Record<string, any>` | No |  |
| `domains` | `Record<string, any>` | No |  |
| `emails` | `Record<string, any>` | No |  |
| `object` | `string` | No | The type of object. |
| `rate_limit` | `Record<string, any>` | No |  |
| `segments` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Usage().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UsageEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebhookEntity

```ts
const webhook = client.Webhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the webhook was created. |
| `endpoint` | `string` | No | The URL where webhook events are sent. |
| `events` | `any[] | null` | No | Array of event types subscribed to. |
| `id` | `string` | No | The ID of the webhook. |
| `object` | `string` | No | The type of object. |
| `signing_secret` | `string` | No | The secret key used to verify webhook payloads. |
| `status` | `string` | No | The status of the webhook. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `endpoint` | - | - | Yes | - | - |
| `events` | - | - | Yes | - | - |
| `id` | - | - | - | - | - |
| `object` | - | - | - | - | - |
| `signing_secret` | - | - | - | - | - |
| `status` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `rotate` | `/webhooks/{webhook_id}/signing-secret/rotate` | `client.Webhook().create({ $action: 'rotate', ... })` |

An action returns that action's OWN response, which is not necessarily a
Webhook record — check the API definition for its shape.

```ts
const result = await client.Webhook().create({
  $action: 'rotate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.Webhook().create({
  endpoint: 'example_endpoint',
  events: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.Webhook().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Webhook().load({ id: 'webhook_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria. Resolves to the entity, marked as deleted.

```ts
const result = await client.Webhook().remove({ id: 'webhook_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`. Resolves to the updated entity.

```ts
const result = await client.Webhook().update({
  id: 'webhook_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebhookEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebhookEventEntity

```ts
const webhook_event = client.WebhookEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the event was created. |
| `id` | `string` | No | The ID of the webhook event. |
| `next_attempt_at` | `string | null` | No | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | `string` | No | The type of object. |
| `payload` | `Record<string, any>` | No | The event payload sent to the webhook endpoint. |
| `status` | `string` | No | The delivery status of the event for this webhook. |
| `type` | `string` | No | The type of the event. |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `replay` | `/webhooks/{webhook_id}/events/{event_id}/replay` | `client.WebhookEvent().create({ $action: 'replay', ... })` |

An action returns that action's OWN response, which is not necessarily a
WebhookEvent record — check the API definition for its shape.

```ts
const result = await client.WebhookEvent().create({
  $action: 'replay',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data. Resolves to the created entity.

```ts
const result = await client.WebhookEvent().create({
  event_id: 'example_event_id',
  webhook_id: 'example_webhook_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.WebhookEvent().list({ webhook_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.WebhookEvent().load({ id: 'webhook_event_id', webhook_id: 'webhook_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebhookEventEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebhookEventAttemptEntity

```ts
const webhook_event_attempt = client.WebhookEventAttempt()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `http_status_code` | `number` | No | The HTTP status code returned by the webhook endpoint. |
| `id` | `string` | No | The ID of the webhook event attempt. |
| `response` | `string` | No | The response body returned by the webhook endpoint. |
| `sent_at` | `string` | No | Timestamp indicating when the attempt was sent. |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Resolves to an array of entities, one per record.

```ts
const results = await client.WebhookEventAttempt().list({ event_id: "example", webhook_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebhookEventAttemptEntity` instance with the same client and
options.

#### `client()`

Return the parent `ResendSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `debug` | 0.0.1 | Debug capture |
| `idempotency` | 0.0.1 | Idempotency |
| `metrics` | 0.0.1 | Metrics |
| `paging` | 0.0.1 | Paging |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```ts
const client = new ResendSDK({
  feature: {
    debug: { active: true },
    idempotency: { active: true },
    metrics: { active: true },
    paging: { active: true },
    ratelimit: { active: true },
    retry: { active: true },
    test: { active: true },
    timeout: { active: true },
  }
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`debug`, `idempotency`, `metrics`, `paging`, `test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `debug`

Debug capture.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

| Option | Type |
|---|---|
| `now` | function |
| `onEntry` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.debug.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `idempotency`

Idempotency.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

| Option | Type |
|---|---|
| `keygen` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.idempotency.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `metrics`

Metrics.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `now` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.metrics.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `paging`

Paging.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

| Option | Type |
|---|---|
| `limit` | number |
| `ops` | list |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.paging.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Inactive by default: leaving it out costs nothing at runtime.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `now` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

