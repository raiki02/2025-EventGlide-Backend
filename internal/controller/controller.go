package controller

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewActController,
	NewCommentController,
	NewFeedController,
	NewInteractionController,
	NewPostController,
	NewUserController,
	NewCallbackAuditorController,
)
