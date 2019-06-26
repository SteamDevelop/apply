<template>
  <div>
    <el-input v-model="password" placeholder="password" clearable show-password></el-input>
    <el-button type="primary" size="mini" @click="apply">生成地址并报名</el-button>
  </div>
</template>

<script>
import {utils} from "../utils/utils.js";
export default {
  name: 'HelloWorld',
  data () {
    return {
      password: ""
    }
  },
  methods: {
    apply: function () {
      utils.apply(this.password);
    }
  }
}
</script>

<style >

</style>
