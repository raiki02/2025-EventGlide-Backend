package dao

import "context"

type PostDaoHdl interface {
	GetPostList(context.Context) error
	InsertPost(context.Context, string, string, int) error
}
