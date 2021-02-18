function loadINFO() {
    loadTags();
    loadTypes();
    bandAction();
}
//页面加载types
function loadTypes(){
    var xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.responseType="json";
    xmlHttpRequest.onreadystatechange=function(){
        var typeJson = this.response;
        var type = document.getElementById("type").value;
        var id = document.getElementById("id").value;
        if(typeJson=="{}"){
            alert("types没有数据");
        }else{
            var x="";
            var menu = document.getElementById("type_menu");
            menu.innerHTML = '';
            for(var i =0;i<typeJson.length;i++){
                if(type == typeJson[i].id && id !="0"){
                    var typeText = document.getElementById("type_text");
                    typeText.setAttribute("class","text");
                    typeText.innerText=typeJson[i].name;
                    x+="<div class=\"item active selected\" data-value=\""+typeJson[i].id+"\">"+typeJson[i].name+"</div>"
                }else{
                    x+="<div class=\"item\" data-value=\""+typeJson[i].id+"\">"+typeJson[i].name+"</div>";
                }
            }
            menu.innerHTML=x;
        }
    }
    xmlHttpRequest.open("GET","/info/types",true);
    xmlHttpRequest.send();
}

function loadTags(){
    var xmlHttpRequest = new XMLHttpRequest();
    xmlHttpRequest.responseType="json";
    xmlHttpRequest.onreadystatechange=function(){
        var tagJson = this.response;
        var tag = document.getElementById("tags").value;
        if(tagJson=="{}"){
            alert("types没有数据");
        }else{
            var x="";
            var menu = document.getElementById("tags_menu");
            var id = document.getElementById("id").value;
            menu.innerHTML = '';
            for(var i =0;i<tagJson.length;i++){
                if(tag == tagJson[i].id && id != "0"){
                    var tagText = document.getElementById("tag_text");
                    tagText.setAttribute("class","text");
                    tagText.innerText=tagJson[i].name;
                    x+="<div class=\"item active selected\" data-value=\""+tagJson[i].id+"\">"+tagJson[i].name+"</div>"
                }else{
                    x+="<div class=\"item\" data-value=\""+tagJson[i].id+"\">"+tagJson[i].name+"</div>"
                }
            }
            menu.innerHTML=x;
        }
    }
    xmlHttpRequest.open("GET","/info/tags",true);
    xmlHttpRequest.send();
}

function bandAction(){
    var btn_save = document.getElementById("btn-save");
    var btn_publish = document.getElementById("btn-publish");
    btn_publish.addEventListener('click',submit);
    btn_save.addEventListener('click',submit);
}

//保存
function submit(){
    var id = document.getElementById("id").value;
    var flag = document.getElementById("flag").value;
    var title = document.getElementById("title").value;
    var description = document.getElementById("description").value;
    var content = document.getElementById("content").value;
    var type = document.getElementById("type").value;
    var tags = document.getElementById("tags").value;
    var first_p = document.getElementById("first_p").value;
    var recommend = document.getElementById("recommend").checked;
    var shareInfo = document.getElementById("shareInfo").checked;
    var appreciation = document.getElementById("appreciation").checked;
    var comment = document.getElementById("comment").checked;
    var published = this.value;
    $.ajax({
        type: 'POST',
        url: '/blog/edit',
        dataType: 'json',
        contentType: 'application/json',
        data: JSON.stringify({"id":id,
                                    "flag":flag,
                                    "title":title,
                                    "description":description,
                                    "content":content,
                                    "type":type,
                                    "tags":tags,
                                    "first_p":first_p,
                                    "recommend":recommend,
                                    "shareInfo":shareInfo,
                                    "appreciation":appreciation,
                                    "comment":comment,
                                    "published":published}),
        success: function (info) {
            alert(info.message);
        },
        error: function () {
            console.log("fucking error...")
        }
    });
}