package processes

import "fmt"

//Because UserManager only have one object in Server, define it as Global Variable
var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers
}

func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[up.UserId]
	if !ok {
		err = fmt.Errorf("User%d does not exit!", userId)
		return
	}
	return
}
