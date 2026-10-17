# Live smoke test of the generated Python SDK against api.resend.com.
#
#   pip install -e py
#   RESEND_APIKEY=re_... python live/smoke.py
#
# Needs a FULL ACCESS key. Sends one email from onboarding@resend.dev to
# Resend's sandbox inbox delivered@resend.dev unless RESEND_SMOKE_TO names
# another address. The one contact it creates is removed at the end.

import os
import sys
import time

from resend_sdk import ResendSDK

apikey = os.environ.get("RESEND_APIKEY")
if not apikey:
    print("Set RESEND_APIKEY to a Resend API key (https://resend.com/api-keys).")
    sys.exit(2)

to = os.environ.get("RESEND_SMOKE_TO", "delivered@resend.dev")
stamp = int(time.time() * 1000)

client = ResendSDK({
    "apikey": apikey,
    "feature": {
        "retry": {"active": True, "retries": 3, "maxDelay": 4000},
        "timeout": {"active": True, "ms": 15000},
        "ratelimit": {"active": True, "rate": 2, "burst": 2},
    },
})

results = []
state = {}


def step(name, fn):
    try:
        detail = fn()
        results.append((name, True))
        print(f"PASS  {name}" + (f"  {detail}" if detail else ""))
    except Exception as err:
        results.append((name, False))
        print(f"FAIL  {name}  {err}")


def usage_load():
    usage = client.Usage().load().data_get()
    assert isinstance(usage, dict), "no usage record"


def domain_list():
    return f"{len(client.Domain().list())} domain(s)"


def email_create():
    sent = client.Email().create({
        "from": "Voxgig Resend SDK <onboarding@resend.dev>",
        "to": [to],
        "subject": f"Resend SDK smoke test (python) {stamp}",
        "html": "<p>Sent by the generated Voxgig Resend Python SDK.</p>",
    }).data_get()
    assert sent.get("id"), "no id in send response"
    state["email_id"] = sent["id"]
    return f"id={sent['id']} to={to}"


def email_load():
    assert state.get("email_id"), "skipped: no email was sent"
    email = client.Email().load({"id": state["email_id"]}).data_get()
    assert email["id"] == state["email_id"], "loaded the wrong email"
    return f"last_event={email.get('last_event')}"


contact_email = f"sdk-smoke-py-{stamp}@example.com"


def contact_create():
    created = client.Contact().create({"email": contact_email, "first_name": "Smoke"}).data_get()
    assert created.get("id"), "no id in create response"
    state["contact_id"] = created["id"]
    return f"id={created['id']}"


def contact_update():
    assert state.get("contact_id"), "skipped: no contact was created"
    client.Contact().update({"id": state["contact_id"], "first_name": "Smoked"})


def contact_load():
    assert state.get("contact_id"), "skipped: no contact was created"
    contact = client.Contact().load({"id": state["contact_id"]}).data_get()
    assert contact["email"] == contact_email, f"email {contact['email']} != {contact_email}"
    assert contact["first_name"] == "Smoked", f"first_name {contact['first_name']} != Smoked"
    return f"first_name={contact['first_name']}"


def contact_remove():
    assert state.get("contact_id"), "skipped: no contact was created"
    client.Contact().remove({"id": state["contact_id"]})


def bad_key():
    bad = ResendSDK({"apikey": "re_invalid_key_for_smoke_test"})
    try:
        bad.Domain().list()
    except Exception as err:
        return f"rejected: {err}"
    raise AssertionError("an invalid key was accepted")


step("Usage.load", usage_load)
step("Domain.list", domain_list)
step("Email.create (send)", email_create)
step("Email.load", email_load)
step("Contact.create", contact_create)
step("Contact.update", contact_update)
step("Contact.load", contact_load)
step("Contact.remove", contact_remove)
step("bad key is rejected", bad_key)

passed = sum(1 for _, ok in results if ok)
print(f"\n{passed}/{len(results)} steps passed")
sys.exit(0 if passed == len(results) else 1)
