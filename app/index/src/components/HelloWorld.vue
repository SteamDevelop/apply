<template>
  <div>
    <el-input v-model="password" placeholder="password" clearable show-password></el-input>
    <el-button type="primary" size="mini" @click="apply">生成地址并报名</el-button>
  </div>
</template>

<script>
import {ApplyService} from "../utils/utils.js";
export default {
  name: 'HelloWorld',
  data () {
    return {
      password: ""
    }
  },
  methods: {
    apply: function () {
      if (localStorage.getItem("sdfd") === null) {
        ApplyService.apply({password: this.password, finger:"sdfd"}).then(data => {
          localStorage.setItem("sdfd", data);
        }).catch((err) => {
          alert(err);
        });
      } else {
        alert("你已完成报名，请不要重复操作！");
      }
    }
  }
}
</script>

<style>

</style>
