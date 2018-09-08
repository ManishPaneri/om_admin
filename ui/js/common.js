$("#header").load("includes/header.html"); 
$("#nav").load("includes/nav.html"); 
$("#footer").load("includes/footer.html"); 
$("#modal").load("includes/modal.html"); 

function ajax_call(url, method ){
  var cookie = GetCookie("frontend") ;
  SetCookie("frontend",cookie)
  var final_result;
  $.ajax({
      type:method,
      dataType : 'JSON',
      url: url,
      async : false,
      // crossDomain: true,
      headers: {
        "Authorization": "Bearer "+cookie ,
        "Content-Type": "application/json; charset=utf-8",
      },
      success: function(res,status,xhr){
         final_result = res;
      },
      error: function (xhr,status,error) {
        console.log(status, error, xhr);
      }
   });
   return final_result;
}

function post_ajax_call(url, data ){
  var cookie = GetCookie("frontend") ;
  SetCookie("frontend",cookie)
  var final_result;
  $.ajax({
      type:"POST",
      url: url,
      data:JSON.stringify(data),
      async : false,
      dataType : 'JSON',
      // crossDomain: true,
      headers: {
        "Authorization": "Bearer "+cookie ,
        "Content-Type": "application/json; charset=utf-8",
      },
      success: function(res,status,xhr){
         final_result = res;
      },
      error: function (xhr,status,error) {
        console.log(status, error, xhr);
      }
   });
   return final_result;
}

function put_ajax_call(url, data ){
  var cookie = GetCookie("frontend") ;
  SetCookie("frontend",cookie)
  var final_result;
  $.ajax({
      type:"PUT",
      url: url,
      data:JSON.stringify(data),
      async : false,
      dataType : 'JSON',
      crossDomain: true,
      // crossDomain: true,
      headers: {
       "Authorization": "Bearer "+cookie ,
        "Content-Type": "application/json; charset=utf-8",
      },
      success: function(res,status,xhr){
          console.log('res',res);
         final_result = res;
      },
      error: function (xhr,status,error) {
        console.log(status, error, xhr);
      }
   });
   return final_result;
}

//post_ajax_call_lambda
function post_ajax_call_lambda(url, data ){
  var cookie = GetCookie("frontend") ;
  var final_result;
  $.ajax({
      type:"POST",
      url: url,
      data:JSON.stringify(data),
      headers: {
        "Content-Type" : "application/json",
        "Authorization": "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImMwNmEyMTQ5YjdmOTU3MjgwNTJhOTg1YWRlY2JmNWRlMDQ3Y2RhNmYifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vYnV5bnN0YS1kZXYiLCJuYW1lIjoiUmFzaG1pIFJhbmdhcmFqYW4iLCJwaWN0dXJlIjoiaHR0cHM6Ly9sb29rYXNpZGUuZmFjZWJvb2suY29tL3BsYXRmb3JtL3Byb2ZpbGVwaWMvP2FzaWQ9MTAyMTU3MDgwMjg3MTQ3ODEmaGVpZ2h0PTEwMCZ3aWR0aD0xMDAiLCJhdWQiOiJidXluc3RhLWRldiIsImF1dGhfdGltZSI6MTUyMzI3MTcwNywidXNlcl9pZCI6ImlHdFVWbkNSNVlNQUlaVzUyN1M2dkpVcmdybDIiLCJzdWIiOiJpR3RVVm5DUjVZTUFJWlc1MjdTNnZKVXJncmwyIiwiaWF0IjoxNTIzNzAyODI5LCJleHAiOjE1MjM3MDY0MjksImVtYWlsIjoiYW5vbnltb3VzLnpvbWJpZUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZmFjZWJvb2suY29tIjpbIjEwMjE1NzA4MDI4NzE0NzgxIl0sImVtYWlsIjpbImFub255bW91cy56b21iaWVAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZmFjZWJvb2suY29tIn19.mSMRyMzcuUXqjQhlYbHIvVcjSU7mAOTTZo4hvD2xNt5wZIw85q5V7p9aoD4uwbRvRVdYFNklj0LJq707Klnc1i_fE25quROB-yM6zaVtMNs9lvYyKqYQEyXcXYzE8kVI42-VZmBYA1rkIUYtMo0bxYGbNGYe6x2ZabL2x-6UeL3u9T1_nw3K2nqWRx_bjt65_hrabciTRV93j1KjXlSlNwkRbQEGWug18tWJeYUUY3zf22AVQ77-MBkxU7OId2OdTgb8YL3LSo1XaMLyrRBrxo5AMpeu3u-fAzCcOxwXdghW4rxsj_a3gGhlubyk8_sdPB4l_-C_tsyIzdq5_cAOtQ"
      },
      success: function(res,status,xhr){
         final_result = res;
      }
   });
   return final_result;
}

function SetCookie(cname, cvalue) {
    var d = new Date();
    d.setTime(d.getTime()+3600*24*1000);
    var expires = "expires="+ d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}


function create_arr(res){
  var result = Object.keys(res).map(function(key) {
    return [res[key]];
  });
   console.log(result);
  return result;
}


function GetURLParameter(sParam){
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
   
    for (var i = 0; i < sURLVariables.length; i++){
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam){
            return sParameterName[1];
        }
    }
}

function GetCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function jsonConvertToArray(result){
    var array = $.map(result[0], function(value, index) {
      return [value];
    }); 
  return array;
}

function convertObjToArray(result){
  var array = $.map(result, function(value, index) {
      return [value];
  });
  return array;
}

function editSuccess(result,id){
  if(result.code == 400){
                $("#"+id).show();
                setTimeout(function(){
                  location.reload();
                }, 2000);
            }
}

function removeBlank(str){
  var newStr = str.replace(/ /g, "_");
  return newStr;
}

var globalData = {
            exportButton : {
                    dom: 'Bfrtip',
                    buttons: [
                        {
                            extend: 'csv',
                            text: 'CSV',
                            className: 'btn btn-default',
                            exportOptions: {
                                columns: 'th:not(:last-child)'
                                            }
                         },
                         {
                            extend: 'excel',
                            text: 'Excel',
                            className: 'btn btn-default',
                            exportOptions: {
                                columns: 'th:not(:last-child)'
                                            }
                         },
                         {
                            extend: 'pdf',
                            text: 'PDF',
                            className: 'btn btn-default',
                            exportOptions: {
                                columns: 'th:not(:last-child)'
                                            }
                         }
                    ],
                    "scrollX": true
                }
            }



function accountNumberMask(str){
  return str.replace(/.(?=.{4})/g, 'x');
}

function makeCapital(str){
}


function dateFormat(date){
 if(date=="0001-01-01T00:00:00Z"){
    return "NA"
 }else{
    var d = new Date(date),
      month = '' + (d.getMonth() + 1),
      day = '' + d.getDate(),
      year = d.getFullYear(),
      time = d.toLocaleTimeString(), date;

    if (month.length < 2) month = '0' + month;
    if (day.length < 2) day = '0' + day;

    date =  [year, month, day].join('-');

    return date + ' ' + time;
 } 
  
}

Number.prototype.padLeft = function(base,chr){
   var  len = (String(base || 10).length - String(this).length)+1;
   return len > 0? new Array(len).join(chr || '0')+this : this;
}


function CustomerSearchFunction(){
    var value = $("#top_customer_search").val();
    if(value!=""){
       var result = ajax_call(global.get_search_url+"?key=user&id="+value,"GET");
       console.log(result);
       if(result.code==200){
          console.log(result.model.id)
          window.open("profile.html?id="+result.model.id, "_blank", "");
       }else{
         swal({     
          title: "Oops",
          text: result.msg,
            type: "error"
          },function () {

        });
      }
    }
}


function LoanSearchFunction(){
    var value = $("#top_loan_search").val();
    if(value!=""){
       var result = ajax_call(global.get_search_url+"?key=loan&id="+value,"GET");
       console.log(result);
       if(result.code==200){
          console.log(result.model.id)
          window.open("order_detail.html?id="+result.model.id, "_blank", "");
       }else{
         swal({     
          title: "Oops",
          text: result.msg,
            type: "error"
          },function () {

        });
      }
    }
}



function login(){
     window.location.replace(global.get_google_login_url);
}



function PermisssionDisplayFunction(RefID,DisplayID){
    if($('#'+RefID).css('display') != 'none'){  
        $("#"+DisplayID).show();
    }else{
      $("#"+DisplayID).hide();
    }  
}


function ModalFunction(id,action){
  $('#'+id+' .form-control').prop("disabled", false);
  $("#"+id).modal(action);
 }


function randomChar(number) {
  var text = "";
  var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  for (var i = 0; i < number; i++)
    text += possible.charAt(Math.floor(Math.random() * possible.length));

  return text;
}



function ActionLink(id, page, columns, table) {
    var row = id.split("=");
    var row_ID = row[1];
    console.log("page",page,"columns", columns)
    var sitename= $("#"+table).getCell(row_ID, columns);

    var url = base_url+""+page+".html?id="+sitename; // sitename will be like google.com or yahoo.com
    window.open(url);
}

function JSONToCSVConvertor(JSONData, ReportTitle, ShowLabel, headers) {
    //If JSONData is not an object then JSON.parse will parse the JSON string in an Object
    var arrData = typeof JSONData != 'object' ? JSON.parse(JSONData) : JSONData;

    var CSV = '';

    //This condition will generate the Label/Header
    if (ShowLabel) {
        var row = "";

        if(headers) {
          for (var i = 0; i < headers.length; i += 1) {
            if(headers[i] !== 'Action') {
              row += headers[i] + ',';
            }
          }
        } else {
          //This loop will extract the label from 1st index of on array
          for (var index in arrData[0]) {
              //Now convert each value to string and comma-seprated
              if(index !== 'Action') {
                row += index + ',';
              }
          }
        }

        row = row.slice(0, -1);
        //append Label row with line break
        CSV += row + '\r\n';
    }

    //1st loop is to extract each row
    for (var i = 0; i < arrData.length; i++) {
        var row = "";

        //2nd loop will extract each column and convert it in string comma-seprated
        for (var index in arrData[i]) {
            if(index !== 'Action') {
              row += '"' + arrData[i][index] + '",';
            }
        }

        row.slice(0, row.length - 1);

        //add a line break after each row
        CSV += row + '\r\n';
    }

    if (CSV == '') {
        alert("Invalid data");
        return;
    }

    //Generate a file name
    var fileName = "";
    //this will remove the blank-spaces from the title and replace it with an underscore
    fileName += ReportTitle.replace(/ /g,"_");

    //Initialize file format you want csv or xls
    var uri = 'data:text/csv;charset=utf-8,' + escape(CSV);

    // Now the little tricky part.
    // you can use either>> window.open(uri);
    // but this will not work in some browsers
    // or you will not get the correct file extension

    //this trick will generate a temp <a /> tag
    var link = document.createElement("a");
    link.href = uri;

    //set the visibility hidden so it will not effect on your web-layout
    link.style = "visibility:hidden";
    link.download = fileName + ".csv";

    //this part will append the anchor tag and remove it after automatic click
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}