package controllers

import (
	"encoding/json"
	// "fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"om_admin/models"
	"om_admin/utilities"
)

type AdminReturnJson struct {
	Details     *models.AdminUser
	Roles       interface{}
	RoleDetails []models.AdminRole
}

// Get All Admin User Details
func GetAllAdminUserFunction() (returnData utilities.ResponseJSON) {
	allData, err := models.GetAllAdminUser()
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:Admin GET REQUEST"
	}
	return
}

//Get Admin User Details By Admin ID
func GetAdminUserDetailsFunction(ID string) (returnData utilities.ResponseJSON) {
	UserObj := AdminReturnJson{}
	allData, err := models.GetAdminUserById(cast.ToInt(ID))
	if err == nil {
		role, _ := models.GetRoleuserByID(allData.AuthId)
		UserObj.Details = allData
		UserObj.Roles = role
		UserObj.RoleDetails, _ = models.GetAllAdminRole()
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = UserObj
	} else {
		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + cast.ToString(ID))
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		rUserObj := make(map[string]interface{})
		json.Unmarshal(contents, &UserObj)
		allData, err := models.GetAdminUserByAuthId(cast.ToString(rUserObj["id"]))
		if err == nil {
			role, _ := models.GetRoleuserByID(allData.AuthId)
			UserObj.Details = allData
			UserObj.Roles = role
			returnData.Code = 200
			returnData.Msg = "Success"
			returnData.Model = UserObj
		} else {
			returnData.Code = 400
			returnData.Msg = "Failure:Admin doesn't exists"
		}
	}
	return
}

//Update Admin User Details By Admin ID
func UpdateAdminUserDetailsFunction(inputobj models.AdminUser, currentUser *models.AdminUser) (returnData utilities.ResponseJSON) {
	err := models.UpdateAdminUserById(&inputobj)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:Admin Update error"
	}
	return
}

// Add New Admin User Details
func AddAdminUserDetailsFunction(inputobj models.AdminUser, currentUser *models.AdminUser) (returnData utilities.ResponseJSON) {
	_, err := models.AddAdminUser(&inputobj)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:Admin Creation error"
	}
	return
}
