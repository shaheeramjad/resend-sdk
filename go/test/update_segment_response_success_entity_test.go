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

	sdk "github.com/voxgig-sdk/resend-sdk/go"
	"github.com/voxgig-sdk/resend-sdk/go/core"

	vs "github.com/voxgig-sdk/resend-sdk/go/utility/struct"
)

// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const update_segment_response_successEntityLiveStrict = true


func TestUpdateSegmentResponseSuccessEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.UpdateSegmentResponseSuccess(nil)
		if ent == nil {
			t.Fatal("expected non-nil UpdateSegmentResponseSuccessEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.UpdateSegmentResponseSuccess(nil).Update(map[string]any{"id": 1}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := update_segment_response_successBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "update_segment_response_success." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		if setup.live {
			liveMiss(t, update_segment_response_successEntityLiveStrict, "Live entity test blocked: %s", "the flow updates a update_segment_response_success record it did not create")
		}
		client := setup.client
		_ = client

		// Bootstrap entity data from existing test data (no create step in flow).
		updateSegmentResponseSuccessRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath(setup.data, "existing.update_segment_response_success")))
		var updateSegmentResponseSuccessRef01Data map[string]any
		if len(updateSegmentResponseSuccessRef01DataRaw) > 0 {
			updateSegmentResponseSuccessRef01Data = core.ToMapAny(updateSegmentResponseSuccessRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = updateSegmentResponseSuccessRef01Data

		// UPDATE
		updateSegmentResponseSuccessRef01Ent := client.UpdateSegmentResponseSuccess(nil)
		updateSegmentResponseSuccessRef01DataUp0Up := map[string]any{
			"id": updateSegmentResponseSuccessRef01Data["id"],
		}

		updateSegmentResponseSuccessRef01MarkdefUp0Name := "name"
		updateSegmentResponseSuccessRef01MarkdefUp0Value := fmt.Sprintf("Mark01-update_segment_response_success_ref01_%d", setup.now)
		updateSegmentResponseSuccessRef01DataUp0Up[updateSegmentResponseSuccessRef01MarkdefUp0Name] = updateSegmentResponseSuccessRef01MarkdefUp0Value

		updateSegmentResponseSuccessRef01ResdataUp0Result, err := updateSegmentResponseSuccessRef01Ent.Update(updateSegmentResponseSuccessRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		updateSegmentResponseSuccessRef01ResdataUp0 := core.ToMapAny(entityData(updateSegmentResponseSuccessRef01ResdataUp0Result))
		if updateSegmentResponseSuccessRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if updateSegmentResponseSuccessRef01ResdataUp0["id"] != updateSegmentResponseSuccessRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if updateSegmentResponseSuccessRef01ResdataUp0[updateSegmentResponseSuccessRef01MarkdefUp0Name] != updateSegmentResponseSuccessRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", updateSegmentResponseSuccessRef01MarkdefUp0Name, updateSegmentResponseSuccessRef01ResdataUp0[updateSegmentResponseSuccessRef01MarkdefUp0Name])
		}

	})
}

func update_segment_response_successBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "update_segment_response_success", "UpdateSegmentResponseSuccessTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read update_segment_response_success test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse update_segment_response_success test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"update_segment_response_success01", "update_segment_response_success02", "update_segment_response_success03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_UPDATE_SEGMENT_RESPONSE_SUCCESS_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_UPDATE_SEGMENT_RESPONSE_SUCCESS_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_UPDATE_SEGMENT_RESPONSE_SUCCESS_ENTID"])
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
