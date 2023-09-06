package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"proto_demo/proto/userService"
)

func main() {

	u := &userService.Userinfo{
		Username: "Hutao",
		Age:      16,
		Hobby:    []string{"Bury", "Funeral", "Burn"},
	}
	fmt.Println(u)
	fmt.Println(u.GetUsername())
	fmt.Println(u.GetAge())

	data, _ := proto.Marshal(u)
	fmt.Println(data)
	user := userService.Userinfo{}
	proto.Unmarshal(data, &user)
	fmt.Println(user)
	fmt.Printf("%#v\n", user)
	fmt.Println(user.GetHobby())
}
