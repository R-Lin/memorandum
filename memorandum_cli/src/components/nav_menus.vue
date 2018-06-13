<template>
<div class='side-bar'>
  <div class='inline-level' v-for='nav in navMenu'>
    <div class='span-title'>{{nav.label}}</div>
    <div class='span-content' >
      <ul>
        <li v-for='(item, index) in nav.items'>
            <router-link :to='item.url'>{{item.label}}</router-link>
        </li>
      </ul>
    </div>
  </div>
</div>
</template>
<script>
export default {
  data () {
    var navMenu = [
      {
        label: '工作',
        items: [
          {
            url: 'undone',
            label: '所有工作'
          },
          {
            url: 'done',
            label: '已完成'
          }
        ]
      }
    ]
    return {
      navMenu: navMenu
    }
  }
}
</script>
<style>
  a{
    text-decoration:none;
    width: 100%;
  }
</style>
