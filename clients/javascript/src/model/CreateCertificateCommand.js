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
 * The CreateCertificateCommand model module.
 * @module model/CreateCertificateCommand
 * @version 1.0
 */
class CreateCertificateCommand {
    /**
     * Constructs a new <code>CreateCertificateCommand</code>.
     * @alias module:model/CreateCertificateCommand
     * @param name {String} 
     * @param publicKey {String} 
     * @param privateKey {String} 
     */
    constructor(name, publicKey, privateKey) { 
        
        CreateCertificateCommand.initialize(this, name, publicKey, privateKey);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, publicKey, privateKey) { 
        obj['name'] = name;
        obj['publicKey'] = publicKey;
        obj['privateKey'] = privateKey;
    }

    /**
     * Constructs a <code>CreateCertificateCommand</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CreateCertificateCommand} obj Optional instance to populate.
     * @return {module:model/CreateCertificateCommand} The populated <code>CreateCertificateCommand</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CreateCertificateCommand();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('publicKey')) {
                obj['publicKey'] = ApiClient.convertToType(data['publicKey'], 'String');
            }
            if (data.hasOwnProperty('privateKey')) {
                obj['privateKey'] = ApiClient.convertToType(data['privateKey'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
CreateCertificateCommand.prototype['name'] = undefined;

/**
 * @member {String} publicKey
 */
CreateCertificateCommand.prototype['publicKey'] = undefined;

/**
 * @member {String} privateKey
 */
CreateCertificateCommand.prototype['privateKey'] = undefined;






export default CreateCertificateCommand;

