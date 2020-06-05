// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package certificate_authority

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// DeleteCaOKCode is the HTTP code returned for type DeleteCaOK
const DeleteCaOKCode int = 200

/*DeleteCaOK The delete request was successful and the resource has been removed

swagger:response deleteCaOK
*/
type DeleteCaOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewDeleteCaOK creates DeleteCaOK with default headers values
func NewDeleteCaOK() *DeleteCaOK {

	return &DeleteCaOK{}
}

// WithPayload adds the payload to the delete ca o k response
func (o *DeleteCaOK) WithPayload(payload *rest_model.Empty) *DeleteCaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ca o k response
func (o *DeleteCaOK) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCaBadRequestCode is the HTTP code returned for type DeleteCaBadRequest
const DeleteCaBadRequestCode int = 400

/*DeleteCaBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response deleteCaBadRequest
*/
type DeleteCaBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteCaBadRequest creates DeleteCaBadRequest with default headers values
func NewDeleteCaBadRequest() *DeleteCaBadRequest {

	return &DeleteCaBadRequest{}
}

// WithPayload adds the payload to the delete ca bad request response
func (o *DeleteCaBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteCaBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ca bad request response
func (o *DeleteCaBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCaBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCaUnauthorizedCode is the HTTP code returned for type DeleteCaUnauthorized
const DeleteCaUnauthorizedCode int = 401

/*DeleteCaUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response deleteCaUnauthorized
*/
type DeleteCaUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDeleteCaUnauthorized creates DeleteCaUnauthorized with default headers values
func NewDeleteCaUnauthorized() *DeleteCaUnauthorized {

	return &DeleteCaUnauthorized{}
}

// WithPayload adds the payload to the delete ca unauthorized response
func (o *DeleteCaUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DeleteCaUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ca unauthorized response
func (o *DeleteCaUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
