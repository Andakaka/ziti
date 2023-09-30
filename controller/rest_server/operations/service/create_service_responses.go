// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/ziti/controller/rest_model"
)

// CreateServiceCreatedCode is the HTTP code returned for type CreateServiceCreated
const CreateServiceCreatedCode int = 201

/*CreateServiceCreated The create request was successful and the resource has been added at the following location

swagger:response createServiceCreated
*/
type CreateServiceCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.CreateEnvelope `json:"body,omitempty"`
}

// NewCreateServiceCreated creates CreateServiceCreated with default headers values
func NewCreateServiceCreated() *CreateServiceCreated {

	return &CreateServiceCreated{}
}

// WithPayload adds the payload to the create service created response
func (o *CreateServiceCreated) WithPayload(payload *rest_model.CreateEnvelope) *CreateServiceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service created response
func (o *CreateServiceCreated) SetPayload(payload *rest_model.CreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServiceBadRequestCode is the HTTP code returned for type CreateServiceBadRequest
const CreateServiceBadRequestCode int = 400

/*CreateServiceBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createServiceBadRequest
*/
type CreateServiceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateServiceBadRequest creates CreateServiceBadRequest with default headers values
func NewCreateServiceBadRequest() *CreateServiceBadRequest {

	return &CreateServiceBadRequest{}
}

// WithPayload adds the payload to the create service bad request response
func (o *CreateServiceBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateServiceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service bad request response
func (o *CreateServiceBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServiceUnauthorizedCode is the HTTP code returned for type CreateServiceUnauthorized
const CreateServiceUnauthorizedCode int = 401

/*CreateServiceUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createServiceUnauthorized
*/
type CreateServiceUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateServiceUnauthorized creates CreateServiceUnauthorized with default headers values
func NewCreateServiceUnauthorized() *CreateServiceUnauthorized {

	return &CreateServiceUnauthorized{}
}

// WithPayload adds the payload to the create service unauthorized response
func (o *CreateServiceUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateServiceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service unauthorized response
func (o *CreateServiceUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
