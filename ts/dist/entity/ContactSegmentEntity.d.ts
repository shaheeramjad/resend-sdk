import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { ContactSegment, ContactSegmentListMatch, ContactSegmentCreateData, ContactSegmentRemoveMatch } from '../ResendTypes';
declare class ContactSegmentEntity extends ResendEntityBase<ContactSegment> {
    constructor(client: ResendSDK, entopts: any);
    make(this: ContactSegmentEntity): ContactSegmentEntity;
    list(this: any, reqmatch?: ContactSegmentListMatch, ctrl?: Control): Promise<ContactSegmentEntity[]>;
    create(this: any, reqdata?: ContactSegmentCreateData, ctrl?: Control): Promise<ContactSegmentEntity>;
    remove(this: any, reqmatch?: ContactSegmentRemoveMatch, ctrl?: Control): Promise<ContactSegmentEntity>;
}
export { ContactSegmentEntity };
