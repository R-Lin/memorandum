<template>
 <main class="main" id="app">
    <!-- 左边导览航 -->
    <nav-menus></nav-menus>
    <!-- 右部主内容区 -->
    <div class="main-content">
      <div id="content">
          <router-view/>
      </div>
    </div>
  </main>
</template>
<script>
import navMenus from '@/components/nav_menus.vue' 
export default {
  name: 'App',
  components: {
    navMenus
  }
}
</script>
