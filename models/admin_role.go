package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AdminRole struct {
	Id       int    `orm:"column(id);auto"`
	RoleName string `orm:"column(role_name);size(225)"`
	Enable   int    `orm:"column(enable);size(11)"`
}

func (t *AdminRole) TableName() string {
	return "_admin_role"
}

func init() {
	orm.RegisterModel(new(AdminRole))
}

// AddAdminRole insert a new AdminRole into database and returns
// last inserted Id on success.
func AddAdminRole(m *AdminRole) (id int64, err error) {
	m.Enable = 1
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetRoleuserByID(authID string) (userRole []orm.Params, err error) {
	fmt.Println("===============Get  Admin User Role==================")
	o := orm.NewOrm()
	_, err = o.Raw("SELECT (SELECT action FROM `_admin_action` where id= _admin_permission.action_id ) as action, enable FROM `_admin_permission` where role_id =(SELECT id FROM `_admin_role` where role_name = (SELECT role FROM `_admin_user` where auth_id = ? )) and (SELECT enable FROM `_admin_action` where id= _admin_permission.action_id )  =1 ", &authID).Values(&userRole)
	return userRole, err
}

// GetAdminRoleById retrieves AdminRole by Id. Returns error if
// Id doesn't exist
func GetAdminRoleById(id int) (v *AdminRole, err error) {
	o := orm.NewOrm()
	v = &AdminRole{}
	err = o.QueryTable(new(AdminRole)).Filter("Id", id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllAdminRole retrieves all AdminRole matches certain condition. Returns empty list if
// no records exist
func GetAllAdminRole() (v []AdminRole, err error) {
	o := orm.NewOrm()
	v = []AdminRole{}
	_, err = o.QueryTable(new(AdminRole)).Filter("enable", "1").All(&v)
	fmt.Println(v)
	return v, err
}

// UpdateAdminRole updates AdminRole by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminRoleById(m *AdminRole) (err error) {
	o := orm.NewOrm()
	v := AdminRole{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdminRole deletes AdminRole by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdminRole(id int) (err error) {
	o := orm.NewOrm()
	v := AdminRole{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdminRole{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
