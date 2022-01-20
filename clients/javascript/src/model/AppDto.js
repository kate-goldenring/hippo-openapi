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
import ChannelDto from './ChannelDto';
import RevisionDto from './RevisionDto';

/**
 * The AppDto model module.
 * @module model/AppDto
 * @version 1.0
 */
class AppDto {
    /**
     * Constructs a new <code>AppDto</code>.
     * @alias module:model/AppDto
     */
    constructor() { 
        
        AppDto.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AppDto</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AppDto} obj Optional instance to populate.
     * @return {module:model/AppDto} The populated <code>AppDto</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AppDto();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('storageId')) {
                obj['storageId'] = ApiClient.convertToType(data['storageId'], 'String');
            }
            if (data.hasOwnProperty('channels')) {
                obj['channels'] = ApiClient.convertToType(data['channels'], [ChannelDto]);
            }
            if (data.hasOwnProperty('revisions')) {
                obj['revisions'] = ApiClient.convertToType(data['revisions'], [RevisionDto]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
AppDto.prototype['id'] = undefined;

/**
 * @member {String} name
 */
AppDto.prototype['name'] = undefined;

/**
 * @member {String} storageId
 */
AppDto.prototype['storageId'] = undefined;

/**
 * @member {Array.<module:model/ChannelDto>} channels
 */
AppDto.prototype['channels'] = undefined;

/**
 * @member {Array.<module:model/RevisionDto>} revisions
 */
AppDto.prototype['revisions'] = undefined;






export default AppDto;

