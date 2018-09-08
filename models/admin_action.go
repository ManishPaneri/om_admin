package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AdminAction struct {
	Id          int    `orm:"column(id);auto"`
	Action      string `orm:"column(action);size(225)"`
	DisplayName string `orm:"column(display_name)"`
	Type        string `orm:"column(type)"`
	Enable      int    `orm:"column(enable);size(11)"`
}

func (t *AdminAction) TableName() string {
	return "_admin_action"
}

func init() {
	orm.RegisterModel(new(AdminAction))
}

// AddAdminAction insert a new AdminAction into database and returns
// last inserted Id on success.
func AddAdminAction(m *AdminAction) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetAllAdminAction retrieves all AdminAction matches certain condition. Returns empty list if
// no records exist
func GetAllAdminAction() (v []AdminAction, err error) {
	o := orm.NewOrm()
	v = []AdminAction{}
	_, err = o.QueryTable(new(AdminAction)).RelatedSel().All(&v)
	fmt.Println(v)
	return v, err
}

// GetAdminActionById retrieves AdminAction by Id. Returns error if
// Id doesn't exist
func GetAdminActionById(id int) (v *AdminAction, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(new(AdminAction)).Filter("id", id).RelatedSel().One(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

func GetPermissionActionByRole(RoleID int) (userRole []orm.Params, err error) {
	fmt.Println("===============Get  Admin User Role==================")
	o := orm.NewOrm()
	_, err = o.Raw("SELECT (select te from _admin_action where id = _admin_permission.action_id) as permission,(select action from _admin_action where id = _admin_permission.action_id) as action, enable FROM ._admin_permission where _admin_permission.role_id= ?", &RoleID).Values(&userRole)
	return userRole, err
}

// UpdateAdminAction updates AdminAction by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminActionById(m *AdminAction) (err error) {
	o := orm.NewOrm()
	v := AdminAction{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdminAction deletes AdminAction by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdminAction(id int) (err error) {
	o := orm.NewOrm()
	v := AdminAction{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdminAction{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
