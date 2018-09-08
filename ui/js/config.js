
/*var go_server = "https://3x6kqh4hd0.execute-api.ap-south-1.amazonaws.com/Prod/";
var base_url = "http://admin.buynsta.com/";*/

var go_server = "http://localhost:9000/";
var base_url = "http://localhost:9000/ui/";

var global = {
	get_google_login_url : go_server+"GoogleLogin",
	get_all_users_list_url : go_server+"user",
	get_auth_url : go_server+"admin",
	get_logout_url:  base_url+"/login.html",
}

