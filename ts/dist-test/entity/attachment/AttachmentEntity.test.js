"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('AttachmentEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('RESEND_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ResendSDK.test();
        const ent = testsdk.Attachment();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('validate', async (t) => {
        if (null == __1.config.feature?.validate) {
            t.skip('feature not present in this SDK: validate');
            return;
        }
        const client = __1.ResendSDK.test(undefined, { feature: { validate: { active: true } } });
        await node_assert_1.default.rejects(client.Attachment().list({ "email_id": 1 }), (err) => 'validate_failed' === err.code);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.RESEND_TEST_LIVE;
        for (const op of ['list', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'attachment.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "content_disposition": { "a": true, "h": "Content Disposition", "n": "content_disposition", "r": false, "sh": "How the attachment should be displayed.", "t": ["`$ONE`", ["`$STRING`", "`$NULL`"]], "key$": "content_disposition", "index$": 0 }, "content_id": { "a": true, "h": "Content Id", "n": "content_id", "r": false, "sh": "The content ID for inline attachments.", "t": "`$STRING`", "key$": "content_id", "index$": 1 }, "content_type": { "a": true, "h": "Content Type", "n": "content_type", "r": false, "sh": "The MIME type of the attachment.", "t": "`$STRING`", "key$": "content_type", "index$": 2 }, "download_url": { "a": true, "h": "Download Url", "n": "download_url", "r": false, "sh": "Signed URL to download the attachment content.", "t": "`$STRING`", "key$": "download_url", "index$": 3 }, "expires_at": { "a": true, "fo": "date-time", "h": "Expires At", "n": "expires_at", "r": false, "sh": "Timestamp when the download URL expires.", "t": "`$STRING`", "key$": "expires_at", "index$": 4 }, "filename": { "a": true, "h": "Filename", "n": "filename", "r": false, "sh": "The filename of the attachment.", "t": ["`$ONE`", ["`$STRING`", "`$NULL`"]], "key$": "filename", "index$": 5 }, "id": { "a": true, "fo": "uuid", "h": "Id", "n": "id", "r": false, "sh": "The ID of the attachment.", "t": "`$STRING`", "key$": "id", "index$": 6 }, "object": { "a": true, "h": "Object", "n": "object", "r": false, "sh": "The type of object.", "t": "`$STRING`", "key$": "object", "index$": 7 }, "size": { "a": true, "h": "Size", "n": "size", "r": false, "sh": "Size of the attachment in bytes.", "t": "`$INTEGER`", "key$": "size", "index$": 8 } }, "id": { "field": "id", "name": "id" }, "name": "attachment", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /emails/{email_id}/attachments", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "email_id", "or": "email_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "k": "query", "n": "after", "or": "after", "r": false, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "query", "n": "before", "or": "before", "r": false, "t": "`$STRING`", "index$": 1 }, { "a": true, "k": "query", "n": "limit", "or": "limit", "r": false, "t": "`$INTEGER`", "index$": 2 }] }, "k": "http", "m": "GET", "o": "/emails/{email_id}/attachments", "q": { "exist": ["email_id"] }, "r": {}, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "emails" }, { "var": "email_id" }, { "lit": "attachments" }], "t": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /emails/{email_id}/attachments/{attachment_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "email_id", "or": "email_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "id", "or": "attachment_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/emails/{email_id}/attachments/{attachment_id}", "q": { "exist": ["email_id", "id"] }, "r": { "param": { "attachment_id": "id" } }, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "emails" }, { "var": "email_id" }, { "lit": "attachments" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [["$.main.kit.entity.email"]] }, "key$": "attachment", "name__orig": "attachment", "Name": "Attachment", "name_": "attachment", "name-": "attachment", "NAME": "ATTACHMENT", "index$": 1 }, { "active": true, "entity": "attachment", "key$": "BasicAttachmentFlow", "kind": "basic", "name": "BasicAttachmentFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": { "email_id": "email01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "attachment_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "attachment_ref01", "srcdatavar": "attachment_ref01_data", "suffix": "_dt0" }, "m": { "email_id": "email01", "id": "attachment01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-attachment_ref01" } }], "index$": 1 }] }, 'Attachment', { "GET /emails/{email_id}/attachments": { "protocol": "http", "parameters": [{ "name": "email_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The ID of the email.", "index$": 0 }, { "name": "limit", "in": "query", "required": false, "schema": { "type": "integer" }, "description": "Maximum number of attachments to return.", "index$": 1 }, { "name": "after", "in": "query", "required": false, "schema": { "type": "string", "format": "uuid" }, "description": "Pagination cursor to fetch results after this attachment ID. Cannot be used with 'before'.", "index$": 2 }, { "name": "before", "in": "query", "required": false, "schema": { "type": "string", "format": "uuid" }, "description": "Pagination cursor to fetch results before this attachment ID. Cannot be used with 'after'.", "index$": 3 }] }, "GET /emails/{email_id}/attachments/{attachment_id}": { "protocol": "http", "parameters": [{ "name": "email_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The ID of the email.", "index$": 0 }, { "name": "attachment_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The ID of the attachment.", "index$": 1 }] } }, { strict: LIVE_STRICT, t });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let attachment_ref01_data = Object.values(setup.data.existing.attachment)[0];
        // LIST
        const attachment_ref01_ent = client.Attachment();
        const attachment_ref01_match = {};
        attachment_ref01_match['email_id'] = setup.idmap['email01'];
        const attachment_ref01_list = (await attachment_ref01_ent.list(attachment_ref01_match)).map((e) => e.data());
        // LOAD
        const attachment_ref01_match_dt0 = {};
        attachment_ref01_match_dt0.id = attachment_ref01_data.id;
        attachment_ref01_match_dt0['email_id'] = setup.idmap['email_id'] ?? setup.idmap['email01'];
        const attachment_ref01_data_dt0 = (await attachment_ref01_ent.load(attachment_ref01_match_dt0)).data();
        (0, node_assert_1.default)(attachment_ref01_data_dt0.id === attachment_ref01_data.id);
    });
});
// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true;
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/attachment/AttachmentTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ResendSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['attachment01', 'attachment02', 'attachment03', 'email01', 'email02', 'email03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'RESEND_TEST_ATTACHMENT_ENTID': idmap,
        'RESEND_TEST_LIVE': 'FALSE',
        'RESEND_TEST_EXPLAIN': 'FALSE',
        'RESEND_APIKEY': '',
    });
    idmap = env['RESEND_TEST_ATTACHMENT_ENTID'];
    const live = 'TRUE' === env.RESEND_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['RESEND_TEST_ATTACHMENT_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.ResendSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
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
        ]));
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
    };
    return setup;
}
//# sourceMappingURL=AttachmentEntity.test.js.map