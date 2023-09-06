package processes

import (
	"IM/common/message"
	"IM/server/model"
	"IM/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (up *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, u := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		u.NotifyMeOnline(userId)
	}
}

func (up *UserProcess) NotifyMeOnline(userId int) {

	//Build NotifyMes type struct

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshl(notifyUserStatusMes) failed! Error:", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(ms) failed! Error:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err=", err)
		return
	}
}

func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&loginMes)failed! Error:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 400
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "Server internal error!"
			fmt.Println("Server intenal error!")
		}

	} else {
		loginResMes.Code = 200
		up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(up)
		up.NotifyOthersOnlineUser(loginMes.UserId)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user.UserName, "login Successful!")
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) failed! Error:", err)
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("jon.Marshal(resMes) failed! Error:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err=", err)
	}

	return
}

func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data),&registerMes) failed! Error:", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Signup(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 500
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "Server internal error!"
			fmt.Println("Server internal error!")
		}

	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.User.UserName, "sign up successful!")
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) failed! Error:", err)
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) failed! Error:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err=", err)
	}

	return
}
