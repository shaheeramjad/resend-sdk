

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


describe('ContactTopicEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.ContactTopic()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ResendSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.ContactTopic().list({"contact_id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'contact_topic.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"contact_id":{"a":true,"h":"Contact Id","n":"contact_id","r":false,"sh":"The ID of the contact.","t":"`$STRING`","key$":"contact_id","index$":0},"description":{"a":true,"h":"Description","n":"description","r":false,"sh":"Description of the topic.","t":"`$STRING`","key$":"description","index$":1},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier for the topic.","t":"`$STRING`","key$":"id","index$":2},"name":{"a":true,"h":"Name","n":"name","r":false,"sh":"Name of the topic.","t":"`$STRING`","key$":"name","index$":3},"object":{"a":true,"h":"Object","n":"object","r":false,"sh":"The object type.","t":"`$STRING`","key$":"object","index$":4},"subscription":{"a":true,"h":"Subscription","n":"subscription","r":false,"sh":"The subscription status for this topic.","t":"`$STRING`","key$":"subscription","index$":5},"topics":{"a":true,"h":"Topics","n":"topics","op":{"update":{"req":true,"type":"`$ARRAY`"}},"r":false,"sh":"Array of updated topic subscriptions.","t":"`$ARRAY`","key$":"topics","index$":6}},"id":{"field":"id","name":"id"},"name":"contact_topic","op":{"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /contacts/{contact_id}/topics","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"contact_id","or":"contact_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"after","or":"after","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"before","or":"before","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":2}]},"k":"http","m":"GET","o":"/contacts/{contact_id}/topics","q":{"exist":["contact_id"]},"r":{"param":{"contact_id":"contact_id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"contacts"},{"var":"contact_id"},{"lit":"topics"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"},"update":{"input":"data","name":"update","points":[{"a":true,"bf":["topics"],"co":{"id":"PATCH /contacts/{contact_id}/topics","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"contact_id","or":"contact_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/contacts/{contact_id}/topics","q":{"exist":["contact_id"]},"r":{"param":{"contact_id":"contact_id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"contacts"},{"var":"contact_id"},{"lit":"topics"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[["$.main.kit.entity.contact"]]},"key$":"contact_topic","name__orig":"contact_topic","Name":"ContactTopic","name_":"contact_topic","name-":"contact-topic","NAME":"CONTACT_TOPIC","index$":12}, {"active":true,"entity":"contact_topic","key$":"BasicContactTopicFlow","kind":"basic","name":"BasicContactTopicFlow","param":{},"step":[{"a":true,"d":{},"i":{},"m":{"contact_id":"contact01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"contact_topic_ref01"}}],"index$":0},{"a":false,"d":{},"i":{"ref":"contact_topic_ref01","srcdatavar":"contact_topic_ref01_data","suffix":"_up0","textfield":"description"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-contact_topic_ref01"}}],"v":[],"unreachable":true}]}, 'ContactTopic', {"GET /contacts/{contact_id}/topics":{"protocol":"http","parameters":[{"name":"contact_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Contact ID or email address.","index$":0},{"in":"query","name":"limit","required":false,"schema":{"type":"integer","minimum":1,"maximum":100},"description":"Number of items to return.","x-ref":"#/components/parameters/PaginationLimit","index$":1},{"in":"query","name":"after","required":false,"schema":{"type":"string"},"description":"Return items after this cursor.","x-ref":"#/components/parameters/PaginationAfter","index$":2},{"in":"query","name":"before","required":false,"schema":{"type":"string"},"description":"Return items before this cursor.","x-ref":"#/components/parameters/PaginationBefore","index$":3}]},"PATCH /contacts/{contact_id}/topics":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","required":["topics"],"properties":{"topics":{"type":"array","items":{"type":"object","properties":{"id":{"type":"string","description":"The ID of the topic."},"subscription":{"type":"string","enum":[],"description":"The subscription status (opt_in or opt_out)."}}},"key$":"topics"}},"x-ref":"#/components/schemas/UpdateContactTopicsOptions","index$":1}}}},"parameters":[{"name":"contact_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Contact ID or email address.","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let contact_topic_ref01_data = Object.values(setup.data.existing.contact_topic)[0] as any

    // LIST
    const contact_topic_ref01_ent = client.ContactTopic()
    const contact_topic_ref01_match: any = {}
    contact_topic_ref01_match['contact_id'] = setup.idmap['contact01']

    const contact_topic_ref01_list = (await contact_topic_ref01_ent.list(contact_topic_ref01_match)).map((e: any) => e.data())


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
      '../../../../.sdk/test/entity/contact_topic/ContactTopicTestData.json')

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
    ['contact_topic01','contact_topic02','contact_topic03','contact01','contact02','contact03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_CONTACT_TOPIC_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_CONTACT_TOPIC_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_CONTACT_TOPIC_ENTID']
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
  
