package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewDebugFeatureFunc func() Feature

var NewIdempotencyFeatureFunc func() Feature

var NewMetricsFeatureFunc func() Feature

var NewPagingFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewAddContactToSegmentResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAddInboxThreadLabelsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewApiKeyEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAudienceEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAutomationEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAutomationRunEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewAutomationRunListItemEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewBatchAddSuppressionsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewBatchRemoveSuppressionsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewBroadcastEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactImportResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactPropertyEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewContactTopicsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewCreateApiKeyEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewCreateBatchEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewCreateBroadcastResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewCreateInboxDraftResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewCreateInboxLabelResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewDomainEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewDomainClaimEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEmailsMetricEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxAgentResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxDraftListItemEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxDraftResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxLabelEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadEmailResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewInboxThreadResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListAttachmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListBroadcastClickedLinksResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListBroadcastRecipientsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListContactSegmentsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListContactsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListWebhookEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewListWebhookEventAttemptEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewLogEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewReceivedEmailEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveAudienceResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveBroadcastResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveContactFromSegmentResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveContactPropertyResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveContactResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveInboxDraftResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveInboxLabelResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveInboxResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveInboxThreadLabelResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveInboxThreadResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveSegmentResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveSuppressionResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveTemplateResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRemoveTopicResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRetrievedAttachmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRevokeOAuthGrantEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewRotateEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSegmentEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSetInboxThreadLabelsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewSuppressionEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewTemplateEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewTopicEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateApiKeyEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateBroadcastResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateContactPropertyResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateContactResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateContactTopicsResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateDomainResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateEmailOptionEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateInboxDraftResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateInboxLabelResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateInboxResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateInboxThreadResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateSegmentResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateTemplateResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateTopicResponseSuccessEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUpdateWebhookEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewUsageEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewWebhookEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

var NewWebhookEventEntityFunc func(client *ResendSDK, entopts map[string]any) ResendEntity

