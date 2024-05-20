// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Emp is the golang structure for table emp.
type Emp struct {
	Id     uint   `json:"id"     orm:"id"      ` // ID
	DeptId uint   `json:"deptId" orm:"dept_id" ` // 所属部门
	Name   string `json:"name"   orm:"name"    ` // 姓名
	Gender int    `json:"gender" orm:"gender"  ` // 性别: 0=男 1=女
	Phone  string `json:"phone"  orm:"phone"   ` // 联系电话
	Email  string `json:"email"  orm:"email"   ` // 邮箱
	Avatar string `json:"avatar" orm:"avatar"  ` // 照片
}
