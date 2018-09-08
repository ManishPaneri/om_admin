package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AdminPermission struct {
	Id       int          `orm:"column(id);auto"`
	RoleId   *AdminRole   `orm:"column(role_id);rel(fk)"`
	ActionId *AdminAction `orm:"column(action_id);rel(fk)"`
	Enable   int8         `orm:"column(enable)"`
}

func (t *AdminPermission) TableName() string {
	return "_admin_permission"
}

func init() {
	orm.RegisterModel(new(AdminPermission))
}

// AddAdminPermission insert a new AdminPermission into database and returns
// last inserted Id on success.
func AddAdminPermission(roleID int, ActionID int, Enable int8) (userRole []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("INSERT INTO `_admin_permission` (`role_id`,`action_id`,`enable`) VALUES (?, ?, ?)", &roleID, &ActionID, &Enable).Values(&userRole)
	fmt.Println("=========", userRole)
	return userRole, err
}

//GetAllAdminAction retrieves all AdminAction matches certain condition. Returns empty list if
// no records exist
func GetAllAdminPermission() (v []AdminPermission, err error) {
	o := orm.NewOrm()
	v = []AdminPermission{}
	_, err = o.QueryTable(new(AdminPermission)).RelatedSel().All(&v)
	fmt.Println(v)
	return v, err
}

// GetAdminPermissionById retrieves AdminPermission by Id. Returns error if
// Id doesn't exist
func GetAdminPermissionById(id int) (v *AdminPermission, err error) {
	o := orm.NewOrm()
	v = &AdminPermission{}
	err = o.QueryTable(new(AdminPermission)).Filter("Id", id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// GetAdminPermissionById retrieves AdminPermission by Id. Returns error if
// Id doesn't exist
func GetAdminPermissionByRoleId(id int) (v []AdminPermission, err error) {
	o := orm.NewOrm()
	v = []AdminPermission{}
	_, err = o.QueryTable(new(AdminPermission)).Filter("role_id", id).RelatedSel().All(&v)
	return v, err
}

// UpdateAdminPermission updates AdminPermission by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminPermissionById(m *AdminPermission) (err error) {
	o := orm.NewOrm()
	v := AdminPermission{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdminPermission deletes AdminPermission by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdminPermission(id int) (err error) {
	o := orm.NewOrm()
	v := AdminPermission{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdminPermission{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
