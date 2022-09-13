package User

import (
	"sample/mocktest/IUser"
)

type User struct {
	IUser IUser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "sample test ")
	//fmt.Printf("u.IUser: %v\n", u.IUser)
}
