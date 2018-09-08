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
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if utilities.Mapping(r, "GETALL") == true {
		utilities.URLReturnResponseJson(w, UserGetAll(w, r))

	} else if utilities.Mapping(r, "EXPORT") == true {
		UserExportAll(w, r)

	} else if utilities.Mapping(r, "IMPORT") == true {
		utilities.URLReturnResponseJson(w, UserLegalNoticeImportAll(w, r))

	} else if utilities.Mapping(r, "GETONE") == true {
		utilities.URLReturnResponseJson(w, UserGetOne(w, r))

	} else if utilities.Mapping(r, "PUT") == true {
		utilities.URLReturnResponseJson(w, UserPutOne(w, r))

	} else if utilities.Mapping(r, "POST") == true {
		utilities.URLReturnResponseJson(w, UserPostOne(w, r))
	} else {
		utilities.URLReturnResponseJson(w, "error")
	}
}

// GetALL ...
// @Description Get User details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	string 		true 	"request te"
// @Param	id  Query		string	 	true 	"user id"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetAll]
func UserGetAll(w http.ResponseWriter, r *http.Request) interface{} {
	returnData := utilities.ResponseJSON{}
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}

	key := r.URL.Query().Get("key")

	switch true {

	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "customer/grid"):
		return controllers.GetAllUsers(w, r)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil
		return returnData

	}
	return returnData
}

// Export ...
// @Description  Export User details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	string 		true 	"request te"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [ExportAll]
func UserExportAll(w http.ResponseWriter, r *http.Request) interface{} {

	returnData := utilities.ResponseJSON{}
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}
	key := r.URL.Query().Get("key")

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "customer/export/csv"):
		return controllers.GetAllUsers(w, r)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil
		return returnData

	}
	return returnData
}

// GetOne ...
// @Description Get User details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	string 		true 	"request te"
// @Param	id  Query		string	 	true 	"user id"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [GetOne]
func UserGetOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}

	ID := r.URL.Query().Get("id")
	key := r.URL.Query().Get("key")

	switch true {

	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "customer/view"):
		return controllers.GetUserDetailsFunction(cast.ToInt(ID), currentUser)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil
		return

	}
	return
}

// PUT ...
// @Description User details Add function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	 string 		true 	"request te"
// @Param	body  Rawdata	 models.User 	true 	"request User Details"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [PUT]
func UserPutOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Failure"
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}

	key := r.URL.Query().Get("key")

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "customer/add"):
		inputobj := models.User{}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &inputobj)
		return controllers.AddUserDetailsFunction(inputobj, currentUser)
		break

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil

	}
	return returnData
}

// POST ...
// @Description User details edit function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key   Query	 	 string 		true 	"request te"
// @Param	body  Rawdata	 models.User 	true 	"request User Details"
// @Failure 404  Error: The reason entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [post]
func UserPostOne(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Failure"
	returnData.Model = nil

	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return returnData
	}
	key := r.URL.Query().Get("key")

	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "customer/edit"):
		inputobj := models.User{}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &inputobj)
		return controllers.UpdateUserAttributeFunction(inputobj, currentUser)

	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure:Authorization error"
		returnData.Model = nil

	}
	return returnData
}

// Import ...
// @Description  Import user details function , key if permission is enable
// @Success 200	 Success {json}
// @Param	key  Query	 	string 		true 	"request te"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [ImportAll]
func UserLegalNoticeImportAll(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	returnData.Msg = "Error: Invalid Session!"
	returnData.Code = 400
	returnData.Model = nil
	currentUser, _ := utilities.GetSessionUserDetails(w, r)
	if currentUser == nil {
		return
	}
	key := r.URL.Query().Get("key")
	switch true {
	case key == "" && utilities.CheckCurrentUserAccessPermission(currentUser, "customer/import"):
		// return controllers.UserLegalNoticeFuncton(currentUser, r)
		break
	default:
		fmt.Println("Not Authorized to access this resource")
		returnData.Code = 406
		returnData.Msg = "Failure: invalid request te error"
		returnData.Model = nil
		return
	}
	return
}
