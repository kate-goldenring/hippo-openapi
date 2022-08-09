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


import ApiClient from './ApiClient';
import AppChannelListItem from './model/AppChannelListItem';
import AppItem from './model/AppItem';
import AppItemPage from './model/AppItemPage';
import AppSummaryDto from './model/AppSummaryDto';
import CertificateItem from './model/CertificateItem';
import CertificateItemPage from './model/CertificateItemPage';
import ChannelItem from './model/ChannelItem';
import ChannelItemPage from './model/ChannelItemPage';
import ChannelJobStatusItem from './model/ChannelJobStatusItem';
import ChannelJobStatusItemPage from './model/ChannelJobStatusItemPage';
import ChannelRevisionSelectionStrategy from './model/ChannelRevisionSelectionStrategy';
import ChannelRevisionSelectionStrategyField from './model/ChannelRevisionSelectionStrategyField';
import CreateAccountCommand from './model/CreateAccountCommand';
import CreateAppCommand from './model/CreateAppCommand';
import CreateCertificateCommand from './model/CreateCertificateCommand';
import CreateChannelCommand from './model/CreateChannelCommand';
import CreateTokenCommand from './model/CreateTokenCommand';
import DesiredStatus from './model/DesiredStatus';
import EnvironmentVariableItem from './model/EnvironmentVariableItem';
import GetChannelLogsVm from './model/GetChannelLogsVm';
import GuidNullableField from './model/GuidNullableField';
import JobStatus from './model/JobStatus';
import PatchChannelCommand from './model/PatchChannelCommand';
import RegisterRevisionCommand from './model/RegisterRevisionCommand';
import RevisionComponentDto from './model/RevisionComponentDto';
import RevisionItem from './model/RevisionItem';
import RevisionItemPage from './model/RevisionItemPage';
import StringField from './model/StringField';
import StringPage from './model/StringPage';
import TokenInfo from './model/TokenInfo';
import UpdateAppCommand from './model/UpdateAppCommand';
import UpdateCertificateCommand from './model/UpdateCertificateCommand';
import UpdateChannelCommand from './model/UpdateChannelCommand';
import UpdateDesiredStatusCommand from './model/UpdateDesiredStatusCommand';
import UpdateEnvironmentVariableDto from './model/UpdateEnvironmentVariableDto';
import UpdateEnvironmentVariableDtoListField from './model/UpdateEnvironmentVariableDtoListField';
import AccountsApi from './api/AccountsApi';
import AppsApi from './api/AppsApi';
import AuthTokensApi from './api/AuthTokensApi';
import CertificatesApi from './api/CertificatesApi';
import ChannelStatusesApi from './api/ChannelStatusesApi';
import ChannelsApi from './api/ChannelsApi';
import RevisionsApi from './api/RevisionsApi';
import StoragesApi from './api/StoragesApi';


/**
* JS API client generated by OpenAPI Generator.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var HippoWeb = require('index'); // See note below*.
* var xxxSvc = new HippoWeb.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new HippoWeb.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new HippoWeb.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new HippoWeb.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 1.0
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The AppChannelListItem model constructor.
     * @property {module:model/AppChannelListItem}
     */
    AppChannelListItem,

    /**
     * The AppItem model constructor.
     * @property {module:model/AppItem}
     */
    AppItem,

    /**
     * The AppItemPage model constructor.
     * @property {module:model/AppItemPage}
     */
    AppItemPage,

    /**
     * The AppSummaryDto model constructor.
     * @property {module:model/AppSummaryDto}
     */
    AppSummaryDto,

    /**
     * The CertificateItem model constructor.
     * @property {module:model/CertificateItem}
     */
    CertificateItem,

    /**
     * The CertificateItemPage model constructor.
     * @property {module:model/CertificateItemPage}
     */
    CertificateItemPage,

    /**
     * The ChannelItem model constructor.
     * @property {module:model/ChannelItem}
     */
    ChannelItem,

    /**
     * The ChannelItemPage model constructor.
     * @property {module:model/ChannelItemPage}
     */
    ChannelItemPage,

    /**
     * The ChannelJobStatusItem model constructor.
     * @property {module:model/ChannelJobStatusItem}
     */
    ChannelJobStatusItem,

    /**
     * The ChannelJobStatusItemPage model constructor.
     * @property {module:model/ChannelJobStatusItemPage}
     */
    ChannelJobStatusItemPage,

    /**
     * The ChannelRevisionSelectionStrategy model constructor.
     * @property {module:model/ChannelRevisionSelectionStrategy}
     */
    ChannelRevisionSelectionStrategy,

    /**
     * The ChannelRevisionSelectionStrategyField model constructor.
     * @property {module:model/ChannelRevisionSelectionStrategyField}
     */
    ChannelRevisionSelectionStrategyField,

    /**
     * The CreateAccountCommand model constructor.
     * @property {module:model/CreateAccountCommand}
     */
    CreateAccountCommand,

    /**
     * The CreateAppCommand model constructor.
     * @property {module:model/CreateAppCommand}
     */
    CreateAppCommand,

    /**
     * The CreateCertificateCommand model constructor.
     * @property {module:model/CreateCertificateCommand}
     */
    CreateCertificateCommand,

    /**
     * The CreateChannelCommand model constructor.
     * @property {module:model/CreateChannelCommand}
     */
    CreateChannelCommand,

    /**
     * The CreateTokenCommand model constructor.
     * @property {module:model/CreateTokenCommand}
     */
    CreateTokenCommand,

    /**
     * The DesiredStatus model constructor.
     * @property {module:model/DesiredStatus}
     */
    DesiredStatus,

    /**
     * The EnvironmentVariableItem model constructor.
     * @property {module:model/EnvironmentVariableItem}
     */
    EnvironmentVariableItem,

    /**
     * The GetChannelLogsVm model constructor.
     * @property {module:model/GetChannelLogsVm}
     */
    GetChannelLogsVm,

    /**
     * The GuidNullableField model constructor.
     * @property {module:model/GuidNullableField}
     */
    GuidNullableField,

    /**
     * The JobStatus model constructor.
     * @property {module:model/JobStatus}
     */
    JobStatus,

    /**
     * The PatchChannelCommand model constructor.
     * @property {module:model/PatchChannelCommand}
     */
    PatchChannelCommand,

    /**
     * The RegisterRevisionCommand model constructor.
     * @property {module:model/RegisterRevisionCommand}
     */
    RegisterRevisionCommand,

    /**
     * The RevisionComponentDto model constructor.
     * @property {module:model/RevisionComponentDto}
     */
    RevisionComponentDto,

    /**
     * The RevisionItem model constructor.
     * @property {module:model/RevisionItem}
     */
    RevisionItem,

    /**
     * The RevisionItemPage model constructor.
     * @property {module:model/RevisionItemPage}
     */
    RevisionItemPage,

    /**
     * The StringField model constructor.
     * @property {module:model/StringField}
     */
    StringField,

    /**
     * The StringPage model constructor.
     * @property {module:model/StringPage}
     */
    StringPage,

    /**
     * The TokenInfo model constructor.
     * @property {module:model/TokenInfo}
     */
    TokenInfo,

    /**
     * The UpdateAppCommand model constructor.
     * @property {module:model/UpdateAppCommand}
     */
    UpdateAppCommand,

    /**
     * The UpdateCertificateCommand model constructor.
     * @property {module:model/UpdateCertificateCommand}
     */
    UpdateCertificateCommand,

    /**
     * The UpdateChannelCommand model constructor.
     * @property {module:model/UpdateChannelCommand}
     */
    UpdateChannelCommand,

    /**
     * The UpdateDesiredStatusCommand model constructor.
     * @property {module:model/UpdateDesiredStatusCommand}
     */
    UpdateDesiredStatusCommand,

    /**
     * The UpdateEnvironmentVariableDto model constructor.
     * @property {module:model/UpdateEnvironmentVariableDto}
     */
    UpdateEnvironmentVariableDto,

    /**
     * The UpdateEnvironmentVariableDtoListField model constructor.
     * @property {module:model/UpdateEnvironmentVariableDtoListField}
     */
    UpdateEnvironmentVariableDtoListField,

    /**
    * The AccountsApi service constructor.
    * @property {module:api/AccountsApi}
    */
    AccountsApi,

    /**
    * The AppsApi service constructor.
    * @property {module:api/AppsApi}
    */
    AppsApi,

    /**
    * The AuthTokensApi service constructor.
    * @property {module:api/AuthTokensApi}
    */
    AuthTokensApi,

    /**
    * The CertificatesApi service constructor.
    * @property {module:api/CertificatesApi}
    */
    CertificatesApi,

    /**
    * The ChannelStatusesApi service constructor.
    * @property {module:api/ChannelStatusesApi}
    */
    ChannelStatusesApi,

    /**
    * The ChannelsApi service constructor.
    * @property {module:api/ChannelsApi}
    */
    ChannelsApi,

    /**
    * The RevisionsApi service constructor.
    * @property {module:api/RevisionsApi}
    */
    RevisionsApi,

    /**
    * The StoragesApi service constructor.
    * @property {module:api/StoragesApi}
    */
    StoragesApi
};
