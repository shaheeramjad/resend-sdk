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

#### `ApiKey(data=None)`

Create a new `ApiKeyEntity` instance. Pass `None` for no initial data.

#### `Attachment(data=None)`

Create a new `AttachmentEntity` instance. Pass `None` for no initial data.

#### `Audience(data=None)`

Create a new `AudienceEntity` instance. Pass `None` for no initial data.

#### `Automation(data=None)`

Create a new `AutomationEntity` instance. Pass `None` for no initial data.

#### `AutomationRun(data=None)`

Create a new `AutomationRunEntity` instance. Pass `None` for no initial data.

#### `Broadcast(data=None)`

Create a new `BroadcastEntity` instance. Pass `None` for no initial data.

#### `BroadcastClickedLink(data=None)`

Create a new `BroadcastClickedLinkEntity` instance. Pass `None` for no initial data.

#### `BroadcastRecipient(data=None)`

Create a new `BroadcastRecipientEntity` instance. Pass `None` for no initial data.

#### `Contact(data=None)`

Create a new `ContactEntity` instance. Pass `None` for no initial data.

#### `ContactImport(data=None)`

Create a new `ContactImportEntity` instance. Pass `None` for no initial data.

#### `ContactProperty(data=None)`

Create a new `ContactPropertyEntity` instance. Pass `None` for no initial data.

#### `ContactSegment(data=None)`

Create a new `ContactSegmentEntity` instance. Pass `None` for no initial data.

#### `ContactTopic(data=None)`

Create a new `ContactTopicEntity` instance. Pass `None` for no initial data.

#### `Domain(data=None)`

Create a new `DomainEntity` instance. Pass `None` for no initial data.

#### `DomainClaim(data=None)`

Create a new `DomainClaimEntity` instance. Pass `None` for no initial data.

#### `Email(data=None)`

Create a new `EmailEntity` instance. Pass `None` for no initial data.

#### `EmailBatch(data=None)`

Create a new `EmailBatchEntity` instance. Pass `None` for no initial data.

#### `EmailsMetric(data=None)`

Create a new `EmailsMetricEntity` instance. Pass `None` for no initial data.

#### `Event(data=None)`

Create a new `EventEntity` instance. Pass `None` for no initial data.

#### `Inbox(data=None)`

Create a new `InboxEntity` instance. Pass `None` for no initial data.

#### `InboxAgent(data=None)`

Create a new `InboxAgentEntity` instance. Pass `None` for no initial data.

#### `InboxDraft(data=None)`

Create a new `InboxDraftEntity` instance. Pass `None` for no initial data.

#### `InboxLabel(data=None)`

Create a new `InboxLabelEntity` instance. Pass `None` for no initial data.

#### `InboxThread(data=None)`

Create a new `InboxThreadEntity` instance. Pass `None` for no initial data.

#### `InboxThreadEmail(data=None)`

Create a new `InboxThreadEmailEntity` instance. Pass `None` for no initial data.

#### `InboxThreadLabel(data=None)`

Create a new `InboxThreadLabelEntity` instance. Pass `None` for no initial data.

#### `Log(data=None)`

Create a new `LogEntity` instance. Pass `None` for no initial data.

#### `OauthGrant(data=None)`

Create a new `OauthGrantEntity` instance. Pass `None` for no initial data.

#### `ReceivedAttachment(data=None)`

Create a new `ReceivedAttachmentEntity` instance. Pass `None` for no initial data.

#### `ReceivedEmail(data=None)`

Create a new `ReceivedEmailEntity` instance. Pass `None` for no initial data.

#### `Segment(data=None)`

Create a new `SegmentEntity` instance. Pass `None` for no initial data.

#### `SegmentContact(data=None)`

Create a new `SegmentContactEntity` instance. Pass `None` for no initial data.

#### `Suppression(data=None)`

Create a new `SuppressionEntity` instance. Pass `None` for no initial data.

#### `SuppressionBatchAdd(data=None)`

Create a new `SuppressionBatchAddEntity` instance. Pass `None` for no initial data.

#### `SuppressionBatchRemove(data=None)`

Create a new `SuppressionBatchRemoveEntity` instance. Pass `None` for no initial data.

#### `Template(data=None)`

Create a new `TemplateEntity` instance. Pass `None` for no initial data.

#### `Topic(data=None)`

Create a new `TopicEntity` instance. Pass `None` for no initial data.

#### `Usage(data=None)`

Create a new `UsageEntity` instance. Pass `None` for no initial data.

#### `Webhook(data=None)`

Create a new `WebhookEntity` instance. Pass `None` for no initial data.

#### `WebhookEvent(data=None)`

Create a new `WebhookEventEntity` instance. Pass `None` for no initial data.

#### `WebhookEventAttempt(data=None)`

Create a new `WebhookEventAttemptEntity` instance. Pass `None` for no initial data.

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

## ApiKeyEntity

```python
api_key = client.ApiKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No | The date and time the API key was created. |
| `domain_id` | `str` | No | Restrict an API key to send emails only from a specific domain. |
| `id` | `str` | No | The ID of the API key. |
| `last_used_at` | `str | None` | No | The date and time the API key was last used. |
| `name` | `str` | Yes | The API key name. |
| `object` | `str` | No | The type of object. |
| `permission` | `str` | No | The API key can have full access to Resend’s API or be only restricted to send emails. |

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

#### `create(reqdata, ctrl=None) -> ApiKeyEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ApiKey().create({
    "name": "example_name",  # str
})
```

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

#### `update(reqdata, ctrl=None) -> ApiKeyEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.ApiKey().update({
    "id": "api_key_id",
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

Create a new `ApiKeyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AttachmentEntity

```python
attachment = client.Attachment()
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

#### `list(reqmatch=None, ctrl=None) -> list[AttachmentEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Attachment().list({"email_id": "example"})
for attachment in results:
    print(attachment.data_get())
```

#### `load(reqmatch, ctrl=None) -> AttachmentEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Attachment().load({"id": "attachment_id", "email_id": "email_id"})
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

Create a new `AttachmentEntity` instance with the same options.

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

| Field | load | list | create | remove |
| --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `object` | - | - | - | - |

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

#### `remove(reqmatch, ctrl=None) -> AudienceEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Audience().remove({"id": "audience_id"})
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

#### `list(reqmatch=None, ctrl=None) -> list[AutomationRunEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.AutomationRun().list({"automation_id": "example"})
for automation_run in results:
    print(automation_run.data_get())
```

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
| `object` | `str` | No | The object type of the response. |
| `preview_text` | `str` | No | The preview text of the email. |
| `reply_to` | `list` | No | The email addresses to which replies should be sent. |
| `scheduled_at` | `str` | No | Timestamp indicating when the broadcast is scheduled to be sent. |
| `segment_id` | `str | None` | No | Unique identifier of the segment this broadcast will be sent to. |
| `send` | `bool` | No | Whether to send the broadcast immediately or keep it as a draft. |
| `sent_at` | `str` | No | Timestamp indicating when the broadcast was sent. |
| `status` | `str` | No | The status of the broadcast. |
| `subject` | `str` | No | The subject line of the email. |
| `text` | `str | None` | No | The plain text version of the broadcast content. |
| `topic_id` | `str | None` | No | The topic ID that the broadcast is scoped to. |

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

#### `create(reqdata, ctrl=None) -> BroadcastEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Broadcast().create({
    "from": "example_from",  # str
    "segment_id": "example_segment_id",  # str | None
    "subject": "example_subject",  # str
})
```

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

#### `remove(reqmatch, ctrl=None) -> BroadcastEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Broadcast().remove({"id": "broadcast_id"})
```

#### `update(reqdata, ctrl=None) -> BroadcastEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Broadcast().update({
    "id": "broadcast_id",
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

Create a new `BroadcastEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BroadcastClickedLinkEntity

```python
broadcast_clicked_link = client.BroadcastClickedLink()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clicks` | `int` | No | Total number of clicks on this URL. |
| `id` | `str` | No | An opaque cursor for this row, used only for pagination. |
| `unique_clicks` | `int` | No | Number of unique clicks on this URL. |
| `url` | `str` | No | The URL that was clicked. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[BroadcastClickedLinkEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.BroadcastClickedLink().list({"broadcast_id": "example"})
for broadcast_clicked_link in results:
    print(broadcast_clicked_link.data_get())
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

Create a new `BroadcastClickedLinkEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BroadcastRecipientEntity

```python
broadcast_recipient = client.BroadcastRecipient()
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

#### `list(reqmatch=None, ctrl=None) -> list[BroadcastRecipientEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.BroadcastRecipient().list({"broadcast_id": "example", "type": "example"})
for broadcast_recipient in results:
    print(broadcast_recipient.data_get())
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

Create a new `BroadcastRecipientEntity` instance with the same options.

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

#### `remove(reqmatch, ctrl=None) -> ContactEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Contact().remove({"id": "contact_id"})
```

#### `update(reqdata, ctrl=None) -> ContactEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Contact().update({
    "id": "contact_id",
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

Create a new `ContactEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactImportEntity

```python
contact_import = client.ContactImport()
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

#### `create(reqdata, ctrl=None) -> ContactImportEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ContactImport().create({
    "file": "example_file",  # str
})
```

Declares a `multipart/form-data` body, which this SDK does not encode yet: it sends the data as JSON.

#### `list(reqmatch=None, ctrl=None) -> list[ContactImportEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ContactImport().list()
for contact_import in results:
    print(contact_import.data_get())
```

#### `load(reqmatch, ctrl=None) -> ContactImportEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.ContactImport().load({"id": "contact_import_id"})
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

Create a new `ContactImportEntity` instance with the same options.

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

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `fallback_value` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `key` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `type` | - | - | Yes | - | - |

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

#### `remove(reqmatch, ctrl=None) -> ContactPropertyEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.ContactProperty().remove({"id": "contact_property_id"})
```

#### `update(reqdata, ctrl=None) -> ContactPropertyEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.ContactProperty().update({
    "id": "contact_property_id",
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

Create a new `ContactPropertyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactSegmentEntity

```python
contact_segment = client.ContactSegment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `str` | No | The ID of the contact. |
| `created_at` | `str` | No | Timestamp indicating when the contact was added to the segment. |
| `id` | `str` | No | Unique identifier for the segment. |
| `name` | `str` | No | Name of the segment. |
| `object` | `str` | No | The object type. |
| `segment_id` | `str` | No | The ID of the segment. |

### Operations

#### `create(reqdata, ctrl=None) -> ContactSegmentEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.ContactSegment().create({
    "contact_id": "example_contact_id",  # str
    "id": "example_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[ContactSegmentEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ContactSegment().list({"contact_id": "example"})
for contact_segment in results:
    print(contact_segment.data_get())
```

#### `remove(reqmatch, ctrl=None) -> ContactSegmentEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.ContactSegment().remove({"id": "contact_segment_id", "contact_id": "contact_id"})
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

Create a new `ContactSegmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactTopicEntity

```python
contact_topic = client.ContactTopic()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `str` | No | The ID of the contact. |
| `description` | `str` | No | Description of the topic. |
| `id` | `str` | No | Unique identifier for the topic. |
| `name` | `str` | No | Name of the topic. |
| `object` | `str` | No | The object type. |
| `subscription` | `str` | No | The subscription status for this topic. |
| `topics` | `list` | No | Array of updated topic subscriptions. |

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

#### `list(reqmatch=None, ctrl=None) -> list[ContactTopicEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ContactTopic().list({"contact_id": "example"})
for contact_topic in results:
    print(contact_topic.data_get())
```

#### `update(reqdata, ctrl=None) -> ContactTopicEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.ContactTopic().update({
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

Create a new `ContactTopicEntity` instance with the same options.

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

#### `update(reqdata, ctrl=None) -> DomainEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Domain().update({
    "id": "domain_id",
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

#### `update(reqdata, ctrl=None) -> EmailEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Email().update({
    "id": "email_id",
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

Create a new `EmailEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EmailBatchEntity

```python
email_batch = client.EmailBatch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | No |  |
| `send_emails` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> EmailBatchEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.EmailBatch().create({
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

Create a new `EmailBatchEntity` instance with the same options.

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

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `object` | - | - | - | - | - |
| `schema` | - | - | - | Yes | - |
| `updated_at` | - | - | - | - | - |

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

#### `remove(reqmatch, ctrl=None) -> EventEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Event().remove({"id": "event_id"})
```

#### `update(reqdata, ctrl=None) -> EventEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Event().update({
    "id": "event_id",
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

#### `remove(reqmatch, ctrl=None) -> InboxEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Inbox().remove({"id": "inbox_id"})
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

## InboxAgentEntity

```python
inbox_agent = client.InboxAgent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `actions` | `list` | Yes | The actions the agent can take, in the order of the enum. |
| `id` | `str` | Yes | The ID of the inbox agent. |
| `instructions` | `str | None` | Yes | Instructions the agent follows when it handles the threads of the inbox. |
| `object` | `str` | Yes | The object type. |
| `tone` | `str | None` | Yes | The tone the agent writes in. |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `actions` | - | Yes |
| `id` | - | - |
| `instructions` | - | Yes |
| `object` | - | - |
| `tone` | - | Yes |

### Operations

#### `load(reqmatch, ctrl=None) -> InboxAgentEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxAgent().load({"inbox_id": "inbox_id"})
```

#### `update(reqdata, ctrl=None) -> InboxAgentEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.InboxAgent().update({
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

Create a new `InboxAgentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InboxDraftEntity

```python
inbox_draft = client.InboxDraft()
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
| `snippet` | `str | None` | Yes | A short preview of the body. |
| `subject` | `str | None` | Yes | The subject. |
| `text` | `str | None` | Yes | The plain text body. |
| `thread_id` | `str | None` | Yes | The Thread ID when the draft is a reply. |
| `to` | `list | None` | Yes | The recipients. |
| `type` | `str` | Yes | `reply` when the draft replies to an email in an existing thread. |
| `updated_at` | `str` | Yes | Timestamp when the draft was last saved. |

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

#### `create(reqdata, ctrl=None) -> InboxDraftEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.InboxDraft().create({
    "inbox_id": "example_inbox_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[InboxDraftEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxDraft().list({"inbox_id": "example"})
for inbox_draft in results:
    print(inbox_draft.data_get())
```

#### `load(reqmatch, ctrl=None) -> InboxDraftEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxDraft().load({"id": "inbox_draft_id", "inbox_id": "inbox_id"})
```

#### `remove(reqmatch, ctrl=None) -> InboxDraftEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.InboxDraft().remove({"id": "inbox_draft_id", "inbox_id": "inbox_id"})
```

#### `update(reqdata, ctrl=None) -> InboxDraftEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.InboxDraft().update({
    "id": "inbox_draft_id",
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

Create a new `InboxDraftEntity` instance with the same options.

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
| `color` | `str` | No | The color of the label, as a `#RRGGBB` hex color. |
| `created_at` | `str` | Yes | Timestamp when the label was created. |
| `id` | `str` | Yes | The ID of the updated label. |
| `name` | `str` | Yes | The name of the label. |
| `object` | `str` | Yes | Type of the response object. |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `color` | Yes | - | - | - |
| `created_at` | - | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `object` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> InboxLabelEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.InboxLabel().create({
    "inbox_id": "example_inbox_id",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[InboxLabelEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxLabel().list({"inbox_id": "example"})
for inbox_label in results:
    print(inbox_label.data_get())
```

#### `remove(reqmatch, ctrl=None) -> InboxLabelEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.InboxLabel().remove({"id": "inbox_label_id", "inbox_id": "inbox_id"})
```

#### `update(reqdata, ctrl=None) -> InboxLabelEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.InboxLabel().update({
    "id": "inbox_label_id",
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
| `object` | `str` | Yes | The object type. |
| `read` | `bool` | Yes | True only when every email in the thread is read. |
| `received_at` | `str` | Yes | Timestamp when the thread was last active. |
| `subject` | `str | None` | Yes | The subject of the thread. |
| `to` | `list` | Yes | The recipients of the most recent email. |

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

#### `list(reqmatch=None, ctrl=None) -> list[InboxThreadEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxThread().list({"inbox_id": "example"})
for inbox_thread in results:
    print(inbox_thread.data_get())
```

#### `load(reqmatch, ctrl=None) -> InboxThreadEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxThread().load({"id": "inbox_thread_id", "inbox_id": "inbox_id"})
```

#### `remove(reqmatch, ctrl=None) -> InboxThreadEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.InboxThread().remove({"id": "inbox_thread_id", "inbox_id": "inbox_id"})
```

#### `update(reqdata, ctrl=None) -> InboxThreadEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.InboxThread().update({
    "id": "inbox_thread_id",
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

#### `create(reqdata, ctrl=None) -> InboxThreadEmailEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

#### `list(reqmatch=None, ctrl=None) -> list[InboxThreadEmailEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.InboxThreadEmail().list({"inbox_id": "example", "thread_id": "example"})
for inbox_thread_email in results:
    print(inbox_thread_email.data_get())
```

#### `load(reqmatch, ctrl=None) -> InboxThreadEmailEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.InboxThreadEmail().load({"id": "inbox_thread_email_id", "inbox_id": "inbox_id", "thread_id": "thread_id"})
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

## InboxThreadLabelEntity

```python
inbox_thread_label = client.InboxThreadLabel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | Yes | Every label on the thread after the change. |
| `id` | `str` | No |  |
| `labels` | `list` | Yes | The IDs of the labels to apply to the thread. |
| `object` | `str` | Yes | Type of the response object. |

### Operations

#### `create(reqdata, ctrl=None) -> InboxThreadLabelEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.InboxThreadLabel().create({
    "inbox_id": "example_inbox_id",  # str
    "thread_id": "example_thread_id",  # str
    "labels": [],  # list
})
```

#### `remove(reqmatch, ctrl=None) -> InboxThreadLabelEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.InboxThreadLabel().remove({"id": "inbox_thread_label_id", "inbox_id": "inbox_id", "thread_id": "thread_id"})
```

#### `update(reqdata, ctrl=None) -> InboxThreadLabelEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.InboxThreadLabel().update({
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

Create a new `InboxThreadLabelEntity` instance with the same options.

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

## OauthGrantEntity

```python
oauth_grant = client.OauthGrant()
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

#### `list(reqmatch=None, ctrl=None) -> list[OauthGrantEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.OauthGrant().list()
for oauth_grant in results:
    print(oauth_grant.data_get())
```

#### `remove(reqmatch, ctrl=None) -> OauthGrantEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.OauthGrant().remove({"id": "oauth_grant_id"})
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

Create a new `OauthGrantEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReceivedAttachmentEntity

```python
received_attachment = client.ReceivedAttachment()
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

#### `list(reqmatch=None, ctrl=None) -> list[ReceivedAttachmentEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.ReceivedAttachment().list({"email_id": "example"})
for received_attachment in results:
    print(received_attachment.data_get())
```

#### `load(reqmatch, ctrl=None) -> ReceivedAttachmentEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.ReceivedAttachment().load({"id": "received_attachment_id", "email_id": "email_id"})
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

Create a new `ReceivedAttachmentEntity` instance with the same options.

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

#### `create(reqdata, ctrl=None) -> SegmentEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Segment().create({
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[SegmentEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Segment().list()
for segment in results:
    print(segment.data_get())
```

#### `load(reqmatch, ctrl=None) -> SegmentEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Segment().load({"id": "segment_id"})
```

#### `remove(reqmatch, ctrl=None) -> SegmentEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Segment().remove({"id": "segment_id"})
```

#### `update(reqdata, ctrl=None) -> SegmentEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Segment().update({
    "id": "segment_id",
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

Create a new `SegmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SegmentContactEntity

```python
segment_contact = client.SegmentContact()
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

#### `list(reqmatch=None, ctrl=None) -> list[SegmentContactEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.SegmentContact().list({"segment_id": "example"})
for segment_contact in results:
    print(segment_contact.data_get())
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

Create a new `SegmentContactEntity` instance with the same options.

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

#### `create(reqdata, ctrl=None) -> SuppressionEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Suppression().create({
    "email": "example_email",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[SuppressionEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Suppression().list()
for suppression in results:
    print(suppression.data_get())
```

#### `load(reqmatch, ctrl=None) -> SuppressionEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Suppression().load({"id": "suppression_id"})
```

#### `remove(reqmatch, ctrl=None) -> SuppressionEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Suppression().remove({"id": "suppression_id"})
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

## SuppressionBatchAddEntity

```python
suppression_batch_add = client.SuppressionBatchAdd()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `emails` | `list` | Yes | Email addresses to suppress. |

### Operations

#### `create(reqdata, ctrl=None) -> SuppressionBatchAddEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.SuppressionBatchAdd().create({
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

Create a new `SuppressionBatchAddEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SuppressionBatchRemoveEntity

```python
suppression_batch_remove = client.SuppressionBatchRemove()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `list` | No | Array containing the removed suppressions. |
| `emails` | `list` | No | Email addresses to remove from the suppression list. |
| `ids` | `list` | No | Suppression IDs to remove from the suppression list. |

### Operations

#### `create(reqdata, ctrl=None) -> SuppressionBatchRemoveEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.SuppressionBatchRemove().create({
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

Create a new `SuppressionBatchRemoveEntity` instance with the same options.

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

#### `create(reqdata, ctrl=None) -> TemplateEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Template().create({
    "html": "example_html",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[TemplateEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Template().list()
for template in results:
    print(template.data_get())
```

#### `load(reqmatch, ctrl=None) -> TemplateEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Template().load({"id": "template_id"})
```

#### `remove(reqmatch, ctrl=None) -> TemplateEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Template().remove({"id": "template_id"})
```

#### `update(reqdata, ctrl=None) -> TemplateEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Template().update({
    "id": "template_id",
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

#### `create(reqdata, ctrl=None) -> TopicEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Topic().create({
    "default_subscription": "example_default_subscription",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list[TopicEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Topic().list()
for topic in results:
    print(topic.data_get())
```

#### `load(reqmatch, ctrl=None) -> TopicEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Topic().load({"id": "topic_id"})
```

#### `remove(reqmatch, ctrl=None) -> TopicEntity`

Remove the entity matching the given criteria. Returns the entity, marked as deleted, and raises on error.

```python
result = client.Topic().remove({"id": "topic_id"})
```

#### `update(reqdata, ctrl=None) -> TopicEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Topic().update({
    "id": "topic_id",
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

Create a new `TopicEntity` instance with the same options.

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

#### `create(reqdata, ctrl=None) -> WebhookEntity`

Create a new entity with the given data. Returns the created entity and raises on error.

```python
result = client.Webhook().create({
    "endpoint": "example_endpoint",  # str
    "events": [],  # list | None
})
```

#### `list(reqmatch=None, ctrl=None) -> list[WebhookEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Webhook().list()
for webhook in results:
    print(webhook.data_get())
```

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

#### `update(reqdata, ctrl=None) -> WebhookEntity`

Update an existing entity. The data must include the entity `id`. Returns the updated entity and raises on error.

```python
result = client.Webhook().update({
    "id": "webhook_id",
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

#### `list(reqmatch=None, ctrl=None) -> list[WebhookEventEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.WebhookEvent().list({"webhook_id": "example"})
for webhook_event in results:
    print(webhook_event.data_get())
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

## WebhookEventAttemptEntity

```python
webhook_event_attempt = client.WebhookEventAttempt()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `http_status_code` | `int` | No | The HTTP status code returned by the webhook endpoint. |
| `id` | `str` | No | The ID of the webhook event attempt. |
| `response` | `str` | No | The response body returned by the webhook endpoint. |
| `sent_at` | `str` | No | Timestamp indicating when the attempt was sent. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[WebhookEventAttemptEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.WebhookEventAttempt().list({"event_id": "example", "webhook_id": "example"})
for webhook_event_attempt in results:
    print(webhook_event_attempt.data_get())
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

Create a new `WebhookEventAttemptEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = ResendSDK({
    "feature": {
        "debug": {"active": True},
        "idempotency": {"active": True},
        "metrics": {"active": True},
        "netsim": {"active": True},
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

