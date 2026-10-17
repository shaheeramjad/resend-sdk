package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewDebugFeatureFunc func() Feature

var NewIdempotencyFeatureFunc func() Feature

var NewMetricsFeatureFunc func() Feature

var NewNetsimFeatureFunc func() Feature

var NewPagingFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewApiKeyEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAttachmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAudienceEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAutomationEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAutomationRunEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewBroadcastEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewBroadcastClickedLinkEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewBroadcastRecipientEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactImportEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactPropertyEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactSegmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactTopicEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewDomainEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewDomainClaimEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEmailBatchEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEmailsMetricEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxAgentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxDraftEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxLabelEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadLabelEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewLogEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewOauthGrantEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewReceivedAttachmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewReceivedEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSegmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSegmentContactEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSuppressionEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSuppressionBatchAddEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSuppressionBatchRemoveEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewTemplateEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewTopicEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUsageEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewWebhookEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewWebhookEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewWebhookEventAttemptEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

