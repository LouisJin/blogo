layui.use(['element', 'form'], function () {
    let element = layui.element;
    let form = layui.form;
})

$('#search').keydown(function (e) {
    if (e.keyCode === 13) {
        window.location.href = "/?groupId=0&title=" + $('#search').val()
    }
})