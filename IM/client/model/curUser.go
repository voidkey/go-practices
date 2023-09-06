package model

import (
	"IM/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
