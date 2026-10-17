# resend-mcp

[MCP](https://modelcontextprotocol.io) server exposing the Resend SDK as
2 agent tools, `resend_list` and `resend_load`, built on the
[official Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk) and the
sibling Go SDK at `../go`. Runs over **stdio** (default, for spawnable installs)
or **streamable HTTP** (one shared server for several agents).

The server only reads. Create, update, patch and remove become tools too when the
SDK's own model sets `main: kit: target: 'go-mcp': tool: write: true`; they are off by default, as an agent calling
them changes the API's data.

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/resend-mcp)
make build

# 2. Provide credentials via the environment
export RESEND_APIKEY=sk_live_xxx

# 3a. Install into Claude Code over stdio (most common)
claude mcp add --scope user resend \
  -- /absolute/path/to/resend-mcp -transport stdio

# 3b. …or run a shared HTTP server instead
./resend-mcp -transport http -addr :8080
```

Tool-call arguments (what an agent sends):

```jsonc
// resend_list: first page of records
{ "entity": "api_key" }

// resend_load: one record
{ "entity": "audience", "query": { "id": "example_id" } }
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: install and call a tool

1. **Build** the server from this `go-mcp/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/resend-mcp
   ```

2. **Set your API key:**

   ```sh
   export RESEND_APIKEY=sk_live_xxx
   ```

3. **Install it into Claude Code** (stdio transport):

   ```sh
   claude mcp add --scope user resend \
     -- "$PWD"/dist/*/resend-mcp -transport stdio
   ```

4. **Restart Claude Code.** The `resend_list` and `resend_load` tools now appear in new
   sessions. Ask the agent to *"list api_key using resend"*
   and it calls `resend_list` with `{"entity":"api_key"}`.

## How-to guides

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export RESEND_APIKEY=sk_live_xxx            # API key
export RESEND_BASE=https://api.example.com  # optional: override the API base URL
```

Set these in the shell that launches the server (or in the `claude mcp add`
environment) so every tool call is authenticated.

### Run as a shared HTTP server

```sh
./resend-mcp -transport http -addr :8080
```

Streamable HTTP lets several agents share one running process; stdio (the
default) spawns a fresh process per client.

### Call the `resend_list` tool

Args: `entity` (required), `query` (optional: optional filter map; omit it for the first page).
Returns the first page of records as JSON:

```jsonc
{ "entity": "api_key" }
```

### Call the `resend_load` tool

Args: `entity` (required), `query` (required: match map naming the record, such as {"id":1}).
Returns the record as JSON:

```jsonc
{ "entity": "audience", "query": { "id": "example_id" } }
```

### Turn on the write tools

In the SDK's own model (`.sdk/model/sdk.aontu`), then regenerate:

```
main: kit: target: 'go-mcp': tool: write: true
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

## Reference

### Tools

| Tool | Args | Returns | MCP hints |
|------|------|---------|-----------|
| `resend_list` | `entity`, `query` (optional map) | The first page of records as JSON | read-only |
| `resend_load` | `entity`, `query` (required map) | The record as JSON | read-only |

On error, a tool returns an MCP error result (`isError: true`) whose text is the
failure message (e.g. unknown entity, or an API error).

### Entities

Each tool takes as its `entity` argument one of the entities that has its
operation, of the 84 the SDK has:

| Tool | Entities |
|------|----------|
| `resend_list` | api_key, audience, automation, automation_run_list_item, broadcast, contact, contact_import_response_success, contact_property, contact_topics_response_success, domain, email, emails_metric, event, inbox, inbox_agent_response_success, inbox_draft_list_item, inbox_label, inbox_thread, inbox_thread_email, list_attachment, list_broadcast_clicked_links_response_success, list_broadcast_recipients_response_success, list_contact_segments_response_success, list_contacts_response_success, list_webhook_event, list_webhook_event_attempt, log, received_email, remove_segment_response_success, remove_suppression_response_success, remove_template_response_success, remove_topic_response_success, revoke_o_auth_grant, update_webhook |
| `resend_load` | audience, automation, automation_run, broadcast, contact, contact_import_response_success, contact_property, domain, domain_claim, email, event, inbox, inbox_draft_response_success, inbox_thread_email_response_success, inbox_thread_response_success, log, received_email, retrieved_attachment, segment, suppression, template, topic, usage, webhook, webhook_event |

JSON schemas are emitted by the SDK from each tool's argument struct's
`json` / `jsonschema` tags — no schema is hand-written. Each tool's
`entity` is an `enum` of the entities in its row, so the server refuses
any other before it runs a call.

### Transports & flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-transport` | `stdio` | `stdio` (spawnable) or `http` (streamable HTTP). |
| `-addr` | `:8080` | Listen address for the `http` transport. |

### Environment variables

| Variable | Purpose |
|----------|---------|
| `RESEND_APIKEY` | API key sent with every request. |
| `RESEND_BASE` | Optional override of the API base URL. |

### Smoke test via HTTP (raw JSON-RPC)

```sh
./resend-mcp -transport http -addr :18080 &

# initialize, grab the session id
curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -D headers \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"smoke","version":"0"}}}'

SESSION=$(awk '/Mcp-Session-Id/ {print $2}' headers | tr -d '\r')

curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -H "Mcp-Session-Id: $SESSION" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"resend_list","arguments":{"entity":"api_key"}}}'
```

## Explanation

### How tools map to the SDK

`main.go` builds the SDK client (configured from the environment) and registers
one tool per operation the SDK's entities have. Each dispatches on the
`entity` argument to the matching entity in the sibling Go SDK at `../go`,
calls its operation, unwraps the `Entity` wrappers to plain data, and returns
it as pretty-printed JSON.

### Why two transports

**stdio** is the standard for agent hosts that spawn a server per client
(Claude Code's `claude mcp add`). **streamable HTTP** keeps one process running
that many agents can share — handy for a long-lived deployment.

### Schema generation

The input schema is derived from each tool's argument struct's `json` /
`jsonschema` tags at registration time, so the advertised tool schema can
never drift from the code that consumes it. The `entity` enum comes from the
same list the tool is registered for.

## Generated by

sdkgen `go-mcp` target. See the target source under `.sdk/src/cmp/go-mcp/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-mcp/`.
