package utilities

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"om_admin/models"
	"strings"
)

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

type AllGridJson struct {
	Page    int         `json:"page"`
	Total   float64     `json:"total"`
	Records int64       `json:"records"`
	Rows    interface{} `json:"rows"`
}

func GetSessionUserDetails(w http.ResponseWriter, r *http.Request) (*models.AdminUser, error) {
	fmt.Println("=====================LOGIN================================")
	// }
	frontend := r.Header.Get("Authorization")
	fmt.Println("==================")
	fmt.Println("Authorization Token: ", frontend)
	fmt.Println("==================")
	if cast.ToString(frontend) == "" {
		for _, cookie := range r.Cookies() { // loop to get all cookies
			fmt.Println("Cookie :", cookie.Value)
			if cookie.Name == "frontend" { // only required to frontend cookies
				frontend = cast.ToString(cookie.Value) // check frontend cookie value
			}
		}
	} else {
		auth := strings.Split(frontend, "Bearer ")
		fmt.Println("Authorization token  :", auth)
		frontend = cast.ToString(auth[1])
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + frontend)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	UserObj := make(map[string]interface{})
	json.Unmarshal(contents, &UserObj)
	frontend = cast.ToString(UserObj["id"])
	userObj, err := models.GetAdminUserByAuthId(frontend)
	LogApiDetailsInDb(r, userObj.Id)
	w.Header().Set("currentUser", "true")
	fmt.Println("User", userObj)
	fmt.Println("=====================END================================")
	return userObj, err
}

func CheckCurrentUserAccessPermission(currentUser *models.AdminUser, requestapi string) bool {
	fmt.Println("=====================Permission Check================================")
	fmt.Println("RequestTe", requestapi)
	flag := false
	UserRole, _ := models.GetRoleuserByID(currentUser.AuthId)
	for _, value := range UserRole {
		fmt.Println("role:", value["action"], value["enable"])
		if requestapi == cast.ToString(value["action"]) && cast.ToInt(value["enable"]) == 1 {
			flag = true
		}
	}
	fmt.Println("=====================END================================")
	return flag
}

func CalculateAmountConvFee(amount float32, feeTe string, feeAmount float32) (totalAmount float32, convFee float32, err error) {
	fmt.Println("CalculateAmountConvFee")
	amountFloat, err := cast.ToFloat32E(amount)
	if err != nil {
		return 0.0, 0.0, err
	}

	feeAmountFloat, err := cast.ToFloat32E(feeAmount)
	if err != nil {
		return 0.0, 0.0, err
	}

	if feeTe == "fixed" {
		convFee = feeAmountFloat
		totalAmount = amountFloat + convFee

	} else {
		convFee = (amountFloat * feeAmountFloat) / 100.0
		totalAmount = amountFloat + convFee
	}
	return
}

func maskString(any string, length int) string {
	if len(any) <= length {
		return any
	}
	return strings.Repeat("x", len(any)-length) + any[len(any)-length:]
}
