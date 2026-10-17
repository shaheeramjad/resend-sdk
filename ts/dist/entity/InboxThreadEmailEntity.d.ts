import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { InboxThreadEmail, InboxThreadEmailLoadMatch, InboxThreadEmailListMatch, InboxThreadEmailCreateData } from '../ResendTypes';
declare class InboxThreadEmailEntity extends ResendEntityBase<InboxThreadEmail> {
    constructor(client: ResendSDK, entopts: any);
    make(this: InboxThreadEmailEntity): InboxThreadEmailEntity;
    load(this: any, reqmatch?: InboxThreadEmailLoadMatch, ctrl?: Control): Promise<InboxThreadEmailEntity>;
    list(this: any, reqmatch?: InboxThreadEmailListMatch, ctrl?: Control): Promise<InboxThreadEmailEntity[]>;
    create(this: any, reqdata?: InboxThreadEmailCreateData, ctrl?: Control): Promise<InboxThreadEmailEntity>;
}
export { InboxThreadEmailEntity };
