# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pkg/transport/gRPC/gRPCService.proto](#pkg/transport/gRPC/gRPCService.proto)
    - [Plant](#grpc.Plant)
    - [PlantAvailabilityRequest](#grpc.PlantAvailabilityRequest)
    - [PlantAvailabilityResponse](#grpc.PlantAvailabilityResponse)
    - [PlantPrice](#grpc.PlantPrice)
    - [PlantPriceRequest](#grpc.PlantPriceRequest)
    - [PlantPriceResponse](#grpc.PlantPriceResponse)
    - [PlantsRequest](#grpc.PlantsRequest)
    - [PlantsResponse](#grpc.PlantsResponse)
  
    - [gRPCService](#grpc.gRPCService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="pkg/transport/gRPC/gRPCService.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pkg/transport/gRPC/gRPCService.proto



<a name="grpc.Plant"></a>

### Plant
Represents a plant


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Id | [int64](#int64) |  | Unique ID of a plant |
| Name | [string](#string) |  | Name of the machine |
| Description | [string](#string) |  | Description of the machine |
| PricePerDay | [double](#double) |  | Daily cost of the machine |






<a name="grpc.PlantAvailabilityRequest"></a>

### PlantAvailabilityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Id | [int64](#int64) |  | ID of the plant to be fetched |
| startDate | [string](#string) |  | Start date of the duration |
| endDate | [string](#string) |  | End date of the duration |






<a name="grpc.PlantAvailabilityResponse"></a>

### PlantAvailabilityResponse
Represents a plants availability for a given time duration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isAvailable | [bool](#bool) |  | Availability of the plant during the selected timeframe |






<a name="grpc.PlantPrice"></a>

### PlantPrice
Represents a plant&#39;s price for a duration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PlantId | [int64](#int64) |  | ID of the plant |
| StartDate | [string](#string) |  | Start date of the duration |
| EndDate | [string](#string) |  | End date of the duration |
| PricePerDuration | [double](#double) |  | Total price for duration (PricePerDay * days) |






<a name="grpc.PlantPriceRequest"></a>

### PlantPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Id | [int64](#int64) |  | ID of the plant to be fetched |
| startDate | [string](#string) |  | Start date of the duration |
| endDate | [string](#string) |  | End date of the duration |






<a name="grpc.PlantPriceResponse"></a>

### PlantPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [PlantPrice](#grpc.PlantPrice) |  | Generated plant price item |






<a name="grpc.PlantsRequest"></a>

### PlantsRequest







<a name="grpc.PlantsResponse"></a>

### PlantsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [Plant](#grpc.Plant) | repeated | Fetched plant items |





 

 

 


<a name="grpc.gRPCService"></a>

### gRPCService
Service for handling plant information

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAllPlants | [PlantsRequest](#grpc.PlantsRequest) | [PlantsResponse](#grpc.PlantsResponse) | Fetches all the plants in the repository |
| GetPlantPrice | [PlantPriceRequest](#grpc.PlantPriceRequest) | [PlantPriceResponse](#grpc.PlantPriceResponse) | Fetches the total plant price for a duration |
| IsPlantAvailable | [PlantAvailabilityRequest](#grpc.PlantAvailabilityRequest) | [PlantAvailabilityResponse](#grpc.PlantAvailabilityResponse) | Checks whether the selected plant is available within a timeframe |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

