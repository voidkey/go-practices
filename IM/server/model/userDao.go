package model

import (
	"IM/common/message"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	rdb *redis.Client
	ctx context.Context
}

func NewUserDao(rdb *redis.Client, ctx context.Context) (userDao *UserDao) {
	userDao = &UserDao{
		rdb: rdb,
		ctx: ctx,
	}
	return
}

func (ud *UserDao) getUserById(id int) (user *User, err error) {
	//res,err:= redis.String(conn.Do("HGet","users",id))
	val, err := ud.rdb.HGet(ud.ctx, "users", strconv.Itoa(id)).Result()
	if err == redis.Nil {
		err = ERROR_USER_NOTEXISTS
		return
	} else if err != nil {
		fmt.Println("HGet failed! Error:", err)
		return
	}
	//user = &User{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(val),&user) failed! Error:", err)
		return
	}
	return
}

func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	user, err = ud.getUserById(userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (ud *UserDao) Signup(user *message.User) (err error) {
	_, err = ud.getUserById(user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	if err == ERROR_USER_NOTEXISTS {
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println("json.Marshal(user) failed! Error:", err)
			return err
		}
		_, err = ud.rdb.HSet(ud.ctx, "users", user.UserId, string(data)).Result()
		if err != nil {
			fmt.Println("HSet(ud.ctx, \"users\", data) failed! Error:", err)
			return err
		} else {
			return nil
		}
	}
	return
}
