package voxgigresendsdk

import (
	"github.com/voxgig-sdk/resend-sdk/go/core"
	"github.com/voxgig-sdk/resend-sdk/go/entity"
	"github.com/voxgig-sdk/resend-sdk/go/feature"
	_ "github.com/voxgig-sdk/resend-sdk/go/utility"
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
	core.NewAddContactToSegmentResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAddContactToSegmentResponseSuccessEntity(client, entopts)
	}
	core.NewAddInboxThreadLabelsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAddInboxThreadLabelsResponseSuccessEntity(client, entopts)
	}
	core.NewApiKeyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewApiKeyEntity(client, entopts)
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
	core.NewAutomationRunListItemEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewAutomationRunListItemEntity(client, entopts)
	}
	core.NewBatchAddSuppressionsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBatchAddSuppressionsResponseSuccessEntity(client, entopts)
	}
	core.NewBatchRemoveSuppressionsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBatchRemoveSuppressionsResponseSuccessEntity(client, entopts)
	}
	core.NewBroadcastEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewBroadcastEntity(client, entopts)
	}
	core.NewContactEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactEntity(client, entopts)
	}
	core.NewContactImportResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactImportResponseSuccessEntity(client, entopts)
	}
	core.NewContactPropertyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactPropertyEntity(client, entopts)
	}
	core.NewContactTopicsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewContactTopicsResponseSuccessEntity(client, entopts)
	}
	core.NewCreateApiKeyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewCreateApiKeyEntity(client, entopts)
	}
	core.NewCreateBatchEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewCreateBatchEmailEntity(client, entopts)
	}
	core.NewCreateBroadcastResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewCreateBroadcastResponseSuccessEntity(client, entopts)
	}
	core.NewCreateInboxDraftResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewCreateInboxDraftResponseSuccessEntity(client, entopts)
	}
	core.NewCreateInboxLabelResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewCreateInboxLabelResponseSuccessEntity(client, entopts)
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
	core.NewEmailsMetricEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEmailsMetricEntity(client, entopts)
	}
	core.NewEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewEventEntity(client, entopts)
	}
	core.NewInboxEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxEntity(client, entopts)
	}
	core.NewInboxAgentResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxAgentResponseSuccessEntity(client, entopts)
	}
	core.NewInboxDraftListItemEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxDraftListItemEntity(client, entopts)
	}
	core.NewInboxDraftResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxDraftResponseSuccessEntity(client, entopts)
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
	core.NewInboxThreadEmailResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadEmailResponseSuccessEntity(client, entopts)
	}
	core.NewInboxThreadResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewInboxThreadResponseSuccessEntity(client, entopts)
	}
	core.NewListAttachmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListAttachmentEntity(client, entopts)
	}
	core.NewListBroadcastClickedLinksResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListBroadcastClickedLinksResponseSuccessEntity(client, entopts)
	}
	core.NewListBroadcastRecipientsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListBroadcastRecipientsResponseSuccessEntity(client, entopts)
	}
	core.NewListContactSegmentsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListContactSegmentsResponseSuccessEntity(client, entopts)
	}
	core.NewListContactsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListContactsResponseSuccessEntity(client, entopts)
	}
	core.NewListWebhookEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListWebhookEventEntity(client, entopts)
	}
	core.NewListWebhookEventAttemptEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewListWebhookEventAttemptEntity(client, entopts)
	}
	core.NewLogEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewLogEntity(client, entopts)
	}
	core.NewReceivedEmailEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewReceivedEmailEntity(client, entopts)
	}
	core.NewRemoveAudienceResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveAudienceResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveBroadcastResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveBroadcastResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveContactFromSegmentResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveContactFromSegmentResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveContactPropertyResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveContactPropertyResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveContactResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveContactResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveEventEntity(client, entopts)
	}
	core.NewRemoveInboxDraftResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveInboxDraftResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveInboxLabelResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveInboxLabelResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveInboxResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveInboxResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveInboxThreadLabelResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveInboxThreadLabelResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveInboxThreadResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveInboxThreadResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveSegmentResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveSegmentResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveSuppressionResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveSuppressionResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveTemplateResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveTemplateResponseSuccessEntity(client, entopts)
	}
	core.NewRemoveTopicResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRemoveTopicResponseSuccessEntity(client, entopts)
	}
	core.NewRetrievedAttachmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRetrievedAttachmentEntity(client, entopts)
	}
	core.NewRevokeOAuthGrantEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRevokeOAuthGrantEntity(client, entopts)
	}
	core.NewRotateEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewRotateEntity(client, entopts)
	}
	core.NewSegmentEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSegmentEntity(client, entopts)
	}
	core.NewSetInboxThreadLabelsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSetInboxThreadLabelsResponseSuccessEntity(client, entopts)
	}
	core.NewSuppressionEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewSuppressionEntity(client, entopts)
	}
	core.NewTemplateEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewTemplateEntity(client, entopts)
	}
	core.NewTopicEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewTopicEntity(client, entopts)
	}
	core.NewUpdateApiKeyEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateApiKeyEntity(client, entopts)
	}
	core.NewUpdateBroadcastResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateBroadcastResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateContactPropertyResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateContactPropertyResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateContactResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateContactResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateContactTopicsResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateContactTopicsResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateDomainResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateDomainResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateEmailOptionEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateEmailOptionEntity(client, entopts)
	}
	core.NewUpdateEventEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateEventEntity(client, entopts)
	}
	core.NewUpdateInboxDraftResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateInboxDraftResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateInboxLabelResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateInboxLabelResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateInboxResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateInboxResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateInboxThreadResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateInboxThreadResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateSegmentResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateSegmentResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateTemplateResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateTemplateResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateTopicResponseSuccessEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateTopicResponseSuccessEntity(client, entopts)
	}
	core.NewUpdateWebhookEntityFunc = func(client *core.ResendSDK, entopts map[string]any) core.ResendEntity {
		return entity.NewUpdateWebhookEntity(client, entopts)
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
var NewPagingFeature = feature.NewPagingFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
