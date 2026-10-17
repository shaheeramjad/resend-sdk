# InboxDraft entity test

import json
import os
import time

import pytest

from resend_sdk.utility.voxgig_struct import voxgig_struct as vs
from resend_sdk import ResendSDK
from resend_sdk.core import helpers
from resend_sdk.config import shared_config
from resend_sdk.feature.base_feature import ResendBaseFeature

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner



# main.kit.test.live.strict is true (the default is true): a live
# request that fails, or a live test missing an input it needs,
# fails the test.
# An account with no record for a test to read skips it either way.
LIVE_STRICT = True


class TestInboxDraftEntity:

    def test_should_create_instance(self):
        testsdk = ResendSDK.test(None, None)
        ent = testsdk.InboxDraft(None)
        assert ent is not None

    def test_should_refuse_an_invalid_request(self):
        if "validate" not in (shared_config().get("feature") or {}):
            pytest.skip("feature not present in this SDK: validate")
        client = ResendSDK.test(
            None, {"feature": {"validate": {"active": True}}})
        with pytest.raises(Exception) as err:
            client.InboxDraft(None).list({"inbox_id": 1}, None)
        assert "validate_failed" == getattr(err.value, "code", None)

    def test_should_run_basic_flow(self):
        setup = _inbox_draft_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "inbox_draft." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        if setup["live"]:
            for _live_key in ["inbox01"]:
                if setup.get("synthetic_only") or setup["idmap"].get(_live_key) is None:
                    runner.live_miss(LIVE_STRICT, f"Live entity test blocked: needs {_live_key} via RESEND_TEST_INBOX_DRAFT_ENTID")
        client = setup["client"]

        # CREATE
        inbox_draft_ref01_ent = client.InboxDraft(None)
        inbox_draft_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.inbox_draft"), "inbox_draft_ref01"))
        inbox_draft_ref01_data["inbox_id"] = setup["idmap"]["inbox01"]

        inbox_draft_ref01_data = helpers.to_map(runner.entity_data(inbox_draft_ref01_ent.create(inbox_draft_ref01_data, None)))
        assert inbox_draft_ref01_data is not None
        assert inbox_draft_ref01_data["id"] is not None

        # LIST
        inbox_draft_ref01_match = {
            "inbox_id": setup["idmap"]["inbox01"],
        }

        inbox_draft_ref01_list_result = inbox_draft_ref01_ent.list(inbox_draft_ref01_match, None)
        assert isinstance(inbox_draft_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(inbox_draft_ref01_list_result),
            {"id": inbox_draft_ref01_data["id"]})
        assert not vs.isempty(found_item)

        # UPDATE
        inbox_draft_ref01_data_up0_up = {
            "id": inbox_draft_ref01_data["id"],
            "inbox_id": setup["idmap"]["inbox_id"],
        }

        inbox_draft_ref01_markdef_up0_name = "created_at"
        inbox_draft_ref01_markdef_up0_value = "Mark01-inbox_draft_ref01_" + str(setup["now"])
        inbox_draft_ref01_data_up0_up[inbox_draft_ref01_markdef_up0_name] = inbox_draft_ref01_markdef_up0_value

        inbox_draft_ref01_resdata_up0 = helpers.to_map(runner.entity_data(inbox_draft_ref01_ent.update(inbox_draft_ref01_data_up0_up, None)))
        assert inbox_draft_ref01_resdata_up0 is not None
        assert inbox_draft_ref01_resdata_up0["id"] == inbox_draft_ref01_data_up0_up["id"]
        assert inbox_draft_ref01_resdata_up0[inbox_draft_ref01_markdef_up0_name] == inbox_draft_ref01_markdef_up0_value

        # LOAD
        inbox_draft_ref01_match_dt0 = {
            "id": inbox_draft_ref01_data["id"],
        }
        inbox_draft_ref01_match_dt0["inbox_id"] = setup["idmap"].get("inbox_id") or setup["idmap"].get("inbox01")
        inbox_draft_ref01_data_dt0_loaded = inbox_draft_ref01_ent.load(inbox_draft_ref01_match_dt0, None)
        inbox_draft_ref01_data_dt0_load_result = helpers.to_map(runner.entity_data(inbox_draft_ref01_data_dt0_loaded))
        assert inbox_draft_ref01_data_dt0_load_result is not None
        assert inbox_draft_ref01_data_dt0_load_result["id"] == inbox_draft_ref01_data["id"]

        # REMOVE
        inbox_draft_ref01_match_rm0 = {
            "id": inbox_draft_ref01_data["id"],
        }
        inbox_draft_ref01_match_rm0["inbox_id"] = setup["idmap"].get("inbox_id") or setup["idmap"].get("inbox01")
        inbox_draft_ref01_ent.remove(inbox_draft_ref01_match_rm0, None)

        # LIST
        inbox_draft_ref01_match_rt0 = {
            "inbox_id": setup["idmap"]["inbox01"],
        }

        inbox_draft_ref01_list_rt0_result = inbox_draft_ref01_ent.list(inbox_draft_ref01_match_rt0, None)
        assert isinstance(inbox_draft_ref01_list_rt0_result, list)

        not_found_item = vs.select(
            runner.entity_list_to_data(inbox_draft_ref01_list_rt0_result),
            {"id": inbox_draft_ref01_data["id"]})
        assert vs.isempty(not_found_item)



def _inbox_draft_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/inbox_draft/InboxDraftTestData.json")
    with open(entity_data_file, "r", encoding="utf-8") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = ResendSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["inbox_draft01", "inbox_draft02", "inbox_draft03", "inbox01", "inbox02", "inbox03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Whether *_ENTID supplied the idmap, read before env_override consumes
    # it: without it, the ids a live flow binds are the fixture's synthetic ones.
    _entid_env_raw = os.environ.get(
        "RESEND_TEST_INBOX_DRAFT_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "RESEND_TEST_INBOX_DRAFT_ENTID": idmap,
        "RESEND_TEST_LIVE": "FALSE",
        "RESEND_TEST_EXPLAIN": "FALSE",
        "RESEND_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("RESEND_TEST_INBOX_DRAFT_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)
    if idmap_resolved.get("inbox_id") is None:
        idmap_resolved["inbox_id"] = idmap_resolved.get("inbox01")

    if env.get("RESEND_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
                "apikey": env.get("RESEND_APIKEY"),
            },
            extra or {},
        ])
        client = ResendSDK(helpers.to_map(merged_opts))

    _live = env.get("RESEND_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("RESEND_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
