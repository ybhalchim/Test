/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"github.com/infobloxopen/bloxone-go-client/internal"
	"github.com/infobloxopen/bloxone-go-client/option"
)

const serviceBasePath = "/api/ddi/v1"

// APIClient manages communication with the DDI Keys API v1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	GenerateTsigAPI GenerateTsigAPI
	KerberosAPI     KerberosAPI
	TsigAPI         TsigAPI
	UploadAPI       UploadAPI
}

// NewAPIClient creates a new API client.
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithCSPUrl(string) sets the URL for BloxOne Cloud Services Portal.
// - WithAPIKey(string) sets the APIKey for accessing the BloxOne API.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultTags(map[string]string) sets the tags the client can set by default for objects that has tags support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	c := &APIClient{}
	c.APIClient = internal.NewAPIClient(serviceBasePath, cfg)

	// API Services
	c.GenerateTsigAPI = (*GenerateTsigAPIService)(&c.Common)
	c.KerberosAPI = (*KerberosAPIService)(&c.Common)
	c.TsigAPI = (*TsigAPIService)(&c.Common)
	c.UploadAPI = (*UploadAPIService)(&c.Common)

	return c
}