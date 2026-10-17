import { ResendEntityBase } from '../ResendEntityBase';
import type { ResendSDK } from '../ResendSDK';
import type { Control } from '../types';
import type { OauthGrant, OauthGrantListMatch, OauthGrantRemoveMatch } from '../ResendTypes';
declare class OauthGrantEntity extends ResendEntityBase<OauthGrant> {
    constructor(client: ResendSDK, entopts: any);
    make(this: OauthGrantEntity): OauthGrantEntity;
    list(this: any, reqmatch?: OauthGrantListMatch, ctrl?: Control): Promise<OauthGrantEntity[]>;
    remove(this: any, reqmatch?: OauthGrantRemoveMatch, ctrl?: Control): Promise<OauthGrantEntity>;
}
export { OauthGrantEntity };
