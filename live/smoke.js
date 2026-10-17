// Live smoke test of the generated TypeScript SDK against api.resend.com.
//
//   cd ts && npm install && npm run build && cd ..
//   RESEND_APIKEY=re_... node live/smoke.js
//
// Needs a FULL ACCESS key (a sending-only key cannot list domains or keys).
// Sends one email from onboarding@resend.dev to Resend's sandbox inbox
// delivered@resend.dev unless RESEND_SMOKE_TO names another address.
// Every record it creates (one contact, one segment) is removed at the end.

const { ResendSDK } = require('../ts/dist/ResendSDK.js')

const apikey = process.env.RESEND_APIKEY
if (!apikey) {
  console.error('Set RESEND_APIKEY to a Resend API key (https://resend.com/api-keys).')
  process.exit(2)
}

const to = process.env.RESEND_SMOKE_TO || 'delivered@resend.dev'
const stamp = Date.now()

const client = new ResendSDK({
  apikey,
  feature: {
    retry: { active: true, retries: 3, maxDelay: 4000 },
    timeout: { active: true, ms: 15000 },
    ratelimit: { active: true, rate: 2, burst: 2 },
    metrics: { active: true },
    debug: { active: true, max: 200 },
  },
})

const results = []
const cleanup = []

async function step(name, fn) {
  const start = Date.now()
  try {
    const detail = await fn()
    results.push({ name, ok: true, ms: Date.now() - start, detail })
    console.log(`PASS  ${name}${detail ? '  ' + detail : ''}`)
  } catch (err) {
    results.push({ name, ok: false, ms: Date.now() - start, detail: err.message })
    console.log(`FAIL  ${name}  ${err.message}`)
  }
}

function must(cond, msg) {
  if (!cond) throw new Error(msg)
}

async function main() {
  await step('Usage.load', async () => {
    const usage = (await client.Usage().load()).data()
    must(usage && 'object' === typeof usage, 'no usage record')
    return JSON.stringify(usage).slice(0, 80)
  })

  await step('Domain.list', async () => {
    const domains = await client.Domain().list()
    must(Array.isArray(domains), 'list did not return an array')
    return `${domains.length} domain(s)`
  })

  await step('ApiKey.list', async () => {
    const keys = await client.ApiKey().list()
    must(Array.isArray(keys) && keys.length > 0, 'expected at least the key in use')
    return `${keys.length} key(s)`
  })

  let emailId
  await step('Email.create (send)', async () => {
    const sent = (await client.Email().create({
      from: 'Voxgig Resend SDK <onboarding@resend.dev>',
      to: [to],
      subject: `Resend SDK smoke test ${stamp}`,
      html: '<p>Sent by the generated Voxgig Resend SDK smoke test.</p>',
    })).data()
    must(sent.id, 'no id in send response')
    emailId = sent.id
    return `id=${emailId} to=${to}`
  })

  await step('Email.load', async () => {
    must(emailId, 'skipped: no email was sent')
    const email = (await client.Email().load({ id: emailId })).data()
    must(email.id === emailId, `loaded id ${email.id} != ${emailId}`)
    return `last_event=${email.last_event}`
  })

  await step('Email.list', async () => {
    const emails = await client.Email().list()
    must(Array.isArray(emails), 'list did not return an array')
    return `${emails.length} email(s)`
  })

  let contactId
  const contactEmail = `sdk-smoke-${stamp}@example.com`
  await step('Contact.create', async () => {
    const created = (await client.Contact().create({
      email: contactEmail,
      first_name: 'Smoke',
    })).data()
    must(created.id, 'no id in create response')
    contactId = created.id
    cleanup.push(['Contact.remove', () => client.Contact().remove({ id: contactId })])
    return `id=${contactId}`
  })

  await step('Contact.update', async () => {
    must(contactId, 'skipped: no contact was created')
    await client.Contact().update({ id: contactId, first_name: 'Smoked' })
  })

  await step('Contact.load', async () => {
    must(contactId, 'skipped: no contact was created')
    const contact = (await client.Contact().load({ id: contactId })).data()
    must(contact.email === contactEmail, `email ${contact.email} != ${contactEmail}`)
    must(contact.first_name === 'Smoked', `first_name ${contact.first_name} != Smoked`)
    return `first_name=${contact.first_name}`
  })

  await step('Contact.list', async () => {
    const contacts = await client.Contact().list()
    must(Array.isArray(contacts), 'list did not return an array')
    return `${contacts.length} contact(s)`
  })

  let segmentId
  await step('Segment.create', async () => {
    const created = (await client.Segment().create({ name: `sdk-smoke-${stamp}` })).data()
    must(created.id, 'no id in create response')
    segmentId = created.id
    cleanup.push(['Segment.remove', () => client.Segment().remove({ id: segmentId })])
    return `id=${segmentId}`
  })

  await step('Segment.load', async () => {
    must(segmentId, 'skipped: no segment was created')
    const segment = (await client.Segment().load({ id: segmentId })).data()
    must(segment.id === segmentId, `loaded id ${segment.id} != ${segmentId}`)
  })

  for (const name of ['Template', 'Webhook', 'Topic', 'Broadcast', 'Log']) {
    await step(`${name}.list`, async () => {
      const items = await client[name]().list()
      must(Array.isArray(items), 'list did not return an array')
      return `${items.length} item(s)`
    })
  }

  await step('direct GET /domains', async () => {
    const res = await client.direct({ path: '/domains', method: 'GET' })
    must(res.ok, `status ${res.status}`)
    return `status=${res.status}`
  })

  await step('bad key is rejected', async () => {
    const bad = new ResendSDK({ apikey: 're_invalid_key_for_smoke_test' })
    try {
      await bad.Domain().list()
    } catch (err) {
      return `rejected: ${err.message}`
    }
    throw new Error('an invalid key was accepted')
  })

  for (const [name, fn] of cleanup.reverse()) {
    await step(name, async () => { await fn() })
  }

  const failed = results.filter((r) => !r.ok)
  const total = client._metrics?.total
  console.log('')
  console.log(`${results.length - failed.length}/${results.length} steps passed`)
  if (total) {
    console.log(`metrics: ${total.count} ops, ${total.ok} ok, ${total.err} err, max ${total.maxMs} ms`)
  }
  process.exit(failed.length ? 1 : 0)
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
