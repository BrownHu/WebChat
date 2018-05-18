<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>YOOZOO_CHAT</title>
    <link rel="stylesheet" href="http://cdn.amazeui.org/amazeui/2.7.2/css/amazeui.min.css">
    <link rel="stylesheet" href="static/css/common.css">
</head>
<body>
<!--群名区域-->
    <div class="am-g name_area_bg">
        <div class="am-u-md-12  am-padding-vertical-sm am-padding-horizontal-lg" >
            <span class="am-icon-users am-success am-margin-right-sm"></span>
            少三后端
        </div>
    </div>

    <!--content_frame-->
    <div class="am-g">
        <!--聊天内容显示、发送区域-->
        <div class="am-u-md-10 left_content_part">
            <!--内容显示-->
            <div class="am-g chat_content_area  am-padding">
                <ul class="am-comments-list am-comments-list-flip" id="ulTag">
                    <!--动态聊天框-->

                </ul>
            </div>
            <!--表情区域，用空白代替-->
            <div class="am-g mid_split">

            </div>
            <!--发送编辑框-->
            <div class="am-g edit_area">
                <form action="" id="chat">
                <filedset>
                    <div class="am-form-group">
                        <textarea name="content" id="content" cols="30" rows="10" class="am-padding-xs"></textarea>
                        <input type="hidden" name="who" value="{{ .who }}">
                    </div>
                </filedset>
                </form>
            </div>
            <!--发送底部-->
            <div class="am-g send_area am-padding-vertical-xs">
                <div class="am-u-md-2 am-u-md-offset-10 am-padding-vertical-sm ">
                    <button class="am-btn am-btn-primary am-btn-default am-btn-xs" id="reset" type="button">重置</button>
                    <button class="am-btn am-btn-primary am-btn-default am-btn-xs" id="submit" type="button">发送</button>
                </div>
            </div>
        </div>
    <!--在线列表显示-->
        <div class="am-u-md-2 right_join_list am-padding-top" id="joinTag">
            <!--动态列表位置-->
        </div>
    </div>

<script src="https://code.jquery.com/jquery-1.12.4.js" integrity="sha256-Qw82+bXyGq6MydymqBxNPYTaUXXq7c8v3CwiYwLLNXU=" crossorigin="anonymous"></script>
<script src="http://cdn.amazeui.org/amazeui/2.7.2/js/amazeui.js"></script>
<script src="http://cdn.amazeui.org/amazeui/2.7.2/js/amazeui.min.js"></script>
<script src="http://cdn.amazeui.org/amazeui/2.7.2/js/amazeui.ie8polyfill.js"></script>
<script src="http://cdn.amazeui.org/amazeui/2.7.2/js/amazeui.ie8polyfill.min.js"></script>
<script src="http://cdn.amazeui.org/amazeui/2.7.2/js/amazeui.widgets.helper.js"></script>
<script src="http://cdn.amazeui.org/amazeui/2.7.2/js/amazeui.widgets.helper.min.js"></script>
<script src="/static/js/letschat.js"></script>
</body>
</html>