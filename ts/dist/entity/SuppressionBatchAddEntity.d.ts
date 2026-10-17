import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { SuppressionBatchAdd, SuppressionBatchAddCreateData } from '../ResendTypes';
declare class SuppressionBatchAddEntity extends ResendEntityBase<SuppressionBatchAdd> {
    constructor(client: ResendSDK, entopts: any);
    make(this: SuppressionBatchAddEntity): SuppressionBatchAddEntity;
    create(this: any, reqdata?: SuppressionBatchAddCreateData, ctrl?: Control): Promise<SuppressionBatchAddEntity>;
}
export { SuppressionBatchAddEntity };
