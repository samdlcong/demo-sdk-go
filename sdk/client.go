package sdk

import (
	"fmt"
	"github.com/samdlcong/demo-sdk-go/sdk/log"
	"net/http"
	"runtime"
	"strings"
)

var defaultUserAgent = fmt.Sprintf("DEMOSDKGo/%s (%s; %s) Golang/%s", Version, runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"))

// Client is the base struct of service clients
type Client struct {
	signMethod  string
	Credential  *Credential
	Config      *Config
	ServiceName string
	Logger      log.Logger
}

type SignFunc func(r *http.Request) error

// ListOptions specifies the optional patameters to various List methods that
// support pagination.
type ListOptions struct {
	Offset *int64 `json:"offset,omitempty"`
	Limit  *int64 `json:"limit,omitempty"`
}

func (c *Client) Init(serviceName string) *Client {
	c.signMethod = "jwt"
	c.Logger = log.New()
	c.ServiceName = serviceName
	return c
}

func (c *Client) WithCredential(cred *Credential) *Client {
	c.Credential = cred
	return c
}

func (c *Client) WithSecret(secretID, secretKey string) *Client {
	c.Credential = NewCredentials(secretID, secretKey)
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	c.Logger.SetLevel(config.LogLevel)
	return c
}
