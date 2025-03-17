package errUtil

import "errors"

const (
	BAD_REQUEST            = "BAD_REQUEST"            // 400
	UNAUTHORIZED           = "UNAUTHORIZED"           // 401
	FORBIDDEN              = "FORBIDDEN"              // 403
	NOT_FOUND              = "NOT_FOUND"              // 404
	PAYLOAD_TOO_LARGE      = "PAYLOAD_TOO_LARGE"      // 413
	UNSUPPORTED_MEDIA_TYPE = "UNSUPPORTED_MEDIA_TYPE" // 415
	INTERNAL_SERVER_ERROR  = "INTERNAL_SERVER_ERROR"  // 500
)

func generateMessage(code int) string {
	switch code {
	case 400:
		return BAD_REQUEST
	case 401:
		return UNAUTHORIZED
	case 403:
		return FORBIDDEN
	case 404:
		return NOT_FOUND
	case 413:
		return PAYLOAD_TOO_LARGE
	case 415:
		return UNSUPPORTED_MEDIA_TYPE
	case 500:
		return INTERNAL_SERVER_ERROR
	}
	return INTERNAL_SERVER_ERROR
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"0" format:"int"`
	Message string `json:"message" example:"success" format:"string"`
}

type applicationError struct {
	error
	Code int
}

func MakeBaseResponse(code int) ErrorResponse {
	res := ErrorResponse{}
	res.Code = code
	res.Message = generateMessage(res.Code) // general message 생성
	return res
}

func WrapWithCode(err error, code int) error {
	return &applicationError{wrap(err, ""), code}
}

func CastApplicationError(err error) (*applicationError, bool) {
	for err != nil {
		switch err.(type) {
		case *applicationError:
			return err.(*applicationError), true
		}
		err = errors.Unwrap(err)
	}
	return nil, false
}
