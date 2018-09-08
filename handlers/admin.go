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
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if utilities.Mapping(r, "GETALL") == true {
		utilities.URLReturnResponseJson(w, AdminGetAll(w, r))

	} else if utilities.Mapping(r, "GETONE") == true {
		utilities.URLReturnResponseJson(w, AdminGetOne(w, r))

	} else if utilities.Mapping(r, "PUT") == true {
		utilities.URLReturnResponseJson(w, AdminPutOne(w, r))

	} else if utilities.Mapping(r, "POST") == true {
		utilities.URLReturnResponseJson(w, AdminPostOne(w, r))
	} else {
		utilities.URLReturnResponseJson(w, "error")
	}
}

// GetAll ...
// @Description Get All Admin User details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	string 		true 	"request te"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [Get]
func AdminGetAll(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 406

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}

	key := r.URL.Query().Get("key")

	switch true {

	case cast.ToString(key) == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "admin/grid"):
		return controllers.GetAllAdminUserFunction()
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
// @Description Get Admin User details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	string 		true 	"request te"
// @Param	id    Query		string	 	true 	"Admin user id"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetOne]
func AdminGetOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 406

	// currentUser, _ := utilities.GetSessionUserDetails(w, r)
	// if currentUser == nil {
	// 	return
	// }

	ID := r.URL.Query().Get("id")
	key := r.URL.Query().Get("key")

	switch true {

	case cast.ToString(key) == "" /*&& utilities.CheckCurrentUserAccessPermission(currentUser, "admin/view")*/ :
		return controllers.GetAdminUserDetailsFunction(ID)
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
// @Description Admin User details Add function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	 string 				true 	"request te"
// @Param	body  Rawdata	 models.AdminUser 	true 	"request Admin User Details"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [PUT]
func AdminPutOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 406

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}

	key := r.URL.Query().Get("key")
	inputobj := models.AdminUser{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "admin/add"):
		return controllers.AddAdminUserDetailsFunction(inputobj, currentUser)
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
// @Description User details edit function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	 string 				true 	"request te"
// @Param	body  Rawdata	 models.AdminUser 	true 	"request Admin User Details"
// @Failure 404  Error: The reason entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [post]
func AdminPostOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 406
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}

	key := r.URL.Query().Get("key")
	inputobj := models.AdminUser{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &inputobj)

	switch true {

	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "admin/edit"):
		return controllers.UpdateAdminUserDetailsFunction(inputobj, currentUser)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: Invalid request te error!"
		returnData.Model = nil

	}
	return returnData
}
