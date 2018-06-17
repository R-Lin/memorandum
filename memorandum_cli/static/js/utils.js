function fullWidth(value){
  return value < 10 ? "0" + value : value;
}

function getYMD(timeStamp){
  var num2Week = {
    0: "Sun",
    1: "Mon",
    2: "Tue",
    3: "Wed",
    4: "Thu",
    5: "Fri",
    6: "Sat"
  }
  var timeObj = new Date(timeStamp * 1000);
  var result = [
    timeObj.getFullYear(),
    fullWidth(timeObj.getMonth() + 1),
    fullWidth(timeObj.getDate()),
    
  ];
  return result.join("-") + " " + num2Week[timeObj.getDay()];
}