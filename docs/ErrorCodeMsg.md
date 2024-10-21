# ErrorCodeMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpStatusCode** | **int32** | The HTTP equivalent status code | 
**ErrorCode** | **string** | A unique machine-friendly identifier for the error from a global list of enumerated identifier strings | 
**Message** | **string** | A user-friendly error message | 
**DebugId** | **string** | A unique identifier for the instance of this error. This can be used to help with troubleshooting. May be the same as a trace id (see Trace Headers) for ease of log tracing | 

## Methods

### NewErrorCodeMsg

`func NewErrorCodeMsg(httpStatusCode int32, errorCode string, message string, debugId string, ) *ErrorCodeMsg`

NewErrorCodeMsg instantiates a new ErrorCodeMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorCodeMsgWithDefaults

`func NewErrorCodeMsgWithDefaults() *ErrorCodeMsg`

NewErrorCodeMsgWithDefaults instantiates a new ErrorCodeMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatusCode

`func (o *ErrorCodeMsg) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *ErrorCodeMsg) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *ErrorCodeMsg) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.


### GetErrorCode

`func (o *ErrorCodeMsg) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorCodeMsg) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorCodeMsg) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMessage

`func (o *ErrorCodeMsg) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorCodeMsg) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorCodeMsg) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDebugId

`func (o *ErrorCodeMsg) GetDebugId() string`

GetDebugId returns the DebugId field if non-nil, zero value otherwise.

### GetDebugIdOk

`func (o *ErrorCodeMsg) GetDebugIdOk() (*string, bool)`

GetDebugIdOk returns a tuple with the DebugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugId

`func (o *ErrorCodeMsg) SetDebugId(v string)`

SetDebugId sets DebugId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


