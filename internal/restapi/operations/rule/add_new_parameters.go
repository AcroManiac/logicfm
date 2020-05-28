// Code generated by go-swagger; DO NOT EDIT.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ahamtat/logicfm/internal/models"
)

// NewAddNewParams creates a new AddNewParams object
// no default values defined in spec.
func NewAddNewParams() AddNewParams {

	return AddNewParams{}
}

// AddNewParams contains all the bound params for the add new operation
// typically these are obtained from a http.Request
//
// swagger:parameters addNew
type AddNewParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	MusrvID int64
	/*
	  In: body
	*/
	Rule *models.Rule
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddNewParams() beforehand.
func (o *AddNewParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMusrvID, rhkMusrvID, _ := route.Params.GetOK("musrvId")
	if err := o.bindMusrvID(rMusrvID, rhkMusrvID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Rule
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("rule", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Rule = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMusrvID binds and validates parameter MusrvID from path.
func (o *AddNewParams) bindMusrvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("musrvId", "path", "int64", raw)
	}
	o.MusrvID = value

	return nil
}
