import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { ContactTopic, ContactTopicListMatch, ContactTopicUpdateData } from '../ResendTypes';
declare class ContactTopicEntity extends ResendEntityBase<ContactTopic> {
    constructor(client: ResendSDK, entopts: any);
    make(this: ContactTopicEntity): ContactTopicEntity;
    list(this: any, reqmatch?: ContactTopicListMatch, ctrl?: Control): Promise<ContactTopicEntity[]>;
    update(this: any, reqdata?: ContactTopicUpdateData, ctrl?: Control): Promise<ContactTopicEntity>;
}
export { ContactTopicEntity };
