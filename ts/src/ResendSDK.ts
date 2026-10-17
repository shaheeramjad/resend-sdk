// Resend Ts SDK

import { ApiKeyEntity } from './entity/ApiKeyEntity'
import { AttachmentEntity } from './entity/AttachmentEntity'
import { AudienceEntity } from './entity/AudienceEntity'
import { AutomationEntity } from './entity/AutomationEntity'
import { AutomationRunEntity } from './entity/AutomationRunEntity'
import { BroadcastEntity } from './entity/BroadcastEntity'
import { BroadcastClickedLinkEntity } from './entity/BroadcastClickedLinkEntity'
import { BroadcastRecipientEntity } from './entity/BroadcastRecipientEntity'
import { ContactEntity } from './entity/ContactEntity'
import { ContactImportEntity } from './entity/ContactImportEntity'
import { ContactPropertyEntity } from './entity/ContactPropertyEntity'
import { ContactSegmentEntity } from './entity/ContactSegmentEntity'
import { ContactTopicEntity } from './entity/ContactTopicEntity'
import { DomainEntity } from './entity/DomainEntity'
import { DomainClaimEntity } from './entity/DomainClaimEntity'
import { EmailEntity } from './entity/EmailEntity'
import { EmailBatchEntity } from './entity/EmailBatchEntity'
import { EmailsMetricEntity } from './entity/EmailsMetricEntity'
import { EventEntity } from './entity/EventEntity'
import { InboxEntity } from './entity/InboxEntity'
import { InboxAgentEntity } from './entity/InboxAgentEntity'
import { InboxDraftEntity } from './entity/InboxDraftEntity'
import { InboxLabelEntity } from './entity/InboxLabelEntity'
import { InboxThreadEntity } from './entity/InboxThreadEntity'
import { InboxThreadEmailEntity } from './entity/InboxThreadEmailEntity'
import { InboxThreadLabelEntity } from './entity/InboxThreadLabelEntity'
import { LogEntity } from './entity/LogEntity'
import { OauthGrantEntity } from './entity/OauthGrantEntity'
import { ReceivedAttachmentEntity } from './entity/ReceivedAttachmentEntity'
import { ReceivedEmailEntity } from './entity/ReceivedEmailEntity'
import { SegmentEntity } from './entity/SegmentEntity'
import { SegmentContactEntity } from './entity/SegmentContactEntity'
import { SuppressionEntity } from './entity/SuppressionEntity'
import { SuppressionBatchAddEntity } from './entity/SuppressionBatchAddEntity'
import { SuppressionBatchRemoveEntity } from './entity/SuppressionBatchRemoveEntity'
import { TemplateEntity } from './entity/TemplateEntity'
import { TopicEntity } from './entity/TopicEntity'
import { UsageEntity } from './entity/UsageEntity'
import { WebhookEntity } from './entity/WebhookEntity'
import { WebhookEventEntity } from './entity/WebhookEventEntity'
import { WebhookEventAttemptEntity } from './entity/WebhookEventAttemptEntity'

export type * from './ResendTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { ResendEntityBase } from './ResendEntityBase'
import { Utility } from './utility/Utility'
import { unreadableBody } from './utility/ResultBodyUtility'
import { abortError } from './utility/MakeRequestUtility'
import { allowed } from './utility/PrepareMethodUtility'


import { BaseFeature } from './feature/base/BaseFeature'



const stdutil = new Utility()


// A request's outcome: ok is false alone on an error with no response, so a
// caller narrowing on it reaches the status and the data.
type DirectResult =
  | { ok: false, err: any, status?: undefined, headers?: undefined, data?: undefined }
  | { ok: boolean, status: number, headers: any, data: any, err?: any }


class ResendSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context
  

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    for (const key of ['_options', '_rootctx', '_features']) {
      Object.defineProperty(this, key, {
        value: (this as any)[key], enumerable: false, writable: true, configurable: true
      })
    }

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const extend = this._options.extend || []

    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        // An active name with no generated class is legal when an
        // extend-supplied instance carries that name (station's adopt
        // path): the instance is added below, positioned by its own
        // __after__ entry, so skip it here rather than fail construction.
        if (!this._rootctx.config.hasFeature(fname) &&
          extend.some((f: any) => fname === f.name)) {
          continue
        }
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    for (let f of extend) {
      featureAdd(this._rootctx, f)
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }

  


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options
    const method = String(fetchargs.method || 'GET').toUpperCase()

    if (!allowed(options.allow.method, method)) {
      return ctx.error('spec_method_allow', 'Method "' + method +
        '" not allowed by SDK option allow.method value: "' + options.allow.method + '"')
    }

    const spec: any = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method,
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    

    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  // Raw endpoint access is operator-controllable, like every entity op.
  // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  // either one reaches the same endpoint.
  async direct(fetchargs?: any): Promise<DirectResult> {
    if (!allowed(this._options.allow.op, 'direct')) {
      return {
        ok: false,
        err: new Error('ResendSDK: direct: operation not allowed by' +
          ' SDK option allow.op value: "' + this._options.allow.op + '"'),
      }
    }

    return this._rawRequest(fetchargs)
  }


  // Ungated request path shared by direct() and graphql(), each of which
  // checks its own allow.op token first. Private, rather than a flag on
  // fetchargs: a caller-supplied marker would let anyone opt straight back
  // out of the gate by passing it.
  async _rawRequest(fetchargs?: any): Promise<DirectResult> {
    const utility = this._utility

    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return { ok: false, err: utility.clean(this._rootctx, fetchdef) }
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      if (true === fetchdef.signal?.aborted) {
        throw fetchdef.signal.reason
      }

      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: utility.clean(ctx, abortError(ctx, fetched)) }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      let err: any = undefined
      if (!noBody) {
        let text: any = undefined
        try {
          const raw: any = fetched
          if ('function' === typeof raw.text) {
            text = await raw.text()
            json = '' === text.trim() ? undefined : JSON.parse(text)
          }
          else {
            json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
          }
        }
        catch (parseErr: any) {
          if ('SyntaxError' !== parseErr?.name) {
            throw parseErr
          }
          err = unreadableBody(ctx, {
            status, headers, text: text ?? parseErr.text, sent: fetchdef.headers,
            failed: 200 <= status && status < 300 ? undefined :
              ctx.error('request_status', 'request: ' + status + ': ' + fetched.statusText),
          })
        }
      }

      return {
        ok: null == err && status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
        ...(null == err ? {} : { err: utility.clean(ctx, err) }),
      }
    }
    catch (err: any) {
      return { ok: false, err: utility.clean(ctx, abortError(ctx, err)) }
    }
  }



  async graphql(query: string, variables?: any, ctrl?: any) {
    const options = this._options

    if (!allowed(options.allow.op, 'graphql')) {
      return {
        ok: false,
        err: new Error('ResendSDK: graphql: operation not allowed by' +
          ' SDK option allow.op value: "' + options.allow.op + '"'),
      }
    }

    const res: any = await this._rawRequest({
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: { query, variables: variables || {} },
      ctrl,
    })

    // Errors are read BEFORE any status check: a GraphQL parse or validation
    // failure comes back as HTTP 400 carrying the standard { errors: [...] }
    // body, and the raw path represents a non-2xx as { ok: false } with no
    // err — so returning early on status would discard the server's own
    // diagnostics, which are the only useful part of that response.
    const errors = null == res.data ? undefined : res.data.errors

    if (null != errors && Array.isArray(errors) && 0 < errors.length) {
      const first = errors[0] || {}
      const err: any = new Error('ResendSDK: graphql: ' +
        (first.message || 'graphql error'))
      err.graphql = errors
      return { ok: false, status: res.status, headers: res.headers, err, data: res.data }
    }

    return res
  }



  // Entity access: `client.ApiKey().list()` / `client.ApiKey().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ApiKey(entopts?: Record<string, any>) {
    const self = this
    return new ApiKeyEntity(self, entopts)
  }


  // Entity access: `client.Attachment().list()` / `client.Attachment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Attachment(entopts?: Record<string, any>) {
    const self = this
    return new AttachmentEntity(self, entopts)
  }


  // Entity access: `client.Audience().list()` / `client.Audience().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Audience(entopts?: Record<string, any>) {
    const self = this
    return new AudienceEntity(self, entopts)
  }


  // Entity access: `client.Automation().list()` / `client.Automation().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Automation(entopts?: Record<string, any>) {
    const self = this
    return new AutomationEntity(self, entopts)
  }


  // Entity access: `client.AutomationRun().list()` / `client.AutomationRun().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  AutomationRun(entopts?: Record<string, any>) {
    const self = this
    return new AutomationRunEntity(self, entopts)
  }


  // Entity access: `client.Broadcast().list()` / `client.Broadcast().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Broadcast(entopts?: Record<string, any>) {
    const self = this
    return new BroadcastEntity(self, entopts)
  }


  // Entity access: `client.BroadcastClickedLink().list()` / `client.BroadcastClickedLink().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  BroadcastClickedLink(entopts?: Record<string, any>) {
    const self = this
    return new BroadcastClickedLinkEntity(self, entopts)
  }


  // Entity access: `client.BroadcastRecipient().list()` / `client.BroadcastRecipient().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  BroadcastRecipient(entopts?: Record<string, any>) {
    const self = this
    return new BroadcastRecipientEntity(self, entopts)
  }


  // Entity access: `client.Contact().list()` / `client.Contact().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Contact(entopts?: Record<string, any>) {
    const self = this
    return new ContactEntity(self, entopts)
  }


  // Entity access: `client.ContactImport().list()` / `client.ContactImport().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactImport(entopts?: Record<string, any>) {
    const self = this
    return new ContactImportEntity(self, entopts)
  }


  // Entity access: `client.ContactProperty().list()` / `client.ContactProperty().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactProperty(entopts?: Record<string, any>) {
    const self = this
    return new ContactPropertyEntity(self, entopts)
  }


  // Entity access: `client.ContactSegment().list()` / `client.ContactSegment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactSegment(entopts?: Record<string, any>) {
    const self = this
    return new ContactSegmentEntity(self, entopts)
  }


  // Entity access: `client.ContactTopic().list()` / `client.ContactTopic().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactTopic(entopts?: Record<string, any>) {
    const self = this
    return new ContactTopicEntity(self, entopts)
  }


  // Entity access: `client.Domain().list()` / `client.Domain().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Domain(entopts?: Record<string, any>) {
    const self = this
    return new DomainEntity(self, entopts)
  }


  // Entity access: `client.DomainClaim().list()` / `client.DomainClaim().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  DomainClaim(entopts?: Record<string, any>) {
    const self = this
    return new DomainClaimEntity(self, entopts)
  }


  // Entity access: `client.Email().list()` / `client.Email().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Email(entopts?: Record<string, any>) {
    const self = this
    return new EmailEntity(self, entopts)
  }


  // Entity access: `client.EmailBatch().list()` / `client.EmailBatch().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailBatch(entopts?: Record<string, any>) {
    const self = this
    return new EmailBatchEntity(self, entopts)
  }


  // Entity access: `client.EmailsMetric().list()` / `client.EmailsMetric().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EmailsMetric(entopts?: Record<string, any>) {
    const self = this
    return new EmailsMetricEntity(self, entopts)
  }


  // Entity access: `client.Event().list()` / `client.Event().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Event(entopts?: Record<string, any>) {
    const self = this
    return new EventEntity(self, entopts)
  }


  // Entity access: `client.Inbox().list()` / `client.Inbox().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Inbox(entopts?: Record<string, any>) {
    const self = this
    return new InboxEntity(self, entopts)
  }


  // Entity access: `client.InboxAgent().list()` / `client.InboxAgent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InboxAgent(entopts?: Record<string, any>) {
    const self = this
    return new InboxAgentEntity(self, entopts)
  }


  // Entity access: `client.InboxDraft().list()` / `client.InboxDraft().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InboxDraft(entopts?: Record<string, any>) {
    const self = this
    return new InboxDraftEntity(self, entopts)
  }


  // Entity access: `client.InboxLabel().list()` / `client.InboxLabel().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InboxLabel(entopts?: Record<string, any>) {
    const self = this
    return new InboxLabelEntity(self, entopts)
  }


  // Entity access: `client.InboxThread().list()` / `client.InboxThread().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InboxThread(entopts?: Record<string, any>) {
    const self = this
    return new InboxThreadEntity(self, entopts)
  }


  // Entity access: `client.InboxThreadEmail().list()` / `client.InboxThreadEmail().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InboxThreadEmail(entopts?: Record<string, any>) {
    const self = this
    return new InboxThreadEmailEntity(self, entopts)
  }


  // Entity access: `client.InboxThreadLabel().list()` / `client.InboxThreadLabel().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InboxThreadLabel(entopts?: Record<string, any>) {
    const self = this
    return new InboxThreadLabelEntity(self, entopts)
  }


  // Entity access: `client.Log().list()` / `client.Log().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Log(entopts?: Record<string, any>) {
    const self = this
    return new LogEntity(self, entopts)
  }


  // Entity access: `client.OauthGrant().list()` / `client.OauthGrant().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OauthGrant(entopts?: Record<string, any>) {
    const self = this
    return new OauthGrantEntity(self, entopts)
  }


  // Entity access: `client.ReceivedAttachment().list()` / `client.ReceivedAttachment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ReceivedAttachment(entopts?: Record<string, any>) {
    const self = this
    return new ReceivedAttachmentEntity(self, entopts)
  }


  // Entity access: `client.ReceivedEmail().list()` / `client.ReceivedEmail().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ReceivedEmail(entopts?: Record<string, any>) {
    const self = this
    return new ReceivedEmailEntity(self, entopts)
  }


  // Entity access: `client.Segment().list()` / `client.Segment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Segment(entopts?: Record<string, any>) {
    const self = this
    return new SegmentEntity(self, entopts)
  }


  // Entity access: `client.SegmentContact().list()` / `client.SegmentContact().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SegmentContact(entopts?: Record<string, any>) {
    const self = this
    return new SegmentContactEntity(self, entopts)
  }


  // Entity access: `client.Suppression().list()` / `client.Suppression().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Suppression(entopts?: Record<string, any>) {
    const self = this
    return new SuppressionEntity(self, entopts)
  }


  // Entity access: `client.SuppressionBatchAdd().list()` / `client.SuppressionBatchAdd().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SuppressionBatchAdd(entopts?: Record<string, any>) {
    const self = this
    return new SuppressionBatchAddEntity(self, entopts)
  }


  // Entity access: `client.SuppressionBatchRemove().list()` / `client.SuppressionBatchRemove().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SuppressionBatchRemove(entopts?: Record<string, any>) {
    const self = this
    return new SuppressionBatchRemoveEntity(self, entopts)
  }


  // Entity access: `client.Template().list()` / `client.Template().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Template(entopts?: Record<string, any>) {
    const self = this
    return new TemplateEntity(self, entopts)
  }


  // Entity access: `client.Topic().list()` / `client.Topic().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Topic(entopts?: Record<string, any>) {
    const self = this
    return new TopicEntity(self, entopts)
  }


  // Entity access: `client.Usage().list()` / `client.Usage().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Usage(entopts?: Record<string, any>) {
    const self = this
    return new UsageEntity(self, entopts)
  }


  // Entity access: `client.Webhook().list()` / `client.Webhook().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Webhook(entopts?: Record<string, any>) {
    const self = this
    return new WebhookEntity(self, entopts)
  }


  // Entity access: `client.WebhookEvent().list()` / `client.WebhookEvent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  WebhookEvent(entopts?: Record<string, any>) {
    const self = this
    return new WebhookEventEntity(self, entopts)
  }


  // Entity access: `client.WebhookEventAttempt().list()` / `client.WebhookEventAttempt().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  WebhookEventAttempt(entopts?: Record<string, any>) {
    const self = this
    return new WebhookEventAttemptEntity(self, entopts)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new ResendSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return ResendSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Resend' }
  }

  toString() {
    return 'Resend ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = ResendSDK


export {
  stdutil,
  config,
  

  BaseFeature,
  ResendEntityBase,

  ResendSDK,
  SDK,
}


export type { DirectResult }
