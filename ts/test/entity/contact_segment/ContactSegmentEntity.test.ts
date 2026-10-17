

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


describe('ContactSegmentEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.ContactSegment()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ResendSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.ContactSegment().list({"contact_id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of ['list', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'contact_segment.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"contact_id":{"a":true,"h":"Contact Id","n":"contact_id","r":false,"sh":"The ID of the contact.","t":"`$STRING`","key$":"contact_id","index$":0},"created_at":{"a":true,"h":"Created At","n":"created_at","r":false,"sh":"Timestamp indicating when the contact was added to the segment.","t":"`$STRING`","key$":"created_at","index$":1},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Unique identifier for the segment.","t":"`$STRING`","key$":"id","index$":2},"name":{"a":true,"h":"Name","n":"name","r":false,"sh":"Name of the segment.","t":"`$STRING`","key$":"name","index$":3},"object":{"a":true,"h":"Object","n":"object","r":false,"sh":"The object type.","t":"`$STRING`","key$":"object","index$":4},"segment_id":{"a":true,"h":"Segment Id","n":"segment_id","r":false,"sh":"The ID of the segment.","t":"`$STRING`","key$":"segment_id","index$":5}},"id":{"field":"id","name":"id"},"name":"contact_segment","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /contacts/{contact_id}/segments/{segment_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"contact_id","or":"contact_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"id","or":"segment_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"POST","o":"/contacts/{contact_id}/segments/{segment_id}","q":{"exist":["contact_id","id"]},"r":{"param":{"segment_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"contacts"},{"var":"contact_id"},{"lit":"segments"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /contacts/{contact_id}/segments","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"contact_id","or":"contact_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"after","or":"after","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"before","or":"before","r":false,"t":"`$STRING`","index$":1},{"a":true,"k":"query","n":"limit","or":"limit","r":false,"t":"`$INTEGER`","index$":2}]},"k":"http","m":"GET","o":"/contacts/{contact_id}/segments","q":{"exist":["contact_id"]},"r":{},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"contacts"},{"var":"contact_id"},{"lit":"segments"}],"t":{"req":"`reqdata`","res":"`body.data`"},"index$":0}],"key$":"list"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /contacts/{contact_id}/segments/{segment_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"contact_id","or":"contact_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"id","or":"segment_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/contacts/{contact_id}/segments/{segment_id}","q":{"exist":["contact_id","id"]},"r":{"param":{"segment_id":"id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"contacts"},{"var":"contact_id"},{"lit":"segments"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"}},"relations":{"ancestors":[["$.main.kit.entity.contact"]]},"key$":"contact_segment","name__orig":"contact_segment","Name":"ContactSegment","name_":"contact_segment","name-":"contact-segment","NAME":"CONTACT_SEGMENT","index$":11}, {"active":true,"entity":"contact_segment","key$":"BasicContactSegmentFlow","kind":"basic","name":"BasicContactSegmentFlow","param":{},"step":[{"a":false,"d":{},"i":{"ref":"contact_segment_ref01"},"m":{"contact_id":"contact01","segment_id":"segment01"},"o":"create","s":[],"v":[],"unreachable":true},{"a":true,"d":{},"i":{},"m":{"contact_id":"contact01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"contact_segment_ref01"}}],"index$":0},{"a":true,"d":{},"i":{"ref":"contact_segment_ref01","suffix":"_rm0"},"m":{"contact_id":"contact01","id":"contact_segment01"},"o":"remove","s":[],"v":[],"index$":1},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{"contact_id":"contact01"},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"contact_segment_ref01"}}],"index$":2}]}, 'ContactSegment', {"POST /contacts/{contact_id}/segments/{segment_id}":{"protocol":"http","parameters":[{"name":"contact_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Contact ID or email address.","index$":0},{"name":"segment_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Segment ID.","index$":1}]},"GET /contacts/{contact_id}/segments":{"protocol":"http","parameters":[{"name":"contact_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Contact ID or email address.","index$":0},{"in":"query","name":"limit","required":false,"schema":{"type":"integer","minimum":1,"maximum":100},"description":"Number of items to return.","x-ref":"#/components/parameters/PaginationLimit","index$":1},{"in":"query","name":"after","required":false,"schema":{"type":"string"},"description":"Return items after this cursor.","x-ref":"#/components/parameters/PaginationAfter","index$":2},{"in":"query","name":"before","required":false,"schema":{"type":"string"},"description":"Return items before this cursor.","x-ref":"#/components/parameters/PaginationBefore","index$":3}]},"DELETE /contacts/{contact_id}/segments/{segment_id}":{"protocol":"http","parameters":[{"name":"contact_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Contact ID or email address.","index$":0},{"name":"segment_id","in":"path","required":true,"schema":{"type":"string"},"description":"The Segment ID.","index$":1}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let contact_segment_ref01_data = Object.values(setup.data.existing.contact_segment)[0] as any

    // LIST
    const contact_segment_ref01_ent = client.ContactSegment()
    const contact_segment_ref01_match: any = {}
    contact_segment_ref01_match['contact_id'] = setup.idmap['contact01']

    const contact_segment_ref01_list = (await contact_segment_ref01_ent.list(contact_segment_ref01_match)).map((e: any) => e.data())


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
      '../../../../.sdk/test/entity/contact_segment/ContactSegmentTestData.json')

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
    ['contact_segment01','contact_segment02','contact_segment03','contact01','contact02','contact03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_CONTACT_SEGMENT_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_CONTACT_SEGMENT_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_CONTACT_SEGMENT_ENTID']
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
  
