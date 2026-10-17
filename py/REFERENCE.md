# Resend Python SDK Reference

Complete API reference for the Resend Python SDK.


## ResendSDK

### Constructor

```python
from resend_sdk import ResendSDK

client = ResendSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `ResendSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = ResendSDK.test()
```


### Instance Methods

#### `AddContactToSegmentResponseSuccess(data=None)`

Create a new `AddContactToSegmentResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `AddInboxThreadLabelsResponseSuccess(data=None)`

Create a new `AddInboxThreadLabelsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ApiKey(data=None)`

Create a new `ApiKeyEntity` instance. Pass `None` for no initial data.

#### `Audience(data=None)`

Create a new `AudienceEntity` instance. Pass `None` for no initial data.

#### `Automation(data=None)`

Create a new `AutomationEntity` instance. Pass `None` for no initial data.

#### `AutomationRun(data=None)`

Create a new `AutomationRunEntity` instance. Pass `None` for no initial data.

#### `AutomationRunListItem(data=None)`

Create a new `AutomationRunListItemEntity` instance. Pass `None` for no initial data.

#### `BatchAddSuppressionsResponseSuccess(data=None)`

Create a new `BatchAddSuppressionsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `BatchRemoveSuppressionsResponseSuccess(data=None)`

Create a new `BatchRemoveSuppressionsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `Broadcast(data=None)`

Create a new `BroadcastEntity` instance. Pass `None` for no initial data.

#### `Contact(data=None)`

Create a new `ContactEntity` instance. Pass `None` for no initial data.

#### `ContactImportResponseSuccess(data=None)`

Create a new `ContactImportResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ContactProperty(data=None)`

Create a new `ContactPropertyEntity` instance. Pass `None` for no initial data.

#### `ContactTopicsResponseSuccess(data=None)`

Create a new `ContactTopicsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `CreateApiKey(data=None)`

Create a new `CreateApiKeyEntity` instance. Pass `None` for no initial data.

#### `CreateBatchEmail(data=None)`

Create a new `CreateBatchEmailEntity` instance. Pass `None` for no initial data.

#### `CreateBroadcastResponseSuccess(data=None)`

Create a new `CreateBroadcastResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `CreateInboxDraftResponseSuccess(data=None)`

Create a new `CreateInboxDraftResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `CreateInboxLabelResponseSuccess(data=None)`

Create a new `CreateInboxLabelResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `Domain(data=None)`

Create a new `DomainEntity` instance. Pass `None` for no initial data.

#### `DomainClaim(data=None)`

Create a new `DomainClaimEntity` instance. Pass `None` for no initial data.

#### `Email(data=None)`

Create a new `EmailEntity` instance. Pass `None` for no initial data.

#### `EmailsMetric(data=None)`

Create a new `EmailsMetricEntity` instance. Pass `None` for no initial data.

#### `Event(data=None)`

Create a new `EventEntity` instance. Pass `None` for no initial data.

#### `Inbox(data=None)`

Create a new `InboxEntity` instance. Pass `None` for no initial data.

#### `InboxAgentResponseSuccess(data=None)`

Create a new `InboxAgentResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `InboxDraftListItem(data=None)`

Create a new `InboxDraftListItemEntity` instance. Pass `None` for no initial data.

#### `InboxDraftResponseSuccess(data=None)`

Create a new `InboxDraftResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `InboxLabel(data=None)`

Create a new `InboxLabelEntity` instance. Pass `None` for no initial data.

#### `InboxThread(data=None)`

Create a new `InboxThreadEntity` instance. Pass `None` for no initial data.

#### `InboxThreadEmail(data=None)`

Create a new `InboxThreadEmailEntity` instance. Pass `None` for no initial data.

#### `InboxThreadEmailResponseSuccess(data=None)`

Create a new `InboxThreadEmailResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `InboxThreadResponseSuccess(data=None)`

Create a new `InboxThreadResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ListAttachment(data=None)`

Create a new `ListAttachmentEntity` instance. Pass `None` for no initial data.

#### `ListBroadcastClickedLinksResponseSuccess(data=None)`

Create a new `ListBroadcastClickedLinksResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ListBroadcastRecipientsResponseSuccess(data=None)`

Create a new `ListBroadcastRecipientsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ListContactSegmentsResponseSuccess(data=None)`

Create a new `ListContactSegmentsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ListContactsResponseSuccess(data=None)`

Create a new `ListContactsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `ListWebhookEvent(data=None)`

Create a new `ListWebhookEventEntity` instance. Pass `None` for no initial data.

#### `ListWebhookEventAttempt(data=None)`

Create a new `ListWebhookEventAttemptEntity` instance. Pass `None` for no initial data.

#### `Log(data=None)`

Create a new `LogEntity` instance. Pass `None` for no initial data.

#### `ReceivedEmail(data=None)`

Create a new `ReceivedEmailEntity` instance. Pass `None` for no initial data.

#### `RemoveAudienceResponseSuccess(data=None)`

Create a new `RemoveAudienceResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveBroadcastResponseSuccess(data=None)`

Create a new `RemoveBroadcastResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveContactFromSegmentResponseSuccess(data=None)`

Create a new `RemoveContactFromSegmentResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveContactPropertyResponseSuccess(data=None)`

Create a new `RemoveContactPropertyResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveContactResponseSuccess(data=None)`

Create a new `RemoveContactResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveEvent(data=None)`

Create a new `RemoveEventEntity` instance. Pass `None` for no initial data.

#### `RemoveInboxDraftResponseSuccess(data=None)`

Create a new `RemoveInboxDraftResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveInboxLabelResponseSuccess(data=None)`

Create a new `RemoveInboxLabelResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveInboxResponseSuccess(data=None)`

Create a new `RemoveInboxResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveInboxThreadLabelResponseSuccess(data=None)`

Create a new `RemoveInboxThreadLabelResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveInboxThreadResponseSuccess(data=None)`

Create a new `RemoveInboxThreadResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveSegmentResponseSuccess(data=None)`

Create a new `RemoveSegmentResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveSuppressionResponseSuccess(data=None)`

Create a new `RemoveSuppressionResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveTemplateResponseSuccess(data=None)`

Create a new `RemoveTemplateResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RemoveTopicResponseSuccess(data=None)`

Create a new `RemoveTopicResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `RetrievedAttachment(data=None)`

Create a new `RetrievedAttachmentEntity` instance. Pass `None` for no initial data.

#### `RevokeOAuthGrant(data=None)`

Create a new `RevokeOAuthGrantEntity` instance. Pass `None` for no initial data.

#### `Rotate(data=None)`

Create a new `RotateEntity` instance. Pass `None` for no initial data.

#### `Segment(data=None)`

Create a new `SegmentEntity` instance. Pass `None` for no initial data.

#### `SetInboxThreadLabelsResponseSuccess(data=None)`

Create a new `SetInboxThreadLabelsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `Suppression(data=None)`

Create a new `SuppressionEntity` instance. Pass `None` for no initial data.

#### `Template(data=None)`

Create a new `TemplateEntity` instance. Pass `None` for no initial data.

#### `Topic(data=None)`

Create a new `TopicEntity` instance. Pass `None` for no initial data.

#### `UpdateApiKey(data=None)`

Create a new `UpdateApiKeyEntity` instance. Pass `None` for no initial data.

#### `UpdateBroadcastResponseSuccess(data=None)`

Create a new `UpdateBroadcastResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateContactPropertyResponseSuccess(data=None)`

Create a new `UpdateContactPropertyResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateContactResponseSuccess(data=None)`

Create a new `UpdateContactResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateContactTopicsResponseSuccess(data=None)`

Create a new `UpdateContactTopicsResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateDomainResponseSuccess(data=None)`

Create a new `UpdateDomainResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateEmailOption(data=None)`

Create a new `UpdateEmailOptionEntity` instance. Pass `None` for no initial data.

#### `UpdateEvent(data=None)`

Create a new `UpdateEventEntity` instance. Pass `None` for no initial data.

#### `UpdateInboxDraftResponseSuccess(data=None)`

Create a new `UpdateInboxDraftResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateInboxLabelResponseSuccess(data=None)`

Create a new `UpdateInboxLabelResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateInboxResponseSuccess(data=None)`

Create a new `UpdateInboxResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateInboxThreadResponseSuccess(data=None)`

Create a new `UpdateInboxThreadResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateSegmentResponseSuccess(data=None)`

Create a new `UpdateSegmentResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateTemplateResponseSuccess(data=None)`

Create a new `UpdateTemplateResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateTopicResponseSuccess(data=None)`

Create a new `UpdateTopicResponseSuccessEntity` instance. Pass `None` for no initial data.

#### `UpdateWebhook(data=None)`

Create a new `UpdateWebhookEntity` instance. Pass `None` for no initial data.

#### `Usage(data=None)`

Create a new `UsageEntity` instance. Pass `None` for no initial data.

#### `Webhook(data=None)`

Create a new `WebhookEntity` instance. Pass `None` for no initial data.

#### `WebhookEvent(data=None)`

Create a new `WebhookEventEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AddContactToSegmentResponseSuccessEntity

```python
add_contact_to_segment_response_success = client.AddContactToSegmentResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `str` | No | The ID of the contact. |
| `object` | `str` | No | The object type. |
| `segment_id` | `str` | No | The ID of the segment. |

### Operations

#### `create(reqdata, ctrl=None) -> AddContactToSegmentResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.AddContactToSegmentResponseSuccess().create({
    "contact_id": "example_contact_id",  # str
    "segment_id": "example_segment_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AddContactToSegmentResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AddInboxThreadLabelsResponseSuccessEntity

```python
add_inbox_thread_labels_response_success = client.AddInboxThreadLabelsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | Yes | Every label on the thread after the change. |
| `labels` | `list` | Yes | The IDs of the labels to apply to the thread. |
| `object` | `str` | Yes | Type of the response object. |

### Operations

#### `create(reqdata, ctrl=None) -> AddInboxThreadLabelsResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.AddInboxThreadLabelsResponseSuccess().create({
    "inbox_id": "example_inbox_id",  # str
    "thread_id": "example_thread_id",  # str
    "labels": [],  # list
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AddInboxThreadLabelsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ApiKeyEntity

```python
api_key = client.ApiKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | The date and time the API key was created. |
| `id` | `str` | No | The ID of the API key. |
| `last_used_at` | `str | None` | No | The date and time the API key was last used. |
| `name` | `str` | No | The name of the API key. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ApiKeyEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ApiKey().list()
for api_key in results:
    print(api_key.data_get())
```

#### `remove(reqmatch, ctrl=None) -> ApiKeyEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.ApiKey().remove({"id": "api_key_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ApiKeyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AudienceEntity

```python
audience = client.Audience()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | The date that the object was created. |
| `id` | `str` | No | The ID of the audience. |
| `name` | `str` | No | The name of the audience. |
| `object` | `str` | No | The object of the audience. |

### Field Usage by Operation

| Field | load | list | create |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `id` | - | - | - |
| `name` | - | - | Yes |
| `object` | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> AudienceEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Audience().create({
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[AudienceEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Audience().list()
for audience in results:
    print(audience.data_get())
```

#### `load(reqmatch, ctrl=None) -> AudienceEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Audience().load({"id": "audience_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AudienceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AutomationEntity

```python
automation = client.Automation()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `connections` | `list` | No | The connections between steps in the active version of the automation. |
| `created_at` | `str` | No | The date and time the automation was created. |
| `id` | `str` | No | The ID of the automation. |
| `name` | `str` | No | The name of the automation. |
| `object` | `str` | No | Type of the response object. |
| `status` | `str` | No | The current status of the automation. |
| `steps` | `list` | No | The steps in the active version of the automation. |
| `updated_at` | `str` | No | The date and time the automation was last updated. |

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

#### `create(reqdata, ctrl=None) -> AutomationEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Automation().create({
    "connections": [],  # list
    "name": "example_name",  # str
    "steps": [],  # list
})
```

#### `list(reqmatch=None, ctrl=None) -> list[AutomationEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Automation().list()
for automation in results:
    print(automation.data_get())
```

#### `load(reqmatch, ctrl=None) -> AutomationEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Automation().load({"id": "automation_id"})
```

#### `remove(reqmatch, ctrl=None) -> AutomationEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Automation().remove({"id": "automation_id"})
```

#### `update(reqdata, ctrl=None) -> AutomationEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Automation().update({
    "id": "automation_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AutomationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AutomationRunEntity

```python
automation_run = client.AutomationRun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `completed_at` | `str | None` | No | The date and time the run completed. |
| `created_at` | `str` | No | The date and time the run was created. |
| `id` | `str` | No | The ID of the automation run. |
| `object` | `str` | No | Type of the response object. |
| `started_at` | `str | None` | No | The date and time the run started. |
| `status` | `str` | No | The current status of the automation run. |
| `steps` | `list` | No | The steps executed in this run, sorted in graph order. |

### Operations

#### `load(reqmatch, ctrl=None) -> AutomationRunEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.AutomationRun().load({"id": "automation_run_id", "automation_id": "automation_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AutomationRunEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AutomationRunListItemEntity

```python
automation_run_list_item = client.AutomationRunListItem()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `completed_at` | `str | None` | No | The date and time the run completed. |
| `created_at` | `str` | No | The date and time the run was created. |
| `id` | `str` | No | The ID of the automation run. |
| `started_at` | `str | None` | No | The date and time the run started. |
| `status` | `str` | No | The current status of the automation run. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[AutomationRunListItemEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.AutomationRunListItem().list({"id": "example"})
for automation_run_list_item in results:
    print(automation_run_list_item.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AutomationRunListItemEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BatchAddSuppressionsResponseSuccessEntity

```python
batch_add_suppressions_response_success = client.BatchAddSuppressionsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `emails` | `list` | Yes | Email addresses to suppress. |

### Operations

#### `create(reqdata, ctrl=None) -> BatchAddSuppressionsResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.BatchAddSuppressionsResponseSuccess().create({
    "emails": [],  # list
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchAddSuppressionsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BatchRemoveSuppressionsResponseSuccessEntity

```python
batch_remove_suppressions_response_success = client.BatchRemoveSuppressionsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | No | Array containing the removed suppressions. |
| `emails` | `list` | No | Email addresses to remove from the suppression list. |
| `ids` | `list` | No | Suppression IDs to remove from the suppression list. |

### Operations

#### `create(reqdata, ctrl=None) -> BatchRemoveSuppressionsResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.BatchRemoveSuppressionsResponseSuccess().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchRemoveSuppressionsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BroadcastEntity

```python
broadcast = client.Broadcast()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `str | None` | No | Deprecated: use `segment_id` instead. |
| `created_at` | `str` | No | Timestamp indicating when the broadcast was created. |
| `from` | `str` | No | The email address of the sender. |
| `html` | `str | None` | No | The HTML version of the broadcast content. |
| `id` | `str` | No | Unique identifier for the broadcast. |
| `name` | `str` | No | Name of the broadcast. |
| `preview_text` | `str` | No | The preview text of the email. |
| `reply_to` | `list` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `str` | No | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `str | None` | No | Unique identifier of the segment this broadcast will be sent to. |
| `sent_at` | `str` | No | Timestamp indicating when the broadcast was sent. |
| `status` | `str` | No | The status of the broadcast. |
| `subject` | `str` | No | The subject line of the email. |
| `text` | `str | None` | No | The plain text version of the broadcast content. |
| `topic_id` | `str | None` | No | The topic ID that the broadcast is scoped to. |

### Operations

#### `create(reqdata, ctrl=None) -> BroadcastEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

#### `list(reqmatch=None, ctrl=None) -> list[BroadcastEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Broadcast().list()
for broadcast in results:
    print(broadcast.data_get())
```

#### `load(reqmatch, ctrl=None) -> BroadcastEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Broadcast().load({"id": "broadcast_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BroadcastEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactEntity

```python
contact = client.Contact()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `str` | No | Unique identifier of the audience to which the contact belongs. |
| `created_at` | `str` | No | Timestamp indicating when the contact was created. |
| `email` | `str` | No | Email address of the contact. |
| `first_name` | `str | None` | No | First name of the contact. |
| `id` | `str` | No | Unique identifier for the contact. |
| `last_name` | `str | None` | No | Last name of the contact. |
| `object` | `str` | No | Type of the response object. |
| `properties` | `dict` | No | A map of custom property keys and values. |
| `segments` | `list` | No | Array of segment IDs to add the contact to. |
| `topics` | `list` | No | Array of topic subscriptions for the contact. |
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

#### `create(reqdata, ctrl=None) -> ContactEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Contact().create({
    "email": "example_email",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[ContactEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Contact().list()
for contact in results:
    print(contact.data_get())
```

#### `load(reqmatch, ctrl=None) -> ContactEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Contact().load({"id": "contact_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactImportResponseSuccessEntity

```python
contact_import_response_success = client.ContactImportResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `column_map` | `str` | No | JSON-encoded object mapping contact fields and custom property keys to CSV column names. |
| `completed_at` | `str | None` | No | Timestamp indicating when the contact import completed. |
| `counts` | `dict` | No |  |
| `created_at` | `str` | No | Timestamp indicating when the contact import was created. |
| `file` | `str` | Yes | CSV file to import. |
| `id` | `str` | No | Unique identifier for the contact import. |
| `object` | `str` | No | Type of the response object. |
| `on_conflict` | `str` | No | Strategy to use when an imported contact already exists. |
| `segments` | `str` | No | JSON-encoded array of segments to add imported contacts to. |
| `status` | `str` | No | Current status of the contact import. |
| `topics` | `str` | No | JSON-encoded array of topic subscriptions to apply to imported contacts. |

### Operations

#### `create(reqdata, ctrl=None) -> ContactImportResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ContactImportResponseSuccess().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

#### `list(reqmatch=None, ctrl=None) -> list[ContactImportResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ContactImportResponseSuccess().list()
for contact_import_response_success in results:
    print(contact_import_response_success.data_get())
```

#### `load(reqmatch, ctrl=None) -> ContactImportResponseSuccessEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.ContactImportResponseSuccess().load({"id": "contact_import_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactImportResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactPropertyEntity

```python
contact_property = client.ContactProperty()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the contact property was created. |
| `fallback_value` | `Any` | No | The default value when the property is not set for a contact. |
| `id` | `str` | No | The ID of the contact property. |
| `key` | `str` | No | The property key. |
| `object` | `str` | No | The object type. |
| `type` | `str` | No | The property type. |

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

#### `create(reqdata, ctrl=None) -> ContactPropertyEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ContactProperty().create({
    "key": "example_key",  # str
    "type": "example_type",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[ContactPropertyEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ContactProperty().list()
for contact_property in results:
    print(contact_property.data_get())
```

#### `load(reqmatch, ctrl=None) -> ContactPropertyEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.ContactProperty().load({"id": "contact_property_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactPropertyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactTopicsResponseSuccessEntity

```python
contact_topics_response_success = client.ContactTopicsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ContactTopicsResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ContactTopicsResponseSuccess().list({"id": "example"})
for contact_topics_response_success in results:
    print(contact_topics_response_success.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactTopicsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CreateApiKeyEntity

```python
create_api_key = client.CreateApiKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `domain_id` | `str` | No | Restrict an API key to send emails only from a specific domain. |
| `name` | `str` | Yes | The API key name. |
| `permission` | `str` | No | The API key can have full access to Resend’s API or be only restricted to send emails. |

### Operations

#### `create(reqdata, ctrl=None) -> CreateApiKeyEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.CreateApiKey().create({
    "name": "example_name",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CreateApiKeyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CreateBatchEmailEntity

```python
create_batch_email = client.CreateBatchEmail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | No |  |
| `send_emails` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> CreateBatchEmailEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.CreateBatchEmail().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CreateBatchEmailEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CreateBroadcastResponseSuccessEntity

```python
create_broadcast_response_success = client.CreateBroadcastResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `str` | No | Use `segment_id` instead. |
| `from` | `str` | Yes | The email address of the sender. |
| `html` | `str` | No | The HTML version of the message. |
| `name` | `str` | No | Name of the broadcast. |
| `preview_text` | `str` | No | The preview text of the email. |
| `reply_to` | `list` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `str` | No | Schedule time to send the broadcast. |
| `segment_id` | `str` | Yes | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `bool` | No | Whether to send the broadcast immediately or keep it as a draft. |
| `subject` | `str` | Yes | The subject line of the email. |
| `text` | `str` | No | The plain text version of the message. |
| `topic_id` | `str` | No | The topic ID that the broadcast will be scoped to. |

### Operations

#### `create(reqdata, ctrl=None) -> CreateBroadcastResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.CreateBroadcastResponseSuccess().create({
    "from": "example_from",  # str
    "segment_id": "example_segment_id",  # str
    "subject": "example_subject",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CreateBroadcastResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CreateInboxDraftResponseSuccessEntity

```python
create_inbox_draft_response_success = client.CreateInboxDraftResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list | None` | No | Files to store on the draft. |
| `bcc` | `Any` | No | The BCC recipients. |
| `cc` | `Any` | No | The CC recipients. |
| `html` | `str | None` | No | The HTML body. |
| `id` | `str` | Yes | The ID of the draft. |
| `object` | `str` | Yes | Type of the response object. |
| `reply_to_email_id` | `str` | No | The Email ID to reply to, as returned in `data[].id` by the list thread emails endpoint. |
| `subject` | `str | None` | No | The subject. |
| `text` | `str | None` | No | The plain text body. |
| `thread_id` | `str` | No | The Thread ID to reply in. |
| `to` | `Any` | No | The recipients. |

### Operations

#### `create(reqdata, ctrl=None) -> CreateInboxDraftResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.CreateInboxDraftResponseSuccess().create({
    "inbox_id": "example_inbox_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CreateInboxDraftResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CreateInboxLabelResponseSuccessEntity

```python
create_inbox_label_response_success = client.CreateInboxLabelResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `str` | No | The color of the label, as a `#RRGGBB` hex color. |
| `name` | `str` | Yes | The name of the label. |

### Operations

#### `create(reqdata, ctrl=None) -> CreateInboxLabelResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.CreateInboxLabelResponseSuccess().create({
    "inbox_id": "example_inbox_id",  # str
    "name": "example_name",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CreateInboxLabelResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DomainEntity

```python
domain = client.Domain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capabilities` | `dict` | No | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | No | Whether click tracking is enabled for this domain. |
| `created_at` | `str` | No | The date and time the domain was created. |
| `custom_return_path` | `str` | No | For advanced use cases, choose a subdomain for the Return-Path address. |
| `id` | `str` | No | The ID of the domain. |
| `name` | `str` | No | The name of the domain. |
| `object` | `str` | No | The type of object. |
| `open_tracking` | `bool` | No | Whether open tracking is enabled for this domain. |
| `records` | `list` | No |  |
| `region` | `str` | No | The region where the domain is hosted. |
| `status` | `str` | No | The status of the domain. |
| `tls` | `str` | No | TLS mode. |
| `tracking_subdomain` | `str` | No | The subdomain used for click and open tracking. |

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

#### `create(reqdata, ctrl=None) -> DomainEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Domain().create({
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[DomainEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Domain().list()
for domain in results:
    print(domain.data_get())
```

#### `load(reqmatch, ctrl=None) -> DomainEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Domain().load({"id": "domain_id"})
```

#### `remove(reqmatch, ctrl=None) -> DomainEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Domain().remove({"id": "domain_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DomainEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DomainClaimEntity

```python
domain_claim = client.DomainClaim()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `blocked_reason` | `str | None` | No | Why the claim is currently blocked, if applicable. |
| `click_tracking` | `bool` | No | Track clicks within the body of each HTML email. |
| `created_at` | `str` | No | The date and time the claim was created. |
| `custom_return_path` | `str` | No | For advanced use cases, choose a subdomain for the Return-Path address. |
| `domain_id` | `str | None` | No | The ID of the placeholder domain created for the claim. |
| `expires_at` | `str` | No | The date and time the claim expires if not verified. |
| `failure_reason` | `str | None` | No | Why the claim failed, if applicable. |
| `id` | `str` | No | The ID of the claim. |
| `name` | `str` | No | The name of the domain being claimed. |
| `object` | `str` | No | The type of object. |
| `open_tracking` | `bool` | No | Track the open rate of each email. |
| `record` | `dict` | No | The TXT record to add to your DNS to prove ownership of the claimed domain. |
| `region` | `str | None` | No | The region where the claimed domain will send from. |
| `status` | `str` | No | The status of the claim. |
| `tracking_subdomain` | `str` | No | The subdomain to use for click and open tracking. |

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

#### `create(reqdata, ctrl=None) -> DomainClaimEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.DomainClaim().create({
    "name": "example_name",  # str
})
```

#### `load(reqmatch, ctrl=None) -> DomainClaimEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.DomainClaim().load({"id": "domain_claim_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DomainClaimEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailEntity

```python
email = client.Email()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | No |  |
| `bcc` | `list` | No | The email addresses of the blind carbon copy recipients. |
| `cc` | `list` | No | The email addresses of the carbon copy recipients. |
| `created_at` | `str` | No | The date and time the email was created. |
| `from` | `str` | No | The email address of the sender. |
| `headers` | `dict` | No | Custom headers to add to the email. |
| `html` | `str` | No | The HTML body of the email. |
| `id` | `str` | No | The ID of the email. |
| `last_event` | `str` | No | The status of the email. |
| `message_id` | `str` | No | The Message-ID header value of the email. |
| `object` | `str` | No | The type of object. |
| `reply_to` | `list` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `str` | No | Schedule email to be sent later. |
| `subject` | `str` | No | The subject line of the email. |
| `tags` | `list` | No |  |
| `template` | `Any` | No |  |
| `text` | `str` | No | The plain text body of the email. |
| `to` | `list` | No | Recipient email address. |
| `topic_id` | `str` | No | The topic ID to scope the email to. |

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

#### `create(reqdata, ctrl=None) -> EmailEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Email().create({
    "from": "example_from",  # str
    "subject": "example_subject",  # str
    "to": [],  # list
})
```

#### `list(reqmatch=None, ctrl=None) -> list[EmailEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Email().list()
for email in results:
    print(email.data_get())
```

#### `load(reqmatch, ctrl=None) -> EmailEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Email().load({"id": "email_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailsMetricEntity

```python
emails_metric = client.EmailsMetric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `broadcast_id` | `str` | No | Present when `broadcast` is in `dimensions`. |
| `broadcast_name` | `str` | No | Present when `broadcast` is in `dimensions`. |
| `domain_id` | `str` | No | Present when `domain` is in `dimensions`. |
| `domain_name` | `str` | No | Present when `domain` is in `dimensions`. |
| `email_id` | `str` | No | Present when `email` is in `dimensions`. |
| `period` | `str` | No | Present when `period` is in `dimensions`. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[EmailsMetricEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.EmailsMetric().list()
for emails_metric in results:
    print(emails_metric.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EmailsMetricEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventEntity

```python
event = client.Event()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | The date and time the event was created. |
| `id` | `str` | No | The event ID. |
| `name` | `str` | No | The event name. |
| `object` | `str` | No | Type of the response object. |
| `schema` | `dict | None` | No | A flat key/type map defining the event payload schema. |
| `updated_at` | `str | None` | No | The date and time the event was last updated. |

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

#### `create(reqdata, ctrl=None) -> EventEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Event().create({
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[EventEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Event().list()
for event in results:
    print(event.data_get())
```

#### `load(reqmatch, ctrl=None) -> EventEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Event().load({"id": "event_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxEntity

```python
inbox = client.Inbox()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes | ISO 8601 timestamp when the inbox was created. |
| `domain_id` | `str` | Yes | The ID of the domain that owns the address. |
| `drafts` | `int` | Yes | The number of unsent drafts. |
| `email_address` | `str` | Yes | The address of the inbox. |
| `forwarding` | `bool` | No | When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record. |
| `from_name` | `str | None` | Yes | The name recipients see when mail is sent from this inbox. |
| `id` | `str` | Yes | The ID of the inbox. |
| `last_received` | `str | None` | Yes | ISO 8601 timestamp of the newest received email in the Inbox folder. |
| `name` | `str` | Yes | Internal name for the inbox. |
| `object` | `str` | Yes | The object type. |
| `receiving_address` | `str | None` | Yes | The address to forward mail to when forwarding is enabled. |
| `unread` | `int` | Yes | The number of unread threads in the Inbox folder. |
| `updated_at` | `str` | Yes | ISO 8601 timestamp when the inbox was last updated. |

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

#### `create(reqdata, ctrl=None) -> InboxEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Inbox().create({
    "email_address": "example_email_address",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[InboxEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Inbox().list()
for inbox in results:
    print(inbox.data_get())
```

#### `load(reqmatch, ctrl=None) -> InboxEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Inbox().load({"id": "inbox_id"})
```

#### `update(reqdata, ctrl=None) -> InboxEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Inbox().update({
    "id": "inbox_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxAgentResponseSuccessEntity

```python
inbox_agent_response_success = client.InboxAgentResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[InboxAgentResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxAgentResponseSuccess().list({"id": "example"})
for inbox_agent_response_success in results:
    print(inbox_agent_response_success.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxAgentResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxDraftListItemEntity

```python
inbox_draft_list_item = client.InboxDraftListItem()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | Yes | The files stored on the draft. |
| `bcc` | `list` | Yes | The BCC recipients. |
| `cc` | `list` | Yes | The CC recipients. |
| `id` | `str` | Yes | The ID of the draft. |
| `reply_to_email_id` | `str | None` | Yes | The Email ID the draft replies to. |
| `snippet` | `str | None` | Yes | A short preview of the body. |
| `subject` | `str | None` | Yes | The subject. |
| `thread_id` | `str | None` | Yes | The Thread ID when the draft is a reply. |
| `to` | `list | None` | Yes | The recipients. |
| `type` | `str` | Yes | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `str` | Yes | Timestamp when the draft was last saved. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[InboxDraftListItemEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxDraftListItem().list({"id": "example"})
for inbox_draft_list_item in results:
    print(inbox_draft_list_item.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxDraftListItemEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxDraftResponseSuccessEntity

```python
inbox_draft_response_success = client.InboxDraftResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | Yes | The files stored on the draft. |
| `bcc` | `list` | Yes | The BCC recipients. |
| `cc` | `list` | Yes | The CC recipients. |
| `created_at` | `str` | Yes | Timestamp when the draft was created. |
| `email_id` | `str | None` | Yes | The ID of the outbound email queued when the draft was sent. |
| `html` | `str | None` | Yes | The HTML body. |
| `id` | `str` | Yes | The ID of the draft. |
| `object` | `str` | Yes | The object type. |
| `reply_to_email_id` | `str | None` | Yes | The Email ID the draft replies to. |
| `subject` | `str | None` | Yes | The subject. |
| `text` | `str | None` | Yes | The plain text body. |
| `thread_id` | `str | None` | Yes | The Thread ID when the draft is a reply. |
| `to` | `list | None` | Yes | The recipients. |
| `type` | `str` | Yes | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `str` | Yes | Timestamp when the draft was last saved. |

### Operations

#### `create(reqdata, ctrl=None) -> InboxDraftResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.InboxDraftResponseSuccess().create({
    "draft_id": "example_draft_id",  # str
    "inbox_id": "example_inbox_id",  # str
})
```

#### `load(reqmatch, ctrl=None) -> InboxDraftResponseSuccessEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxDraftResponseSuccess().load({"draft_id": "draft_id", "inbox_id": "inbox_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxDraftResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxLabelEntity

```python
inbox_label = client.InboxLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `str` | Yes | The color of the label, as a `#RRGGBB` hex color in uppercase. |
| `created_at` | `str` | Yes | Timestamp when the label was created. |
| `id` | `str` | Yes | The ID of the label. |
| `name` | `str` | Yes | The name of the label. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[InboxLabelEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxLabel().list({"id": "example"})
for inbox_label in results:
    print(inbox_label.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxLabelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxThreadEntity

```python
inbox_thread = client.InboxThread()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bcc` | `list` | Yes | The BCC recipients of the most recent email. |
| `cc` | `list` | Yes | The CC recipients of the most recent email. |
| `email_count` | `int` | Yes | The number of emails in the thread. |
| `folder` | `str` | Yes | The folder the thread is in. |
| `from` | `str | None` | Yes | The sender email address. |
| `has_attachment` | `bool` | Yes | Whether the most recent email in the thread has an attachment. |
| `has_draft` | `bool` | Yes | Whether the thread has an unsent draft. |
| `id` | `str` | Yes | The ID of the thread. |
| `labels` | `list` | Yes | The labels attached to the thread. |
| `matched_email_id` | `str | None` | Yes | The ID of the email in the thread that matched the search. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `received_at` | `str` | Yes | Timestamp when the thread was last active. |
| `subject` | `str | None` | Yes | The subject of the thread. |
| `to` | `list` | Yes | The recipients of the most recent email. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[InboxThreadEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxThread().list({"id": "example"})
for inbox_thread in results:
    print(inbox_thread.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxThreadEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxThreadEmailEntity

```python
inbox_thread_email = client.InboxThreadEmail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | Yes | The attachments of the message. |
| `bcc` | `list` | Yes | The BCC recipients of the message. |
| `cc` | `list` | Yes | The CC recipients of the message. |
| `direction` | `str` | Yes | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `str` | Yes | Sender email address. |
| `html` | `str | None` | Yes | The HTML body. |
| `id` | `str` | Yes | The ID of the message. |
| `message_id` | `str | None` | Yes | The Message-ID header of the message. |
| `read` | `bool` | Yes | Whether the message has been read. |
| `received_at` | `str` | Yes | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `list` | Yes | The Reply-To addresses of the message. |
| `subject` | `str | None` | Yes | The subject of the message. |
| `text` | `str | None` | Yes | The plain-text body. |
| `to` | `list` | Yes | The recipients of the message. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[InboxThreadEmailEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxThreadEmail().list({"inbox_id": "example", "thread_id": "example"})
for inbox_thread_email in results:
    print(inbox_thread_email.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxThreadEmailEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxThreadEmailResponseSuccessEntity

```python
inbox_thread_email_response_success = client.InboxThreadEmailResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | Yes | The attachments of the message. |
| `bcc` | `list` | Yes | The BCC recipients of the message. |
| `cc` | `list` | Yes | The CC recipients of the message. |
| `direction` | `str` | Yes | Whether the inbox received the message (`inbound`) or sent it (`outbound`). |
| `from` | `str` | Yes | Sender email address. |
| `html` | `str | None` | Yes | The HTML body. |
| `id` | `str` | Yes | The ID of the message. |
| `message_id` | `str | None` | Yes | The Message-ID header of the message. |
| `read` | `bool` | Yes | Whether the message has been read. |
| `received_at` | `str` | Yes | Timestamp when the inbox received the message, or when the inbox sent it. |
| `reply_to` | `list` | Yes | The Reply-To addresses of the message. |
| `subject` | `str | None` | Yes | The subject of the message. |
| `text` | `str | None` | Yes | The plain-text body. |
| `to` | `list` | Yes | The recipients of the message. |

### Operations

#### `create(reqdata, ctrl=None) -> InboxThreadEmailResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

#### `load(reqmatch, ctrl=None) -> InboxThreadEmailResponseSuccessEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxThreadEmailResponseSuccess().load({"email_id": "email_id", "inbox_id": "inbox_id", "thread_id": "thread_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxThreadEmailResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxThreadResponseSuccessEntity

```python
inbox_thread_response_success = client.InboxThreadResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `folder` | `str` | Yes | The folder the thread is in. |
| `id` | `str` | Yes | The ID of the thread. |
| `labels` | `list` | Yes | The labels attached to the thread. |
| `object` | `str` | Yes | The object type. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `subject` | `str | None` | Yes | The subject of the thread. |

### Operations

#### `load(reqmatch, ctrl=None) -> InboxThreadResponseSuccessEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxThreadResponseSuccess().load({"inbox_id": "inbox_id", "thread_id": "thread_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InboxThreadResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListAttachmentEntity

```python
list_attachment = client.ListAttachment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_disposition` | `str | None` | No | How the attachment should be displayed. |
| `content_id` | `str` | No | The content ID for inline attachments. |
| `content_type` | `str` | No | The MIME type of the attachment. |
| `download_url` | `str` | No | Signed URL to download the attachment content. |
| `expires_at` | `str` | No | Timestamp when the download URL expires. |
| `filename` | `str | None` | No | The filename of the attachment. |
| `id` | `str` | No | The ID of the attachment. |
| `size` | `int` | No | Size of the attachment in bytes. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListAttachmentEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListAttachment().list({"email_id": "example"})
for list_attachment in results:
    print(list_attachment.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListAttachmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListBroadcastClickedLinksResponseSuccessEntity

```python
list_broadcast_clicked_links_response_success = client.ListBroadcastClickedLinksResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clicks` | `int` | No | Total number of clicks on this URL. |
| `id` | `str` | No | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `int` | No | Number of unique clicks on this URL. |
| `url` | `str` | No | The URL that was clicked. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListBroadcastClickedLinksResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListBroadcastClickedLinksResponseSuccess().list({"broadcast_id": "example"})
for list_broadcast_clicked_links_response_success in results:
    print(list_broadcast_clicked_links_response_success.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListBroadcastClickedLinksResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListBroadcastRecipientsResponseSuccessEntity

```python
list_broadcast_recipients_response_success = client.ListBroadcastRecipientsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bounce_type` | `str` | No | The type of bounce. |
| `clicked_links` | `list` | No | The links this recipient clicked. |
| `contact_id` | `str | None` | No | The ID of the contact associated with this recipient, if one exists. |
| `count` | `int` | No | The number of times this recipient triggered the event. |
| `email` | `str` | No | The recipient's email address. |
| `id` | `str` | No | Opaque cursor identifying this row, used for pagination. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListBroadcastRecipientsResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListBroadcastRecipientsResponseSuccess().list({"broadcast_id": "example", "type": "example"})
for list_broadcast_recipients_response_success in results:
    print(list_broadcast_recipients_response_success.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListBroadcastRecipientsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListContactSegmentsResponseSuccessEntity

```python
list_contact_segments_response_success = client.ListContactSegmentsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the contact was added to the segment. |
| `id` | `str` | No | Unique identifier for the segment. |
| `name` | `str` | No | Name of the segment. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListContactSegmentsResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListContactSegmentsResponseSuccess().list({"contact_id": "example"})
for list_contact_segments_response_success in results:
    print(list_contact_segments_response_success.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListContactSegmentsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListContactsResponseSuccessEntity

```python
list_contacts_response_success = client.ListContactsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the contact was created. |
| `email` | `str` | No | Email address of the contact. |
| `first_name` | `str | None` | No | First name of the contact. |
| `id` | `str` | No | Unique identifier for the contact. |
| `last_name` | `str | None` | No | Last name of the contact. |
| `unsubscribed` | `bool` | No | Indicates if the contact is unsubscribed. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListContactsResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListContactsResponseSuccess().list({"segment_id": "example"})
for list_contacts_response_success in results:
    print(list_contacts_response_success.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListContactsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListWebhookEventEntity

```python
list_webhook_event = client.ListWebhookEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the event was created. |
| `id` | `str` | No | The ID of the webhook event. |
| `status` | `str` | No | The delivery status of the event for this webhook. |
| `type` | `str` | No | The type of the event. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListWebhookEventEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListWebhookEvent().list({"webhook_id": "example"})
for list_webhook_event in results:
    print(list_webhook_event.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListWebhookEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListWebhookEventAttemptEntity

```python
list_webhook_event_attempt = client.ListWebhookEventAttempt()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `http_status_code` | `int` | No | The HTTP status code returned by the webhook endpoint. |
| `id` | `str` | No | The ID of the webhook event attempt. |
| `response` | `str` | No | The response body returned by the webhook endpoint. |
| `sent_at` | `str` | No | Timestamp indicating when the attempt was sent. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ListWebhookEventAttemptEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ListWebhookEventAttempt().list({"event_id": "example", "webhook_id": "example"})
for list_webhook_event_attempt in results:
    print(list_webhook_event_attempt.data_get())
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListWebhookEventAttemptEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## LogEntity

```python
log = client.Log()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | The date the log was created. |
| `endpoint` | `str` | No | The API endpoint that was called. |
| `id` | `str` | No | The log ID. |
| `method` | `str` | No | The HTTP method used. |
| `object` | `str` | No | Type of the response object. |
| `request_body` | `dict | None` | No | The request body sent to the API. |
| `response_body` | `dict | None` | No | The response body returned by the API. |
| `response_status` | `int` | No | The HTTP status code of the response. |
| `user_agent` | `str | None` | No | The user agent of the request. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[LogEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Log().list()
for log in results:
    print(log.data_get())
```

#### `load(reqmatch, ctrl=None) -> LogEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Log().load({"id": "log_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `LogEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReceivedEmailEntity

```python
received_email = client.ReceivedEmail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list` | No | Array of attachments. |
| `bcc` | `list | None` | No | The BCC recipients. |
| `cc` | `list | None` | No | The CC recipients. |
| `created_at` | `str` | No | Timestamp when the email was received. |
| `from` | `str` | No | The sender email address. |
| `headers` | `dict | None` | No | The email headers. |
| `html` | `str | None` | No | The HTML content of the email. |
| `id` | `str` | No | The ID of the received email. |
| `message_id` | `str` | No | The unique message ID from the email headers. |
| `object` | `str` | No | The type of object. |
| `received_for` | `list` | No | The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers. |
| `reply_to` | `list | None` | No | The reply-to addresses. |
| `subject` | `str` | No | The email subject. |
| `text` | `str | None` | No | The plain text content of the email. |
| `to` | `list` | No | The recipient email addresses. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ReceivedEmailEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ReceivedEmail().list()
for received_email in results:
    print(received_email.data_get())
```

#### `load(reqmatch, ctrl=None) -> ReceivedEmailEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.ReceivedEmail().load({"email_id": "email_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReceivedEmailEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveAudienceResponseSuccessEntity

```python
remove_audience_response_success = client.RemoveAudienceResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveAudienceResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveAudienceResponseSuccess().remove({"id": "remove_audience_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveAudienceResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveBroadcastResponseSuccessEntity

```python
remove_broadcast_response_success = client.RemoveBroadcastResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveBroadcastResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveBroadcastResponseSuccess().remove({"id": "remove_broadcast_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveBroadcastResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveContactFromSegmentResponseSuccessEntity

```python
remove_contact_from_segment_response_success = client.RemoveContactFromSegmentResponseSuccess()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveContactFromSegmentResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveContactFromSegmentResponseSuccess().remove({"contact_id": "contact_id", "segment_id": "segment_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveContactFromSegmentResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveContactPropertyResponseSuccessEntity

```python
remove_contact_property_response_success = client.RemoveContactPropertyResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveContactPropertyResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveContactPropertyResponseSuccess().remove({"id": "remove_contact_property_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveContactPropertyResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveContactResponseSuccessEntity

```python
remove_contact_response_success = client.RemoveContactResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveContactResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveContactResponseSuccess().remove({"id": "remove_contact_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveContactResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveEventEntity

```python
remove_event = client.RemoveEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveEventEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveEvent().remove({"id": "remove_event_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveInboxDraftResponseSuccessEntity

```python
remove_inbox_draft_response_success = client.RemoveInboxDraftResponseSuccess()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveInboxDraftResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveInboxDraftResponseSuccess().remove({"draft_id": "draft_id", "inbox_id": "inbox_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveInboxDraftResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveInboxLabelResponseSuccessEntity

```python
remove_inbox_label_response_success = client.RemoveInboxLabelResponseSuccess()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveInboxLabelResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveInboxLabelResponseSuccess().remove({"inbox_id": "inbox_id", "label_id": "label_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveInboxLabelResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveInboxResponseSuccessEntity

```python
remove_inbox_response_success = client.RemoveInboxResponseSuccess()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveInboxResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveInboxResponseSuccess().remove({"inbox_id": "inbox_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveInboxResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveInboxThreadLabelResponseSuccessEntity

```python
remove_inbox_thread_label_response_success = client.RemoveInboxThreadLabelResponseSuccess()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveInboxThreadLabelResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveInboxThreadLabelResponseSuccess().remove({"inbox_id": "inbox_id", "label_id": "label_id", "thread_id": "thread_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveInboxThreadLabelResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveInboxThreadResponseSuccessEntity

```python
remove_inbox_thread_response_success = client.RemoveInboxThreadResponseSuccess()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> RemoveInboxThreadResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveInboxThreadResponseSuccess().remove({"inbox_id": "inbox_id", "thread_id": "thread_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveInboxThreadResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveSegmentResponseSuccessEntity

```python
remove_segment_response_success = client.RemoveSegmentResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `str` | No | The ID of the audience this segment belongs to. |
| `created_at` | `str` | No | Timestamp indicating when the segment was created. |
| `filter` | `dict` | No | Filter conditions for the segment. |
| `id` | `str` | No | Unique identifier for the segment. |
| `name` | `str` | Yes | The name of the segment. |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `audience_id` | - | - | - |
| `created_at` | - | - | - |
| `filter` | - | - | - |
| `id` | - | - | - |
| `name` | Yes | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> RemoveSegmentResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.RemoveSegmentResponseSuccess().create({
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[RemoveSegmentResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.RemoveSegmentResponseSuccess().list()
for remove_segment_response_success in results:
    print(remove_segment_response_success.data_get())
```

#### `remove(reqmatch, ctrl=None) -> RemoveSegmentResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveSegmentResponseSuccess().remove({"id": "remove_segment_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveSegmentResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveSuppressionResponseSuccessEntity

```python
remove_suppression_response_success = client.RemoveSuppressionResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the suppression was created. |
| `email` | `str` | Yes | Email address to suppress. |
| `id` | `str` | No | Unique identifier for the suppression. |
| `origin` | `str` | No | Origin of the suppression. |
| `source_id` | `str | None` | No | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `created_at` | - | - | - |
| `email` | Yes | - | - |
| `id` | - | - | - |
| `origin` | - | - | - |
| `source_id` | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> RemoveSuppressionResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.RemoveSuppressionResponseSuccess().create({
    "email": "example_email",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[RemoveSuppressionResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.RemoveSuppressionResponseSuccess().list()
for remove_suppression_response_success in results:
    print(remove_suppression_response_success.data_get())
```

#### `remove(reqmatch, ctrl=None) -> RemoveSuppressionResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveSuppressionResponseSuccess().remove({"suppression": "suppression"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveSuppressionResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveTemplateResponseSuccessEntity

```python
remove_template_response_success = client.RemoveTemplateResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `str` | No | The alias of the template. |
| `created_at` | `str` | No | Timestamp indicating when the template was created. |
| `from` | `str` | No | Sender email address. |
| `html` | `str` | Yes | The HTML version of the template. |
| `id` | `str` | No | The ID of the template. |
| `name` | `str` | Yes | The name of the template. |
| `published_at` | `str | None` | No | Timestamp indicating when the template was published. |
| `reply_to` | `list` | No | Reply-to email addresses. |
| `status` | `str` | No | The publication status of the template. |
| `subject` | `str` | No | Email subject. |
| `text` | `str` | No | The plain text version of the template. |
| `updated_at` | `str` | No | Timestamp indicating when the template was last updated. |
| `variables` | `list` | No |  |

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

#### `create(reqdata, ctrl=None) -> RemoveTemplateResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.RemoveTemplateResponseSuccess().create({
    "html": "example_html",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[RemoveTemplateResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.RemoveTemplateResponseSuccess().list()
for remove_template_response_success in results:
    print(remove_template_response_success.data_get())
```

#### `remove(reqmatch, ctrl=None) -> RemoveTemplateResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveTemplateResponseSuccess().remove({"id": "remove_template_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveTemplateResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RemoveTopicResponseSuccessEntity

```python
remove_topic_response_success = client.RemoveTopicResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the topic was created. |
| `default_subscription` | `str` | Yes | The default subscription status for the topic. |
| `description` | `str` | No | A description of the topic. |
| `id` | `str` | No | Unique identifier for the topic. |
| `name` | `str` | Yes | The name of the topic. |
| `visibility` | `str` | No | The visibility of the topic. |

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

#### `create(reqdata, ctrl=None) -> RemoveTopicResponseSuccessEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.RemoveTopicResponseSuccess().create({
    "default_subscription": "example_default_subscription",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[RemoveTopicResponseSuccessEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.RemoveTopicResponseSuccess().list()
for remove_topic_response_success in results:
    print(remove_topic_response_success.data_get())
```

#### `remove(reqmatch, ctrl=None) -> RemoveTopicResponseSuccessEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RemoveTopicResponseSuccess().remove({"id": "remove_topic_response_success_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RemoveTopicResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RetrievedAttachmentEntity

```python
retrieved_attachment = client.RetrievedAttachment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_disposition` | `str | None` | No | How the attachment should be displayed. |
| `content_id` | `str` | No | The content ID for inline attachments. |
| `content_type` | `str` | No | The MIME type of the attachment. |
| `download_url` | `str` | No | Signed URL to download the attachment content. |
| `expires_at` | `str` | No | Timestamp when the download URL expires. |
| `filename` | `str | None` | No | The filename of the attachment. |
| `id` | `str` | No | The ID of the attachment. |
| `object` | `str` | No | The type of object. |
| `size` | `int` | No | Size of the attachment in bytes. |

### Operations

#### `load(reqmatch, ctrl=None) -> RetrievedAttachmentEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.RetrievedAttachment().load({"id": "retrieved_attachment_id", "email_id": "email_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RetrievedAttachmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RevokeOAuthGrantEntity

```python
revoke_o_auth_grant = client.RevokeOAuthGrant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client` | `dict` | No | The OAuth client the grant was issued to. |
| `client_id` | `str` | No | The ID of the OAuth client the grant was issued to. |
| `created_at` | `str` | No | The date and time the OAuth grant was created. |
| `id` | `str` | No | The ID of the OAuth grant. |
| `revoked_at` | `str | None` | No | The date and time the OAuth grant was revoked, or null if it is still active. |
| `revoked_reason` | `str | None` | No | The reason the OAuth grant was revoked, or null if it is still active. |
| `scopes` | `list` | No | The scopes granted to the OAuth client. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[RevokeOAuthGrantEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.RevokeOAuthGrant().list()
for revoke_o_auth_grant in results:
    print(revoke_o_auth_grant.data_get())
```

#### `remove(reqmatch, ctrl=None) -> RevokeOAuthGrantEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.RevokeOAuthGrant().remove({"id": "revoke_o_auth_grant_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RevokeOAuthGrantEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RotateEntity

```python
rotate = client.Rotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No | The ID of the webhook. |
| `object` | `str` | No | The type of object. |
| `signing_secret` | `str` | No | The new secret key used to verify webhook payloads. |

### Operations

#### `create(reqdata, ctrl=None) -> RotateEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Rotate().create({
    "webhook_id": "example_webhook_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RotateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SegmentEntity

```python
segment = client.Segment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `str` | No | The ID of the audience this segment belongs to. |
| `created_at` | `str` | No | Timestamp indicating when the segment was created. |
| `filter` | `dict` | No | Filter conditions for the segment. |
| `id` | `str` | No | The ID of the segment. |
| `name` | `str` | No | The name of the segment. |
| `object` | `str` | No | The object type. |

### Operations

#### `load(reqmatch, ctrl=None) -> SegmentEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Segment().load({"id": "segment_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SegmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SetInboxThreadLabelsResponseSuccessEntity

```python
set_inbox_thread_labels_response_success = client.SetInboxThreadLabelsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | Yes | Every label on the thread after the change. |
| `labels` | `list` | Yes | The IDs of the labels the thread carries after the change. |
| `object` | `str` | Yes | Type of the response object. |

### Operations

#### `update(reqdata, ctrl=None) -> SetInboxThreadLabelsResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.SetInboxThreadLabelsResponseSuccess().update({
    "inbox_id": "inbox_id",
    "thread_id": "thread_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SetInboxThreadLabelsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SuppressionEntity

```python
suppression = client.Suppression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the suppression was created. |
| `email` | `str` | No | Email address that is suppressed. |
| `id` | `str` | No | Unique identifier for the suppression. |
| `object` | `str` | No | Type of the response object. |
| `origin` | `str` | No | Origin of the suppression. |
| `source_id` | `str | None` | No | Identifier of the event that caused the suppression, such as the email that bounced or complained. |

### Operations

#### `load(reqmatch, ctrl=None) -> SuppressionEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Suppression().load({"id": "suppression_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SuppressionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TemplateEntity

```python
template = client.Template()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `str` | No | The alias of the template. |
| `created_at` | `str` | No | Timestamp indicating when the template was created. |
| `current_version_id` | `str` | No | The ID of the current version of the template. |
| `from` | `str` | No | Sender email address. |
| `has_unpublished_versions` | `bool` | No | Indicates whether the template has unpublished versions. |
| `html` | `str` | No | The HTML version of the template. |
| `id` | `str` | No | The ID of the template. |
| `name` | `str` | No | The name of the template. |
| `object` | `str` | No | The type of object. |
| `published_at` | `str | None` | No | Timestamp indicating when the template was published. |
| `reply_to` | `list | None` | No | Reply-to email addresses. |
| `status` | `str` | No | The publication status of the template. |
| `subject` | `str` | No | Email subject. |
| `text` | `str` | No | The plain text version of the template. |
| `updated_at` | `str` | No | Timestamp indicating when the template was last updated. |
| `variables` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> TemplateEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

#### `load(reqmatch, ctrl=None) -> TemplateEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Template().load({"id": "template_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TemplateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TopicEntity

```python
topic = client.Topic()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the topic was created. |
| `default_subscription` | `str` | No | The default subscription status for the topic. |
| `description` | `str` | No | A description of the topic. |
| `id` | `str` | No | The ID of the topic. |
| `name` | `str` | No | The name of the topic. |
| `object` | `str` | No | The object type. |
| `visibility` | `str` | No | The visibility of the topic. |

### Operations

#### `load(reqmatch, ctrl=None) -> TopicEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Topic().load({"id": "topic_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TopicEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateApiKeyEntity

```python
update_api_key = client.UpdateApiKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No | The ID of the API key. |
| `name` | `str` | Yes | The API key name. |
| `object` | `str` | No | The type of object. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateApiKeyEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateApiKey().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateApiKeyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateBroadcastResponseSuccessEntity

```python
update_broadcast_response_success = client.UpdateBroadcastResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `audience_id` | `str` | No | Use `segment_id` instead. |
| `from` | `str` | No | The email address of the sender. |
| `html` | `str` | No | The HTML version of the message. |
| `id` | `str` | No | The ID of the broadcast. |
| `name` | `str` | No | Name of the broadcast. |
| `object` | `str` | No | The object type of the response. |
| `preview_text` | `str` | No | The preview text of the email. |
| `reply_to` | `list` | No | The email addresses to which replies should be sent. |
| `segment_id` | `str` | No | Unique identifier of the segment this broadcast will be sent to. |
| `subject` | `str` | No | The subject line of the email. |
| `text` | `str` | No | The plain text version of the message. |
| `topic_id` | `str` | No | The topic ID that the broadcast will be scoped to. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateBroadcastResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateBroadcastResponseSuccess().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateBroadcastResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateContactPropertyResponseSuccessEntity

```python
update_contact_property_response_success = client.UpdateContactPropertyResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fallback_value` | `Any` | No | The default value to use when the property is not set for a contact. |
| `id` | `str` | No | The ID of the contact property. |
| `object` | `str` | No | The object type. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateContactPropertyResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateContactPropertyResponseSuccess().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateContactPropertyResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateContactResponseSuccessEntity

```python
update_contact_response_success = client.UpdateContactResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `str` | No | Email address of the contact. |
| `first_name` | `str` | No | First name of the contact. |
| `id` | `str` | No | Unique identifier for the updated contact. |
| `last_name` | `str` | No | Last name of the contact. |
| `object` | `str` | No | Type of the response object. |
| `properties` | `dict` | No | A map of custom property keys and values to update. |
| `unsubscribed` | `bool` | No | The Contact's global subscription status. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateContactResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateContactResponseSuccess().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateContactResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateContactTopicsResponseSuccessEntity

```python
update_contact_topics_response_success = client.UpdateContactTopicsResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `str` | No | The ID of the contact. |
| `object` | `str` | No | The object type. |
| `topics` | `list` | No | Array of updated topic subscriptions. |

### Field Usage by Operation

| Field | update |
| --- | --- |
| `contact_id` | - |
| `object` | - |
| `topics` | Yes |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateContactTopicsResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateContactTopicsResponseSuccess().update({
    "contact_id": "contact_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateContactTopicsResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateDomainResponseSuccessEntity

```python
update_domain_response_success = client.UpdateDomainResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capabilities` | `dict` | No | Configure the domain capabilities for sending and receiving emails. |
| `click_tracking` | `bool` | No | Track clicks within the body of each HTML email. |
| `id` | `str` | No | The ID of the updated domain. |
| `object` | `str` | No | The object type representing the updated domain. |
| `open_tracking` | `bool` | No | Track the open rate of each email. |
| `tls` | `str` | No | enforced | opportunistic. |
| `tracking_subdomain` | `str` | No | The subdomain to use for click and open tracking. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateDomainResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateDomainResponseSuccess().update({
    "domain_id": "domain_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateDomainResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateEmailOptionEntity

```python
update_email_option = client.UpdateEmailOption()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `scheduled_at` | `str` | No | Schedule email to be sent later. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateEmailOptionEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateEmailOption().update({
    "email_id": "email_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateEmailOptionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateEventEntity

```python
update_event = client.UpdateEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No | The ID of the updated event. |
| `object` | `str` | No | Type of the response object. |
| `schema` | `dict | None` | Yes | A flat key/type map defining the event payload schema. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateEventEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateEvent().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateInboxDraftResponseSuccessEntity

```python
update_inbox_draft_response_success = client.UpdateInboxDraftResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `attachments` | `list | None` | No | Files to store on the draft. |
| `bcc` | `Any` | No | The BCC recipients. |
| `cc` | `Any` | No | The CC recipients. |
| `html` | `str | None` | No | The HTML body. |
| `id` | `str` | Yes | The ID of the draft. |
| `object` | `str` | Yes | The object type. |
| `subject` | `str | None` | No | The subject. |
| `text` | `str | None` | No | The plain text body. |
| `to` | `Any` | No | The recipients. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateInboxDraftResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateInboxDraftResponseSuccess().update({
    "draft_id": "draft_id",
    "inbox_id": "inbox_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateInboxDraftResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateInboxLabelResponseSuccessEntity

```python
update_inbox_label_response_success = client.UpdateInboxLabelResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `str` | No | A new color for the label, as a `#RRGGBB` hex color. |
| `id` | `str` | Yes | The ID of the updated label. |
| `name` | `str` | No | A new name for the label. |
| `object` | `str` | Yes | Type of the response object. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateInboxLabelResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateInboxLabelResponseSuccess().update({
    "inbox_id": "inbox_id",
    "label_id": "label_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateInboxLabelResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateInboxResponseSuccessEntity

```python
update_inbox_response_success = client.UpdateInboxResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from_name` | `str` | No | The name recipients see when mail is sent from this inbox. |
| `id` | `str` | Yes | The ID of the inbox. |
| `name` | `str` | No | New internal name for the inbox. |
| `object` | `str` | Yes | The object type. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateInboxResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateInboxResponseSuccess().update({
    "inbox_id": "inbox_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateInboxResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateInboxThreadResponseSuccessEntity

```python
update_inbox_thread_response_success = client.UpdateInboxThreadResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `folder` | `str` | Yes | The folder the thread is in. |
| `id` | `str` | Yes | The ID of the thread. |
| `labels` | `list` | Yes | The labels attached to the thread. |
| `object` | `str` | Yes | The object type. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `subject` | `str | None` | Yes | The subject of the thread. |

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

#### `update(reqdata, ctrl=None) -> UpdateInboxThreadResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateInboxThreadResponseSuccess().update({
    "inbox_id": "inbox_id",
    "thread_id": "thread_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateInboxThreadResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateSegmentResponseSuccessEntity

```python
update_segment_response_success = client.UpdateSegmentResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No | The ID of the segment. |
| `name` | `str` | Yes | The name of the segment. |
| `object` | `str` | No | The object type. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateSegmentResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateSegmentResponseSuccess().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateSegmentResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateTemplateResponseSuccessEntity

```python
update_template_response_success = client.UpdateTemplateResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `alias` | `str` | No | The alias of the template. |
| `from` | `str` | No | Sender email address. |
| `html` | `str` | No | The HTML version of the template. |
| `id` | `str` | No | The ID of the template. |
| `name` | `str` | No | The name of the template. |
| `object` | `str` | No | The object type of the response. |
| `reply_to` | `list` | No | Reply-to email addresses. |
| `subject` | `str` | No | Email subject. |
| `text` | `str` | No | The plain text version of the template. |
| `variables` | `list` | No |  |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateTemplateResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateTemplateResponseSuccess().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateTemplateResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateTopicResponseSuccessEntity

```python
update_topic_response_success = client.UpdateTopicResponseSuccess()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `str` | No | A description of the topic. |
| `id` | `str` | No | The ID of the topic. |
| `name` | `str` | No | The name of the topic. |
| `object` | `str` | No | The object type. |
| `visibility` | `str` | No | The visibility of the topic. |

### Operations

#### `update(reqdata, ctrl=None) -> UpdateTopicResponseSuccessEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateTopicResponseSuccess().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateTopicResponseSuccessEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UpdateWebhookEntity

```python
update_webhook = client.UpdateWebhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the webhook was created. |
| `endpoint` | `str` | Yes | The URL where webhook events will be sent. |
| `events` | `list` | Yes | Array of event types to subscribe to. |
| `id` | `str` | No | The ID of the updated webhook. |
| `object` | `str` | No | The type of object. |
| `status` | `str` | No | The status of the webhook. |

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

#### `create(reqdata, ctrl=None) -> UpdateWebhookEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.UpdateWebhook().create({
    "endpoint": "example_endpoint",  # str
    "events": [],  # list
})
```

#### `list(reqmatch=None, ctrl=None) -> list[UpdateWebhookEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.UpdateWebhook().list()
for update_webhook in results:
    print(update_webhook.data_get())
```

#### `update(reqdata, ctrl=None) -> UpdateWebhookEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.UpdateWebhook().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UpdateWebhookEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UsageEntity

```python
usage = client.Usage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ai_credits` | `dict` | No |  |
| `automation_runs` | `dict` | No |  |
| `broadcasts` | `dict` | No |  |
| `contacts` | `dict` | No |  |
| `domains` | `dict` | No |  |
| `emails` | `dict` | No |  |
| `object` | `str` | No | The type of object. |
| `rate_limit` | `dict` | No |  |
| `segments` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> UsageEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Usage().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UsageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WebhookEntity

```python
webhook = client.Webhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the webhook was created. |
| `endpoint` | `str` | No | The URL where webhook events are sent. |
| `events` | `list | None` | No | Array of event types subscribed to. |
| `id` | `str` | No | The ID of the webhook. |
| `object` | `str` | No | The type of object. |
| `signing_secret` | `str` | No | The secret key used to verify webhook payloads. |
| `status` | `str` | No | The status of the webhook. |

### Operations

#### `load(reqmatch, ctrl=None) -> WebhookEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Webhook().load({"id": "webhook_id"})
```

#### `remove(reqmatch, ctrl=None) -> WebhookEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Webhook().remove({"id": "webhook_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WebhookEventEntity

```python
webhook_event = client.WebhookEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | Timestamp indicating when the event was created. |
| `id` | `str` | No | The ID of the webhook event. |
| `next_attempt_at` | `str | None` | No | Timestamp of the next scheduled delivery attempt, or null when none is scheduled. |
| `object` | `str` | No | The type of object. |
| `payload` | `dict` | No | The event payload sent to the webhook endpoint. |
| `status` | `str` | No | The delivery status of the event for this webhook. |
| `type` | `str` | No | The type of the event. |

### Operations

#### `create(reqdata, ctrl=None) -> WebhookEventEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.WebhookEvent().create({
    "event_id": "example_event_id",  # str
    "webhook_id": "example_webhook_id",  # str
})
```

#### `load(reqmatch, ctrl=None) -> WebhookEventEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.WebhookEvent().load({"id": "webhook_event_id", "webhook_id": "webhook_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEventEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = ResendSDK({
    "feature": {
        "debug": {"active": True},
        "idempotency": {"active": True},
        "metrics": {"active": True},
        "paging": {"active": True},
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

