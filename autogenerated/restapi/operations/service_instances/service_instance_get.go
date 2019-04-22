// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ServiceInstanceGetHandlerFunc turns a function with the right signature into a service instance get handler
type ServiceInstanceGetHandlerFunc func(ServiceInstanceGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceInstanceGetHandlerFunc) Handle(params ServiceInstanceGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ServiceInstanceGetHandler interface for that can handle valid service instance get params
type ServiceInstanceGetHandler interface {
	Handle(ServiceInstanceGetParams, interface{}) middleware.Responder
}

// NewServiceInstanceGet creates a new http.Handler for the service instance get operation
func NewServiceInstanceGet(ctx *middleware.Context, handler ServiceInstanceGetHandler) *ServiceInstanceGet {
	return &ServiceInstanceGet{Context: ctx, Handler: handler}
}

/*ServiceInstanceGet swagger:route GET /v2/service_instances/{instance_id} ServiceInstances serviceInstanceGet

gets a service instance

*/
type ServiceInstanceGet struct {
	Context *middleware.Context
	Handler ServiceInstanceGetHandler
}

func (o *ServiceInstanceGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewServiceInstanceGetParams()

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
