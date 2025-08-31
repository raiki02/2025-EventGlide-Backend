package dao

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewActDao,
	NewAuditorRepo,
	NewCommentDao,
	NewFeedDao,
	NewInteractionDao,
	NewPostDao,
	NewUserDao,
)
