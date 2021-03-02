layui.use(['table'], function () {
    let table = layui.table;

    table.render({
        elem: "#article-table",
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
    table.on('row(article-table)', function (row) {
        window.open("article/" + row.data.Id, "_blank");
    })
})

$('#search').keydown(function (e) {
    if (e.keyCode === 13) {
        q = $('#search').val().trim();
        filterArticles();
    }
})

let groupId = 0;
let q = ""

function setGroupId(id) {
    groupId = id;
    filterArticles();
}

function filterArticles() {
    layui.use('table', function () {
        let table = layui.table;
        table.reload('article-table', {
            where: {"groupId": groupId, "q": q}
        })
    })
}