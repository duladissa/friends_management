// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostFriendsUpdatesListHandlerFunc turns a function with the right signature into a post friends updates list handler
type PostFriendsUpdatesListHandlerFunc func(PostFriendsUpdatesListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostFriendsUpdatesListHandlerFunc) Handle(params PostFriendsUpdatesListParams) middleware.Responder {
	return fn(params)
}

// PostFriendsUpdatesListHandler interface for that can handle valid post friends updates list params
type PostFriendsUpdatesListHandler interface {
	Handle(PostFriendsUpdatesListParams) middleware.Responder
}

// NewPostFriendsUpdatesList creates a new http.Handler for the post friends updates list operation
func NewPostFriendsUpdatesList(ctx *middleware.Context, handler PostFriendsUpdatesListHandler) *PostFriendsUpdatesList {
	return &PostFriendsUpdatesList{Context: ctx, Handler: handler}
}

/*PostFriendsUpdatesList swagger:route POST /friends/updates/list friends postFriendsUpdatesList

6. Retrieve all users who can receive updates from given user.

6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.</br> Eligibility for receiving updates from i.e. john@example.com</br> <ul> <li>has not blocked updates from john@example.com, and</li> <li>at least one of the following:</li> <ul> <li>has a friend connection with john@example.com</li> <li>has subscribed to updates from john@example.com</li> <li>has been @mentioned in the update</li> </ul> </ul>

*/
type PostFriendsUpdatesList struct {
	Context *middleware.Context
	Handler PostFriendsUpdatesListHandler
}

func (o *PostFriendsUpdatesList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostFriendsUpdatesListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
