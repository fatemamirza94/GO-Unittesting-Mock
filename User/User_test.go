package User_test

import (
	"sample/mocktest/User"
	"sample/mocktest/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser := &User.User{IUser: mockUser}
	mockUser.EXPECT().AddUser(1, "sample test ").Return(nil).Times(1)
	testUser.Use()
}
