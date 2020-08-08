// Code generated by go-swagger; DO NOT EDIT.

package healthcheck

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// HealthcheckHandlerFunc turns a function with the right signature into a healthcheck handler
type HealthcheckHandlerFunc func(HealthcheckParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HealthcheckHandlerFunc) Handle(params HealthcheckParams) middleware.Responder {
	return fn(params)
}

// HealthcheckHandler interface for that can handle valid healthcheck params
type HealthcheckHandler interface {
	Handle(HealthcheckParams) middleware.Responder
}

// NewHealthcheck creates a new http.Handler for the healthcheck operation
func NewHealthcheck(ctx *middleware.Context, handler HealthcheckHandler) *Healthcheck {
	return &Healthcheck{Context: ctx, Handler: handler}
}

/*Healthcheck swagger:route GET /healthcheck Healthcheck healthcheck

Healthcheck application

Healthcheck

*/
type Healthcheck struct {
	Context *middleware.Context
	Handler HealthcheckHandler
}

func (o *Healthcheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHealthcheckParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
