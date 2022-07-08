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
import AppSummaryDto from './AppSummaryDto';
import CertificateItem from './CertificateItem';
import ChannelRevisionSelectionStrategy from './ChannelRevisionSelectionStrategy';
import EnvironmentVariableItem from './EnvironmentVariableItem';
import RevisionItem from './RevisionItem';

/**
 * The ChannelItem model module.
 * @module model/ChannelItem
 * @version 1.0
 */
class ChannelItem {
    /**
     * Constructs a new <code>ChannelItem</code>.
     * @alias module:model/ChannelItem
     * @param id {String} 
     * @param appId {String} 
     * @param name {String} 
     * @param domain {String} 
     * @param revisionSelectionStrategy {module:model/ChannelRevisionSelectionStrategy} 
     * @param environmentVariables {Array.<module:model/EnvironmentVariableItem>} 
     */
    constructor(id, appId, name, domain, revisionSelectionStrategy, environmentVariables) { 
        
        ChannelItem.initialize(this, id, appId, name, domain, revisionSelectionStrategy, environmentVariables);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, appId, name, domain, revisionSelectionStrategy, environmentVariables) { 
        obj['id'] = id;
        obj['appId'] = appId;
        obj['name'] = name;
        obj['domain'] = domain;
        obj['revisionSelectionStrategy'] = revisionSelectionStrategy;
        obj['environmentVariables'] = environmentVariables;
    }

    /**
     * Constructs a <code>ChannelItem</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ChannelItem} obj Optional instance to populate.
     * @return {module:model/ChannelItem} The populated <code>ChannelItem</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ChannelItem();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('appId')) {
                obj['appId'] = ApiClient.convertToType(data['appId'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('domain')) {
                obj['domain'] = ApiClient.convertToType(data['domain'], 'String');
            }
            if (data.hasOwnProperty('revisionSelectionStrategy')) {
                obj['revisionSelectionStrategy'] = ChannelRevisionSelectionStrategy.constructFromObject(data['revisionSelectionStrategy']);
            }
            if (data.hasOwnProperty('activeRevision')) {
                obj['activeRevision'] = RevisionItem.constructFromObject(data['activeRevision']);
            }
            if (data.hasOwnProperty('lastPublishAt')) {
                obj['lastPublishAt'] = ApiClient.convertToType(data['lastPublishAt'], 'Date');
            }
            if (data.hasOwnProperty('rangeRule')) {
                obj['rangeRule'] = ApiClient.convertToType(data['rangeRule'], 'String');
            }
            if (data.hasOwnProperty('certificate')) {
                obj['certificate'] = CertificateItem.constructFromObject(data['certificate']);
            }
            if (data.hasOwnProperty('environmentVariables')) {
                obj['environmentVariables'] = ApiClient.convertToType(data['environmentVariables'], [EnvironmentVariableItem]);
            }
            if (data.hasOwnProperty('appSummary')) {
                obj['appSummary'] = AppSummaryDto.constructFromObject(data['appSummary']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
ChannelItem.prototype['id'] = undefined;

/**
 * @member {String} appId
 */
ChannelItem.prototype['appId'] = undefined;

/**
 * @member {String} name
 */
ChannelItem.prototype['name'] = undefined;

/**
 * @member {String} domain
 */
ChannelItem.prototype['domain'] = undefined;

/**
 * @member {module:model/ChannelRevisionSelectionStrategy} revisionSelectionStrategy
 */
ChannelItem.prototype['revisionSelectionStrategy'] = undefined;

/**
 * @member {module:model/RevisionItem} activeRevision
 */
ChannelItem.prototype['activeRevision'] = undefined;

/**
 * @member {Date} lastPublishAt
 */
ChannelItem.prototype['lastPublishAt'] = undefined;

/**
 * @member {String} rangeRule
 */
ChannelItem.prototype['rangeRule'] = undefined;

/**
 * @member {module:model/CertificateItem} certificate
 */
ChannelItem.prototype['certificate'] = undefined;

/**
 * @member {Array.<module:model/EnvironmentVariableItem>} environmentVariables
 */
ChannelItem.prototype['environmentVariables'] = undefined;

/**
 * @member {module:model/AppSummaryDto} appSummary
 */
ChannelItem.prototype['appSummary'] = undefined;






export default ChannelItem;

