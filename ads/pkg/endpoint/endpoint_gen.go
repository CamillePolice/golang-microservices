// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "golangmicroservices/ads/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateEndpoint          endpoint.Endpoint
	UpdateEndpoint          endpoint.Endpoint
	DeleteEndpoint          endpoint.Endpoint
	GetEndpoint             endpoint.Endpoint
	GetAllByKeyWordEndpoint endpoint.Endpoint
	GetAllByUserEndpoint    endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.AdsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateEndpoint:          MakeCreateEndpoint(s),
		DeleteEndpoint:          MakeDeleteEndpoint(s),
		GetAllByKeyWordEndpoint: MakeGetAllByKeyWordEndpoint(s),
		GetAllByUserEndpoint:    MakeGetAllByUserEndpoint(s),
		GetEndpoint:             MakeGetEndpoint(s),
		UpdateEndpoint:          MakeUpdateEndpoint(s),
	}
	for _, m := range mdw["Create"] {
		eps.CreateEndpoint = m(eps.CreateEndpoint)
	}
	for _, m := range mdw["Update"] {
		eps.UpdateEndpoint = m(eps.UpdateEndpoint)
	}
	for _, m := range mdw["Delete"] {
		eps.DeleteEndpoint = m(eps.DeleteEndpoint)
	}
	for _, m := range mdw["Get"] {
		eps.GetEndpoint = m(eps.GetEndpoint)
	}
	for _, m := range mdw["GetAllByKeyWord"] {
		eps.GetAllByKeyWordEndpoint = m(eps.GetAllByKeyWordEndpoint)
	}
	for _, m := range mdw["GetAllByUser"] {
		eps.GetAllByUserEndpoint = m(eps.GetAllByUserEndpoint)
	}
	return eps
}
