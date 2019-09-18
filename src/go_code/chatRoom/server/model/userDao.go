package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatRoom/common/message"

	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后,就初始化一个userDao的实例,把他做成全局的,在需要redis操作时,就可以直接使用

var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体,完成对User结构体的各种操作
type UserDao struct {
	Pool *redis.Pool
}

// 使用工厂模式,创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

// 1. 根据用户id返回一个user实例加error
func (u *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	// 通过给的id,到redis查询
	ret, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			// 说明没有找到此用户
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// 我们需要反序列化变成User实例
	err = json.Unmarshal([]byte(ret), user)
	if err != nil {
		fmt.Println("反序列化用户失败", err)
		return
	}
	return
}

//登录校验
func (u *UserDao) LoginCheck(userId int, userPwd string) (user *User, err error) {
	// 1. 完成用户的校验
	// 2. 如果用户的id和pwd都正确,则返回一个user实例
	// 3. 如果用户的id和pwd有错误,则返回对应的错误信息
	// 从连接池中取出一条连接
	conn := u.Pool.Get()
	defer conn.Close()
	user, err = u.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 此时的user是一个反序列化的结构体
	// 取出密码检验,虽然用户id是对的,但是需要校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 注册校验
func (u *UserDao) RegisterCheck(user message.User) (err error) {

	conn := u.Pool.Get()
	defer conn.Close()
	// 到redis中去去相应的数据,如果有说明此用户已被
	_, err = u.getUserById(conn, user.UserId)
	if err == nil {
		// 说明用户已经存在了
		err = ERROR_USER_EXISTS
		return
	}

	// 如果有错误,说明id在redis还没有注册,就可以完成注册
	data, err := json.Marshal(user) // 序列化然后转成字符串,存入redis中
	if err != nil {
		fmt.Println("序列化注册信息失败", err)
		return
	}

	// 入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("到redis中保存注册用户错误", err)
		return
	}
	return
}
