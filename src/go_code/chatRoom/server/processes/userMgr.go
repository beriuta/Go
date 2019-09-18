package processes

import (
	"fmt"
)

// 因为UserMgr实例在服务器有且只有一个,并且在很多地方都能使用到,所以定义为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUser map[int]*UserProcess
}

// 完成多userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUser:make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUser添加
func (u *UserMgr) AddOnlineUser(up *UserProcess){

	u.onlineUser[up.UserId] = up
}

// 对onlineUser删除
func (u *UserMgr) DeleteOnlineUser(userId int) {
	delete(u.onlineUser, userId)  // 删除字典中的值

}

// 返回当前所有在线的用户
func (u *UserMgr) GetAllOnlineUser() (map[int]*UserProcess) {
	return u.onlineUser
}

// 根据id返回对应的值
func (u *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	// 如何从map取出一个值,带检测方式的
	up, ok := u.onlineUser[userId]
	if !ok {
		// 说明查找的用户当前不在线
		err = fmt.Errorf("用户%d不在线或不存在", userId)
		return
	}
	return
}