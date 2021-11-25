/* $(".tab-wizard").steps({
    headerTag: "h5",
    bodyTag: "section",
    transitionEffect: "fade",
    titleTemplate: '<span class="step">#index#</span> #title#',
    labels: {
        finish: "Submit"
    },
    onStepChanged: function (event, currentIndex, priorIndex) {
        $('.steps .current').prevAll().addClass('disabled');
    },
    onFinished: function (event, currentIndex) {
        $('#success-modal').modal('show');
    }
});

$(".tab-wizard2").steps({
    headerTag: "h5",
    bodyTag: "section",
    transitionEffect: "fade",
    titleTemplate: '<span class="step">#index#</span> <span class="info">#title#</span>',
    labels: {
        finish: "Submit",
        next: "Next",
        previous: "Previous",
    },
    onStepChanged: function (event, currentIndex, priorIndex) {
        $('.steps .current').prevAll().addClass('disabled');
    },
    onFinished: function (event, currentIndex) {
        try_join();
    }
}); */

function on_chnage_confirm_password() {
    $("input#password, input#password_confirm").on("propertychange change keyup paste input", function () {
        let password = document.querySelector("input#password"),
            password_confirm = document.querySelector("input#password_confirm");
        let join_result = document.querySelector("div#join_result")

        if (password.value == password_confirm.value) {
            join_result.innerHTML = "비밀번호가 동일합니다."
        }
        else {
            join_result.innerHTML = "비밀번호가 불일치합니다."
        }
    });
}


function try_join() {
    let username = document.querySelector("input#username"),
        password = document.querySelector("input#password"),
        password_confirm = document.querySelector("input#password_confirm"),
        userlevel = document.querySelector("select#userlevel");

    let join_result = document.querySelector("div#join_result");

    if (username.value == "")
        return
    if (password.value == "")
        return
    if (password_confirm.value == "")
        return
    if (password.value != password_confirm.value)
        return
    if (!userlevel.value in ['1', '2', '3'])
        return

    $.ajax({
        type: 'POST',
        url: "/newuser",
        dataType: "json",
        accept: "application/json",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            username: username.value,
            password: password.value,
            userlevel: parseInt(userlevel.value),
        }),
        statusCode: {
            200: function (response) {
                console.log(response);
                $('#success-modal-btn').trigger('click');
            },
            400: function (response) {
                // console.log(response);
                join_result.innerHTML = response.responseJSON.error;
            },
        },
        error: function (jqXHR, textStatus, errorThrown) {
            if (jqXHR.status != 400) {
                alert("캡쳐후 개발팀에 보여주세요" + errorThrown);
            }
        }
    });
}
