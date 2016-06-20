var sock = null;
var wsuri = "ws://127.0.0.1:8080/hello";
window.onload = function() {

};
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
