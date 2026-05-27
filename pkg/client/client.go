package client

import (
	v1 "github.com/resonatehq/resonate/pkg/client/v1"
)

type Client interface {
	V1() v1.ClientWithResponsesInterface
	Setup(string) error
	SetBasicAuth(string, string)
	SetBearerToken(string)
}

// Client

type client struct {
	v1       v1.ClientWithResponsesInterface
	username string
	password string
	token    string
}

func New() Client { _ = "STUB: not implemented"; return *new(Client) }

func (c *client) Setup(server string) error { _ = "STUB: not implemented"; return nil }

func (c *client) V1() v1.ClientWithResponsesInterface {
	_ = "STUB: not implemented"
	return *new(v1.ClientWithResponsesInterface)
}

func (c *client) SetBasicAuth(username, password string) { _ = "STUB: not implemented"; return }

func (c *client) SetBearerToken(token string) {
	_ = "STUB: not implemented"

	// Mock Client
	return
}

type mockClient struct {
	v1 *v1.MockClientWithResponsesInterface
}

func MockClient(v1 *v1.MockClientWithResponsesInterface) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func (c *mockClient) Setup(string) error { _ = "STUB: not implemented"; return nil }

func (c *mockClient) V1() v1.ClientWithResponsesInterface {
	_ = "STUB: not implemented"
	return *new(v1.ClientWithResponsesInterface)
}

func (c *mockClient) SetBasicAuth(string, string) { _ = "STUB: not implemented"; return }

func (c *mockClient) SetBearerToken(string) {
	_ = "STUB: not implemented"

	// Helper functions
	return
}

func basicAuth(username, password string) v1.ClientOption {
	_ = "STUB: not implemented"
	return *new(v1.ClientOption)
}

func bearerToken(token string) v1.ClientOption {
	_ = "STUB: not implemented"
	return *new(v1.ClientOption)
}
