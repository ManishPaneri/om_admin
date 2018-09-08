package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"time"
)

type User struct {
	Id           int       `orm:"column(id);auto"`
	Name         string    `orm:"column(name);size(64);null"`
	Email        string    `orm:"column(email);size(64);null"`
	Mobile       string    `orm:"column(mobile);size(13);null"`
	CreationTime time.Time `orm:"column(creationTime);te(datetime);null"`
}

func (t *User) TableName() string {
	return "_user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	m.Id = cast.ToInt(id)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{}
	err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		fmt.Println("==================upadteUser===========")
		fmt.Println(m)
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			fmt.Println("user update failed :", err)
		}
	}

	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
