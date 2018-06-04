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
		var date = getYMD(i.modifyTime);
		if(!record[date]){
			dateList.push(date);
			record[date] = [];
		}
		record[date].push(i);
	}
	dateList.sort();
	var statuList = ["进行中", "计划中", "已完成"];
	var statusClass = ["class-running", "class-plan", "class-done"];
	var priorityStar = {
		high: 3,
		normal: 2,
		low: 1
	}
	var addition;
	$("#runnging-task").empty();
	for(var index = dateList.length - 1; index >= 0; index--){
		for(var _item of record[dateList[index]]){
			var trItem;
			trItem = "<tr id='" +  _item.uuid + "'>";
			trItem += "<td title=" + _item.prioriry + ">" 
			for(var starN = 0; starN < (priorityStar[_item.prioriry] || 0); starN++){
				trItem += '<i class="fa fa-star priority-star"></i>'
			}
			trItem += "</td>";
			trItem += "<td>" + _item.taskType + "</td>";
			trItem += "<td>" + _item.taskDesc + "</td>";
			trItem += "<td>" + dateList[index] + "</td>";
			trItem += "<td title='切换任务状态' onclick='changeStatus(event)'>";

			// 生成状态切换按钮
			for(var s_index = 0; s_index < statuList.length; s_index++){
				if (statuList[s_index] == _item.status){
					addition = statusClass[s_index] + "' disabled ";
				}else{
					addition = "' ";
				}

				trItem += "<button id=" + _item.uuid + " class='stausSpan " + addition + " >" + statuList[s_index] + "</button>";
			}
			trItem += "</td>";
			trItem += "<td><button class='detele-task' onclick=deleteTask('" + _item.uuid + "')>删除任务</button></td>";
			trItem += "</tr>";
			$("#runnging-task").append(trItem);
		}
	}

}

function changeStatus(event){
	event.preventDefault();
	var comfirmResult = confirm("确定改状态么？");
	if(!comfirmResult){
		return;
	}
	var target = event.target;
	if(target.tagName.toLowerCase() !== "button"){
		return;
	}
	$.ajax({
		dataType: "json",
		url: "/change_status",
		type: "GET",
		data: {
			uuid: target.id,
			status: target.innerText
		},
		success: function(data){
			console.log("changeStatus", data);
			location.reload();
		}
	});
}

function deleteTask(uuid){
	var confirmResult = confirm("确定要删除吗 ？");
	if(!confirmResult){
		return;
	}
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