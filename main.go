package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"om_admin/controllers"
	"om_admin/handlers"
	"om_admin/utilities"
)

func init() {
	setUpViper()
	registerDatabase()
	utilities.LogS3BucketInitialization()
}

func main() {
	runmode := cast.ToString(viper.Get("runmode"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.HandleMain)
	mux.HandleFunc("/GoogleLogin", controllers.HandleGoogleLogin)
	mux.HandleFunc("/GoogleCallback", controllers.HandleGoogleCallback)
	mux.HandleFunc("/user", handlers.UserHandler)
	mux.HandleFunc("/permission", handlers.PermissionHandler)
	mux.HandleFunc("/role", handlers.RoleHandler)
	mux.HandleFunc("/admin", handlers.AdminHandler)
	// mux.HandleFunc("/statics", http.FileServer(http.Dir("./ui")))
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui"))))
	handler := cors.Default().Handler(mux)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowCredentials: true,
	})
	// Insert the middleware
	fmt.Println(cast.ToString(viper.Get(runmode + ".base_url")))
	handler = c.Handler(handler)
	http.ListenAndServe(":9000", handler)
}

//function to register the database to beego orm
func registerDatabase() {
	runmode := cast.ToString(viper.Get("runmode"))

	mysql := viper.Get(runmode + ".mysql").(map[string]interface{})
	mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string)
	log.Println("conf", mysqlConf)
	orm.RegisterDataBase("default", "mysql", mysqlConf)
	orm.Debug = true
}

//set up config file from conf folder
func setUpViper() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("global")
}
