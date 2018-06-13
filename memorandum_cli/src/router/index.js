import Vue from 'vue'
import Router from 'vue-router'
import taskUndone from '@/components/undone.vue'
import taskDone from '@/components/done.vue'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/undone',
      component: taskUndone
    },
    {
      path: '/done',
      component: taskDone
    },
    {  
        path:'*',  
        redirect:"/undone"  
    }  
  ]
})
