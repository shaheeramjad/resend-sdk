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
const inbox_threadEntityLiveStrict = true


func TestInboxThreadEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.InboxThread(nil)
		if ent == nil {
			t.Fatal("expected non-nil InboxThreadEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.InboxThread(nil).List(map[string]any{"inbox_id": 1}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := inbox_threadBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "inbox_thread." + _op, _mode); _shouldSkip {
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
					liveMiss(t, inbox_threadEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_INBOX_THREAD_ENTID", _liveKey)
				}
			}
		}
		if setup.live {
			liveMiss(t, inbox_threadEntityLiveStrict, "Live entity test blocked: %s", "the flow updates a inbox_thread record it did not create")
		}
		client := setup.client
		_ = client

		// Bootstrap entity data from existing test data (no create step in flow).
		inboxThreadRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath(setup.data, "existing.inbox_thread")))
		var inboxThreadRef01Data map[string]any
		if len(inboxThreadRef01DataRaw) > 0 {
			inboxThreadRef01Data = core.ToMapAny(inboxThreadRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = inboxThreadRef01Data

		// LIST
		inboxThreadRef01Ent := client.InboxThread(nil)
		inboxThreadRef01Match := map[string]any{
			"inbox_id": setup.idmap["inbox01"],
		}

		inboxThreadRef01ListResult, err := inboxThreadRef01Ent.List(inboxThreadRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, inboxThreadRef01ListOk := inboxThreadRef01ListResult.([]any)
		if !inboxThreadRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", inboxThreadRef01ListResult)
		}

		// UPDATE
		inboxThreadRef01DataUp0Up := map[string]any{
			"id": inboxThreadRef01Data["id"],
			"inbox_id": setup.idmap["inbox_id"],
		}

		inboxThreadRef01MarkdefUp0Name := "folder"
		inboxThreadRef01MarkdefUp0Value := fmt.Sprintf("Mark01-inbox_thread_ref01_%d", setup.now)
		inboxThreadRef01DataUp0Up[inboxThreadRef01MarkdefUp0Name] = inboxThreadRef01MarkdefUp0Value

		inboxThreadRef01ResdataUp0Result, err := inboxThreadRef01Ent.Update(inboxThreadRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		inboxThreadRef01ResdataUp0 := core.ToMapAny(entityData(inboxThreadRef01ResdataUp0Result))
		if inboxThreadRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if inboxThreadRef01ResdataUp0["id"] != inboxThreadRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if inboxThreadRef01ResdataUp0[inboxThreadRef01MarkdefUp0Name] != inboxThreadRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", inboxThreadRef01MarkdefUp0Name, inboxThreadRef01ResdataUp0[inboxThreadRef01MarkdefUp0Name])
		}

		// LOAD
		inboxThreadRef01MatchDt0 := map[string]any{
			"id": inboxThreadRef01Data["id"],
		}
		if v := setup.idmap["inbox_id"]; v != nil {
			inboxThreadRef01MatchDt0["inbox_id"] = v
		} else {
			inboxThreadRef01MatchDt0["inbox_id"] = setup.idmap["inbox01"]
		}
		inboxThreadRef01DataDt0Loaded, err := inboxThreadRef01Ent.Load(inboxThreadRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		inboxThreadRef01DataDt0LoadResult := core.ToMapAny(entityData(inboxThreadRef01DataDt0Loaded))
		if inboxThreadRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if inboxThreadRef01DataDt0LoadResult["id"] != inboxThreadRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func inbox_threadBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "inbox_thread", "InboxThreadTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read inbox_thread test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse inbox_thread test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"inbox_thread01", "inbox_thread02", "inbox_thread03", "inbox01", "inbox02", "inbox03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_INBOX_THREAD_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_INBOX_THREAD_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_INBOX_THREAD_ENTID"])
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
