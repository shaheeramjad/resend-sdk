

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { ResendSDK, BaseFeature, config, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('InboxDraftEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.InboxDraft()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ResendSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.InboxDraft().list({"inbox_id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'inbox_draft.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"attachments":{"a":true,"h":"Attachments","n":"attachments","op":{"create":{"req":false,"type":["`$ONE`",["`$ARRAY`","`$NULL`"]]},"update":{"req":false,"type":["`$ONE`",["`$ARRAY`","`$NULL`"]]}},"r":true,"sh":"The files stored on the draft.","t":"`$ARRAY`","key$":"attachments","index$":0},"bcc":{"a":true,"h":"Bcc","n":"bcc","op":{"create":{"req":false,"type":"`$ANY`"},"update":{"req":false,"type":"`$ANY`"}},"r":true,"sh":"The BCC recipients.","t":"`$ARRAY`","key$":"bcc","index$":1},"cc":{"a":true,"h":"Cc","n":"cc","op":{"create":{"req":false,"type":"`$ANY`"},"update":{"req":false,"type":"`$ANY`"}},"r":true,"sh":"The CC recipients.","t":"`$ARRAY`","key$":"cc","index$":2},"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":true,"sh":"Timestamp when the draft was created.","t":"`$STRING`","key$":"created_at","index$":3},"email_id":{"a":true,"fo":"uuid","h":"Email Id","n":"email_id","r":true,"sh":"The ID of the outbound email queued when the draft was sent.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"email_id","index$":4},"html":{"a":true,"h":"Html","n":"html","op":{"create":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]},"update":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]}},"r":true,"sh":"The HTML body.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"html","index$":5},"id":{"a":true,"fo":"uuid","h":"Id","n":"id","r":true,"sh":"The ID of the draft.","t":"`$STRING`","key$":"id","index$":6},"object":{"a":true,"h":"Object","n":"object","r":true,"sh":"The object type.","t":"`$STRING`","key$":"object","index$":7},"reply_to_email_id":{"a":true,"fo":"uuid","h":"Reply To Email Id","n":"reply_to_email_id","op":{"create":{"req":false,"type":"`$STRING`"}},"r":true,"sh":"The Email ID the draft replies to.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"reply_to_email_id","index$":8},"snippet":{"a":true,"h":"Snippet","n":"snippet","r":true,"sh":"A short preview of the body.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"snippet","index$":9},"subject":{"a":true,"h":"Subject","n":"subject","op":{"create":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]},"update":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]}},"r":true,"sh":"The subject.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"subject","index$":10},"text":{"a":true,"h":"Text","n":"text","op":{"create":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]},"update":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]}},"r":true,"sh":"The plain text body.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"text","index$":11},"thread_id":{"a":true,"fo":"uuid","h":"Thread Id","n":"thread_id","op":{"create":{"req":false,"type":"`$STRING`"}},"r":true,"sh":"The Thread ID when the draft is a reply.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"thread_id","index$":12},"to":{"a":true,"h":"To","n":"to","op":{"create":{"req":false,"type":"`$ANY`"},"update":{"req":false,"type":"`$ANY`"}},"r":true,"sh":"The recipients.","t":["`$ONE`",["`$ARRAY`","`$NULL`"]],"key$":"to","index$":13},"type":{"a":true,"h":"Type","n":"type","r":true,"sh":"`reply` when the draft replies to an email in an existing thread.","t":"`$STRING`","key$":"type","index$":14},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":true,"sh":"Timestamp when the draft was last saved.","t":"`$STRING`","key$":"updated_at","index$":15}},"id":{"field":"id","name":"id"},"name":"inbox_draft","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["headers"],"co":{"id":"POST /inboxes/{inbox_id}/drafts/{draft_id}/send","source":"openapi3","version":2},"g":{"header":[{"a":true,"k":"header","n":"idempotency_key","or":"Idempotency-Key","r":false,"t":"`$STRING`","index$":0}],"params":[{"a":true,"k":"param","n":"id","or":"draft_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"POST","o":"/inboxes/{inbox_id}/drafts/{draft_id}/send","q":{"$action":"send","exist":["id","inbox_id"]},"r":{"param":{"draft_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"drafts"},{"var":"id"},{"lit":"send"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"bf":["attachments","bcc","cc","html","reply_to_email_id","subject","text","thread_id","to"],"co":{"id":"POST /inboxes/{inbox_id}/drafts","source":"openapi3","version":2},"g":{"header":[{"a":true,"k":"header","n":"idempotency_key","or":"Idempotency-Key","r":false,"t":"`$STRING`","index$":0}],"params":[{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/inboxes/{inbox_id}/drafts","q":{"exist":["inbox_id"]},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"drafts"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /inboxes/{inbox_id}/drafts","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"after","or":"after","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"before","or":"before","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":2}]},"k":"http","m":"GET","o":"/inboxes/{inbox_id}/drafts","q":{"exist":["inbox_id"]},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"drafts"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /inboxes/{inbox_id}/drafts/{draft_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"draft_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/inboxes/{inbox_id}/drafts/{draft_id}","q":{"exist":["id","inbox_id"]},"r":{"param":{"draft_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"drafts"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /inboxes/{inbox_id}/drafts/{draft_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"draft_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/inboxes/{inbox_id}/drafts/{draft_id}","q":{"exist":["id","inbox_id"]},"r":{"param":{"draft_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"drafts"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"bf":["attachments","bcc","cc","html","subject","text","to"],"co":{"id":"PATCH /inboxes/{inbox_id}/drafts/{draft_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"draft_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PATCH","o":"/inboxes/{inbox_id}/drafts/{draft_id}","q":{"exist":["id","inbox_id"]},"r":{"param":{"draft_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"drafts"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[["$.main.kit.entity.inbox"]]},"key$":"inbox_draft","name__orig":"inbox_draft","Name":"InboxDraft","name_":"inbox_draft","name-":"inbox-draft","NAME":"INBOX_DRAFT","index$":21}, {"active":true,"entity":"inbox_draft","key$":"BasicInboxDraftFlow","kind":"basic","name":"BasicInboxDraftFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"inbox_draft_ref01"},"m":{"inbox_id":"inbox01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{"inbox_id":"inbox01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"inbox_draft_ref01"}}],"index$":1},{"a":true,"d":{"inbox_id":"inbox01"},"i":{"ref":"inbox_draft_ref01","srcdatavar":"inbox_draft_ref01_data","suffix":"_up0","textfield":"created_at"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-inbox_draft_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"inbox_draft_ref01","srcdatavar":"inbox_draft_ref01_data","suffix":"_dt0"},"m":{"id":"inbox_draft01","inbox_id":"inbox01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-inbox_draft_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"inbox_draft_ref01","suffix":"_rm0"},"m":{"id":"inbox_draft01","inbox_id":"inbox01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{"inbox_id":"inbox01"},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"inbox_draft_ref01"}}],"index$":5}]}, 'InboxDraft', {"POST /inboxes/{inbox_id}/drafts/{draft_id}/send":{"protocol":"http","requestBody":{"required":false,"content":{"application/json":{"schema":{"type":"object","properties":{"headers":{"type":["object","null"],"additionalProperties":{"type":"string"},"description":"Custom headers to add to the email. The `In-Reply-To`, `References` and `Reply-To` headers are set by Resend and cannot be set here. Sending one of them, in any letter case, returns a 422. Header names cannot be empty, and names and values cannot contain line breaks or null characters, or the request returns a 422. `null` means no custom headers.","example":{"X-Ticket-Id":"4821"}}},"x-ref":"#/components/schemas/SendInboxDraftOptions"}}}},"parameters":[{"in":"header","name":"Idempotency-Key","required":false,"schema":{"type":"string","minLength":1,"maxLength":256},"description":"A unique key for the request, so the draft is sent only once. A retry with the same key within 24 hours returns the first successful (2xx) response, with the `idempotent-replayed: true` header. A failed request is not stored, so it can be retried with the same key. Reusing the key with a different body, or while the first request is still running, returns a 409. An empty key, or one longer than 256 characters, returns a 400 `invalid_idempotency_key`. [Learn more](https://resend.com/docs/dashboard/emails/idempotency-keys)","index$":0},{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":1},{"name":"draft_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Draft ID.","index$":2}]},"POST /inboxes/{inbox_id}/drafts":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","description":"At least one of `to`, `cc`, `bcc`, `subject`, `text`, `html` or `attachments` must be non-empty. The combined recipients of `to`, `cc` and `bcc` cannot exceed 50. `thread_id` and `reply_to_email_id` must be sent together or not at all.","properties":{"to":{"description":"The recipients. For multiple addresses, send an array of strings. Each value can include a friendly name, for example `Ada Lovelace <ada@example.org>`. Each address is `email@example.com` or `Name <email@example.com>`. An address that is not valid returns a 422.","oneOf":[{"type":"string"},{"type":"array","items":{"type":"string"},"maxItems":50},{"type":"null"}],"example":["ada@example.org"],"key$":"to"},"cc":{"description":"The CC recipients. For multiple addresses, send an array of strings. Each address is `email@example.com` or `Name <email@example.com>`. An address that is not valid returns a 422.","oneOf":[{"type":"string"},{"type":"array","items":{"type":"string"},"maxItems":50},{"type":"null"}],"key$":"cc"},"bcc":{"description":"The BCC recipients. For multiple addresses, send an array of strings. Each address is `email@example.com` or `Name <email@example.com>`. An address that is not valid returns a 422.","oneOf":[{"type":"string"},{"type":"array","items":{"type":"string"},"maxItems":50},{"type":"null"}],"key$":"bcc"},"subject":{"type":["string","null"],"minLength":1,"maxLength":2000,"description":"The subject. Max 2000 characters. A string that is only whitespace, or that contains a line break, is rejected.","example":"Refund for order 1041","key$":"subject"},"html":{"type":["string","null"],"description":"The HTML body.","example":"<p>Refund issued for order 1041.</p>","key$":"html"},"text":{"type":["string","null"],"description":"The plain text body.","example":"Refund issued for order 1041.","key$":"text"},"attachments":{"type":["array","null"],"description":"Files to store on the draft. Each one needs `content`, `filename` and `content_type`. A `path` without `content` returns a 422, and when both are sent, `path` is ignored. A `content_id` returns a 422. The files on a draft can total at most 25 MB. An empty file or a blocked file type returns a 422.","items":{"type":"object","description":"A file to store on a draft. Unlike `POST /emails` attachments, a draft attachment cannot use `path` or `content_id`. A `path` without `content` returns a 422, and when both are sent, `path` is ignored.","required":["content","filename","content_type"],"properties":{"content":{"type":"string","description":"The content of the file, encoded in base64. It cannot be empty.","example":"aGVsbG8gZTJl"},"filename":{"type":"string","minLength":1,"description":"The filename of the attachment. A blocked file type returns a 422.","example":"receipt.pdf"},"content_type":{"type":"string","minLength":1,"description":"The MIME type of the attachment. It must be a valid MIME type, such as `application/pdf`, or the request returns a 422.","example":"application/pdf"}},"x-ref":"#/components/schemas/InboxDraftAttachmentInput"},"key$":"attachments"},"thread_id":{"type":"string","format":"uuid","description":"The Thread ID to reply in. Must be sent with `reply_to_email_id`.","example":"4d8e2a1c-9b3f-4c6d-8a21-3e5f7c9c0d12","key$":"thread_id"},"reply_to_email_id":{"type":"string","format":"uuid","description":"The Email ID to reply to, as returned in `data[].id` by the list thread emails endpoint. Must be sent with `thread_id`.","example":"5b1a9f47-2c8d-4e6f-9a03-1d2e3f4a5b60","key$":"reply_to_email_id"}},"x-ref":"#/components/schemas/CreateInboxDraftOptions","index$":1}}}},"parameters":[{"in":"header","name":"Idempotency-Key","required":false,"schema":{"type":"string","minLength":1,"maxLength":256},"description":"A unique key for the request, so the draft is created only once. A retry with the same key within 24 hours returns the first successful (2xx) response, with the `idempotent-replayed: true` header. A failed request is not stored, so it can be retried with the same key. Reusing the key with a different body, or while the first request is still running, returns a 409. An empty key, or one longer than 256 characters, returns a 400 `invalid_idempotency_key`. [Learn more](https://resend.com/docs/dashboard/emails/idempotency-keys)","index$":0},{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":1}]},"GET /inboxes/{inbox_id}/drafts":{"protocol":"http","parameters":[{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":0},{"in":"query","name":"limit","required":false,"schema":{"type":"integer","minimum":1,"maximum":100},"description":"Number of items to return.","x-ref":"#/components/parameters/PaginationLimit","index$":1},{"in":"query","name":"after","required":false,"schema":{"type":"string"},"description":"Return items after this cursor.","x-ref":"#/components/parameters/PaginationAfter","index$":2},{"in":"query","name":"before","required":false,"schema":{"type":"string"},"description":"Return items before this cursor.","x-ref":"#/components/parameters/PaginationBefore","index$":3}]},"GET /inboxes/{inbox_id}/drafts/{draft_id}":{"protocol":"http","parameters":[{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":0},{"name":"draft_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Draft ID.","index$":1}]},"DELETE /inboxes/{inbox_id}/drafts/{draft_id}":{"protocol":"http","parameters":[{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":0},{"name":"draft_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Draft ID.","index$":1}]},"PATCH /inboxes/{inbox_id}/drafts/{draft_id}":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","description":"Send at least one of `to`, `cc`, `bcc`, `subject`, `text`, `html` or `attachments`. Omitted fields are left unchanged. Setting `text` or `html` replaces the whole body, so the body part you omit becomes `null`. The merged draft must contain at least one non-empty field, and its combined recipients cannot exceed 50.","anyOf":[{"required":["to"]},{"required":["cc"]},{"required":["bcc"]},{"required":["subject"]},{"required":["text"]},{"required":["html"]},{"required":["attachments"]}],"properties":{"to":{"description":"The recipients. For multiple addresses, send an array of strings. Each value can include a friendly name. Pass `null` to clear. Each address is `email@example.com` or `Name <email@example.com>`. An address that is not valid returns a 422.","oneOf":[{"type":"string"},{"type":"array","items":{"type":"string"},"maxItems":50},{"type":"null"}],"example":["ada@example.org"],"key$":"to"},"cc":{"description":"The CC recipients. Pass `null` or `[]` to clear. Each address is `email@example.com` or `Name <email@example.com>`. An address that is not valid returns a 422.","oneOf":[{"type":"string"},{"type":"array","items":{"type":"string"},"maxItems":50},{"type":"null"}],"key$":"cc"},"bcc":{"description":"The BCC recipients. Pass `null` or `[]` to clear. Each address is `email@example.com` or `Name <email@example.com>`. An address that is not valid returns a 422.","oneOf":[{"type":"string"},{"type":"array","items":{"type":"string"},"maxItems":50},{"type":"null"}],"key$":"bcc"},"subject":{"type":["string","null"],"minLength":1,"maxLength":2000,"description":"The subject. Max 2000 characters. Cannot be only whitespace or contain a line break. Pass `null` to clear.","example":"Refund issued for order 1041","key$":"subject"},"html":{"type":["string","null"],"description":"The HTML body. Pass `null` to clear.","key$":"html"},"text":{"type":["string","null"],"description":"The plain text body. Pass `null` to clear.","key$":"text"},"attachments":{"type":["array","null"],"description":"Files to store on the draft. Each one needs `content`, `filename` and `content_type`. A `path` without `content` returns a 422, and when both are sent, `path` is ignored. A `content_id` returns a 422. The files on a draft can total at most 25 MB. An empty file or a blocked file type returns a 422. Sending the field replaces every attachment on the draft. Send `[]` to remove them all. `null` also removes them all.","items":{"type":"object","description":"A file to store on a draft. Unlike `POST /emails` attachments, a draft attachment cannot use `path` or `content_id`. A `path` without `content` returns a 422, and when both are sent, `path` is ignored.","required":["content","filename","content_type"],"properties":{"content":{"type":"string","description":"The content of the file, encoded in base64. It cannot be empty.","example":"aGVsbG8gZTJl"},"filename":{"type":"string","minLength":1,"description":"The filename of the attachment. A blocked file type returns a 422.","example":"receipt.pdf"},"content_type":{"type":"string","minLength":1,"description":"The MIME type of the attachment. It must be a valid MIME type, such as `application/pdf`, or the request returns a 422.","example":"application/pdf"}},"x-ref":"#/components/schemas/InboxDraftAttachmentInput"},"key$":"attachments"}},"x-ref":"#/components/schemas/UpdateInboxDraftOptions","index$":1}}}},"parameters":[{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":0},{"name":"draft_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Draft ID.","index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const inbox_draft_ref01_ent = client.InboxDraft()
    let inbox_draft_ref01_data = setup.data.new.inbox_draft['inbox_draft_ref01']
    inbox_draft_ref01_data['inbox_id'] = setup.idmap['inbox01']

    inbox_draft_ref01_data = (await inbox_draft_ref01_ent.create(inbox_draft_ref01_data)).data()
    assert(null != inbox_draft_ref01_data.id)


    // LIST
    const inbox_draft_ref01_match: any = {}
    inbox_draft_ref01_match['inbox_id'] = setup.idmap['inbox01']

    const inbox_draft_ref01_list = (await inbox_draft_ref01_ent.list(inbox_draft_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(inbox_draft_ref01_list, { id: inbox_draft_ref01_data.id })))


    // UPDATE
    const inbox_draft_ref01_data_up0: any = {}
    inbox_draft_ref01_data_up0.id = inbox_draft_ref01_data.id
    inbox_draft_ref01_data_up0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01']

    const inbox_draft_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-inbox_draft_ref01_' + setup.now }
    ;(inbox_draft_ref01_data_up0 as any)[inbox_draft_ref01_markdef_up0.name] = inbox_draft_ref01_markdef_up0.value

    const inbox_draft_ref01_resdata_up0 = (await inbox_draft_ref01_ent.update(inbox_draft_ref01_data_up0)).data()
    assert(inbox_draft_ref01_resdata_up0.id === inbox_draft_ref01_data_up0.id)

    assert((inbox_draft_ref01_resdata_up0 as any)[inbox_draft_ref01_markdef_up0.name] === inbox_draft_ref01_markdef_up0.value)


    // LOAD
    const inbox_draft_ref01_match_dt0: any = {}
    inbox_draft_ref01_match_dt0.id = inbox_draft_ref01_data.id
    inbox_draft_ref01_match_dt0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01']
    const inbox_draft_ref01_data_dt0 = (await inbox_draft_ref01_ent.load(inbox_draft_ref01_match_dt0)).data()
    assert(inbox_draft_ref01_data_dt0.id === inbox_draft_ref01_data.id)


    // REMOVE
    const inbox_draft_ref01_match_rm0: any = { id: inbox_draft_ref01_data.id }
    inbox_draft_ref01_match_rm0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01']
    await inbox_draft_ref01_ent.remove(inbox_draft_ref01_match_rm0)
  

    // LIST
    const inbox_draft_ref01_match_rt0: any = {}
    inbox_draft_ref01_match_rt0['inbox_id'] = setup.idmap['inbox01']

    const inbox_draft_ref01_list_rt0 = (await inbox_draft_ref01_ent.list(inbox_draft_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(inbox_draft_ref01_list_rt0, { id: inbox_draft_ref01_data.id })))


  })
})



// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true

function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/inbox_draft/InboxDraftTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = ResendSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['inbox_draft01','inbox_draft02','inbox_draft03','inbox01','inbox02','inbox03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_INBOX_DRAFT_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_INBOX_DRAFT_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_INBOX_DRAFT_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new ResendSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.RESEND_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.RESEND_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
