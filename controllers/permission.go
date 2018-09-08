package controllers

import (
	"om_admin/models"
	"om_admin/utilities"
)

// Get All Permission Details
func GetPermissionDetailsFunction() (returnData utilities.ResponseJSON) {
	allData, err := models.GetAllAdminPermission()
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:Order GET REQUEST"
	}
	return
}

// Get Permission details by Id
func GetPermissionDetailsByID(ID int) (returnData utilities.ResponseJSON) {
	allData, err := models.GetAdminPermissionById(ID)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure: permission doesn't exists"

	}
	return
}

// Get Permission details by role Id
func GetPermissionDetailsByRoleID(ID int) (returnData utilities.ResponseJSON) {
	allData, err := models.GetAdminPermissionByRoleId(ID)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure: permission doesn't exists"
	}
	return
}

// Update Permission details By id
func UpdatePermissionDetailsByID(inputobj models.AdminPermission) (returnData utilities.ResponseJSON) {
	err := models.UpdateAdminPermissionById(&inputobj)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure: Permission Update error"
	}
	return
}
