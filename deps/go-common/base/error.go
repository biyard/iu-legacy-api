package base

import "fmt"

type Error struct {
	Code        int    `json:"code"`
	StatusCode  int    `json:"-"`
	Message     string `json:"message"`
	shouldRetry bool
}

func NewError(code int, statusCode int, message string) *Error {
	ret := &Error{
		Code:       code,
		StatusCode: statusCode,
		Message:    message,
	}

	return ret
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) WithMessage(message string) *Error {
	ret := &Error{
		Code:       e.Code,
		StatusCode: e.StatusCode,
		Message:    fmt.Sprintf("%s - %s", e.Message, message),
	}

	return ret
}

func (e *Error) WithCode(code int) *Error {
	ret := &Error{
		Code:       code,
		StatusCode: e.StatusCode,
		Message:    e.Message,
	}

	return ret
}

func (e *Error) WithStatusCode(statusCode int) Error {
	ret := Error{
		Code:       e.Code,
		StatusCode: statusCode,
		Message:    e.Message,
	}

	return ret
}

func (e *Error) WithRetry() *Error {
	e.shouldRetry = true
	return e
}

func (e *Error) ShouldRetry() bool {
	return e.shouldRetry
}

func (e *Error) Equal(err *Error) bool {
	return e.Code == err.Code
}

var (
	ErrNotImplemented = NewError(0, 501, "not implemented")
	ErrNotFound       = NewError(1, 404, "not found")
	ErrRequestParsing = NewError(2, 400, "request parsing error")
	ErrUnknown        = NewError(3, 500, "unknown error")

	// 10 - 50: common errors
	ErrInvalidLambdaType = NewError(10, 500, "invalid lambda type")

	// 100 - 150: klaytn util errors
	ErrKlaytnRLPUnsupported      = NewError(100, 400, "unsupported rlp type; legacy transaction does not support rlp encoding")
	ErrKlaytnFilteredTransaction = NewError(101, 400, "filtered transaction")
	ErrKlaytnFailedToSignTx      = NewError(102, 500, "failed to sign transaction")
	ErrKlaytnFailedToEncodeTx    = NewError(103, 500, "failed to encode transaction")
	ErrKlaytnFailedToDecodeTx    = NewError(104, 400, "failed to decode RLP")
	ErrKlaytnFailedToSendTx      = NewError(105, 500, "failed to send transaction")
	ErrKlaytnFailedToGetNonce    = NewError(106, 500, "failed to get nonce")
	ErrKlaytnTxRejected          = NewError(107, 500, "transaction reverted")
	ErrKlaytnFailedToCall        = NewError(108, 500, "failed to call data")

	// 200 - 250: kakao errors
	ErrKakaoInvalieCode = NewError(200, 400, "invalid code")
	ErrKakaoTokenInfo   = NewError(201, 500, "failed to get token info")

	// 300 - 350: firebase errors
	ErrFirebaseAuthClient    = NewError(300, 500, "failed to get auth client")
	ErrFirebaseVerifyIDToken = NewError(301, 500, "failed to verify ID token")

	// 400 - 450: database errors
	ErrDatabaseUnknown = NewError(400, 500, "database error")
	ErrDataNotFound    = NewError(401, 400, "data not found")
	ErrUnsupportedType = NewError(402, 500, "unsupported type")
	ErrDataParsing     = NewError(403, 500, "data parsing error")

	// 500 - 550: klaytn errors
	ErrKlaytnFailedToParsePrivateKey = NewError(500, 500, "failed to parse private key")

	// 600 - 650: cache errors
	ErrRedisConfigParse = NewError(600, 500, "failed to parse redis config")
	ErrCachePush        = NewError(601, 500, "failed to push cache")
	ErrCacheDelete      = NewError(602, 500, "failed to delete cache")
	ErrCacheGet         = NewError(603, 500, "failed to get cache")
	ErrCacheMarshal     = NewError(604, 500, "failed to marshal cache")
	ErrCacheUnmarshal   = NewError(605, 500, "failed to unmarshal cache")
)
