// new Vditor('article', {
//
// })

Vditor.preview(document.getElementById('article'), $('#articleSrc').val(), {
    hljs: {
      lineNumber: true
    },
    theme: {
        current: "dark"
    }
});