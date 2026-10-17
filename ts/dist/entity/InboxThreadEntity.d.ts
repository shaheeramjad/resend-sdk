import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { InboxThread, InboxThreadLoadMatch, InboxThreadListMatch, InboxThreadUpdateData, InboxThreadRemoveMatch } from '../ResendTypes';
declare class InboxThreadEntity extends ResendEntityBase<InboxThread> {
    constructor(client: ResendSDK, entopts: any);
    make(this: InboxThreadEntity): InboxThreadEntity;
    load(this: any, reqmatch?: InboxThreadLoadMatch, ctrl?: Control): Promise<InboxThreadEntity>;
    list(this: any, reqmatch?: InboxThreadListMatch, ctrl?: Control): Promise<InboxThreadEntity[]>;
    update(this: any, reqdata?: InboxThreadUpdateData, ctrl?: Control): Promise<InboxThreadEntity>;
    remove(this: any, reqmatch?: InboxThreadRemoveMatch, ctrl?: Control): Promise<InboxThreadEntity>;
}
export { InboxThreadEntity };
