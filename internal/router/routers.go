package router

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewActRouter,
	NewCommentRouter,
	NewFeedRouter,
	NewInteractionRouter,
	NewPostRouter,
	NewUserRouter,
	NewRouter,
)
