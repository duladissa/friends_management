// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/duladissa/friends_management/models"
)

// PostFriendsCommonListOKCode is the HTTP code returned for type PostFriendsCommonListOK
const PostFriendsCommonListOKCode int = 200

/*PostFriendsCommonListOK successful operation

swagger:response postFriendsCommonListOK
*/
type PostFriendsCommonListOK struct {

	/*
	  In: Body
	*/
	Payload *models.FriendsListResponse `json:"body,omitempty"`
}

// NewPostFriendsCommonListOK creates PostFriendsCommonListOK with default headers values
func NewPostFriendsCommonListOK() *PostFriendsCommonListOK {

	return &PostFriendsCommonListOK{}
}

// WithPayload adds the payload to the post friends common list o k response
func (o *PostFriendsCommonListOK) WithPayload(payload *models.FriendsListResponse) *PostFriendsCommonListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post friends common list o k response
func (o *PostFriendsCommonListOK) SetPayload(payload *models.FriendsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFriendsCommonListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFriendsCommonListUnauthorizedCode is the HTTP code returned for type PostFriendsCommonListUnauthorized
const PostFriendsCommonListUnauthorizedCode int = 401

/*PostFriendsCommonListUnauthorized Unauthorized

swagger:response postFriendsCommonListUnauthorized
*/
type PostFriendsCommonListUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostFriendsCommonListUnauthorized creates PostFriendsCommonListUnauthorized with default headers values
func NewPostFriendsCommonListUnauthorized() *PostFriendsCommonListUnauthorized {

	return &PostFriendsCommonListUnauthorized{}
}

// WithPayload adds the payload to the post friends common list unauthorized response
func (o *PostFriendsCommonListUnauthorized) WithPayload(payload *models.ErrorResponse) *PostFriendsCommonListUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post friends common list unauthorized response
func (o *PostFriendsCommonListUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFriendsCommonListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFriendsCommonListNotFoundCode is the HTTP code returned for type PostFriendsCommonListNotFound
const PostFriendsCommonListNotFoundCode int = 404

/*PostFriendsCommonListNotFound Not Found.

swagger:response postFriendsCommonListNotFound
*/
type PostFriendsCommonListNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostFriendsCommonListNotFound creates PostFriendsCommonListNotFound with default headers values
func NewPostFriendsCommonListNotFound() *PostFriendsCommonListNotFound {

	return &PostFriendsCommonListNotFound{}
}

// WithPayload adds the payload to the post friends common list not found response
func (o *PostFriendsCommonListNotFound) WithPayload(payload *models.ErrorResponse) *PostFriendsCommonListNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post friends common list not found response
func (o *PostFriendsCommonListNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFriendsCommonListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
