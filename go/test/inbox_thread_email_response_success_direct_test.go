package sdktest

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	sdk "github.com/voxgig-sdk/resend-sdk/go"
	"github.com/voxgig-sdk/resend-sdk/go/core"
)

// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const inbox_thread_email_response_successDirectLiveStrict = true

func TestInboxThreadEmailResponseSuccessDirect(t *testing.T) {
	t.Run("direct-load-inbox_thread_email_response_success", func(t *testing.T) {
		setup := inbox_thread_email_response_successDirectSetup(map[string]any{"id": "direct01"})
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		if _shouldSkip, _reason := isControlSkipped("direct", "direct-load-inbox_thread_email_response_success", _mode); _shouldSkip {
			if _reason == "" {
				_reason = "skipped via sdk-test-control.json"
			}
			t.Skip(_reason)
			return
		}
		if setup.live {
			for _, _liveKey := range []string{"email01", "inbox01", "thread01"} {
				if v := setup.idmap[_liveKey]; v == nil {
					liveMiss(t, inbox_thread_email_response_successDirectLiveStrict, "Live test blocked: needs %s via RESEND_TEST_INBOX_THREAD_EMAIL_RESPONSE_SUCCESS_ENTID", _liveKey)
					return
				}
			}
		}
		client := setup.client

		params := map[string]any{}
		query := map[string]any{}
		if setup.live {
			params["email_id"] = setup.idmap["email01"]
			params["inbox_id"] = setup.idmap["inbox01"]
			params["thread_id"] = setup.idmap["thread01"]
		} else {
			params["email_id"] = "direct01"
			params["inbox_id"] = "direct02"
			params["thread_id"] = "direct03"
		}

		result, err := client.Direct(map[string]any{
			"path":   "inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}",
			"method": "GET",
			"params": params,
			"query":  query,
		})
		if setup.live {
			if err != nil {
				liveMiss(t, inbox_thread_email_response_successDirectLiveStrict, "Live load failed: %v", err)
			}
			if status := core.ToInt(result["status"]); result["ok"] != true || status < 200 || status >= 300 {
				liveMiss(t, inbox_thread_email_response_successDirectLiveStrict, "Live load failed: %s", liveDescribe(result))
			}
			if result["data"] == nil {
				liveMiss(t, inbox_thread_email_response_successDirectLiveStrict, "Live load returned no data: %s", liveDescribe(result))
			}
		} else {
			if err != nil {
				t.Fatalf("direct failed: %v", err)
			}
			if result["ok"] != true {
				t.Fatalf("expected ok to be true, got %v", result["ok"])
			}
			if core.ToInt(result["status"]) != 200 {
				t.Fatalf("expected status 200, got %v", result["status"])
			}
			if result["data"] == nil {
				t.Fatal("expected data to be non-nil")
			}
		}

		if !setup.live {
			if dataMap, ok := result["data"].(map[string]any); ok {
				if dataMap["id"] != "direct01" {
					t.Fatalf("expected data.id to be direct01, got %v", dataMap["id"])
				}
			}

			if len(*setup.calls) != 1 {
				t.Fatalf("expected 1 call, got %d", len(*setup.calls))
			}
			call := (*setup.calls)[0]
			if initMap, ok := call["init"].(map[string]any); ok {
				if initMap["method"] != "GET" {
					t.Fatalf("expected method GET, got %v", initMap["method"])
				}
			}
			if url, ok := call["url"].(string); ok {
				if !strings.Contains(url, "direct01") {
					t.Fatalf("expected url to contain direct01, got %v", url)
				}
				if !strings.Contains(url, "direct02") {
					t.Fatalf("expected url to contain direct02, got %v", url)
				}
				if !strings.Contains(url, "direct03") {
					t.Fatalf("expected url to contain direct03, got %v", url)
				}
			}
		}
	})

}

type inbox_thread_email_response_successDirectSetupResult struct {
	client *sdk.ResendSDK
	calls  *[]map[string]any
	live   bool
	idmap  map[string]any
}

func inbox_thread_email_response_successDirectSetup(mockres any) *inbox_thread_email_response_successDirectSetupResult {
	loadEnvLocal()

	calls := &[]map[string]any{}

	env := envOverride(map[string]any{
		"RESEND_TEST_INBOX_THREAD_EMAIL_RESPONSE_SUCCESS_ENTID": map[string]any{},
		"RESEND_TEST_LIVE":    "FALSE",
		"RESEND_APIKEY":       "",
	})

	live := env["RESEND_TEST_LIVE"] == "TRUE"

	if live {
		// sdk-test-control.json's test.client.options seeds the live
		// client; the generated fields below overwrite anything they name.
		mergedOpts := map[string]any{}
		for k, v := range liveClientOptions() {
			mergedOpts[k] = v
		}
		for k, v := range map[string]any{
			"apikey": env["RESEND_APIKEY"],
		} {
			mergedOpts[k] = v
		}
		client := sdk.NewResendSDK(mergedOpts)

		idmap := map[string]any{}
		if entidRaw, ok := env["RESEND_TEST_INBOX_THREAD_EMAIL_RESPONSE_SUCCESS_ENTID"]; ok {
			if entidStr, ok := entidRaw.(string); ok && strings.HasPrefix(entidStr, "{") {
				json.Unmarshal([]byte(entidStr), &idmap)
			} else if entidMap, ok := entidRaw.(map[string]any); ok {
				idmap = entidMap
			}
		}

		return &inbox_thread_email_response_successDirectSetupResult{client: client, calls: calls, live: true, idmap: idmap}
	}

	mockFetch := func(url string, init map[string]any) (map[string]any, error) {
		*calls = append(*calls, map[string]any{"url": url, "init": init})
		return map[string]any{
			"status":     200,
			"statusText": "OK",
			"headers":    map[string]any{},
			"json": (func() any)(func() any {
				if mockres != nil {
					return mockres
				}
				return map[string]any{"id": "direct01"}
			}),
		}, nil
	}

	client := sdk.NewResendSDK(map[string]any{
		"base": "http://localhost:8080",
		"system": map[string]any{
			"fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
		},
	})

	return &inbox_thread_email_response_successDirectSetupResult{client: client, calls: calls, live: false, idmap: map[string]any{}}
}

var _ = os.Getenv
var _ = json.Unmarshal
