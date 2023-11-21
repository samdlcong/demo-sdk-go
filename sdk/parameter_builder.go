package sdk

import (
	"encoding/json"
	"github.com/samdlcong/demo-sdk-go/sdk/log"
	"github.com/samdlcong/demo-sdk-go/sdk/request"
	"reflect"
)

var baseRequestFileds []string

func init() {
	req := request.BaseRequest{}
	reqType := reflect.TypeOf(req)
	for i := 0; i < reqType.NumField(); i++ {
		baseRequestFileds = append(baseRequestFileds, reqType.Field(i).Name)
	}
}

type ParameterBuilder interface {
	BuildURL(url string, paramJson []byte) (string, error)
	BuildBody(paramJson []byte) (string, error)
}

func GetParameterBuilder(method string, logger log.Logger) ParameterBuilder {
	if method == MethodGet || method == MethodDelete || method == MethodHead {
		return &WithoutBodyBuilder{logger}
	} else {
		return &WithBodyBuilder{logger}
	}
}

// WithBodyBuilder supports PUT/POST/PATCH methods.
// It has path and body (json) parameters, but no query parameters.
type WithBodyBuilder struct {
	Logger log.Logger
}

func (b WithBodyBuilder) BuildURL(url string, paramJson []byte) (string, error) {
	paramMap := make(map[string]interface{})
	err := json.Unmarshal(paramJson, &paramMap)
	if err != nil {
		b.Logger.Errorf("%s", err.Error())
	}
}

func (b WithBodyBuilder) BuildBody(paramJson []byte) (string, error) {

}

type WithoutBodyBuilder struct {
	Logger log.Logger
}

func (b WithoutBodyBuilder) BuildURL(url string, paramJson []byte) (string, error) {

}

func (b WithoutBodyBuilder) BuildBody(paramJson []byte) (string, error) {

}
