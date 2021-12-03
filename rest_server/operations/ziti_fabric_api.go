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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/openziti/fabric/rest_server/operations/circuit"
	"github.com/openziti/fabric/rest_server/operations/inspect"
	"github.com/openziti/fabric/rest_server/operations/link"
	"github.com/openziti/fabric/rest_server/operations/router"
	"github.com/openziti/fabric/rest_server/operations/service"
	"github.com/openziti/fabric/rest_server/operations/terminator"
)

// NewZitiFabricAPI creates a new ZitiFabric instance
func NewZitiFabricAPI(spec *loads.Document) *ZitiFabricAPI {
	return &ZitiFabricAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		RouterCreateRouterHandler: router.CreateRouterHandlerFunc(func(params router.CreateRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.CreateRouter has not yet been implemented")
		}),
		ServiceCreateServiceHandler: service.CreateServiceHandlerFunc(func(params service.CreateServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.CreateService has not yet been implemented")
		}),
		TerminatorCreateTerminatorHandler: terminator.CreateTerminatorHandlerFunc(func(params terminator.CreateTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.CreateTerminator has not yet been implemented")
		}),
		CircuitDeleteCircuitHandler: circuit.DeleteCircuitHandlerFunc(func(params circuit.DeleteCircuitParams) middleware.Responder {
			return middleware.NotImplemented("operation circuit.DeleteCircuit has not yet been implemented")
		}),
		RouterDeleteRouterHandler: router.DeleteRouterHandlerFunc(func(params router.DeleteRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.DeleteRouter has not yet been implemented")
		}),
		ServiceDeleteServiceHandler: service.DeleteServiceHandlerFunc(func(params service.DeleteServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteService has not yet been implemented")
		}),
		TerminatorDeleteTerminatorHandler: terminator.DeleteTerminatorHandlerFunc(func(params terminator.DeleteTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.DeleteTerminator has not yet been implemented")
		}),
		CircuitDetailCircuitHandler: circuit.DetailCircuitHandlerFunc(func(params circuit.DetailCircuitParams) middleware.Responder {
			return middleware.NotImplemented("operation circuit.DetailCircuit has not yet been implemented")
		}),
		LinkDetailLinkHandler: link.DetailLinkHandlerFunc(func(params link.DetailLinkParams) middleware.Responder {
			return middleware.NotImplemented("operation link.DetailLink has not yet been implemented")
		}),
		RouterDetailRouterHandler: router.DetailRouterHandlerFunc(func(params router.DetailRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.DetailRouter has not yet been implemented")
		}),
		ServiceDetailServiceHandler: service.DetailServiceHandlerFunc(func(params service.DetailServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DetailService has not yet been implemented")
		}),
		TerminatorDetailTerminatorHandler: terminator.DetailTerminatorHandlerFunc(func(params terminator.DetailTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.DetailTerminator has not yet been implemented")
		}),
		InspectInspectHandler: inspect.InspectHandlerFunc(func(params inspect.InspectParams) middleware.Responder {
			return middleware.NotImplemented("operation inspect.Inspect has not yet been implemented")
		}),
		CircuitListCircuitsHandler: circuit.ListCircuitsHandlerFunc(func(params circuit.ListCircuitsParams) middleware.Responder {
			return middleware.NotImplemented("operation circuit.ListCircuits has not yet been implemented")
		}),
		LinkListLinksHandler: link.ListLinksHandlerFunc(func(params link.ListLinksParams) middleware.Responder {
			return middleware.NotImplemented("operation link.ListLinks has not yet been implemented")
		}),
		RouterListRouterTerminatorsHandler: router.ListRouterTerminatorsHandlerFunc(func(params router.ListRouterTerminatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation router.ListRouterTerminators has not yet been implemented")
		}),
		RouterListRoutersHandler: router.ListRoutersHandlerFunc(func(params router.ListRoutersParams) middleware.Responder {
			return middleware.NotImplemented("operation router.ListRouters has not yet been implemented")
		}),
		ServiceListServiceTerminatorsHandler: service.ListServiceTerminatorsHandlerFunc(func(params service.ListServiceTerminatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ListServiceTerminators has not yet been implemented")
		}),
		ServiceListServicesHandler: service.ListServicesHandlerFunc(func(params service.ListServicesParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ListServices has not yet been implemented")
		}),
		TerminatorListTerminatorsHandler: terminator.ListTerminatorsHandlerFunc(func(params terminator.ListTerminatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.ListTerminators has not yet been implemented")
		}),
		LinkPatchLinkHandler: link.PatchLinkHandlerFunc(func(params link.PatchLinkParams) middleware.Responder {
			return middleware.NotImplemented("operation link.PatchLink has not yet been implemented")
		}),
		RouterPatchRouterHandler: router.PatchRouterHandlerFunc(func(params router.PatchRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.PatchRouter has not yet been implemented")
		}),
		ServicePatchServiceHandler: service.PatchServiceHandlerFunc(func(params service.PatchServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.PatchService has not yet been implemented")
		}),
		TerminatorPatchTerminatorHandler: terminator.PatchTerminatorHandlerFunc(func(params terminator.PatchTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.PatchTerminator has not yet been implemented")
		}),
		RouterUpdateRouterHandler: router.UpdateRouterHandlerFunc(func(params router.UpdateRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.UpdateRouter has not yet been implemented")
		}),
		ServiceUpdateServiceHandler: service.UpdateServiceHandlerFunc(func(params service.UpdateServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.UpdateService has not yet been implemented")
		}),
		TerminatorUpdateTerminatorHandler: terminator.UpdateTerminatorHandlerFunc(func(params terminator.UpdateTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.UpdateTerminator has not yet been implemented")
		}),
	}
}

/*ZitiFabricAPI the ziti fabric API */
type ZitiFabricAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// RouterCreateRouterHandler sets the operation handler for the create router operation
	RouterCreateRouterHandler router.CreateRouterHandler
	// ServiceCreateServiceHandler sets the operation handler for the create service operation
	ServiceCreateServiceHandler service.CreateServiceHandler
	// TerminatorCreateTerminatorHandler sets the operation handler for the create terminator operation
	TerminatorCreateTerminatorHandler terminator.CreateTerminatorHandler
	// CircuitDeleteCircuitHandler sets the operation handler for the delete circuit operation
	CircuitDeleteCircuitHandler circuit.DeleteCircuitHandler
	// RouterDeleteRouterHandler sets the operation handler for the delete router operation
	RouterDeleteRouterHandler router.DeleteRouterHandler
	// ServiceDeleteServiceHandler sets the operation handler for the delete service operation
	ServiceDeleteServiceHandler service.DeleteServiceHandler
	// TerminatorDeleteTerminatorHandler sets the operation handler for the delete terminator operation
	TerminatorDeleteTerminatorHandler terminator.DeleteTerminatorHandler
	// CircuitDetailCircuitHandler sets the operation handler for the detail circuit operation
	CircuitDetailCircuitHandler circuit.DetailCircuitHandler
	// LinkDetailLinkHandler sets the operation handler for the detail link operation
	LinkDetailLinkHandler link.DetailLinkHandler
	// RouterDetailRouterHandler sets the operation handler for the detail router operation
	RouterDetailRouterHandler router.DetailRouterHandler
	// ServiceDetailServiceHandler sets the operation handler for the detail service operation
	ServiceDetailServiceHandler service.DetailServiceHandler
	// TerminatorDetailTerminatorHandler sets the operation handler for the detail terminator operation
	TerminatorDetailTerminatorHandler terminator.DetailTerminatorHandler
	// InspectInspectHandler sets the operation handler for the inspect operation
	InspectInspectHandler inspect.InspectHandler
	// CircuitListCircuitsHandler sets the operation handler for the list circuits operation
	CircuitListCircuitsHandler circuit.ListCircuitsHandler
	// LinkListLinksHandler sets the operation handler for the list links operation
	LinkListLinksHandler link.ListLinksHandler
	// RouterListRouterTerminatorsHandler sets the operation handler for the list router terminators operation
	RouterListRouterTerminatorsHandler router.ListRouterTerminatorsHandler
	// RouterListRoutersHandler sets the operation handler for the list routers operation
	RouterListRoutersHandler router.ListRoutersHandler
	// ServiceListServiceTerminatorsHandler sets the operation handler for the list service terminators operation
	ServiceListServiceTerminatorsHandler service.ListServiceTerminatorsHandler
	// ServiceListServicesHandler sets the operation handler for the list services operation
	ServiceListServicesHandler service.ListServicesHandler
	// TerminatorListTerminatorsHandler sets the operation handler for the list terminators operation
	TerminatorListTerminatorsHandler terminator.ListTerminatorsHandler
	// LinkPatchLinkHandler sets the operation handler for the patch link operation
	LinkPatchLinkHandler link.PatchLinkHandler
	// RouterPatchRouterHandler sets the operation handler for the patch router operation
	RouterPatchRouterHandler router.PatchRouterHandler
	// ServicePatchServiceHandler sets the operation handler for the patch service operation
	ServicePatchServiceHandler service.PatchServiceHandler
	// TerminatorPatchTerminatorHandler sets the operation handler for the patch terminator operation
	TerminatorPatchTerminatorHandler terminator.PatchTerminatorHandler
	// RouterUpdateRouterHandler sets the operation handler for the update router operation
	RouterUpdateRouterHandler router.UpdateRouterHandler
	// ServiceUpdateServiceHandler sets the operation handler for the update service operation
	ServiceUpdateServiceHandler service.UpdateServiceHandler
	// TerminatorUpdateTerminatorHandler sets the operation handler for the update terminator operation
	TerminatorUpdateTerminatorHandler terminator.UpdateTerminatorHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ZitiFabricAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ZitiFabricAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ZitiFabricAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ZitiFabricAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ZitiFabricAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ZitiFabricAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ZitiFabricAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ZitiFabricAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ZitiFabricAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ZitiFabricAPI
func (o *ZitiFabricAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.RouterCreateRouterHandler == nil {
		unregistered = append(unregistered, "router.CreateRouterHandler")
	}
	if o.ServiceCreateServiceHandler == nil {
		unregistered = append(unregistered, "service.CreateServiceHandler")
	}
	if o.TerminatorCreateTerminatorHandler == nil {
		unregistered = append(unregistered, "terminator.CreateTerminatorHandler")
	}
	if o.CircuitDeleteCircuitHandler == nil {
		unregistered = append(unregistered, "circuit.DeleteCircuitHandler")
	}
	if o.RouterDeleteRouterHandler == nil {
		unregistered = append(unregistered, "router.DeleteRouterHandler")
	}
	if o.ServiceDeleteServiceHandler == nil {
		unregistered = append(unregistered, "service.DeleteServiceHandler")
	}
	if o.TerminatorDeleteTerminatorHandler == nil {
		unregistered = append(unregistered, "terminator.DeleteTerminatorHandler")
	}
	if o.CircuitDetailCircuitHandler == nil {
		unregistered = append(unregistered, "circuit.DetailCircuitHandler")
	}
	if o.LinkDetailLinkHandler == nil {
		unregistered = append(unregistered, "link.DetailLinkHandler")
	}
	if o.RouterDetailRouterHandler == nil {
		unregistered = append(unregistered, "router.DetailRouterHandler")
	}
	if o.ServiceDetailServiceHandler == nil {
		unregistered = append(unregistered, "service.DetailServiceHandler")
	}
	if o.TerminatorDetailTerminatorHandler == nil {
		unregistered = append(unregistered, "terminator.DetailTerminatorHandler")
	}
	if o.InspectInspectHandler == nil {
		unregistered = append(unregistered, "inspect.InspectHandler")
	}
	if o.CircuitListCircuitsHandler == nil {
		unregistered = append(unregistered, "circuit.ListCircuitsHandler")
	}
	if o.LinkListLinksHandler == nil {
		unregistered = append(unregistered, "link.ListLinksHandler")
	}
	if o.RouterListRouterTerminatorsHandler == nil {
		unregistered = append(unregistered, "router.ListRouterTerminatorsHandler")
	}
	if o.RouterListRoutersHandler == nil {
		unregistered = append(unregistered, "router.ListRoutersHandler")
	}
	if o.ServiceListServiceTerminatorsHandler == nil {
		unregistered = append(unregistered, "service.ListServiceTerminatorsHandler")
	}
	if o.ServiceListServicesHandler == nil {
		unregistered = append(unregistered, "service.ListServicesHandler")
	}
	if o.TerminatorListTerminatorsHandler == nil {
		unregistered = append(unregistered, "terminator.ListTerminatorsHandler")
	}
	if o.LinkPatchLinkHandler == nil {
		unregistered = append(unregistered, "link.PatchLinkHandler")
	}
	if o.RouterPatchRouterHandler == nil {
		unregistered = append(unregistered, "router.PatchRouterHandler")
	}
	if o.ServicePatchServiceHandler == nil {
		unregistered = append(unregistered, "service.PatchServiceHandler")
	}
	if o.TerminatorPatchTerminatorHandler == nil {
		unregistered = append(unregistered, "terminator.PatchTerminatorHandler")
	}
	if o.RouterUpdateRouterHandler == nil {
		unregistered = append(unregistered, "router.UpdateRouterHandler")
	}
	if o.ServiceUpdateServiceHandler == nil {
		unregistered = append(unregistered, "service.UpdateServiceHandler")
	}
	if o.TerminatorUpdateTerminatorHandler == nil {
		unregistered = append(unregistered, "terminator.UpdateTerminatorHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ZitiFabricAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ZitiFabricAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *ZitiFabricAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ZitiFabricAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ZitiFabricAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ZitiFabricAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the ziti fabric API
func (o *ZitiFabricAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ZitiFabricAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/routers"] = router.NewCreateRouter(o.context, o.RouterCreateRouterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/services"] = service.NewCreateService(o.context, o.ServiceCreateServiceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/terminators"] = terminator.NewCreateTerminator(o.context, o.TerminatorCreateTerminatorHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/circuits/{id}"] = circuit.NewDeleteCircuit(o.context, o.CircuitDeleteCircuitHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/routers/{id}"] = router.NewDeleteRouter(o.context, o.RouterDeleteRouterHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/services/{id}"] = service.NewDeleteService(o.context, o.ServiceDeleteServiceHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/terminators/{id}"] = terminator.NewDeleteTerminator(o.context, o.TerminatorDeleteTerminatorHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/circuits/{id}"] = circuit.NewDetailCircuit(o.context, o.CircuitDetailCircuitHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/links/{id}"] = link.NewDetailLink(o.context, o.LinkDetailLinkHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/routers/{id}"] = router.NewDetailRouter(o.context, o.RouterDetailRouterHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{id}"] = service.NewDetailService(o.context, o.ServiceDetailServiceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/terminators/{id}"] = terminator.NewDetailTerminator(o.context, o.TerminatorDetailTerminatorHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/inspections"] = inspect.NewInspect(o.context, o.InspectInspectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/circuits"] = circuit.NewListCircuits(o.context, o.CircuitListCircuitsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/links"] = link.NewListLinks(o.context, o.LinkListLinksHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/routers/{id}/terminators"] = router.NewListRouterTerminators(o.context, o.RouterListRouterTerminatorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/routers"] = router.NewListRouters(o.context, o.RouterListRoutersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{id}/terminators"] = service.NewListServiceTerminators(o.context, o.ServiceListServiceTerminatorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services"] = service.NewListServices(o.context, o.ServiceListServicesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/terminators"] = terminator.NewListTerminators(o.context, o.TerminatorListTerminatorsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/links/{id}"] = link.NewPatchLink(o.context, o.LinkPatchLinkHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/routers/{id}"] = router.NewPatchRouter(o.context, o.RouterPatchRouterHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/services/{id}"] = service.NewPatchService(o.context, o.ServicePatchServiceHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/terminators/{id}"] = terminator.NewPatchTerminator(o.context, o.TerminatorPatchTerminatorHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/routers/{id}"] = router.NewUpdateRouter(o.context, o.RouterUpdateRouterHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/services/{id}"] = service.NewUpdateService(o.context, o.ServiceUpdateServiceHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/terminators/{id}"] = terminator.NewUpdateTerminator(o.context, o.TerminatorUpdateTerminatorHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ZitiFabricAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ZitiFabricAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ZitiFabricAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ZitiFabricAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ZitiFabricAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
