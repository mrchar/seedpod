package client

import "fmt"

// ErrorServerReported 服务器报告的错误
type ErrorServerReported struct {
	statusCode int
	message    string
}

func (e ErrorServerReported) Error() string {
	return fmt.Sprintf("StatusCode: %d, Description: %s", e.statusCode, e.message)
}
