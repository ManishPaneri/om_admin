package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
	Id            int    `orm:"column(id);auto"`
	Name          string `orm:"column(name);size(225)"`
	Email         string `orm:"column(email);size(225)"`
	Role          string `orm:"column(role)"`
	Picture       string `orm:"column(picture)"`
	Hd            string `orm:"column(hd)"`
	VerifiedEmail int8   `orm:"column(verified_email)"`
	AuthId        string `orm:"column(auth_id);size(256)"`
	Enable        int    `orm:"column(enable);size(11);null"`
}

func (t *AdminUser) TableName() string {
	return "_admin_user"
}

func init() {
	orm.RegisterModel(new(AdminUser))
}

// AddAdminUser insert a new AdminUser into database and returns
// last inserted Id on success.
func AddAdminUser(m *AdminUser) (id int64, err error) {
	m.Enable = 1
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdminUserById retrieves AdminUser by Id. Returns error if
// Id doesn't exist
func GetAdminUserById(id int) (v *AdminUser, err error) {
	o := orm.NewOrm()
	v = &AdminUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetAllAdminUser retrieves all AdminUser matches certain condition. Returns empty list if
// no records exist
func GetAllAdminUser() (v []AdminUser, err error) {
	o := orm.NewOrm()
	v = []AdminUser{}
	_, err = o.QueryTable(new(AdminUser)).Filter("enable", 1).All(&v)
	return v, err
}

//GetAdminUserByAuthId get user row by authID
func GetAdminUserByAuthId(authid string) (v *AdminUser, err error) {

	o := orm.NewOrm()
	v = &AdminUser{}
	err = o.QueryTable(new(AdminUser)).Filter("authid", authid).One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateAdminUser updates AdminUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdminUserById(m *AdminUser) (err error) {
	m.Enable = 1
	o := orm.NewOrm()
	v := AdminUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdminUser deletes AdminUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdminUser(id int) (err error) {
	o := orm.NewOrm()
	v := AdminUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdminUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
