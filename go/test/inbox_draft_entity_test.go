package sdktest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/shaheeramjad/resend-sdk/go"
	"github.com/shaheeramjad/resend-sdk/go/core"

	vs "github.com/shaheeramjad/resend-sdk/go/utility/struct"
)

// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const inbox_draftEntityLiveStrict = true


func TestInboxDraftEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.InboxDraft(nil)
		if ent == nil {
			t.Fatal("expected non-nil InboxDraftEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.InboxDraft(nil).List(map[string]any{"inbox_id": 1}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := inbox_draftBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "inbox_draft." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		if setup.live {
			for _, _liveKey := range []string{"inbox01"} {
				if setup.syntheticOnly || setup.idmap[_liveKey] == nil {
					liveMiss(t, inbox_draftEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_INBOX_DRAFT_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client

		// CREATE
		inboxDraftRef01Ent := client.InboxDraft(nil)
		inboxDraftRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "inbox_draft"}), "inbox_draft_ref01"))
		inboxDraftRef01Data["inbox_id"] = setup.idmap["inbox01"]

		inboxDraftRef01DataResult, err := inboxDraftRef01Ent.Create(inboxDraftRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		inboxDraftRef01Data = core.ToMapAny(entityData(inboxDraftRef01DataResult))
		if inboxDraftRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if inboxDraftRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		inboxDraftRef01Match := map[string]any{
			"inbox_id": setup.idmap["inbox01"],
		}

		inboxDraftRef01ListResult, err := inboxDraftRef01Ent.List(inboxDraftRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		inboxDraftRef01List, inboxDraftRef01ListOk := inboxDraftRef01ListResult.([]any)
		if !inboxDraftRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", inboxDraftRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(inboxDraftRef01List), map[string]any{"id": inboxDraftRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		inboxDraftRef01DataUp0Up := map[string]any{
			"id": inboxDraftRef01Data["id"],
			"inbox_id": setup.idmap["inbox_id"],
		}

		inboxDraftRef01MarkdefUp0Name := "created_at"
		inboxDraftRef01MarkdefUp0Value := fmt.Sprintf("Mark01-inbox_draft_ref01_%d", setup.now)
		inboxDraftRef01DataUp0Up[inboxDraftRef01MarkdefUp0Name] = inboxDraftRef01MarkdefUp0Value

		inboxDraftRef01ResdataUp0Result, err := inboxDraftRef01Ent.Update(inboxDraftRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		inboxDraftRef01ResdataUp0 := core.ToMapAny(entityData(inboxDraftRef01ResdataUp0Result))
		if inboxDraftRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if inboxDraftRef01ResdataUp0["id"] != inboxDraftRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if inboxDraftRef01ResdataUp0[inboxDraftRef01MarkdefUp0Name] != inboxDraftRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", inboxDraftRef01MarkdefUp0Name, inboxDraftRef01ResdataUp0[inboxDraftRef01MarkdefUp0Name])
		}

		// LOAD
		inboxDraftRef01MatchDt0 := map[string]any{
			"id": inboxDraftRef01Data["id"],
		}
		if v := setup.idmap["inbox_id"]; v != nil {
			inboxDraftRef01MatchDt0["inbox_id"] = v
		} else {
			inboxDraftRef01MatchDt0["inbox_id"] = setup.idmap["inbox01"]
		}
		inboxDraftRef01DataDt0Loaded, err := inboxDraftRef01Ent.Load(inboxDraftRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		inboxDraftRef01DataDt0LoadResult := core.ToMapAny(entityData(inboxDraftRef01DataDt0Loaded))
		if inboxDraftRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if inboxDraftRef01DataDt0LoadResult["id"] != inboxDraftRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		inboxDraftRef01MatchRm0 := map[string]any{
			"id": inboxDraftRef01Data["id"],
		}
		if v := setup.idmap["inbox_id"]; v != nil {
			inboxDraftRef01MatchRm0["inbox_id"] = v
		} else {
			inboxDraftRef01MatchRm0["inbox_id"] = setup.idmap["inbox01"]
		}
		_, err = inboxDraftRef01Ent.Remove(inboxDraftRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		inboxDraftRef01MatchRt0 := map[string]any{
			"inbox_id": setup.idmap["inbox01"],
		}

		inboxDraftRef01ListRt0Result, err := inboxDraftRef01Ent.List(inboxDraftRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		inboxDraftRef01ListRt0, inboxDraftRef01ListRt0Ok := inboxDraftRef01ListRt0Result.([]any)
		if !inboxDraftRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", inboxDraftRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(inboxDraftRef01ListRt0), map[string]any{"id": inboxDraftRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func inbox_draftBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "inbox_draft", "InboxDraftTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read inbox_draft test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse inbox_draft test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"inbox_draft01", "inbox_draft02", "inbox_draft03", "inbox01", "inbox02", "inbox03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_INBOX_DRAFT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_INBOX_DRAFT_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_INBOX_DRAFT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add inbox_id alias for update test.
	if idmapResolved["inbox_id"] == nil {
		idmapResolved["inbox_id"] = idmapResolved["inbox01"]
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
