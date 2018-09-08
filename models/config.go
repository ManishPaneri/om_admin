package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
)

type Config struct {
	Id         int    `orm:"column(id);auto"`
	ConfigType string `orm:"column(configType);size(45)" description:"config type - email/bank/otp/digio/..."`
	Key        string `orm:"column(key);size(45)" description:"Name of the config"`
	Value      string `orm:"column(value);size(45)" description:"Value of config"`
}

func (t *Config) TableName() string {
	return "_config"
}

func init() {
	orm.RegisterModel(new(Config))
}

// AddConfig insert a new Config into database and returns
// last inserted Id on success.
func AddConfig(m *Config) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetAllYpAdminUser retrieves all YpAdminUser matches certain condition. Returns empty list if
// no records exist
func GetAllConfig() (v []Config, err error) {
	o := orm.NewOrm()
	v = []Config{}
	_, err = o.QueryTable(new(Config)).All(&v)
	fmt.Println(v)
	return v, err
}

// GetConfigById retrieves Config by Id. Returns error if
// Id doesn't exist
func GetConfigById(id int) (v *Config, err error) {
	o := orm.NewOrm()
	v = &Config{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//get configs by type
func GetConfigsByType(configType string) (config map[string]string, err error) {
	o := orm.NewOrm()
	var v []*Config
	_, err = o.QueryTable(new(Config)).Filter("configType", configType).All(&v)
	if err != nil {
		return nil, err
	}
	config = make(map[string]string)
	for _, value := range v {
		config[value.Key] = value.Value
	}
	return config, nil
}

// GetAllConfig retrieves all Config matches certain condition. Returns empty list if
// no records exist

// UpdateConfig updates Config by Id and returns error if
// the record to be updated doesn't exist
func UpdateConfigById(m *Config) (err error) {
	o := orm.NewOrm()
	v := Config{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteConfig deletes Config by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConfig(id int) (err error) {
	o := orm.NewOrm()
	v := Config{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Config{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetConfigByType(configType string) (conf []*Config, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("_config").Filter("configType__startswith", configType).All(&conf)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		return conf, nil
	}
	return nil, err
}

// GetConfigByTypeAndKey... get a config by type and key
func GetConfigByTypeAndKey(configType string, key string) (conf Config, err error) {
	o := orm.NewOrm()
	conf = Config{}
	err = o.QueryTable("_config").Filter("configType", configType).Filter("key", key).One(&conf)
	return conf, err
}

// GetConfigByTypeAndKey ... get a config by type and key
func GetConfigValueNameByTypeAndKey(configType string, key string) (string, string, error) {
	o := orm.NewOrm()
	var value []orm.Params
	num, err := o.QueryTable(new(Config)).Filter("configType", configType).Filter("key", key).Values(&value)
	if err == nil && num != 0 && len(value) != 0 {
		fmt.Printf("Result Nums: %d\n", num)
		return cast.ToString(value[0]["Value"]), cast.ToString(value[0]["Name"]), nil
	}
	fmt.Println(err)
	return "", "", err
}
