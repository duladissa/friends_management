// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/duladissa/friends_management/models"
)

// PostFriendsUpdatesListOKCode is the HTTP code returned for type PostFriendsUpdatesListOK
const PostFriendsUpdatesListOKCode int = 200

/*PostFriendsUpdatesListOK successful operation

swagger:response postFriendsUpdatesListOK
*/
type PostFriendsUpdatesListOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReceiveUpdatesFriendsResponse `json:"body,omitempty"`
}

// NewPostFriendsUpdatesListOK creates PostFriendsUpdatesListOK with default headers values
func NewPostFriendsUpdatesListOK() *PostFriendsUpdatesListOK {

	return &PostFriendsUpdatesListOK{}
}

// WithPayload adds the payload to the post friends updates list o k response
func (o *PostFriendsUpdatesListOK) WithPayload(payload *models.ReceiveUpdatesFriendsResponse) *PostFriendsUpdatesListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post friends updates list o k response
func (o *PostFriendsUpdatesListOK) SetPayload(payload *models.ReceiveUpdatesFriendsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFriendsUpdatesListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFriendsUpdatesListUnauthorizedCode is the HTTP code returned for type PostFriendsUpdatesListUnauthorized
const PostFriendsUpdatesListUnauthorizedCode int = 401

/*PostFriendsUpdatesListUnauthorized Unauthorized

swagger:response postFriendsUpdatesListUnauthorized
*/
type PostFriendsUpdatesListUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostFriendsUpdatesListUnauthorized creates PostFriendsUpdatesListUnauthorized with default headers values
func NewPostFriendsUpdatesListUnauthorized() *PostFriendsUpdatesListUnauthorized {

	return &PostFriendsUpdatesListUnauthorized{}
}

// WithPayload adds the payload to the post friends updates list unauthorized response
func (o *PostFriendsUpdatesListUnauthorized) WithPayload(payload *models.ErrorResponse) *PostFriendsUpdatesListUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post friends updates list unauthorized response
func (o *PostFriendsUpdatesListUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFriendsUpdatesListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFriendsUpdatesListNotFoundCode is the HTTP code returned for type PostFriendsUpdatesListNotFound
const PostFriendsUpdatesListNotFoundCode int = 404

/*PostFriendsUpdatesListNotFound Not Found.

swagger:response postFriendsUpdatesListNotFound
*/
type PostFriendsUpdatesListNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostFriendsUpdatesListNotFound creates PostFriendsUpdatesListNotFound with default headers values
func NewPostFriendsUpdatesListNotFound() *PostFriendsUpdatesListNotFound {

	return &PostFriendsUpdatesListNotFound{}
}

// WithPayload adds the payload to the post friends updates list not found response
func (o *PostFriendsUpdatesListNotFound) WithPayload(payload *models.ErrorResponse) *PostFriendsUpdatesListNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post friends updates list not found response
func (o *PostFriendsUpdatesListNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFriendsUpdatesListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
