package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/resend-sdk/go"
	"github.com/voxgig-sdk/resend-sdk/go/core"

	vs "github.com/voxgig-sdk/resend-sdk/go/utility/struct"
)

// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const inbox_draft_response_successEntityLiveStrict = true


func TestInboxDraftResponseSuccessEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.InboxDraftResponseSuccess(nil)
		if ent == nil {
			t.Fatal("expected non-nil InboxDraftResponseSuccessEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.InboxDraftResponseSuccess(nil).Load(map[string]any{"draft_id": 1, "inbox_id": "x"}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := inbox_draft_response_successBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "inbox_draft_response_success." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		if setup.live {
			for _, _liveKey := range []string{"draft01", "inbox01"} {
				if setup.syntheticOnly || setup.idmap[_liveKey] == nil {
					liveMiss(t, inbox_draft_response_successEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_INBOX_DRAFT_RESPONSE_SUCCESS_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client

		// CREATE
		inboxDraftResponseSuccessRef01Ent := client.InboxDraftResponseSuccess(nil)
		inboxDraftResponseSuccessRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "inbox_draft_response_success"}), "inbox_draft_response_success_ref01"))
		inboxDraftResponseSuccessRef01Data["draft_id"] = setup.idmap["draft01"]
		inboxDraftResponseSuccessRef01Data["inbox_id"] = setup.idmap["inbox01"]

		inboxDraftResponseSuccessRef01DataResult, err := inboxDraftResponseSuccessRef01Ent.Create(inboxDraftResponseSuccessRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		inboxDraftResponseSuccessRef01Data = core.ToMapAny(entityData(inboxDraftResponseSuccessRef01DataResult))
		if inboxDraftResponseSuccessRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if inboxDraftResponseSuccessRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LOAD
		inboxDraftResponseSuccessRef01MatchDt0 := map[string]any{
			"id": inboxDraftResponseSuccessRef01Data["id"],
		}
		if v := setup.idmap["inbox_id"]; v != nil {
			inboxDraftResponseSuccessRef01MatchDt0["inbox_id"] = v
		} else {
			inboxDraftResponseSuccessRef01MatchDt0["inbox_id"] = setup.idmap["inbox01"]
		}
		inboxDraftResponseSuccessRef01DataDt0Loaded, err := inboxDraftResponseSuccessRef01Ent.Load(inboxDraftResponseSuccessRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		inboxDraftResponseSuccessRef01DataDt0LoadResult := core.ToMapAny(entityData(inboxDraftResponseSuccessRef01DataDt0Loaded))
		if inboxDraftResponseSuccessRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if inboxDraftResponseSuccessRef01DataDt0LoadResult["id"] != inboxDraftResponseSuccessRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func inbox_draft_response_successBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "inbox_draft_response_success", "InboxDraftResponseSuccessTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read inbox_draft_response_success test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse inbox_draft_response_success test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"inbox_draft_response_success01", "inbox_draft_response_success02", "inbox_draft_response_success03", "inbox01", "inbox02", "inbox03", "draft01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_INBOX_DRAFT_RESPONSE_SUCCESS_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_INBOX_DRAFT_RESPONSE_SUCCESS_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_INBOX_DRAFT_RESPONSE_SUCCESS_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["RESEND_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["RESEND_APIKEY"],
			},
			extraOpts,
		})
		client = sdk.NewResendSDK(core.ToMapAny(mergedOpts))
	}

	live := env["RESEND_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["RESEND_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
