# hippo_web

HippoWeb - JavaScript client for hippo_web
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
This SDK is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: 1.0
- Package version: 1.0
- Build package: org.openapitools.codegen.languages.JavascriptClientCodegen

## Installation

### For [Node.js](https://nodejs.org/)

#### npm

To publish the library as a [npm](https://www.npmjs.com/), please follow the procedure in ["Publishing npm packages"](https://docs.npmjs.com/getting-started/publishing-npm-packages).

Then install it via:

```shell
npm install hippo_web --save
```

Finally, you need to build the module:

```shell
npm run build
```

##### Local development

To use the library locally without publishing to a remote npm registry, first install the dependencies by changing into the directory containing `package.json` (and this README). Let's call this `JAVASCRIPT_CLIENT_DIR`. Then run:

```shell
npm install
```

Next, [link](https://docs.npmjs.com/cli/link) it globally in npm with the following, also from `JAVASCRIPT_CLIENT_DIR`:

```shell
npm link
```

To use the link you just defined in your project, switch to the directory you want to use your hippo_web from, and run:

```shell
npm link /path/to/<JAVASCRIPT_CLIENT_DIR>
```

Finally, you need to build the module:

```shell
npm run build
```

#### git

If the library is hosted at a git repository, e.g.https://github.com/deislabs/deislabs/hippo-openapi
then install it via:

```shell
    npm install deislabs/deislabs/hippo-openapi --save
```

### For browser

The library also works in the browser environment via npm and [browserify](http://browserify.org/). After following
the above steps with Node.js and installing browserify with `npm install -g browserify`,
perform the following (assuming *main.js* is your entry file):

```shell
browserify main.js > bundle.js
```

Then include *bundle.js* in the HTML pages.

### Webpack Configuration

Using Webpack you may encounter the following error: "Module not found: Error:
Cannot resolve module", most certainly you should disable AMD loader. Add/merge
the following section to your webpack config:

```javascript
module: {
  rules: [
    {
      parser: {
        amd: false
      }
    }
  ]
}
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following JS code:

```javascript
var HippoWeb = require('hippo_web');

var defaultClient = HippoWeb.ApiClient.instance;
// Configure API key authorization: Bearer
var Bearer = defaultClient.authentications['Bearer'];
Bearer.apiKey = "YOUR API KEY"
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//Bearer.apiKeyPrefix['Authorization'] = "Token"

var api = new HippoWeb.AccountApi()
var opts = {
  'createTokenCommand': new HippoWeb.CreateTokenCommand() // {CreateTokenCommand} 
};
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
api.apiAccountCreatetokenPost(opts, callback);

```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HippoWeb.AccountApi* | [**apiAccountCreatetokenPost**](docs/AccountApi.md#apiAccountCreatetokenPost) | **POST** /api/account/createtoken | 
*HippoWeb.AccountApi* | [**apiAccountPost**](docs/AccountApi.md#apiAccountPost) | **POST** /api/account | 
*HippoWeb.AppApi* | [**apiAppExportGet**](docs/AppApi.md#apiAppExportGet) | **GET** /api/app/export | 
*HippoWeb.AppApi* | [**apiAppGet**](docs/AppApi.md#apiAppGet) | **GET** /api/app | 
*HippoWeb.AppApi* | [**apiAppIdDelete**](docs/AppApi.md#apiAppIdDelete) | **DELETE** /api/app/{id} | 
*HippoWeb.AppApi* | [**apiAppPost**](docs/AppApi.md#apiAppPost) | **POST** /api/app | 
*HippoWeb.CertificateApi* | [**apiCertificateGet**](docs/CertificateApi.md#apiCertificateGet) | **GET** /api/certificate | 
*HippoWeb.CertificateApi* | [**apiCertificateIdDelete**](docs/CertificateApi.md#apiCertificateIdDelete) | **DELETE** /api/certificate/{id} | 
*HippoWeb.CertificateApi* | [**apiCertificateIdGet**](docs/CertificateApi.md#apiCertificateIdGet) | **GET** /api/certificate/{id} | 
*HippoWeb.CertificateApi* | [**apiCertificatePost**](docs/CertificateApi.md#apiCertificatePost) | **POST** /api/certificate | 
*HippoWeb.ChannelApi* | [**apiChannelGet**](docs/ChannelApi.md#apiChannelGet) | **GET** /api/channel | 
*HippoWeb.ChannelApi* | [**apiChannelIdDelete**](docs/ChannelApi.md#apiChannelIdDelete) | **DELETE** /api/channel/{id} | 
*HippoWeb.ChannelApi* | [**apiChannelIdGet**](docs/ChannelApi.md#apiChannelIdGet) | **GET** /api/channel/{id} | 
*HippoWeb.ChannelApi* | [**apiChannelPost**](docs/ChannelApi.md#apiChannelPost) | **POST** /api/channel | 
*HippoWeb.EnvironmentVariableApi* | [**apiEnvironmentvariableGet**](docs/EnvironmentVariableApi.md#apiEnvironmentvariableGet) | **GET** /api/environmentvariable | 
*HippoWeb.EnvironmentVariableApi* | [**apiEnvironmentvariableIdDelete**](docs/EnvironmentVariableApi.md#apiEnvironmentvariableIdDelete) | **DELETE** /api/environmentvariable/{id} | 
*HippoWeb.EnvironmentVariableApi* | [**apiEnvironmentvariableIdGet**](docs/EnvironmentVariableApi.md#apiEnvironmentvariableIdGet) | **GET** /api/environmentvariable/{id} | 
*HippoWeb.EnvironmentVariableApi* | [**apiEnvironmentvariablePost**](docs/EnvironmentVariableApi.md#apiEnvironmentvariablePost) | **POST** /api/environmentvariable | 
*HippoWeb.RevisionApi* | [**apiRevisionGet**](docs/RevisionApi.md#apiRevisionGet) | **GET** /api/revision | 
*HippoWeb.RevisionApi* | [**apiRevisionIdDelete**](docs/RevisionApi.md#apiRevisionIdDelete) | **DELETE** /api/revision/{id} | 
*HippoWeb.RevisionApi* | [**apiRevisionIdGet**](docs/RevisionApi.md#apiRevisionIdGet) | **GET** /api/revision/{id} | 
*HippoWeb.RevisionApi* | [**apiRevisionPost**](docs/RevisionApi.md#apiRevisionPost) | **POST** /api/revision | 


## Documentation for Models

 - [HippoWeb.App](docs/App.md)
 - [HippoWeb.AppDto](docs/AppDto.md)
 - [HippoWeb.AppsVm](docs/AppsVm.md)
 - [HippoWeb.Certificate](docs/Certificate.md)
 - [HippoWeb.CertificateDto](docs/CertificateDto.md)
 - [HippoWeb.CertificatesVm](docs/CertificatesVm.md)
 - [HippoWeb.Channel](docs/Channel.md)
 - [HippoWeb.ChannelDto](docs/ChannelDto.md)
 - [HippoWeb.ChannelRevisionSelectionStrategy](docs/ChannelRevisionSelectionStrategy.md)
 - [HippoWeb.ChannelsVm](docs/ChannelsVm.md)
 - [HippoWeb.CreateAccountCommand](docs/CreateAccountCommand.md)
 - [HippoWeb.CreateAppCommand](docs/CreateAppCommand.md)
 - [HippoWeb.CreateCertificateCommand](docs/CreateCertificateCommand.md)
 - [HippoWeb.CreateChannelCommand](docs/CreateChannelCommand.md)
 - [HippoWeb.CreateEnvironmentVariableCommand](docs/CreateEnvironmentVariableCommand.md)
 - [HippoWeb.CreateTokenCommand](docs/CreateTokenCommand.md)
 - [HippoWeb.DomainEvent](docs/DomainEvent.md)
 - [HippoWeb.EnvironmentVariable](docs/EnvironmentVariable.md)
 - [HippoWeb.EnvironmentVariableDto](docs/EnvironmentVariableDto.md)
 - [HippoWeb.EnvironmentVariablesVm](docs/EnvironmentVariablesVm.md)
 - [HippoWeb.RegisterRevisionCommand](docs/RegisterRevisionCommand.md)
 - [HippoWeb.Revision](docs/Revision.md)
 - [HippoWeb.RevisionDto](docs/RevisionDto.md)
 - [HippoWeb.RevisionsVm](docs/RevisionsVm.md)
 - [HippoWeb.TokenInfo](docs/TokenInfo.md)


## Documentation for Authorization



### Bearer


- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header
