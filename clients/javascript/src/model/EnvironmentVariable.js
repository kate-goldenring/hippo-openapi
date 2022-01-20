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
import Channel from './Channel';
import DomainEvent from './DomainEvent';

/**
 * The EnvironmentVariable model module.
 * @module model/EnvironmentVariable
 * @version 1.0
 */
class EnvironmentVariable {
    /**
     * Constructs a new <code>EnvironmentVariable</code>.
     * @alias module:model/EnvironmentVariable
     */
    constructor() { 
        
        EnvironmentVariable.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>EnvironmentVariable</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/EnvironmentVariable} obj Optional instance to populate.
     * @return {module:model/EnvironmentVariable} The populated <code>EnvironmentVariable</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new EnvironmentVariable();

            if (data.hasOwnProperty('created')) {
                obj['created'] = ApiClient.convertToType(data['created'], 'Date');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = ApiClient.convertToType(data['createdBy'], 'String');
            }
            if (data.hasOwnProperty('lastModified')) {
                obj['lastModified'] = ApiClient.convertToType(data['lastModified'], 'Date');
            }
            if (data.hasOwnProperty('lastModifiedBy')) {
                obj['lastModifiedBy'] = ApiClient.convertToType(data['lastModifiedBy'], 'String');
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('key')) {
                obj['key'] = ApiClient.convertToType(data['key'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], 'String');
            }
            if (data.hasOwnProperty('channelId')) {
                obj['channelId'] = ApiClient.convertToType(data['channelId'], 'String');
            }
            if (data.hasOwnProperty('channel')) {
                obj['channel'] = Channel.constructFromObject(data['channel']);
            }
            if (data.hasOwnProperty('domainEvents')) {
                obj['domainEvents'] = ApiClient.convertToType(data['domainEvents'], [DomainEvent]);
            }
        }
        return obj;
    }


}

/**
 * @member {Date} created
 */
EnvironmentVariable.prototype['created'] = undefined;

/**
 * @member {String} createdBy
 */
EnvironmentVariable.prototype['createdBy'] = undefined;

/**
 * @member {Date} lastModified
 */
EnvironmentVariable.prototype['lastModified'] = undefined;

/**
 * @member {String} lastModifiedBy
 */
EnvironmentVariable.prototype['lastModifiedBy'] = undefined;

/**
 * @member {String} id
 */
EnvironmentVariable.prototype['id'] = undefined;

/**
 * @member {String} key
 */
EnvironmentVariable.prototype['key'] = undefined;

/**
 * @member {String} value
 */
EnvironmentVariable.prototype['value'] = undefined;

/**
 * @member {String} channelId
 */
EnvironmentVariable.prototype['channelId'] = undefined;

/**
 * @member {module:model/Channel} channel
 */
EnvironmentVariable.prototype['channel'] = undefined;

/**
 * @member {Array.<module:model/DomainEvent>} domainEvents
 */
EnvironmentVariable.prototype['domainEvents'] = undefined;






export default EnvironmentVariable;

