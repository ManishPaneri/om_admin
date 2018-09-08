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
func PermissionHandler(w http.ResponseWriter, r *http.Request) {
	if utilities.Mapping(r, "GETALL") == true {
		utilities.URLReturnResponseJson(w, PermissionGetAll(w, r))

	} else if utilities.Mapping(r, "GETONE") == true {
		utilities.URLReturnResponseJson(w, PermissionGetOne(w, r))

	} else if utilities.Mapping(r, "POST") == true {
		utilities.URLReturnResponseJson(w, PermissionPostOne(w, r))
	} else {
		utilities.URLReturnResponseJson(w, "error")
	}
}

// GetAll ...
// @Description  Get All Admin Permission details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	string 		true 	"request te"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetOne]
func PermissionGetAll(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}
	key := r.URL.Query().Get("key")

	switch true {
	case cast.ToString(key) == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/view"):
		return controllers.GetPermissionDetailsFunction()
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
// @Description Get Admin Permission User details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	string 		true 	"request te"
// @Param	id   Query		string	 	true 	"Admin Permission id"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetOne]
func PermissionGetOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}
	ID := r.URL.Query().Get("id")
	key := r.URL.Query().Get("key")

	switch true {
	case cast.ToString(key) == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/view/id"):
		return controllers.GetPermissionDetailsByID(cast.ToInt(ID))
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

// POST ...
// @Description Admin Permission details edit function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	 string 				true 	"request te"
// @Param	body Rawdata	 models.AdminPermission 	true 	"request Admin Permission Details"
// @Failure 404  Error: The reason entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [post]
func PermissionPostOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}

	key := r.URL.Query().Get("key")
	inputobj := models.AdminPermission{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "role/edit"):
		return controllers.UpdatePermissionDetailsByID(inputobj)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: Invalid request te error!"
		returnData.Model = nil

	}
	return returnData
}
