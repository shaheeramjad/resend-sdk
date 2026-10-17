import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { InboxLabel, InboxLabelListMatch, InboxLabelCreateData, InboxLabelUpdateData, InboxLabelRemoveMatch } from '../ResendTypes';
declare class InboxLabelEntity extends ResendEntityBase<InboxLabel> {
    constructor(client: ResendSDK, entopts: any);
    make(this: InboxLabelEntity): InboxLabelEntity;
    list(this: any, reqmatch?: InboxLabelListMatch, ctrl?: Control): Promise<InboxLabelEntity[]>;
    create(this: any, reqdata?: InboxLabelCreateData, ctrl?: Control): Promise<InboxLabelEntity>;
    update(this: any, reqdata?: InboxLabelUpdateData, ctrl?: Control): Promise<InboxLabelEntity>;
    remove(this: any, reqmatch?: InboxLabelRemoveMatch, ctrl?: Control): Promise<InboxLabelEntity>;
}
export { InboxLabelEntity };
