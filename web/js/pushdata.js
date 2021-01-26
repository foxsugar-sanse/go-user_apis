var getJson = function (url) {
    return new Promise(function (resolve, reject) {
        var xhr = new XMLHttpRequest();
        xhr.open('get', url, true);
        xhr.responseType = 'json';
        xhr.onload = function () {
            var status = xhr.status;
            if (status === 200) {
                resolve(xhr.response);
            } else {
                reject(status)
            }
        }
        xhr.send();
    })
}

function ajax_submit() {
    var usernames = $("#oldPassword").val();//获取表单的输入值;
    var userpassword = $("#newPassword").val();//获取表单的输入值;
    $.ajax({
        type: "POST",  //数据提交方式（post/get）
        url: '/v1/user/new',  //提交到的url
        data: {"username":usernames,"password":userpassword},//提交的数据
        dataType: "json",//返回的数据类型格式
        success: function(msg){
            if (msg.success){  //修改成功
                //修改成功处理代码...
                document.write(msg[0].username)
            }else {  //修改失败
                //修改失败处理代码...
                document.write(msg[0].username)
            }
        }
    });
}

function ajax_fetch() {
    fetch('/v1/user/new',{
        method: 'post',
        body: new FormData(document.getElementById('username'), document.getElementById('password')),
    })
}

function fromPost(URL, PARAMS) {
    var temp = document.createElement("form");
    temp.action = URL;
    temp.method = "post";
    temp.style.display = "none";
    for (var x in PARAMS) {
        var opt = document.createElement("textarea");
        opt.name = x;
        opt.value = PARAMS[x];
        // alert(opt.name)
        temp.appendChild(opt);
    }
    document.body.appendChild(temp);
    temp.submit();
    return temp;
}