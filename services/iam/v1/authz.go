package v1

import (
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

}
