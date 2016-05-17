package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/deis/workflow-manager/pkg/swagger/client/operations"
)

// Default workflow manager HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new workflow manager HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WorkflowManager {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new workflow manager client
func New(transport client.Transport, formats strfmt.Registry) *WorkflowManager {
	cli := new(WorkflowManager)
	cli.Transport = transport

	cli.Operations = operations.New(transport, formats)

	return cli
}

// WorkflowManager is a client for workflow manager
type WorkflowManager struct {
	Operations *operations.Client

	Transport client.Transport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WorkflowManager) SetTransport(transport client.Transport) {
	c.Transport = transport

	c.Operations.SetTransport(transport)

}
