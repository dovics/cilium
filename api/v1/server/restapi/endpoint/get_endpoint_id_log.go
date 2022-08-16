// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEndpointIDLogHandlerFunc turns a function with the right signature into a get endpoint ID log handler
type GetEndpointIDLogHandlerFunc func(GetEndpointIDLogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEndpointIDLogHandlerFunc) Handle(params GetEndpointIDLogParams) middleware.Responder {
	return fn(params)
}

// GetEndpointIDLogHandler interface for that can handle valid get endpoint ID log params
type GetEndpointIDLogHandler interface {
	Handle(GetEndpointIDLogParams) middleware.Responder
}

// NewGetEndpointIDLog creates a new http.Handler for the get endpoint ID log operation
func NewGetEndpointIDLog(ctx *middleware.Context, handler GetEndpointIDLogHandler) *GetEndpointIDLog {
	return &GetEndpointIDLog{Context: ctx, Handler: handler}
}

/*
GetEndpointIDLog swagger:route GET /endpoint/{id}/log endpoint getEndpointIdLog

Retrieves the status logs associated with this endpoint.
*/
type GetEndpointIDLog struct {
	Context *middleware.Context
	Handler GetEndpointIDLogHandler
}

func (o *GetEndpointIDLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEndpointIDLogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
