package voxgigresendsdk

import (
	"github.com/shaheeramjad/resend-sdk/go/core"
	"github.com/shaheeramjad/resend-sdk/go/entity"
	"github.com/shaheeramjad/resend-sdk/go/feature"
	_ "github.com/shaheeramjad/resend-sdk/go/utility"
)

// Type aliases preserve external API.
type ResendSDK = core.ResendSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type ResendEntity = core.ResendEntity
type StreamItem = core.StreamItem
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type ResendError = core.ResendError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewDebugFeatureFunc = func() core.Feature {
		return feature.NewDebugFeature()
	}
	core.NewIdempotencyFeatureFunc = func() core.Feature {
		return feature.NewIdempotencyFeature()
	}
	core.NewMetricsFeatureFunc = func() core.Feature {
		return feature.NewMetricsFeature()
	}
	core.NewNetsimFeatureFunc = func() core.Feature {
		return feature.NewNetsimFeature()
	}
	core.NewPagingFeatureFunc = func() core.Feature {
		return feature.NewPagingFeature()
	}
	core.NewRatelimitFeatureFunc = func() core.Feature {
		return feature.NewRatelimitFeature()
	}
	core.NewRetryFeatureFunc = func() core.Feature {
		return feature.NewRetryFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewTimeoutFeatureFunc = func() core.Feature {
		return feature.NewTimeoutFeature()
	}
	core.NewApiKeyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewApiKeyEntity(client, entopts)
	}
	core.NewAttachmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAttachmentEntity(client, entopts)
	}
	core.NewAudienceEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAudienceEntity(client, entopts)
	}
	core.NewAutomationEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAutomationEntity(client, entopts)
	}
	core.NewAutomationRunEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAutomationRunEntity(client, entopts)
	}
	core.NewBroadcastEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastEntity(client, entopts)
	}
	core.NewBroadcastClickedLinkEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastClickedLinkEntity(client, entopts)
	}
	core.NewBroadcastRecipientEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastRecipientEntity(client, entopts)
	}
	core.NewContactEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactEntity(client, entopts)
	}
	core.NewContactImportEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactImportEntity(client, entopts)
	}
	core.NewContactPropertyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactPropertyEntity(client, entopts)
	}
	core.NewContactSegmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactSegmentEntity(client, entopts)
	}
	core.NewContactTopicEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactTopicEntity(client, entopts)
	}
	core.NewDomainEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewDomainEntity(client, entopts)
	}
	core.NewDomainClaimEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewDomainClaimEntity(client, entopts)
	}
	core.NewEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailEntity(client, entopts)
	}
	core.NewEmailBatchEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailBatchEntity(client, entopts)
	}
	core.NewEmailsMetricEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailsMetricEntity(client, entopts)
	}
	core.NewEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEventEntity(client, entopts)
	}
	core.NewInboxEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxEntity(client, entopts)
	}
	core.NewInboxAgentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxAgentEntity(client, entopts)
	}
	core.NewInboxDraftEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxDraftEntity(client, entopts)
	}
	core.NewInboxLabelEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxLabelEntity(client, entopts)
	}
	core.NewInboxThreadEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadEntity(client, entopts)
	}
	core.NewInboxThreadEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadEmailEntity(client, entopts)
	}
	core.NewInboxThreadLabelEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadLabelEntity(client, entopts)
	}
	core.NewLogEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewLogEntity(client, entopts)
	}
	core.NewOauthGrantEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewOauthGrantEntity(client, entopts)
	}
	core.NewReceivedAttachmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewReceivedAttachmentEntity(client, entopts)
	}
	core.NewReceivedEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewReceivedEmailEntity(client, entopts)
	}
	core.NewSegmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSegmentEntity(client, entopts)
	}
	core.NewSegmentContactEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSegmentContactEntity(client, entopts)
	}
	core.NewSuppressionEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSuppressionEntity(client, entopts)
	}
	core.NewSuppressionBatchAddEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSuppressionBatchAddEntity(client, entopts)
	}
	core.NewSuppressionBatchRemoveEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSuppressionBatchRemoveEntity(client, entopts)
	}
	core.NewTemplateEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewTemplateEntity(client, entopts)
	}
	core.NewTopicEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewTopicEntity(client, entopts)
	}
	core.NewUsageEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUsageEntity(client, entopts)
	}
	core.NewWebhookEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewWebhookEntity(client, entopts)
	}
	core.NewWebhookEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewWebhookEventEntity(client, entopts)
	}
	core.NewWebhookEventAttemptEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewWebhookEventAttemptEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewResendSDK = core.NewResendSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewResendSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *ResendSDK  { return NewResendSDK(nil) }
func Test() *ResendSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewDebugFeature = feature.NewDebugFeature
var NewIdempotencyFeature = feature.NewIdempotencyFeature
var NewMetricsFeature = feature.NewMetricsFeature
var NewNetsimFeature = feature.NewNetsimFeature
var NewPagingFeature = feature.NewPagingFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
