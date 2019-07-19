// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostFriendsUpdatesSubscribeHandlerFunc turns a function with the right signature into a post friends updates subscribe handler
type PostFriendsUpdatesSubscribeHandlerFunc func(PostFriendsUpdatesSubscribeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostFriendsUpdatesSubscribeHandlerFunc) Handle(params PostFriendsUpdatesSubscribeParams) middleware.Responder {
	return fn(params)
}

// PostFriendsUpdatesSubscribeHandler interface for that can handle valid post friends updates subscribe params
type PostFriendsUpdatesSubscribeHandler interface {
	Handle(PostFriendsUpdatesSubscribeParams) middleware.Responder
}

// NewPostFriendsUpdatesSubscribe creates a new http.Handler for the post friends updates subscribe operation
func NewPostFriendsUpdatesSubscribe(ctx *middleware.Context, handler PostFriendsUpdatesSubscribeHandler) *PostFriendsUpdatesSubscribe {
	return &PostFriendsUpdatesSubscribe{Context: ctx, Handler: handler}
}

/*PostFriendsUpdatesSubscribe swagger:route POST /friends/updates/subscribe friends postFriendsUpdatesSubscribe

Subscribe to a given users updates.

4. As a user, I need an API to subscribe to updates from an email address.</br> Please note that subscribing to updates is NOT equivalent to adding a friend connection

*/
type PostFriendsUpdatesSubscribe struct {
	Context *middleware.Context
	Handler PostFriendsUpdatesSubscribeHandler
}

func (o *PostFriendsUpdatesSubscribe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostFriendsUpdatesSubscribeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
