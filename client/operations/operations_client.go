/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WeaviateBatchingActionsCreate creates new actions based on an action template as a batch

Register new Actions in bulk. Given meta-data and schema values are validated.
*/
func (a *Client) WeaviateBatchingActionsCreate(params *WeaviateBatchingActionsCreateParams) (*WeaviateBatchingActionsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateBatchingActionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.batching.actions.create",
		Method:             "POST",
		PathPattern:        "/batching/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateBatchingActionsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateBatchingActionsCreateOK), nil

}

/*
WeaviateBatchingReferencesCreate creates new cross references between arbitrary classes in bulk

Register cross-references between any class items (things or actions) in bulk.
*/
func (a *Client) WeaviateBatchingReferencesCreate(params *WeaviateBatchingReferencesCreateParams) (*WeaviateBatchingReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateBatchingReferencesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.batching.references.create",
		Method:             "POST",
		PathPattern:        "/batching/references",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateBatchingReferencesCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateBatchingReferencesCreateOK), nil

}

/*
WeaviateBatchingThingsCreate creates new things based on a thing template as a batch

Register new Things in bulk. Provided meta-data and schema values are validated.
*/
func (a *Client) WeaviateBatchingThingsCreate(params *WeaviateBatchingThingsCreateParams) (*WeaviateBatchingThingsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateBatchingThingsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.batching.things.create",
		Method:             "POST",
		PathPattern:        "/batching/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateBatchingThingsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateBatchingThingsCreateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
