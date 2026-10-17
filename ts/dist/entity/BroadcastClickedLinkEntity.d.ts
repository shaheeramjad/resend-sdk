import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { BroadcastClickedLink, BroadcastClickedLinkListMatch } from '../ResendTypes';
declare class BroadcastClickedLinkEntity extends ResendEntityBase<BroadcastClickedLink> {
    constructor(client: ResendSDK, entopts: any);
    make(this: BroadcastClickedLinkEntity): BroadcastClickedLinkEntity;
    list(this: any, reqmatch?: BroadcastClickedLinkListMatch, ctrl?: Control): Promise<BroadcastClickedLinkEntity[]>;
}
export { BroadcastClickedLinkEntity };
