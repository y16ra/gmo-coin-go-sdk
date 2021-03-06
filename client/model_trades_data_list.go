/*
GMO Coin APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: dev@tricoro.tech
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TradesDataList struct for TradesDataList
type TradesDataList struct {
	Price *int32 `json:"price,omitempty"`
	Side *string `json:"side,omitempty"`
	Size *string `json:"size,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewTradesDataList instantiates a new TradesDataList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradesDataList() *TradesDataList {
	this := TradesDataList{}
	return &this
}

// NewTradesDataListWithDefaults instantiates a new TradesDataList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradesDataListWithDefaults() *TradesDataList {
	this := TradesDataList{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TradesDataList) GetPrice() int32 {
	if o == nil || o.Price == nil {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesDataList) GetPriceOk() (*int32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TradesDataList) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *TradesDataList) SetPrice(v int32) {
	o.Price = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *TradesDataList) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesDataList) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *TradesDataList) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *TradesDataList) SetSide(v string) {
	o.Side = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TradesDataList) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesDataList) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TradesDataList) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *TradesDataList) SetSize(v string) {
	o.Size = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TradesDataList) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradesDataList) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TradesDataList) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TradesDataList) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o TradesDataList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Side != nil {
		toSerialize["side"] = o.Side
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableTradesDataList struct {
	value *TradesDataList
	isSet bool
}

func (v NullableTradesDataList) Get() *TradesDataList {
	return v.value
}

func (v *NullableTradesDataList) Set(val *TradesDataList) {
	v.value = val
	v.isSet = true
}

func (v NullableTradesDataList) IsSet() bool {
	return v.isSet
}

func (v *NullableTradesDataList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradesDataList(val *TradesDataList) *NullableTradesDataList {
	return &NullableTradesDataList{value: val, isSet: true}
}

func (v NullableTradesDataList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradesDataList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


