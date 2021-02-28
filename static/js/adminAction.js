function loadINFO() {
    loadBlogs();
    loadTypes();
}
function loadBlogs(){
    var xhttp = new XMLHttpRequest();
    xhttp.responseType="json";
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myJson = this.response;
            if(myJson=="{}"){
                alert("空的jSON");
                document.getElementById("admin-blogs-table").innerHTML='';
            }else{
                var x="";
                document.getElementById("admin-blogs-table").innerHTML =  '';
                for(var i = 0;i<myJson.length;i++){
                    x +="<tr>" +
                        "<td>"+myJson[i].id+"</td>" +
                        "<td>"+myJson[i].title+"</td>" +
                        "<td>"+myJson[i].type+"</td>" +
                        "<td>"+myJson[i].recommend+"</td>" +
                        "<td>"+myJson[i].updateTime+"</td>" +
                        "<td>" +
                        "<button class=\"ui mini teal basic button btn_edit\" value=\""+ myJson[i].id +"\">编辑</button>" +
                        "<button class=\"ui mini red basic button btn_del\" value=\""+myJson[i].id+"\">删除</button>" +
                        "</td>" +
                        "</tr>"
                }
                document.getElementById("admin-blogs-table").innerHTML = x;
                //为按钮绑定事件
                bandAction();
            }
        }
    };
    xhttp.open("GET", "/info/blogs", true);
    xhttp.send();
}
//页面加载types
function loadTypes(){
    var xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.responseType="json";
    xmlHttpRequest.onreadystatechange=function(){
        var typeJson = this.response;
        if(typeJson=="{}"){
            alert("types没有数据")
        }else{
            var x="<div class=\"item\" data-value=\"\"></div>";
            var menu = document.getElementById("type_menu");
            menu.innerHTML = '';
            for(var i =0;i<typeJson.length;i++){
                x+="<div class=\"item\" data-value=\""+typeJson[i].id+"\">"+typeJson[i].name+"</div>"
            }
            menu.innerHTML=x;
        }
    }
    xmlHttpRequest.open("GET","/info/types",true);
    xmlHttpRequest.send();
}



function getSearchCondition(){
    var title = document.getElementById("title_bolgs").value;
    //type获取的是选项的value 对应的是相应的类型的id
    var type = document.getElementById("type_bolgs").value;
    var recommend = document.getElementById("recommend_bolgs").checked;
    console.log("出现了！！！");
    $.ajax({
        type:'POST',
        url: '/info/search',
        dataType:'json',
        contentType:'application/json',
        data:JSON.stringify({"title":title,"type":type,"recommend":recommend}),
        success: function (json) {
            console.log("success!!");
            var myJson = json;
            if(myJson=="{}"){
                alert("空的jSON");
                document.getElementById("admin-blogs-table").innerHTML='';
            }else{
                var x="";
                document.getElementById("admin-blogs-table").innerHTML =  '';
                for(var i = 0;i<myJson.length;i++){
                    x +="<tr>" +
                        "<td>"+myJson[i].id+"</td>" +
                        "<td>"+myJson[i].title+"</td>" +
                        "<td>"+myJson[i].type+"</td>" +
                        "<td>"+myJson[i].recommend+"</td>" +
                        "<td>"+myJson[i].updateTime+"</td>" +
                        "<td>" +
                        "<button class=\"ui mini teal basic button btn_edit\" value=\""+myJson[i].id+"\">编辑</button>" +
                        "<button class=\"ui mini red basic button btn_del\" value=\""+myJson[i].id+"\">删除</button>" +
                        "</td>" +
                        "</tr>"
                }
                document.getElementById("admin-blogs-table").innerHTML = x;
                //为按钮绑定事件
                bandAction();

            }
        },
        error: function () {
            console.log("fucking error...");
        }
    });
}


/**
 * 这里给信息列表中每个条目的按钮绑定事件，（undone）
 */
//给所有删除、修改button绑定事件
function bandAction(){
    var btn_del = document.getElementsByClassName("btn_del");
    var btn_edit = document.getElementsByClassName("btn_edit");

    for(var i = 0;i<btn_del.length;i++){
        btn_del[i].addEventListener("click",delAction);
        btn_edit[i].addEventListener("click",edAction);
    }
}

function delAction(){
    var value = this.value;
    var confirm_w = confirm("确定删除？");
    if(confirm_w==true){
        console.log(value);
        var xmlHttpRequest = new XMLHttpRequest();
        //利用拿到的json数据刷新相关标签内容
        xmlHttpRequest.onreadystatechange = function(){
            if (this.readyState == 4 && this.status == 200) {
                loadBlogs();
            }
        };
        xmlHttpRequest.open("GET","/info/del_blog?id="+value,true);
        xmlHttpRequest.send();
    }
}

function edAction(){
    var id = this.value;
    window.parent.location.href="/admin/edit?id="+id;
}