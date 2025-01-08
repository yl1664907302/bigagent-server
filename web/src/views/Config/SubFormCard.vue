<template>
  <el-card class="sub-form-card">
    <div class="card-title">{{subForm.data_name}}:{{subForm.slot_name  }}</div>
    <!-- 添加标题 -->

    <el-form-item label="认证模式" :prop="`${prefix}.auth_name`">
      <el-select v-model="subForm.auth_name" placeholder="请选择认证模式">
        <el-option label="token认证" value="token" />
        <el-option label="无" value="" />
      </el-select>
    </el-form-item>

    <el-form-item v-if="subForm.auth_name === 'token'" label="Token" :prop="`${prefix}.token`">
      <el-input v-model="subForm.token" placeholder="请输入Token" />
    </el-form-item>

    <el-form-item label="通信地址" :prop="`${prefix}.host`">
      <el-input v-model="subForm.host" placeholder="请输入主机" />
    </el-form-item>
  </el-card>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  subForm: {
    type: Object,
    required: true
  },
  prefix: {
    type: String,
    required: true
  }
})
</script>

<style scoped>
.card-title {
  font-size: 16px; /* 设置标题字体大小 */
  font-weight: bold; /* 设置标题字体加粗 */
  margin-bottom: 10px; /* 设置标题与内容之间的间距 */
}

.sub-form-card {
  padding: 20px; /* 增加卡片内边距 */
  border-radius: 8px; /* 设置圆角 */
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
}
</style>
