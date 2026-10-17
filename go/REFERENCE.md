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

#### `AddContactToSegmentResponseSuccess(data map[string]any) ResendEntity`

Create a new `AddContactToSegmentResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `AddInboxThreadLabelsResponseSuccess(data map[string]any) ResendEntity`

Create a new `AddInboxThreadLabelsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ApiKey(data map[string]any) ResendEntity`

Create a new `ApiKey` entity instance. Pass `nil` for no initial data.

#### `Audience(data map[string]any) ResendEntity`

Create a new `Audience` entity instance. Pass `nil` for no initial data.

#### `Automation(data map[string]any) ResendEntity`

Create a new `Automation` entity instance. Pass `nil` for no initial data.

#### `AutomationRun(data map[string]any) ResendEntity`

Create a new `AutomationRun` entity instance. Pass `nil` for no initial data.

#### `AutomationRunListItem(data map[string]any) ResendEntity`

Create a new `AutomationRunListItem` entity instance. Pass `nil` for no initial data.

#### `BatchAddSuppressionsResponseSuccess(data map[string]any) ResendEntity`

Create a new `BatchAddSuppressionsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `BatchRemoveSuppressionsResponseSuccess(data map[string]any) ResendEntity`

Create a new `BatchRemoveSuppressionsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `Broadcast(data map[string]any) ResendEntity`

Create a new `Broadcast` entity instance. Pass `nil` for no initial data.

#### `Contact(data map[string]any) ResendEntity`

Create a new `Contact` entity instance. Pass `nil` for no initial data.

#### `ContactImportResponseSuccess(data map[string]any) ResendEntity`

Create a new `ContactImportResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ContactProperty(data map[string]any) ResendEntity`

Create a new `ContactProperty` entity instance. Pass `nil` for no initial data.

#### `ContactTopicsResponseSuccess(data map[string]any) ResendEntity`

Create a new `ContactTopicsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `CreateApiKey(data map[string]any) ResendEntity`

Create a new `CreateApiKey` entity instance. Pass `nil` for no initial data.

#### `CreateBatchEmail(data map[string]any) ResendEntity`

Create a new `CreateBatchEmail` entity instance. Pass `nil` for no initial data.

#### `CreateBroadcastResponseSuccess(data map[string]any) ResendEntity`

Create a new `CreateBroadcastResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `CreateInboxDraftResponseSuccess(data map[string]any) ResendEntity`

Create a new `CreateInboxDraftResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `CreateInboxLabelResponseSuccess(data map[string]any) ResendEntity`

Create a new `CreateInboxLabelResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `Domain(data map[string]any) ResendEntity`

Create a new `Domain` entity instance. Pass `nil` for no initial data.

#### `DomainClaim(data map[string]any) ResendEntity`

Create a new `DomainClaim` entity instance. Pass `nil` for no initial data.

#### `Email(data map[string]any) ResendEntity`

Create a new `Email` entity instance. Pass `nil` for no initial data.

#### `EmailsMetric(data map[string]any) ResendEntity`

Create a new `EmailsMetric` entity instance. Pass `nil` for no initial data.

#### `Event(data map[string]any) ResendEntity`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `Inbox(data map[string]any) ResendEntity`

Create a new `Inbox` entity instance. Pass `nil` for no initial data.

#### `InboxAgentResponseSuccess(data map[string]any) ResendEntity`

Create a new `InboxAgentResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `InboxDraftListItem(data map[string]any) ResendEntity`

Create a new `InboxDraftListItem` entity instance. Pass `nil` for no initial data.

#### `InboxDraftResponseSuccess(data map[string]any) ResendEntity`

Create a new `InboxDraftResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `InboxLabel(data map[string]any) ResendEntity`

Create a new `InboxLabel` entity instance. Pass `nil` for no initial data.

#### `InboxThread(data map[string]any) ResendEntity`

Create a new `InboxThread` entity instance. Pass `nil` for no initial data.

#### `InboxThreadEmail(data map[string]any) ResendEntity`

Create a new `InboxThreadEmail` entity instance. Pass `nil` for no initial data.

#### `InboxThreadEmailResponseSuccess(data map[string]any) ResendEntity`

Create a new `InboxThreadEmailResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `InboxThreadResponseSuccess(data map[string]any) ResendEntity`

Create a new `InboxThreadResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ListAttachment(data map[string]any) ResendEntity`

Create a new `ListAttachment` entity instance. Pass `nil` for no initial data.

#### `ListBroadcastClickedLinksResponseSuccess(data map[string]any) ResendEntity`

Create a new `ListBroadcastClickedLinksResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ListBroadcastRecipientsResponseSuccess(data map[string]any) ResendEntity`

Create a new `ListBroadcastRecipientsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ListContactSegmentsResponseSuccess(data map[string]any) ResendEntity`

Create a new `ListContactSegmentsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ListContactsResponseSuccess(data map[string]any) ResendEntity`

Create a new `ListContactsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `ListWebhookEvent(data map[string]any) ResendEntity`

Create a new `ListWebhookEvent` entity instance. Pass `nil` for no initial data.

#### `ListWebhookEventAttempt(data map[string]any) ResendEntity`

Create a new `ListWebhookEventAttempt` entity instance. Pass `nil` for no initial data.

#### `Log(data map[string]any) ResendEntity`

Create a new `Log` entity instance. Pass `nil` for no initial data.

#### `ReceivedEmail(data map[string]any) ResendEntity`

Create a new `ReceivedEmail` entity instance. Pass `nil` for no initial data.

#### `RemoveAudienceResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveAudienceResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveBroadcastResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveBroadcastResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveContactFromSegmentResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveContactFromSegmentResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveContactPropertyResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveContactPropertyResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveContactResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveContactResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveEvent(data map[string]any) ResendEntity`

Create a new `RemoveEvent` entity instance. Pass `nil` for no initial data.

#### `RemoveInboxDraftResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveInboxDraftResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveInboxLabelResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveInboxLabelResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveInboxResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveInboxResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveInboxThreadLabelResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveInboxThreadLabelResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveInboxThreadResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveInboxThreadResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveSegmentResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveSegmentResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveSuppressionResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveSuppressionResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveTemplateResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveTemplateResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RemoveTopicResponseSuccess(data map[string]any) ResendEntity`

Create a new `RemoveTopicResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `RetrievedAttachment(data map[string]any) ResendEntity`

Create a new `RetrievedAttachment` entity instance. Pass `nil` for no initial data.

#### `RevokeOAuthGrant(data map[string]any) ResendEntity`

Create a new `RevokeOAuthGrant` entity instance. Pass `nil` for no initial data.

#### `Rotate(data map[string]any) ResendEntity`

Create a new `Rotate` entity instance. Pass `nil` for no initial data.

#### `Segment(data map[string]any) ResendEntity`

Create a new `Segment` entity instance. Pass `nil` for no initial data.

#### `SetInboxThreadLabelsResponseSuccess(data map[string]any) ResendEntity`

Create a new `SetInboxThreadLabelsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `Suppression(data map[string]any) ResendEntity`

Create a new `Suppression` entity instance. Pass `nil` for no initial data.

#### `Template(data map[string]any) ResendEntity`

Create a new `Template` entity instance. Pass `nil` for no initial data.

#### `Topic(data map[string]any) ResendEntity`

Create a new `Topic` entity instance. Pass `nil` for no initial data.

#### `UpdateApiKey(data map[string]any) ResendEntity`

Create a new `UpdateApiKey` entity instance. Pass `nil` for no initial data.

#### `UpdateBroadcastResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateBroadcastResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateContactPropertyResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateContactPropertyResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateContactResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateContactResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateContactTopicsResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateContactTopicsResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateDomainResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateDomainResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateEmailOption(data map[string]any) ResendEntity`

Create a new `UpdateEmailOption` entity instance. Pass `nil` for no initial data.

#### `UpdateEvent(data map[string]any) ResendEntity`

Create a new `UpdateEvent` entity instance. Pass `nil` for no initial data.

#### `UpdateInboxDraftResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateInboxDraftResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateInboxLabelResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateInboxLabelResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateInboxResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateInboxResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateInboxThreadResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateInboxThreadResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateSegmentResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateSegmentResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateTemplateResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateTemplateResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateTopicResponseSuccess(data map[string]any) ResendEntity`

Create a new `UpdateTopicResponseSuccess` entity instance. Pass `nil` for no initial data.

#### `UpdateWebhook(data map[string]any) ResendEntity`

Create a new `UpdateWebhook` entity instance. Pass `nil` for no initial data.

#### `Usage(data map[string]any) ResendEntity`

Create a new `Usage` entity instance. Pass `nil` for no initial data.

#### `Webhook(data map[string]any) ResendEntity`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `WebhookEvent(data map[string]any) ResendEntity`

Create a new `WebhookEvent` entity instance. Pass `nil` for no initial data.

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

## AddContactToSegmentResponseSuccessEntity

```go
addContactToSegmentResponseSuccess := client.AddContactToSegmentResponseSuccess(nil)
fmt.Println(addContactToSegmentResponseSuccess.GetName()) // "add_contact_to_segment_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No | The ID of the contact. |
| `object` | `string` | No | The object type. |
| `segment_id` | `string` | No | The ID of the segment. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.AddContactToSegmentResponseSuccess(nil).Create(map[string]any{
    "contact_id": "example_contact_id",
    "segment_id": "example_segment_id",
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

Create a new `AddContactToSegmentResponseSuccessEntity` instance with the same client and
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

## AddInboxThreadLabelsResponseSuccessEntity

```go
addInboxThreadLabelsResponseSuccess := client.AddInboxThreadLabelsResponseSuccess(nil)
fmt.Println(addInboxThreadLabelsResponseSuccess.GetName()) // "add_inbox_thread_labels_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | Yes | Every label on the thread after the change. |
| `labels` | `[]any` | Yes | The IDs of the labels to apply to the thread. |
| `object` | `string` | Yes | Type of the response object. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.AddInboxThreadLabelsResponseSuccess(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
    "thread_id": "example_thread_id",
    "labels": []any{},
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

Create a new `AddInboxThreadLabelsResponseSuccessEntity` instance with the same client and
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

## ApiKeyEntity

```go
apiKey := client.ApiKey(nil)
fmt.Println(apiKey.GetName()) // "api_key"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | The date and time the API key was created. |
| `id` | `string` | No | The ID of the API key. |
| `last_used_at` | `any` | No | The date and time the API key was last used. |
| `name` | `string` | No | The name of the API key. |

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

| Field | load | list | create |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `id` | - | - | - |
| `name` | - | - | Yes |
| `object` | - | - | - |

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

## AutomationRunListItemEntity

```go
automationRunListItem := client.AutomationRunListItem(nil)
fmt.Println(automationRunListItem.GetName()) // "automation_run_list_item"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `completed_at` | `any` | No | The date and time the run completed. |
| `created_at` | `string` | No | The date and time the run was created. |
| `id` | `string` | No | The ID of the automation run. |
| `started_at` | `any` | No | The date and time the run started. |
| `status` | `string` | No | The current status of the automation run. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.AutomationRunListItem(nil).List(map[string]any{"id": "example"}, nil)
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

Create a new `AutomationRunListItemEntity` instance with the same client and
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

## BatchAddSuppressionsResponseSuccessEntity

```go
batchAddSuppressionsResponseSuccess := client.BatchAddSuppressionsResponseSuccess(nil)
fmt.Println(batchAddSuppressionsResponseSuccess.GetName()) // "batch_add_suppressions_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `emails` | `[]any` | Yes | Email addresses to suppress. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.BatchAddSuppressionsResponseSuccess(nil).Create(map[string]any{
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

Create a new `BatchAddSuppressionsResponseSuccessEntity` instance with the same client and
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

## BatchRemoveSuppressionsResponseSuccessEntity

```go
batchRemoveSuppressionsResponseSuccess := client.BatchRemoveSuppressionsResponseSuccess(nil)
fmt.Println(batchRemoveSuppressionsResponseSuccess.GetName()) // "batch_remove_suppressions_response_success"
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
result, err := client.BatchRemoveSuppressionsResponseSuccess(nil).Create(map[string]any{
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

Create a new `BatchRemoveSuppressionsResponseSuccessEntity` instance with the same client and
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
| `preview_text` | `string` | No | The preview text of the email. |
| `reply_to` | `[]any` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | No | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `any` | No | Unique identifier of the segment this broadcast will be sent to. |
| `sent_at` | `string` | No | Timestamp indicating when the broadcast was sent. |
| `status` | `string` | No | The status of the broadcast. |
| `subject` | `string` | No | The subject line of the email. |
| `text` | `any` | No | The plain text version of the broadcast content. |
| `topic_id` | `any` | No | The topic ID that the broadcast is scoped to. |

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

| Field | load | list | create |
| --- | --- | --- | --- |
| `audience_id` | - | - | - |
| `created_at` | - | - | - |
| `email` | - | - | Yes |
| `first_name` | - | - | - |
| `id` | - | - | - |
| `last_name` | - | - | - |
| `object` | - | - | - |
| `properties` | - | - | - |
| `segments` | - | - | - |
| `topics` | - | - | - |
| `unsubscribed` | - | - | - |

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

## ContactImportResponseSuccessEntity

```go
contactImportResponseSuccess := client.ContactImportResponseSuccess(nil)
fmt.Println(contactImportResponseSuccess.GetName()) // "contact_import_response_success"
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
results, err := client.ContactImportResponseSuccess(nil).List(nil, nil)
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
result, err := client.ContactImportResponseSuccess(nil).Load(map[string]any{"id": "contact_import_response_success_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.ContactImportResponseSuccess(nil).Create(map[string]any{
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

Create a new `ContactImportResponseSuccessEntity` instance with the same client and
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

| Field | load | list | create |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `fallback_value` | - | - | - |
| `id` | - | - | - |
| `key` | - | - | Yes |
| `object` | - | - | - |
| `type` | - | - | Yes |

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

## ContactTopicsResponseSuccessEntity

```go
contactTopicsResponseSuccess := client.ContactTopicsResponseSuccess(nil)
fmt.Println(contactTopicsResponseSuccess.GetName()) // "contact_topics_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ContactTopicsResponseSuccess(nil).List(map[string]any{"id": "example"}, nil)
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

Create a new `ContactTopicsResponseSuccessEntity` instance with the same client and
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

## CreateApiKeyEntity

```go
createApiKey := client.CreateApiKey(nil)
fmt.Println(createApiKey.GetName()) // "create_api_key"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `domain_id` | `string` | No | Restrict an API key to send emails only from a specific domain. |
| `name` | `string` | Yes | The API key name. |
| `permission` | `string` | No | The API key can have full access to Resend’s API or be only restricted to send emails. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.CreateApiKey(nil).Create(map[string]any{
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

Create a new `CreateApiKeyEntity` instance with the same client and
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

## CreateBatchEmailEntity

```go
createBatchEmail := client.CreateBatchEmail(nil)
fmt.Println(createBatchEmail.GetName()) // "create_batch_email"
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
result, err := client.CreateBatchEmail(nil).Create(map[string]any{
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

Create a new `CreateBatchEmailEntity` instance with the same client and
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

## CreateBroadcastResponseSuccessEntity

```go
createBroadcastResponseSuccess := client.CreateBroadcastResponseSuccess(nil)
fmt.Println(createBroadcastResponseSuccess.GetName()) // "create_broadcast_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | Use `segment_id` instead. |
| `from` | `string` | Yes | The email address of the sender. |
| `html` | `string` | No | The HTML version of the message. |
| `name` | `string` | No | Name of the broadcast. |
| `preview_text` | `string` | No | The preview text of the email. |
| `reply_to` | `[]any` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | No | Schedule time to send the broadcast. |
| `segment_id` | `string` | Yes | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `bool` | No | Whether to send the broadcast immediately or keep it as a draft. |
| `subject` | `string` | Yes | The subject line of the email. |
| `text` | `string` | No | The plain text version of the message. |
| `topic_id` | `string` | No | The topic ID that the broadcast will be scoped to. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.CreateBroadcastResponseSuccess(nil).Create(map[string]any{
    "from": "example_from",
    "segment_id": "example_segment_id",
    "subject": "example_subject",
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

Create a new `CreateBroadcastResponseSuccessEntity` instance with the same client and
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

## CreateInboxDraftResponseSuccessEntity

```go
createInboxDraftResponseSuccess := client.CreateInboxDraftResponseSuccess(nil)
fmt.Println(createInboxDraftResponseSuccess.GetName()) // "create_inbox_draft_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any` | No | Files to store on the draft. |
| `bcc` | `any` | No | The BCC recipients. |
| `cc` | `any` | No | The CC recipients. |
| `html` | `any` | No | The HTML body. |
| `id` | `string` | Yes | The ID of the draft. |
| `object` | `string` | Yes | Type of the response object. |
| `reply_to_email_id` | `string` | No | The Email ID to reply to, as returned in `data[].id` by the list thread emails endpoint. |
| `subject` | `any` | No | The subject. |
| `text` | `any` | No | The plain text body. |
| `thread_id` | `string` | No | The Thread ID to reply in. |
| `to` | `any` | No | The recipients. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.CreateInboxDraftResponseSuccess(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
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

Create a new `CreateInboxDraftResponseSuccessEntity` instance with the same client and
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

## CreateInboxLabelResponseSuccessEntity

```go
createInboxLabelResponseSuccess := client.CreateInboxLabelResponseSuccess(nil)
fmt.Println(createInboxLabelResponseSuccess.GetName()) // "create_inbox_label_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `string` | No | The color of the label, as a `#RRGGBB` hex color. |
| `name` | `string` | Yes | The name of the label. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.CreateInboxLabelResponseSuccess(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
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

Create a new `CreateInboxLabelResponseSuccessEntity` instance with the same client and
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

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `capabilities` | - | - | - | - |
| `click_tracking` | - | - | - | - |
| `created_at` | - | - | - | - |
| `custom_return_path` | - | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `object` | - | - | - | - |
| `open_tracking` | - | - | - | - |
| `records` | - | - | - | - |
| `region` | - | - | - | - |
| `status` | - | - | - | - |
| `tls` | - | - | - | - |
| `tracking_subdomain` | - | - | - | - |

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

| Field | load | list | create |
| --- | --- | --- | --- |
| `attachments` | - | - | - |
| `bcc` | - | - | - |
| `cc` | - | - | - |
| `created_at` | - | - | - |
| `from` | - | - | Yes |
| `headers` | - | - | - |
| `html` | - | - | - |
| `id` | - | - | - |
| `last_event` | - | - | - |
| `message_id` | - | - | - |
| `object` | - | - | - |
| `reply_to` | - | - | - |
| `scheduled_at` | - | - | - |
| `subject` | - | - | Yes |
| `tags` | - | - | - |
| `template` | - | - | - |
| `text` | - | - | - |
| `to` | - | - | Yes |
| `topic_id` | - | - | - |

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

| Field | load | list | create |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `id` | - | - | - |
| `name` | - | - | Yes |
| `object` | - | - | - |
| `schema` | - | - | - |
| `updated_at` | - | - | - |

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

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `domain_id` | - | - | - | - |
| `drafts` | - | - | - | - |
| `email_address` | - | - | - | - |
| `forwarding` | - | - | - | - |
| `from_name` | - | - | Yes | - |
| `id` | - | - | - | - |
| `last_received` | - | - | - | - |
| `name` | - | - | Yes | - |
| `object` | - | - | - | - |
| `receiving_address` | - | - | - | - |
| `unread` | - | - | - | - |
| `updated_at` | - | - | - | - |

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

## InboxAgentResponseSuccessEntity

```go
inboxAgentResponseSuccess := client.InboxAgentResponseSuccess(nil)
fmt.Println(inboxAgentResponseSuccess.GetName()) // "inbox_agent_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxAgentResponseSuccess(nil).List(map[string]any{"id": "example"}, nil)
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

Create a new `InboxAgentResponseSuccessEntity` instance with the same client and
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

## InboxDraftListItemEntity

```go
inboxDraftListItem := client.InboxDraftListItem(nil)
fmt.Println(inboxDraftListItem.GetName()) // "inbox_draft_list_item"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `[]any` | Yes | The files stored on the draft. |
| `bcc` | `[]any` | Yes | The BCC recipients. |
| `cc` | `[]any` | Yes | The CC recipients. |
| `id` | `string` | Yes | The ID of the draft. |
| `reply_to_email_id` | `any` | Yes | The Email ID the draft replies to. |
| `snippet` | `any` | Yes | A short preview of the body. |
| `subject` | `any` | Yes | The subject. |
| `thread_id` | `any` | Yes | The Thread ID when the draft is a reply. |
| `to` | `any` | Yes | The recipients. |
| `type` | `string` | Yes | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `string` | Yes | Timestamp when the draft was last saved. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxDraftListItem(nil).List(map[string]any{"id": "example"}, nil)
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

Create a new `InboxDraftListItemEntity` instance with the same client and
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

## InboxDraftResponseSuccessEntity

```go
inboxDraftResponseSuccess := client.InboxDraftResponseSuccess(nil)
fmt.Println(inboxDraftResponseSuccess.GetName()) // "inbox_draft_response_success"
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
| `subject` | `any` | Yes | The subject. |
| `text` | `any` | Yes | The plain text body. |
| `thread_id` | `any` | Yes | The Thread ID when the draft is a reply. |
| `to` | `any` | Yes | The recipients. |
| `type` | `string` | Yes | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `string` | Yes | Timestamp when the draft was last saved. |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxDraftResponseSuccess(nil).Load(map[string]any{"draft_id": "draft_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.InboxDraftResponseSuccess(nil).Create(map[string]any{
    "draft_id": "example_draft_id",
    "inbox_id": "example_inbox_id",
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

Create a new `InboxDraftResponseSuccessEntity` instance with the same client and
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
| `color` | `string` | Yes | The color of the label, as a `#RRGGBB` hex color in uppercase. |
| `created_at` | `string` | Yes | Timestamp when the label was created. |
| `id` | `string` | Yes | The ID of the label. |
| `name` | `string` | Yes | The name of the label. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxLabel(nil).List(map[string]any{"id": "example"}, nil)
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
| `matched_email_id` | `any` | Yes | The ID of the email in the thread that matched the search. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `received_at` | `string` | Yes | Timestamp when the thread was last active. |
| `subject` | `any` | Yes | The subject of the thread. |
| `to` | `[]any` | Yes | The recipients of the most recent email. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.InboxThread(nil).List(map[string]any{"id": "example"}, nil)
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

## InboxThreadEmailResponseSuccessEntity

```go
inboxThreadEmailResponseSuccess := client.InboxThreadEmailResponseSuccess(nil)
fmt.Println(inboxThreadEmailResponseSuccess.GetName()) // "inbox_thread_email_response_success"
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxThreadEmailResponseSuccess(nil).Load(map[string]any{"email_id": "email_id", "inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
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

Create a new `InboxThreadEmailResponseSuccessEntity` instance with the same client and
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

## InboxThreadResponseSuccessEntity

```go
inboxThreadResponseSuccess := client.InboxThreadResponseSuccess(nil)
fmt.Println(inboxThreadResponseSuccess.GetName()) // "inbox_thread_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `folder` | `string` | Yes | The folder the thread is in. |
| `id` | `string` | Yes | The ID of the thread. |
| `labels` | `[]any` | Yes | The labels attached to the thread. |
| `object` | `string` | Yes | The object type. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `subject` | `any` | Yes | The subject of the thread. |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.InboxThreadResponseSuccess(nil).Load(map[string]any{"inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
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

Create a new `InboxThreadResponseSuccessEntity` instance with the same client and
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

## ListAttachmentEntity

```go
listAttachment := client.ListAttachment(nil)
fmt.Println(listAttachment.GetName()) // "list_attachment"
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
| `size` | `int` | No | Size of the attachment in bytes. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ListAttachment(nil).List(map[string]any{"email_id": "example"}, nil)
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

Create a new `ListAttachmentEntity` instance with the same client and
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

## ListBroadcastClickedLinksResponseSuccessEntity

```go
listBroadcastClickedLinksResponseSuccess := client.ListBroadcastClickedLinksResponseSuccess(nil)
fmt.Println(listBroadcastClickedLinksResponseSuccess.GetName()) // "list_broadcast_clicked_links_response_success"
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
results, err := client.ListBroadcastClickedLinksResponseSuccess(nil).List(map[string]any{"broadcast_id": "example"}, nil)
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

Create a new `ListBroadcastClickedLinksResponseSuccessEntity` instance with the same client and
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

## ListBroadcastRecipientsResponseSuccessEntity

```go
listBroadcastRecipientsResponseSuccess := client.ListBroadcastRecipientsResponseSuccess(nil)
fmt.Println(listBroadcastRecipientsResponseSuccess.GetName()) // "list_broadcast_recipients_response_success"
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
results, err := client.ListBroadcastRecipientsResponseSuccess(nil).List(map[string]any{"broadcast_id": "example", "type": "example"}, nil)
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

Create a new `ListBroadcastRecipientsResponseSuccessEntity` instance with the same client and
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

## ListContactSegmentsResponseSuccessEntity

```go
listContactSegmentsResponseSuccess := client.ListContactSegmentsResponseSuccess(nil)
fmt.Println(listContactSegmentsResponseSuccess.GetName()) // "list_contact_segments_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the contact was added to the segment. |
| `id` | `string` | No | Unique identifier for the segment. |
| `name` | `string` | No | Name of the segment. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ListContactSegmentsResponseSuccess(nil).List(map[string]any{"contact_id": "example"}, nil)
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

Create a new `ListContactSegmentsResponseSuccessEntity` instance with the same client and
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

## ListContactsResponseSuccessEntity

```go
listContactsResponseSuccess := client.ListContactsResponseSuccess(nil)
fmt.Println(listContactsResponseSuccess.GetName()) // "list_contacts_response_success"
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
results, err := client.ListContactsResponseSuccess(nil).List(map[string]any{"segment_id": "example"}, nil)
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

Create a new `ListContactsResponseSuccessEntity` instance with the same client and
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

## ListWebhookEventEntity

```go
listWebhookEvent := client.ListWebhookEvent(nil)
fmt.Println(listWebhookEvent.GetName()) // "list_webhook_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the event was created. |
| `id` | `string` | No | The ID of the webhook event. |
| `status` | `string` | No | The delivery status of the event for this webhook. |
| `type` | `string` | No | The type of the event. |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.ListWebhookEvent(nil).List(map[string]any{"webhook_id": "example"}, nil)
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

Create a new `ListWebhookEventEntity` instance with the same client and
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

## ListWebhookEventAttemptEntity

```go
listWebhookEventAttempt := client.ListWebhookEventAttempt(nil)
fmt.Println(listWebhookEventAttempt.GetName()) // "list_webhook_event_attempt"
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
results, err := client.ListWebhookEventAttempt(nil).List(map[string]any{"event_id": "example", "webhook_id": "example"}, nil)
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

Create a new `ListWebhookEventAttemptEntity` instance with the same client and
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

## RemoveAudienceResponseSuccessEntity

```go
removeAudienceResponseSuccess := client.RemoveAudienceResponseSuccess(nil)
fmt.Println(removeAudienceResponseSuccess.GetName()) // "remove_audience_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveAudienceResponseSuccess(nil).Remove(map[string]any{"id": "remove_audience_response_success_id"}, nil)
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

Create a new `RemoveAudienceResponseSuccessEntity` instance with the same client and
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

## RemoveBroadcastResponseSuccessEntity

```go
removeBroadcastResponseSuccess := client.RemoveBroadcastResponseSuccess(nil)
fmt.Println(removeBroadcastResponseSuccess.GetName()) // "remove_broadcast_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveBroadcastResponseSuccess(nil).Remove(map[string]any{"id": "remove_broadcast_response_success_id"}, nil)
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

Create a new `RemoveBroadcastResponseSuccessEntity` instance with the same client and
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

## RemoveContactFromSegmentResponseSuccessEntity

```go
removeContactFromSegmentResponseSuccess := client.RemoveContactFromSegmentResponseSuccess(nil)
fmt.Println(removeContactFromSegmentResponseSuccess.GetName()) // "remove_contact_from_segment_response_success"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveContactFromSegmentResponseSuccess(nil).Remove(map[string]any{"contact_id": "contact_id", "segment_id": "segment_id"}, nil)
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

Create a new `RemoveContactFromSegmentResponseSuccessEntity` instance with the same client and
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

## RemoveContactPropertyResponseSuccessEntity

```go
removeContactPropertyResponseSuccess := client.RemoveContactPropertyResponseSuccess(nil)
fmt.Println(removeContactPropertyResponseSuccess.GetName()) // "remove_contact_property_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveContactPropertyResponseSuccess(nil).Remove(map[string]any{"id": "remove_contact_property_response_success_id"}, nil)
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

Create a new `RemoveContactPropertyResponseSuccessEntity` instance with the same client and
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

## RemoveContactResponseSuccessEntity

```go
removeContactResponseSuccess := client.RemoveContactResponseSuccess(nil)
fmt.Println(removeContactResponseSuccess.GetName()) // "remove_contact_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveContactResponseSuccess(nil).Remove(map[string]any{"id": "remove_contact_response_success_id"}, nil)
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

Create a new `RemoveContactResponseSuccessEntity` instance with the same client and
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

## RemoveEventEntity

```go
removeEvent := client.RemoveEvent(nil)
fmt.Println(removeEvent.GetName()) // "remove_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveEvent(nil).Remove(map[string]any{"id": "remove_event_id"}, nil)
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

Create a new `RemoveEventEntity` instance with the same client and
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

## RemoveInboxDraftResponseSuccessEntity

```go
removeInboxDraftResponseSuccess := client.RemoveInboxDraftResponseSuccess(nil)
fmt.Println(removeInboxDraftResponseSuccess.GetName()) // "remove_inbox_draft_response_success"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveInboxDraftResponseSuccess(nil).Remove(map[string]any{"draft_id": "draft_id", "inbox_id": "inbox_id"}, nil)
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

Create a new `RemoveInboxDraftResponseSuccessEntity` instance with the same client and
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

## RemoveInboxLabelResponseSuccessEntity

```go
removeInboxLabelResponseSuccess := client.RemoveInboxLabelResponseSuccess(nil)
fmt.Println(removeInboxLabelResponseSuccess.GetName()) // "remove_inbox_label_response_success"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveInboxLabelResponseSuccess(nil).Remove(map[string]any{"inbox_id": "inbox_id", "label_id": "label_id"}, nil)
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

Create a new `RemoveInboxLabelResponseSuccessEntity` instance with the same client and
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

## RemoveInboxResponseSuccessEntity

```go
removeInboxResponseSuccess := client.RemoveInboxResponseSuccess(nil)
fmt.Println(removeInboxResponseSuccess.GetName()) // "remove_inbox_response_success"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveInboxResponseSuccess(nil).Remove(map[string]any{"inbox_id": "inbox_id"}, nil)
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

Create a new `RemoveInboxResponseSuccessEntity` instance with the same client and
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

## RemoveInboxThreadLabelResponseSuccessEntity

```go
removeInboxThreadLabelResponseSuccess := client.RemoveInboxThreadLabelResponseSuccess(nil)
fmt.Println(removeInboxThreadLabelResponseSuccess.GetName()) // "remove_inbox_thread_label_response_success"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveInboxThreadLabelResponseSuccess(nil).Remove(map[string]any{"inbox_id": "inbox_id", "label_id": "label_id", "thread_id": "thread_id"}, nil)
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

Create a new `RemoveInboxThreadLabelResponseSuccessEntity` instance with the same client and
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

## RemoveInboxThreadResponseSuccessEntity

```go
removeInboxThreadResponseSuccess := client.RemoveInboxThreadResponseSuccess(nil)
fmt.Println(removeInboxThreadResponseSuccess.GetName()) // "remove_inbox_thread_response_success"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria. Returns the entity, marked as deleted; `err` is non-nil on failure.

```go
result, err := client.RemoveInboxThreadResponseSuccess(nil).Remove(map[string]any{"inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
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

Create a new `RemoveInboxThreadResponseSuccessEntity` instance with the same client and
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

## RemoveSegmentResponseSuccessEntity

```go
removeSegmentResponseSuccess := client.RemoveSegmentResponseSuccess(nil)
fmt.Println(removeSegmentResponseSuccess.GetName()) // "remove_segment_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | The ID of the audience this segment belongs to. |
| `created_at` | `string` | No | Timestamp indicating when the segment was created. |
| `filter` | `map[string]any` | No | Filter conditions for the segment. |
| `id` | `string` | No | Unique identifier for the segment. |
| `name` | `string` | Yes | The name of the segment. |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `audience_id` | - | - | - |
| `created_at` | - | - | - |
| `filter` | - | - | - |
| `id` | - | - | - |
| `name` | Yes | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.RemoveSegmentResponseSuccess(nil).List(nil, nil)
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
result, err := client.RemoveSegmentResponseSuccess(nil).Create(map[string]any{
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
result, err := client.RemoveSegmentResponseSuccess(nil).Remove(map[string]any{"id": "remove_segment_response_success_id"}, nil)
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

Create a new `RemoveSegmentResponseSuccessEntity` instance with the same client and
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

## RemoveSuppressionResponseSuccessEntity

```go
removeSuppressionResponseSuccess := client.RemoveSuppressionResponseSuccess(nil)
fmt.Println(removeSuppressionResponseSuccess.GetName()) // "remove_suppression_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the suppression was created. |
| `email` | `string` | Yes | Email address to suppress. |
| `id` | `string` | No | Unique identifier for the suppression. |
| `origin` | `string` | No | Origin of the suppression. |
| `source_id` | `any` | No | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `email` | Yes | - | - |
| `id` | - | - | - |
| `origin` | - | - | - |
| `source_id` | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.RemoveSuppressionResponseSuccess(nil).List(nil, nil)
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
result, err := client.RemoveSuppressionResponseSuccess(nil).Create(map[string]any{
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
result, err := client.RemoveSuppressionResponseSuccess(nil).Remove(map[string]any{"suppression": "suppression"}, nil)
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

Create a new `RemoveSuppressionResponseSuccessEntity` instance with the same client and
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

## RemoveTemplateResponseSuccessEntity

```go
removeTemplateResponseSuccess := client.RemoveTemplateResponseSuccess(nil)
fmt.Println(removeTemplateResponseSuccess.GetName()) // "remove_template_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `string` | No | The alias of the template. |
| `created_at` | `string` | No | Timestamp indicating when the template was created. |
| `from` | `string` | No | Sender email address. |
| `html` | `string` | Yes | The HTML version of the template. |
| `id` | `string` | No | The ID of the template. |
| `name` | `string` | Yes | The name of the template. |
| `published_at` | `any` | No | Timestamp indicating when the template was published. |
| `reply_to` | `[]any` | No | Reply-to email addresses. |
| `status` | `string` | No | The publication status of the template. |
| `subject` | `string` | No | Email subject. |
| `text` | `string` | No | The plain text version of the template. |
| `updated_at` | `string` | No | Timestamp indicating when the template was last updated. |
| `variables` | `[]any` | No |  |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `alias` | - | - | - |
| `created_at` | - | - | - |
| `from` | - | - | - |
| `html` | - | - | - |
| `id` | - | - | - |
| `name` | Yes | - | - |
| `published_at` | - | - | - |
| `reply_to` | - | - | - |
| `status` | - | - | - |
| `subject` | - | - | - |
| `text` | - | - | - |
| `updated_at` | - | - | - |
| `variables` | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.RemoveTemplateResponseSuccess(nil).List(nil, nil)
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
result, err := client.RemoveTemplateResponseSuccess(nil).Create(map[string]any{
    "html": "example_html",
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
result, err := client.RemoveTemplateResponseSuccess(nil).Remove(map[string]any{"id": "remove_template_response_success_id"}, nil)
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

Create a new `RemoveTemplateResponseSuccessEntity` instance with the same client and
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

## RemoveTopicResponseSuccessEntity

```go
removeTopicResponseSuccess := client.RemoveTopicResponseSuccess(nil)
fmt.Println(removeTopicResponseSuccess.GetName()) // "remove_topic_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the topic was created. |
| `default_subscription` | `string` | Yes | The default subscription status for the topic. |
| `description` | `string` | No | A description of the topic. |
| `id` | `string` | No | Unique identifier for the topic. |
| `name` | `string` | Yes | The name of the topic. |
| `visibility` | `string` | No | The visibility of the topic. |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `default_subscription` | Yes | - | - |
| `description` | - | - | - |
| `id` | - | - | - |
| `name` | Yes | - | - |
| `visibility` | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.RemoveTopicResponseSuccess(nil).List(nil, nil)
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
result, err := client.RemoveTopicResponseSuccess(nil).Create(map[string]any{
    "default_subscription": "example_default_subscription",
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
result, err := client.RemoveTopicResponseSuccess(nil).Remove(map[string]any{"id": "remove_topic_response_success_id"}, nil)
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

Create a new `RemoveTopicResponseSuccessEntity` instance with the same client and
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

## RetrievedAttachmentEntity

```go
retrievedAttachment := client.RetrievedAttachment(nil)
fmt.Println(retrievedAttachment.GetName()) // "retrieved_attachment"
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.RetrievedAttachment(nil).Load(map[string]any{"id": "retrieved_attachment_id", "email_id": "email_id"}, nil)
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

Create a new `RetrievedAttachmentEntity` instance with the same client and
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

## RevokeOAuthGrantEntity

```go
revokeOAuthGrant := client.RevokeOAuthGrant(nil)
fmt.Println(revokeOAuthGrant.GetName()) // "revoke_o_auth_grant"
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
results, err := client.RevokeOAuthGrant(nil).List(nil, nil)
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
result, err := client.RevokeOAuthGrant(nil).Remove(map[string]any{"id": "revoke_o_auth_grant_id"}, nil)
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

Create a new `RevokeOAuthGrantEntity` instance with the same client and
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

## RotateEntity

```go
rotate := client.Rotate(nil)
fmt.Println(rotate.GetName()) // "rotate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No | The ID of the webhook. |
| `object` | `string` | No | The type of object. |
| `signing_secret` | `string` | No | The new secret key used to verify webhook payloads. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data. Returns the created entity; `err` is non-nil on failure.

```go
result, err := client.Rotate(nil).Create(map[string]any{
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

Create a new `RotateEntity` instance with the same client and
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

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Segment(nil).Load(map[string]any{"id": "segment_id"}, nil)
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

## SetInboxThreadLabelsResponseSuccessEntity

```go
setInboxThreadLabelsResponseSuccess := client.SetInboxThreadLabelsResponseSuccess(nil)
fmt.Println(setInboxThreadLabelsResponseSuccess.GetName()) // "set_inbox_thread_labels_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `[]any` | Yes | Every label on the thread after the change. |
| `labels` | `[]any` | Yes | The IDs of the labels the thread carries after the change. |
| `object` | `string` | Yes | Type of the response object. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.SetInboxThreadLabelsResponseSuccess(nil).Update(map[string]any{
    "inbox_id": "inbox_id",
    "thread_id": "thread_id",
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

Create a new `SetInboxThreadLabelsResponseSuccessEntity` instance with the same client and
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

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Suppression(nil).Load(map[string]any{"id": "suppression_id"}, nil)
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

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Template(nil).Load(map[string]any{"id": "template_id"}, nil)
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

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Topic(nil).Load(map[string]any{"id": "topic_id"}, nil)
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

## UpdateApiKeyEntity

```go
updateApiKey := client.UpdateApiKey(nil)
fmt.Println(updateApiKey.GetName()) // "update_api_key"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No | The ID of the API key. |
| `name` | `string` | Yes | The API key name. |
| `object` | `string` | No | The type of object. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateApiKey(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateApiKeyEntity` instance with the same client and
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

## UpdateBroadcastResponseSuccessEntity

```go
updateBroadcastResponseSuccess := client.UpdateBroadcastResponseSuccess(nil)
fmt.Println(updateBroadcastResponseSuccess.GetName()) // "update_broadcast_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `string` | No | Use `segment_id` instead. |
| `from` | `string` | No | The email address of the sender. |
| `html` | `string` | No | The HTML version of the message. |
| `id` | `string` | No | The ID of the broadcast. |
| `name` | `string` | No | Name of the broadcast. |
| `object` | `string` | No | The object type of the response. |
| `preview_text` | `string` | No | The preview text of the email. |
| `reply_to` | `[]any` | No | The email addresses to which replies should be sent. |
| `segment_id` | `string` | No | Unique identifier of the segment this broadcast will be sent to. |
| `subject` | `string` | No | The subject line of the email. |
| `text` | `string` | No | The plain text version of the message. |
| `topic_id` | `string` | No | The topic ID that the broadcast will be scoped to. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateBroadcastResponseSuccess(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateBroadcastResponseSuccessEntity` instance with the same client and
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

## UpdateContactPropertyResponseSuccessEntity

```go
updateContactPropertyResponseSuccess := client.UpdateContactPropertyResponseSuccess(nil)
fmt.Println(updateContactPropertyResponseSuccess.GetName()) // "update_contact_property_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fallback_value` | `any` | No | The default value to use when the property is not set for a contact. |
| `id` | `string` | No | The ID of the contact property. |
| `object` | `string` | No | The object type. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateContactPropertyResponseSuccess(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateContactPropertyResponseSuccessEntity` instance with the same client and
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

## UpdateContactResponseSuccessEntity

```go
updateContactResponseSuccess := client.UpdateContactResponseSuccess(nil)
fmt.Println(updateContactResponseSuccess.GetName()) // "update_contact_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | No | Email address of the contact. |
| `first_name` | `string` | No | First name of the contact. |
| `id` | `string` | No | Unique identifier for the updated contact. |
| `last_name` | `string` | No | Last name of the contact. |
| `object` | `string` | No | Type of the response object. |
| `properties` | `map[string]any` | No | A map of custom property keys and values to update. |
| `unsubscribed` | `bool` | No | The Contact's global subscription status. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateContactResponseSuccess(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateContactResponseSuccessEntity` instance with the same client and
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

## UpdateContactTopicsResponseSuccessEntity

```go
updateContactTopicsResponseSuccess := client.UpdateContactTopicsResponseSuccess(nil)
fmt.Println(updateContactTopicsResponseSuccess.GetName()) // "update_contact_topics_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No | The ID of the contact. |
| `object` | `string` | No | The object type. |
| `topics` | `[]any` | No | Array of updated topic subscriptions. |

### Field Usage by Operation

| Field | update |
| --- | --- |
| `contact_id` | - |
| `object` | - |
| `topics` | Yes |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateContactTopicsResponseSuccess(nil).Update(map[string]any{
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

Create a new `UpdateContactTopicsResponseSuccessEntity` instance with the same client and
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

## UpdateDomainResponseSuccessEntity

```go
updateDomainResponseSuccess := client.UpdateDomainResponseSuccess(nil)
fmt.Println(updateDomainResponseSuccess.GetName()) // "update_domain_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capabilities` | `map[string]any` | No | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | No | Track clicks within the body of each HTML email. |
| `id` | `string` | No | The ID of the updated domain. |
| `object` | `string` | No | The object type representing the updated domain. |
| `open_tracking` | `bool` | No | Track the open rate of each email. |
| `tls` | `string` | No | enforced | opportunistic. |
| `tracking_subdomain` | `string` | No | The subdomain to use for click and open tracking. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateDomainResponseSuccess(nil).Update(map[string]any{
    "domain_id": "domain_id",
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

Create a new `UpdateDomainResponseSuccessEntity` instance with the same client and
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

## UpdateEmailOptionEntity

```go
updateEmailOption := client.UpdateEmailOption(nil)
fmt.Println(updateEmailOption.GetName()) // "update_email_option"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `scheduled_at` | `string` | No | Schedule email to be sent later. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateEmailOption(nil).Update(map[string]any{
    "email_id": "email_id",
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

Create a new `UpdateEmailOptionEntity` instance with the same client and
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

## UpdateEventEntity

```go
updateEvent := client.UpdateEvent(nil)
fmt.Println(updateEvent.GetName()) // "update_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No | The ID of the updated event. |
| `object` | `string` | No | Type of the response object. |
| `schema` | `any` | Yes | A flat key/type map defining the event payload schema. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateEvent(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateEventEntity` instance with the same client and
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

## UpdateInboxDraftResponseSuccessEntity

```go
updateInboxDraftResponseSuccess := client.UpdateInboxDraftResponseSuccess(nil)
fmt.Println(updateInboxDraftResponseSuccess.GetName()) // "update_inbox_draft_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `any` | No | Files to store on the draft. |
| `bcc` | `any` | No | The BCC recipients. |
| `cc` | `any` | No | The CC recipients. |
| `html` | `any` | No | The HTML body. |
| `id` | `string` | Yes | The ID of the draft. |
| `object` | `string` | Yes | The object type. |
| `subject` | `any` | No | The subject. |
| `text` | `any` | No | The plain text body. |
| `to` | `any` | No | The recipients. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateInboxDraftResponseSuccess(nil).Update(map[string]any{
    "draft_id": "draft_id",
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

Create a new `UpdateInboxDraftResponseSuccessEntity` instance with the same client and
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

## UpdateInboxLabelResponseSuccessEntity

```go
updateInboxLabelResponseSuccess := client.UpdateInboxLabelResponseSuccess(nil)
fmt.Println(updateInboxLabelResponseSuccess.GetName()) // "update_inbox_label_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `string` | No | A new color for the label, as a `#RRGGBB` hex color. |
| `id` | `string` | Yes | The ID of the updated label. |
| `name` | `string` | No | A new name for the label. |
| `object` | `string` | Yes | Type of the response object. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateInboxLabelResponseSuccess(nil).Update(map[string]any{
    "inbox_id": "inbox_id",
    "label_id": "label_id",
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

Create a new `UpdateInboxLabelResponseSuccessEntity` instance with the same client and
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

## UpdateInboxResponseSuccessEntity

```go
updateInboxResponseSuccess := client.UpdateInboxResponseSuccess(nil)
fmt.Println(updateInboxResponseSuccess.GetName()) // "update_inbox_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from_name` | `string` | No | The name recipients see when mail is sent from this inbox. |
| `id` | `string` | Yes | The ID of the inbox. |
| `name` | `string` | No | New internal name for the inbox. |
| `object` | `string` | Yes | The object type. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateInboxResponseSuccess(nil).Update(map[string]any{
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

Create a new `UpdateInboxResponseSuccessEntity` instance with the same client and
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

## UpdateInboxThreadResponseSuccessEntity

```go
updateInboxThreadResponseSuccess := client.UpdateInboxThreadResponseSuccess(nil)
fmt.Println(updateInboxThreadResponseSuccess.GetName()) // "update_inbox_thread_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `folder` | `string` | Yes | The folder the thread is in. |
| `id` | `string` | Yes | The ID of the thread. |
| `labels` | `[]any` | Yes | The labels attached to the thread. |
| `object` | `string` | Yes | The object type. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `subject` | `any` | Yes | The subject of the thread. |

### Field Usage by Operation

| Field | update |
| --- | --- |
| `folder` | Yes |
| `id` | - |
| `labels` | - |
| `object` | - |
| `read` | Yes |
| `subject` | - |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateInboxThreadResponseSuccess(nil).Update(map[string]any{
    "inbox_id": "inbox_id",
    "thread_id": "thread_id",
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

Create a new `UpdateInboxThreadResponseSuccessEntity` instance with the same client and
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

## UpdateSegmentResponseSuccessEntity

```go
updateSegmentResponseSuccess := client.UpdateSegmentResponseSuccess(nil)
fmt.Println(updateSegmentResponseSuccess.GetName()) // "update_segment_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No | The ID of the segment. |
| `name` | `string` | Yes | The name of the segment. |
| `object` | `string` | No | The object type. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateSegmentResponseSuccess(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateSegmentResponseSuccessEntity` instance with the same client and
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

## UpdateTemplateResponseSuccessEntity

```go
updateTemplateResponseSuccess := client.UpdateTemplateResponseSuccess(nil)
fmt.Println(updateTemplateResponseSuccess.GetName()) // "update_template_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `string` | No | The alias of the template. |
| `from` | `string` | No | Sender email address. |
| `html` | `string` | No | The HTML version of the template. |
| `id` | `string` | No | The ID of the template. |
| `name` | `string` | No | The name of the template. |
| `object` | `string` | No | The object type of the response. |
| `reply_to` | `[]any` | No | Reply-to email addresses. |
| `subject` | `string` | No | Email subject. |
| `text` | `string` | No | The plain text version of the template. |
| `variables` | `[]any` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateTemplateResponseSuccess(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateTemplateResponseSuccessEntity` instance with the same client and
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

## UpdateTopicResponseSuccessEntity

```go
updateTopicResponseSuccess := client.UpdateTopicResponseSuccess(nil)
fmt.Println(updateTopicResponseSuccess.GetName()) // "update_topic_response_success"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | No | A description of the topic. |
| `id` | `string` | No | The ID of the topic. |
| `name` | `string` | No | The name of the topic. |
| `object` | `string` | No | The object type. |
| `visibility` | `string` | No | The visibility of the topic. |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`. Returns the updated entity; `err` is non-nil on failure.

```go
result, err := client.UpdateTopicResponseSuccess(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateTopicResponseSuccessEntity` instance with the same client and
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

## UpdateWebhookEntity

```go
updateWebhook := client.UpdateWebhook(nil)
fmt.Println(updateWebhook.GetName()) // "update_webhook"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No | Timestamp indicating when the webhook was created. |
| `endpoint` | `string` | Yes | The URL where webhook events will be sent. |
| `events` | `[]any` | Yes | Array of event types to subscribe to. |
| `id` | `string` | No | The ID of the updated webhook. |
| `object` | `string` | No | The type of object. |
| `status` | `string` | No | The status of the webhook. |

### Field Usage by Operation

| Field | list | create | update |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `endpoint` | Yes | - | Yes |
| `events` | Yes | - | Yes |
| `id` | - | - | - |
| `object` | - | - | - |
| `status` | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns a `[]any` of entities, one per record; `err` is non-nil on failure.

```go
results, err := client.UpdateWebhook(nil).List(nil, nil)
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
result, err := client.UpdateWebhook(nil).Create(map[string]any{
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
result, err := client.UpdateWebhook(nil).Update(map[string]any{
    "id": "id",
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

Create a new `UpdateWebhookEntity` instance with the same client and
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

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
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

```go
client := sdk.NewResendSDK(map[string]any{
    "feature": map[string]any{
        "debug": map[string]any{"active": true},
        "idempotency": map[string]any{"active": true},
        "metrics": map[string]any{"active": true},
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

