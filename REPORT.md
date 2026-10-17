# Building a Resend SDK with the Voxgig SDK generator

Observations from building this repository with `@voxgig/sdkgen`, written
for the Voxgig team.

## Summary

The generator took the official Resend OpenAPI spec to TypeScript, Python and
Go SDKs, a Go CLI and a Go MCP server in about 45 seconds of tool time
(scaffold, add targets and features, generate), and the tooling is fast,
deterministic and very well documented. The
weak point is the entity model the spec produces: out of the box it gave 84
entities, many named after response schemas such as
`RemoveTemplateResponseSuccess`, with one resource's operations scattered
across several of them. That split also caused one generated test to fail in
all three languages. Correcting it in `guide.aontu` worked exactly as the docs
describe and brought the SDK to 41 resource-shaped entities, with every test
green.

## What I built

- **API:** [Resend](https://resend.com), the email API. It has no SDK in the
  voxgig-sdk catalogue (checked against all 804 repositories), publishes an
  official [OpenAPI 3.1.2 spec](https://github.com/resend/resend-openapi)
  (v1.5.1: 89 paths, 142 operations), uses a Bearer API key, and has a free
  tier that needs no card.
- **Toolchain:** create-sdkgen 0.30.8, sdkgen 4.35.0, apidef 8.25.3,
  model 12.0.1, docgen 0.30.2, Node 24.22.
- **Targets:** `ts`, `py`, `go`, `go-cli`, `go-mcp`. Only these were built and
  tested, because PHP, Ruby and Lua toolchains were not installed on the
  build machine.
- **Features:** `retry`, `timeout`, `idempotency`, `ratelimit`, `paging`,
  `metrics`, `debug`, `netsim`, plus the default `test`.
- **How:** an AI coding agent (Cursor) ran the generator by following
  create-sdkgen's agent guide. All SDK source is generator output. The hand
  work is the model files in `.sdk/model/` (`guide.aontu`, `project.aontu`,
  `text.aontu`) and the live smoke tests in `live/`.

## Timeline

Times are wall-clock time on the build machine.

| Step | Time |
| --- | --- |
| Scaffold (`npm create @voxgig/sdkgen`, including `npm install`) | 37 s |
| `target add` five targets, then `feature add` seven features | 2 s |
| First `npm run generate` (1,234 files) | 5 s |
| Build and test: TypeScript, Python, Go | 32 s, 19 s, 23 s |
| Build the CLI and the MCP server | 9 s, 14 s |
| Correct the entity model in `guide.aontu`, regenerate, retest | about 10 min |
| End to end, from picking the API to a green, identity-corrected SDK | about 35 min |

## Test results

| | TypeScript | Python | Go |
| --- | --- | --- | --- |
| Unmodified output | 710 pass, **1 fail**, 90 skip | 530 pass, **1 fail**, 157 skip | **1 fail** |
| After the guide fix, with `netsim` | 637 pass, 0 fail, 39 skip | 466 pass, 0 fail, 96 skip | ok |

The fixed SDK has fewer tests because it has fewer entities. Every remaining
skip but one is for a feature this SDK does not include (`validate`, `cost`,
`cache` and similar), and each says so in its skip reason. The exception is a
Python type check that needs `mypy` installed.

## What worked well

- **Speed and determinism.** Generation takes about 5 seconds. Regenerating
  with no model change produced zero diff across 1,835 tracked files.
- **The documentation is excellent.** The apidef guide reference explains
  every classification rule with real-API examples, and the "Correcting the
  guide" section told me exactly how to fix the entity model. Each fix worked
  on the first try.
- **Model-first customisation holds up.** Repository identity, publisher,
  package names, README tagline and descriptions are all set in
  `project.aontu` and `text.aontu`, and they flowed into every manifest,
  `go.mod`, LICENSE and README. Hand-written files (`live/`, this report)
  survive regeneration, and switching an entity off also deletes its
  generated files.
- **Doc examples are tested.** Between 271 and 312 README and REFERENCE code
  blocks per language are compiled, and the runnable ones are executed.
- **Safe defaults.** Every feature is off until enabled. The MCP server
  exposes only read tools (`resend_list`, `resend_load`); write tools are
  opt-in.
- **Consistent behaviour across languages.** The same test failed in the same
  way in all three languages, and the TypeScript and Python SDKs return the
  same error text. The parity claim holds.
- **Clear warnings.** For example, apidef warned that thread search could
  never be selected because it shared a selector with thread list.

## Issues and bugs

Ordered by impact.

1. **Entity naming follows response schemas, not resources.** Resend names
   responses per operation (`UpdateContactResponseSuccess`,
   `RemoveContactResponseSuccess`), and the heuristic turned many of those
   names into entities, giving 84 entities where 41 describe the API. The worst
   cases moved whole collections to the wrong entity:
   `remove_template_response_success` owned template create, list and
   remove, and `update_webhook` owned webhook create, list and update.
   Others included `create_api_key`, `revoke_o_auth_grant` and `rotate`. This
   naming reaches the SDK, CLI and MCP surfaces, for example
   `client.RemoveTemplateResponseSuccess().create(...)`. Fixing it took a
   359-line `guide.aontu`.
   *Suggestion:* treat verb prefixes (`Create`, `Update`, `Remove`, `List`,
   `Get`) and suffixes (`ResponseSuccess`, `Response`) as wrapper noise, the
   way `UserResponse` is already cleaned. Prefer the path's resource when one
   path's methods would otherwise split across entities.

2. **One generated test fails out of the box, in all three languages.** The
   path `/suppressions/{suppression}` was split across two entities. The GET
   side renamed `suppression` to `id`, but the DELETE side did not, so the
   generated basic flow deleted nothing and then found the record still
   listed. A user's first `npm test` is red without any obvious reason.

3. **API error messages are dropped.** For a bad key, Resend answers
   `{"message":"API key is invalid"}`, but the SDK error's `message` (and the
   CLI's output) is `ResendSDK: list: request: 400: Bad Request`. The API's
   text is only in `err.result.body.message`. Appending `body.message` (or
   `error`/`detail`) when present would make errors much easier to act on.

4. **Python typed models omit `from`.** `from` is a Python keyword, so it is
   left out of the generated `TypedDict`s (18 warnings on the first
   generation). For an email API, the
   sender is the most important field of the most important call. The
   functional form `TypedDict('EmailCreateData', {'from': str, ...})`
   supports keyword keys.

5. **Switching entities off leaves orphan flow files.** After the guide fix,
   59 `model/flow/*.aontu` files for disabled entities stayed on disk.
   `doctor` listed them, then printed ".sdk has drifted: 0 forked, 0 edited,
   0 stale, 0 missing, 0 outdated, 0 unresolved" and exited non-zero. The
   summary should count orphans, and `generate` could remove them as it does
   for entity files.

6. **Default identity is Voxgig's.** `create-sdkgen` hard-codes
   `origin: 'voxgig-sdk'`, so the package names, Go module paths and URLs
   point at `github.com/voxgig-sdk/resend-sdk`. The LICENSE copyright and
   SECURITY.md contact default to Voxgig, and the scaffold's own
   `.sdk/tm/LICENSE` says "Copyright (c) 2026 Resend", which is the API
   vendor. All of this is fixable in `project.aontu`, but a scaffold flag or
   prompt for owner and publisher would prevent shipping it by accident.

7. **The live suite has no safe mode.** `RESEND_TEST_LIVE=TRUE` runs create,
   update and remove flows for every entity with synthetic data. On a real
   account that would create API keys and domains and send email to made-up
   addresses. `sdk-test-control.json` only has a skip list. An allowlist or
   a read-only live mode would make live testing of SaaS APIs practical, so I
   wrote a curated smoke test instead (`live/smoke.js`, `live/smoke.py`).

8. **The retry and timeout tests need `netsim`, which is not suggested.**
   With `retry` and `timeout` enabled but no `netsim`, the tests that
   exercise them are skipped. `feature add retry` could suggest
   `feature add netsim`, or add it.

9. **Smaller items.**
   - Apidef printed "PARTIAL BUILD! There were 1 warnings", but `generate`
     exited 0.
   - In `text.aontu`, `about_md` is Markdown, but `"""` strings fail with
     "unprintable character". Backtick strings work but clash with
     Markdown inline code. The reference could show a multi-line example.
   - Node: the tutorial says "a recent LTS" and the agent guide says
     Node 24+. sdkgen and aontu declare `>=24`, but on Node 22 the install
     only prints `EBADENGINE` warnings, and scaffold and generate succeed.
   - The CLI and live tests read `RESEND_APIKEY`, while Resend's convention
     is `RESEND_API_KEY`. The MCP README's placeholder key is `sk_live_xxx`
     (Stripe style); Resend keys start with `re_`.
   - The README quickstart uses the alphabetically first entity (`ApiKey`).
     For this API, `Email().create()` is the natural first example.
   - The Go SDK's `go.mod` says `go 1.21` while the CLI and MCP modules say
     `go 1.24.7`.
   - 6 generated flow steps are switched off as unreachable. They are nested
     routes that need two identifiers, such as adding a contact to a
     segment.

## Live testing

`live/smoke.js` (TypeScript) and `live/smoke.py` (Python) call the real API
with `retry`, `timeout`, `ratelimit`, `metrics` and `debug` switched on. They
read usage, domains and API keys; send one email to Resend's sandbox address
`delivered@resend.dev`; load it back; create, update, load and remove a
contact; create, load and remove a segment; list templates, webhooks,
topics, broadcasts and logs; make a `direct()` call; and check that a bad key
is rejected. Both were checked against the live endpoint with an invalid key:
every request reached Resend and was rejected on authentication.

Live run with a real key: *pending.*

## Repository map

- `.sdk/model/guide/guide.aontu`: the entity corrections (issue 1).
- `.sdk/model/project.aontu`, `.sdk/model/text.aontu`: identity and README
  wording.
- `.sdk/def/resend-openapi.yaml`: the unmodified upstream spec.
- `ts/`, `py/`, `go/`, `go-cli/`, `go-mcp/`: generated output.
- `live/`: hand-written live smoke tests.
