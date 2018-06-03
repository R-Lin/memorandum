$("#content").load("inline_content/doing")


// 将通过 serialize 获取的对象列表，拼接成一个对象
function formDataHandler(data){
	var result = {};
	for(var value of data){
		result[value.name] = value.value;
	}
	result["uuid"] = guid();

	return JSON.stringify(result);
}

// 根据状态获取记录列表
function getRecordsList(status, targetFunc){
	$.ajax({
		type: "GET",
		dataType: "json",
		data: {"status": status},
		url: "/getRecord",
		success: function(data){

			targetFunc(data)
		}
	})

}

function createRunningRecord(recordList){
	var record = {};
	var dateList = [];
	for(var i of recordList){
		var date = getYMD(i.createTime);
		if(!record[date]){
			dateList.push(date);
			record[date] = [];
		}
		record[date].push(i);
	}
	dateList.sort();
	var statuList = ["进行中", "计划中", "已完成"];
	var addition;
	$("#runnging-task").empty();
	for(var index = dateList.length - 1; index >= 0; index--){
		for(var _item of record[dateList[index]]){
			var trItem;
			trItem = "<tr id='" +  _item.uuid + "'>";
			trItem += "<td>" + _item.prioriry + "</td>";
			trItem += "<td>" + _item.taskType + "</td>";
			trItem += "<td>" + _item.taskDesc + "</td>";
			trItem += "<td>" + dateList[index] + "</td>";
			trItem += "<td title='切换任务状态' onclick='changeStatus(event)'>";

			// 生成状态切换按钮
			for(_statusSpan of statuList){
				if (_statusSpan == _item.status){
					addition = "currentStatus' disabled ";
				}else{
					addition = "' ";
				}

				trItem += "<input  type='button' id=" + _item.uuid + " class='stausSpan " + addition + " value='" + _statusSpan + "'/>";
			}
			trItem += "</td>";
			trItem += "<td><input type='button' onclick=deleteTask('" + _item.uuid + "') value='删除任务'></td>";
			trItem += "</tr>";
			$("#runnging-task").append(trItem);
		}
	}

}

function changeStatus(event){
	event.preventDefault();
	var target = event.target;
	if(target.tagName.toLowerCase() !== "input"){
		return;
	}
	$.ajax({
		dataType: "json",
		url: "/change_status",
		type: "GET",
		data: {
			uuid: target.id,
			status: target.value
		},
		success: function(data){
			console.log("changeStatus", data);
			location.reload();
		}
	});
}

function deleteTask(uuid){
	deleteElement(uuid)
	$.ajax({
		type: "GET",
		dataTye: "json",
		data: {"uuid": uuid},
		url: "/del_task",
		success: function(data){
			console.log(data);
		}
	})
}

function taskCleanAndHiden(){
	// 清空数据
	$("form div label input").each(function(i, value){
		$(value).val("")
	});
	$("#addTask").hide();
}

function deleteElement(id){
	$("#" + id).remove();
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

function listTask(value){
	getRecordsList(value, createRunningRecord);}

function fullWidth(value){
	return value < 10 ? "0" + value : value;
}

function S4() {
   return (((1+Math.random())*0x10000)|0).toString(16).substring(1);
}
function guid() {
   return (S4()+S4()+"-"+S4()+"-"+S4()+"-"+S4()+"-"+S4()+S4()+S4());
}