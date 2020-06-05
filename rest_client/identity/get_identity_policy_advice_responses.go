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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// GetIdentityPolicyAdviceReader is a Reader for the GetIdentityPolicyAdvice structure.
type GetIdentityPolicyAdviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityPolicyAdviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityPolicyAdviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetIdentityPolicyAdviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityPolicyAdviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIdentityPolicyAdviceOK creates a GetIdentityPolicyAdviceOK with default headers values
func NewGetIdentityPolicyAdviceOK() *GetIdentityPolicyAdviceOK {
	return &GetIdentityPolicyAdviceOK{}
}

/*GetIdentityPolicyAdviceOK handles this case with default header values.

Returns the document that represents the policy advice
*/
type GetIdentityPolicyAdviceOK struct {
	Payload *rest_model.GetIdentityPolicyAdviceEnvelope
}

func (o *GetIdentityPolicyAdviceOK) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/policy-advice/{serviceId}][%d] getIdentityPolicyAdviceOK  %+v", 200, o.Payload)
}

func (o *GetIdentityPolicyAdviceOK) GetPayload() *rest_model.GetIdentityPolicyAdviceEnvelope {
	return o.Payload
}

func (o *GetIdentityPolicyAdviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.GetIdentityPolicyAdviceEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityPolicyAdviceUnauthorized creates a GetIdentityPolicyAdviceUnauthorized with default headers values
func NewGetIdentityPolicyAdviceUnauthorized() *GetIdentityPolicyAdviceUnauthorized {
	return &GetIdentityPolicyAdviceUnauthorized{}
}

/*GetIdentityPolicyAdviceUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type GetIdentityPolicyAdviceUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *GetIdentityPolicyAdviceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/policy-advice/{serviceId}][%d] getIdentityPolicyAdviceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityPolicyAdviceUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetIdentityPolicyAdviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityPolicyAdviceNotFound creates a GetIdentityPolicyAdviceNotFound with default headers values
func NewGetIdentityPolicyAdviceNotFound() *GetIdentityPolicyAdviceNotFound {
	return &GetIdentityPolicyAdviceNotFound{}
}

/*GetIdentityPolicyAdviceNotFound handles this case with default header values.

The requested resource does not exist
*/
type GetIdentityPolicyAdviceNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *GetIdentityPolicyAdviceNotFound) Error() string {
	return fmt.Sprintf("[GET /identities/{id}/policy-advice/{serviceId}][%d] getIdentityPolicyAdviceNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityPolicyAdviceNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *GetIdentityPolicyAdviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
