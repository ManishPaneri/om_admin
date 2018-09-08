$("document").ready(function(e){
	tokenCheck();
	authConfiguration();
});

function tokenCheck(){
	var accesstoken = GetURLParameter("token");
	var id = GetCookie("frontend")
	if(id=="" && accesstoken!=""){
		SetCookie("frontend",accesstoken)
	}

}

function authConfiguration(sessionRole){
	sessionRole = authConfigurationCallAPI()
	//log.info("Start Auth validationfrontend UI Function");
	var array = $.map(sessionRole, function(value, index) {
  		if(value.enable==1){
  			$("#"+removeSplChar(value.action)).show();
  		}else{
  			$("#"+removeSplChar(value.action)).hide();
  		}
	});  
	//log.info("END Auth validation frontend UI Function");
}

function authConfigurationCallAPI(){
	try {
		var id = GetCookie("frontend") ;
		if(id==""){
			window.location.replace(base_url);	
		}
		//log.info("Auth validation API Call :"+global.get_auth_url+"?id="+id);
		var result = ajax_call(global.get_auth_url+"?id="+id,"GET");
		window.adminData = result["model"];
		//log.info("Auth validation API Call Response :"+result);
		var data = result["model"]["Details"];
		if(data==null){
			window.location.replace(base_url);
		}
		$("#authUserEmail").text(data.Email);
		$("#authUserName").text(data.Name);
		$("#authUserImage").html('<img alt="image" class="img-circle" src="'+data.Picture+'" style="width: 50px;"/>');
		return result["model"]["Roles"];
	} catch (e) {
		//log.error("An Auth error occurred", e);
		return 
	}

}

function logout() {
	SetCookie("frontend","");
	window.location.replace(global.get_logout_url);
}


function removeSplChar(sourceString){
	var r = sourceString.replace(/[`~!@#$%^&*()_|+\-=÷¿?;:'",.<>\{\}\[\]\\\/]/gi, '_');
	// console.log(r+"_ID");
	return r+"_ID";
}