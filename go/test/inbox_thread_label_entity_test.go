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
const inbox_thread_labelEntityLiveStrict = true


func TestInboxThreadLabelEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.InboxThreadLabel(nil)
		if ent == nil {
			t.Fatal("expected non-nil InboxThreadLabelEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.InboxThreadLabel(nil).Create(map[string]any{"inbox_id": 1, "thread_id": "x", "labels": "x"}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := inbox_thread_labelBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "inbox_thread_label." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		if setup.live {
			for _, _liveKey := range []string{"inbox01", "thread01"} {
				if setup.syntheticOnly || setup.idmap[_liveKey] == nil {
					liveMiss(t, inbox_thread_labelEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_INBOX_THREAD_LABEL_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client

		// CREATE
		inboxThreadLabelRef01Ent := client.InboxThreadLabel(nil)
		inboxThreadLabelRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "inbox_thread_label"}), "inbox_thread_label_ref01"))
		inboxThreadLabelRef01Data["inbox_id"] = setup.idmap["inbox01"]
		inboxThreadLabelRef01Data["thread_id"] = setup.idmap["thread01"]

		inboxThreadLabelRef01DataResult, err := inboxThreadLabelRef01Ent.Create(inboxThreadLabelRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		inboxThreadLabelRef01Data = core.ToMapAny(entityData(inboxThreadLabelRef01DataResult))
		if inboxThreadLabelRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if inboxThreadLabelRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		inboxThreadLabelRef01DataUp0Up := map[string]any{
			"id": inboxThreadLabelRef01Data["id"],
			"inbox_id": setup.idmap["inbox_id"],
		}

		inboxThreadLabelRef01MarkdefUp0Name := "object"
		inboxThreadLabelRef01MarkdefUp0Value := fmt.Sprintf("Mark01-inbox_thread_label_ref01_%d", setup.now)
		inboxThreadLabelRef01DataUp0Up[inboxThreadLabelRef01MarkdefUp0Name] = inboxThreadLabelRef01MarkdefUp0Value

		inboxThreadLabelRef01ResdataUp0Result, err := inboxThreadLabelRef01Ent.Update(inboxThreadLabelRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		inboxThreadLabelRef01ResdataUp0 := core.ToMapAny(entityData(inboxThreadLabelRef01ResdataUp0Result))
		if inboxThreadLabelRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if inboxThreadLabelRef01ResdataUp0["id"] != inboxThreadLabelRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if inboxThreadLabelRef01ResdataUp0[inboxThreadLabelRef01MarkdefUp0Name] != inboxThreadLabelRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", inboxThreadLabelRef01MarkdefUp0Name, inboxThreadLabelRef01ResdataUp0[inboxThreadLabelRef01MarkdefUp0Name])
		}

		// REMOVE
		inboxThreadLabelRef01MatchRm0 := map[string]any{
			"id": inboxThreadLabelRef01Data["id"],
		}
		if v := setup.idmap["inbox_id"]; v != nil {
			inboxThreadLabelRef01MatchRm0["inbox_id"] = v
		} else {
			inboxThreadLabelRef01MatchRm0["inbox_id"] = setup.idmap["inbox01"]
		}
		if v := setup.idmap["thread_id"]; v != nil {
			inboxThreadLabelRef01MatchRm0["thread_id"] = v
		} else {
			inboxThreadLabelRef01MatchRm0["thread_id"] = setup.idmap["thread01"]
		}
		_, err = inboxThreadLabelRef01Ent.Remove(inboxThreadLabelRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

	})
}

func inbox_thread_labelBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "inbox_thread_label", "InboxThreadLabelTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read inbox_thread_label test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse inbox_thread_label test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"inbox_thread_label01", "inbox_thread_label02", "inbox_thread_label03", "inbox01", "inbox02", "inbox03", "thread01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_INBOX_THREAD_LABEL_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_INBOX_THREAD_LABEL_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_INBOX_THREAD_LABEL_ENTID"])
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
