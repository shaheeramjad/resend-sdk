# resend-cli

boru-driven command-line client **and** interactive REPL for the Resend
SDK. Each command line is parsed as a single [boru](https://github.com/boru-lang/boru)
expression and evaluated against the live API; run it with no arguments to drop
into a REPL. Built on `github.com/boru-lang/boru/eng/go` and the sibling Go SDK
at `../go`.

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/resend-cli)
make build

# 2. See usage (words, entities, env vars)
./resend-cli --help

# 3. Provide credentials once, via the environment
export RESEND_APIKEY=sk_live_xxx

# 4. Each command line is ONE boru expression, run against the API:
./resend-cli list api_key

# 5. Override the API base URL for a single call
RESEND_BASE=https://api.example.com ./resend-cli --help

# 6. No arguments -> interactive REPL
./resend-cli
resend> /help
resend> /quit
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: your first query in under a minute

1. **Build the binary.** From this `go-cli/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/resend-cli
   ```

2. **Set your API key** (read from the environment):

   ```sh
   export RESEND_APIKEY=sk_live_xxx
   ```

3. **Run a query.** Evaluate an boru expression against the API (or run with no
   arguments to open the REPL):

   ```sh
   ./dist/*/resend-cli --help
   ```

4. **Go interactive.** Run the binary with no arguments to open the REPL, then
   type `/help` for the word and entity lists and `/quit` to leave.

That is the whole loop: *build → set key → evaluate boru expressions*.

## How-to guides

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export RESEND_APIKEY=sk_live_xxx            # API key
export RESEND_BASE=https://api.example.com  # optional: override the API base URL
./resend-cli --help
```

Both are injectable by a secrets vault, so the key never has to be typed inline.

### Explore interactively with the REPL

Run with no arguments to open a REPL (prompt `resend>`). Each line is
evaluated as its own boru expression:

```text
$ ./resend-cli
resend> /help
resend> /quit
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

### Discover the available entities

`/help` in the REPL prints the full entity list, or see [Entities](#entities)
below — this SDK exposes 84 entities.

## Reference

### Words

The CLI registers these boru words, each bound to the SDK:

| Word     | Signatures                                    | Returns                        |
|----------|-----------------------------------------------|--------------------------------|
| `list`   | `list <entity>` · `list <query> <entity>`     | First page of records          |
| `load`   | `load <entity>` · `load <query> <entity>`     | A single record                |
| `update` | `update <query> <entity>`                     | Update a record, return it     |

- `<entity>` is a bareword, auto-quoted as an boru atom (e.g. `add_contact_to_segment_response_success`).
- `<query>` is either a **Map** (`{id:1}`) or a **Scalar** (`1`, treated as
  `{id:1}`). A scalar is always wrapped as `{id:<value>}`.

### Environment variables

| Variable | Purpose |
|----------|---------|
| `RESEND_APIKEY` | API key sent with every request. |
| `RESEND_BASE` | Optional override of the API base URL. |

Unset variables fall back to the SDK's built-in defaults.

### CLI flags

- `--help` / `-h` — print usage (words, entities, env vars) and exit.

### REPL commands

Meta-commands use the `/` prefix (everything else on a line is evaluated as boru):

- `/quit` / `/q` / `/exit` — exit the REPL
- `/help` / `/h` / `/?`     — show the word list, entity list and meta commands

### Exit codes

| Code | Meaning |
|------|---------|
| `0` | Success (also the normal REPL exit). |
| `1` | Parse error, word-registration error, or an API/evaluation error. |

### Build targets

| Target | Result |
|--------|--------|
| `make build` | Native binary at `dist/<os>-<arch>/resend-cli`. |
| `make build-all` | linux/darwin/windows x amd64/arm64, each under its own `dist/<os>-<arch>/`. |
| `make clean` | Remove `dist/` and any stray binaries. |

### Entities

The 84 entities this SDK exposes (any is valid as `<entity>`):

add_contact_to_segment_response_success add_inbox_thread_labels_response_success api_key audience automation automation_run automation_run_list_item batch_add_suppressions_response_success batch_remove_suppressions_response_success broadcast contact contact_import_response_success contact_property contact_topics_response_success create_api_key create_batch_email create_broadcast_response_success create_inbox_draft_response_success create_inbox_label_response_success domain domain_claim email emails_metric event inbox inbox_agent_response_success inbox_draft_list_item inbox_draft_response_success inbox_label inbox_thread inbox_thread_email inbox_thread_email_response_success inbox_thread_response_success list_attachment list_broadcast_clicked_links_response_success list_broadcast_recipients_response_success list_contact_segments_response_success list_contacts_response_success list_webhook_event list_webhook_event_attempt log received_email remove_audience_response_success remove_broadcast_response_success remove_contact_from_segment_response_success remove_contact_property_response_success remove_contact_response_success remove_event remove_inbox_draft_response_success remove_inbox_label_response_success remove_inbox_response_success remove_inbox_thread_label_response_success remove_inbox_thread_response_success remove_segment_response_success remove_suppression_response_success remove_template_response_success remove_topic_response_success retrieved_attachment revoke_o_auth_grant rotate segment set_inbox_thread_labels_response_success suppression template topic update_api_key update_broadcast_response_success update_contact_property_response_success update_contact_response_success update_contact_topics_response_success update_domain_response_success update_email_option update_event update_inbox_draft_response_success update_inbox_label_response_success update_inbox_response_success update_inbox_thread_response_success update_segment_response_success update_template_response_success update_topic_response_success update_webhook usage webhook webhook_event

## Explanation

### Why boru?

The whole command line is one [boru](https://github.com/boru-lang/boru) expression,
not a fixed `verb --flag` grammar. That means the same binary works one-shot
(`./resend-cli <expr>`) and interactively (the REPL), and expressions compose the
same way in both. `list` / `load` / `update` are ordinary boru *words* bound to
the SDK — adding SDK operations is adding words, not re-parsing flags.

### How it is wired

`main.go` builds the SDK client (configured from the environment), creates an
boru registry, and `words.go` registers `list` / `load` / `update` as native
words that dispatch on the entity atom and call the sibling Go SDK at `../go`.
Results are unwrapped from their `Entity` wrappers to plain data before being
printed.

### Output format

Each result value is printed as its boru string form (a JSON-like rendering of
the record or list of records). One-shot mode prints to stdout; errors go to
stderr with a non-zero exit code.

## Generated by

sdkgen `go-cli` target. See the target source under `.sdk/src/cmp/go-cli/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-cli/`.
