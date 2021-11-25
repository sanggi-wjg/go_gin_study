$("a#try_logout").click(function () {
    $.ajax({
        type: 'GET',
        url: "/logout",
        dataType: "json",
        accept: "application/json",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({}),
        statusCode: {
            200: function (response) {
                console.log(response);
                location.replace("/login")
            },
        },
        error: function (jqXHR, textStatus, errorThrown) {
            if (jqXHR.status != 400) {
                alert("캡쳐후 개발팀에 보여주세요" + errorThrown);
            }
        }
    });
});