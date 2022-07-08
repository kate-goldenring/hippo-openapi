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
import ChannelItem from './ChannelItem';

/**
 * The CertificateItem model module.
 * @module model/CertificateItem
 * @version 1.0
 */
class CertificateItem {
    /**
     * Constructs a new <code>CertificateItem</code>.
     * @alias module:model/CertificateItem
     * @param id {String} 
     * @param name {String} 
     * @param publicKey {String} 
     * @param privateKey {String} 
     * @param channels {Array.<module:model/ChannelItem>} 
     */
    constructor(id, name, publicKey, privateKey, channels) { 
        
        CertificateItem.initialize(this, id, name, publicKey, privateKey, channels);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, name, publicKey, privateKey, channels) { 
        obj['id'] = id;
        obj['name'] = name;
        obj['publicKey'] = publicKey;
        obj['privateKey'] = privateKey;
        obj['channels'] = channels;
    }

    /**
     * Constructs a <code>CertificateItem</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CertificateItem} obj Optional instance to populate.
     * @return {module:model/CertificateItem} The populated <code>CertificateItem</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CertificateItem();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('publicKey')) {
                obj['publicKey'] = ApiClient.convertToType(data['publicKey'], 'String');
            }
            if (data.hasOwnProperty('privateKey')) {
                obj['privateKey'] = ApiClient.convertToType(data['privateKey'], 'String');
            }
            if (data.hasOwnProperty('channels')) {
                obj['channels'] = ApiClient.convertToType(data['channels'], [ChannelItem]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
CertificateItem.prototype['id'] = undefined;

/**
 * @member {String} name
 */
CertificateItem.prototype['name'] = undefined;

/**
 * @member {String} publicKey
 */
CertificateItem.prototype['publicKey'] = undefined;

/**
 * @member {String} privateKey
 */
CertificateItem.prototype['privateKey'] = undefined;

/**
 * @member {Array.<module:model/ChannelItem>} channels
 */
CertificateItem.prototype['channels'] = undefined;






export default CertificateItem;

