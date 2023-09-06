package main

import (
	"IM/common/message"
	"IM/server/processes"
	"IM/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//According to the type of messages, choose to call the correct function
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		sp := &processes.SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("Message Type does not exist!")
	}

	return
}

func (p *Processor) Dispatcher() (err error) {
	for {

		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client Communication is over!")
				return err
			} else {
				fmt.Println("readPkg() err:", err)
				return err
			}
		}
		err = p.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
