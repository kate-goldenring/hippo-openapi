/**
 * Hippo.Web
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import EnvironmentVariableDto from './EnvironmentVariableDto';

/**
 * The EnvironmentVariablesVm model module.
 * @module model/EnvironmentVariablesVm
 * @version 1.0
 */
class EnvironmentVariablesVm {
    /**
     * Constructs a new <code>EnvironmentVariablesVm</code>.
     * @alias module:model/EnvironmentVariablesVm
     */
    constructor() { 
        
        EnvironmentVariablesVm.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>EnvironmentVariablesVm</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/EnvironmentVariablesVm} obj Optional instance to populate.
     * @return {module:model/EnvironmentVariablesVm} The populated <code>EnvironmentVariablesVm</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new EnvironmentVariablesVm();

            if (data.hasOwnProperty('environmentVariables')) {
                obj['environmentVariables'] = ApiClient.convertToType(data['environmentVariables'], [EnvironmentVariableDto]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/EnvironmentVariableDto>} environmentVariables
 */
EnvironmentVariablesVm.prototype['environmentVariables'] = undefined;






export default EnvironmentVariablesVm;
