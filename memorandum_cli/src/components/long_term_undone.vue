<template>
  <div>
    <button 
      @click="longTermTask=!longTermTask"
      style="background-color: rgb(216, 188, 43);">新建长期任务</button>
    <div v-show="longTermTask">
      <label>标题： <input type="text" placeholer="请输入标题"></label>
      <label>简要： <input type="text" placeholder="一句话说明次任务"></label>
      <label for="">
        马上开始：
        <select name="" id="">
          <option value=true>是</option>
          <option value=false>否</option>
        </select>

      </label>
      
    </div>
    <hr>
    <long-term-task-list 
      :taskList="taskList"
      ></long-term-task-list>
  </div>

</template>

<script>
  import longTermTaskList from "@/components/long_term_task_list.vue"
  export default {
    data(){
      return {
        longTermTask: 0,
        taskList: [
          {
            title: '测试标题1',
            desc: 'Just test Just test Just test Just test Just ' + 
                  'test Just test \nJust test Just test Just test \
                  1. asdasd\n \
                  2. asdasd\n',
            status: '计划中',
            steps: [
              {timestamp: 1529139543, desc: "test1"},
              {timestamp: 1529139540, desc: "test12"},
            ]
          },
          {
            title: 'The test',
            desc: 'test Just test \nJust test Just test Just test \
                  1. asdasd \
                  2. asdasd',
            status: '进行中',
            steps: [
              {timestamp: 1529139543, desc: "test1"},
               {timestamp: 1529139540, desc: "test12"},
               {timestamp: 1529139540, desc: "test12"},
               {timestamp: 1529139540, desc: "test12"},
            ]
          },
           {
            title: '测试标题1123123123123123131231231测试测试测测试asdadadad',
            desc: 'Just test\n\n Just test Just test Just test Just test Just test Just test Just test Just test ',
            status: '进行中',
            steps: [
              {timestamp: 1529139543, desc: "test1"},
               {timestamp: 1529139540, desc: "test12"},
               {timestamp: 1529139540, desc: "test12"},
            ]
          }
        ]

      }
    },
    methods: {
     
    },
    components: {
      longTermTaskList
    }
  }

</script>