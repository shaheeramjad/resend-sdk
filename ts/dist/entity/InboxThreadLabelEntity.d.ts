import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { InboxThreadLabel, InboxThreadLabelCreateData, InboxThreadLabelUpdateData, InboxThreadLabelRemoveMatch } from '../ResendTypes';
declare class InboxThreadLabelEntity extends ResendEntityBase<InboxThreadLabel> {
    constructor(client: ResendSDK, entopts: any);
    make(this: InboxThreadLabelEntity): InboxThreadLabelEntity;
    create(this: any, reqdata?: InboxThreadLabelCreateData, ctrl?: Control): Promise<InboxThreadLabelEntity>;
    update(this: any, reqdata?: InboxThreadLabelUpdateData, ctrl?: Control): Promise<InboxThreadLabelEntity>;
    remove(this: any, reqmatch?: InboxThreadLabelRemoveMatch, ctrl?: Control): Promise<InboxThreadLabelEntity>;
}
export { InboxThreadLabelEntity };
