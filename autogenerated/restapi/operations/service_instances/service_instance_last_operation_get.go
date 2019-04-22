// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ServiceInstanceLastOperationGetHandlerFunc turns a function with the right signature into a service instance last operation get handler
type ServiceInstanceLastOperationGetHandlerFunc func(ServiceInstanceLastOperationGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceInstanceLastOperationGetHandlerFunc) Handle(params ServiceInstanceLastOperationGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ServiceInstanceLastOperationGetHandler interface for that can handle valid service instance last operation get params
type ServiceInstanceLastOperationGetHandler interface {
	Handle(ServiceInstanceLastOperationGetParams, interface{}) middleware.Responder
}

// NewServiceInstanceLastOperationGet creates a new http.Handler for the service instance last operation get operation
func NewServiceInstanceLastOperationGet(ctx *middleware.Context, handler ServiceInstanceLastOperationGetHandler) *ServiceInstanceLastOperationGet {
	return &ServiceInstanceLastOperationGet{Context: ctx, Handler: handler}
}

/*ServiceInstanceLastOperationGet swagger:route GET /v2/service_instances/{instance_id}/last_operation ServiceInstances serviceInstanceLastOperationGet

last requested operation state for service instance

*/
type ServiceInstanceLastOperationGet struct {
	Context *middleware.Context
	Handler ServiceInstanceLastOperationGetHandler
}

func (o *ServiceInstanceLastOperationGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewServiceInstanceLastOperationGetParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
