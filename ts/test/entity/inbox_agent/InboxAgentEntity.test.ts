

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


describe('InboxAgentEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
  afterEach(liveDelay('RESEND_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ResendSDK.test()
    const ent = testsdk.InboxAgent()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = ResendSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.InboxAgent().load({"inbox_id":1} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.RESEND_TEST_LIVE
    for (const op of []) {
      if (!live && maybeSkipControl(t, 'entityOp', 'inbox_agent.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"actions":{"a":true,"h":"Actions","n":"actions","op":{"update":{"req":false,"type":"`$ARRAY`"}},"r":true,"sh":"The actions the agent can take, in the order of the enum.","t":"`$ARRAY`","key$":"actions","index$":0},"id":{"a":true,"fo":"uuid","h":"Id","n":"id","r":true,"sh":"The ID of the inbox agent.","t":"`$STRING`","key$":"id","index$":1},"instructions":{"a":true,"h":"Instructions","n":"instructions","op":{"update":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]}},"r":true,"sh":"Instructions the agent follows when it handles the threads of the inbox.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"instructions","index$":2},"object":{"a":true,"h":"Object","n":"object","r":true,"sh":"The object type.","t":"`$STRING`","key$":"object","index$":3},"tone":{"a":true,"h":"Tone","n":"tone","op":{"update":{"req":false,"type":["`$ONE`",["`$STRING`","`$NULL`"]]}},"r":true,"sh":"The tone the agent writes in.","t":["`$ONE`",["`$STRING`","`$NULL`"]],"key$":"tone","index$":4}},"id":{"field":"id","name":"id"},"name":"inbox_agent","op":{"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /inboxes/{inbox_id}/agent","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/inboxes/{inbox_id}/agent","q":{"exist":["inbox_id"]},"r":{"param":{"inbox_id":"inbox_id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"agent"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"update":{"input":"data","name":"update","points":[{"a":true,"bf":["actions","instructions","tone"],"co":{"id":"PATCH /inboxes/{inbox_id}/agent","source":"openapi3","version":2},"g":{"params":[{"a":true,"ex":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","k":"param","n":"inbox_id","or":"inbox_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/inboxes/{inbox_id}/agent","q":{"exist":["inbox_id"]},"r":{"param":{"inbox_id":"inbox_id"}},"rs":{"kind":"json","media":"application/json"},"s":[{"lit":"inboxes"},{"var":"inbox_id"},{"lit":"agent"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[["$.main.kit.entity.inbox"]]},"key$":"inbox_agent","name__orig":"inbox_agent","Name":"InboxAgent","name_":"inbox_agent","name-":"inbox-agent","NAME":"INBOX_AGENT","index$":20}, {"active":true,"entity":"inbox_agent","key$":"BasicInboxAgentFlow","kind":"basic","name":"BasicInboxAgentFlow","param":{},"step":[{"a":false,"d":{},"i":{"ref":"inbox_agent_ref01","srcdatavar":"inbox_agent_ref01_data","suffix":"_up0","textfield":"object"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-inbox_agent_ref01"}}],"v":[],"unreachable":true},{"a":false,"d":{},"i":{"ref":"inbox_agent_ref01","srcdatavar":"inbox_agent_ref01_data","suffix":"_dt0"},"m":{"id":"inbox_agent01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-inbox_agent_ref01"}}],"unreachable":true}]}, 'InboxAgent', {"GET /inboxes/{inbox_id}/agent":{"protocol":"http","parameters":[{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":0}]},"PATCH /inboxes/{inbox_id}/agent":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"type":"object","anyOf":[{"required":["instructions"]},{"required":["tone"]},{"required":["actions"]}],"description":"At least one of `instructions`, `tone`, or `actions` must be provided. Omitted fields keep their current value.","properties":{"instructions":{"type":["string","null"],"minLength":1,"maxLength":4000,"description":"Instructions the agent follows when it handles the threads of the inbox. Max 4000 characters. Leading and trailing whitespace is trimmed, and the trimmed value must not be empty. Send null to clear the instructions.","example":"Answer refund questions yourself. Escalate legal threats.","key$":"instructions"},"tone":{"type":["string","null"],"enum":["concise_professional","concise_casual","expanded_professional","expanded_casual",null],"description":"The tone the agent writes in. Send null to clear the tone.","example":"concise_professional","key$":"tone"},"actions":{"type":"array","description":"The actions the agent can take. Replaces the whole set. Duplicates are removed. Send an empty array to disable all actions. The old names `archive_thread` and `delete_thread` return a 422. Use `archive` and `move_to_trash`.","items":{"type":"string","enum":["draft_reply","forward_thread","add_labels","assign_thread","mark_as_spam","archive","move_to_trash"]},"example":["draft_reply","add_labels"],"key$":"actions"}},"x-ref":"#/components/schemas/UpdateInboxAgentOptions","index$":1}}}},"parameters":[{"in":"path","name":"inbox_id","required":true,"schema":{"type":"string"},"description":"The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.","example":"b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1","x-ref":"#/components/parameters/InboxIdOrEmailAddress","index$":0}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let inbox_agent_ref01_data = Object.values(setup.data.existing.inbox_agent)[0] as any

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
      '../../../../.sdk/test/entity/inbox_agent/InboxAgentTestData.json')

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
    ['inbox_agent01','inbox_agent02','inbox_agent03','inbox01','inbox02','inbox03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'RESEND_TEST_INBOX_AGENT_ENTID': idmap,
    'RESEND_TEST_LIVE': 'FALSE',
    'RESEND_TEST_EXPLAIN': 'FALSE',
    'RESEND_APIKEY': '',
  })

  idmap = env['RESEND_TEST_INBOX_AGENT_ENTID']

  const live = 'TRUE' === env.RESEND_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['RESEND_TEST_INBOX_AGENT_ENTID']
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
  
