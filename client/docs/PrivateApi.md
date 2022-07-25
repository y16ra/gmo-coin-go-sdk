# \PrivateApi

All URIs are relative to *https://api.coin.z.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateV1AccountAssetsGet**](PrivateApi.md#PrivateV1AccountAssetsGet) | **Get** /private/v1/account/assets | Get account assets
[**PrivateV1AccountMarginGet**](PrivateApi.md#PrivateV1AccountMarginGet) | **Get** /private/v1/account/margin | Get account margin
[**PrivateV1AccountTradingVolumeGet**](PrivateApi.md#PrivateV1AccountTradingVolumeGet) | **Get** /private/v1/account/tradingVolume | Get trading volume
[**PrivateV1OrderGet**](PrivateApi.md#PrivateV1OrderGet) | **Get** /private/v1/order | Get orders
[**PrivateV1OrderPost**](PrivateApi.md#PrivateV1OrderPost) | **Post** /private/v1/order | Order



## PrivateV1AccountAssetsGet

> Assets PrivateV1AccountAssetsGet(ctx).Execute()

Get account assets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateApi.PrivateV1AccountAssetsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateApi.PrivateV1AccountAssetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateV1AccountAssetsGet`: Assets
    fmt.Fprintf(os.Stdout, "Response from `PrivateApi.PrivateV1AccountAssetsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateV1AccountAssetsGetRequest struct via the builder pattern


### Return type

[**Assets**](Assets.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiSign](../README.md#ApiSign), [ApiTimestamp](../README.md#ApiTimestamp)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateV1AccountMarginGet

> AccountMargin PrivateV1AccountMarginGet(ctx).Execute()

Get account margin

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateApi.PrivateV1AccountMarginGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateApi.PrivateV1AccountMarginGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateV1AccountMarginGet`: AccountMargin
    fmt.Fprintf(os.Stdout, "Response from `PrivateApi.PrivateV1AccountMarginGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateV1AccountMarginGetRequest struct via the builder pattern


### Return type

[**AccountMargin**](AccountMargin.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiSign](../README.md#ApiSign), [ApiTimestamp](../README.md#ApiTimestamp)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateV1AccountTradingVolumeGet

> TradingVolume PrivateV1AccountTradingVolumeGet(ctx).Execute()

Get trading volume

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateApi.PrivateV1AccountTradingVolumeGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateApi.PrivateV1AccountTradingVolumeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateV1AccountTradingVolumeGet`: TradingVolume
    fmt.Fprintf(os.Stdout, "Response from `PrivateApi.PrivateV1AccountTradingVolumeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateV1AccountTradingVolumeGetRequest struct via the builder pattern


### Return type

[**TradingVolume**](TradingVolume.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiSign](../README.md#ApiSign), [ApiTimestamp](../README.md#ApiTimestamp)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateV1OrderGet

> Order PrivateV1OrderGet(ctx).OrderId(orderId).Execute()

Get orders

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := "orderId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateApi.PrivateV1OrderGet(context.Background()).OrderId(orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateApi.PrivateV1OrderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateV1OrderGet`: Order
    fmt.Fprintf(os.Stdout, "Response from `PrivateApi.PrivateV1OrderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateV1OrderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** |  | 

### Return type

[**Order**](Order.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiSign](../README.md#ApiSign), [ApiTimestamp](../README.md#ApiTimestamp)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateV1OrderPost

> InlineResponse200 PrivateV1OrderPost(ctx).InlineObject(inlineObject).Execute()

Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject := *openapiclient.NewInlineObject() // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateApi.PrivateV1OrderPost(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateApi.PrivateV1OrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateV1OrderPost`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `PrivateApi.PrivateV1OrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateV1OrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[ApiKey](../README.md#ApiKey), [ApiSign](../README.md#ApiSign), [ApiTimestamp](../README.md#ApiTimestamp)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

