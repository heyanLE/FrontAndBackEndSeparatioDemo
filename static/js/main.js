function login() {
    FetchPost("/api/v1/login",{
        username:$("#username_login").val(),
        password:$("#password_login").val()
    },(code,message,value) => {
        window.alert(message)
    })
}

function register() {
    FetchPost("/api/v1/register",{
        username:$("#username_register").val(),
        password:$("#password_register").val()
    },(code,message,value) => {
        window.alert(message)
    })
}

function loginOut() {
    FetchGet("/api/v1/login-out",(code,message,value) => {
        $("#login_p").css("display","none");
        $("#dis_login_p").css("display","block");
        window.alert(message)
    })
}

function getUserNow() {
    FetchGet("/api/v1/user",(code,message,value) => {
        if (code === 200){
            $("#login_p").css("display","block");
            $("#dis_login_p").css("display","none");
            $("#username_span").text(value.username)
        }else{
            $("#login_p").css("display","none");
            $("#dis_login_p").css("display","block");
        }
    })
}

function backToIndex() {
    window.location.href = "/index.html"
}