import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { ReceivedAttachment, ReceivedAttachmentLoadMatch, ReceivedAttachmentListMatch } from '../ResendTypes';
declare class ReceivedAttachmentEntity extends ResendEntityBase<ReceivedAttachment> {
    constructor(client: ResendSDK, entopts: any);
    make(this: ReceivedAttachmentEntity): ReceivedAttachmentEntity;
    load(this: any, reqmatch?: ReceivedAttachmentLoadMatch, ctrl?: Control): Promise<ReceivedAttachmentEntity>;
    list(this: any, reqmatch?: ReceivedAttachmentListMatch, ctrl?: Control): Promise<ReceivedAttachmentEntity[]>;
}
export { ReceivedAttachmentEntity };
