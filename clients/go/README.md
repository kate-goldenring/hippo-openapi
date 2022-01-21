# Go API client for hippo-openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 0.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./hippo-openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**ApiAccountCreatetokenPost**](docs/AccountApi.md#apiaccountcreatetokenpost) | **Post** /api/account/createtoken | 
*AccountApi* | [**ApiAccountPost**](docs/AccountApi.md#apiaccountpost) | **Post** /api/account | 
*AppApi* | [**ApiAppExportGet**](docs/AppApi.md#apiappexportget) | **Get** /api/app/export | 
*AppApi* | [**ApiAppGet**](docs/AppApi.md#apiappget) | **Get** /api/app | 
*AppApi* | [**ApiAppIdDelete**](docs/AppApi.md#apiappiddelete) | **Delete** /api/app/{id} | 
*AppApi* | [**ApiAppPost**](docs/AppApi.md#apiapppost) | **Post** /api/app | 
*CertificateApi* | [**ApiCertificateGet**](docs/CertificateApi.md#apicertificateget) | **Get** /api/certificate | 
*CertificateApi* | [**ApiCertificateIdDelete**](docs/CertificateApi.md#apicertificateiddelete) | **Delete** /api/certificate/{id} | 
*CertificateApi* | [**ApiCertificateIdGet**](docs/CertificateApi.md#apicertificateidget) | **Get** /api/certificate/{id} | 
*CertificateApi* | [**ApiCertificatePost**](docs/CertificateApi.md#apicertificatepost) | **Post** /api/certificate | 
*ChannelApi* | [**ApiChannelGet**](docs/ChannelApi.md#apichannelget) | **Get** /api/channel | 
*ChannelApi* | [**ApiChannelIdDelete**](docs/ChannelApi.md#apichanneliddelete) | **Delete** /api/channel/{id} | 
*ChannelApi* | [**ApiChannelIdGet**](docs/ChannelApi.md#apichannelidget) | **Get** /api/channel/{id} | 
*ChannelApi* | [**ApiChannelPost**](docs/ChannelApi.md#apichannelpost) | **Post** /api/channel | 
*EnvironmentVariableApi* | [**ApiEnvironmentvariableGet**](docs/EnvironmentVariableApi.md#apienvironmentvariableget) | **Get** /api/environmentvariable | 
*EnvironmentVariableApi* | [**ApiEnvironmentvariableIdDelete**](docs/EnvironmentVariableApi.md#apienvironmentvariableiddelete) | **Delete** /api/environmentvariable/{id} | 
*EnvironmentVariableApi* | [**ApiEnvironmentvariableIdGet**](docs/EnvironmentVariableApi.md#apienvironmentvariableidget) | **Get** /api/environmentvariable/{id} | 
*EnvironmentVariableApi* | [**ApiEnvironmentvariablePost**](docs/EnvironmentVariableApi.md#apienvironmentvariablepost) | **Post** /api/environmentvariable | 
*RevisionApi* | [**ApiRevisionGet**](docs/RevisionApi.md#apirevisionget) | **Get** /api/revision | 
*RevisionApi* | [**ApiRevisionIdDelete**](docs/RevisionApi.md#apirevisioniddelete) | **Delete** /api/revision/{id} | 
*RevisionApi* | [**ApiRevisionIdGet**](docs/RevisionApi.md#apirevisionidget) | **Get** /api/revision/{id} | 
*RevisionApi* | [**ApiRevisionPost**](docs/RevisionApi.md#apirevisionpost) | **Post** /api/revision | 


## Documentation For Models

 - [App](docs/App.md)
 - [AppDto](docs/AppDto.md)
 - [AppsVm](docs/AppsVm.md)
 - [Certificate](docs/Certificate.md)
 - [CertificateDto](docs/CertificateDto.md)
 - [CertificatesVm](docs/CertificatesVm.md)
 - [Channel](docs/Channel.md)
 - [ChannelDto](docs/ChannelDto.md)
 - [ChannelRevisionSelectionStrategy](docs/ChannelRevisionSelectionStrategy.md)
 - [ChannelsVm](docs/ChannelsVm.md)
 - [CreateAccountCommand](docs/CreateAccountCommand.md)
 - [CreateAppCommand](docs/CreateAppCommand.md)
 - [CreateCertificateCommand](docs/CreateCertificateCommand.md)
 - [CreateChannelCommand](docs/CreateChannelCommand.md)
 - [CreateEnvironmentVariableCommand](docs/CreateEnvironmentVariableCommand.md)
 - [CreateTokenCommand](docs/CreateTokenCommand.md)
 - [DomainEvent](docs/DomainEvent.md)
 - [EnvironmentVariable](docs/EnvironmentVariable.md)
 - [EnvironmentVariableDto](docs/EnvironmentVariableDto.md)
 - [EnvironmentVariablesVm](docs/EnvironmentVariablesVm.md)
 - [RegisterRevisionCommand](docs/RegisterRevisionCommand.md)
 - [Revision](docs/Revision.md)
 - [RevisionDto](docs/RevisionDto.md)
 - [RevisionsVm](docs/RevisionsVm.md)
 - [TokenInfo](docs/TokenInfo.md)


## Documentation For Authorization



### Bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


