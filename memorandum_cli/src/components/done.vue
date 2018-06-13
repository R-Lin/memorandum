<template>
<div id="content">
  <task-list :taskStatus="taskStatus" :button-items="buttonItems"></task-list>
</div>
</template>

<script>
  import taskList from '@/components/task_list.vue'
  export default {
    data () {
      return {
        taskStatus: '已完成',
        buttonItems: [
          { name: '已完成', className: "class-done"}
        ],
      }
    },
    components: {
      taskList
    }
  }
</script>
