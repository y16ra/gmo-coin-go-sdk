# TradingVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | 0 Nomal end | 
**Data** | **interface{}** |  | 
**Responsetime** | **time.Time** |  | 

## Methods

### NewTradingVolume

`func NewTradingVolume(status int32, data interface{}, responsetime time.Time, ) *TradingVolume`

NewTradingVolume instantiates a new TradingVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingVolumeWithDefaults

`func NewTradingVolumeWithDefaults() *TradingVolume`

NewTradingVolumeWithDefaults instantiates a new TradingVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TradingVolume) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradingVolume) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradingVolume) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetData

`func (o *TradingVolume) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TradingVolume) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TradingVolume) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *TradingVolume) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TradingVolume) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetResponsetime

`func (o *TradingVolume) GetResponsetime() time.Time`

GetResponsetime returns the Responsetime field if non-nil, zero value otherwise.

### GetResponsetimeOk

`func (o *TradingVolume) GetResponsetimeOk() (*time.Time, bool)`

GetResponsetimeOk returns a tuple with the Responsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetime

`func (o *TradingVolume) SetResponsetime(v time.Time)`

SetResponsetime sets Responsetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


