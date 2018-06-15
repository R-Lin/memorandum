import Vue from 'vue'
import Router from 'vue-router'
import taskUndone from '@/components/undone.vue'
import taskDone from '@/components/done.vue'
import long_term_undone from '@/components/long_term_undone.vue'
import long_term_done from '@/components/long_term_done.vue'
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
      path: '/long_term_undone',
      component: long_term_undone
    },
    {
      path: '/long_term_done',
      component: long_term_done
    },
    {  
        path:'*',  
        redirect:"/short_term_undone"  
    }  
  ]
})
