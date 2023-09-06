package processes

import (
	"IM/client/utils"
	"IM/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User.UserId = CurUser.User.UserId
	smsMes.User.UserStatus = CurUser.User.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) failed! Error:", err)
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
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg() failed! Error:", err)
		return
	}
	return
}
