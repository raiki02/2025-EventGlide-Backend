package middleware

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewJwt,
	NewCors,
)
