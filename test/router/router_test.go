package tests

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestRegisterRouters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockRouterHdl(ctrl)
	mock.EXPECT().RegisterRouters().Times(1)

	mock.RegisterRouters()
	mock.EXPECT().Run().Times(1)
	mock.Run()
}
