import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { Attachment, AttachmentLoadMatch, AttachmentListMatch } from '../ResendTypes';
declare class AttachmentEntity extends ResendEntityBase<Attachment> {
    constructor(client: ResendSDK, entopts: any);
    make(this: AttachmentEntity): AttachmentEntity;
    load(this: any, reqmatch?: AttachmentLoadMatch, ctrl?: Control): Promise<AttachmentEntity>;
    list(this: any, reqmatch?: AttachmentListMatch, ctrl?: Control): Promise<AttachmentEntity[]>;
}
export { AttachmentEntity };
