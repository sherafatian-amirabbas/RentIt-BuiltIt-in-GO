## Acceptance Tests Documentation
### Cancellation request (CC3)
CC3 - The system should allow site engineers to cancel a plant hire request. Cancellations are  allowed until the day before the plant is due to be delivered. If a cancellation is requested after  the PO has been sent, a request for cancellation should be sent to the supplier. 
* Endpoint: request/cancel/{orderId}
* Request type: POST
* Request (input) 
  * The endpoint must accept the one parameters through the URL (as shown in the endpoint)
  * Purchase order ID - integer
* Response (output)
  * Cancellation request successful - the response must be with HTTP status code **200 OK**. No response body is required.
  * Cancellation request failure - the response must be with HTTP status code **400 Bad Request**. No response body is required.
* Examples
  * Call to "request/cancel/1" (the purchase order with ID 1 exists)
    * Cancellation request successful
  * Call to "request/cancel/1" (the purchase order with ID 1 doesn't exist)
    * Cancellation request failure
  * Call to "request/cancel/asd"
    * Cancellation request failed

### Purchase order confirmation (CC6)
CC3 - The system should produce a PO for every approved plant hire request and forward it to the  corresponding supplier. The supplier may respond that the plant being requested is no longer  available (which means the PO is rejected), or it may respond with a confirmation of the PO. 
* Endpoint: request/purchase_order
* Request type: POST
* Request (input) 
  * The request includes the complete purchase order in the request body as JSON.
  * The object includes - ID [int64], PlantName [string], SiteName [string], SupplierName [string], RequesterName [string], StartDate [time.Time], EndDate [time.Time], TotalHiringCost [float64], Regulator [string], WorkEngineerComment [string], StatusCode [int64], StatusDesc [string]
* Response (output)
  * Purchase order confirmation successful - the response must be with HTTP status code **200 OK**. No response body is required.
  * Purchase order confirmation failure - the response must be with HTTP status code **400 Bad Request**. No response body is required.
* Examples
  * Call to "request/purchase_order" (the body of the request contains the required object as JSON and the plant is available)
    * Purchase order confirmation successful
  * Call to "request/purchase_order" (the body of the request contains the required object as JSON and the plant is NOT available)
    * Purchase order confirmation failure
  * Call to "request/purchase_order" (the body of the request isn't valid JSON or doesn't contain the required object)
    * Purchase order confirmation failure


### Extension request (CC8)
CC3 - The system should allow site engineers to request an extension in order to keep a plant longer  than its initial period of engagement. When an extension is requested, the system should  produce a modified PO and forward it to the supplier. The supplier may accept/reject the  modified PO. 
* Endpoint: request/purchase_order
* Request type: POST
* Request (input) 
  * The request includes the complete purchase order in the request body as JSON.
  * The object includes - ID [int64], PlantName [string], SiteName [string], SupplierName [string], RequesterName [string], StartDate [time.Time], EndDate [time.Time], TotalHiringCost [float64], Regulator [string], WorkEngineerComment [string], StatusCode [int64], StatusDesc [string]
* Response (output)
  * Purchase order confirmation successful - the response must be with HTTP status code **200 OK**. No response body is required.
  * Purchase order confirmation failure - the response must be with HTTP status code **400 Bad Request**. No response body is required.
* Examples
  * Call to "request/purchase_order" (the body of the request contains the required object as JSON and the plant is available)
    * Purchase order confirmation successful
  * Call to "request/purchase_order" (the body of the request contains the required object as JSON and the plant is NOT available)
    * Purchase order confirmation failure
  * Call to "request/purchase_order" (the body of the request isn't valid JSON or doesn't contain the required object)
    * Purchase order confirmation failure


### Remittance advice submission (CC12)
CC12 - The system must submit a remittance advice to the supplier after the invoice is approved
* Endpoint: remittance/create/{purchaseOrderId}/{referenceNumber}
* Request type: POST
* Request (input) 
  * The endpoint must accept the two parameters through the URL (as shown in the endpoint)
  * Purchase order ID - integer
  * Invoice ID (reference number) - integer
* Response (output)
  * Remittance submission successful - the response must be with HTTP status code **200 OK**. No response body is required.
  * Remittance submission failure - the response must be with HTTP status code **400 Bad Request**. No response body is required.
* Examples
  * Call to "remittance/create/1/1" (if the purchase order with ID 1 and the associated invoice with ID 1 both exist)
    * Submission successful
  * Call to "remittance/create/1/1" (if either the purchase order with ID 1 or the associated invoice with ID 1 do not exist)
    * Submission failed 
  * Call to "remittance/create/asd/123"
    * Submission failed
  * Call to "remittance/create/123/asd"
    * Submission failed
  * Call to "remittance/create/asd/asd"
    * Submission failed
