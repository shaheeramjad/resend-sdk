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
(0, node_test_1.describe)('InboxLabelEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when RESEND_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('RESEND_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.ResendSDK.test();
        const ent = testsdk.InboxLabel();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('validate', async (t) => {
        if (null == __1.config.feature?.validate) {
            t.skip('feature not present in this SDK: validate');
            return;
        }
        const client = __1.ResendSDK.test(undefined, { feature: { validate: { active: true } } });
        await node_assert_1.default.rejects(client.InboxLabel().list({ "inbox_id": 1 }), (err) => 'validate_failed' === err.code);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.RESEND_TEST_LIVE;
        for (const op of ['create', 'list', 'update', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'inbox_label.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "color": { "a": true, "h": "Color", "n": "color", "op": { "list": { "req": true, "type": "`$STRING`" } }, "r": false, "sh": "The color of the label, as a `#RRGGBB` hex color.", "t": "`$STRING`", "key$": "color", "index$": 0 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": true, "sh": "Timestamp when the label was created.", "t": "`$STRING`", "key$": "created_at", "index$": 1 }, "id": { "a": true, "fo": "uuid", "h": "Id", "n": "id", "r": true, "sh": "The ID of the updated label.", "t": "`$STRING`", "key$": "id", "index$": 2 }, "name": { "a": true, "h": "Name", "n": "name", "op": { "update": { "req": false, "type": "`$STRING`" } }, "r": true, "sh": "The name of the label.", "t": "`$STRING`", "key$": "name", "index$": 3 }, "object": { "a": true, "h": "Object", "n": "object", "r": true, "sh": "Type of the response object.", "t": "`$STRING`", "key$": "object", "index$": 4 } }, "id": { "field": "id", "name": "id" }, "name": "inbox_label", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "bf": ["color", "name"], "co": { "id": "POST /inboxes/{inbox_id}/labels", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/inboxes/{inbox_id}/labels", "q": { "exist": ["inbox_id"] }, "r": { "param": { "inbox_id": "inbox_id" } }, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "labels" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /inboxes/{inbox_id}/labels", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/inboxes/{inbox_id}/labels", "q": { "exist": ["inbox_id"] }, "r": { "param": { "inbox_id": "inbox_id" } }, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "labels" }], "t": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }], "key$": "list" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /inboxes/{inbox_id}/labels/{label_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "label_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/inboxes/{inbox_id}/labels/{label_id}", "q": { "exist": ["id", "inbox_id"] }, "r": { "param": { "label_id": "id" } }, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "labels" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "bf": ["color", "name"], "co": { "id": "PATCH /inboxes/{inbox_id}/labels/{label_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "label_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "ex": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "k": "param", "n": "inbox_id", "or": "inbox_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/inboxes/{inbox_id}/labels/{label_id}", "q": { "exist": ["id", "inbox_id"] }, "r": { "param": { "label_id": "id" } }, "rs": { "kind": "json", "media": "application/json" }, "s": [{ "lit": "inboxes" }, { "var": "inbox_id" }, { "lit": "labels" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.inbox"]] }, "key$": "inbox_label", "name__orig": "inbox_label", "Name": "InboxLabel", "name_": "inbox_label", "name-": "inbox-label", "NAME": "INBOX_LABEL", "index$": 22 }, { "active": true, "entity": "inbox_label", "key$": "BasicInboxLabelFlow", "kind": "basic", "name": "BasicInboxLabelFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "inbox_label_ref01" }, "m": { "inbox_id": "inbox01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "inbox_id": "inbox01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "inbox_label_ref01" } }], "index$": 1 }, { "a": true, "d": { "inbox_id": "inbox01" }, "i": { "ref": "inbox_label_ref01", "srcdatavar": "inbox_label_ref01_data", "suffix": "_up0", "textfield": "color" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-inbox_label_ref01" } }], "v": [], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "inbox_label_ref01", "suffix": "_rm0" }, "m": { "id": "inbox_label01", "inbox_id": "inbox01" }, "o": "remove", "s": [], "v": [], "index$": 3 }, { "a": true, "d": {}, "i": { "suffix": "_rt0" }, "m": { "inbox_id": "inbox01" }, "o": "list", "s": [], "v": [{ "apply": "ItemNotExists", "def": { "ref": "inbox_label_ref01" } }], "index$": 4 }] }, 'InboxLabel', { "POST /inboxes/{inbox_id}/labels": { "protocol": "http", "requestBody": { "required": true, "content": { "application/json": { "schema": { "type": "object", "required": ["name"], "properties": { "name": { "type": "string", "minLength": 1, "maxLength": 64, "description": "The name of the label. Max 64 characters. Leading and trailing whitespace is trimmed. Must not match another label on the inbox, compared without case.", "example": "Urgent", "key$": "name" }, "color": { "type": "string", "pattern": "^#[0-9A-Fa-f]{6}$", "description": "The color of the label, as a `#RRGGBB` hex color. It is stored in uppercase. When omitted, a random color from the dashboard palette is used.", "example": "#E93D82", "key$": "color" } }, "x-ref": "#/components/schemas/CreateInboxLabelOptions", "index$": 1 } } } }, "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }] }, "GET /inboxes/{inbox_id}/labels": { "protocol": "http", "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }] }, "DELETE /inboxes/{inbox_id}/labels/{label_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }, { "name": "label_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The Label ID.", "index$": 1 }] }, "PATCH /inboxes/{inbox_id}/labels/{label_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "name": { "type": "string", "minLength": 1, "maxLength": 64, "description": "A new name for the label. Max 64 characters. Leading and trailing whitespace is trimmed. Must not match another label on the inbox, compared without case.", "example": "Billing", "key$": "name" }, "color": { "type": "string", "pattern": "^#[0-9A-Fa-f]{6}$", "description": "A new color for the label, as a `#RRGGBB` hex color. It is stored in uppercase.", "example": "#F76B15", "key$": "color" } }, "x-ref": "#/components/schemas/UpdateInboxLabelOptions", "index$": 1 } } } }, "parameters": [{ "in": "path", "name": "inbox_id", "required": true, "schema": { "type": "string" }, "description": "The Inbox ID or the inbox email address, such as `support@acme.com`. A value that contains `@` must be a valid email address, and any other value must be a UUID, or the request returns a 422. The address match is case-insensitive and never falls back to the inbox `name`. An address that matches no inbox returns a 404.", "example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1", "x-ref": "#/components/parameters/InboxIdOrEmailAddress", "index$": 0 }, { "name": "label_id", "in": "path", "required": true, "schema": { "type": "string", "format": "uuid" }, "description": "The Label ID.", "index$": 1 }] } }, { strict: LIVE_STRICT, t });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const inbox_label_ref01_ent = client.InboxLabel();
        let inbox_label_ref01_data = setup.data.new.inbox_label['inbox_label_ref01'];
        inbox_label_ref01_data['inbox_id'] = setup.idmap['inbox01'];
        inbox_label_ref01_data = (await inbox_label_ref01_ent.create(inbox_label_ref01_data)).data();
        (0, node_assert_1.default)(null != inbox_label_ref01_data.id);
        // LIST
        const inbox_label_ref01_match = {};
        inbox_label_ref01_match['inbox_id'] = setup.idmap['inbox01'];
        const inbox_label_ref01_list = (await inbox_label_ref01_ent.list(inbox_label_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(inbox_label_ref01_list, { id: inbox_label_ref01_data.id })));
        // UPDATE
        const inbox_label_ref01_data_up0 = {};
        inbox_label_ref01_data_up0.id = inbox_label_ref01_data.id;
        inbox_label_ref01_data_up0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01'];
        const inbox_label_ref01_markdef_up0 = { name: 'color', value: 'Mark01-inbox_label_ref01_' + setup.now };
        inbox_label_ref01_data_up0[inbox_label_ref01_markdef_up0.name] = inbox_label_ref01_markdef_up0.value;
        const inbox_label_ref01_resdata_up0 = (await inbox_label_ref01_ent.update(inbox_label_ref01_data_up0)).data();
        (0, node_assert_1.default)(inbox_label_ref01_resdata_up0.id === inbox_label_ref01_data_up0.id);
        (0, node_assert_1.default)(inbox_label_ref01_resdata_up0[inbox_label_ref01_markdef_up0.name] === inbox_label_ref01_markdef_up0.value);
        // REMOVE
        const inbox_label_ref01_match_rm0 = { id: inbox_label_ref01_data.id };
        inbox_label_ref01_match_rm0['inbox_id'] = setup.idmap['inbox_id'] ?? setup.idmap['inbox01'];
        await inbox_label_ref01_ent.remove(inbox_label_ref01_match_rm0);
        // LIST
        const inbox_label_ref01_match_rt0 = {};
        inbox_label_ref01_match_rt0['inbox_id'] = setup.idmap['inbox01'];
        const inbox_label_ref01_list_rt0 = (await inbox_label_ref01_ent.list(inbox_label_ref01_match_rt0)).map((e) => e.data());
        (0, node_assert_1.default)(isempty(select(inbox_label_ref01_list_rt0, { id: inbox_label_ref01_data.id })));
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
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/inbox_label/InboxLabelTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.ResendSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['inbox_label01', 'inbox_label02', 'inbox_label03', 'inbox01', 'inbox02', 'inbox03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'RESEND_TEST_INBOX_LABEL_ENTID': idmap,
        'RESEND_TEST_LIVE': 'FALSE',
        'RESEND_TEST_EXPLAIN': 'FALSE',
        'RESEND_APIKEY': '',
    });
    idmap = env['RESEND_TEST_INBOX_LABEL_ENTID'];
    const live = 'TRUE' === env.RESEND_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['RESEND_TEST_INBOX_LABEL_ENTID'];
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
//# sourceMappingURL=InboxLabelEntity.test.js.map