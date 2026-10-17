import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { InboxAgent, InboxAgentLoadMatch, InboxAgentUpdateData } from '../ResendTypes';
declare class InboxAgentEntity extends ResendEntityBase<InboxAgent> {
    constructor(client: ResendSDK, entopts: any);
    make(this: InboxAgentEntity): InboxAgentEntity;
    load(this: any, reqmatch?: InboxAgentLoadMatch, ctrl?: Control): Promise<InboxAgentEntity>;
    update(this: any, reqdata?: InboxAgentUpdateData, ctrl?: Control): Promise<InboxAgentEntity>;
}
export { InboxAgentEntity };
