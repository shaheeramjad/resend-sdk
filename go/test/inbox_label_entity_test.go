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
const inbox_labelEntityLiveStrict = true


func TestInboxLabelEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.InboxLabel(nil)
		if ent == nil {
			t.Fatal("expected non-nil InboxLabelEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.InboxLabel(nil).List(map[string]any{"inbox_id": 1}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := inbox_labelBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "inbox_label." + _op, _mode); _shouldSkip {
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
					liveMiss(t, inbox_labelEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_INBOX_LABEL_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client

		// CREATE
		inboxLabelRef01Ent := client.InboxLabel(nil)
		inboxLabelRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "inbox_label"}), "inbox_label_ref01"))
		inboxLabelRef01Data["inbox_id"] = setup.idmap["inbox01"]

		inboxLabelRef01DataResult, err := inboxLabelRef01Ent.Create(inboxLabelRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		inboxLabelRef01Data = core.ToMapAny(entityData(inboxLabelRef01DataResult))
		if inboxLabelRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if inboxLabelRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		inboxLabelRef01Match := map[string]any{
			"inbox_id": setup.idmap["inbox01"],
		}

		inboxLabelRef01ListResult, err := inboxLabelRef01Ent.List(inboxLabelRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		inboxLabelRef01List, inboxLabelRef01ListOk := inboxLabelRef01ListResult.([]any)
		if !inboxLabelRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", inboxLabelRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(inboxLabelRef01List), map[string]any{"id": inboxLabelRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		inboxLabelRef01DataUp0Up := map[string]any{
			"id": inboxLabelRef01Data["id"],
			"inbox_id": setup.idmap["inbox_id"],
		}

		inboxLabelRef01MarkdefUp0Name := "color"
		inboxLabelRef01MarkdefUp0Value := fmt.Sprintf("Mark01-inbox_label_ref01_%d", setup.now)
		inboxLabelRef01DataUp0Up[inboxLabelRef01MarkdefUp0Name] = inboxLabelRef01MarkdefUp0Value

		inboxLabelRef01ResdataUp0Result, err := inboxLabelRef01Ent.Update(inboxLabelRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		inboxLabelRef01ResdataUp0 := core.ToMapAny(entityData(inboxLabelRef01ResdataUp0Result))
		if inboxLabelRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if inboxLabelRef01ResdataUp0["id"] != inboxLabelRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if inboxLabelRef01ResdataUp0[inboxLabelRef01MarkdefUp0Name] != inboxLabelRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", inboxLabelRef01MarkdefUp0Name, inboxLabelRef01ResdataUp0[inboxLabelRef01MarkdefUp0Name])
		}

		// REMOVE
		inboxLabelRef01MatchRm0 := map[string]any{
			"id": inboxLabelRef01Data["id"],
		}
		if v := setup.idmap["inbox_id"]; v != nil {
			inboxLabelRef01MatchRm0["inbox_id"] = v
		} else {
			inboxLabelRef01MatchRm0["inbox_id"] = setup.idmap["inbox01"]
		}
		_, err = inboxLabelRef01Ent.Remove(inboxLabelRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		inboxLabelRef01MatchRt0 := map[string]any{
			"inbox_id": setup.idmap["inbox01"],
		}

		inboxLabelRef01ListRt0Result, err := inboxLabelRef01Ent.List(inboxLabelRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		inboxLabelRef01ListRt0, inboxLabelRef01ListRt0Ok := inboxLabelRef01ListRt0Result.([]any)
		if !inboxLabelRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", inboxLabelRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(inboxLabelRef01ListRt0), map[string]any{"id": inboxLabelRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func inbox_labelBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "inbox_label", "InboxLabelTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read inbox_label test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse inbox_label test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"inbox_label01", "inbox_label02", "inbox_label03", "inbox01", "inbox02", "inbox03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_INBOX_LABEL_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_INBOX_LABEL_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_INBOX_LABEL_ENTID"])
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
