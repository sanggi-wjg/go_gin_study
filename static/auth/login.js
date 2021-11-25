function try_login() {
    let username = document.querySelector("input#username"),
        password = document.querySelector("input#password");

    let loginResult = document.querySelector("div#try_login_result");

    if (username.value == "")
        return
    if (password.value == "")
        return

    $.ajax({
        type: 'POST',
        url: "/login",
        dataType: "json",
        accept: "application/json",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            username: username.value,
            password: password.value,
        }),
        statusCode: {
            200: function (response) {
                console.log(response);
                loginResult.innerHTML = response.status;
                save_jwt(response.token);
                location.replace("/")
            },
            400: function (response) {
                console.log(response);
                loginResult.innerHTML = response.responseJSON.error;
            },
        },
        error: function (jqXHR, textStatus, errorThrown) {
            if (jqXHR.status != 400) {
                alert("캡쳐후 개발팀에 보여주세요" + errorThrown);
            }
        }
    });
}


function parse_jwt(token) {
    let base64Url = token.split('.')[1];
    let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    let jsonPayload = decodeURIComponent(atob(base64).split('').map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}

function unixtime_to_date(unixtime) {
    return new Date(unixtime * 1000);
}

function save_jwt(token) {
    let jwt = parse_jwt(token);
    // let exp_date = unixtime_to_date(jwt.exp);
    localStorage.removeItem("token")
    localStorage.setItem("token", token);
}