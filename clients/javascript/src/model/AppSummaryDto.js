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
import AppChannelSummary from './AppChannelSummary';

/**
 * The AppSummaryDto model module.
 * @module model/AppSummaryDto
 * @version 1.0
 */
class AppSummaryDto {
    /**
     * Constructs a new <code>AppSummaryDto</code>.
     * @alias module:model/AppSummaryDto
     * @param id {String} 
     * @param name {String} 
     * @param storageId {String} 
     * @param channels {Array.<module:model/AppChannelSummary>} 
     */
    constructor(id, name, storageId, channels) { 
        
        AppSummaryDto.initialize(this, id, name, storageId, channels);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, name, storageId, channels) { 
        obj['id'] = id;
        obj['name'] = name;
        obj['storageId'] = storageId;
        obj['channels'] = channels;
    }

    /**
     * Constructs a <code>AppSummaryDto</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AppSummaryDto} obj Optional instance to populate.
     * @return {module:model/AppSummaryDto} The populated <code>AppSummaryDto</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AppSummaryDto();

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
                obj['channels'] = ApiClient.convertToType(data['channels'], [AppChannelSummary]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
AppSummaryDto.prototype['id'] = undefined;

/**
 * @member {String} name
 */
AppSummaryDto.prototype['name'] = undefined;

/**
 * @member {String} storageId
 */
AppSummaryDto.prototype['storageId'] = undefined;

/**
 * @member {Array.<module:model/AppChannelSummary>} channels
 */
AppSummaryDto.prototype['channels'] = undefined;






export default AppSummaryDto;

