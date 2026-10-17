import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { WebhookEventAttempt, WebhookEventAttemptListMatch } from '../ResendTypes';
declare class WebhookEventAttemptEntity extends ResendEntityBase<WebhookEventAttempt> {
    constructor(client: ResendSDK, entopts: any);
    make(this: WebhookEventAttemptEntity): WebhookEventAttemptEntity;
    list(this: any, reqmatch?: WebhookEventAttemptListMatch, ctrl?: Control): Promise<WebhookEventAttemptEntity[]>;
}
export { WebhookEventAttemptEntity };
