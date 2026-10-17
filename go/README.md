# Resend Golang SDK



The Golang SDK for the Resend API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.AddContactToSegmentResponseSuccess(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `py`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/resend-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Tags](https://github.com/voxgig-sdk/resend-sdk/tags) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/resend-sdk/go=../resend-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the entity, and for
`List` a `[]any` of entities, one per record (there is no `{ok, data}`
wrapper), so check `err` and read a record through the entity's
`Data()`.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/resend-sdk/go"
)

func main() {
    client := sdk.NewResendSDK(map[string]any{
        "apikey": os.Getenv("RESEND_APIKEY"),
    })

    // Create a addContactToSegmentResponseSuccess.
    created, err := client.AddContactToSegmentResponseSuccess(nil).Create(map[string]any{"contact_id": "example_contact_id", "segment_id": "example_segment_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created.(sdk.Entity).Data())
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
automationrun, err := client.AutomationRun(nil).Load(map[string]any{"automation_id": "example", "id": "example_id"}, nil)
if err != nil {
    // handle err
    return
}
_ = automationrun
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

automationRun, err := client.AutomationRun(nil).Load(
    map[string]any{"id": "test01", "automation_id": "example"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(automationRun.(sdk.Entity).Data()) // the entity's mock record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewResendSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewResendSDK

```go
func NewResendSDK(options map[string]any) *ResendSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *ResendSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### ResendSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `AddContactToSegmentResponseSuccess` | `(data map[string]any) ResendEntity` | Create an AddContactToSegmentResponseSuccess entity instance. |
| `AddInboxThreadLabelsResponseSuccess` | `(data map[string]any) ResendEntity` | Create an AddInboxThreadLabelsResponseSuccess entity instance. |
| `ApiKey` | `(data map[string]any) ResendEntity` | Create an ApiKey entity instance. |
| `Audience` | `(data map[string]any) ResendEntity` | Create an Audience entity instance. |
| `Automation` | `(data map[string]any) ResendEntity` | Create an Automation entity instance. |
| `AutomationRun` | `(data map[string]any) ResendEntity` | Create an AutomationRun entity instance. |
| `AutomationRunListItem` | `(data map[string]any) ResendEntity` | Create an AutomationRunListItem entity instance. |
| `BatchAddSuppressionsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a BatchAddSuppressionsResponseSuccess entity instance. |
| `BatchRemoveSuppressionsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a BatchRemoveSuppressionsResponseSuccess entity instance. |
| `Broadcast` | `(data map[string]any) ResendEntity` | Create a Broadcast entity instance. |
| `Contact` | `(data map[string]any) ResendEntity` | Create a Contact entity instance. |
| `ContactImportResponseSuccess` | `(data map[string]any) ResendEntity` | Create a ContactImportResponseSuccess entity instance. |
| `ContactProperty` | `(data map[string]any) ResendEntity` | Create a ContactProperty entity instance. |
| `ContactTopicsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a ContactTopicsResponseSuccess entity instance. |
| `CreateApiKey` | `(data map[string]any) ResendEntity` | Create a CreateApiKey entity instance. |
| `CreateBatchEmail` | `(data map[string]any) ResendEntity` | Create a CreateBatchEmail entity instance. |
| `CreateBroadcastResponseSuccess` | `(data map[string]any) ResendEntity` | Create a CreateBroadcastResponseSuccess entity instance. |
| `CreateInboxDraftResponseSuccess` | `(data map[string]any) ResendEntity` | Create a CreateInboxDraftResponseSuccess entity instance. |
| `CreateInboxLabelResponseSuccess` | `(data map[string]any) ResendEntity` | Create a CreateInboxLabelResponseSuccess entity instance. |
| `Domain` | `(data map[string]any) ResendEntity` | Create a Domain entity instance. |
| `DomainClaim` | `(data map[string]any) ResendEntity` | Create a DomainClaim entity instance. |
| `Email` | `(data map[string]any) ResendEntity` | Create an Email entity instance. |
| `EmailsMetric` | `(data map[string]any) ResendEntity` | Create an EmailsMetric entity instance. |
| `Event` | `(data map[string]any) ResendEntity` | Create an Event entity instance. |
| `Inbox` | `(data map[string]any) ResendEntity` | Create an Inbox entity instance. |
| `InboxAgentResponseSuccess` | `(data map[string]any) ResendEntity` | Create an InboxAgentResponseSuccess entity instance. |
| `InboxDraftListItem` | `(data map[string]any) ResendEntity` | Create an InboxDraftListItem entity instance. |
| `InboxDraftResponseSuccess` | `(data map[string]any) ResendEntity` | Create an InboxDraftResponseSuccess entity instance. |
| `InboxLabel` | `(data map[string]any) ResendEntity` | Create an InboxLabel entity instance. |
| `InboxThread` | `(data map[string]any) ResendEntity` | Create an InboxThread entity instance. |
| `InboxThreadEmail` | `(data map[string]any) ResendEntity` | Create an InboxThreadEmail entity instance. |
| `InboxThreadEmailResponseSuccess` | `(data map[string]any) ResendEntity` | Create an InboxThreadEmailResponseSuccess entity instance. |
| `InboxThreadResponseSuccess` | `(data map[string]any) ResendEntity` | Create an InboxThreadResponseSuccess entity instance. |
| `ListAttachment` | `(data map[string]any) ResendEntity` | Create a ListAttachment entity instance. |
| `ListBroadcastClickedLinksResponseSuccess` | `(data map[string]any) ResendEntity` | Create a ListBroadcastClickedLinksResponseSuccess entity instance. |
| `ListBroadcastRecipientsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a ListBroadcastRecipientsResponseSuccess entity instance. |
| `ListContactSegmentsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a ListContactSegmentsResponseSuccess entity instance. |
| `ListContactsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a ListContactsResponseSuccess entity instance. |
| `ListWebhookEvent` | `(data map[string]any) ResendEntity` | Create a ListWebhookEvent entity instance. |
| `ListWebhookEventAttempt` | `(data map[string]any) ResendEntity` | Create a ListWebhookEventAttempt entity instance. |
| `Log` | `(data map[string]any) ResendEntity` | Create a Log entity instance. |
| `ReceivedEmail` | `(data map[string]any) ResendEntity` | Create a ReceivedEmail entity instance. |
| `RemoveAudienceResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveAudienceResponseSuccess entity instance. |
| `RemoveBroadcastResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveBroadcastResponseSuccess entity instance. |
| `RemoveContactFromSegmentResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveContactFromSegmentResponseSuccess entity instance. |
| `RemoveContactPropertyResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveContactPropertyResponseSuccess entity instance. |
| `RemoveContactResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveContactResponseSuccess entity instance. |
| `RemoveEvent` | `(data map[string]any) ResendEntity` | Create a RemoveEvent entity instance. |
| `RemoveInboxDraftResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveInboxDraftResponseSuccess entity instance. |
| `RemoveInboxLabelResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveInboxLabelResponseSuccess entity instance. |
| `RemoveInboxResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveInboxResponseSuccess entity instance. |
| `RemoveInboxThreadLabelResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveInboxThreadLabelResponseSuccess entity instance. |
| `RemoveInboxThreadResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveInboxThreadResponseSuccess entity instance. |
| `RemoveSegmentResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveSegmentResponseSuccess entity instance. |
| `RemoveSuppressionResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveSuppressionResponseSuccess entity instance. |
| `RemoveTemplateResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveTemplateResponseSuccess entity instance. |
| `RemoveTopicResponseSuccess` | `(data map[string]any) ResendEntity` | Create a RemoveTopicResponseSuccess entity instance. |
| `RetrievedAttachment` | `(data map[string]any) ResendEntity` | Create a RetrievedAttachment entity instance. |
| `RevokeOAuthGrant` | `(data map[string]any) ResendEntity` | Create a RevokeOAuthGrant entity instance. |
| `Rotate` | `(data map[string]any) ResendEntity` | Create a Rotate entity instance. |
| `Segment` | `(data map[string]any) ResendEntity` | Create a Segment entity instance. |
| `SetInboxThreadLabelsResponseSuccess` | `(data map[string]any) ResendEntity` | Create a SetInboxThreadLabelsResponseSuccess entity instance. |
| `Suppression` | `(data map[string]any) ResendEntity` | Create a Suppression entity instance. |
| `Template` | `(data map[string]any) ResendEntity` | Create a Template entity instance. |
| `Topic` | `(data map[string]any) ResendEntity` | Create a Topic entity instance. |
| `UpdateApiKey` | `(data map[string]any) ResendEntity` | Create an UpdateApiKey entity instance. |
| `UpdateBroadcastResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateBroadcastResponseSuccess entity instance. |
| `UpdateContactPropertyResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateContactPropertyResponseSuccess entity instance. |
| `UpdateContactResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateContactResponseSuccess entity instance. |
| `UpdateContactTopicsResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateContactTopicsResponseSuccess entity instance. |
| `UpdateDomainResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateDomainResponseSuccess entity instance. |
| `UpdateEmailOption` | `(data map[string]any) ResendEntity` | Create an UpdateEmailOption entity instance. |
| `UpdateEvent` | `(data map[string]any) ResendEntity` | Create an UpdateEvent entity instance. |
| `UpdateInboxDraftResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateInboxDraftResponseSuccess entity instance. |
| `UpdateInboxLabelResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateInboxLabelResponseSuccess entity instance. |
| `UpdateInboxResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateInboxResponseSuccess entity instance. |
| `UpdateInboxThreadResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateInboxThreadResponseSuccess entity instance. |
| `UpdateSegmentResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateSegmentResponseSuccess entity instance. |
| `UpdateTemplateResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateTemplateResponseSuccess entity instance. |
| `UpdateTopicResponseSuccess` | `(data map[string]any) ResendEntity` | Create an UpdateTopicResponseSuccess entity instance. |
| `UpdateWebhook` | `(data map[string]any) ResendEntity` | Create an UpdateWebhook entity instance. |
| `Usage` | `(data map[string]any) ResendEntity` | Create an Usage entity instance. |
| `Webhook` | `(data map[string]any) ResendEntity` | Create a Webhook entity instance. |
| `WebhookEvent` | `(data map[string]any) ResendEntity` | Create a WebhookEvent entity instance. |

### Entity interface (ResendEntity)

All entities implement the `ResendEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria, and return it. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria, one per record. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity, and return it. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity, and return it. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity, and return it marked as deleted. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the entity
itself — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity, whose `Data()` reads its record (`map[string]any`) |
| `List` | a `[]any` of entities, one per record |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    addContactToSegmentResponseSuccess, err := client.AddContactToSegmentResponseSuccess(nil).Create(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // addContactToSegmentResponseSuccess is the entity; addContactToSegmentResponseSuccess.(sdk.Entity).Data() reads its record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### AddContactToSegmentResponseSuccess

| Field | Description |
| --- | --- |
| `"contact_id"` | The ID of the contact. |
| `"object"` | The object type. |
| `"segment_id"` | The ID of the segment. |

Operations: Create.

API path: `/contacts/{contact_id}/segments/{segment_id}`

#### AddInboxThreadLabelsResponseSuccess

| Field | Description |
| --- | --- |
| `"data"` | Every label on the thread after the change. |
| `"labels"` | The IDs of the labels to apply to the thread. |
| `"object"` | Type of the response object. |

Operations: Create.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/labels`

#### ApiKey

| Field | Description |
| --- | --- |
| `"created_at"` | The date and time the API key was created. |
| `"id"` | The ID of the API key. |
| `"last_used_at"` | The date and time the API key was last used. |
| `"name"` | The name of the API key. |

Operations: List, Remove.

API path: `/api-keys`

#### Audience

| Field | Description |
| --- | --- |
| `"created_at"` | The date that the object was created. |
| `"id"` | The ID of the audience. |
| `"name"` | The name of the audience. |
| `"object"` | The object of the audience. |

Operations: Create, List, Load.

API path: `/audiences`

#### Automation

| Field | Description |
| --- | --- |
| `"connections"` | The connections between steps in the active version of the automation. |
| `"created_at"` | The date and time the automation was created. |
| `"id"` | The ID of the automation. |
| `"name"` | The name of the automation. |
| `"object"` | Type of the response object. |
| `"status"` | The current status of the automation. |
| `"steps"` | The steps in the active version of the automation. |
| `"updated_at"` | The date and time the automation was last updated. |

Operations: Create, List, Load, Remove, Update.

API path: `/automations/{automation_id}/duplicate`

#### AutomationRun

| Field | Description |
| --- | --- |
| `"completed_at"` | The date and time the run completed. |
| `"created_at"` | The date and time the run was created. |
| `"id"` | The ID of the automation run. |
| `"object"` | Type of the response object. |
| `"started_at"` | The date and time the run started. |
| `"status"` | The current status of the automation run. |
| `"steps"` | The steps executed in this run, sorted in graph order. |

Operations: Load.

API path: `/automations/{automation_id}/runs/{run_id}`

#### AutomationRunListItem

| Field | Description |
| --- | --- |
| `"completed_at"` | The date and time the run completed. |
| `"created_at"` | The date and time the run was created. |
| `"id"` | The ID of the automation run. |
| `"started_at"` | The date and time the run started. |
| `"status"` | The current status of the automation run. |

Operations: List.

API path: `/automations/{automation_id}/runs`

#### BatchAddSuppressionsResponseSuccess

| Field | Description |
| --- | --- |
| `"emails"` | Email addresses to suppress. |

Operations: Create.

API path: `/suppressions/batch/add`

#### BatchRemoveSuppressionsResponseSuccess

| Field | Description |
| --- | --- |
| `"data"` | Array containing the removed suppressions. |
| `"emails"` | Email addresses to remove from the suppression list. |
| `"ids"` | Suppression IDs to remove from the suppression list. |

Operations: Create.

API path: `/suppressions/batch/remove`

#### Broadcast

| Field | Description |
| --- | --- |
| `"audience_id"` | Deprecated: use `segment_id` instead. |
| `"created_at"` | Timestamp indicating when the broadcast was created. |
| `"from"` | The email address of the sender. |
| `"html"` | The HTML version of the broadcast content. |
| `"id"` | Unique identifier for the broadcast. |
| `"name"` | Name of the broadcast. |
| `"preview_text"` | The preview text of the email. |
| `"reply_to"` | The email addresses to which replies should be sent. |
| `"scheduled_at"` | Timestamp indicating when the broadcast is scheduled to be sent. |
| `"segment_id"` | Unique identifier of the segment this broadcast will be sent to. |
| `"sent_at"` | Timestamp indicating when the broadcast was sent. |
| `"status"` | The status of the broadcast. |
| `"subject"` | The subject line of the email. |
| `"text"` | The plain text version of the broadcast content. |
| `"topic_id"` | The topic ID that the broadcast is scoped to. |

Operations: Create, List, Load.

API path: `/broadcasts/{id}/cancel`

#### Contact

| Field | Description |
| --- | --- |
| `"audience_id"` | Unique identifier of the audience to which the contact belongs. |
| `"created_at"` | Timestamp indicating when the contact was created. |
| `"email"` | Email address of the contact. |
| `"first_name"` | First name of the contact. |
| `"id"` | Unique identifier for the contact. |
| `"last_name"` | Last name of the contact. |
| `"object"` | Type of the response object. |
| `"properties"` | A map of custom property keys and values. |
| `"segments"` | Array of segment IDs to add the contact to. |
| `"topics"` | Array of topic subscriptions for the contact. |
| `"unsubscribed"` | Indicates if the contact is unsubscribed. |

Operations: Create, List, Load.

API path: `/contacts`

#### ContactImportResponseSuccess

| Field | Description |
| --- | --- |
| `"column_map"` | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `"completed_at"` | Timestamp indicating when the contact import completed. |
| `"counts"` |  |
| `"created_at"` | Timestamp indicating when the contact import was created. |
| `"file"` | CSV file to import. |
| `"id"` | Unique identifier for the contact import. |
| `"object"` | Type of the response object. |
| `"on_conflict"` | Strategy to use when an imported contact already exists. |
| `"segments"` | JSON-encoded array of segments to add imported contacts to. |
| `"status"` | Current status of the contact import. |
| `"topics"` | JSON-encoded array of topic subscriptions to apply to imported contacts. |

Operations: Create, List, Load.

API path: `/contacts/imports`

#### ContactProperty

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the contact property was created. |
| `"fallback_value"` | The default value when the property is not set for a contact. |
| `"id"` | The ID of the contact property. |
| `"key"` | The property key. |
| `"object"` | The object type. |
| `"type"` | The property type. |

Operations: Create, List, Load.

API path: `/contact-properties`

#### ContactTopicsResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: List.

API path: `/contacts/{contact_id}/topics`

#### CreateApiKey

| Field | Description |
| --- | --- |
| `"domain_id"` | Restrict an API key to send emails only from a specific domain. |
| `"name"` | The API key name. |
| `"permission"` | The API key can have full access to Resend’s API or be only restricted to send emails. |

Operations: Create.

API path: `/api-keys`

#### CreateBatchEmail

| Field | Description |
| --- | --- |
| `"data"` |  |
| `"send_emails"` |  |

Operations: Create.

API path: `/emails/batch`

#### CreateBroadcastResponseSuccess

| Field | Description |
| --- | --- |
| `"audience_id"` | Use `segment_id` instead. |
| `"from"` | The email address of the sender. |
| `"html"` | The HTML version of the message. |
| `"name"` | Name of the broadcast. |
| `"preview_text"` | The preview text of the email. |
| `"reply_to"` | The email addresses to which replies should be sent. |
| `"scheduled_at"` | Schedule time to send the broadcast. |
| `"segment_id"` | Unique identifier of the segment this broadcast will be sent to. |
| `"send"` | Whether to send the broadcast immediately or keep it as a draft. |
| `"subject"` | The subject line of the email. |
| `"text"` | The plain text version of the message. |
| `"topic_id"` | The topic ID that the broadcast will be scoped to. |

Operations: Create.

API path: `/broadcasts`

#### CreateInboxDraftResponseSuccess

| Field | Description |
| --- | --- |
| `"attachments"` | Files to store on the draft. |
| `"bcc"` | The BCC recipients. |
| `"cc"` | The CC recipients. |
| `"html"` | The HTML body. |
| `"id"` | The ID of the draft. |
| `"object"` | Type of the response object. |
| `"reply_to_email_id"` | The Email ID to reply to, as returned in `data[].id` by the list thread emails endpoint. |
| `"subject"` | The subject. |
| `"text"` | The plain text body. |
| `"thread_id"` | The Thread ID to reply in. |
| `"to"` | The recipients. |

Operations: Create.

API path: `/inboxes/{inbox_id}/drafts`

#### CreateInboxLabelResponseSuccess

| Field | Description |
| --- | --- |
| `"color"` | The color of the label, as a `#RRGGBB` hex color. |
| `"name"` | The name of the label. |

Operations: Create.

API path: `/inboxes/{inbox_id}/labels`

#### Domain

| Field | Description |
| --- | --- |
| `"capabilities"` | Configure the domain capabilities for sending and receiving emails. |
| `"click_tracking"` | Whether click tracking is enabled for this domain. |
| `"created_at"` | The date and time the domain was created. |
| `"custom_return_path"` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `"id"` | The ID of the domain. |
| `"name"` | The name of the domain. |
| `"object"` | The type of object. |
| `"open_tracking"` | Whether open tracking is enabled for this domain. |
| `"records"` |  |
| `"region"` | The region where the domain is hosted. |
| `"status"` | The status of the domain. |
| `"tls"` | TLS mode. |
| `"tracking_subdomain"` | The subdomain used for click and open tracking. |

Operations: Create, List, Load, Remove.

API path: `/domains/{domain_id}/verify`

#### DomainClaim

| Field | Description |
| --- | --- |
| `"blocked_reason"` | Why the claim is currently blocked, if applicable. |
| `"click_tracking"` | Track clicks within the body of each HTML email. |
| `"created_at"` | The date and time the claim was created. |
| `"custom_return_path"` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `"domain_id"` | The ID of the placeholder domain created for the claim. |
| `"expires_at"` | The date and time the claim expires if not verified. |
| `"failure_reason"` | Why the claim failed, if applicable. |
| `"id"` | The ID of the claim. |
| `"name"` | The name of the domain being claimed. |
| `"object"` | The type of object. |
| `"open_tracking"` | Track the open rate of each email. |
| `"record"` | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `"region"` | The region where the claimed domain will send from. |
| `"status"` | The status of the claim. |
| `"tracking_subdomain"` | The subdomain to use for click and open tracking. |

Operations: Create, Load.

API path: `/domains/{domain_id}/claim/verify`

#### Email

| Field | Description |
| --- | --- |
| `"attachments"` |  |
| `"bcc"` | The email addresses of the blind carbon copy recipients. |
| `"cc"` | The email addresses of the carbon copy recipients. |
| `"created_at"` | The date and time the email was created. |
| `"from"` | The email address of the sender. |
| `"headers"` | Custom headers to add to the email. |
| `"html"` | The HTML body of the email. |
| `"id"` | The ID of the email. |
| `"last_event"` | The status of the email. |
| `"message_id"` | The Message-ID header value of the email. |
| `"object"` | The type of object. |
| `"reply_to"` | The email addresses to which replies should be sent. |
| `"scheduled_at"` | Schedule email to be sent later. |
| `"subject"` | The subject line of the email. |
| `"tags"` |  |
| `"template"` |  |
| `"text"` | The plain text body of the email. |
| `"to"` | Recipient email address. |
| `"topic_id"` | The topic ID to scope the email to. |

Operations: Create, List, Load.

API path: `/emails/{email_id}/cancel`

#### EmailsMetric

| Field | Description |
| --- | --- |
| `"broadcast_id"` | Present when `broadcast` is in `dimensions`. |
| `"broadcast_name"` | Present when `broadcast` is in `dimensions`. |
| `"domain_id"` | Present when `domain` is in `dimensions`. |
| `"domain_name"` | Present when `domain` is in `dimensions`. |
| `"email_id"` | Present when `email` is in `dimensions`. |
| `"period"` | Present when `period` is in `dimensions`. |

Operations: List.

API path: `/emails/metrics`

#### Event

| Field | Description |
| --- | --- |
| `"created_at"` | The date and time the event was created. |
| `"id"` | The event ID. |
| `"name"` | The event name. |
| `"object"` | Type of the response object. |
| `"schema"` | A flat key/type map defining the event payload schema. |
| `"updated_at"` | The date and time the event was last updated. |

Operations: Create, List, Load.

API path: `/events`

#### Inbox

| Field | Description |
| --- | --- |
| `"created_at"` | ISO 8601 timestamp when the inbox was created. |
| `"domain_id"` | The ID of the domain that owns the address. |
| `"drafts"` | The number of unsent drafts. |
| `"email_address"` | The address of the inbox. |
| `"forwarding"` | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `"from_name"` | The name recipients see when mail is sent from this inbox. |
| `"id"` | The ID of the inbox. |
| `"last_received"` | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `"name"` | Internal name for the inbox. |
| `"object"` | The object type. |
| `"receiving_address"` | The address to forward mail to when forwarding is enabled. |
| `"unread"` | The number of unread threads in the Inbox folder. |
| `"updated_at"` | ISO 8601 timestamp when the inbox was last updated. |

Operations: Create, List, Load, Update.

API path: `/inboxes`

#### InboxAgentResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: List.

API path: `/inboxes/{inbox_id}/agent`

#### InboxDraftListItem

| Field | Description |
| --- | --- |
| `"attachments"` | The files stored on the draft. |
| `"bcc"` | The BCC recipients. |
| `"cc"` | The CC recipients. |
| `"id"` | The ID of the draft. |
| `"reply_to_email_id"` | The Email ID the draft replies to. |
| `"snippet"` | A short preview of the body. |
| `"subject"` | The subject. |
| `"thread_id"` | The Thread ID when the draft is a reply. |
| `"to"` | The recipients. |
| `"type"` | `reply` when the draft replies to an email in an existing thread. |
| `"updated_at"` | Timestamp when the draft was last saved. |

Operations: List.

API path: `/inboxes/{inbox_id}/drafts`

#### InboxDraftResponseSuccess

| Field | Description |
| --- | --- |
| `"attachments"` | The files stored on the draft. |
| `"bcc"` | The BCC recipients. |
| `"cc"` | The CC recipients. |
| `"created_at"` | Timestamp when the draft was created. |
| `"email_id"` | The ID of the outbound email queued when the draft was sent. |
| `"html"` | The HTML body. |
| `"id"` | The ID of the draft. |
| `"object"` | The object type. |
| `"reply_to_email_id"` | The Email ID the draft replies to. |
| `"subject"` | The subject. |
| `"text"` | The plain text body. |
| `"thread_id"` | The Thread ID when the draft is a reply. |
| `"to"` | The recipients. |
| `"type"` | `reply` when the draft replies to an email in an existing thread. |
| `"updated_at"` | Timestamp when the draft was last saved. |

Operations: Create, Load.

API path: `/inboxes/{inbox_id}/drafts/{draft_id}/send`

#### InboxLabel

| Field | Description |
| --- | --- |
| `"color"` | The color of the label, as a `#RRGGBB` hex color in uppercase. |
| `"created_at"` | Timestamp when the label was created. |
| `"id"` | The ID of the label. |
| `"name"` | The name of the label. |

Operations: List.

API path: `/inboxes/{inbox_id}/labels`

#### InboxThread

| Field | Description |
| --- | --- |
| `"bcc"` | The BCC recipients of the most recent email. |
| `"cc"` | The CC recipients of the most recent email. |
| `"email_count"` | The number of emails in the thread. |
| `"folder"` | The folder the thread is in. |
| `"from"` | The sender email address. |
| `"has_attachment"` | Whether the most recent email in the thread has an attachment. |
| `"has_draft"` | Whether the thread has an unsent draft. |
| `"id"` | The ID of the thread. |
| `"labels"` | The labels attached to the thread. |
| `"matched_email_id"` | The ID of the email in the thread that matched the search. |
| `"read"` | True only when every email in the thread is read. |
| `"received_at"` | Timestamp when the thread was last active. |
| `"subject"` | The subject of the thread. |
| `"to"` | The recipients of the most recent email. |

Operations: List.

API path: `/inboxes/{inbox_id}/threads`

#### InboxThreadEmail

| Field | Description |
| --- | --- |
| `"attachments"` | The attachments of the message. |
| `"bcc"` | The BCC recipients of the message. |
| `"cc"` | The CC recipients of the message. |
| `"direction"` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `"from"` | Sender email address. |
| `"html"` | The HTML body. |
| `"id"` | The ID of the message. |
| `"message_id"` | The Message-ID header of the message. |
| `"read"` | Whether the message has been read. |
| `"received_at"` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `"reply_to"` | The Reply-To addresses of the message. |
| `"subject"` | The subject of the message. |
| `"text"` | The plain-text body. |
| `"to"` | The recipients of the message. |

Operations: List.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/emails`

#### InboxThreadEmailResponseSuccess

| Field | Description |
| --- | --- |
| `"attachments"` | The attachments of the message. |
| `"bcc"` | The BCC recipients of the message. |
| `"cc"` | The CC recipients of the message. |
| `"direction"` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `"from"` | Sender email address. |
| `"html"` | The HTML body. |
| `"id"` | The ID of the message. |
| `"message_id"` | The Message-ID header of the message. |
| `"read"` | Whether the message has been read. |
| `"received_at"` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `"reply_to"` | The Reply-To addresses of the message. |
| `"subject"` | The subject of the message. |
| `"text"` | The plain-text body. |
| `"to"` | The recipients of the message. |

Operations: Create, Load.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/forward`

#### InboxThreadResponseSuccess

| Field | Description |
| --- | --- |
| `"folder"` | The folder the thread is in. |
| `"id"` | The ID of the thread. |
| `"labels"` | The labels attached to the thread. |
| `"object"` | The object type. |
| `"read"` | True only when every email in the thread is read. |
| `"subject"` | The subject of the thread. |

Operations: Load.

API path: `/inboxes/{inbox_id}/threads/{thread_id}`

#### ListAttachment

| Field | Description |
| --- | --- |
| `"content_disposition"` | How the attachment should be displayed. |
| `"content_id"` | The content ID for inline attachments. |
| `"content_type"` | The MIME type of the attachment. |
| `"download_url"` | Signed URL to download the attachment content. |
| `"expires_at"` | Timestamp when the download URL expires. |
| `"filename"` | The filename of the attachment. |
| `"id"` | The ID of the attachment. |
| `"size"` | Size of the attachment in bytes. |

Operations: List.

API path: `/emails/{email_id}/attachments`

#### ListBroadcastClickedLinksResponseSuccess

| Field | Description |
| --- | --- |
| `"clicks"` | Total number of clicks on this URL. |
| `"id"` | An opaque cursor for this row, used only for pagination. |
| `"unique_clicks"` | Number of unique clicks on this URL. |
| `"url"` | The URL that was clicked. |

Operations: List.

API path: `/broadcasts/{id}/clicked-links`

#### ListBroadcastRecipientsResponseSuccess

| Field | Description |
| --- | --- |
| `"bounce_type"` | The type of bounce. |
| `"clicked_links"` | The links this recipient clicked. |
| `"contact_id"` | The ID of the contact associated with this recipient, if one exists. |
| `"count"` | The number of times this recipient triggered the event. |
| `"email"` | The recipient's email address. |
| `"id"` | Opaque cursor identifying this row, used for pagination. |

Operations: List.

API path: `/broadcasts/{id}/recipients`

#### ListContactSegmentsResponseSuccess

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the contact was added to the segment. |
| `"id"` | Unique identifier for the segment. |
| `"name"` | Name of the segment. |

Operations: List.

API path: `/contacts/{contact_id}/segments`

#### ListContactsResponseSuccess

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the contact was created. |
| `"email"` | Email address of the contact. |
| `"first_name"` | First name of the contact. |
| `"id"` | Unique identifier for the contact. |
| `"last_name"` | Last name of the contact. |
| `"unsubscribed"` | Indicates if the contact is unsubscribed. |

Operations: List.

API path: `/segments/{id}/contacts`

#### ListWebhookEvent

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the event was created. |
| `"id"` | The ID of the webhook event. |
| `"status"` | The delivery status of the event for this webhook. |
| `"type"` | The type of the event. |

Operations: List.

API path: `/webhooks/{webhook_id}/events`

#### ListWebhookEventAttempt

| Field | Description |
| --- | --- |
| `"http_status_code"` | The HTTP status code returned by the webhook endpoint. |
| `"id"` | The ID of the webhook event attempt. |
| `"response"` | The response body returned by the webhook endpoint. |
| `"sent_at"` | Timestamp indicating when the attempt was sent. |

Operations: List.

API path: `/webhooks/{webhook_id}/events/{event_id}/attempts`

#### Log

| Field | Description |
| --- | --- |
| `"created_at"` | The date the log was created. |
| `"endpoint"` | The API endpoint that was called. |
| `"id"` | The log ID. |
| `"method"` | The HTTP method used. |
| `"object"` | Type of the response object. |
| `"request_body"` | The request body sent to the API. |
| `"response_body"` | The response body returned by the API. |
| `"response_status"` | The HTTP status code of the response. |
| `"user_agent"` | The user agent of the request. |

Operations: List, Load.

API path: `/logs`

#### ReceivedEmail

| Field | Description |
| --- | --- |
| `"attachments"` | Array of attachments. |
| `"bcc"` | The BCC recipients. |
| `"cc"` | The CC recipients. |
| `"created_at"` | Timestamp when the email was received. |
| `"from"` | The sender email address. |
| `"headers"` | The email headers. |
| `"html"` | The HTML content of the email. |
| `"id"` | The ID of the received email. |
| `"message_id"` | The unique message ID from the email headers. |
| `"object"` | The type of object. |
| `"received_for"` | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `"reply_to"` | The reply-to addresses. |
| `"subject"` | The email subject. |
| `"text"` | The plain text content of the email. |
| `"to"` | The recipient email addresses. |

Operations: List, Load.

API path: `/emails/receiving`

#### RemoveAudienceResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: Remove.

API path: `/audiences/{id}`

#### RemoveBroadcastResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: Remove.

API path: `/broadcasts/{id}`

#### RemoveContactFromSegmentResponseSuccess

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/contacts/{contact_id}/segments/{segment_id}`

#### RemoveContactPropertyResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: Remove.

API path: `/contact-properties/{id}`

#### RemoveContactResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: Remove.

API path: `/contacts/{id}`

#### RemoveEvent

| Field | Description |
| --- | --- |
| `"id"` |  |

Operations: Remove.

API path: `/events/{identifier}`

#### RemoveInboxDraftResponseSuccess

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/inboxes/{inbox_id}/drafts/{draft_id}`

#### RemoveInboxLabelResponseSuccess

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/inboxes/{inbox_id}/labels/{label_id}`

#### RemoveInboxResponseSuccess

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/inboxes/{inbox_id}`

#### RemoveInboxThreadLabelResponseSuccess

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/labels/{label_id}`

#### RemoveInboxThreadResponseSuccess

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/inboxes/{inbox_id}/threads/{thread_id}`

#### RemoveSegmentResponseSuccess

| Field | Description |
| --- | --- |
| `"audience_id"` | The ID of the audience this segment belongs to. |
| `"created_at"` | Timestamp indicating when the segment was created. |
| `"filter"` | Filter conditions for the segment. |
| `"id"` | Unique identifier for the segment. |
| `"name"` | The name of the segment. |

Operations: Create, List, Remove.

API path: `/segments`

#### RemoveSuppressionResponseSuccess

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the suppression was created. |
| `"email"` | Email address to suppress. |
| `"id"` | Unique identifier for the suppression. |
| `"origin"` | Origin of the suppression. |
| `"source_id"` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

Operations: Create, List, Remove.

API path: `/suppressions`

#### RemoveTemplateResponseSuccess

| Field | Description |
| --- | --- |
| `"alias"` | The alias of the template. |
| `"created_at"` | Timestamp indicating when the template was created. |
| `"from"` | Sender email address. |
| `"html"` | The HTML version of the template. |
| `"id"` | The ID of the template. |
| `"name"` | The name of the template. |
| `"published_at"` | Timestamp indicating when the template was published. |
| `"reply_to"` | Reply-to email addresses. |
| `"status"` | The publication status of the template. |
| `"subject"` | Email subject. |
| `"text"` | The plain text version of the template. |
| `"updated_at"` | Timestamp indicating when the template was last updated. |
| `"variables"` |  |

Operations: Create, List, Remove.

API path: `/templates`

#### RemoveTopicResponseSuccess

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the topic was created. |
| `"default_subscription"` | The default subscription status for the topic. |
| `"description"` | A description of the topic. |
| `"id"` | Unique identifier for the topic. |
| `"name"` | The name of the topic. |
| `"visibility"` | The visibility of the topic. |

Operations: Create, List, Remove.

API path: `/topics`

#### RetrievedAttachment

| Field | Description |
| --- | --- |
| `"content_disposition"` | How the attachment should be displayed. |
| `"content_id"` | The content ID for inline attachments. |
| `"content_type"` | The MIME type of the attachment. |
| `"download_url"` | Signed URL to download the attachment content. |
| `"expires_at"` | Timestamp when the download URL expires. |
| `"filename"` | The filename of the attachment. |
| `"id"` | The ID of the attachment. |
| `"object"` | The type of object. |
| `"size"` | Size of the attachment in bytes. |

Operations: Load.

API path: `/emails/{email_id}/attachments/{attachment_id}`

#### RevokeOAuthGrant

| Field | Description |
| --- | --- |
| `"client"` | The OAuth client the grant was issued to. |
| `"client_id"` | The ID of the OAuth client the grant was issued to. |
| `"created_at"` | The date and time the OAuth grant was created. |
| `"id"` | The ID of the OAuth grant. |
| `"revoked_at"` | The date and time the OAuth grant was revoked, or null if it is still active. |
| `"revoked_reason"` | The reason the OAuth grant was revoked, or null if it is still active. |
| `"scopes"` | The scopes granted to the OAuth client. |

Operations: List, Remove.

API path: `/oauth/grants`

#### Rotate

| Field | Description |
| --- | --- |
| `"id"` | The ID of the webhook. |
| `"object"` | The type of object. |
| `"signing_secret"` | The new secret key used to verify webhook payloads. |

Operations: Create.

API path: `/webhooks/{webhook_id}/signing-secret/rotate`

#### Segment

| Field | Description |
| --- | --- |
| `"audience_id"` | The ID of the audience this segment belongs to. |
| `"created_at"` | Timestamp indicating when the segment was created. |
| `"filter"` | Filter conditions for the segment. |
| `"id"` | The ID of the segment. |
| `"name"` | The name of the segment. |
| `"object"` | The object type. |

Operations: Load.

API path: `/segments/{id}`

#### SetInboxThreadLabelsResponseSuccess

| Field | Description |
| --- | --- |
| `"data"` | Every label on the thread after the change. |
| `"labels"` | The IDs of the labels the thread carries after the change. |
| `"object"` | Type of the response object. |

Operations: Update.

API path: `/inboxes/{inbox_id}/threads/{thread_id}/labels`

#### Suppression

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the suppression was created. |
| `"email"` | Email address that is suppressed. |
| `"id"` | Unique identifier for the suppression. |
| `"object"` | Type of the response object. |
| `"origin"` | Origin of the suppression. |
| `"source_id"` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

Operations: Load.

API path: `/suppressions/{suppression}`

#### Template

| Field | Description |
| --- | --- |
| `"alias"` | The alias of the template. |
| `"created_at"` | Timestamp indicating when the template was created. |
| `"current_version_id"` | The ID of the current version of the template. |
| `"from"` | Sender email address. |
| `"has_unpublished_versions"` | Indicates whether the template has unpublished versions. |
| `"html"` | The HTML version of the template. |
| `"id"` | The ID of the template. |
| `"name"` | The name of the template. |
| `"object"` | The type of object. |
| `"published_at"` | Timestamp indicating when the template was published. |
| `"reply_to"` | Reply-to email addresses. |
| `"status"` | The publication status of the template. |
| `"subject"` | Email subject. |
| `"text"` | The plain text version of the template. |
| `"updated_at"` | Timestamp indicating when the template was last updated. |
| `"variables"` |  |

Operations: Create, Load.

API path: `/templates/{id}/duplicate`

#### Topic

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the topic was created. |
| `"default_subscription"` | The default subscription status for the topic. |
| `"description"` | A description of the topic. |
| `"id"` | The ID of the topic. |
| `"name"` | The name of the topic. |
| `"object"` | The object type. |
| `"visibility"` | The visibility of the topic. |

Operations: Load.

API path: `/topics/{id}`

#### UpdateApiKey

| Field | Description |
| --- | --- |
| `"id"` | The ID of the API key. |
| `"name"` | The API key name. |
| `"object"` | The type of object. |

Operations: Update.

API path: `/api-keys/{api_key_id}`

#### UpdateBroadcastResponseSuccess

| Field | Description |
| --- | --- |
| `"audience_id"` | Use `segment_id` instead. |
| `"from"` | The email address of the sender. |
| `"html"` | The HTML version of the message. |
| `"id"` | The ID of the broadcast. |
| `"name"` | Name of the broadcast. |
| `"object"` | The object type of the response. |
| `"preview_text"` | The preview text of the email. |
| `"reply_to"` | The email addresses to which replies should be sent. |
| `"segment_id"` | Unique identifier of the segment this broadcast will be sent to. |
| `"subject"` | The subject line of the email. |
| `"text"` | The plain text version of the message. |
| `"topic_id"` | The topic ID that the broadcast will be scoped to. |

Operations: Update.

API path: `/broadcasts/{id}`

#### UpdateContactPropertyResponseSuccess

| Field | Description |
| --- | --- |
| `"fallback_value"` | The default value to use when the property is not set for a contact. |
| `"id"` | The ID of the contact property. |
| `"object"` | The object type. |

Operations: Update.

API path: `/contact-properties/{id}`

#### UpdateContactResponseSuccess

| Field | Description |
| --- | --- |
| `"email"` | Email address of the contact. |
| `"first_name"` | First name of the contact. |
| `"id"` | Unique identifier for the updated contact. |
| `"last_name"` | Last name of the contact. |
| `"object"` | Type of the response object. |
| `"properties"` | A map of custom property keys and values to update. |
| `"unsubscribed"` | The Contact's global subscription status. |

Operations: Update.

API path: `/contacts/{id}`

#### UpdateContactTopicsResponseSuccess

| Field | Description |
| --- | --- |
| `"contact_id"` | The ID of the contact. |
| `"object"` | The object type. |
| `"topics"` | Array of updated topic subscriptions. |

Operations: Update.

API path: `/contacts/{contact_id}/topics`

#### UpdateDomainResponseSuccess

| Field | Description |
| --- | --- |
| `"capabilities"` | Configure the domain capabilities for sending and receiving emails. |
| `"click_tracking"` | Track clicks within the body of each HTML email. |
| `"id"` | The ID of the updated domain. |
| `"object"` | The object type representing the updated domain. |
| `"open_tracking"` | Track the open rate of each email. |
| `"tls"` | enforced | opportunistic. |
| `"tracking_subdomain"` | The subdomain to use for click and open tracking. |

Operations: Update.

API path: `/domains/{domain_id}`

#### UpdateEmailOption

| Field | Description |
| --- | --- |
| `"scheduled_at"` | Schedule email to be sent later. |

Operations: Update.

API path: `/emails/{email_id}`

#### UpdateEvent

| Field | Description |
| --- | --- |
| `"id"` | The ID of the updated event. |
| `"object"` | Type of the response object. |
| `"schema"` | A flat key/type map defining the event payload schema. |

Operations: Update.

API path: `/events/{identifier}`

#### UpdateInboxDraftResponseSuccess

| Field | Description |
| --- | --- |
| `"attachments"` | Files to store on the draft. |
| `"bcc"` | The BCC recipients. |
| `"cc"` | The CC recipients. |
| `"html"` | The HTML body. |
| `"id"` | The ID of the draft. |
| `"object"` | The object type. |
| `"subject"` | The subject. |
| `"text"` | The plain text body. |
| `"to"` | The recipients. |

Operations: Update.

API path: `/inboxes/{inbox_id}/drafts/{draft_id}`

#### UpdateInboxLabelResponseSuccess

| Field | Description |
| --- | --- |
| `"color"` | A new color for the label, as a `#RRGGBB` hex color. |
| `"id"` | The ID of the updated label. |
| `"name"` | A new name for the label. |
| `"object"` | Type of the response object. |

Operations: Update.

API path: `/inboxes/{inbox_id}/labels/{label_id}`

#### UpdateInboxResponseSuccess

| Field | Description |
| --- | --- |
| `"from_name"` | The name recipients see when mail is sent from this inbox. |
| `"id"` | The ID of the inbox. |
| `"name"` | New internal name for the inbox. |
| `"object"` | The object type. |

Operations: Update.

API path: `/inboxes/{inbox_id}`

#### UpdateInboxThreadResponseSuccess

| Field | Description |
| --- | --- |
| `"folder"` | The folder the thread is in. |
| `"id"` | The ID of the thread. |
| `"labels"` | The labels attached to the thread. |
| `"object"` | The object type. |
| `"read"` | True only when every email in the thread is read. |
| `"subject"` | The subject of the thread. |

Operations: Update.

API path: `/inboxes/{inbox_id}/threads/{thread_id}`

#### UpdateSegmentResponseSuccess

| Field | Description |
| --- | --- |
| `"id"` | The ID of the segment. |
| `"name"` | The name of the segment. |
| `"object"` | The object type. |

Operations: Update.

API path: `/segments/{id}`

#### UpdateTemplateResponseSuccess

| Field | Description |
| --- | --- |
| `"alias"` | The alias of the template. |
| `"from"` | Sender email address. |
| `"html"` | The HTML version of the template. |
| `"id"` | The ID of the template. |
| `"name"` | The name of the template. |
| `"object"` | The object type of the response. |
| `"reply_to"` | Reply-to email addresses. |
| `"subject"` | Email subject. |
| `"text"` | The plain text version of the template. |
| `"variables"` |  |

Operations: Update.

API path: `/templates/{id}`

#### UpdateTopicResponseSuccess

| Field | Description |
| --- | --- |
| `"description"` | A description of the topic. |
| `"id"` | The ID of the topic. |
| `"name"` | The name of the topic. |
| `"object"` | The object type. |
| `"visibility"` | The visibility of the topic. |

Operations: Update.

API path: `/topics/{id}`

#### UpdateWebhook

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the webhook was created. |
| `"endpoint"` | The URL where webhook events will be sent. |
| `"events"` | Array of event types to subscribe to. |
| `"id"` | The ID of the updated webhook. |
| `"object"` | The type of object. |
| `"status"` | The status of the webhook. |

Operations: Create, List, Update.

API path: `/webhooks`

#### Usage

| Field | Description |
| --- | --- |
| `"ai_credits"` |  |
| `"automation_runs"` |  |
| `"broadcasts"` |  |
| `"contacts"` |  |
| `"domains"` |  |
| `"emails"` |  |
| `"object"` | The type of object. |
| `"rate_limit"` |  |
| `"segments"` |  |

Operations: Load.

API path: `/usage`

#### Webhook

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the webhook was created. |
| `"endpoint"` | The URL where webhook events are sent. |
| `"events"` | Array of event types subscribed to. |
| `"id"` | The ID of the webhook. |
| `"object"` | The type of object. |
| `"signing_secret"` | The secret key used to verify webhook payloads. |
| `"status"` | The status of the webhook. |

Operations: Load, Remove.

API path: `/webhooks/{webhook_id}`

#### WebhookEvent

| Field | Description |
| --- | --- |
| `"created_at"` | Timestamp indicating when the event was created. |
| `"id"` | The ID of the webhook event. |
| `"next_attempt_at"` | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `"object"` | The type of object. |
| `"payload"` | The event payload sent to the webhook endpoint. |
| `"status"` | The delivery status of the event for this webhook. |
| `"type"` | The type of the event. |

Operations: Create, Load.

API path: `/webhooks/{webhook_id}/events/{event_id}/replay`



## Entities


### AddContactToSegmentResponseSuccess

Create an instance: `addContactToSegmentResponseSuccess := client.AddContactToSegmentResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` | The ID of the contact. |
| `object` | `string` | The object type. |
| `segment_id` | `string` | The ID of the segment. |

#### Example: Create

```go
result, err := client.AddContactToSegmentResponseSuccess(nil).Create(map[string]any{
    "contact_id": "example_contact_id",
    "segment_id": "example_segment_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### AddInboxThreadLabelsResponseSuccess

Create an instance: `addInboxThreadLabelsResponseSuccess := client.AddInboxThreadLabelsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `[]any` | Every label on the thread after the change. |
| `labels` | `[]any` | The IDs of the labels to apply to the thread. |
| `object` | `string` | Type of the response object. |

#### Example: Create

```go
result, err := client.AddInboxThreadLabelsResponseSuccess(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
    "thread_id": "example_thread_id",
    "labels": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### ApiKey

Create an instance: `apiKey := client.ApiKey(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date and time the API key was created. |
| `id` | `string` | The ID of the API key. |
| `last_used_at` | `any` | The date and time the API key was last used. |
| `name` | `string` | The name of the API key. |

#### Example: List

```go
apiKeys, err := client.ApiKey(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range apiKeys.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Audience

Create an instance: `audience := client.Audience(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date that the object was created. |
| `id` | `string` | The ID of the audience. |
| `name` | `string` | The name of the audience. |
| `object` | `string` | The object of the audience. |

#### Example: Load

```go
audience, err := client.Audience(nil).Load(map[string]any{"id": "audience_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(audience.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
audiences, err := client.Audience(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range audiences.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Audience(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Automation

Create an instance: `automation := client.Automation(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `connections` | `[]any` | The connections between steps in the active version of the automation. |
| `created_at` | `string` | The date and time the automation was created. |
| `id` | `string` | The ID of the automation. |
| `name` | `string` | The name of the automation. |
| `object` | `string` | Type of the response object. |
| `status` | `string` | The current status of the automation. |
| `steps` | `[]any` | The steps in the active version of the automation. |
| `updated_at` | `string` | The date and time the automation was last updated. |

#### Example: Load

```go
automation, err := client.Automation(nil).Load(map[string]any{"id": "automation_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(automation.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
automations, err := client.Automation(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range automations.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Automation(nil).Create(map[string]any{
    "connections": []any{},
    "name": "example_name",
    "steps": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### AutomationRun

Create an instance: `automationRun := client.AutomationRun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `completed_at` | `any` | The date and time the run completed. |
| `created_at` | `string` | The date and time the run was created. |
| `id` | `string` | The ID of the automation run. |
| `object` | `string` | Type of the response object. |
| `started_at` | `any` | The date and time the run started. |
| `status` | `string` | The current status of the automation run. |
| `steps` | `[]any` | The steps executed in this run, sorted in graph order. |

#### Example: Load

```go
automationRun, err := client.AutomationRun(nil).Load(map[string]any{"id": "automation_run_id", "automation_id": "automation_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(automationRun.(sdk.Entity).Data()) // the loaded entity's record
```


### AutomationRunListItem

Create an instance: `automationRunListItem := client.AutomationRunListItem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `completed_at` | `any` | The date and time the run completed. |
| `created_at` | `string` | The date and time the run was created. |
| `id` | `string` | The ID of the automation run. |
| `started_at` | `any` | The date and time the run started. |
| `status` | `string` | The current status of the automation run. |

#### Example: List

```go
automationRunListItems, err := client.AutomationRunListItem(nil).List(map[string]any{"id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range automationRunListItems.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### BatchAddSuppressionsResponseSuccess

Create an instance: `batchAddSuppressionsResponseSuccess := client.BatchAddSuppressionsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `emails` | `[]any` | Email addresses to suppress. |

#### Example: Create

```go
result, err := client.BatchAddSuppressionsResponseSuccess(nil).Create(map[string]any{
    "emails": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### BatchRemoveSuppressionsResponseSuccess

Create an instance: `batchRemoveSuppressionsResponseSuccess := client.BatchRemoveSuppressionsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `[]any` | Array containing the removed suppressions. |
| `emails` | `[]any` | Email addresses to remove from the suppression list. |
| `ids` | `[]any` | Suppression IDs to remove from the suppression list. |

#### Example: Create

```go
result, err := client.BatchRemoveSuppressionsResponseSuccess(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Broadcast

Create an instance: `broadcast := client.Broadcast(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `any` | Deprecated: use `segment_id` instead. |
| `created_at` | `string` | Timestamp indicating when the broadcast was created. |
| `from` | `string` | The email address of the sender. |
| `html` | `any` | The HTML version of the broadcast content. |
| `id` | `string` | Unique identifier for the broadcast. |
| `name` | `string` | Name of the broadcast. |
| `preview_text` | `string` | The preview text of the email. |
| `reply_to` | `[]any` | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `any` | Unique identifier of the segment this broadcast will be sent to. |
| `sent_at` | `string` | Timestamp indicating when the broadcast was sent. |
| `status` | `string` | The status of the broadcast. |
| `subject` | `string` | The subject line of the email. |
| `text` | `any` | The plain text version of the broadcast content. |
| `topic_id` | `any` | The topic ID that the broadcast is scoped to. |

#### Example: Load

```go
broadcast, err := client.Broadcast(nil).Load(map[string]any{"id": "broadcast_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(broadcast.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
broadcasts, err := client.Broadcast(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range broadcasts.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Contact

Create an instance: `contact := client.Contact(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | Unique identifier of the audience to which the contact belongs. |
| `created_at` | `string` | Timestamp indicating when the contact was created. |
| `email` | `string` | Email address of the contact. |
| `first_name` | `any` | First name of the contact. |
| `id` | `string` | Unique identifier for the contact. |
| `last_name` | `any` | Last name of the contact. |
| `object` | `string` | Type of the response object. |
| `properties` | `map[string]any` | A map of custom property keys and values. |
| `segments` | `[]any` | Array of segment IDs to add the contact to. |
| `topics` | `[]any` | Array of topic subscriptions for the contact. |
| `unsubscribed` | `bool` | Indicates if the contact is unsubscribed. |

#### Example: Load

```go
contact, err := client.Contact(nil).Load(map[string]any{"id": "contact_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(contact.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
contacts, err := client.Contact(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range contacts.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Contact(nil).Create(map[string]any{
    "email": "example_email",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### ContactImportResponseSuccess

Create an instance: `contactImportResponseSuccess := client.ContactImportResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `column_map` | `string` | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | `any` | Timestamp indicating when the contact import completed. |
| `counts` | `map[string]any` |  |
| `created_at` | `string` | Timestamp indicating when the contact import was created. |
| `file` | `string` | CSV file to import. |
| `id` | `string` | Unique identifier for the contact import. |
| `object` | `string` | Type of the response object. |
| `on_conflict` | `string` | Strategy to use when an imported contact already exists. |
| `segments` | `string` | JSON-encoded array of segments to add imported contacts to. |
| `status` | `string` | Current status of the contact import. |
| `topics` | `string` | JSON-encoded array of topic subscriptions to apply to imported contacts. |

#### Example: Load

```go
contactImportResponseSuccess, err := client.ContactImportResponseSuccess(nil).Load(map[string]any{"id": "contact_import_response_success_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(contactImportResponseSuccess.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
contactImportResponseSuccesss, err := client.ContactImportResponseSuccess(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range contactImportResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.ContactImportResponseSuccess(nil).Create(map[string]any{
    "file": "example_file",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### ContactProperty

Create an instance: `contactProperty := client.ContactProperty(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

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

```go
contactProperty, err := client.ContactProperty(nil).Load(map[string]any{"id": "contact_property_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(contactProperty.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
contactPropertys, err := client.ContactProperty(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range contactPropertys.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.ContactProperty(nil).Create(map[string]any{
    "key": "example_key",
    "type": "example_type",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### ContactTopicsResponseSuccess

Create an instance: `contactTopicsResponseSuccess := client.ContactTopicsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |

#### Example: List

```go
contactTopicsResponseSuccesss, err := client.ContactTopicsResponseSuccess(nil).List(map[string]any{"id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range contactTopicsResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### CreateApiKey

Create an instance: `createApiKey := client.CreateApiKey(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `domain_id` | `string` | Restrict an API key to send emails only from a specific domain. |
| `name` | `string` | The API key name. |
| `permission` | `string` | The API key can have full access to Resend’s API or be only restricted to send emails. |

#### Example: Create

```go
result, err := client.CreateApiKey(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### CreateBatchEmail

Create an instance: `createBatchEmail := client.CreateBatchEmail(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `[]any` |  |
| `send_emails` | `[]any` |  |

#### Example: Create

```go
result, err := client.CreateBatchEmail(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### CreateBroadcastResponseSuccess

Create an instance: `createBroadcastResponseSuccess := client.CreateBroadcastResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | Use `segment_id` instead. |
| `from` | `string` | The email address of the sender. |
| `html` | `string` | The HTML version of the message. |
| `name` | `string` | Name of the broadcast. |
| `preview_text` | `string` | The preview text of the email. |
| `reply_to` | `[]any` | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | Schedule time to send the broadcast. |
| `segment_id` | `string` | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `bool` | Whether to send the broadcast immediately or keep it as a draft. |
| `subject` | `string` | The subject line of the email. |
| `text` | `string` | The plain text version of the message. |
| `topic_id` | `string` | The topic ID that the broadcast will be scoped to. |

#### Example: Create

```go
result, err := client.CreateBroadcastResponseSuccess(nil).Create(map[string]any{
    "from": "example_from",
    "segment_id": "example_segment_id",
    "subject": "example_subject",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### CreateInboxDraftResponseSuccess

Create an instance: `createInboxDraftResponseSuccess := client.CreateInboxDraftResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any` | Files to store on the draft. |
| `bcc` | `any` | The BCC recipients. |
| `cc` | `any` | The CC recipients. |
| `html` | `any` | The HTML body. |
| `id` | `string` | The ID of the draft. |
| `object` | `string` | Type of the response object. |
| `reply_to_email_id` | `string` | The Email ID to reply to, as returned in `data[].id` by the list thread emails endpoint. |
| `subject` | `any` | The subject. |
| `text` | `any` | The plain text body. |
| `thread_id` | `string` | The Thread ID to reply in. |
| `to` | `any` | The recipients. |

#### Example: Create

```go
result, err := client.CreateInboxDraftResponseSuccess(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### CreateInboxLabelResponseSuccess

Create an instance: `createInboxLabelResponseSuccess := client.CreateInboxLabelResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `string` | The color of the label, as a `#RRGGBB` hex color. |
| `name` | `string` | The name of the label. |

#### Example: Create

```go
result, err := client.CreateInboxLabelResponseSuccess(nil).Create(map[string]any{
    "inbox_id": "example_inbox_id",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Domain

Create an instance: `domain := client.Domain(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `capabilities` | `map[string]any` | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | Whether click tracking is enabled for this domain. |
| `created_at` | `string` | The date and time the domain was created. |
| `custom_return_path` | `string` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | `string` | The ID of the domain. |
| `name` | `string` | The name of the domain. |
| `object` | `string` | The type of object. |
| `open_tracking` | `bool` | Whether open tracking is enabled for this domain. |
| `records` | `[]any` |  |
| `region` | `string` | The region where the domain is hosted. |
| `status` | `string` | The status of the domain. |
| `tls` | `string` | TLS mode. |
| `tracking_subdomain` | `string` | The subdomain used for click and open tracking. |

#### Example: Load

```go
domain, err := client.Domain(nil).Load(map[string]any{"id": "domain_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(domain.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
domains, err := client.Domain(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range domains.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Domain(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### DomainClaim

Create an instance: `domainClaim := client.DomainClaim(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `blocked_reason` | `any` | Why the claim is currently blocked, if applicable. |
| `click_tracking` | `bool` | Track clicks within the body of each HTML email. |
| `created_at` | `string` | The date and time the claim was created. |
| `custom_return_path` | `string` | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | `any` | The ID of the placeholder domain created for the claim. |
| `expires_at` | `string` | The date and time the claim expires if not verified. |
| `failure_reason` | `any` | Why the claim failed, if applicable. |
| `id` | `string` | The ID of the claim. |
| `name` | `string` | The name of the domain being claimed. |
| `object` | `string` | The type of object. |
| `open_tracking` | `bool` | Track the open rate of each email. |
| `record` | `map[string]any` | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | `any` | The region where the claimed domain will send from. |
| `status` | `string` | The status of the claim. |
| `tracking_subdomain` | `string` | The subdomain to use for click and open tracking. |

#### Example: Load

```go
domainClaim, err := client.DomainClaim(nil).Load(map[string]any{"id": "domain_claim_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(domainClaim.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: Create

```go
result, err := client.DomainClaim(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Email

Create an instance: `email := client.Email(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` |  |
| `bcc` | `[]any` | The email addresses of the blind carbon copy recipients. |
| `cc` | `[]any` | The email addresses of the carbon copy recipients. |
| `created_at` | `string` | The date and time the email was created. |
| `from` | `string` | The email address of the sender. |
| `headers` | `map[string]any` | Custom headers to add to the email. |
| `html` | `string` | The HTML body of the email. |
| `id` | `string` | The ID of the email. |
| `last_event` | `string` | The status of the email. |
| `message_id` | `string` | The Message-ID header value of the email. |
| `object` | `string` | The type of object. |
| `reply_to` | `[]any` | The email addresses to which replies should be sent. |
| `scheduled_at` | `string` | Schedule email to be sent later. |
| `subject` | `string` | The subject line of the email. |
| `tags` | `[]any` |  |
| `template` | `any` |  |
| `text` | `string` | The plain text body of the email. |
| `to` | `[]any` | Recipient email address. |
| `topic_id` | `string` | The topic ID to scope the email to. |

#### Example: Load

```go
email, err := client.Email(nil).Load(map[string]any{"id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(email.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
emails, err := client.Email(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range emails.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Email(nil).Create(map[string]any{
    "from": "example_from",
    "subject": "example_subject",
    "to": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### EmailsMetric

Create an instance: `emailsMetric := client.EmailsMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

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

```go
emailsMetrics, err := client.EmailsMetric(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range emailsMetrics.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Event

Create an instance: `event := client.Event(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date and time the event was created. |
| `id` | `string` | The event ID. |
| `name` | `string` | The event name. |
| `object` | `string` | Type of the response object. |
| `schema` | `any` | A flat key/type map defining the event payload schema. |
| `updated_at` | `any` | The date and time the event was last updated. |

#### Example: Load

```go
event, err := client.Event(nil).Load(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(event.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
events, err := client.Event(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range events.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Event(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Inbox

Create an instance: `inbox := client.Inbox(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | `string` | The ID of the domain that owns the address. |
| `drafts` | `int` | The number of unsent drafts. |
| `email_address` | `string` | The address of the inbox. |
| `forwarding` | `bool` | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | `any` | The name recipients see when mail is sent from this inbox. |
| `id` | `string` | The ID of the inbox. |
| `last_received` | `any` | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | `string` | Internal name for the inbox. |
| `object` | `string` | The object type. |
| `receiving_address` | `any` | The address to forward mail to when forwarding is enabled. |
| `unread` | `int` | The number of unread threads in the Inbox folder. |
| `updated_at` | `string` | ISO 8601 timestamp when the inbox was last updated. |

#### Example: Load

```go
inbox, err := client.Inbox(nil).Load(map[string]any{"id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(inbox.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
inboxs, err := client.Inbox(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range inboxs.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.Inbox(nil).Create(map[string]any{
    "email_address": "example_email_address",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### InboxAgentResponseSuccess

Create an instance: `inboxAgentResponseSuccess := client.InboxAgentResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |

#### Example: List

```go
inboxAgentResponseSuccesss, err := client.InboxAgentResponseSuccess(nil).List(map[string]any{"id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range inboxAgentResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### InboxDraftListItem

Create an instance: `inboxDraftListItem := client.InboxDraftListItem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` | The files stored on the draft. |
| `bcc` | `[]any` | The BCC recipients. |
| `cc` | `[]any` | The CC recipients. |
| `id` | `string` | The ID of the draft. |
| `reply_to_email_id` | `any` | The Email ID the draft replies to. |
| `snippet` | `any` | A short preview of the body. |
| `subject` | `any` | The subject. |
| `thread_id` | `any` | The Thread ID when the draft is a reply. |
| `to` | `any` | The recipients. |
| `type` | `string` | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `string` | Timestamp when the draft was last saved. |

#### Example: List

```go
inboxDraftListItems, err := client.InboxDraftListItem(nil).List(map[string]any{"id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range inboxDraftListItems.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### InboxDraftResponseSuccess

Create an instance: `inboxDraftResponseSuccess := client.InboxDraftResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` | The files stored on the draft. |
| `bcc` | `[]any` | The BCC recipients. |
| `cc` | `[]any` | The CC recipients. |
| `created_at` | `string` | Timestamp when the draft was created. |
| `email_id` | `any` | The ID of the outbound email queued when the draft was sent. |
| `html` | `any` | The HTML body. |
| `id` | `string` | The ID of the draft. |
| `object` | `string` | The object type. |
| `reply_to_email_id` | `any` | The Email ID the draft replies to. |
| `subject` | `any` | The subject. |
| `text` | `any` | The plain text body. |
| `thread_id` | `any` | The Thread ID when the draft is a reply. |
| `to` | `any` | The recipients. |
| `type` | `string` | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `string` | Timestamp when the draft was last saved. |

#### Example: Load

```go
inboxDraftResponseSuccess, err := client.InboxDraftResponseSuccess(nil).Load(map[string]any{"draft_id": "draft_id", "inbox_id": "inbox_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(inboxDraftResponseSuccess.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: Create

```go
result, err := client.InboxDraftResponseSuccess(nil).Create(map[string]any{
    "draft_id": "example_draft_id",
    "inbox_id": "example_inbox_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### InboxLabel

Create an instance: `inboxLabel := client.InboxLabel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `string` | The color of the label, as a `#RRGGBB` hex color in uppercase. |
| `created_at` | `string` | Timestamp when the label was created. |
| `id` | `string` | The ID of the label. |
| `name` | `string` | The name of the label. |

#### Example: List

```go
inboxLabels, err := client.InboxLabel(nil).List(map[string]any{"id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range inboxLabels.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### InboxThread

Create an instance: `inboxThread := client.InboxThread(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bcc` | `[]any` | The BCC recipients of the most recent email. |
| `cc` | `[]any` | The CC recipients of the most recent email. |
| `email_count` | `int` | The number of emails in the thread. |
| `folder` | `string` | The folder the thread is in. |
| `from` | `any` | The sender email address. |
| `has_attachment` | `bool` | Whether the most recent email in the thread has an attachment. |
| `has_draft` | `bool` | Whether the thread has an unsent draft. |
| `id` | `string` | The ID of the thread. |
| `labels` | `[]any` | The labels attached to the thread. |
| `matched_email_id` | `any` | The ID of the email in the thread that matched the search. |
| `read` | `bool` | True only when every email in the thread is read. |
| `received_at` | `string` | Timestamp when the thread was last active. |
| `subject` | `any` | The subject of the thread. |
| `to` | `[]any` | The recipients of the most recent email. |

#### Example: List

```go
inboxThreads, err := client.InboxThread(nil).List(map[string]any{"id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range inboxThreads.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### InboxThreadEmail

Create an instance: `inboxThreadEmail := client.InboxThreadEmail(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` | The attachments of the message. |
| `bcc` | `[]any` | The BCC recipients of the message. |
| `cc` | `[]any` | The CC recipients of the message. |
| `direction` | `string` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `string` | Sender email address. |
| `html` | `any` | The HTML body. |
| `id` | `string` | The ID of the message. |
| `message_id` | `any` | The Message-ID header of the message. |
| `read` | `bool` | Whether the message has been read. |
| `received_at` | `string` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `[]any` | The Reply-To addresses of the message. |
| `subject` | `any` | The subject of the message. |
| `text` | `any` | The plain-text body. |
| `to` | `[]any` | The recipients of the message. |

#### Example: List

```go
inboxThreadEmails, err := client.InboxThreadEmail(nil).List(map[string]any{"inbox_id": "example", "thread_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range inboxThreadEmails.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### InboxThreadEmailResponseSuccess

Create an instance: `inboxThreadEmailResponseSuccess := client.InboxThreadEmailResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` | The attachments of the message. |
| `bcc` | `[]any` | The BCC recipients of the message. |
| `cc` | `[]any` | The CC recipients of the message. |
| `direction` | `string` | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `string` | Sender email address. |
| `html` | `any` | The HTML body. |
| `id` | `string` | The ID of the message. |
| `message_id` | `any` | The Message-ID header of the message. |
| `read` | `bool` | Whether the message has been read. |
| `received_at` | `string` | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `[]any` | The Reply-To addresses of the message. |
| `subject` | `any` | The subject of the message. |
| `text` | `any` | The plain-text body. |
| `to` | `[]any` | The recipients of the message. |

#### Example: Load

```go
inboxThreadEmailResponseSuccess, err := client.InboxThreadEmailResponseSuccess(nil).Load(map[string]any{"email_id": "email_id", "inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(inboxThreadEmailResponseSuccess.(sdk.Entity).Data()) // the loaded entity's record
```


### InboxThreadResponseSuccess

Create an instance: `inboxThreadResponseSuccess := client.InboxThreadResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `folder` | `string` | The folder the thread is in. |
| `id` | `string` | The ID of the thread. |
| `labels` | `[]any` | The labels attached to the thread. |
| `object` | `string` | The object type. |
| `read` | `bool` | True only when every email in the thread is read. |
| `subject` | `any` | The subject of the thread. |

#### Example: Load

```go
inboxThreadResponseSuccess, err := client.InboxThreadResponseSuccess(nil).Load(map[string]any{"inbox_id": "inbox_id", "thread_id": "thread_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(inboxThreadResponseSuccess.(sdk.Entity).Data()) // the loaded entity's record
```


### ListAttachment

Create an instance: `listAttachment := client.ListAttachment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_disposition` | `any` | How the attachment should be displayed. |
| `content_id` | `string` | The content ID for inline attachments. |
| `content_type` | `string` | The MIME type of the attachment. |
| `download_url` | `string` | Signed URL to download the attachment content. |
| `expires_at` | `string` | Timestamp when the download URL expires. |
| `filename` | `any` | The filename of the attachment. |
| `id` | `string` | The ID of the attachment. |
| `size` | `int` | Size of the attachment in bytes. |

#### Example: List

```go
listAttachments, err := client.ListAttachment(nil).List(map[string]any{"email_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listAttachments.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ListBroadcastClickedLinksResponseSuccess

Create an instance: `listBroadcastClickedLinksResponseSuccess := client.ListBroadcastClickedLinksResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clicks` | `int` | Total number of clicks on this URL. |
| `id` | `string` | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `int` | Number of unique clicks on this URL. |
| `url` | `string` | The URL that was clicked. |

#### Example: List

```go
listBroadcastClickedLinksResponseSuccesss, err := client.ListBroadcastClickedLinksResponseSuccess(nil).List(map[string]any{"broadcast_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listBroadcastClickedLinksResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ListBroadcastRecipientsResponseSuccess

Create an instance: `listBroadcastRecipientsResponseSuccess := client.ListBroadcastRecipientsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bounce_type` | `string` | The type of bounce. |
| `clicked_links` | `[]any` | The links this recipient clicked. |
| `contact_id` | `any` | The ID of the contact associated with this recipient, if one exists. |
| `count` | `int` | The number of times this recipient triggered the event. |
| `email` | `string` | The recipient's email address. |
| `id` | `string` | Opaque cursor identifying this row, used for pagination. |

#### Example: List

```go
listBroadcastRecipientsResponseSuccesss, err := client.ListBroadcastRecipientsResponseSuccess(nil).List(map[string]any{"broadcast_id": "example", "type": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listBroadcastRecipientsResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ListContactSegmentsResponseSuccess

Create an instance: `listContactSegmentsResponseSuccess := client.ListContactSegmentsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the contact was added to the segment. |
| `id` | `string` | Unique identifier for the segment. |
| `name` | `string` | Name of the segment. |

#### Example: List

```go
listContactSegmentsResponseSuccesss, err := client.ListContactSegmentsResponseSuccess(nil).List(map[string]any{"contact_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listContactSegmentsResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ListContactsResponseSuccess

Create an instance: `listContactsResponseSuccess := client.ListContactsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the contact was created. |
| `email` | `string` | Email address of the contact. |
| `first_name` | `any` | First name of the contact. |
| `id` | `string` | Unique identifier for the contact. |
| `last_name` | `any` | Last name of the contact. |
| `unsubscribed` | `bool` | Indicates if the contact is unsubscribed. |

#### Example: List

```go
listContactsResponseSuccesss, err := client.ListContactsResponseSuccess(nil).List(map[string]any{"segment_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listContactsResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ListWebhookEvent

Create an instance: `listWebhookEvent := client.ListWebhookEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the event was created. |
| `id` | `string` | The ID of the webhook event. |
| `status` | `string` | The delivery status of the event for this webhook. |
| `type` | `string` | The type of the event. |

#### Example: List

```go
listWebhookEvents, err := client.ListWebhookEvent(nil).List(map[string]any{"webhook_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listWebhookEvents.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ListWebhookEventAttempt

Create an instance: `listWebhookEventAttempt := client.ListWebhookEventAttempt(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `http_status_code` | `int` | The HTTP status code returned by the webhook endpoint. |
| `id` | `string` | The ID of the webhook event attempt. |
| `response` | `string` | The response body returned by the webhook endpoint. |
| `sent_at` | `string` | Timestamp indicating when the attempt was sent. |

#### Example: List

```go
listWebhookEventAttempts, err := client.ListWebhookEventAttempt(nil).List(map[string]any{"event_id": "example", "webhook_id": "example"}, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range listWebhookEventAttempts.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Log

Create an instance: `log := client.Log(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | The date the log was created. |
| `endpoint` | `string` | The API endpoint that was called. |
| `id` | `string` | The log ID. |
| `method` | `string` | The HTTP method used. |
| `object` | `string` | Type of the response object. |
| `request_body` | `any` | The request body sent to the API. |
| `response_body` | `any` | The response body returned by the API. |
| `response_status` | `int` | The HTTP status code of the response. |
| `user_agent` | `any` | The user agent of the request. |

#### Example: Load

```go
log, err := client.Log(nil).Load(map[string]any{"id": "log_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(log.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
logs, err := client.Log(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range logs.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### ReceivedEmail

Create an instance: `receivedEmail := client.ReceivedEmail(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `[]any` | Array of attachments. |
| `bcc` | `any` | The BCC recipients. |
| `cc` | `any` | The CC recipients. |
| `created_at` | `string` | Timestamp when the email was received. |
| `from` | `string` | The sender email address. |
| `headers` | `any` | The email headers. |
| `html` | `any` | The HTML content of the email. |
| `id` | `string` | The ID of the received email. |
| `message_id` | `string` | The unique message ID from the email headers. |
| `object` | `string` | The type of object. |
| `received_for` | `[]any` | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | `any` | The reply-to addresses. |
| `subject` | `string` | The email subject. |
| `text` | `any` | The plain text content of the email. |
| `to` | `[]any` | The recipient email addresses. |

#### Example: Load

```go
receivedEmail, err := client.ReceivedEmail(nil).Load(map[string]any{"email_id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(receivedEmail.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
receivedEmails, err := client.ReceivedEmail(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range receivedEmails.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### RemoveAudienceResponseSuccess

Create an instance: `removeAudienceResponseSuccess := client.RemoveAudienceResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RemoveBroadcastResponseSuccess

Create an instance: `removeBroadcastResponseSuccess := client.RemoveBroadcastResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RemoveContactFromSegmentResponseSuccess

Create an instance: `removeContactFromSegmentResponseSuccess := client.RemoveContactFromSegmentResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### RemoveContactPropertyResponseSuccess

Create an instance: `removeContactPropertyResponseSuccess := client.RemoveContactPropertyResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RemoveContactResponseSuccess

Create an instance: `removeContactResponseSuccess := client.RemoveContactResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RemoveEvent

Create an instance: `removeEvent := client.RemoveEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |


### RemoveInboxDraftResponseSuccess

Create an instance: `removeInboxDraftResponseSuccess := client.RemoveInboxDraftResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### RemoveInboxLabelResponseSuccess

Create an instance: `removeInboxLabelResponseSuccess := client.RemoveInboxLabelResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### RemoveInboxResponseSuccess

Create an instance: `removeInboxResponseSuccess := client.RemoveInboxResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### RemoveInboxThreadLabelResponseSuccess

Create an instance: `removeInboxThreadLabelResponseSuccess := client.RemoveInboxThreadLabelResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### RemoveInboxThreadResponseSuccess

Create an instance: `removeInboxThreadResponseSuccess := client.RemoveInboxThreadResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### RemoveSegmentResponseSuccess

Create an instance: `removeSegmentResponseSuccess := client.RemoveSegmentResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | The ID of the audience this segment belongs to. |
| `created_at` | `string` | Timestamp indicating when the segment was created. |
| `filter` | `map[string]any` | Filter conditions for the segment. |
| `id` | `string` | Unique identifier for the segment. |
| `name` | `string` | The name of the segment. |

#### Example: List

```go
removeSegmentResponseSuccesss, err := client.RemoveSegmentResponseSuccess(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range removeSegmentResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.RemoveSegmentResponseSuccess(nil).Create(map[string]any{
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### RemoveSuppressionResponseSuccess

Create an instance: `removeSuppressionResponseSuccess := client.RemoveSuppressionResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the suppression was created. |
| `email` | `string` | Email address to suppress. |
| `id` | `string` | Unique identifier for the suppression. |
| `origin` | `string` | Origin of the suppression. |
| `source_id` | `any` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

#### Example: List

```go
removeSuppressionResponseSuccesss, err := client.RemoveSuppressionResponseSuccess(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range removeSuppressionResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.RemoveSuppressionResponseSuccess(nil).Create(map[string]any{
    "email": "example_email",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### RemoveTemplateResponseSuccess

Create an instance: `removeTemplateResponseSuccess := client.RemoveTemplateResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alias` | `string` | The alias of the template. |
| `created_at` | `string` | Timestamp indicating when the template was created. |
| `from` | `string` | Sender email address. |
| `html` | `string` | The HTML version of the template. |
| `id` | `string` | The ID of the template. |
| `name` | `string` | The name of the template. |
| `published_at` | `any` | Timestamp indicating when the template was published. |
| `reply_to` | `[]any` | Reply-to email addresses. |
| `status` | `string` | The publication status of the template. |
| `subject` | `string` | Email subject. |
| `text` | `string` | The plain text version of the template. |
| `updated_at` | `string` | Timestamp indicating when the template was last updated. |
| `variables` | `[]any` |  |

#### Example: List

```go
removeTemplateResponseSuccesss, err := client.RemoveTemplateResponseSuccess(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range removeTemplateResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.RemoveTemplateResponseSuccess(nil).Create(map[string]any{
    "html": "example_html",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### RemoveTopicResponseSuccess

Create an instance: `removeTopicResponseSuccess := client.RemoveTopicResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the topic was created. |
| `default_subscription` | `string` | The default subscription status for the topic. |
| `description` | `string` | A description of the topic. |
| `id` | `string` | Unique identifier for the topic. |
| `name` | `string` | The name of the topic. |
| `visibility` | `string` | The visibility of the topic. |

#### Example: List

```go
removeTopicResponseSuccesss, err := client.RemoveTopicResponseSuccess(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range removeTopicResponseSuccesss.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.RemoveTopicResponseSuccess(nil).Create(map[string]any{
    "default_subscription": "example_default_subscription",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### RetrievedAttachment

Create an instance: `retrievedAttachment := client.RetrievedAttachment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_disposition` | `any` | How the attachment should be displayed. |
| `content_id` | `string` | The content ID for inline attachments. |
| `content_type` | `string` | The MIME type of the attachment. |
| `download_url` | `string` | Signed URL to download the attachment content. |
| `expires_at` | `string` | Timestamp when the download URL expires. |
| `filename` | `any` | The filename of the attachment. |
| `id` | `string` | The ID of the attachment. |
| `object` | `string` | The type of object. |
| `size` | `int` | Size of the attachment in bytes. |

#### Example: Load

```go
retrievedAttachment, err := client.RetrievedAttachment(nil).Load(map[string]any{"id": "retrieved_attachment_id", "email_id": "email_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(retrievedAttachment.(sdk.Entity).Data()) // the loaded entity's record
```


### RevokeOAuthGrant

Create an instance: `revokeOAuthGrant := client.RevokeOAuthGrant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client` | `map[string]any` | The OAuth client the grant was issued to. |
| `client_id` | `string` | The ID of the OAuth client the grant was issued to. |
| `created_at` | `string` | The date and time the OAuth grant was created. |
| `id` | `string` | The ID of the OAuth grant. |
| `revoked_at` | `any` | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | `any` | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | `[]any` | The scopes granted to the OAuth client. |

#### Example: List

```go
revokeOAuthGrants, err := client.RevokeOAuthGrant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range revokeOAuthGrants.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Rotate

Create an instance: `rotate := client.Rotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` | The ID of the webhook. |
| `object` | `string` | The type of object. |
| `signing_secret` | `string` | The new secret key used to verify webhook payloads. |

#### Example: Create

```go
result, err := client.Rotate(nil).Create(map[string]any{
    "webhook_id": "example_webhook_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Segment

Create an instance: `segment := client.Segment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | The ID of the audience this segment belongs to. |
| `created_at` | `string` | Timestamp indicating when the segment was created. |
| `filter` | `map[string]any` | Filter conditions for the segment. |
| `id` | `string` | The ID of the segment. |
| `name` | `string` | The name of the segment. |
| `object` | `string` | The object type. |

#### Example: Load

```go
segment, err := client.Segment(nil).Load(map[string]any{"id": "segment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(segment.(sdk.Entity).Data()) // the loaded entity's record
```


### SetInboxThreadLabelsResponseSuccess

Create an instance: `setInboxThreadLabelsResponseSuccess := client.SetInboxThreadLabelsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `[]any` | Every label on the thread after the change. |
| `labels` | `[]any` | The IDs of the labels the thread carries after the change. |
| `object` | `string` | Type of the response object. |


### Suppression

Create an instance: `suppression := client.Suppression(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the suppression was created. |
| `email` | `string` | Email address that is suppressed. |
| `id` | `string` | Unique identifier for the suppression. |
| `object` | `string` | Type of the response object. |
| `origin` | `string` | Origin of the suppression. |
| `source_id` | `any` | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

#### Example: Load

```go
suppression, err := client.Suppression(nil).Load(map[string]any{"id": "suppression_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(suppression.(sdk.Entity).Data()) // the loaded entity's record
```


### Template

Create an instance: `template := client.Template(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alias` | `string` | The alias of the template. |
| `created_at` | `string` | Timestamp indicating when the template was created. |
| `current_version_id` | `string` | The ID of the current version of the template. |
| `from` | `string` | Sender email address. |
| `has_unpublished_versions` | `bool` | Indicates whether the template has unpublished versions. |
| `html` | `string` | The HTML version of the template. |
| `id` | `string` | The ID of the template. |
| `name` | `string` | The name of the template. |
| `object` | `string` | The type of object. |
| `published_at` | `any` | Timestamp indicating when the template was published. |
| `reply_to` | `any` | Reply-to email addresses. |
| `status` | `string` | The publication status of the template. |
| `subject` | `string` | Email subject. |
| `text` | `string` | The plain text version of the template. |
| `updated_at` | `string` | Timestamp indicating when the template was last updated. |
| `variables` | `[]any` |  |

#### Example: Load

```go
template, err := client.Template(nil).Load(map[string]any{"id": "template_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(template.(sdk.Entity).Data()) // the loaded entity's record
```


### Topic

Create an instance: `topic := client.Topic(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
topic, err := client.Topic(nil).Load(map[string]any{"id": "topic_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(topic.(sdk.Entity).Data()) // the loaded entity's record
```


### UpdateApiKey

Create an instance: `updateApiKey := client.UpdateApiKey(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` | The ID of the API key. |
| `name` | `string` | The API key name. |
| `object` | `string` | The type of object. |


### UpdateBroadcastResponseSuccess

Create an instance: `updateBroadcastResponseSuccess := client.UpdateBroadcastResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `audience_id` | `string` | Use `segment_id` instead. |
| `from` | `string` | The email address of the sender. |
| `html` | `string` | The HTML version of the message. |
| `id` | `string` | The ID of the broadcast. |
| `name` | `string` | Name of the broadcast. |
| `object` | `string` | The object type of the response. |
| `preview_text` | `string` | The preview text of the email. |
| `reply_to` | `[]any` | The email addresses to which replies should be sent. |
| `segment_id` | `string` | Unique identifier of the segment this broadcast will be sent to. |
| `subject` | `string` | The subject line of the email. |
| `text` | `string` | The plain text version of the message. |
| `topic_id` | `string` | The topic ID that the broadcast will be scoped to. |


### UpdateContactPropertyResponseSuccess

Create an instance: `updateContactPropertyResponseSuccess := client.UpdateContactPropertyResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fallback_value` | `any` | The default value to use when the property is not set for a contact. |
| `id` | `string` | The ID of the contact property. |
| `object` | `string` | The object type. |


### UpdateContactResponseSuccess

Create an instance: `updateContactResponseSuccess := client.UpdateContactResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` | Email address of the contact. |
| `first_name` | `string` | First name of the contact. |
| `id` | `string` | Unique identifier for the updated contact. |
| `last_name` | `string` | Last name of the contact. |
| `object` | `string` | Type of the response object. |
| `properties` | `map[string]any` | A map of custom property keys and values to update. |
| `unsubscribed` | `bool` | The Contact's global subscription status. |


### UpdateContactTopicsResponseSuccess

Create an instance: `updateContactTopicsResponseSuccess := client.UpdateContactTopicsResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` | The ID of the contact. |
| `object` | `string` | The object type. |
| `topics` | `[]any` | Array of updated topic subscriptions. |


### UpdateDomainResponseSuccess

Create an instance: `updateDomainResponseSuccess := client.UpdateDomainResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `capabilities` | `map[string]any` | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | Track clicks within the body of each HTML email. |
| `id` | `string` | The ID of the updated domain. |
| `object` | `string` | The object type representing the updated domain. |
| `open_tracking` | `bool` | Track the open rate of each email. |
| `tls` | `string` | enforced | opportunistic. |
| `tracking_subdomain` | `string` | The subdomain to use for click and open tracking. |


### UpdateEmailOption

Create an instance: `updateEmailOption := client.UpdateEmailOption(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `scheduled_at` | `string` | Schedule email to be sent later. |


### UpdateEvent

Create an instance: `updateEvent := client.UpdateEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` | The ID of the updated event. |
| `object` | `string` | Type of the response object. |
| `schema` | `any` | A flat key/type map defining the event payload schema. |


### UpdateInboxDraftResponseSuccess

Create an instance: `updateInboxDraftResponseSuccess := client.UpdateInboxDraftResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `attachments` | `any` | Files to store on the draft. |
| `bcc` | `any` | The BCC recipients. |
| `cc` | `any` | The CC recipients. |
| `html` | `any` | The HTML body. |
| `id` | `string` | The ID of the draft. |
| `object` | `string` | The object type. |
| `subject` | `any` | The subject. |
| `text` | `any` | The plain text body. |
| `to` | `any` | The recipients. |


### UpdateInboxLabelResponseSuccess

Create an instance: `updateInboxLabelResponseSuccess := client.UpdateInboxLabelResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `string` | A new color for the label, as a `#RRGGBB` hex color. |
| `id` | `string` | The ID of the updated label. |
| `name` | `string` | A new name for the label. |
| `object` | `string` | Type of the response object. |


### UpdateInboxResponseSuccess

Create an instance: `updateInboxResponseSuccess := client.UpdateInboxResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from_name` | `string` | The name recipients see when mail is sent from this inbox. |
| `id` | `string` | The ID of the inbox. |
| `name` | `string` | New internal name for the inbox. |
| `object` | `string` | The object type. |


### UpdateInboxThreadResponseSuccess

Create an instance: `updateInboxThreadResponseSuccess := client.UpdateInboxThreadResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `folder` | `string` | The folder the thread is in. |
| `id` | `string` | The ID of the thread. |
| `labels` | `[]any` | The labels attached to the thread. |
| `object` | `string` | The object type. |
| `read` | `bool` | True only when every email in the thread is read. |
| `subject` | `any` | The subject of the thread. |


### UpdateSegmentResponseSuccess

Create an instance: `updateSegmentResponseSuccess := client.UpdateSegmentResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` | The ID of the segment. |
| `name` | `string` | The name of the segment. |
| `object` | `string` | The object type. |


### UpdateTemplateResponseSuccess

Create an instance: `updateTemplateResponseSuccess := client.UpdateTemplateResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `alias` | `string` | The alias of the template. |
| `from` | `string` | Sender email address. |
| `html` | `string` | The HTML version of the template. |
| `id` | `string` | The ID of the template. |
| `name` | `string` | The name of the template. |
| `object` | `string` | The object type of the response. |
| `reply_to` | `[]any` | Reply-to email addresses. |
| `subject` | `string` | Email subject. |
| `text` | `string` | The plain text version of the template. |
| `variables` | `[]any` |  |


### UpdateTopicResponseSuccess

Create an instance: `updateTopicResponseSuccess := client.UpdateTopicResponseSuccess(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` | A description of the topic. |
| `id` | `string` | The ID of the topic. |
| `name` | `string` | The name of the topic. |
| `object` | `string` | The object type. |
| `visibility` | `string` | The visibility of the topic. |


### UpdateWebhook

Create an instance: `updateWebhook := client.UpdateWebhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the webhook was created. |
| `endpoint` | `string` | The URL where webhook events will be sent. |
| `events` | `[]any` | Array of event types to subscribe to. |
| `id` | `string` | The ID of the updated webhook. |
| `object` | `string` | The type of object. |
| `status` | `string` | The status of the webhook. |

#### Example: List

```go
updateWebhooks, err := client.UpdateWebhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range updateWebhooks.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

#### Example: Create

```go
result, err := client.UpdateWebhook(nil).Create(map[string]any{
    "endpoint": "example_endpoint",
    "events": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
```


### Usage

Create an instance: `usage := client.Usage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ai_credits` | `map[string]any` |  |
| `automation_runs` | `map[string]any` |  |
| `broadcasts` | `map[string]any` |  |
| `contacts` | `map[string]any` |  |
| `domains` | `map[string]any` |  |
| `emails` | `map[string]any` |  |
| `object` | `string` | The type of object. |
| `rate_limit` | `map[string]any` |  |
| `segments` | `map[string]any` |  |

#### Example: Load

```go
usage, err := client.Usage(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(usage.(sdk.Entity).Data()) // the loaded entity's record
```


### Webhook

Create an instance: `webhook := client.Webhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the webhook was created. |
| `endpoint` | `string` | The URL where webhook events are sent. |
| `events` | `any` | Array of event types subscribed to. |
| `id` | `string` | The ID of the webhook. |
| `object` | `string` | The type of object. |
| `signing_secret` | `string` | The secret key used to verify webhook payloads. |
| `status` | `string` | The status of the webhook. |

#### Example: Load

```go
webhook, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(webhook.(sdk.Entity).Data()) // the loaded entity's record
```


### WebhookEvent

Create an instance: `webhookEvent := client.WebhookEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` | Timestamp indicating when the event was created. |
| `id` | `string` | The ID of the webhook event. |
| `next_attempt_at` | `any` | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | `string` | The type of object. |
| `payload` | `map[string]any` | The event payload sent to the webhook endpoint. |
| `status` | `string` | The delivery status of the event for this webhook. |
| `type` | `string` | The type of the event. |

#### Example: Load

```go
webhookEvent, err := client.WebhookEvent(nil).Load(map[string]any{"id": "webhook_event_id", "webhook_id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(webhookEvent.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: Create

```go
result, err := client.WebhookEvent(nil).Create(map[string]any{
    "event_id": "example_event_id",
    "webhook_id": "example_webhook_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data()) // the created entity's record
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

3 fields are carried as open values rather than typed structures.
This follows from the API definition, not from a gap in this SDK: the
definition describes them with untagged unions —
`oneOf`/`anyOf` branches with no `discriminator` — so it never states which
variant a given value is. Nothing can select a branch reliably, so the SDK
passes the value through unchanged rather than assert a shape the API does not
guarantee.

| Entity | Field | Variants | Nesting |
| --- | --- | --- | --- |
| `remove_template_response_success` | `variables` | 5 | 3 levels |
| `template` | `variables` | 5 | 3 levels |
| `update_template_response_success` | `variables` | 5 | 3 levels |

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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

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

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

An operation returns the entity, and its `Data()` returns the record. Use
`core.ToMapAny()` to safely cast that record, or data nested in it, to
`map[string]any`: it returns `nil` for anything else, an entity included.

### Package structure

```
github.com/voxgig-sdk/resend-sdk/go/
├── resend.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/resend-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
automationrun := client.AutomationRun(nil)
automationrun.Load(map[string]any{"automation_id": "example", "id": "example_id"}, nil)

// automationrun.Data() now returns the automationrun data from the last load
// automationrun.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
