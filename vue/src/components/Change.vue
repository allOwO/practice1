<template>
  <div class="am-tab-panel am-fade am-in am-active" id="change">
    <form class="am-form am-form-inline">
      <div class="am-g am-margin-top">
        <div class="col-sm-2 am-text-right">邮箱</div>
        <div class="col-sm-4">
          <input type="text" class="am-input-sm" v-model="mess.user_mail">
        </div>
        <div class="col-sm-4 col-end">
          <button type="button" class="am-btn am-btn-primary am-btn-xs" @click="check()">查询</button>
        </div>
      </div>
      <div class="am-g am-margin-top">
        <div class="col-sm-2 am-text-right">
          用户名
        </div>
        <div class="col-sm-4 col-end">
          <input type="text" class="am-input-sm" v-model="mess.user_name">
        </div>
        <!--<div class="col-sm-4">-->
        <!--&lt;!&ndash;{{mess.user_name}}&ndash;&gt;-->
        <!--</div>-->
      </div>

      <div class="am-g am-margin-top">
        <div class="col-sm-2 am-text-right">
          手机
        </div>
        <div class="col-sm-4 col-end">
          <input type="text" class="am-input-sm" v-model="mess.user_phone">
        </div>
        <!--<div class="col-sm-4">-->
        <!--&lt;!&ndash;{{mess.user_phone}}&ndash;&gt;-->
        <!--</div>-->
      </div>
      <div class="am-g am-margin-top">
        <div class="col-sm-2 am-text-right">小组</div>
        <div class="col-sm-10">
          <div class="am-btn-group" data-am-button>
            <label class="">
              <input type="checkbox" value="system_user" v-model="mess.groups"> 管理员
            </label>
            <label class="">
              <input type="checkbox" value="service_staff" v-model="mess.groups"> 客服人员
            </label>
            <label class="">
              <input type="checkbox" value="worker" v-model="mess.groups"> 运营
            </label>
          </div>
        </div>
      </div>
    </form>
    <div class="am-margin">
      <button type="button" class="am-btn am-btn-primary am-btn-xs" @click="onSubmit()">提交保存</button>
    </div>
  </div>
</template>
<script>
import axios from 'axios'

export default {
  name: 'change',
  data () {
    return {
      mess: {
        user_name: '',
        user_phone: '',
        user_mail: '',
        groups: []
      }
    }
  },
  methods: {
    onSubmit () {
      console.log(this.mess)
      // let formData =JSON.stringify(this.mess)
      axios({
        method: 'post',
        url: 'http://localhost:8000/changeuser',
        data: this.mess
      }).then((res) => {
        console.log(res)
        if (res.data.code === 200) {
          alert('修改成功')
        }
        if (res.data.code === 300) {
          alert(res.data.msg)
        }
      })
    },
    check () {
      axios({
        method: 'get',
        url: 'http://localhost:8000/checkuser',
        params: {
          'mail': this.mess.user_mail
        }
      }).then((res) => {
        console.log(res.data.data.user_name)
        console.log(res.data.data.user_phone)
        console.log(res.data.data.user_mail)
        console.log(res.data.data.groups)
        if (res.data.code === 200) {
          this.mess.user_name = res.data.data.user_name
          this.mess.user_phone = res.data.data.user_phone
          this.mess.user_mail = res.data.data.user_mail
          this.mess.groups = res.data.data.groups
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
