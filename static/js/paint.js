var mouserFlag =false;//鼠标是否可执行画图操作标识
var role="watcher";//角色
var context=  ($("#canvas")[0]).getContext("2d");
  $("#canvas").on("mousedown",function(e) {
      console.log("heheh");
      if (role!="draw"){
        // return false;
      }
      mouserFlag=true;
      context.beginPath();
      var x=e.clientX;
      var y=e.clientY;
      context.moveTo(x,y)
      gamesock.send("move:"+x+","+y);
    });
  $("#canvas").on("mouseup",function() {
        mouserFlag=false;
        console.log("leave");
        context.closePath();
    });
  $("#canvas").on("mouseout",function() {
          mouserFlag=false;
          console.log("leave");
          context.closePath();
  });
  $("#canvas").on("mousemove",function(e) {//画图
    if (!mouserFlag) {
        return false;
    }
    var x=e.clientX;
    var y=e.clientY;
    gamesock.send(x+","+y);
    context.lineTo(x,y);
    context.stroke()
  });
