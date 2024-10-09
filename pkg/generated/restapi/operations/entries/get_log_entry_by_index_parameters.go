// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetLogEntryByIndexParams creates a new GetLogEntryByIndexParams object
//
// There are no default values defined in the spec.
func NewGetLogEntryByIndexParams() GetLogEntryByIndexParams {

	return GetLogEntryByIndexParams{}
}

// GetLogEntryByIndexParams contains all the bound params for the get log entry by index operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLogEntryByIndex
type GetLogEntryByIndexParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*specifies the index of the entry in the transparency log to be retrieved
	  Required: true
	  Minimum: 0
	  In: query
	*/
	LogIndex int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLogEntryByIndexParams() beforehand.
func (o *GetLogEntryByIndexParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLogIndex, qhkLogIndex, _ := qs.GetOK("logIndex")
	if err := o.bindLogIndex(qLogIndex, qhkLogIndex, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLogIndex binds and validates parameter LogIndex from query.
func (o *GetLogEntryByIndexParams) bindLogIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("logIndex", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("logIndex", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("logIndex", "query", "int64", raw)
	}
	o.LogIndex = value

	if err := o.validateLogIndex(formats); err != nil {
		return err
	}

	return nil
}

// validateLogIndex carries on validations for parameter LogIndex
func (o *GetLogEntryByIndexParams) validateLogIndex(formats strfmt.Registry) error {

	if err := validate.MinimumInt("logIndex", "query", o.LogIndex, 0, false); err != nil {
		return err
	}

	return nil
}
