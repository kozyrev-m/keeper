package httpserver

type ctxKey int8

const (
	sessionName = "session"
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)