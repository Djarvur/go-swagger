// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/Djarvur/go-swagger/examples/oauth2/models"
)

// GetIDHandlerFunc turns a function with the right signature into a get Id handler
type GetIDHandlerFunc func(GetIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIDHandlerFunc) Handle(params GetIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetIDHandler interface for that can handle valid get Id params
type GetIDHandler interface {
	Handle(GetIDParams, *models.Principal) middleware.Responder
}

// NewGetID creates a new http.Handler for the get Id operation
func NewGetID(ctx *middleware.Context, handler GetIDHandler) *GetID {
	return &GetID{Context: ctx, Handler: handler}
}

/*GetID swagger:route GET /customers customers getId

Get a customerId given an SSN

*/
type GetID struct {
	Context *middleware.Context
	Handler GetIDHandler
}

func (o *GetID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetIDParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
