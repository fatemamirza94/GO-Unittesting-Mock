# GO-Unittesting-Mock

Mocking is a process used in unit testing when the unit being tested has external dependencies. The purpose of mocking is to isolate and focus on the code being tested and not on the behavior or state of external dependencies.

##############################################################################

1. In VSCode: go mod init sample/mocktest
2. Create IUser and User folders
3. Create IUser.go and User.go 
4. Download gomock and mockgen using terminal in VSCode:
   a) go get github.com/golang/mock/gomock
   b) go get github.com/golang/mock/mockgen
5. Create User_test.go (the part that is commented)
6. Run in CMD at root directory of the project (this only works on CMD and not on powershell or VSCode Terminal - I am sure there is a workaround but I am too lazy to figure it out): mockgen -destination=mocks/mockRunner.go  -package=mocks sample/mocktest/IUser IUserInterface
7. Complete the code for User_test.go (as shown in the non commented part)
8. In cmd terminal: go test -v sample/mocktest/User
9. To clear test caching: go clean -testcache
