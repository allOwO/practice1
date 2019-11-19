<template>
  <div>
    <div id="app">
      <div class="am-tabs am-margin" data-am-tabs>
        <ul class="am-tabs-nav am-nav am-nav-tabs">
          <li>
            <router-link to="/add">添加</router-link>
          </li>
          <li>
            <router-link to="/change">修改</router-link>
          </li>
        </ul>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'App'
}
</script>
<style>
  @import "assets/css/amazeui.min.css";
  @import "assets/css/admin.css";
</style>
