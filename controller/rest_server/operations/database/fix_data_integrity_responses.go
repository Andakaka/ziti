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

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/ziti/controller/rest_model"
)

// FixDataIntegrityAcceptedCode is the HTTP code returned for type FixDataIntegrityAccepted
const FixDataIntegrityAcceptedCode int = 202

/*FixDataIntegrityAccepted Base empty response

swagger:response fixDataIntegrityAccepted
*/
type FixDataIntegrityAccepted struct {

	/*
	  In: Body
	*/
	Payload *rest_model.Empty `json:"body,omitempty"`
}

// NewFixDataIntegrityAccepted creates FixDataIntegrityAccepted with default headers values
func NewFixDataIntegrityAccepted() *FixDataIntegrityAccepted {

	return &FixDataIntegrityAccepted{}
}

// WithPayload adds the payload to the fix data integrity accepted response
func (o *FixDataIntegrityAccepted) WithPayload(payload *rest_model.Empty) *FixDataIntegrityAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fix data integrity accepted response
func (o *FixDataIntegrityAccepted) SetPayload(payload *rest_model.Empty) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FixDataIntegrityAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FixDataIntegrityUnauthorizedCode is the HTTP code returned for type FixDataIntegrityUnauthorized
const FixDataIntegrityUnauthorizedCode int = 401

/*FixDataIntegrityUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response fixDataIntegrityUnauthorized
*/
type FixDataIntegrityUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewFixDataIntegrityUnauthorized creates FixDataIntegrityUnauthorized with default headers values
func NewFixDataIntegrityUnauthorized() *FixDataIntegrityUnauthorized {

	return &FixDataIntegrityUnauthorized{}
}

// WithPayload adds the payload to the fix data integrity unauthorized response
func (o *FixDataIntegrityUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *FixDataIntegrityUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fix data integrity unauthorized response
func (o *FixDataIntegrityUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FixDataIntegrityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FixDataIntegrityTooManyRequestsCode is the HTTP code returned for type FixDataIntegrityTooManyRequests
const FixDataIntegrityTooManyRequestsCode int = 429

/*FixDataIntegrityTooManyRequests The resource requested is rate limited and the rate limit has been exceeded

swagger:response fixDataIntegrityTooManyRequests
*/
type FixDataIntegrityTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewFixDataIntegrityTooManyRequests creates FixDataIntegrityTooManyRequests with default headers values
func NewFixDataIntegrityTooManyRequests() *FixDataIntegrityTooManyRequests {

	return &FixDataIntegrityTooManyRequests{}
}

// WithPayload adds the payload to the fix data integrity too many requests response
func (o *FixDataIntegrityTooManyRequests) WithPayload(payload *rest_model.APIErrorEnvelope) *FixDataIntegrityTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fix data integrity too many requests response
func (o *FixDataIntegrityTooManyRequests) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FixDataIntegrityTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
