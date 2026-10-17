import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { AutomationRun, AutomationRunLoadMatch, AutomationRunListMatch } from '../ResendTypes';
declare class AutomationRunEntity extends ResendEntityBase<AutomationRun> {
    constructor(client: ResendSDK, entopts: any);
    make(this: AutomationRunEntity): AutomationRunEntity;
    load(this: any, reqmatch?: AutomationRunLoadMatch, ctrl?: Control): Promise<AutomationRunEntity>;
    list(this: any, reqmatch?: AutomationRunListMatch, ctrl?: Control): Promise<AutomationRunEntity[]>;
}
export { AutomationRunEntity };
