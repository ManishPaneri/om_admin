package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"om_admin/controllers"
	"om_admin/models"
	"om_admin/utilities"
)

// URLMapping ...
func RoleHandler(w http.ResponseWriter, r *http.Request) {
	if utilities.Mapping(r, "GETALL") == true {
		utilities.URLReturnResponseJson(w, RoleGetAll(w, r))

	} else if utilities.Mapping(r, "GETONE") == true {
		utilities.URLReturnResponseJson(w, RoleGetOne(w, r))

	} else if utilities.Mapping(r, "PUT") == true {
		utilities.URLReturnResponseJson(w, RolePutOne(w, r))

	} else if utilities.Mapping(r, "POST") == true {
		utilities.URLReturnResponseJson(w, RolePostOne(w, r))
	} else {
		utilities.URLReturnResponseJson(w, "error")
	}
}

// GetAll ...
// @Description  Get All Admin Role details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	string 		true 	"request te"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetOne]
func RoleGetAll(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}
	key := r.URL.Query().Get("key")

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/view"):
		return controllers.GetRoleDetailsFunction()
		break

	case key == "action" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/view"):
		return controllers.GetRoleDetailsBermissionAction()
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: Invalid request te error!"
		returnData.Model = nil
		return

	}
	return
}

// GetOne ...
// @Description  Get Admin Role details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	string 		true 	"request te"
// @Param	id   Query		string	 	true 	"Admin Role id"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetOne]
func RoleGetOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}
	ID := r.URL.Query().Get("id")
	key := r.URL.Query().Get("key")

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/view/id"):
		return controllers.GetRoleDetailsByID(cast.ToInt(ID))
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: Invalid request te error!"
		returnData.Model = nil
		return

	}
	return
}

// PUT ...
// @Description  Admin Role details Add function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	 string 				true 	"request te"
// @Param	body Rawdata	 models.AdminRole 	true 	"request Admin Role Details"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [PUT]
func RolePutOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}

	key := r.URL.Query().Get("key")
	inputobj := models.AdminRole{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/add"):
		return controllers.AddNewRoleDetailsFunction(inputobj)
		break
	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: Invalid request te error!"
		returnData.Model = nil

	}
	return returnData
}

// POST ...
// @Description  ROLE details edit function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	 string 				true 	"request te"
// @Param	body Rawdata	 models.AdminRole 	true 	"request Admin role Details"
// @Failure 404  Error: The reason entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [post]
func RolePostOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}
	key := r.URL.Query().Get("key")

	inputobj := models.AdminRole{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/edit"):
		return controllers.UpdateRoleDetailsFunctionByID(inputobj)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: Invalid request te error!"
		returnData.Model = nil

	}
	return returnData
}
