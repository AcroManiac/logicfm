// Code generated by go-swagger; DO NOT EDIT.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteParams creates a new DeleteParams object
// with the default values initialized.
func NewDeleteParams() *DeleteParams {
	var ()
	return &DeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteParamsWithTimeout creates a new DeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteParamsWithTimeout(timeout time.Duration) *DeleteParams {
	var ()
	return &DeleteParams{

		timeout: timeout,
	}
}

// NewDeleteParamsWithContext creates a new DeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteParamsWithContext(ctx context.Context) *DeleteParams {
	var ()
	return &DeleteParams{

		Context: ctx,
	}
}

// NewDeleteParamsWithHTTPClient creates a new DeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteParamsWithHTTPClient(client *http.Client) *DeleteParams {
	var ()
	return &DeleteParams{
		HTTPClient: client,
	}
}

/*DeleteParams contains all the parameters to send to the API endpoint
for the delete operation typically these are written to a http.Request
*/
type DeleteParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete params
func (o *DeleteParams) WithTimeout(timeout time.Duration) *DeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete params
func (o *DeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete params
func (o *DeleteParams) WithContext(ctx context.Context) *DeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete params
func (o *DeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) WithHTTPClient(client *http.Client) *DeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete params
func (o *DeleteParams) WithID(id int64) *DeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete params
func (o *DeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
