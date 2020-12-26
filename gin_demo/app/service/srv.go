package service

var (
	UserSrvCli *UserService
)

func init() {
	UserSrvCli = NewUserService()
}
