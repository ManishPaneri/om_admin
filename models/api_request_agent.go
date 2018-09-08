package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"time"
)

type ApiRequestAgent struct {
	Id           int       `orm:"column(id);auto"`
	RequestAgent string    `orm:"column(requestAgent);size(250)"`
	CreatedAt    time.Time `orm:"column(createdAt);type(datetime);null"`
}

func (t *ApiRequestAgent) TableName() string {
	return "_api_request_agent"
}

func init() {
	orm.RegisterModel(new(ApiRequestAgent))
}

// AddApiRequestAgent insert a new ApiRequestAgent into database and returns
// last inserted Id on success.
func AddApiRequestAgent(m *ApiRequestAgent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApiRequestAgentById retrieves ApiRequestAgent by requestAgent. Returns error if
// requestAgent doesn't exist
func GetApiRequestAgentByRequestAgent(requestAgent string) int {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("SELECT id  FROM _api_request_agent WHERE requestAgent = ?", requestAgent).Values(&maps)
	if len(maps) < 1 {
		return 0
	}
	return cast.ToInt(maps[0]["id"])

}

// UpdateApiRequestAgent updates ApiRequestAgent by Id and returns error if
// the record to be updated doesn't exist
func UpdateApiRequestAgentById(m *ApiRequestAgent) (err error) {
	o := orm.NewOrm()
	v := ApiRequestAgent{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApiRequestAgent deletes ApiRequestAgent by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApiRequestAgent(id int) (err error) {
	o := orm.NewOrm()
	v := ApiRequestAgent{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ApiRequestAgent{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetRequestAgentId(reqAgentString string) (int64, error) {
	reqAgent := GetApiRequestAgentByRequestAgent(reqAgentString)
	if reqAgent == 0 {
		reqAgentObj := &ApiRequestAgent{}
		reqAgentObj.RequestAgent = reqAgentString
		reqAgentObj.CreatedAt = time.Now()
		return AddApiRequestAgent(reqAgentObj)
	}
	return cast.ToInt64(reqAgent), nil
}
