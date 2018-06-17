<template>
  <div id="long-term-task">
      <div v-for="task in taskList"
        class="long-term-task_div">
        <div>
          <h3 :title="task.title">{{task.title}}</h3>
          <span>{{task.status}}</span>
        </div>
        <div class="task-desc">{{task.desc}}</div>
        <div class="long-term-task-edit">
          <i class="fa fa-plus-square" aria-hidden="true"></i>
        </div>
         
        <ol>
          <li 
            v-for="(step, index) in task.steps"
            v-if="index < 1 || 1"
            >
            
            <span class="step-task_span">
              {{fullWidth(task.steps.length - index)}}
            </span>
            <div class="step-task_desc">
              <div>
                {{getYMD(step.timestamp)}}
                <!--
                <i class="fa fa-window-close" aria-hidden="true"></i>
                -->
              </div>
              <div class="step-task-desc_div">
                {{step.desc}}
              </div>
              
            </div>
          </li>
          
        </ol>
      </div>

  </div>  
</template>
<script>
  export default {
    props: ["taskList"],
    data(){
      return {

      }
    },
    methods: {
      fullWidth: fullWidth,
      getYMD: getYMD
    }
  }
</script>