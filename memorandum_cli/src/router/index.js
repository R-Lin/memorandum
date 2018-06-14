import Vue from 'vue'
import Router from 'vue-router'
import taskUndone from '@/components/undone.vue'
import taskDone from '@/components/done.vue'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/short_term_undone',
      component: taskUndone
    },
    {
      path: '/short_term_done',
      component: taskDone
    },
    {  
        path:'*',  
        redirect:"/short_term_undone"  
    }  
  ]
})
