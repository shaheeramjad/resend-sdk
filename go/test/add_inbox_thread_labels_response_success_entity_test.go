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
const add_inbox_thread_labels_response_successEntityLiveStrict = true


func TestAddInboxThreadLabelsResponseSuccessEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.AddInboxThreadLabelsResponseSuccess(nil)
		if ent == nil {
			t.Fatal("expected non-nil AddInboxThreadLabelsResponseSuccessEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.AddInboxThreadLabelsResponseSuccess(nil).Create(map[string]any{"inbox_id": 1, "thread_id": "x", "labels": "x"}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := add_inbox_thread_labels_response_successBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "add_inbox_thread_labels_response_success." + _op, _mode); _shouldSkip {
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
					liveMiss(t, add_inbox_thread_labels_response_successEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_ADD_INBOX_THREAD_LABELS_RESPONSE_SUCCESS_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client

		// CREATE
		addInboxThreadLabelsResponseSuccessRef01Ent := client.AddInboxThreadLabelsResponseSuccess(nil)
		addInboxThreadLabelsResponseSuccessRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "add_inbox_thread_labels_response_success"}), "add_inbox_thread_labels_response_success_ref01"))
		addInboxThreadLabelsResponseSuccessRef01Data["inbox_id"] = setup.idmap["inbox01"]
		addInboxThreadLabelsResponseSuccessRef01Data["thread_id"] = setup.idmap["thread01"]

		addInboxThreadLabelsResponseSuccessRef01DataResult, err := addInboxThreadLabelsResponseSuccessRef01Ent.Create(addInboxThreadLabelsResponseSuccessRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		addInboxThreadLabelsResponseSuccessRef01Data = core.ToMapAny(entityData(addInboxThreadLabelsResponseSuccessRef01DataResult))
		if addInboxThreadLabelsResponseSuccessRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func add_inbox_thread_labels_response_successBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "add_inbox_thread_labels_response_success", "AddInboxThreadLabelsResponseSuccessTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read add_inbox_thread_labels_response_success test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse add_inbox_thread_labels_response_success test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"add_inbox_thread_labels_response_success01", "add_inbox_thread_labels_response_success02", "add_inbox_thread_labels_response_success03", "inbox01", "inbox02", "inbox03", "thread01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_ADD_INBOX_THREAD_LABELS_RESPONSE_SUCCESS_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_ADD_INBOX_THREAD_LABELS_RESPONSE_SUCCESS_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_ADD_INBOX_THREAD_LABELS_RESPONSE_SUCCESS_ENTID"])
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
