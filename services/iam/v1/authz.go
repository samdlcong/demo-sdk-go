package v1

import (
	"encoding/json"
	"github.com/ory/ladon"
	"github.com/samdlcong/demo-sdk-go/sdk/request"
	"github.com/samdlcong/demo-sdk-go/sdk/response"
)

type AuthzRequest struct {
	*request.BaseRequest
	Resource *string `json:"resource"`

	Action *string `json:"action"`

	Subject *string `json:"subject"`
	Context *ladon.Context
}

type AuthzResponse struct {
	*response.BaseResponse
	Allowed bool   `json:"allowed,omitempty"`
	Denied  bool   `json:"denied,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewAuthzRequest() (req *AuthzRequest) {
	req = &AuthzRequest{
		BaseRequest: &request.BaseRequest{
			URL:     "/authz",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
	return
}

func NewAuthzResponse() *AuthzResponse {
	return &AuthzResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

func (r *AuthzResponse) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (c *Client) Authz(req *AuthzRequest) (resp *AuthzResponse, err error) {
	if req == nil {
		req = NewAuthzRequest()
	}

	resp = NewAuthzResponse()
	err = c.Send(req, resp)
	return
}
