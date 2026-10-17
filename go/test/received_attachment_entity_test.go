package sdktest

import (
	"encoding/json"
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
const received_attachmentEntityLiveStrict = true


func TestReceivedAttachmentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ReceivedAttachment(nil)
		if ent == nil {
			t.Fatal("expected non-nil ReceivedAttachmentEntity")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.ReceivedAttachment(nil).List(map[string]any{"email_id": 1}, nil)
		if sdkerr, ok := err.(*core.ResendError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
		setup := received_attachmentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "received_attachment." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		if setup.live {
			for _, _liveKey := range []string{"email01"} {
				if setup.syntheticOnly || setup.idmap[_liveKey] == nil {
					liveMiss(t, received_attachmentEntityLiveStrict, "Live entity test blocked: needs %s via RESEND_TEST_RECEIVED_ATTACHMENT_ENTID", _liveKey)
				}
			}
		}
		client := setup.client
		_ = client
		if setup.live {
			liveExisting(t, received_attachmentEntityLiveStrict, setup.data, "received_attachment", func() (any, error) {
				return client.ReceivedAttachment(nil).List(map[string]any{"email_id": setup.idmap["email01"]}, nil)
			})
		}

		// Bootstrap entity data from existing test data (no create step in flow).
		receivedAttachmentRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath(setup.data, "existing.received_attachment")))
		var receivedAttachmentRef01Data map[string]any
		if len(receivedAttachmentRef01DataRaw) > 0 {
			receivedAttachmentRef01Data = core.ToMapAny(receivedAttachmentRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = receivedAttachmentRef01Data

		// LIST
		receivedAttachmentRef01Ent := client.ReceivedAttachment(nil)
		receivedAttachmentRef01Match := map[string]any{
			"email_id": setup.idmap["email01"],
		}

		receivedAttachmentRef01ListResult, err := receivedAttachmentRef01Ent.List(receivedAttachmentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, receivedAttachmentRef01ListOk := receivedAttachmentRef01ListResult.([]any)
		if !receivedAttachmentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", receivedAttachmentRef01ListResult)
		}

		// LOAD
		receivedAttachmentRef01MatchDt0 := map[string]any{
			"id": receivedAttachmentRef01Data["id"],
		}
		if v := setup.idmap["email_id"]; v != nil {
			receivedAttachmentRef01MatchDt0["email_id"] = v
		} else {
			receivedAttachmentRef01MatchDt0["email_id"] = setup.idmap["email01"]
		}
		receivedAttachmentRef01DataDt0Loaded, err := receivedAttachmentRef01Ent.Load(receivedAttachmentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		receivedAttachmentRef01DataDt0LoadResult := core.ToMapAny(entityData(receivedAttachmentRef01DataDt0Loaded))
		if receivedAttachmentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if receivedAttachmentRef01DataDt0LoadResult["id"] != receivedAttachmentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func received_attachmentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "received_attachment", "ReceivedAttachmentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read received_attachment test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse received_attachment test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap, _ := vs.Transform(
		[]any{"received_attachment01", "received_attachment02", "received_attachment03", "email01"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("RESEND_TEST_RECEIVED_ATTACHMENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"RESEND_TEST_RECEIVED_ATTACHMENT_ENTID": idmap,
		"RESEND_TEST_LIVE":      "FALSE",
		"RESEND_TEST_EXPLAIN":   "FALSE",
		"RESEND_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["RESEND_TEST_RECEIVED_ATTACHMENT_ENTID"])
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
