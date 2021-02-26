layui.use('form', function () {

});

function login() {
    $.ajax({
        url: '/admin/login',
        method: 'post',
        data: $("form").serialize(),
        success(res) {
            if (res.code === 0) {
                window.location.href = "/admin";
            } else {
                layer.msg(res.msg, {time: 1000, offset: '80%'});
            }
        }, error(xhr, textStatus, errorThrown) {
            layer.msg("未知错误", {time: 1000, offset: '80%'});
        }
    })
}