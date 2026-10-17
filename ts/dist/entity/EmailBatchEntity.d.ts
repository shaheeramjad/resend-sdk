import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { EmailBatch, EmailBatchCreateData } from '../ResendTypes';
declare class EmailBatchEntity extends ResendEntityBase<EmailBatch> {
    constructor(client: ResendSDK, entopts: any);
    make(this: EmailBatchEntity): EmailBatchEntity;
    create(this: any, reqdata?: EmailBatchCreateData, ctrl?: Control): Promise<EmailBatchEntity>;
}
export { EmailBatchEntity };
