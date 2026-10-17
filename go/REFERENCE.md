# Resend Golang SDK Reference

Complete API reference for the Resend Golang SDK.


## ResendSDK

### Constructor

```go
func NewResendSDK(options map[string]any) *ResendSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *ResendSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *ResendSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `ApiKey(data map[string]any) ResendEntity`

Create a new `ApiKey` entity instance. Pass `nil` for no initial data.

#### `Attachment(data map[string]any) ResendEntity`

Create a new `Attachment` entity instance. Pass `nil` for no initial data.

#### `Audience(data map[string]any) ResendEntity`

Create a new `Audience` entity instance. Pass `nil` for no initial data.

#### `Automation(data map[string]any) ResendEntity`

Create a new `Automation` entity instance. Pass `nil` for no initial data.

#### `AutomationRun(data map[string]any) ResendEntity`

Create a new `AutomationRun` entity instance. Pass `nil` for no initial data.

#### `Broadcast(data map[string]any) ResendEntity`

Create a new `Broadcast` entity instance. Pass `nil` for no initial data.

#### `BroadcastClickedLink(data map[string]any) ResendEntity`

Create a new `BroadcastClickedLink` entity instance. Pass `nil` for no initial data.

#### `BroadcastRecipient(data map[string]any) ResendEntity`

Create a new `BroadcastRecipient` entity instance. Pass `nil` for no initial data.

#### `Contact(data map[string]any) ResendEntity`

Create a new `Contact` entity instance. Pass `nil` for no initial data.

#### `ContactImport(data map[string]any) ResendEntity`

Create a new `ContactImport` entity instance. Pass `nil` for no initial data.

#### `ContactProperty(data map[string]any) ResendEntity`

Create a new `ContactProperty` entity instance. Pass `nil` for no initial data.

#### `ContactSegment(data map[string]any) ResendEntity`

Create a new `ContactSegment` entity instance. Pass `nil` for no initial data.

#### `ContactTopic(data map[string]any) ResendEntity`

Create a new `ContactTopic` entity instance. Pass `nil` for no initial data.

#### `Domain(data map[string]any) ResendEntity`

Create a new `Domain` entity instance. Pass `nil` for no initial data.

#### `DomainClaim(data map[string]any) ResendEntity`

Create a new `DomainClaim` entity instance. Pass `nil` for no initial data.

#### `Email(data map[string]any) ResendEntity`

Create a new `Email` entity instance. Pass `nil` for no initial data.

#### `EmailBatch(data map[string]any) ResendEntity`

Create a new `EmailBatch` entity instance. Pass `nil` for no initial data.

#### `EmailsMetric(data map[string]any) ResendEntity`

Create a new `EmailsMetric` entity instance. Pass `nil` for no initial data.

#### `Event(data map[string]any) ResendEntity`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `Inbox(data map[string]any) ResendEntity`

Create a new `Inbox` entity instance. Pass `nil` for no initial data.

#### `InboxAgent(data map[string]any) ResendEntity`

Create a new `InboxAgent` entity instance. Pass `nil` for no initial data.

#### `InboxDraft(data map[string]any) ResendEntity`

Create a new `InboxDraft` entity instance. Pass `nil` for no initial data.

#### `InboxLabel(data map[string]any) ResendEntity`

Create a new `InboxLabel` entity instance. Pass `nil` for no initial data.

#### `InboxThread(data map[string]any) ResendEntity`

Create a new `InboxThread` entity instance. Pass `nil` for no initial data.

#### `InboxThreadEmail(data map[string]any) ResendEntity`

Create a new `InboxThreadEmail` entity instance. Pass `nil` for no initial data.

#### `InboxThreadLabel(data map[string]any) ResendEntity`

Create a new `InboxThreadLabel` entity instance. Pass `nil` for no initial data.

#### `Log(data map[string]any) ResendEntity`

Create a new `Log` entity instance. Pass `nil` for no initial data.

#### `OauthGrant(data map[string]any) ResendEntity`

Create a new `OauthGrant` entity instance. Pass `nil` for no initial data.

#### `ReceivedAttachment(data map[string]any) ResendEntity`

Create a new `ReceivedAttachment` entity instance. Pass `nil` for no initial data.

#### `ReceivedEmail(data map[string]any) ResendEntity`

Create a new `ReceivedEmail` entity instance. Pass `nil` for no initial data.

#### `Segment(data map[string]any) ResendEntity`

Create a new `Segment` entity instance. Pass `nil` for no initial data.

#### `SegmentContact(data map[string]any) ResendEntity`

Create a new `SegmentContact` entity instance. Pass `nil` for no initial data.

#### `Suppression(data map[string]any) ResendEntity`

Create a new `Suppression` entity instance. Pass `nil` for no initial data.

#### `SuppressionBatchAdd(data map[string]any) ResendEntity`

Create a new `SuppressionBatchAdd` entity instance. Pass `nil` for no initial data.

#### `SuppressionBatchRemove(data map[string]any) ResendEntity`

Create a new `SuppressionBatchRemove` entity instance. Pass `nil` for no initial data.

#### `Template(data map[string]any) ResendEntity`

Create a new `Template` entity instance. Pass `nil` for no initial data.

#### `Topic(data map[string]any) ResendEntity`

Create a new `Topic` entity instance. Pass `nil` for no initial data.

#### `Usage(data map[string]any) ResendEntity`

Create a new `Usage` entity instance. Pass `nil` for no initial data.

#### `Webhook(data map[string]any) ResendEntity`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `WebhookEvent(data map[string]any) ResendEntity`

Create a new `WebhookEvent` entity instance. Pass `nil` for no initial data.

#### `WebhookEventAttempt(data map[string]any) ResendEntity`

Create a new `WebhookEventAttempt` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## ApiKeyEntity

```go
apiKey := client.ApiKey(nil)
fmt.Println(apiKey.GetName()) // "api_key"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date and time the API key was created. |
| `domain_id` | `string` | No | Restrict an API key to send emails only from a specific domain. |
| `id` | `string` | No | The ID of the API key. |
| `last_used_at` | `any` | No | The date and time the API key was last used. |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ApiKey(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.ApiKey(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.ApiKey(nil).Update(map[string]any{
    "id": "api_key_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.ApiKey(nil).Remove(map[string]any{"id": "api_key_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ApiKeyEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## AttachmentEntity

```go
attachment := client.Attachment(nil)
fmt.Println(attachment.GetName()) // "attachment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_disposition` | `any` | No | How the attachment should be displayed. |
| `content_id` | `string` | No | The content ID for inline attachments. |
| `content_type` | `string` | No | The MIME type of the attachment. |
| `download_url` | `string` | No | Signed URL to download the attachment content. |
| `expires_at` | `string` | No | Timestamp when the download URL expires. |
| `filename` | `any` | No | The filename of the attachment. |
| `id` | `string` | No | The ID of the attachment. |
| `object` | `string` | No | The type of object. |
| `size` | `int` | No | Size of the attachment in bytes. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Attachment(nil).List(map[string]any{"email_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Attachment(nil).Load(map[string]any{"id": "attachment_id", "email_id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AttachmentEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## AudienceEntity

```go
audience := client.Audience(nil)
fmt.Println(audience.GetName()) // "audience"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Audience(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Audience(nil).Load(map[string]any{"id": "audience_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Audience(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Audience(nil).Remove(map[string]any{"id": "audience_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AudienceEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## AutomationEntity

```go
automation := client.Automation(nil)
fmt.Println(automation.GetName()) // "automation"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `connections` | `[]any` | No | The connections between steps in the active version of the automation. |
| `created_at` | `string` | No | The date and time the automation was created. |
| `id` | `string` | No | The ID of the automation. |
| `name` | `string` | No | The name of the automation. |
| `object` | `string` | No | Type of the response object. |
| `status` | `string` | No | The current status of the automation. |
| `steps` | `[]any` | No | The steps in the active version of the automation. |
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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Automation(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Automation(nil).Load(map[string]any{"id": "automation_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Automation(nil).Create(map[string]any{
    "connections": []any{},
    "name": "example_name",
    "steps": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Automation(nil).Update(map[string]any{
    "id": "automation_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Automation(nil).Remove(map[string]any{"id": "automation_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AutomationEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## AutomationRunEntity

```go
automationRun := client.AutomationRun(nil)
fmt.Println(automationRun.GetName()) // "automation_run"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `completed_at` | `any` | No | The date and time the run completed. |
| `created_at` | `string` | No | The date and time the run was created. |
| `id` | `string` | No | The ID of the automation run. |
| `object` | `string` | No | Type of the response object. |
| `started_at` | `any` | No | The date and time the run started. |
| `status` | `string` | No | The current status of the automation run. |
| `steps` | `[]any` | No | The steps executed in this run, sorted in graph order. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.AutomationRun(nil).List(map[string]any{"automation_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.AutomationRun(nil).Load(map[string]any{"id": "automation_run_id", "automation_id": "automation_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AutomationRunEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## BroadcastEntity

```go
broadcast := client.Broadcast(nil)
fmt.Println(broadcast.GetName()) // "broadcast"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `any` | No | Deprecated: use `segment_id` instead. |
| `created_at` | `string` | No | Timestamp indicating when the broadcast was created. |
| `from` | `string` | No | The email address of the sender. |
| `html` | `any` | No | The HTML version of the broadcast content. |
| `id` | `string` | No | Unique identifier for the broadcast. |
| `name` | `string` | No | Name of the broadcast. |
| `object` | `string` | No | The object type of the response. |
| `preview_text` | `string` | No | The preview text of the email. |
| `reply_to` | `[]any` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | No | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `any` | No | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `bool` | No | Whether to send the broadcast immediately or keep it as a draft. |
| `sent_at` | `string` | No | Timestamp indicating when the broadcast was sent. |
| `status` | `string` | No | The status of the broadcast. |
| `subject` | `string` | No | The subject line of the email. |
| `text` | `any` | No | The plain text version of the broadcast content. |
| `topic_id` | `any` | No | The topic ID that the broadcast is scoped to. |

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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Broadcast(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Broadcast(nil).Load(map[string]any{"id": "broadcast_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Broadcast(nil).Create(map[string]any{
    "from": "example_from",
    "segment_id": "example_segment_id",
    "subject": "example_subject",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Broadcast(nil).Update(map[string]any{
    "id": "broadcast_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Broadcast(nil).Remove(map[string]any{"id": "broadcast_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BroadcastEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## BroadcastClickedLinkEntity

```go
broadcastClickedLink := client.BroadcastClickedLink(nil)
fmt.Println(broadcastClickedLink.GetName()) // "broadcast_clicked_link"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clicks` | `int` | No | Total number of clicks on this URL. |
| `id` | `string` | No | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `int` | No | Number of unique clicks on this URL. |
| `url` | `string` | No | The URL that was clicked. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.BroadcastClickedLink(nil).List(map[string]any{"broadcast_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BroadcastClickedLinkEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## BroadcastRecipientEntity

```go
broadcastRecipient := client.BroadcastRecipient(nil)
fmt.Println(broadcastRecipient.GetName()) // "broadcast_recipient"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bounce_type` | `string` | No | The type of bounce. |
| `clicked_links` | `[]any` | No | The links this recipient clicked. |
| `contact_id` | `any` | No | The ID of the contact associated with this recipient, if one exists. |
| `count` | `int` | No | The number of times this recipient triggered the event. |
| `email` | `string` | No | The recipient's email address. |
| `id` | `string` | No | Opaque cursor identifying this row, used for pagination. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.BroadcastRecipient(nil).List(map[string]any{"broadcast_id": "example", "type": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BroadcastRecipientEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ContactEntity

```go
contact := client.Contact(nil)
fmt.Println(contact.GetName()) // "contact"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | Unique identifier of the audience to which the contact belongs. |
| `created_at` | `string` | No | Timestamp indicating when the contact was created. |
| `email` | `string` | No | Email address of the contact. |
| `first_name` | `any` | No | First name of the contact. |
| `id` | `string` | No | Unique identifier for the contact. |
| `last_name` | `any` | No | Last name of the contact. |
| `object` | `string` | No | Type of the response object. |
| `properties` | `map[string]any` | No | A map of custom property keys and values. |
| `segments` | `[]any` | No | Array of segment IDs to add the contact to. |
| `topics` | `[]any` | No | Array of topic subscriptions for the contact. |
| `unsubscribed` | `bool` | No | Indicates if the contact is unsubscribed. |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Contact(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Contact(nil).Load(map[string]any{"id": "contact_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Contact(nil).Create(map[string]any{
    "email": "example_email",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Contact(nil).Update(map[string]any{
    "id": "contact_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Contact(nil).Remove(map[string]any{"id": "contact_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ContactImportEntity

```go
contactImport := client.ContactImport(nil)
fmt.Println(contactImport.GetName()) // "contact_import"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `column_map` | `string` | No | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | `any` | No | Timestamp indicating when the contact import completed. |
| `counts` | `map[string]any` | No |  |
| `created_at` | `string` | No | Timestamp indicating when the contact import was created. |
| `file` | `string` | Yes | CSV file to import. |
| `id` | `string` | No | Unique identifier for the contact import. |
| `object` | `string` | No | Type of the response object. |
| `on_conflict` | `string` | No | Strategy to use when an imported contact already exists. |
| `segments` | `string` | No | JSON-encoded array of segments to add imported contacts to. |
| `status` | `string` | No | Current status of the contact import. |
| `topics` | `string` | No | JSON-encoded array of topic subscriptions to apply to imported contacts. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ContactImport(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.ContactImport(nil).Load(map[string]any{"id": "contact_import_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.ContactImport(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactImportEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ContactPropertyEntity

```go
contactProperty := client.ContactProperty(nil)
fmt.Println(contactProperty.GetName()) // "contact_property"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ContactProperty(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.ContactProperty(nil).Load(map[string]any{"id": "contact_property_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.ContactProperty(nil).Create(map[string]any{
    "key": "example_key",
    "type": "example_type",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.ContactProperty(nil).Update(map[string]any{
    "id": "contact_property_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.ContactProperty(nil).Remove(map[string]any{"id": "contact_property_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactPropertyEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ContactSegmentEntity

```go
contactSegment := client.ContactSegment(nil)
fmt.Println(contactSegment.GetName()) // "contact_segment"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ContactSegment(nil).List(map[string]any{"contact_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.ContactSegment(nil).Create(map[string]any{
    "contact_id": "example_contact_id",
    "id": "example_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.ContactSegment(nil).Remove(map[string]any{"id": "contact_segment_id", "contact_id": "contact_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactSegmentEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ContactTopicEntity

```go
contactTopic := client.ContactTopic(nil)
fmt.Println(contactTopic.GetName()) // "contact_topic"
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
| `topics` | `[]any` | No | Array of updated topic subscriptions. |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ContactTopic(nil).List(map[string]any{"contact_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.ContactTopic(nil).Update(map[string]any{
    "contact_id": "contact_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactTopicEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## DomainEntity

```go
domain := client.Domain(nil)
fmt.Println(domain.GetName()) // "domain"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capabilities` | `map[string]any` | No | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | No | Whether click tracking is enabled for this domain. |
| `created_at` | `string` | No | The date and time the domain was created. |
| `custom_return_path` | `string` | No | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | `string` | No | The ID of the domain. |
| `name` | `string` | No | The name of the domain. |
| `object` | `string` | No | The type of object. |
| `open_tracking` | `bool` | No | Whether open tracking is enabled for this domain. |
| `records` | `[]any` | No |  |
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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Domain(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Domain(nil).Load(map[string]any{"id": "domain_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Domain(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Domain(nil).Update(map[string]any{
    "id": "domain_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Domain(nil).Remove(map[string]any{"id": "domain_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DomainEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## DomainClaimEntity

```go
domainClaim := client.DomainClaim(nil)
fmt.Println(domainClaim.GetName()) // "domain_claim"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `blocked_reason` | `any` | No | Why the claim is currently blocked, if applicable. |
| `click_tracking` | `bool` | No | Track clicks within the body of each HTML email. |
| `created_at` | `string` | No | The date and time the claim was created. |
| `custom_return_path` | `string` | No | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | `any` | No | The ID of the placeholder domain created for the claim. |
| `expires_at` | `string` | No | The date and time the claim expires if not verified. |
| `failure_reason` | `any` | No | Why the claim failed, if applicable. |
| `id` | `string` | No | The ID of the claim. |
| `name` | `string` | No | The name of the domain being claimed. |
| `object` | `string` | No | The type of object. |
| `open_tracking` | `bool` | No | Track the open rate of each email. |
| `record` | `map[string]any` | No | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | `any` | No | The region where the claimed domain will send from. |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.DomainClaim(nil).Load(map[string]any{"id": "domain_claim_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.DomainClaim(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DomainClaimEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## EmailEntity

```go
email := client.Email(nil)
fmt.Println(email.GetName()) // "email"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `[]any` | No |  |
| `bcc` | `[]any` | No | The email addresses of the blind carbon copy recipients. |
| `cc` | `[]any` | No | The email addresses of the carbon copy recipients. |
| `created_at` | `string` | No | The date and time the email was created. |
| `from` | `string` | No | The email address of the sender. |
| `headers` | `map[string]any` | No | Custom headers to add to the email. |
| `html` | `string` | No | The HTML body of the email. |
| `id` | `string` | No | The ID of the email. |
| `last_event` | `string` | No | The status of the email. |
| `message_id` | `string` | No | The Message-ID header value of the email. |
| `object` | `string` | No | The type of object. |
| `reply_to` | `[]any` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | No | Schedule email to be sent later. |
| `subject` | `string` | No | The subject line of the email. |
| `tags` | `[]any` | No |  |
| `template` | `any` | No |  |
| `text` | `string` | No | The plain text body of the email. |
| `to` | `[]any` | No | Recipient email address. |
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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Email(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Email(nil).Load(map[string]any{"id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Email(nil).Create(map[string]any{
    "from": "example_from",
    "subject": "example_subject",
    "to": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Email(nil).Update(map[string]any{
    "id": "email_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## EmailBatchEntity

```go
emailBatch := client.EmailBatch(nil)
fmt.Println(emailBatch.GetName()) // "email_batch"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | No |  |
| `send_emails` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.EmailBatch(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailBatchEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## EmailsMetricEntity

```go
emailsMetric := client.EmailsMetric(nil)
fmt.Println(emailsMetric.GetName()) // "emails_metric"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.EmailsMetric(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EmailsMetricEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## EventEntity

```go
event := client.Event(nil)
fmt.Println(event.GetName()) // "event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date and time the event was created. |
| `id` | `string` | No | The event ID. |
| `name` | `string` | No | The event name. |
| `object` | `string` | No | Type of the response object. |
| `schema` | `any` | No | A flat key/type map defining the event payload schema. |
| `updated_at` | `any` | No | The date and time the event was last updated. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `schema` | - | - | - | Yes | - |
| `updated_at` | - | - | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Event(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Event(nil).Load(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Event(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Event(nil).Update(map[string]any{
    "id": "event_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Event(nil).Remove(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxEntity

```go
inbox := client.Inbox(nil)
fmt.Println(inbox.GetName()) // "inbox"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | `string` | Yes | The ID of the domain that owns the address. |
| `drafts` | `int` | Yes | The number of unsent drafts. |
| `email_address` | `string` | Yes | The address of the inbox. |
| `forwarding` | `bool` | No | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | `any` | Yes | The name recipients see when mail is sent from this inbox. |
| `id` | `string` | Yes | The ID of the inbox. |
| `last_received` | `any` | Yes | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | `string` | Yes | Internal name for the inbox. |
| `object` | `string` | Yes | The object type. |
| `receiving_address` | `any` | Yes | The address to forward mail to when forwarding is enabled. |
| `unread` | `int` | Yes | The number of unread threads in the Inbox folder. |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Inbox(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Inbox(nil).Load(map[string]any{"id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Inbox(nil).Create(map[string]any{
    "email_address": "example_email_address",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Inbox(nil).Update(map[string]any{
    "id": "inbox_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Inbox(nil).Remove(map[string]any{"id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxAgentEntity

```go
inboxAgent := client.InboxAgent(nil)
fmt.Println(inboxAgent.GetName()) // "inbox_agent"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `[]any` | Yes | The actions the agent can take, in the order of the enum. |
| `id` | `string` | Yes | The ID of the inbox agent. |
| `instructions` | `any` | Yes | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | `string` | Yes | The object type. |
| `tone` | `any` | Yes | The tone the agent writes in. |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `actions` | - | Yes |
| `id` | - | - |
| `instructions` | - | Yes |
| `object` | - | - |
| `tone` | - | Yes |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxAgent(nil).Load(map[string]any{"inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.InboxAgent(nil).Update(map[string]any{
    "inbox_id": "inbox_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxAgentEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxDraftEntity

```go
inboxDraft := client.InboxDraft(nil)
fmt.Println(inboxDraft.GetName()) // "inbox_draft"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `[]any` | Yes | The files stored on the draft. |
| `bcc` | `[]any` | Yes | The BCC recipients. |
| `cc` | `[]any` | Yes | The CC recipients. |
| `created_at` | `string` | Yes | Timestamp when the draft was created. |
| `email_id` | `any` | Yes | The ID of the outbound email queued when the draft was sent. |
| `html` | `any` | Yes | The HTML body. |
| `id` | `string` | Yes | The ID of the draft. |
| `object` | `string` | Yes | The object type. |
| `reply_to_email_id` | `any` | Yes | The Email ID the draft replies to. |
| `snippet` | `any` | Yes | A short preview of the body. |
| `subject` | `any` | Yes | The subject. |
| `text` | `any` | Yes | The plain text body. |
| `thread_id` | `any` | Yes | The Thread ID when the draft is a reply. |
| `to` | `any` | Yes | The recipients. |
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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxDraft(nil).List(map[string]any{"inbox_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxDraft(nil).Load(map[string]any{"id": "inbox_draft_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.InboxDraft(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.InboxDraft(nil).Update(map[string]any{
    "id": "inbox_draft_id",
    "inbox_id": "inbox_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.InboxDraft(nil).Remove(map[string]any{"id": "inbox_draft_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxDraftEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxLabelEntity

```go
inboxLabel := client.InboxLabel(nil)
fmt.Println(inboxLabel.GetName()) // "inbox_label"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxLabel(nil).List(map[string]any{"inbox_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.InboxLabel(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.InboxLabel(nil).Update(map[string]any{
    "id": "inbox_label_id",
    "inbox_id": "inbox_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.InboxLabel(nil).Remove(map[string]any{"id": "inbox_label_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxLabelEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxThreadEntity

```go
inboxThread := client.InboxThread(nil)
fmt.Println(inboxThread.GetName()) // "inbox_thread"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bcc` | `[]any` | Yes | The BCC recipients of the most recent email. |
| `cc` | `[]any` | Yes | The CC recipients of the most recent email. |
| `email_count` | `int` | Yes | The number of emails in the thread. |
| `folder` | `string` | Yes | The folder the thread is in. |
| `from` | `any` | Yes | The sender email address. |
| `has_attachment` | `bool` | Yes | Whether the most recent email in the thread has an attachment. |
| `has_draft` | `bool` | Yes | Whether the thread has an unsent draft. |
| `id` | `string` | Yes | The ID of the thread. |
| `labels` | `[]any` | Yes | The labels attached to the thread. |
| `object` | `string` | Yes | The object type. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `received_at` | `string` | Yes | Timestamp when the thread was last active. |
| `subject` | `any` | Yes | The subject of the thread. |
| `to` | `[]any` | Yes | The recipients of the most recent email. |

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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxThread(nil).List(map[string]any{"inbox_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxThread(nil).Load(map[string]any{"id": "inbox_thread_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.InboxThread(nil).Update(map[string]any{
    "id": "inbox_thread_id",
    "inbox_id": "inbox_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.InboxThread(nil).Remove(map[string]any{"id": "inbox_thread_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxThreadEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxThreadEmailEntity

```go
inboxThreadEmail := client.InboxThreadEmail(nil)
fmt.Println(inboxThreadEmail.GetName()) // "inbox_thread_email"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `[]any` | Yes | The attachments of the message. |
| `bcc` | `[]any` | Yes | The BCC recipients of the message. |
| `cc` | `[]any` | Yes | The CC recipients of the message. |
| `direction` | `string` | Yes | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `string` | Yes | Sender email address. |
| `html` | `any` | Yes | The HTML body. |
| `id` | `string` | Yes | The ID of the message. |
| `message_id` | `any` | Yes | The Message-ID header of the message. |
| `read` | `bool` | Yes | Whether the message has been read. |
| `received_at` | `string` | Yes | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `[]any` | Yes | The Reply-To addresses of the message. |
| `subject` | `any` | Yes | The subject of the message. |
| `text` | `any` | Yes | The plain-text body. |
| `to` | `[]any` | Yes | The recipients of the message. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxThreadEmail(nil).List(map[string]any{"inbox_id": "example", "thread_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxThreadEmail(nil).Load(map[string]any{"id": "inbox_thread_email_id", "inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxThreadEmailEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## InboxThreadLabelEntity

```go
inboxThreadLabel := client.InboxThreadLabel(nil)
fmt.Println(inboxThreadLabel.GetName()) // "inbox_thread_label"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | Yes | Every label on the thread after the change. |
| `id` | `string` | No |  |
| `labels` | `[]any` | Yes | The IDs of the labels to apply to the thread. |
| `object` | `string` | Yes | Type of the response object. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.InboxThreadLabel(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
    "thread_id": "example_thread_id",
    "labels": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.InboxThreadLabel(nil).Update(map[string]any{
    "inbox_id": "inbox_id",
    "thread_id": "thread_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.InboxThreadLabel(nil).Remove(map[string]any{"id": "inbox_thread_label_id", "inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InboxThreadLabelEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## LogEntity

```go
log := client.Log(nil)
fmt.Println(log.GetName()) // "log"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date the log was created. |
| `endpoint` | `string` | No | The API endpoint that was called. |
| `id` | `string` | No | The log ID. |
| `method` | `string` | No | The HTTP method used. |
| `object` | `string` | No | Type of the response object. |
| `request_body` | `any` | No | The request body sent to the API. |
| `response_body` | `any` | No | The response body returned by the API. |
| `response_status` | `int` | No | The HTTP status code of the response. |
| `user_agent` | `any` | No | The user agent of the request. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Log(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Log(nil).Load(map[string]any{"id": "log_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `LogEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## OauthGrantEntity

```go
oauthGrant := client.OauthGrant(nil)
fmt.Println(oauthGrant.GetName()) // "oauth_grant"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client` | `map[string]any` | No | The OAuth client the grant was issued to. |
| `client_id` | `string` | No | The ID of the OAuth client the grant was issued to. |
| `created_at` | `string` | No | The date and time the OAuth grant was created. |
| `id` | `string` | No | The ID of the OAuth grant. |
| `revoked_at` | `any` | No | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | `any` | No | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | `[]any` | No | The scopes granted to the OAuth client. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.OauthGrant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.OauthGrant(nil).Remove(map[string]any{"id": "oauth_grant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OauthGrantEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ReceivedAttachmentEntity

```go
receivedAttachment := client.ReceivedAttachment(nil)
fmt.Println(receivedAttachment.GetName()) // "received_attachment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_disposition` | `any` | No | How the attachment should be displayed. |
| `content_id` | `string` | No | The content ID for inline attachments. |
| `content_type` | `string` | No | The MIME type of the attachment. |
| `download_url` | `string` | No | Signed URL to download the attachment content. |
| `expires_at` | `string` | No | Timestamp when the download URL expires. |
| `filename` | `any` | No | The filename of the attachment. |
| `id` | `string` | No | The ID of the attachment. |
| `object` | `string` | No | The type of object. |
| `size` | `int` | No | Size of the attachment in bytes. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ReceivedAttachment(nil).List(map[string]any{"email_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.ReceivedAttachment(nil).Load(map[string]any{"id": "received_attachment_id", "email_id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReceivedAttachmentEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ReceivedEmailEntity

```go
receivedEmail := client.ReceivedEmail(nil)
fmt.Println(receivedEmail.GetName()) // "received_email"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `[]any` | No | Array of attachments. |
| `bcc` | `any` | No | The BCC recipients. |
| `cc` | `any` | No | The CC recipients. |
| `created_at` | `string` | No | Timestamp when the email was received. |
| `from` | `string` | No | The sender email address. |
| `headers` | `any` | No | The email headers. |
| `html` | `any` | No | The HTML content of the email. |
| `id` | `string` | No | The ID of the received email. |
| `message_id` | `string` | No | The unique message ID from the email headers. |
| `object` | `string` | No | The type of object. |
| `received_for` | `[]any` | No | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | `any` | No | The reply-to addresses. |
| `subject` | `string` | No | The email subject. |
| `text` | `any` | No | The plain text content of the email. |
| `to` | `[]any` | No | The recipient email addresses. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ReceivedEmail(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.ReceivedEmail(nil).Load(map[string]any{"email_id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReceivedEmailEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SegmentEntity

```go
segment := client.Segment(nil)
fmt.Println(segment.GetName()) // "segment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | The ID of the audience this segment belongs to. |
| `created_at` | `string` | No | Timestamp indicating when the segment was created. |
| `filter` | `map[string]any` | No | Filter conditions for the segment. |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Segment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Segment(nil).Load(map[string]any{"id": "segment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Segment(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Segment(nil).Update(map[string]any{
    "id": "segment_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Segment(nil).Remove(map[string]any{"id": "segment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SegmentEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SegmentContactEntity

```go
segmentContact := client.SegmentContact(nil)
fmt.Println(segmentContact.GetName()) // "segment_contact"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the contact was created. |
| `email` | `string` | No | Email address of the contact. |
| `first_name` | `any` | No | First name of the contact. |
| `id` | `string` | No | Unique identifier for the contact. |
| `last_name` | `any` | No | Last name of the contact. |
| `unsubscribed` | `bool` | No | Indicates if the contact is unsubscribed. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.SegmentContact(nil).List(map[string]any{"segment_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SegmentContactEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SuppressionEntity

```go
suppression := client.Suppression(nil)
fmt.Println(suppression.GetName()) // "suppression"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the suppression was created. |
| `email` | `string` | No | Email address that is suppressed. |
| `id` | `string` | No | Unique identifier for the suppression. |
| `object` | `string` | No | Type of the response object. |
| `origin` | `string` | No | Origin of the suppression. |
| `source_id` | `any` | No | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Suppression(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Suppression(nil).Load(map[string]any{"id": "suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Suppression(nil).Create(map[string]any{
    "email": "example_email",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Suppression(nil).Remove(map[string]any{"id": "suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SuppressionEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SuppressionBatchAddEntity

```go
suppressionBatchAdd := client.SuppressionBatchAdd(nil)
fmt.Println(suppressionBatchAdd.GetName()) // "suppression_batch_add"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `emails` | `[]any` | Yes | Email addresses to suppress. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.SuppressionBatchAdd(nil).Create(map[string]any{
    "emails": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SuppressionBatchAddEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SuppressionBatchRemoveEntity

```go
suppressionBatchRemove := client.SuppressionBatchRemove(nil)
fmt.Println(suppressionBatchRemove.GetName()) // "suppression_batch_remove"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | No | Array containing the removed suppressions. |
| `emails` | `[]any` | No | Email addresses to remove from the suppression list. |
| `ids` | `[]any` | No | Suppression IDs to remove from the suppression list. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.SuppressionBatchRemove(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SuppressionBatchRemoveEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## TemplateEntity

```go
template := client.Template(nil)
fmt.Println(template.GetName()) // "template"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `string` | No | The alias of the template. |
| `created_at` | `string` | No | Timestamp indicating when the template was created. |
| `current_version_id` | `string` | No | The ID of the current version of the template. |
| `from` | `string` | No | Sender email address. |
| `has_unpublished_versions` | `bool` | No | Indicates whether the template has unpublished versions. |
| `html` | `string` | No | The HTML version of the template. |
| `id` | `string` | No | The ID of the template. |
| `name` | `string` | No | The name of the template. |
| `object` | `string` | No | The type of object. |
| `published_at` | `any` | No | Timestamp indicating when the template was published. |
| `reply_to` | `any` | No | Reply-to email addresses. |
| `status` | `string` | No | The publication status of the template. |
| `subject` | `string` | No | Email subject. |
| `text` | `string` | No | The plain text version of the template. |
| `updated_at` | `string` | No | Timestamp indicating when the template was last updated. |
| `variables` | `[]any` | No |  |

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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Template(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Template(nil).Load(map[string]any{"id": "template_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Template(nil).Create(map[string]any{
    "html": "example_html",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Template(nil).Update(map[string]any{
    "id": "template_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Template(nil).Remove(map[string]any{"id": "template_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TemplateEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## TopicEntity

```go
topic := client.Topic(nil)
fmt.Println(topic.GetName()) // "topic"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Topic(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Topic(nil).Load(map[string]any{"id": "topic_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Topic(nil).Create(map[string]any{
    "default_subscription": "example_default_subscription",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Topic(nil).Update(map[string]any{
    "id": "topic_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Topic(nil).Remove(map[string]any{"id": "topic_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TopicEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## UsageEntity

```go
usage := client.Usage(nil)
fmt.Println(usage.GetName()) // "usage"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ai_credits` | `map[string]any` | No |  |
| `automation_runs` | `map[string]any` | No |  |
| `broadcasts` | `map[string]any` | No |  |
| `contacts` | `map[string]any` | No |  |
| `domains` | `map[string]any` | No |  |
| `emails` | `map[string]any` | No |  |
| `object` | `string` | No | The type of object. |
| `rate_limit` | `map[string]any` | No |  |
| `segments` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Usage(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UsageEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## WebhookEntity

```go
webhook := client.Webhook(nil)
fmt.Println(webhook.GetName()) // "webhook"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the webhook was created. |
| `endpoint` | `string` | No | The URL where webhook events are sent. |
| `events` | `any` | No | Array of event types subscribed to. |
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

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.Webhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Webhook(nil).Create(map[string]any{
    "endpoint": "example_endpoint",
    "events": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.Webhook(nil).Update(map[string]any{
    "id": "webhook_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.Webhook(nil).Remove(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## WebhookEventEntity

```go
webhookEvent := client.WebhookEvent(nil)
fmt.Println(webhookEvent.GetName()) // "webhook_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the event was created. |
| `id` | `string` | No | The ID of the webhook event. |
| `next_attempt_at` | `any` | No | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | `string` | No | The type of object. |
| `payload` | `map[string]any` | No | The event payload sent to the webhook endpoint. |
| `status` | `string` | No | The delivery status of the event for this webhook. |
| `type` | `string` | No | The type of the event. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.WebhookEvent(nil).List(map[string]any{"webhook_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.WebhookEvent(nil).Load(map[string]any{"id": "webhook_event_id", "webhook_id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.WebhookEvent(nil).Create(map[string]any{
    "event_id": "example_event_id",
    "webhook_id": "example_webhook_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebhookEventEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## WebhookEventAttemptEntity

```go
webhookEventAttempt := client.WebhookEventAttempt(nil)
fmt.Println(webhookEventAttempt.GetName()) // "webhook_event_attempt"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `http_status_code` | `int` | No | The HTTP status code returned by the webhook endpoint. |
| `id` | `string` | No | The ID of the webhook event attempt. |
| `response` | `string` | No | The response body returned by the webhook endpoint. |
| `sent_at` | `string` | No | Timestamp indicating when the attempt was sent. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.WebhookEventAttempt(nil).List(map[string]any{"event_id": "example", "webhook_id": "example"}, nil)
if err != nil {
    panic(err)
}
for _, item := range results.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebhookEventAttemptEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `debug` | 0.0.1 | Debug capture |
| `idempotency` | 0.0.1 | Idempotency |
| `metrics` | 0.0.1 | Metrics |
| `netsim` | 0.0.1 | Network simulation |
| `paging` | 0.0.1 | Paging |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```go
client := sdk.NewResendSDK(map[string]any{
    "feature": map[string]any{
        "debug": map[string]any{"active": true},
        "idempotency": map[string]any{"active": true},
        "metrics": map[string]any{"active": true},
        "netsim": map[string]any{"active": true},
        "paging": map[string]any{"active": true},
        "ratelimit": map[string]any{"active": true},
        "retry": map[string]any{"active": true},
        "test": map[string]any{"active": true},
        "timeout": map[string]any{"active": true},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`netsim`, `ratelimit`, `retry`, `timeout` wrap the transport. Each
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

#### `netsim`

Network simulation.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `errorTimes` | `0` |
| `failEvery` | `0` |
| `failRate` | `0` |
| `failStatus` | `503` |
| `failTimes` | `0` |
| `latency` | `0` |
| `offline` | `false` |
| `rateLimitTimes` | `0` |
| `retryAfter` | `0` |
| `seed` | `1` |

| Option | Type |
|---|---|
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.netsim.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
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

