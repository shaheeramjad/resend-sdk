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
(0, node_test_1.describe)('InboxThreadLabelEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('RESEND_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ResendSDK.test();
        const ent = testsdk.InboxThreadLabel();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('validate', async (t) => {
        if (null == __1.config.feature?.validate) {
            t.skip('feature not present in this SDK: validate');
            return;
        }
        const client = __1.ResendSDK.test(undefined, { feature: { validate: { active: true } } });
        await node_assert_1.default.rejects(client.InboxThreadLabel().create({ "inbox_id": 1, "thread_id": "x", "labels": "x" }), (err) => 'validate_failed' === err.code);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.RESEND_TEST_LIVE;
        for (const op of ['create', 'update', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'inbox_thread_label.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "data": { "a": true, "h": "Data", "n": "data", "r": true, "sh": "Every label on the thread after the change.", "t": "`$ARRAY`", "key$": "data", "index$": 0 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "t": "`$STRING`", "key$": "id", "index$": 1 }, "labels": { "a": true, "h": "Labels", "n": "labels", "r": true, "sh": "The IDs of the labels to apply to the thread.", "t": "`$ARRAY`", "key$": "labels", "index$": 2 }, "object": { "a": true, "h": "Object", "n": "object", "r": true, "sh": "Type of the response object.", "t": "`$STRING`", "key$": "object", "index$": 3 } }, "id": { "field": "id", "name": "id" }, "name": "inbox_thread_label", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "bf": ["labels"], "co": { "id": "POST /inboxes/{inbox_id}/threads/{thread_id}/labels", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "thread_id", "or": "thread_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/inboxes/{inbox_id}/threads/{thread_id}/labels", "q": { "exist": ["inbox_id", "thread_id"] }, "r": {}, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "threads" }, { "var": "thread_id" }, { "lit": "labels" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /inboxes/{inbox_id}/threads/{thread_id}/labels/{label_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "label_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 1 }, { "a": true, "k": "param", "n": "thread_id", "or": "thread_id", "r": true, "t": "`$STRING`", "index$": 2 }] }, "k": "http", "m": "DELETE", "o": "/inboxes/{inbox_id}/threads/{thread_id}/labels/{label_id}", "q": { "exist": ["id", "inbox_id", "thread_id"] }, "r": { "param": { "label_id": "id" } }, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "threads" }, { "var": "thread_id" }, { "lit": "labels" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "bf": ["labels"], "co": { "id": "PUT /inboxes/{inbox_id}/threads/{thread_id}/labels", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "thread_id", "or": "thread_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PUT", "o": "/inboxes/{inbox_id}/threads/{thread_id}/labels", "q": { "exist": ["inbox_id", "thread_id"] }, "r": {}, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "threads" }, { "var": "thread_id" }, { "lit": "labels" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.inbox"]] }, "key$": "inbox_thread_label", "name__orig": "inbox_thread_label", "Name": "InboxThreadLabel", "name_": "inbox_thread_label", "name-": "inbox-thread-label", "NAME": "INBOX_THREAD_LABEL", "index$": 25 }, { "active": true, "entity": "inbox_thread_label", "key$": "BasicInboxThreadLabelFlow", "kind": "basic", "name": "BasicInboxThreadLabelFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "inbox_thread_label_ref01" }, "m": { "inbox_id": "inbox01", "thread_id": "thread01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": { "inbox_id": "inbox01" }, "i": { "ref": "inbox_thread_label_ref01", "srcdatavar": "inbox_thread_label_ref01_data", "suffix": "_up0", "textfield": "object" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-inbox_thread_label_ref01" } }], "v": [], "index$": 1 }, { "a": true, "d": {}, "i": { "ref": "inbox_thread_label_ref01", "suffix": "_rm0" }, "m": { "id": "inbox_thread_label01", "inbox_id": "inbox01", "thread_id": "thread01" }, "o": "remove", "s": [], "v": [], "index$": 2 }] }, 'InboxThreadLabel', { "POST /inboxes/{inbox_id}/threads/{thread_id}/labels": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "required": ["labels"], "properties": { "labels": { "type": "array", "minItems": 1, "maxItems": 10, "items": { "type": "string", "format": "uuid" }, "description": "The IDs of the labels to apply to the thread. Between 1 and 10 IDs, each a label of the inbox.", "example": ["7f9c1d2e-4a6b-4c3d-8e15-9b0a7c6d5e34"], "key$": "labels" } }, "x-ref": "#/components/schemas/AddInboxThreadLabelsOptions", "index$": 1 } } } }, "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }, { "name": "thread_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The Thread ID.", "index$": 1 }] }, "DELETE /inboxes/{inbox_id}/threads/{thread_id}/labels/{label_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }, { "name": "thread_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The Thread ID.", "index$": 1 }, { "name": "label_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The Label ID.", "index$": 2 }] }, "PUT /inboxes/{inbox_id}/threads/{thread_id}/labels": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "required": ["labels"], "properties": { "labels": { "type": "array", "maxItems": 100, "items": { "type": "string", "format": "uuid" }, "description": "The IDs of the labels the thread carries after the change. Up to 100 IDs, each a label of the inbox. Pass an empty array to take every label off the thread.", "example": ["7f9c1d2e-4a6b-4c3d-8e15-9b0a7c6d5e34"], "key$": "labels" } }, "x-ref": "#/components/schemas/SetInboxThreadLabelsOptions", "index$": 1 } } } }, "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }, { "name": "thread_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The Thread ID.", "index$": 1 }] } }, { strict: LIVE_STRICT, t });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const inbox_thread_label_ref01_ent = client.InboxThreadLabel();
        let inbox_thread_label_ref01_data = setup.data.new.inbox_thread_label['inbox_thread_label_ref01'];
        inbox_thread_label_ref01_data['inbox_id'] = setup.idmap['inbox01'];
        inbox_thread_label_ref01_data['thread_id'] = setup.idmap['thread01'];
        inbox_thread_label_ref01_data = (await inbox_thread_label_ref01_ent.create(inbox_thread_label_ref01_data)).data();
        (0, node_assert_1.default)(null != inbox_thread_label_ref01_data.id);
        // UPDATE
        const inbox_thread_label_ref01_data_up0 = {};
        inbox_thread_label_ref01_data_up0.id = inbox_thread_label_ref01_data.id;
        inbox_thread_label_ref01_data_up0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01'];
        const inbox_thread_label_ref01_markdef_up0 = { name: 'object', value: 'Mark01-inbox_thread_label_ref01_' + setup.now };
        inbox_thread_label_ref01_data_up0[inbox_thread_label_ref01_markdef_up0.name] = inbox_thread_label_ref01_markdef_up0.value;
        const inbox_thread_label_ref01_resdata_up0 = (await inbox_thread_label_ref01_ent.update(inbox_thread_label_ref01_data_up0)).data();
        (0, node_assert_1.default)(inbox_thread_label_ref01_resdata_up0.id === inbox_thread_label_ref01_data_up0.id);
        (0, node_assert_1.default)(inbox_thread_label_ref01_resdata_up0[inbox_thread_label_ref01_markdef_up0.name] === inbox_thread_label_ref01_markdef_up0.value);
        // REMOVE
        const inbox_thread_label_ref01_match_rm0 = { id: inbox_thread_label_ref01_data.id };
        inbox_thread_label_ref01_match_rm0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01'];
        inbox_thread_label_ref01_match_rm0['thread_id'] = setup.idmap['thread_id'] ?? setup.idmap['thread01'];
        await inbox_thread_label_ref01_ent.remove(inbox_thread_label_ref01_match_rm0);
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
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/inbox_thread_label/InboxThreadLabelTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ResendSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['inbox_thread_label01', 'inbox_thread_label02', 'inbox_thread_label03', 'inbox01', 'inbox02', 'inbox03', 'thread01'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'RESEND_TEST_INBOX_THREAD_LABEL_ENTID': idmap,
        'RESEND_TEST_LIVE': 'FALSE',
        'RESEND_TEST_EXPLAIN': 'FALSE',
        'RESEND_APIKEY': '',
    });
    idmap = env['RESEND_TEST_INBOX_THREAD_LABEL_ENTID'];
    const live = 'TRUE' === env.RESEND_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['RESEND_TEST_INBOX_THREAD_LABEL_ENTID'];
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
//# sourceMappingURL=InboxThreadLabelEntity.test.js.map