import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { BroadcastRecipient, BroadcastRecipientListMatch } from '../ResendTypes';
declare class BroadcastRecipientEntity extends ResendEntityBase<BroadcastRecipient> {
    constructor(client: ResendSDK, entopts: any);
    make(this: BroadcastRecipientEntity): BroadcastRecipientEntity;
    list(this: any, reqmatch?: BroadcastRecipientListMatch, ctrl?: Control): Promise<BroadcastRecipientEntity[]>;
}
export { BroadcastRecipientEntity };
