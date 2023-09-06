package processes

import (
	"IM/client/utils"
	"IM/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	smsProcess := &SmsProcess{}
	for {
		fmt.Println("------------------Congratulations on your successful login------------------")
		fmt.Println("\t\t\t1.Show online numbers")
		fmt.Println("\t\t\t2.Send message")
		fmt.Println("\t\t\t3.Message list")
		fmt.Println("\t\t\t4.Quit System")
		fmt.Println("Please input your option(1~4):")
		var key int
		fmt.Scanf("%d\n", &key)
		var content string
		switch key {
		case 1:
			showOnlineUser()
		case 2:
			fmt.Println("Please input your message")
			fmt.Scanf("%s\n", &content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("test3")
		case 4:
			fmt.Println("You choose to quit system")
			os.Exit(0)
		default:
			fmt.Println("Unexpected option, please input correctly!")
		}
		fmt.Println("------------------------------------------------------")
	}
}

func ProccessSerMes(conn net.Conn) (err error) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("Client is waiting for the meesage sent by Server\n")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() failed! Error:", err)
			return err
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			fmt.Println("Notify USER!")
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			showGroupMes(&mes)
		default:
			fmt.Println("Unknown Error!")
		}
	}

}
