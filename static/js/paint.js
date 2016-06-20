// $("#canvas").click(function() {
//     console.log("heheh");
//   });
var mouserFlag =false;
var context=  ($("#canvas")[0]).getContext("2d");
  $("#canvas").on("mousedown",function(e) {
      console.log("heheh");
      mouserFlag=true;
      context.moveTo(e.clientX,e.clientY)
      // context.moveTo(10,20)
    });
  $("#canvas").on("mouseup",function() {
        mouserFlag=false;
        // context.lineTo(10,30);
        // context.lineTo(10,40);
        // context.lineTo(10,50);

        console.log("leave");

        // context.closePath();
    });
  $("#canvas").on("mousemove",function(e) {
    if (!mouserFlag) {
        return false;
    }

    var x=e.clientX;
    var y=e.clientY;
    gamesock.send(x+","+y);
    context.lineTo(x,y);
    context.stroke()
  });
