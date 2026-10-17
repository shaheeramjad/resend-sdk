import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { SegmentContact, SegmentContactListMatch } from '../ResendTypes';
declare class SegmentContactEntity extends ResendEntityBase<SegmentContact> {
    constructor(client: ResendSDK, entopts: any);
    make(this: SegmentContactEntity): SegmentContactEntity;
    list(this: any, reqmatch?: SegmentContactListMatch, ctrl?: Control): Promise<SegmentContactEntity[]>;
}
export { SegmentContactEntity };
