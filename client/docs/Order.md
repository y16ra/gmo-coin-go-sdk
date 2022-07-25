# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | 0 Nomal end | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Responsetime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Order) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Order) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Order) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Order) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Order) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Order) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Order) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetResponsetime

`func (o *Order) GetResponsetime() time.Time`

GetResponsetime returns the Responsetime field if non-nil, zero value otherwise.

### GetResponsetimeOk

`func (o *Order) GetResponsetimeOk() (*time.Time, bool)`

GetResponsetimeOk returns a tuple with the Responsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetime

`func (o *Order) SetResponsetime(v time.Time)`

SetResponsetime sets Responsetime field to given value.

### HasResponsetime

`func (o *Order) HasResponsetime() bool`

HasResponsetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


