// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Dept is the golang structure for table dept.
type Dept struct {
	Id     uint   `json:"id"     orm:"id"     ` // ID
	Pid    uint   `json:"pid"    orm:"pid"    ` // 上级部门ID
	Name   string `json:"name"   orm:"name"   ` // 部门名称
	Leader string `json:"leader" orm:"leader" ` // 部门领导
	Phone  string `json:"phone"  orm:"phone"  ` // 联系电话
	Emps []Emp		`json:"emps" orm:"with:dept_id=id"` // 员工, 当前表的id放在后面
}
