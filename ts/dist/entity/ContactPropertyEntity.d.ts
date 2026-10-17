import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { ContactProperty, ContactPropertyLoadMatch, ContactPropertyListMatch, ContactPropertyCreateData, ContactPropertyUpdateData, ContactPropertyRemoveMatch } from '../ResendTypes';
declare class ContactPropertyEntity extends ResendEntityBase<ContactProperty> {
    constructor(client: ResendSDK, entopts: any);
    make(this: ContactPropertyEntity): ContactPropertyEntity;
    load(this: any, reqmatch?: ContactPropertyLoadMatch, ctrl?: Control): Promise<ContactPropertyEntity>;
    list(this: any, reqmatch?: ContactPropertyListMatch, ctrl?: Control): Promise<ContactPropertyEntity[]>;
    create(this: any, reqdata?: ContactPropertyCreateData, ctrl?: Control): Promise<ContactPropertyEntity>;
    update(this: any, reqdata?: ContactPropertyUpdateData, ctrl?: Control): Promise<ContactPropertyEntity>;
    remove(this: any, reqmatch?: ContactPropertyRemoveMatch, ctrl?: Control): Promise<ContactPropertyEntity>;
}
export { ContactPropertyEntity };
