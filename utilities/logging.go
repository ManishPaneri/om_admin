package utilities

import (
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"om_admin/models"
	"strings"
	"time"
)

var requestID int64

func LogApiDetailsInDb(r *http.Request, currentUser int) {
	currentDate := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println("reqest id is : ", r.Requ)
	// create api log insert object
	apiLog := models.AdminApiLog{}
	apiLog.Method = r.Method
	apiLog.RequestAgent, _ = models.GetRequestAgentId(r.UserAgent())
	apiLog.RequestPath = r.URL.Path
	apiLog.UrlQueryParams = r.URL.RawQuery
	apiLog.RequestIp = r.RemoteAddr
	apiLog.Permission = 0
	// check if this user had permission
	apiLog.CurrentUser = currentUser
	apiLog.CreatedTime, _ = time.Parse("2006-01-02 15:04:05", currentDate)
	apiLog.UpdatedOn, _ = time.Parse("2006-01-02 15:04:05", currentDate)
	// get request parameters
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close() //  must close
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	contentType := r.Header.Get("Content-type")

	fmt.Println("API Request ContentType :", contentType)

	if err != nil || strings.Contains(contentType, "csv") || strings.Contains(contentType, "multipart/form-data") {
		// request file name store
		apiLog.RequestParameter = strings.Trim(r.URL.Path, "/") + "/" + cast.ToString(apiLog.CreatedTime) + ".csv"

		//Import S3 bucket in log
		ImportS3BucketInCSVLogs(apiLog.RequestParameter, string(body))
	} else {
		//request body json store in mysql
		apiLog.RequestParameter = string(body)

	}

	fmt.Println("*************ADMIN LOG******************")
	// inserting Api logs
	insertedId, insertErr := models.AddAdminApiLog(&apiLog)
	if insertErr != nil {
		fmt.Println("Error inserting Api logs :", insertErr)

	} else {
		requestID = insertedId
		fmt.Println("Api Logs Added Successfully :", "**************Request Id : ", insertedId, "*************")

	}
	return
}

func UpdateLoggedApiByRequestID(data interface{}, permission string) {
	code := escapeResponseCode(data)

	if cast.ToInt(strings.Trim(code, " ")) == 406 || cast.ToBool(permission) == false {
		permission = "0"
	} else {
		permission = "1"
	}
	_, err := models.UpdateApiLogBYRequestId(cast.ToInt(requestID), code, permission)
	if err != nil {
		fmt.Println("Error updating APi logs with insert Id : ", requestID, err)
	} else {
		fmt.Println("Api Logs Updated Successfully | Insert Id :", requestID)
	}

}

func escapeResponseCode(data interface{}) (ResponseCode string) {
	e := spew.Sdump(data)
	// fmt.Print("code:", e)
	response := strings.Split(e, "Code: (int)")
	if len(response) > 1 || strings.Contains(e, "Code:") {
		code := strings.Split(response[1], ",")
		return code[0]
	} else {
		if strings.Contains(e, "Page:") {
			return "1000"
		} else {
			return "400"
		}

	}

}
