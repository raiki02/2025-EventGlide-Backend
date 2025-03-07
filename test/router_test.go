package test

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestRegisterRouters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
