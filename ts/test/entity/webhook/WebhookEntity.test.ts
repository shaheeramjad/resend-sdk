

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


describe('WebhookEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.Webhook()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('webhook hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ResendSDK.test(offline).Webhook().stream('list')) { }
    }, /offline/)

    for await (const _item of ResendSDK.test(offline).Webhook()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ResendSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.Webhook().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ResendSDK.test().Webhook().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ResendSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.Webhook().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.Webhook().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ResendSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Webhook().list({"after":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'webhook.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"created_at":{"a":true,"h":"Created At","n":"created_at","r":false,"sh":"Timestamp indicating when the webhook was created.","t":"`$STRING`","key$":"created_at","index$":0},"endpoint":{"a":true,"h":"Endpoint","n":"endpoint","op":{"create":{"req":true,"type":"`$STRING`"}},"r":false,"sh":"The URL where webhook events are sent.","t":"`$STRING`","key$":"endpoint","index$":1},"events":{"a":true,"h":"Events","n":"events","op":{"create":{"req":true,"type":"`$ARRAY`"}},"r":false,"sh":"Array of event types subscribed to.","t":["`$ONE`",["`$ARRAY`","`$NULL`"]],"key$":"events","index$":2},"id":{"a":true,"fo":"uuid","h":"Id","n":"id","r":false,"sh":"The ID of the webhook.","t":"`$STRING`","key$":"id","index$":3},"object":{"a":true,"h":"Object","n":"object","r":false,"sh":"The type of object.","t":"`$STRING`","key$":"object","index$":4},"signing_secret":{"a":true,"h":"Signing Secret","n":"signing_secret","r":false,"sh":"The secret key used to verify webhook payloads.","t":"`$STRING`","key$":"signing_secret","index$":5},"status":{"a":true,"h":"Status","n":"status","r":false,"sh":"The status of the webhook.","t":"`$STRING`","key$":"status","index$":6}},"id":{"field":"id","name":"id"},"name":"webhook","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /webhooks/{webhook_id}/signing-secret/rotate","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"webhook_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/webhooks/{webhook_id}/signing-secret/rotate","q":{"$action":"rotate","exist":["id"]},"r":{"param":{"webhook_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"webhooks"},{"var":"id"},{"lit":"signing-secret"},{"lit":"rotate"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"bf":["endpoint","events"],"co":{"id":"POST /webhooks","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/webhooks","q":{},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"webhooks"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /webhooks","source":"openapi3","version":2},"g":{"query":[{"a":true,"k":"query","n":"after","or":"after","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"before","or":"before","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":2}]},"k":"http","m":"GET","o":"/webhooks","q":{},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"webhooks"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /webhooks/{webhook_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"webhook_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/webhooks/{webhook_id}","q":{"exist":["id"]},"r":{"param":{"webhook_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"webhooks"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /webhooks/{webhook_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"webhook_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/webhooks/{webhook_id}","q":{"exist":["id"]},"r":{"param":{"webhook_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"webhooks"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"bf":["endpoint","events","status"],"co":{"id":"PATCH /webhooks/{webhook_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"webhook_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/webhooks/{webhook_id}","q":{"exist":["id"]},"r":{"param":{"webhook_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"webhooks"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"webhook","name__orig":"webhook","Name":"Webhook","name_":"webhook","name-":"webhook","NAME":"WEBHOOK","index$":38}, {"active":true,"entity":"webhook","key$":"BasicWebhookFlow","kind":"basic","name":"BasicWebhookFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"webhook_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"webhook_ref01"}}],"index$":1},{"a":true,"d":{},"i":{"ref":"webhook_ref01","srcdatavar":"webhook_ref01_data","suffix":"_up0","textfield":"created_at"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-webhook_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"webhook_ref01","srcdatavar":"webhook_ref01_data","suffix":"_dt0"},"m":{"id":"webhook01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-webhook_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"webhook_ref01","suffix":"_rm0"},"m":{"id":"webhook01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"webhook_ref01"}}],"index$":5}]}, 'Webhook', {"POST /webhooks/{webhook_id}/signing-secret/rotate":{"protocol":"http","parameters":[{"name":"webhook_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Webhook ID.","index$":0}]},"POST /webhooks":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","required":["endpoint","events"],"properties":{"endpoint":{"type":"string","description":"The URL where webhook events will be sent.","example":"https://webhook.example.com/handler","key$":"endpoint"},"events":{"type":"array","items":{"type":"string"},"minItems":1,"description":"Array of event types to subscribe to.","example":["email.sent","email.delivered","email.bounced","email.suppressed"],"key$":"events"}},"x-ref":"#/components/schemas/CreateWebhookRequest","index$":1}}}},"parameters":[]},"GET /webhooks":{"protocol":"http","parameters":[{"name":"limit","in":"query","required":false,"schema":{"type":"integer"},"description":"Maximum number of webhooks to return.","index$":0},{"name":"after","in":"query","required":false,"schema":{"type":"string","format":"uuid"},"description":"Pagination cursor to fetch results after this webhook ID. Cannot be used with 'before'.","index$":1},{"name":"before","in":"query","required":false,"schema":{"type":"string","format":"uuid"},"description":"Pagination cursor to fetch results before this webhook ID. Cannot be used with 'after'.","index$":2}]},"GET /webhooks/{webhook_id}":{"protocol":"http","parameters":[{"name":"webhook_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Webhook ID.","index$":0}]},"DELETE /webhooks/{webhook_id}":{"protocol":"http","parameters":[{"name":"webhook_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Webhook ID.","index$":0}]},"PATCH /webhooks/{webhook_id}":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"endpoint":{"type":"string","description":"The URL where webhook events will be sent.","example":"https://webhook.example.com/new-handler","key$":"endpoint"},"events":{"type":"array","items":{"type":"string"},"minItems":1,"description":"Array of event types to subscribe to.","example":["email.sent","email.delivered"],"key$":"events"},"status":{"type":"string","enum":["enabled","disabled"],"description":"The status of the webhook.","example":"enabled","key$":"status"}},"x-ref":"#/components/schemas/UpdateWebhookRequest","index$":1}}}},"parameters":[{"name":"webhook_id","in":"path","required":true,"schema":{"type":"string","format":"uuid"},"description":"The Webhook ID.","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const webhook_ref01_ent = client.Webhook()
    let webhook_ref01_data = setup.data.new.webhook['webhook_ref01']

    webhook_ref01_data = (await webhook_ref01_ent.create(webhook_ref01_data)).data()
    assert(null != webhook_ref01_data.id)


    // LIST
    const webhook_ref01_match: any = {}

    const webhook_ref01_list = (await webhook_ref01_ent.list(webhook_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(webhook_ref01_list, { id: webhook_ref01_data.id })))


    // UPDATE
    const webhook_ref01_data_up0: any = {}
    webhook_ref01_data_up0.id = webhook_ref01_data.id

    const webhook_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-webhook_ref01_' + setup.now }
    ;(webhook_ref01_data_up0 as any)[webhook_ref01_markdef_up0.name] = webhook_ref01_markdef_up0.value

    const webhook_ref01_resdata_up0 = (await webhook_ref01_ent.update(webhook_ref01_data_up0)).data()
    assert(webhook_ref01_resdata_up0.id === webhook_ref01_data_up0.id)

    assert((webhook_ref01_resdata_up0 as any)[webhook_ref01_markdef_up0.name] === webhook_ref01_markdef_up0.value)


    // LOAD
    const webhook_ref01_match_dt0: any = {}
    webhook_ref01_match_dt0.id = webhook_ref01_data.id
    const webhook_ref01_data_dt0 = (await webhook_ref01_ent.load(webhook_ref01_match_dt0)).data()
    assert(webhook_ref01_data_dt0.id === webhook_ref01_data.id)


    // REMOVE
    const webhook_ref01_match_rm0: any = { id: webhook_ref01_data.id }
    await webhook_ref01_ent.remove(webhook_ref01_match_rm0)
  

    // LIST
    const webhook_ref01_match_rt0: any = {}

    const webhook_ref01_list_rt0 = (await webhook_ref01_ent.list(webhook_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(webhook_ref01_list_rt0, { id: webhook_ref01_data.id })))


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
      '../../../../.sdk/test/entity/webhook/WebhookTestData.json')

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
    ['webhook01','webhook02','webhook03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_WEBHOOK_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_WEBHOOK_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_WEBHOOK_ENTID']
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
  
