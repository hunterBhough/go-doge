// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hunterBhough/bitmask.api/models"
)

// GetAllTransactionsOKCode is the HTTP code returned for type GetAllTransactionsOK
const GetAllTransactionsOKCode int = 200

/*GetAllTransactionsOK successful operation

swagger:response getAllTransactionsOK
*/
type GetAllTransactionsOK struct {

	/*
	  In: Body
	*/
	Payload []models.Question `json:"body,omitempty"`
}

// NewGetAllTransactionsOK creates GetAllTransactionsOK with default headers values
func NewGetAllTransactionsOK() *GetAllTransactionsOK {
	return &GetAllTransactionsOK{}
}

// WithPayload adds the payload to the get all transactions o k response
func (o *GetAllTransactionsOK) WithPayload(payload []models.Question) *GetAllTransactionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all transactions o k response
func (o *GetAllTransactionsOK) SetPayload(payload []models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTransactionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]models.Question, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetAllTransactionsBadRequestCode is the HTTP code returned for type GetAllTransactionsBadRequest
const GetAllTransactionsBadRequestCode int = 400

/*GetAllTransactionsBadRequest Invalid or not enough inputs.

swagger:response getAllTransactionsBadRequest
*/
type GetAllTransactionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllTransactionsBadRequest creates GetAllTransactionsBadRequest with default headers values
func NewGetAllTransactionsBadRequest() *GetAllTransactionsBadRequest {
	return &GetAllTransactionsBadRequest{}
}

// WithPayload adds the payload to the get all transactions bad request response
func (o *GetAllTransactionsBadRequest) WithPayload(payload *models.Error) *GetAllTransactionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all transactions bad request response
func (o *GetAllTransactionsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTransactionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllTransactionsTooManyRequestsCode is the HTTP code returned for type GetAllTransactionsTooManyRequests
const GetAllTransactionsTooManyRequestsCode int = 429

/*GetAllTransactionsTooManyRequests Too many requests.

swagger:response getAllTransactionsTooManyRequests
*/
type GetAllTransactionsTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllTransactionsTooManyRequests creates GetAllTransactionsTooManyRequests with default headers values
func NewGetAllTransactionsTooManyRequests() *GetAllTransactionsTooManyRequests {
	return &GetAllTransactionsTooManyRequests{}
}

// WithPayload adds the payload to the get all transactions too many requests response
func (o *GetAllTransactionsTooManyRequests) WithPayload(payload *models.Error) *GetAllTransactionsTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all transactions too many requests response
func (o *GetAllTransactionsTooManyRequests) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTransactionsTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}