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

/**
 * The EnvironmentVariableDto model module.
 * @module model/EnvironmentVariableDto
 * @version 1.0
 */
class EnvironmentVariableDto {
    /**
     * Constructs a new <code>EnvironmentVariableDto</code>.
     * @alias module:model/EnvironmentVariableDto
     */
    constructor() { 
        
        EnvironmentVariableDto.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>EnvironmentVariableDto</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/EnvironmentVariableDto} obj Optional instance to populate.
     * @return {module:model/EnvironmentVariableDto} The populated <code>EnvironmentVariableDto</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new EnvironmentVariableDto();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('channelId')) {
                obj['channelId'] = ApiClient.convertToType(data['channelId'], 'String');
            }
            if (data.hasOwnProperty('key')) {
                obj['key'] = ApiClient.convertToType(data['key'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
EnvironmentVariableDto.prototype['id'] = undefined;

/**
 * @member {String} channelId
 */
EnvironmentVariableDto.prototype['channelId'] = undefined;

/**
 * @member {String} key
 */
EnvironmentVariableDto.prototype['key'] = undefined;

/**
 * @member {String} value
 */
EnvironmentVariableDto.prototype['value'] = undefined;






export default EnvironmentVariableDto;
