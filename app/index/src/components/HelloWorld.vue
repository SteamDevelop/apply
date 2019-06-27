<template>
  <div v-show="showControl">
    <el-input v-model="password" placeholder="password" clearable show-password></el-input><br/><br/>
    <el-button type="primary" size="mini" @click="sign">签到</el-button>
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
    sign: function () {
      if (localStorage.getItem("sign") === null) {
        ApplyService.sign({addr: localStorage.getItem("apply"), finger: "sign"}).then(data => {
          localStorage.setItem("sign", "signed");
        }).catch((err) => {
          alert("操作失败！", err);
        })
      }
    }
  },
  created() {
    if (localStorage.getItem("apply") === null) {
      alert("你还没有报名！");
      this.showControl = false;
    }
    if (localStorage.getItem("sign") !== null) {
      alert("你已经完成签到，请不要重复操作！");
      this.showControl = false;
    }
  }
}
</script>

<style>

</style>
