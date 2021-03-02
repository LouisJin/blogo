layui.use(['form', 'table'], function () {
    let table = layui.table;

    table.render({
        elem: "#admin-article-table",
        limit: 20,
        page: true,
        cols: [[
            {field: 'Title', title: '标题', unresize: true, align: 'center', style: 'color: white;'},
            {
                field: 'Group.Name',
                title: '分类',
                width: 130,
                unresize: true,
                align: 'center',
                style: 'color: white;',
                templet: function (row) {
                    return row.Group.Name;
                }
            },
            {field: 'ThumbsupNum', title: '点赞', width: 70, unresize: true, align: 'center', style: 'color: white;'},
            {
                field: 'CommentNum',
                title: '评论',
                width: 70,
                unresize: true,
                align: 'center',
                style: 'color: white;',
                templet: function (row) {
                    if (row.IsComment === 0) {
                        return "X";
                    } else {
                        return row.CommentNum;
                    }
                }
            },
            {
                field: 'UpdateTime',
                title: '编辑时间',
                width: 110,
                unresize: true,
                align: 'center', style: 'color: white;',
                templet: function (row) {
                    return row.UpdateTime.substr(0, 10);
                }
            },
        ]],
        url: "article/api/list",
        request: {
            limitName: "size"
        },
        parseData: function (res) {
            return {
                "code": res.code,
                "msg": res.msg,
                "data": res.data.list,
                "count": res.data.count,
            }
        }
    });
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

