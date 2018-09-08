package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/go-sessions"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"om_admin/models"
	"time"
)

const htmlIndex = `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "",
		ClientID:     "",
		ClientSecret: "",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "random"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	runmode := cast.ToString(viper.Get("runmode"))
	URL := cast.ToString(viper.Get(runmode+".base_url")) + "login.html"
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	runmode := cast.ToString(viper.Get("runmode"))
	googleOauthConfig.RedirectURL = cast.ToString(viper.Get(runmode+".redirectURL")) + "GoogleCallback"
	googleOauthConfig.ClientID = cast.ToString(viper.Get(runmode + ".clientID"))
	googleOauthConfig.ClientSecret = cast.ToString(viper.Get(runmode + ".clientSecret"))
	fmt.Println("GoogleCallBack", googleOauthConfig)

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("==================Login=========================")
	fmt.Println("Login Token :", token.AccessToken)
	fmt.Println("================================================")
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	UserObj := make(map[string]interface{})
	json.Unmarshal(contents, &UserObj)
	runmode := cast.ToString(viper.Get("runmode"))
	URL := cast.ToString(viper.Get(runmode+".base_url")) + "dashboard.html?token=" + cast.ToString(token.AccessToken)
	fmt.Println("UI Redirect Url: ", URL)
	AdminUser, err := models.GetAdminUserByAuthId(cast.ToString(UserObj["id"]))

	if err != nil {
		newAdminUser := models.AdminUser{}
		newAdminUser.AuthId = cast.ToString(UserObj["id"])
		newAdminUser.VerifiedEmail = cast.ToInt8(UserObj["verified_email"])
		newAdminUser.Email = cast.ToString(UserObj["email"])
		newAdminUser.Hd = cast.ToString(UserObj["hd"])
		newAdminUser.Name = cast.ToString(UserObj["name"])
		newAdminUser.Picture = cast.ToString(UserObj["picture"])
		newAdminUser.Role = "user"

		models.AddAdminUser(&newAdminUser)
		sess := sessions.Start(w, r)
		sess.Set(cast.ToString(UserObj["id"]), newAdminUser)
		fmt.Println("user not exist ")
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "frontend", Value: cast.ToString(token.AccessToken), Expires: expiration}
		http.SetCookie(w, &cookie)

		fmt.Println(sess.Get(cast.ToString(UserObj["id"])))
		http.Redirect(w, r, URL, http.StatusTemporaryRedirect)

	} else {
		sess := sessions.Start(w, r)
		sess.Set(cast.ToString(UserObj["id"]), AdminUser)

		fmt.Println("user exist already")
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "frontend", Value: cast.ToString(token.AccessToken), Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Println(sess.Get(cast.ToString(UserObj["id"])))
		http.Redirect(w, r, URL, http.StatusTemporaryRedirect)

	}
}
