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
import Revision from './Revision';

/**
 * The App model module.
 * @module model/App
 * @version 1.0
 */
class App {
    /**
     * Constructs a new <code>App</code>.
     * @alias module:model/App
     */
    constructor() { 
        
        App.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>App</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/App} obj Optional instance to populate.
     * @return {module:model/App} The populated <code>App</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new App();

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
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('storageId')) {
                obj['storageId'] = ApiClient.convertToType(data['storageId'], 'String');
            }
            if (data.hasOwnProperty('channels')) {
                obj['channels'] = ApiClient.convertToType(data['channels'], [Channel]);
            }
            if (data.hasOwnProperty('revisions')) {
                obj['revisions'] = ApiClient.convertToType(data['revisions'], [Revision]);
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
App.prototype['created'] = undefined;

/**
 * @member {String} createdBy
 */
App.prototype['createdBy'] = undefined;

/**
 * @member {Date} lastModified
 */
App.prototype['lastModified'] = undefined;

/**
 * @member {String} lastModifiedBy
 */
App.prototype['lastModifiedBy'] = undefined;

/**
 * @member {String} id
 */
App.prototype['id'] = undefined;

/**
 * @member {String} name
 */
App.prototype['name'] = undefined;

/**
 * @member {String} storageId
 */
App.prototype['storageId'] = undefined;

/**
 * @member {Array.<module:model/Channel>} channels
 */
App.prototype['channels'] = undefined;

/**
 * @member {Array.<module:model/Revision>} revisions
 */
App.prototype['revisions'] = undefined;

/**
 * @member {Array.<module:model/DomainEvent>} domainEvents
 */
App.prototype['domainEvents'] = undefined;






export default App;
