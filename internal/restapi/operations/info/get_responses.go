// Code generated by go-swagger; DO NOT EDIT.

package info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ahamtat/logicfm/internal/models"
)

// GetOKCode is the HTTP code returned for type GetOK
const GetOKCode int = 200

/*GetOK displays microservice's name and version

swagger:response getOK
*/
type GetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Info `json:"body,omitempty"`
}

// NewGetOK creates GetOK with default headers values
func NewGetOK() *GetOK {

	return &GetOK{}
}

// WithPayload adds the payload to the get o k response
func (o *GetOK) WithPayload(payload *models.Info) *GetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get o k response
func (o *GetOK) SetPayload(payload *models.Info) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDefault error

swagger:response getDefault
*/
type GetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDefault creates GetDefault with default headers values
func NewGetDefault(code int) *GetDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get default response
func (o *GetDefault) WithStatusCode(code int) *GetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get default response
func (o *GetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get default response
func (o *GetDefault) WithPayload(payload *models.Error) *GetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get default response
func (o *GetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
