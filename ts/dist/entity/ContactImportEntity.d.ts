import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { ContactImport, ContactImportLoadMatch, ContactImportListMatch, ContactImportCreateData } from '../ResendTypes';
declare class ContactImportEntity extends ResendEntityBase<ContactImport> {
    constructor(client: ResendSDK, entopts: any);
    make(this: ContactImportEntity): ContactImportEntity;
    load(this: any, reqmatch?: ContactImportLoadMatch, ctrl?: Control): Promise<ContactImportEntity>;
    list(this: any, reqmatch?: ContactImportListMatch, ctrl?: Control): Promise<ContactImportEntity[]>;
    create(this: any, reqdata?: ContactImportCreateData, ctrl?: Control): Promise<ContactImportEntity>;
}
export { ContactImportEntity };
