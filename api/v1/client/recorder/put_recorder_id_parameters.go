// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

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

	"github.com/cilium/cilium/api/v1/models"
)

// NewPutRecorderIDParams creates a new PutRecorderIDParams object
// with the default values initialized.
func NewPutRecorderIDParams() *PutRecorderIDParams {
	var ()
	return &PutRecorderIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRecorderIDParamsWithTimeout creates a new PutRecorderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRecorderIDParamsWithTimeout(timeout time.Duration) *PutRecorderIDParams {
	var ()
	return &PutRecorderIDParams{

		timeout: timeout,
	}
}

// NewPutRecorderIDParamsWithContext creates a new PutRecorderIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRecorderIDParamsWithContext(ctx context.Context) *PutRecorderIDParams {
	var ()
	return &PutRecorderIDParams{

		Context: ctx,
	}
}

// NewPutRecorderIDParamsWithHTTPClient creates a new PutRecorderIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRecorderIDParamsWithHTTPClient(client *http.Client) *PutRecorderIDParams {
	var ()
	return &PutRecorderIDParams{
		HTTPClient: client,
	}
}

/*
PutRecorderIDParams contains all the parameters to send to the API endpoint
for the put recorder ID operation typically these are written to a http.Request
*/
type PutRecorderIDParams struct {

	/*Config
	  Recorder configuration

	*/
	Config *models.RecorderSpec
	/*ID
	  ID of recorder

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put recorder ID params
func (o *PutRecorderIDParams) WithTimeout(timeout time.Duration) *PutRecorderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put recorder ID params
func (o *PutRecorderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put recorder ID params
func (o *PutRecorderIDParams) WithContext(ctx context.Context) *PutRecorderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put recorder ID params
func (o *PutRecorderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put recorder ID params
func (o *PutRecorderIDParams) WithHTTPClient(client *http.Client) *PutRecorderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put recorder ID params
func (o *PutRecorderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put recorder ID params
func (o *PutRecorderIDParams) WithConfig(config *models.RecorderSpec) *PutRecorderIDParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put recorder ID params
func (o *PutRecorderIDParams) SetConfig(config *models.RecorderSpec) {
	o.Config = config
}

// WithID adds the id to the put recorder ID params
func (o *PutRecorderIDParams) WithID(id int64) *PutRecorderIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put recorder ID params
func (o *PutRecorderIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutRecorderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
