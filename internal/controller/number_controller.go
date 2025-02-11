package controller

type NumberControllerHdl interface {
	HandleLikesNum()
	HandleCommentsNum()
}

type NumberController struct {
}
