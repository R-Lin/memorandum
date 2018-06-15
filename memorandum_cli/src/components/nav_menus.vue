<template>
<div class='side-bar'>
  <div class='inline-level' v-for='nav in navMenu'>
    <div class='span-title'>{{nav.label}}</div>
    <div class='span-content' >
      <ul>
        <li v-for='(item, index) in nav.items'>
            <i v-show="$route.path == '/' + item.url" class="fa item-chosen fa-arrow-right" aria-hidden="true"></i>
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
        label: '短期计划',
        items: [
          {
            url: 'short_term_undone',
            label: '所有计划'
          },
          {
            url: 'short_term_done',
            label: '已完成'
          }
        ]
      },
      {
        label: '长期计划',
        items: [
          {
            url: 'long_term_undone',
            label: '所有计划'
          },
          {
            url: 'long_term_done',
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
