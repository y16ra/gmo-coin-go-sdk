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

// Status struct for Status
type Status struct {
	// 0 Nomal end
	Status int32 `json:"status"`
	Data interface{} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus(status int32, data interface{}, responsetime time.Time) *Status {
	this := Status{}
	this.Status = status
	this.Data = data
	this.Responsetime = responsetime
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetStatus returns the Status field value
func (o *Status) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Status) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Status) SetStatus(v int32) {
	o.Status = v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Status) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetDataOk() (*interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Status) SetData(v interface{}) {
	o.Data = v
}

// GetResponsetime returns the Responsetime field value
func (o *Status) GetResponsetime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Responsetime
}

// GetResponsetimeOk returns a tuple with the Responsetime field value
// and a boolean to check if the value has been set.
func (o *Status) GetResponsetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Responsetime, true
}

// SetResponsetime sets field value
func (o *Status) SetResponsetime(v time.Time) {
	o.Responsetime = v
}

func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["responsetime"] = o.Responsetime
	}
	return json.Marshal(toSerialize)
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


