<template>
<div class="runnging-task-block">
  <label style="margin: 0 0 0 5px;" v-show="showFilter">任务状态：
    <select onchange="listTask(value)">
      <option value="all">全部</option>
      <option value="计划中">计划中</option>
      <!-- <option value="已完成">已完成</option> -->
      <option value="进行中">进行中</option>
    </select>
  </label>
  <div class="runnging-task-item">
    <table>
      <thead>
        <tr>
          <th>优先级</th>
          <th>类别</th>
          <th style="width:40%">任务描述</th>
          <th>创建日期</th>
          <th>状态</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody id="runnging-task">
        <tr v-for="task in tasks" :id="task.uuid">
          <td :title="task.priority">
            <i v-for="cnt in task.priority" class="fa fa-star priority-star"></i>
            
          
          </td>
          <td>{{task.taskType}}</td>
          <td>{{task.taskDesc}}</td>
          <td>{{task.modifyTime}}</td>
          <td title="切换任务状态" @click="changeStatus($event)">
            
            <button v-for="buttonItem in buttonItems"
              class="stausSpan"
              :data-uuid="task.uuid"
              :class="[buttonItem.name == task.status ? buttonItem.className : '']"
              :disabled="buttonItem.name == task.status"
              >{{buttonItem.name}}</button>
          </td>
          <td>
            <button class="detele-task" @click="deleteTask(task.uuid)">删除任务</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
</div>
</template>
<script>
import axios from 'axios'
import querystring from 'querystring'
export default {
  props: ['showFilter', 'taskStatus', 'buttonItems'],
  data () {
    return {
      tasks: [],
      statsButton: [
        { name: '已完成', className: "class-done"},
        { name: '进行中', className: "class-running"},
        { name: '计划中', className: "class-plan"},
      ]
    }
  },
  methods: {
    changeStatus: function(event){
      var target = event.target;
      if(target.tagName.toLowerCase() !== 'button'){
        return;
      }
      event.stopPropagation();
      var comfirmResult = confirm("确定改状态么？");
      if(!comfirmResult){
        return;
      }
      var params = {
        params: {
          uuid: target.dataset.uuid,
          status: target.innerText
        }
      }
      axios.get('http://localhost:9999/change_status', params).then(res =>{
            this.getTasks();
        }).catch(function(err){
            console.log(123123, err)
        })
      },
    
    deleteTask: function (uuid){
      var confirmResult = confirm("确定要删除吗 ？");
      if(!confirmResult){
        return;
      }
      var params = {
        params: {
          uuid: uuid,
        }
      }
      axios.get('http://localhost:9999/del_task', params).then(res =>{
            this.getTasks();
        }).catch(function(err){
            console.log(123123, err)
        })
    },
    
    getTasks: function(){
      var params = {
        params: {
          status: this.taskStatus
        }
      }
      axios.get('http://localhost:9999/getRecord', params).then(res =>{
          for(var item of res.data){
            item['modifyTime'] = getYMD(item['modifyTime']);
          }
          this.tasks = res.data;
        }).catch(function(err){
          console.log(123123, err)  
        })
    }
  },
  created: function(){
    this.getTasks();
  }
}
</script>