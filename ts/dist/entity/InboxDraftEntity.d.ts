import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { InboxDraft, InboxDraftLoadMatch, InboxDraftListMatch, InboxDraftCreateData, InboxDraftUpdateData, InboxDraftRemoveMatch } from '../ResendTypes';
declare class InboxDraftEntity extends ResendEntityBase<InboxDraft> {
    constructor(client: ResendSDK, entopts: any);
    make(this: InboxDraftEntity): InboxDraftEntity;
    load(this: any, reqmatch?: InboxDraftLoadMatch, ctrl?: Control): Promise<InboxDraftEntity>;
    list(this: any, reqmatch?: InboxDraftListMatch, ctrl?: Control): Promise<InboxDraftEntity[]>;
    create(this: any, reqdata?: InboxDraftCreateData, ctrl?: Control): Promise<InboxDraftEntity>;
    update(this: any, reqdata?: InboxDraftUpdateData, ctrl?: Control): Promise<InboxDraftEntity>;
    remove(this: any, reqmatch?: InboxDraftRemoveMatch, ctrl?: Control): Promise<InboxDraftEntity>;
}
export { InboxDraftEntity };
