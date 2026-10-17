import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { SuppressionBatchRemove, SuppressionBatchRemoveCreateData } from '../ResendTypes';
declare class SuppressionBatchRemoveEntity extends ResendEntityBase<SuppressionBatchRemove> {
    constructor(client: ResendSDK, entopts: any);
    make(this: SuppressionBatchRemoveEntity): SuppressionBatchRemoveEntity;
    create(this: any, reqdata?: SuppressionBatchRemoveCreateData, ctrl?: Control): Promise<SuppressionBatchRemoveEntity>;
}
export { SuppressionBatchRemoveEntity };
