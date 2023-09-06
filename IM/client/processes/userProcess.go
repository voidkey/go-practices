package processes

import (
	"IM/client/utils"
	"IM/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {

	//1.Connect to Sever
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Dial failed! Error:", err)
		return err
	}
	defer conn.Close()
	//2.Prepare to send message to sever by conn
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.Create a LoginMes struct
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) failed! Error:", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) failed! Error:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	//Send the lenth of data firstly for Security
	// conn.Write(len(data)) is wrong
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Login WritePkg() failed! Error:", err)
		return
	}

	//Receive Response Message
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Login ReadPkg() err:", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal(mes,&loginResMes) failed! Error:", err)
		return
	}
	if loginResMes.Code == 200 {

		CurUser.Conn = conn
		CurUser.User.UserId = userId
		CurUser.User.UserStatus = message.UserOnline

		fmt.Println("Current Online members:")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("UserId: ", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		go ProccessSerMes(conn)
		ShowMenu()

	} else {
		fmt.Println(loginResMes.Error)
	}
	return err
}

func (up *UserProcess) Signup(userId int, userPwd string, userName string) (err error) {
	//1.Connect to Sever
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Dial failed! Error:", err)
		return err
	}
	defer conn.Close()
	//2.Prepare to send message to sever by conn
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.Create a RegisterMes struct
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) failed! Error:", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) failed! Error:", err)
		return
	}

	//Send Message
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Sign up WritePkg() failed! Error:", err)
		return
	}

	//Receive Response Message
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Sign up readPkg() err:", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal(mes,&registerResMes) failed! Error:", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("Sign up successful! Please log in again!")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return err
}
