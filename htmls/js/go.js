$('.first_btn').click(function () {
    var words = $('#words').val()
    $.ajax({
        url:'/api/send/'+words,
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