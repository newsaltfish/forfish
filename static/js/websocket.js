// 聊天信息
var sock = null;
var wsuri = "ws://127.0.0.1:8081/hello";
$("#send").click(function () {
  var msg= $("#msg").val();
  sock.send(msg);
});
$(".button").hide();
$("#nameBottun").click(function () {
  $(".namediv").hide();
  $(".button").show();
  console.log("onload");
  sock = new WebSocket(wsuri);
  sock.onopen = function() {
    console.log("connected to " + wsuri);
    sock.send($("#name").val());
  }
  sock.onclose = function(e) {
    console.log("connection closed (" + e.code + ")");
  }
  sock.onmessage = function(e) {
    console.log("message received: " + e.data);
    var html='<p></p>'+e.data+'<br>'
    $("#text").append(html);
  }
});
// 图像信息
var gamesock=new WebSocket("ws://127.0.0.1:8081/game")
gamesock.onopen = function() {
console.log("connected to ws://127.0.0.1:8081/game");
};
$("#role").click(function () {
gamesock.send("draw");
});
  // console.log("dengl");
  gamesock.onclose = function(e) {
    console.log("connection closed (" + e.code + ")");
  }
  gamesock.onmessage = function(e) {
    // console.log(e.data);
    if ((e.data).indexOf("move:")>-1) {
      var loc = ((e.data).split(":")[1]).split(",");
      context.moveTo(loc[0],loc[1]);
    }else{
      var loc = (e.data).split(",");
      context.lineTo(loc[0],loc[1]);
    }
    context.stroke()
  }
