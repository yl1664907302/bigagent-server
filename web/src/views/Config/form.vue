<template>
  <el-form :model="formData" :rules="rules" ref="formRef" label-width="120px" @submit.prevent>
    <!-- 标题 -->
    <el-form-item label="标题" prop="title">
      <el-input v-model="formData.title" placeholder="请输入标题" />
    </el-form-item>

    <!-- 备注 -->
    <el-form-item label="备注" prop="details">
      <el-input type="textarea" v-model="formData.details" placeholder="请输入备注" />
    </el-form-item>

    <!-- 主机 -->
    <el-form-item label="通信地址" prop="host">
      <el-input v-model="formData.host" placeholder="请输入主机" />
    </el-form-item>

    <!-- 认证模式 -->
    <el-form-item label="认证模式" prop="authMode">
      <el-select v-model="formData.auth_name" placeholder="请选择认证模式">
        <el-option label="token认证" value="token" />
      </el-select>
    </el-form-item>

    <!-- 数据名称 -->
    <el-form-item label="数据名称" prop="data_name">
      <el-select v-model="formData.data_name" placeholder="请输入数据名称">
        <el-option label="标准数据类型1" value="stand1" />
        <!--        <el-option label="标准数据类型2" value="stand2"></el-option>-->
        <!--        <el-option label="标准数据类型3" value="stand3"></el-option>-->
      </el-select>
    </el-form-item>

    <!-- 槽位 -->
    <el-form-item label="槽位名称" prop="slot_name">
      <el-select v-model="formData.slot_name" placeholder="请输入数据名称">
        <el-option label="槽位1" value="1" />
        <el-option label="槽位2" value="2" />
        <el-option label="槽位3" value="3" />
      </el-select>
    </el-form-item>

    <!-- Token (仅在token认证模式下显示) -->
    <el-form-item v-if="formData.auth_name === 'token'" label="Token" prop="token">
      <el-input v-model="formData.token" placeholder="请输入Token" />
    </el-form-item>

    <el-form-item label="上报与采集频次" prop="collection_frequency">
      <el-select v-model="formData.collection_frequency" placeholder="请输入频次">
        <el-option label="每秒一次" value="1s" />
        <el-option label="每秒二次" value="2s" />
        <el-option label="每秒三次" value="3s" />
      </el-select>
    </el-form-item>

    <!-- 范围 -->
    <!--    <el-form-item label="范围" prop="ranges">-->
    <!--      <el-select v-model="formData.ranges" placeholder="请选择范围">-->
    <!--        <el-option label="全部范围" value="全部范围"></el-option>-->
    <!--        <el-option label="部分范围" value="部分范围"></el-option>-->
    <!--      </el-select>-->
    <!--    </el-form-item>-->

    <!-- 按钮 -->
    <el-form-item>
      <el-button type="primary" @click="onSubmit">提交</el-button>
      <el-button @click="onReset">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { ref, watch } from 'vue'
import { addagentconf } from '@/api/login'
import { ElMessage } from 'element-plus'

// 表单验证规则
const rules = {
  title: [
    { required: true, message: '请输入标题', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符之间', trigger: 'blur' }
  ],
  details: [{ required: false, max: 200, message: '长度不能超过 200 个字符', trigger: 'blur' }],
  host: [
    { required: true, message: '请输入通信地址', trigger: 'blur' },
    {
      message: '请输入有效的grpc套接字地址，可选端口号(192.168.1.1:8080)',
      trigger: 'blur'
    }
  ],
  auth_name: [{ required: true, message: '请选择认证模式', trigger: 'change' }],
  data_name: [{ required: true, message: '请选择数据名称', trigger: 'change' }],
  solt_name: [
    { required: true, message: '请选择槽位', trigger: 'change' },
    { message: '请输入数据接收方配置存放槽位' }
  ],
  token: [{ required: true, message: '请输入Token', trigger: 'blur' }],
  ranges: [{ required: true, message: '请选择范围', trigger: 'change' }]
}

// 定义props接收父组件传值
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

// 定义emit向父组件发送事件
const emit = defineEmits(['update:modelValue', 'submit', 'reset'])

// 创建表单数据的响应式副本
const formData = ref({ ...props.modelValue })

// 监听props变化,同步更新表单数据
watch(
  () => props.modelValue,
  (newVal) => {
    formData.value = { ...newVal }
  },
  { deep: true }
)

// 表单ref
const formRef = ref(null)

// 提交方法
const onSubmit = async () => {
  if (!formRef.value) return
  try {
    // 表单验证
    await formRef.value.validate(async (valid) => {
      if (valid) {
        // 调用API发送数据
        const res = await addagentconf(formData.value)
        if (res.code === 0) {
          // 成功提示
          ElMessage.success('提交成功')
          onReset()
          // 触发父组件的submit事件
          emit('submit', formData.value)
        } else {
          ElMessage.error(res.msg || '提交失败')
        }
      }
    })
  } catch (error) {
    console.error('提交出错：', error)
    ElMessage.error('提交出错，请重试')
  }
}

// 重置方法
const onReset = () => {
  formRef.value?.resetFields()
  emit('reset')
}
</script>
