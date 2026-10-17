# Typed models for the Resend SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class ApiKeyRequired(TypedDict):
    name: str


class ApiKey(ApiKeyRequired, total=False):
    created_at: str
    domain_id: str
    id: str
    last_used_at: str | None
    object: str
    permission: str


class ApiKeyListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ApiKeyCreateDataRequired(TypedDict):
    name: str


class ApiKeyCreateData(ApiKeyCreateDataRequired, total=False):
    domain_id: str
    permission: str


class ApiKeyUpdateDataRequired(TypedDict):
    id: str


class ApiKeyUpdateData(ApiKeyUpdateDataRequired, total=False):
    name: str


class ApiKeyRemoveMatch(TypedDict):
    id: str


class Attachment(TypedDict, total=False):
    content_disposition: str | None
    content_id: str
    content_type: str
    download_url: str
    expires_at: str
    filename: str | None
    id: str
    object: str
    size: int


class AttachmentLoadMatch(TypedDict):
    email_id: str
    id: str


class AttachmentListMatchRequired(TypedDict):
    email_id: str


class AttachmentListMatch(AttachmentListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class Audience(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str


class AudienceLoadMatch(TypedDict):
    id: str


class AudienceListMatch(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str


class AudienceCreateData(TypedDict):
    name: str


class AudienceRemoveMatch(TypedDict):
    id: str


class Automation(TypedDict, total=False):
    connections: list
    created_at: str
    id: str
    name: str
    object: str
    status: str
    steps: list
    updated_at: str


class AutomationLoadMatch(TypedDict):
    id: str


class AutomationListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int
    status: str


class AutomationCreateDataRequired(TypedDict):
    connections: list
    name: str
    steps: list


class AutomationCreateData(AutomationCreateDataRequired, total=False):
    status: str


class AutomationUpdateDataRequired(TypedDict):
    id: str


class AutomationUpdateData(AutomationUpdateDataRequired, total=False):
    connections: list
    name: str
    status: str
    steps: list


class AutomationRemoveMatch(TypedDict):
    id: str


class AutomationRun(TypedDict, total=False):
    completed_at: str | None
    created_at: str
    id: str
    object: str
    started_at: str | None
    status: str
    steps: list


class AutomationRunLoadMatch(TypedDict):
    automation_id: str
    id: str


class AutomationRunListMatchRequired(TypedDict):
    automation_id: str


class AutomationRunListMatch(AutomationRunListMatchRequired, total=False):
    after: str
    before: str
    limit: int
    status: str


class Broadcast(TypedDict, total=False):
    audience_id: str | None
    created_at: str
    html: str | None
    id: str
    name: str
    object: str
    preview_text: str
    reply_to: list
    scheduled_at: str
    segment_id: str | None
    send: bool
    sent_at: str
    status: str
    subject: str
    text: str | None
    topic_id: str | None


class BroadcastLoadMatch(TypedDict):
    id: str


class BroadcastListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class BroadcastCreateDataRequired(TypedDict):
    segment_id: str | None
    subject: str


class BroadcastCreateData(BroadcastCreateDataRequired, total=False):
    audience_id: str | None
    html: str | None
    name: str
    preview_text: str
    reply_to: list
    scheduled_at: str
    send: bool
    text: str | None
    topic_id: str | None


class BroadcastUpdateDataRequired(TypedDict):
    id: str


class BroadcastUpdateData(BroadcastUpdateDataRequired, total=False):
    audience_id: str | None
    html: str | None
    name: str
    preview_text: str
    reply_to: list
    segment_id: str | None
    subject: str
    text: str | None
    topic_id: str | None


class BroadcastRemoveMatch(TypedDict):
    id: str


class BroadcastClickedLink(TypedDict, total=False):
    clicks: int
    id: str
    unique_clicks: int
    url: str


class BroadcastClickedLinkListMatchRequired(TypedDict):
    broadcast_id: str


class BroadcastClickedLinkListMatch(BroadcastClickedLinkListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class BroadcastRecipient(TypedDict, total=False):
    bounce_type: str
    clicked_links: list
    contact_id: str | None
    count: int
    email: str
    id: str


class BroadcastRecipientListMatchRequired(TypedDict):
    broadcast_id: str
    type: str


class BroadcastRecipientListMatch(BroadcastRecipientListMatchRequired, total=False):
    after: str
    before: str
    bounce_type: str
    email: str
    limit: int


class Contact(TypedDict, total=False):
    audience_id: str
    created_at: str
    email: str
    first_name: str | None
    id: str
    last_name: str | None
    object: str
    properties: dict
    segments: list
    topics: list
    unsubscribed: bool


class ContactLoadMatch(TypedDict):
    id: str


class ContactListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ContactCreateDataRequired(TypedDict):
    email: str


class ContactCreateData(ContactCreateDataRequired, total=False):
    audience_id: str
    first_name: str | None
    last_name: str | None
    properties: dict
    segments: list
    topics: list
    unsubscribed: bool


class ContactUpdateDataRequired(TypedDict):
    id: str


class ContactUpdateData(ContactUpdateDataRequired, total=False):
    email: str
    first_name: str | None
    last_name: str | None
    properties: dict
    unsubscribed: bool


class ContactRemoveMatch(TypedDict):
    id: str


class ContactImportRequired(TypedDict):
    file: str


class ContactImport(ContactImportRequired, total=False):
    column_map: str
    completed_at: str | None
    counts: dict
    created_at: str
    id: str
    object: str
    on_conflict: str
    segments: str
    status: str
    topics: str


class ContactImportLoadMatch(TypedDict):
    id: str


class ContactImportListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int
    status: str


class ContactImportCreateDataRequired(TypedDict):
    file: str


class ContactImportCreateData(ContactImportCreateDataRequired, total=False):
    column_map: str
    completed_at: str | None
    counts: dict
    created_at: str
    id: str
    object: str
    on_conflict: str
    segments: str
    status: str
    topics: str


class ContactProperty(TypedDict, total=False):
    created_at: str
    fallback_value: Any
    id: str
    key: str
    object: str
    type: str


class ContactPropertyLoadMatch(TypedDict):
    id: str


class ContactPropertyListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class ContactPropertyCreateDataRequired(TypedDict):
    key: str
    type: str


class ContactPropertyCreateData(ContactPropertyCreateDataRequired, total=False):
    fallback_value: Any


class ContactPropertyUpdateDataRequired(TypedDict):
    id: str


class ContactPropertyUpdateData(ContactPropertyUpdateDataRequired, total=False):
    fallback_value: Any


class ContactPropertyRemoveMatch(TypedDict):
    id: str


class ContactSegment(TypedDict, total=False):
    contact_id: str
    created_at: str
    id: str
    name: str
    object: str
    segment_id: str


class ContactSegmentListMatchRequired(TypedDict):
    contact_id: str


class ContactSegmentListMatch(ContactSegmentListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class ContactSegmentCreateDataRequired(TypedDict):
    contact_id: str
    id: str


class ContactSegmentCreateData(ContactSegmentCreateDataRequired, total=False):
    created_at: str
    name: str
    object: str
    segment_id: str


class ContactSegmentRemoveMatch(TypedDict):
    contact_id: str
    id: str


class ContactTopic(TypedDict, total=False):
    contact_id: str
    description: str
    id: str
    name: str
    object: str
    subscription: str
    topics: list


class ContactTopicListMatchRequired(TypedDict):
    contact_id: str


class ContactTopicListMatch(ContactTopicListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class ContactTopicUpdateDataRequired(TypedDict):
    contact_id: str


class ContactTopicUpdateData(ContactTopicUpdateDataRequired, total=False):
    topics: list


class Domain(TypedDict, total=False):
    capabilities: dict
    click_tracking: bool
    created_at: str
    custom_return_path: str
    id: str
    name: str
    object: str
    open_tracking: bool
    records: list
    region: str
    status: str
    tls: str
    tracking_subdomain: str


class DomainLoadMatch(TypedDict):
    id: str


class DomainListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class DomainCreateDataRequired(TypedDict):
    name: str


class DomainCreateData(DomainCreateDataRequired, total=False):
    capabilities: dict
    click_tracking: bool
    custom_return_path: str
    open_tracking: bool
    region: str
    tls: str
    tracking_subdomain: str


class DomainUpdateDataRequired(TypedDict):
    id: str


class DomainUpdateData(DomainUpdateDataRequired, total=False):
    capabilities: dict
    click_tracking: bool
    open_tracking: bool
    tls: str
    tracking_subdomain: str


class DomainRemoveMatch(TypedDict):
    id: str


class DomainClaim(TypedDict, total=False):
    blocked_reason: str | None
    click_tracking: bool
    created_at: str
    custom_return_path: str
    domain_id: str | None
    expires_at: str
    failure_reason: str | None
    id: str
    name: str
    object: str
    open_tracking: bool
    record: dict
    region: str | None
    status: str
    tracking_subdomain: str


class DomainClaimLoadMatch(TypedDict):
    id: str


class DomainClaimCreateDataRequired(TypedDict):
    name: str


class DomainClaimCreateData(DomainClaimCreateDataRequired, total=False):
    blocked_reason: str | None
    click_tracking: bool
    created_at: str
    custom_return_path: str
    domain_id: str | None
    expires_at: str
    failure_reason: str | None
    id: str
    object: str
    open_tracking: bool
    record: dict
    region: str | None
    status: str
    tracking_subdomain: str


class Email(TypedDict, total=False):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    headers: dict
    html: str
    id: str
    last_event: str
    message_id: str
    object: str
    reply_to: list
    scheduled_at: str
    subject: str
    tags: list
    template: Any
    text: str
    to: list
    topic_id: str


class EmailLoadMatch(TypedDict):
    id: str


class EmailListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class EmailCreateDataRequired(TypedDict):
    subject: str
    to: list


class EmailCreateData(EmailCreateDataRequired, total=False):
    attachments: list
    bcc: list
    cc: list
    headers: dict
    html: str
    reply_to: list
    scheduled_at: str
    tags: list
    template: Any
    text: str
    topic_id: str


class EmailUpdateDataRequired(TypedDict):
    id: str


class EmailUpdateData(EmailUpdateDataRequired, total=False):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    headers: dict
    html: str
    last_event: str
    message_id: str
    object: str
    reply_to: list
    scheduled_at: str
    subject: str
    tags: list
    template: Any
    text: str
    to: list
    topic_id: str


class EmailBatch(TypedDict, total=False):
    data: list
    send_emails: list


class EmailBatchCreateData(TypedDict, total=False):
    send_emails: list


class EmailsMetric(TypedDict, total=False):
    broadcast_id: str
    broadcast_name: str
    domain_id: str
    domain_name: str
    email_id: str
    period: str


class EmailsMetricListMatch(TypedDict, total=False):
    broadcast_id: list
    dimension: list
    domain_id: list
    email_id: list
    end_date: str
    granularity: str
    metric: list
    start_date: str
    timezone: str


class Event(TypedDict, total=False):
    created_at: str
    id: str
    name: str
    object: str
    schema: dict | None
    updated_at: str | None


class EventLoadMatch(TypedDict):
    id: str


class EventListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class EventCreateDataRequired(TypedDict):
    name: str


class EventCreateData(EventCreateDataRequired, total=False):
    schema: dict | None


class EventUpdateDataRequired(TypedDict):
    id: str


class EventUpdateData(EventUpdateDataRequired, total=False):
    schema: dict | None


class EventRemoveMatch(TypedDict):
    id: str


class InboxRequired(TypedDict):
    created_at: str
    domain_id: str
    drafts: int
    email_address: str
    from_name: str | None
    id: str
    last_received: str | None
    name: str
    object: str
    receiving_address: str | None
    unread: int
    updated_at: str


class Inbox(InboxRequired, total=False):
    forwarding: bool


class InboxLoadMatch(TypedDict):
    id: str


class InboxListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class InboxCreateDataRequired(TypedDict):
    email_address: str


class InboxCreateData(InboxCreateDataRequired, total=False):
    forwarding: bool
    from_name: str | None
    name: str


class InboxUpdateDataRequired(TypedDict):
    id: str


class InboxUpdateData(InboxUpdateDataRequired, total=False):
    from_name: str | None
    name: str


class InboxRemoveMatch(TypedDict):
    id: str


class InboxAgent(TypedDict):
    actions: list
    id: str
    instructions: str | None
    object: str
    tone: str | None


class InboxAgentLoadMatch(TypedDict):
    inbox_id: str


class InboxAgentUpdateDataRequired(TypedDict):
    inbox_id: str


class InboxAgentUpdateData(InboxAgentUpdateDataRequired, total=False):
    actions: list
    instructions: str | None
    tone: str | None


class InboxDraft(TypedDict):
    attachments: list
    bcc: list
    cc: list
    created_at: str
    email_id: str | None
    html: str | None
    id: str
    object: str
    reply_to_email_id: str | None
    snippet: str | None
    subject: str | None
    text: str | None
    thread_id: str | None
    to: list | None
    type: str
    updated_at: str


class InboxDraftLoadMatch(TypedDict):
    id: str
    inbox_id: str


class InboxDraftListMatchRequired(TypedDict):
    inbox_id: str


class InboxDraftListMatch(InboxDraftListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class InboxDraftCreateDataRequired(TypedDict):
    inbox_id: str


class InboxDraftCreateData(InboxDraftCreateDataRequired, total=False):
    attachments: list
    bcc: list
    cc: list
    html: str | None
    reply_to_email_id: str | None
    subject: str | None
    text: str | None
    thread_id: str | None
    to: list | None


class InboxDraftUpdateDataRequired(TypedDict):
    id: str
    inbox_id: str


class InboxDraftUpdateData(InboxDraftUpdateDataRequired, total=False):
    attachments: list
    bcc: list
    cc: list
    html: str | None
    subject: str | None
    text: str | None
    to: list | None


class InboxDraftRemoveMatch(TypedDict):
    id: str
    inbox_id: str


class InboxLabelRequired(TypedDict):
    created_at: str
    id: str
    name: str
    object: str


class InboxLabel(InboxLabelRequired, total=False):
    color: str


class InboxLabelListMatch(TypedDict):
    inbox_id: str


class InboxLabelCreateDataRequired(TypedDict):
    inbox_id: str
    name: str


class InboxLabelCreateData(InboxLabelCreateDataRequired, total=False):
    color: str


class InboxLabelUpdateDataRequired(TypedDict):
    id: str
    inbox_id: str


class InboxLabelUpdateData(InboxLabelUpdateDataRequired, total=False):
    color: str
    name: str


class InboxLabelRemoveMatch(TypedDict):
    id: str
    inbox_id: str


class InboxThread(TypedDict):
    bcc: list
    cc: list
    email_count: int
    folder: str
    has_attachment: bool
    has_draft: bool
    id: str
    labels: list
    object: str
    read: bool
    received_at: str
    subject: str | None
    to: list


class InboxThreadLoadMatch(TypedDict):
    id: str
    inbox_id: str


class InboxThreadListMatchRequired(TypedDict):
    inbox_id: str


class InboxThreadListMatch(InboxThreadListMatchRequired, total=False):
    after: str
    before: str
    folder: list
    label: list
    limit: int
    read: bool


class InboxThreadUpdateDataRequired(TypedDict):
    id: str
    inbox_id: str


class InboxThreadUpdateData(InboxThreadUpdateDataRequired, total=False):
    folder: str
    read: bool


class InboxThreadRemoveMatch(TypedDict):
    id: str
    inbox_id: str


class InboxThreadEmail(TypedDict):
    attachments: list
    bcc: list
    cc: list
    direction: str
    html: str | None
    id: str
    message_id: str | None
    read: bool
    received_at: str
    reply_to: list
    subject: str | None
    text: str | None
    to: list


class InboxThreadEmailLoadMatch(TypedDict):
    id: str
    inbox_id: str
    thread_id: str


class InboxThreadEmailListMatchRequired(TypedDict):
    inbox_id: str
    thread_id: str


class InboxThreadEmailListMatch(InboxThreadEmailListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class InboxThreadEmailCreateDataRequired(TypedDict):
    id: str
    inbox_id: str
    thread_id: str
    attachments: list
    bcc: list
    cc: list
    html: str | None
    subject: str | None
    text: str | None
    to: list


class InboxThreadEmailCreateData(InboxThreadEmailCreateDataRequired, total=False):
    headers: Any
    reply_all: Any


class InboxThreadLabelRequired(TypedDict):
    data: list
    labels: list
    object: str


class InboxThreadLabel(InboxThreadLabelRequired, total=False):
    id: str


class InboxThreadLabelCreateData(TypedDict):
    inbox_id: str
    thread_id: str
    labels: list


class InboxThreadLabelUpdateDataRequired(TypedDict):
    inbox_id: str
    thread_id: str


class InboxThreadLabelUpdateData(InboxThreadLabelUpdateDataRequired, total=False):
    labels: list


class InboxThreadLabelRemoveMatch(TypedDict):
    id: str
    inbox_id: str
    thread_id: str


class Log(TypedDict, total=False):
    created_at: str
    endpoint: str
    id: str
    method: str
    object: str
    request_body: dict | None
    response_body: dict | None
    response_status: int
    user_agent: str | None


class LogLoadMatch(TypedDict):
    id: str


class LogListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class OauthGrant(TypedDict, total=False):
    client: dict
    client_id: str
    created_at: str
    id: str
    revoked_at: str | None
    revoked_reason: str | None
    scopes: list


class OauthGrantListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class OauthGrantRemoveMatch(TypedDict):
    id: str


class ReceivedAttachment(TypedDict, total=False):
    content_disposition: str | None
    content_id: str
    content_type: str
    download_url: str
    expires_at: str
    filename: str | None
    id: str
    object: str
    size: int


class ReceivedAttachmentLoadMatch(TypedDict):
    email_id: str
    id: str


class ReceivedAttachmentListMatchRequired(TypedDict):
    email_id: str


class ReceivedAttachmentListMatch(ReceivedAttachmentListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class ReceivedEmail(TypedDict, total=False):
    attachments: list
    bcc: list | None
    cc: list | None
    created_at: str
    headers: dict | None
    html: str | None
    id: str
    message_id: str
    object: str
    received_for: list
    reply_to: list | None
    subject: str
    text: str | None
    to: list


class ReceivedEmailLoadMatch(TypedDict):
    email_id: str


class ReceivedEmailListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class Segment(TypedDict, total=False):
    audience_id: str
    created_at: str
    filter: dict
    id: str
    name: str
    object: str


class SegmentLoadMatch(TypedDict):
    id: str


class SegmentListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class SegmentCreateDataRequired(TypedDict):
    name: str


class SegmentCreateData(SegmentCreateDataRequired, total=False):
    audience_id: str
    filter: dict


class SegmentUpdateDataRequired(TypedDict):
    id: str


class SegmentUpdateData(SegmentUpdateDataRequired, total=False):
    name: str


class SegmentRemoveMatch(TypedDict):
    id: str


class SegmentContact(TypedDict, total=False):
    created_at: str
    email: str
    first_name: str | None
    id: str
    last_name: str | None
    unsubscribed: bool


class SegmentContactListMatchRequired(TypedDict):
    segment_id: str


class SegmentContactListMatch(SegmentContactListMatchRequired, total=False):
    after: str
    before: str
    limit: int


class Suppression(TypedDict, total=False):
    created_at: str
    email: str
    id: str
    object: str
    origin: str
    source_id: str | None


class SuppressionLoadMatch(TypedDict):
    id: str


class SuppressionListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int
    origin: str


class SuppressionCreateData(TypedDict):
    email: str


class SuppressionRemoveMatch(TypedDict):
    id: str


class SuppressionBatchAdd(TypedDict):
    emails: list


class SuppressionBatchAddCreateData(TypedDict):
    emails: list


class SuppressionBatchRemove(TypedDict, total=False):
    data: list
    emails: list
    ids: list


class SuppressionBatchRemoveCreateData(TypedDict, total=False):
    emails: list
    ids: list


class Template(TypedDict, total=False):
    alias: str
    created_at: str
    current_version_id: str
    has_unpublished_versions: bool
    html: str
    id: str
    name: str
    object: str
    published_at: str | None
    reply_to: list | None
    status: str
    subject: str
    text: str
    updated_at: str
    variables: list


class TemplateLoadMatch(TypedDict):
    id: str


class TemplateListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class TemplateCreateDataRequired(TypedDict):
    html: str
    name: str


class TemplateCreateData(TemplateCreateDataRequired, total=False):
    alias: str
    reply_to: list | None
    subject: str
    text: str
    variables: list


class TemplateUpdateDataRequired(TypedDict):
    id: str


class TemplateUpdateData(TemplateUpdateDataRequired, total=False):
    alias: str
    html: str
    name: str
    reply_to: list | None
    subject: str
    text: str
    variables: list


class TemplateRemoveMatch(TypedDict):
    id: str


class Topic(TypedDict, total=False):
    created_at: str
    default_subscription: str
    description: str
    id: str
    name: str
    object: str
    visibility: str


class TopicLoadMatch(TypedDict):
    id: str


class TopicListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class TopicCreateDataRequired(TypedDict):
    default_subscription: str
    name: str


class TopicCreateData(TopicCreateDataRequired, total=False):
    description: str
    visibility: str


class TopicUpdateDataRequired(TypedDict):
    id: str


class TopicUpdateData(TopicUpdateDataRequired, total=False):
    description: str
    name: str
    visibility: str


class TopicRemoveMatch(TypedDict):
    id: str


class Usage(TypedDict, total=False):
    ai_credits: dict
    automation_runs: dict
    broadcasts: dict
    contacts: dict
    domains: dict
    emails: dict
    object: str
    rate_limit: dict
    segments: dict


class UsageLoadMatch(TypedDict, total=False):
    ai_credits: dict
    automation_runs: dict
    broadcasts: dict
    contacts: dict
    domains: dict
    emails: dict
    object: str
    rate_limit: dict
    segments: dict


class Webhook(TypedDict, total=False):
    created_at: str
    endpoint: str
    events: list | None
    id: str
    object: str
    signing_secret: str
    status: str


class WebhookLoadMatch(TypedDict):
    id: str


class WebhookListMatch(TypedDict, total=False):
    after: str
    before: str
    limit: int


class WebhookCreateData(TypedDict):
    endpoint: str
    events: list | None


class WebhookUpdateDataRequired(TypedDict):
    id: str


class WebhookUpdateData(WebhookUpdateDataRequired, total=False):
    endpoint: str
    events: list | None
    status: str


class WebhookRemoveMatch(TypedDict):
    id: str


class WebhookEvent(TypedDict, total=False):
    created_at: str
    id: str
    next_attempt_at: str | None
    object: str
    payload: dict
    status: str
    type: str


class WebhookEventLoadMatch(TypedDict):
    id: str
    webhook_id: str


class WebhookEventListMatchRequired(TypedDict):
    webhook_id: str


class WebhookEventListMatch(WebhookEventListMatchRequired, total=False):
    after: str
    limit: int


class WebhookEventCreateDataRequired(TypedDict):
    event_id: str
    webhook_id: str


class WebhookEventCreateData(WebhookEventCreateDataRequired, total=False):
    created_at: str
    id: str
    next_attempt_at: str | None
    object: str
    payload: dict
    status: str
    type: str


class WebhookEventAttempt(TypedDict, total=False):
    http_status_code: int
    id: str
    response: str
    sent_at: str


class WebhookEventAttemptListMatchRequired(TypedDict):
    event_id: str
    webhook_id: str


class WebhookEventAttemptListMatch(WebhookEventAttemptListMatchRequired, total=False):
    after: str
    limit: int
