package controllers

import (
	"fmt"
	"github.com/spf13/cast"
	"om_admin/models"
	"om_admin/utilities"
)

// Get Role Details Function
func GetRoleDetailsFunction() (returnData utilities.ResponseJSON) {
	allData, err := models.GetAllAdminRole()
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:role GET REQUEST"
	}
	return
}

// Get Role Details By Permission Action
func GetRoleDetailsBermissionAction() (returnData utilities.ResponseJSON) {
	allData, err := models.GetAllAdminAction()
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure: role get error"
	}
	return
}

// Get Role Details By ID
func GetRoleDetailsByID(ID int) (returnData utilities.ResponseJSON) {
	allData, err := models.GetAdminRoleById(ID)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:role doesn't exists"

	}
	return
}

// Update Role Details Id
func UpdateRoleDetailsFunctionByID(inputobj models.AdminRole) (returnData utilities.ResponseJSON) {
	err := models.UpdateAdminRoleById(&inputobj)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:role  Update error"
	}
	return
}

// Add New Role Details
func AddNewRoleDetailsFunction(inputobj models.AdminRole) (returnData utilities.ResponseJSON) {
	roleID, err := models.AddAdminRole(&inputobj)
	arrAction, err_action := models.GetAllAdminAction()
	fmt.Println("========================Role Add & Permission=============================")
	fmt.Println("Add Role :", roleID, arrAction)
	if err == nil && err_action == nil {
		fmt.Println(roleID)
		// adminrole, err1 := models.GetAdminRoleById(cast.ToInt(roleID))
		for _, value := range arrAction {

			fmt.Println("=======================Permission Add & Action==============================")
			models.AddAdminPermission(cast.ToInt(roleID), cast.ToInt(value.Id), cast.ToInt8(0))

		}
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:role Creation error"
	}
	return
}
