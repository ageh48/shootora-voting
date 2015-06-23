function showChart(vs) {
  var ctx = document.getElementById("chart").getContext("2d");
  var colors = ["#5AD3D1", "#F7464A", "#"]

  var data = [
    {
      value: vs[1].toFixed(2),
      color: "#46BFBD",
      highlight: "#5AD3D1",
      label: "参加してもよいと思う"
    },
    {
      value: vs[2].toFixed(2),
      color:"#F7464A",
      highlight: "#FF5A5E",
      label: "参加は許されない"
    },
    {
      value: vs[3].toFixed(2),
      color: "#FDB45C",
      highlight: "#FFC870",
      label: "参加しても良いが、3倍多く参加費を払うべきだ"
    },
    {
      value: vs[0].toFixed(2),
      color: "#B2B2B2",
      highlight: "#DEDEDE",
      label: "未投票"
    }
  ]
  var chart = new Chart(ctx).Pie(data);
  var helpers = Chart.helpers;
  var legendHolder = document.createElement('div');
  legendHolder.innerHTML = chart.generateLegend();
  helpers.each(legendHolder.firstChild.childNodes, function(legendNode, index){
    helpers.addEvent(legendNode, 'mouseover', function(){
      var activeSegment = chart.segments[index];
      activeSegment.save();
      chart.showTooltip([activeSegment]);
      activeSegment.restore();
    });
  });
  helpers.addEvent(legendHolder.firstChild, 'mouseout', function(){
    chart.draw();
  });
  document.getElementById("legend").appendChild(legendHolder.firstChild);
  //chart.chart.canvas.parentNode.parentNode.appendChild(legendHolder.firstChild);
}
