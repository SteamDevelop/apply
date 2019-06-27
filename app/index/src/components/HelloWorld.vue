<template>
  <div v-show="showControl">
    <el-input v-model="password" placeholder="password" clearable show-password></el-input><br/><br/>
    <el-button type="primary" size="mini" @click="apply">生成地址并报名</el-button>
  </div>
</template>

<script>
import {ApplyService} from "../utils/utils.js";
export default {
  name: 'HelloWorld',
  data () {
    return {
      password: "",
      showControl: true,
    }
  },
  methods: {
    apply: function () {
      if (localStorage.getItem("apply") === null) {
        ApplyService.apply({password: this.password, finger:"apply"}).then(data => {
          localStorage.setItem("apply", data.getAddr());
        }).catch((err) => {
          alert("操作失败！", err);
        });
      }
    }
  },
  created() {
    if (localStorage.getItem("apply") !== null) {
      alert("你已经完成报名，请不要重复操作！");
      this.showControl = false;
    }
  }
}
</script>

<style>

</style>
