package errors

import "fmt"

type DEMOSDKError struct {
	Code      int
	Message   string
	RequestID string
}

func (err *DEMOSDKError) Error() string {
	return fmt.Sprintf("[DEMOSDKError] code=%d, message=%s, requestID=%s", err.Code, err.Message, err.RequestID)
}

func NewDEMOSDKError(code int, message, requestID string) error {
	return &DEMOSDKError{
		Code:      code,
		Message:   message,
		RequestID: requestID,
	}
}

func (err *DEMOSDKError) GetCode() int {
	return err.Code
}

func (err *DEMOSDKError) GetMessage() string {
	return err.Message
}

func (err *DEMOSDKError) GetRequestID() string {
	return err.RequestID
}
