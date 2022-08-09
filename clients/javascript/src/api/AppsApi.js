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


import ApiClient from "../ApiClient";
import AppItem from '../model/AppItem';
import AppItemPage from '../model/AppItemPage';
import CreateAppCommand from '../model/CreateAppCommand';
import UpdateAppCommand from '../model/UpdateAppCommand';

/**
* Apps service.
* @module api/AppsApi
* @version 1.0
*/
export default class AppsApi {

    /**
    * Constructs a new AppsApi. 
    * @alias module:api/AppsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the apiAppsGet operation.
     * @callback module:api/AppsApi~apiAppsGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AppItemPage} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {Object} opts Optional parameters
     * @param {String} opts.searchText  (default to '')
     * @param {Number} opts.pageIndex  (default to 0)
     * @param {Number} opts.pageSize  (default to 50)
     * @param {String} opts.sortBy  (default to 'Name')
     * @param {Boolean} opts.isSortedAscending  (default to true)
     * @param {module:api/AppsApi~apiAppsGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AppItemPage}
     */
    apiAppsGet(opts, callback) {
      opts = opts || {};
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
        'searchText': opts['searchText'],
        'pageIndex': opts['pageIndex'],
        'pageSize': opts['pageSize'],
        'sortBy': opts['sortBy'],
        'IsSortedAscending': opts['isSortedAscending']
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['Bearer'];
      let contentTypes = [];
      let accepts = ['text/plain', 'application/json', 'text/json'];
      let returnType = AppItemPage;
      return this.apiClient.callApi(
        '/api/apps', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the apiAppsIdDelete operation.
     * @callback module:api/AppsApi~apiAppsIdDeleteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} id 
     * @param {module:api/AppsApi~apiAppsIdDeleteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    apiAppsIdDelete(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling apiAppsIdDelete");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['Bearer'];
      let contentTypes = [];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/apps/{id}', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the apiAppsIdGet operation.
     * @callback module:api/AppsApi~apiAppsIdGetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/AppItem} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} id 
     * @param {module:api/AppsApi~apiAppsIdGetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/AppItem}
     */
    apiAppsIdGet(id, callback) {
      let postBody = null;
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling apiAppsIdGet");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['Bearer'];
      let contentTypes = [];
      let accepts = ['text/plain', 'application/json', 'text/json'];
      let returnType = AppItem;
      return this.apiClient.callApi(
        '/api/apps/{id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the apiAppsIdPut operation.
     * @callback module:api/AppsApi~apiAppsIdPutCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {String} id 
     * @param {Object} opts Optional parameters
     * @param {module:model/UpdateAppCommand} opts.updateAppCommand 
     * @param {module:api/AppsApi~apiAppsIdPutCallback} callback The callback function, accepting three arguments: error, data, response
     */
    apiAppsIdPut(id, opts, callback) {
      opts = opts || {};
      let postBody = opts['updateAppCommand'];
      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling apiAppsIdPut");
      }

      let pathParams = {
        'id': id
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['Bearer'];
      let contentTypes = ['application/json-patch+json', 'application/json', 'text/json', 'application/*+json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/apps/{id}', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the apiAppsPost operation.
     * @callback module:api/AppsApi~apiAppsPostCallback
     * @param {String} error Error message, if any.
     * @param {String} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * @param {Object} opts Optional parameters
     * @param {module:model/CreateAppCommand} opts.createAppCommand 
     * @param {module:api/AppsApi~apiAppsPostCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link String}
     */
    apiAppsPost(opts, callback) {
      opts = opts || {};
      let postBody = opts['createAppCommand'];

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['Bearer'];
      let contentTypes = ['application/json-patch+json', 'application/json', 'text/json', 'application/*+json'];
      let accepts = ['text/plain', 'application/json', 'text/json'];
      let returnType = 'String';
      return this.apiClient.callApi(
        '/api/apps', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}