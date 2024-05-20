// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        orm:"id"         ` // ID
	Username  string      `json:"username"  orm:"username"   ` // 用户名
	Nickname  string      `json:"nickname"  orm:"nickname"   ` // 昵称
	Password  string      `json:"password"  orm:"password"   ` // 密码
	Avatar    string      `json:"avatar"    orm:"avatar"     ` // 头像
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
}
