# Resend SDK feature factory

from resend_sdk.feature.base_feature import ResendBaseFeature
from resend_sdk.feature.debug_feature import ResendDebugFeature
from resend_sdk.feature.idempotency_feature import ResendIdempotencyFeature
from resend_sdk.feature.metrics_feature import ResendMetricsFeature
from resend_sdk.feature.paging_feature import ResendPagingFeature
from resend_sdk.feature.ratelimit_feature import ResendRatelimitFeature
from resend_sdk.feature.retry_feature import ResendRetryFeature
from resend_sdk.feature.test_feature import ResendTestFeature
from resend_sdk.feature.timeout_feature import ResendTimeoutFeature


_FEATURES = {
    "base": lambda: ResendBaseFeature(),
    "debug": lambda: ResendDebugFeature(),
    "idempotency": lambda: ResendIdempotencyFeature(),
    "metrics": lambda: ResendMetricsFeature(),
    "paging": lambda: ResendPagingFeature(),
    "ratelimit": lambda: ResendRatelimitFeature(),
    "retry": lambda: ResendRetryFeature(),
    "test": lambda: ResendTestFeature(),
    "timeout": lambda: ResendTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
