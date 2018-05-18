/**
 * Created by hu on 2018/5/17.
 */
$(function () {

    //开始页面 
    $("#begin").bind("click",function () {

        if($("input[name='chatName']").val().trim()==""){
            alert("空输入");
            return false
        }
        $.ajax({
            type:"post",
            url:"http://localhost:3000",
            data:$("form#join").serializeArray(),
            contentType:"application/x-www-form-urlencoded",
            success:function (res) {
                var toast = ""
                switch (res){
                    case 0:
                        window.location.href="chat?name="+$("input[name='chatName']").val().trim()
                        break;
                    case 1:
                        alert("输入有问题，确定都不为空")
                        break;
                    case 2:
                        alert("名字已被占用，重置")
                        break;
                }
            },
            error:function () {
                alert("错误，请重试")
            }

        });


    })
    //进入聊天室刷新 长轮训 300毫秒刷新一次 修改址1000
    Refresh()
    var iv=setInterval("Refresh()",1000);

    //聊天页面
    $("#content").focus();
    $("button#submit").bind("click",function () {
        checkAndSendForm()
    });

    $("button#reset").bind("click",function () {
        $("#content").val("");
        $("#content").focus();
    });

    //关闭页面清除聊者身份 todo:有大问题 chrome 不能 triger

    window.onbeforeunload = function () {
        offline();
    };

});

function offline(){
        $.ajax({
            url:"/offline",
            type:"post",
            data:{"who":$("input[name='who']").val().trim()},
            success:function () {
            }
        })
}

function checkAndSendForm()  {
    var content=$("#content").val().trim();
    var tip= content.length !=0  ? (content.length > 20 ? "超长（20字限制）":"")  :  "空内容";
    if  (tip != ""){
        $("#content").focus();
        alert(tip);
    }else{
        $.ajax({
            url:"/chat",
            type:"post",
            data:$("form#chat").serializeArray(),
            contentType:"application/x-www-form-urlencoded",
            success:function ($data) {
                //请求成功 刷新界面
                Refresh();
                //聊天区域置底，编辑区域置空
                // $("div.chat_content_area").scrollTop=$("div.chat_content_area").scrollHeight
                //todo:暂时没发现怎么获取整个scroll的高度
                $("div.chat_content_area").scrollTop(100000);
                $("textarea#content").val("");

            },
            error:function () {
                alert("error")
            }
        })
    }
}


function Refresh() {
        $.ajax({
            url:"/refresh",
            success:function ($data) {
                //刷新页面

                var  records=[],joiners=[];
                $.each($data.Records,function (i,one) {
                    if(one.Who ==$("input[name='who']").val()){
                        one.Who="<span class='tip'>"+one.Who+"(我)</span>"
                    }
                    records.push("<li class=\"am-text-break am-comment\">\n" +
                        "                        <a href=\"#\">\n" +
                        "                            <img src=\"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnregyyDrMvhEDpfC4wFetzykulWRVMGF-jp7RXIIqZ5ffEdawIA\" alt=\"\" class=\"am-comment-avatar\" width=\"48\" height=\"48\">\n" +
                        "                        </a>\n" +
                        "                        <div class=\"am-comment-main\">\n" +
                        "                            <header class=\"am-comment-hd\">\n" +
                        "                                <div class=\"am-comment-meta\">\n" +
                        "                                    <a href=\"#\" class=\"am-comment-author\">"+one.Who+"</a>\n" +
                        "                                    <time datetime=\"\">"+stampToStand(one.Unix)+"</time>\n" +
                        "                                </div>\n" +
                        "                            </header>\n" +
                        "                            <div class=\"am-comment-bd\">\n" +
                        "                                <p>"+one.Content+"</p>\n" +
                        "                            </div>\n" +
                        "                        </div>\n" +
                        "                    </li>")
                })
                var recordEls = records.join('');
                tmpcontent = $(recordEls);
                $("ul#ulTag").empty().append(tmpcontent);

                $.each($data.Joiner,function (i,one) {

                    if(one==$("input[name='who']").val()){
                        one="<span class='tip'>"+one+"(我)</span>"
                    }

                    joiners.push("<div class=\"am-g\">\n" +
                        "                <div class=\"am-u-md-12\">\n" +
                        "                    <a class=\"am-icon-user am-margin-right-xs am-primary\"></a>"+one+"\n" +
                        "                </div>\n" +
                        "            </div>")
                })
                var joinerEls = joiners.join('');
                tmpJoiner = $(joinerEls);
                $("div#joinTag").empty().append(tmpJoiner);


            },
            error:function () {
                alert("出错了")
            }
        })
}

function stampToStand(stamp){
    return new Date(parseInt(stamp) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
}