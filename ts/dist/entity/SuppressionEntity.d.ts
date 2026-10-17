import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { Suppression, SuppressionLoadMatch, SuppressionListMatch, SuppressionCreateData, SuppressionRemoveMatch } from '../ResendTypes';
declare class SuppressionEntity extends ResendEntityBase<Suppression> {
    constructor(client: ResendSDK, entopts: any);
    make(this: SuppressionEntity): SuppressionEntity;
    load(this: any, reqmatch?: SuppressionLoadMatch, ctrl?: Control): Promise<SuppressionEntity>;
    list(this: any, reqmatch?: SuppressionListMatch, ctrl?: Control): Promise<SuppressionEntity[]>;
    create(this: any, reqdata?: SuppressionCreateData, ctrl?: Control): Promise<SuppressionEntity>;
    remove(this: any, reqmatch?: SuppressionRemoveMatch, ctrl?: Control): Promise<SuppressionEntity>;
}
export { SuppressionEntity };
