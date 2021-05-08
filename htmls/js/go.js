$(function () {
    $('.first_btn').click(function () {

        var a = $("input[name='selectedfile']:checked")[0].defaultValue;
        console.log("选中的radio的值是：" + a);

        if (a === 'NEWFILE') {
            a = $('#newfilename').val() + '.txt'
            console.log("选择新文件，文件名：", a)
        }

        console.log("输入的文件内容：", $('#words').val())

        var sendValue = a + "|!!..++|" + $('#words').val();

        $.ajax({
            url:'/api/send/'+sendValue,
            type:'GET',
            dataType:'json',
            contentType:"application/json",
            success:function(data){
                if( data.result == -990) {
                    alert('error！');
                } else {
                    alert(data.msg);
                }
            }
        });
    })
})
