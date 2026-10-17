

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


describe('TopicEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.Topic()
    assert(null != ent)
  })


  class FailHook extends BaseFeature {
    name = 'failhook'
    version = '0.0.1'
    active = true
    unexpected = 0
    init() { }
    PreSpec() { throw new Error('topic hook failed') }
    PreUnexpected() { this.unexpected++ }
  }

  test('stream-error', async () => {
    const offline = { net: { offline: true } }
    await assert.rejects(async () => {
      for await (const _item of ResendSDK.test(offline).Topic().stream('list')) { }
    }, /offline/)

    for await (const _item of ResendSDK.test(offline).Topic()
      .stream('list', undefined, { ctrl: { throw: false } })) { }

    if (null != (config as any).feature?.rbac) {
      const denied = ResendSDK.test(undefined, { feature: { rbac: { active: true, deny: true } } })
      await assert.rejects(async () => {
        for await (const _item of denied.Topic().stream('list')) { }
      }, (err: any) => 'rbac_denied' === err.code)
    }
  })

  test('stream-ctrl', async () => {
    const explain: any = {}
    const ctrl: any = { explain }
    for await (const _item of ResendSDK.test().Topic().stream('list', undefined, { ctrl })) { }
    assert.deepStrictEqual(Object.keys(ctrl), ['explain'])
    assert(explain === ctrl.explain && 0 < Object.keys(explain).length)
  })

  test('unexpected', async () => {
    const hook = new FailHook()
    const client = new ResendSDK({ feature: { test: { active: true } }, extend: [hook] })
    await assert.rejects(client.Topic().list(), /hook failed/)
    assert(0 < hook.unexpected)

    const fired = hook.unexpected
    assert.strictEqual(await client.Topic().list(undefined, { throw: false }), undefined)
    assert(fired < hook.unexpected)
  })

  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ResendSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Topic().list({"after":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'topic.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"created_at":{"a":true,"h":"Created At","n":"created_at","r":false,"sh":"Timestamp indicating when the topic was created.","t":"`$STRING`","key$":"created_at","index$":0},"default_subscription":{"a":true,"h":"Default Subscription","n":"default_subscription","op":{"create":{"req":true,"type":"`$STRING`"}},"r":false,"sh":"The default subscription status for the topic.","t":"`$STRING`","key$":"default_subscription","index$":1},"description":{"a":true,"h":"Description","n":"description","r":false,"sh":"A description of the topic.","t":"`$STRING`","key$":"description","index$":2},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"The ID of the topic.","t":"`$STRING`","key$":"id","index$":3},"name":{"a":true,"h":"Name","n":"name","op":{"create":{"req":true,"type":"`$STRING`"}},"r":false,"sh":"The name of the topic.","t":"`$STRING`","key$":"name","index$":4},"object":{"a":true,"h":"Object","n":"object","r":false,"sh":"The object type.","t":"`$STRING`","key$":"object","index$":5},"visibility":{"a":true,"h":"Visibility","n":"visibility","r":false,"sh":"The visibility of the topic.","t":"`$STRING`","key$":"visibility","index$":6}},"id":{"field":"id","name":"id"},"name":"topic","op":{"create":{"input":"data","name":"create","points":[{"a":true,"bf":["default_subscription","description","name","visibility"],"co":{"id":"POST /topics","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/topics","q":{},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"topics"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /topics","source":"openapi3","version":2},"g":{"query":[{"a":true,"k":"query","n":"after","or":"after","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"before","or":"before","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":2}]},"k":"http","m":"GET","o":"/topics","q":{},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"topics"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /topics/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/topics/{id}","q":{"exist":["id"]},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"topics"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /topics/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/topics/{id}","q":{"exist":["id"]},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"topics"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"bf":["description","name","visibility"],"co":{"id":"PATCH /topics/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/topics/{id}","q":{"exist":["id"]},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"topics"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"topic","name__orig":"topic","Name":"Topic","name_":"topic","name-":"topic","NAME":"TOPIC","index$":36}, {"active":true,"entity":"topic","key$":"BasicTopicFlow","kind":"basic","name":"BasicTopicFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"topic_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"topic_ref01"}}],"index$":1},{"a":true,"d":{},"i":{"ref":"topic_ref01","srcdatavar":"topic_ref01_data","suffix":"_up0","textfield":"created_at"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-topic_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"topic_ref01","srcdatavar":"topic_ref01_data","suffix":"_dt0"},"m":{"id":"topic01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-topic_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"topic_ref01","suffix":"_rm0"},"m":{"id":"topic01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"topic_ref01"}}],"index$":5}]}, 'Topic', {"POST /topics":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","required":["name","default_subscription"],"properties":{"name":{"type":"string","description":"The name of the topic. Max 50 characters.","maxLength":50,"key$":"name"},"default_subscription":{"type":"string","enum":["opt_in","opt_out"],"description":"The default subscription status for the topic. Cannot be changed after creation.","key$":"default_subscription"},"description":{"type":"string","description":"A description of the topic. Max 200 characters.","maxLength":200,"key$":"description"},"visibility":{"type":"string","enum":["public","private"],"default":"private","description":"The visibility of the topic. Public topics are visible to all contacts on the unsubscribe page. Private topics are only visible to opted-in contacts.","key$":"visibility"}},"x-ref":"#/components/schemas/CreateTopicOptions","index$":1}}}},"parameters":[]},"GET /topics":{"protocol":"http","parameters":[{"in":"query","name":"limit","required":false,"schema":{"type":"integer","minimum":1,"maximum":100},"description":"Number of items to return.","x-ref":"#/components/parameters/PaginationLimit","index$":0},{"in":"query","name":"after","required":false,"schema":{"type":"string"},"description":"Return items after this cursor.","x-ref":"#/components/parameters/PaginationAfter","index$":1},{"in":"query","name":"before","required":false,"schema":{"type":"string"},"description":"Return items before this cursor.","x-ref":"#/components/parameters/PaginationBefore","index$":2}]},"GET /topics/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"string"},"description":"The Topic ID.","index$":0}]},"DELETE /topics/{id}":{"protocol":"http","parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"string"},"description":"The Topic ID.","index$":0}]},"PATCH /topics/{id}":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"name":{"type":"string","description":"The name of the topic. Max 50 characters.","maxLength":50,"key$":"name"},"description":{"type":"string","description":"A description of the topic. Max 200 characters.","maxLength":200,"key$":"description"},"visibility":{"type":"string","enum":["public","private"],"description":"The visibility of the topic.","key$":"visibility"}},"x-ref":"#/components/schemas/UpdateTopicOptions","index$":1}}}},"parameters":[{"name":"id","in":"path","required":true,"schema":{"type":"string"},"description":"The Topic ID.","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const topic_ref01_ent = client.Topic()
    let topic_ref01_data = setup.data.new.topic['topic_ref01']

    topic_ref01_data = (await topic_ref01_ent.create(topic_ref01_data)).data()
    assert(null != topic_ref01_data.id)


    // LIST
    const topic_ref01_match: any = {}

    const topic_ref01_list = (await topic_ref01_ent.list(topic_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(topic_ref01_list, { id: topic_ref01_data.id })))


    // UPDATE
    const topic_ref01_data_up0: any = {}
    topic_ref01_data_up0.id = topic_ref01_data.id

    const topic_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-topic_ref01_' + setup.now }
    ;(topic_ref01_data_up0 as any)[topic_ref01_markdef_up0.name] = topic_ref01_markdef_up0.value

    const topic_ref01_resdata_up0 = (await topic_ref01_ent.update(topic_ref01_data_up0)).data()
    assert(topic_ref01_resdata_up0.id === topic_ref01_data_up0.id)

    assert((topic_ref01_resdata_up0 as any)[topic_ref01_markdef_up0.name] === topic_ref01_markdef_up0.value)


    // LOAD
    const topic_ref01_match_dt0: any = {}
    topic_ref01_match_dt0.id = topic_ref01_data.id
    const topic_ref01_data_dt0 = (await topic_ref01_ent.load(topic_ref01_match_dt0)).data()
    assert(topic_ref01_data_dt0.id === topic_ref01_data.id)


    // REMOVE
    const topic_ref01_match_rm0: any = { id: topic_ref01_data.id }
    await topic_ref01_ent.remove(topic_ref01_match_rm0)
  

    // LIST
    const topic_ref01_match_rt0: any = {}

    const topic_ref01_list_rt0 = (await topic_ref01_ent.list(topic_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(topic_ref01_list_rt0, { id: topic_ref01_data.id })))


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
      '../../../../.sdk/test/entity/topic/TopicTestData.json')

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
    ['topic01','topic02','topic03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_TOPIC_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_TOPIC_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_TOPIC_ENTID']
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
  
