package processes

import (
	"IM/common/message"
	"IM/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("unmarshal() failed! Error:", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal() failed! Error:", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.User.UserId {
			continue
		}
		sp.SemdMesEachOnlineUser(data, up.Conn)
	}

}

func (sp *SmsProcess) SemdMesEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesEachOnlineUser WritePkg() failed! Error:", err)
		return
	}
}
