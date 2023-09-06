package processes

import (
	"IM/common/message"
	"encoding/json"
	"fmt"
)

func showGroupMes(mes *message.Message) {
	var smsMes message.SmsMes

	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("unmarshal() failed! Error:", err)
		return
	}
	info := fmt.Sprintf("UserId %d said: %s\n", smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
}
