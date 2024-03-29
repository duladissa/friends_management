// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReceiveUpdatesFriendsResponse List of friends along with the success.
// swagger:model ReceiveUpdatesFriendsResponse
type ReceiveUpdatesFriendsResponse struct {

	// List of friends.
	// Required: true
	Recipients []string `json:"recipients"`

	// success.
	// Required: true
	Success *bool `json:"success"`
}

// Validate validates this receive updates friends response
func (m *ReceiveUpdatesFriendsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReceiveUpdatesFriendsResponse) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("recipients", "body", m.Recipients); err != nil {
		return err
	}

	return nil
}

func (m *ReceiveUpdatesFriendsResponse) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", m.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReceiveUpdatesFriendsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceiveUpdatesFriendsResponse) UnmarshalBinary(b []byte) error {
	var res ReceiveUpdatesFriendsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
