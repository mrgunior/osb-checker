// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
)

// ServiceInstanceProvisionOKCode is the HTTP code returned for type ServiceInstanceProvisionOK
const ServiceInstanceProvisionOKCode int = 200

/*ServiceInstanceProvisionOK OK

swagger:response serviceInstanceProvisionOK
*/
type ServiceInstanceProvisionOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceInstanceProvision `json:"body,omitempty"`
}

// NewServiceInstanceProvisionOK creates ServiceInstanceProvisionOK with default headers values
func NewServiceInstanceProvisionOK() *ServiceInstanceProvisionOK {

	return &ServiceInstanceProvisionOK{}
}

// WithPayload adds the payload to the service instance provision o k response
func (o *ServiceInstanceProvisionOK) WithPayload(payload *models.ServiceInstanceProvision) *ServiceInstanceProvisionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance provision o k response
func (o *ServiceInstanceProvisionOK) SetPayload(payload *models.ServiceInstanceProvision) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceProvisionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceProvisionCreatedCode is the HTTP code returned for type ServiceInstanceProvisionCreated
const ServiceInstanceProvisionCreatedCode int = 201

/*ServiceInstanceProvisionCreated Created

swagger:response serviceInstanceProvisionCreated
*/
type ServiceInstanceProvisionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceInstanceProvision `json:"body,omitempty"`
}

// NewServiceInstanceProvisionCreated creates ServiceInstanceProvisionCreated with default headers values
func NewServiceInstanceProvisionCreated() *ServiceInstanceProvisionCreated {

	return &ServiceInstanceProvisionCreated{}
}

// WithPayload adds the payload to the service instance provision created response
func (o *ServiceInstanceProvisionCreated) WithPayload(payload *models.ServiceInstanceProvision) *ServiceInstanceProvisionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance provision created response
func (o *ServiceInstanceProvisionCreated) SetPayload(payload *models.ServiceInstanceProvision) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceProvisionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceProvisionAcceptedCode is the HTTP code returned for type ServiceInstanceProvisionAccepted
const ServiceInstanceProvisionAcceptedCode int = 202

/*ServiceInstanceProvisionAccepted Accepted

swagger:response serviceInstanceProvisionAccepted
*/
type ServiceInstanceProvisionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceInstanceAsyncOperation `json:"body,omitempty"`
}

// NewServiceInstanceProvisionAccepted creates ServiceInstanceProvisionAccepted with default headers values
func NewServiceInstanceProvisionAccepted() *ServiceInstanceProvisionAccepted {

	return &ServiceInstanceProvisionAccepted{}
}

// WithPayload adds the payload to the service instance provision accepted response
func (o *ServiceInstanceProvisionAccepted) WithPayload(payload *models.ServiceInstanceAsyncOperation) *ServiceInstanceProvisionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance provision accepted response
func (o *ServiceInstanceProvisionAccepted) SetPayload(payload *models.ServiceInstanceAsyncOperation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceProvisionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceProvisionBadRequestCode is the HTTP code returned for type ServiceInstanceProvisionBadRequest
const ServiceInstanceProvisionBadRequestCode int = 400

/*ServiceInstanceProvisionBadRequest Bad Request

swagger:response serviceInstanceProvisionBadRequest
*/
type ServiceInstanceProvisionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceProvisionBadRequest creates ServiceInstanceProvisionBadRequest with default headers values
func NewServiceInstanceProvisionBadRequest() *ServiceInstanceProvisionBadRequest {

	return &ServiceInstanceProvisionBadRequest{}
}

// WithPayload adds the payload to the service instance provision bad request response
func (o *ServiceInstanceProvisionBadRequest) WithPayload(payload *models.Error) *ServiceInstanceProvisionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance provision bad request response
func (o *ServiceInstanceProvisionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceProvisionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceProvisionConflictCode is the HTTP code returned for type ServiceInstanceProvisionConflict
const ServiceInstanceProvisionConflictCode int = 409

/*ServiceInstanceProvisionConflict Conflict

swagger:response serviceInstanceProvisionConflict
*/
type ServiceInstanceProvisionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceProvisionConflict creates ServiceInstanceProvisionConflict with default headers values
func NewServiceInstanceProvisionConflict() *ServiceInstanceProvisionConflict {

	return &ServiceInstanceProvisionConflict{}
}

// WithPayload adds the payload to the service instance provision conflict response
func (o *ServiceInstanceProvisionConflict) WithPayload(payload *models.Error) *ServiceInstanceProvisionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance provision conflict response
func (o *ServiceInstanceProvisionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceProvisionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceProvisionUnprocessableEntityCode is the HTTP code returned for type ServiceInstanceProvisionUnprocessableEntity
const ServiceInstanceProvisionUnprocessableEntityCode int = 422

/*ServiceInstanceProvisionUnprocessableEntity Unprocessable Entity

swagger:response serviceInstanceProvisionUnprocessableEntity
*/
type ServiceInstanceProvisionUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceProvisionUnprocessableEntity creates ServiceInstanceProvisionUnprocessableEntity with default headers values
func NewServiceInstanceProvisionUnprocessableEntity() *ServiceInstanceProvisionUnprocessableEntity {

	return &ServiceInstanceProvisionUnprocessableEntity{}
}

// WithPayload adds the payload to the service instance provision unprocessable entity response
func (o *ServiceInstanceProvisionUnprocessableEntity) WithPayload(payload *models.Error) *ServiceInstanceProvisionUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance provision unprocessable entity response
func (o *ServiceInstanceProvisionUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceProvisionUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
