package processes

import (
	"IM/client/model"
	"IM/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //After user log in successful, initial the CurUser

func showOnlineUser() {
	fmt.Println("Current Online Members:")
	for id, _ := range onlineUsers {
		fmt.Println("UserId: ", id)
	}
}

//Deal with NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	showOnlineUser()
}
