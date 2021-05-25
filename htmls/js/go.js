$(function () {
    $('.first_btn').click(function () {
        var contents = $('#words').val();

        console.log(contents)

        if( contents ===  ''){
            alert('没有输入内容！清重新输入');
            return
        }

        var a = $("input[name='selectedfile']:checked")[0].defaultValue;
        console.log("选中的radio的值是：" + a);

        if (a === 'NEWFILE') {
            a = $('#newfilename').val();
            // 如果是新文件，文件名不能为空
            if( a === '' ){
                alert('文件名不能为空');
                return
            }
            a = a + '.txt'
            console.log("选择新文件，文件名：", a)
        }

        console.log("输入的文件内容：", $('#words').val())

        var _data = {
            'msgCode': '2',
            'msgValue': a,
            'msgBody': contents,
        };

        $.ajax({
            url: '/api/send',
            type: 'POST',
            dataType: 'json',
            contentType: "application/json",
            data: JSON.stringify(_data),
            success: function (data) {
                if(data.result==1){
                    alert("保存成功！！！！")
                    window.location.href='/api/index';
                }else{
                    if ( data.result == -2 ) {
                        alert('保存出错!')
                    } else if (data.result == -989) {
                        alert('同编号的房间已经存在，请重新输入！')
                    } else {
                        alert(data.result)
                    }
                }
            }
        });

    })
})
