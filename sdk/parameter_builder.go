package sdk

import (
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

type Parameter 