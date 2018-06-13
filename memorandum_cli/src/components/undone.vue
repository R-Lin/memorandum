<template>
<div id="content">
  <button style="background-color:#d8bc2b" @click="_initFormData()">新建任务</button>
  <form action="" method="POST" id="addTask" v-show="showForm">
    <div>
      <label>
        优先级别：
        <select v-model="formData.priority">
          <option v-for="item in priorityList" :value="item.value">{{item.name}}</option>
        </select>
      </label>
    </div>
    <div>
      <label>
        任务类别：
        <input 
          type="text" 
          autofocus="" 
          placeholder="请输入任务类别" 
          required="1" 
          v-model="formData.taskType" 
          autocomplete="off">
      </label>
    </div>
    <div>
      <label>
        任务描述：
        <input 
          style="width:400px" 
          type="text" 
          placeholder="请输入任务描述" 
          required="" 
          v-model="formData.taskDesc" 
          autocomplete="off">
      </label>        
    </div>
    <div>
      <label>
        马上开始：
        <select v-model="formData.isStart">
          <option v-for="item in isStartList" :value="item.value">{{item.name}}</option>
        </select>
      </label>
    </div>
    <input type="button" value="取消任务" @click="taskCleanAndHiden()">
    <input 
      type="button" 
      @click="addTask()" 
      style="background-color: #c37574; color: #fff; border-radius: 4px" 
      value="添加任务">
  </form>
  <task-list 
    :taskStatus="taskStatus" 
    :show-filter="showFilter" 
    ref="taskList"
    :button-items="buttonItems">
    </task-list>
</div>
</template>

<script>
  import taskList from '@/components/task_list.vue'
  import axios from 'axios'
  import Qs from 'qs'
  export default {
    components: {
      taskList
    },
    data () {
      return {
        showForm: false,
        taskStatus: 'undone',
        showFilter: true,
        buttonItems: [
          { name: '进行中', className: "class-running"},
          { name: '计划中', className: "class-plan"},
          { name: '已完成', className: "class-done"},
        ],
        priorityList: [
          {value: 1, name: '贼低'},
          {value: 2, name: '一般'},
          {value: 3, name: '贼高'},
        ],
        isStartList: [
          {value: true, name: '是'},
          {value: false, name: '否'},
        ],
        initFormData: {
          priority: 1,
          taskType: '',
          taskDesc: '',
          isStart: false
        },
        formData: {}
      }
    },
    methods: {
      _initFormData: function(){
        // 深拷贝
        this.formData = JSON.parse(JSON.stringify(this.initFormData));
        this.showForm = true;
      },
      S4: function() {
         return (((1+Math.random())*0x10000)|0).toString(16).substring(1);
      },
      guid: function () {
         var S4 = this.S4;
         return (S4()+S4()+"-"+S4()+"-"+S4()+"-"+S4()+"-"+S4()+S4()+S4());
      },
      checkFormVaild: function(formData){
        for(var key of Object.keys(formData)){
          if(formData[key] === undefined || 
             formData[key] === null ||
             formData[key] === ""){
            
            return false;
          }
          
        }
        return true;
        

      },
      addTask: function(){
        if(!this.checkFormVaild(this.formData)){
          alert("输入框不能为空")
          return;
        }
        this.formData.uuid = this.guid();
        this.showForm = false;
        axios.post('http://localhost:9999/add_task', this.formData, {
                      headers: {
                            'Content-Type': 'application/x-www-form-urlencoded'
                      }
                  })
        .then(res =>{
          this.$refs.taskList.getTasks();
        }).catch(function(err){
          console.log(123123, err)  
        })
      },
      taskCleanAndHiden: function(){
        this.showForm = false;
      }
    }
  }
</script>