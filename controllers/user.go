package controllers

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"net/http"
	"om_admin/models"
	"om_admin/utilities"
)

// GetAll...
// @Description user grid data api
// @Success 200	 Success {json}
// @Param	rows 		Query	 	string 		true 	"Number row get"
// @Param	page  		Query		string	 	true 	"which page data required"
// @Param	sortBy  	Query		string	 	true 	"sort by field"
// @Param	sortField  	Query		string	 	true 	"sort fields"
// @Param	searchAll  	Query		string	 	true 	"search all data "
// @Param	export  Query		string	 	true 	"user id"
// @Failure 404  Error: The key entered is incorrect. Please check it and not available condition {json}
// @Failure 408  Error: Please try again after sometime {json}
// @Failure 500  Error: Invalid request te {json}
// @router /:key [Get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) utilities.AllGridJson {

	var returnData utilities.AllGridJson

	rows := r.URL.Query().Get("rows")
	page := r.URL.Query().Get("page")
	sortBy := r.URL.Query().Get("sord")
	sortField := r.URL.Query().Get("sidx")
	searchAll := r.URL.Query().Get("_search")
	export := r.URL.Query().Get("export")

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	mobile := r.URL.Query().Get("mobile")
	creationTime := r.URL.Query().Get("creationTime")

	querarams := make(map[string]string)
	if searchAll == "true" {
		if id != "" {
			querarams["id"] = id
		}
		if name != "" {
			querarams["name"] = name
		}
		if email != "" {
			querarams["email"] = email
		}
		if mobile != "" {
			querarams["mobile"] = mobile
		}
		if creationTime != "" {
			querarams["creationTime"] = creationTime
		}
	}

	validSortFields := []string{"id", "name", "email", "mobile", "creationTime"}
	foundSortField := false
	for _, v := range validSortFields {
		if v == sortField {
			foundSortField = true
			break
		}
	}

	if !foundSortField {
		sortField = ""
	}

	start := 0
	end := cast.ToInt(rows)
	if rows != "" && page != "" {
		start = (cast.ToInt(rows) * cast.ToInt(page)) - cast.ToInt(rows)
		if cast.ToInt(page) == 1 {
			start = cast.ToInt(rows)
			end = 0
		}
	}

	baseQuery := `Select id, name, email, mobile, creationTime from _user`

	if cast.ToBool(export) == true {
		data, err := models.GetTotalNumberRecordForCSV(baseQuery, querarams)
		if err == nil {
			models.CSVExport(data, validSortFields, w)
		}

	} else {
		allData, err := models.GetDetails(cast.ToInt(start), cast.ToInt(end), sortField, sortBy, querarams, baseQuery)
		fmt.Println(allData)
		if err != nil {
			returnData.Records = 0
			return returnData
		}
		records, _ := models.GetTotalNumberRecordNEW(baseQuery, querarams, false)
		fmt.Println(records)
		fmt.Println("rows :: ", rows)
		fmt.Println("page :: ", page)

		returnData.Page = cast.ToInt(page)
		if len(allData) != 0 && len(records) > 0 {
			fmt.Println("records :: ", cast.ToInt(records[0]["Total"]))
			returnData.Records = cast.ToInt64(records[0]["Total"])
			returnData.Total = math.Floor(cast.ToFloat64(cast.ToInt64(records[0]["Total"])/cast.ToInt64(rows))) + 1
			returnData.Rows = allData
		}

	}

	return returnData
}

func GetUserDetailsFunction(ID int, currentUser *models.AdminUser) (returnData utilities.ResponseJSON) {
	allData, err := models.GetUserById(ID)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = allData
	} else {
		returnData.Code = 400
		returnData.Msg = "User Not Exist!"
	}
	return
}

func AddUserDetailsFunction(inputobj models.User, currentUser *models.AdminUser) (returnData utilities.ResponseJSON) {
	_, err := models.AddUser(&inputobj)
	if err == nil {
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:User Creation error"
	}
	return
}

func UpdateUserAttributeFunction(inputobj models.User, currentUser *models.AdminUser) (returnData utilities.ResponseJSON) {
	err := models.UpdateUserById(&inputobj)
	if err == nil {
		// return response
		returnData.Code = 200
		returnData.Msg = "Success"
		returnData.Model = nil
	} else {
		returnData.Code = 400
		returnData.Msg = "Failure:USER update error"
	}
	return
}
