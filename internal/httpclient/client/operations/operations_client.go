// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AdminIntrospectOAuth2Token(params *AdminIntrospectOAuth2TokenParams, opts ...ClientOption) (*AdminIntrospectOAuth2TokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AdminIntrospectOAuth2Token introspects o auth2 access or refresh tokens

  The introspection endpoint allows to check if a token (both refresh and access) is active or not. An active token
is neither expired nor revoked. If a token is active, additional information on the token will be included. You can
set additional data for a token by setting `accessTokenExtra` during the consent flow.

For more information [read this blog post](https://www.oauth.com/oauth2-servers/token-introspection-endpoint/).
*/
func (a *Client) AdminIntrospectOAuth2Token(params *AdminIntrospectOAuth2TokenParams, opts ...ClientOption) (*AdminIntrospectOAuth2TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminIntrospectOAuth2TokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "adminIntrospectOAuth2Token",
		Method:             "POST",
		PathPattern:        "/admin/oauth2/introspect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdminIntrospectOAuth2TokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AdminIntrospectOAuth2TokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AdminIntrospectOAuth2TokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
