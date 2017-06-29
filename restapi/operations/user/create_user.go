package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateUserHandlerFunc turns a function with the right signature into a create user handler
type CreateUserHandlerFunc func(CreateUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUserHandlerFunc) Handle(params CreateUserParams) middleware.Responder {
	return fn(params)
}

// CreateUserHandler interface for that can handle valid create user params
type CreateUserHandler interface {
	Handle(CreateUserParams) middleware.Responder
}

// NewCreateUser creates a new http.Handler for the create user operation
func NewCreateUser(ctx *middleware.Context, handler CreateUserHandler) *CreateUser {
	return &CreateUser{Context: ctx, Handler: handler}
}

/*CreateUser swagger:route POST /user user createUser

Create user

This can only be done by the logged in user.

*/
type CreateUser struct {
	Context *middleware.Context
	Handler CreateUserHandler
}

func (o *CreateUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
