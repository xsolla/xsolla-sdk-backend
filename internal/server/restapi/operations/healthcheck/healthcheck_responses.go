// Code generated by go-swagger; DO NOT EDIT.

package healthcheck

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HealthcheckOKCode is the HTTP code returned for type HealthcheckOK
const HealthcheckOKCode int = 200

/*HealthcheckOK Service available.

swagger:response healthcheckOK
*/
type HealthcheckOK struct {
}

// NewHealthcheckOK creates HealthcheckOK with default headers values
func NewHealthcheckOK() *HealthcheckOK {

	return &HealthcheckOK{}
}

// WriteResponse to the client
func (o *HealthcheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// HealthcheckServiceUnavailableCode is the HTTP code returned for type HealthcheckServiceUnavailable
const HealthcheckServiceUnavailableCode int = 503

/*HealthcheckServiceUnavailable Service unavailable.

swagger:response healthcheckServiceUnavailable
*/
type HealthcheckServiceUnavailable struct {
}

// NewHealthcheckServiceUnavailable creates HealthcheckServiceUnavailable with default headers values
func NewHealthcheckServiceUnavailable() *HealthcheckServiceUnavailable {

	return &HealthcheckServiceUnavailable{}
}

// WriteResponse to the client
func (o *HealthcheckServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}
