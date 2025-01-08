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

    <el-form-item>
      <el-row gutter="20">
        <el-col :span="8">
          <el-button
            type="warning"
            @click="addSubForm('stand1')"
            :disabled="subForms.filter((subForm) => subForm.data_name === 'stand1').length >= 3"
            style="width: 100%"
            >新增标准类型1</el-button
          >
        </el-col>
        <el-col :span="8">
          <el-button
            type="warning"
            @click="addSubForm('stand2')"
            :disabled="subForms.filter((subForm) => subForm.data_name === 'stand2').length >= 3"
            style="width: 100%"
            >新增标准类型2</el-button
          >
        </el-col>
        <el-col :span="8">
          <el-button
            type="warning"
            @click="addSubForm('stand3')"
            :disabled="subForms.filter((subForm) => subForm.data_name === 'stand3').length >= 3"
            style="width: 100%"
            >新增标准类型3</el-button
          >
        </el-col>
      </el-row>
      <el-button type="danger" @click="removeSubForm" style="margin-top: 10px; width: 100%">
        减少接收模块
      </el-button>
    </el-form-item>

    <!-- 小表单列表 -->
    <el-scrollbar style="max-height: 400px; overflow-y: auto">
      <div v-for="(subForm, index) in subForms" :key="index" style="margin-bottom: 20px">
        <SubFormCard :subForm="subForm" :prefix="'subForms.' + index" />
      </div>
    </el-scrollbar>

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
import SubFormCard from '@/views/Config/SubFormCard.vue'
import { ElScrollbar } from 'element-plus'

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
  slot_name: [
    { required: true, message: '请选择槽位', trigger: 'change' },
    { message: '请输入数据接收方配置存放槽位' }
  ],
  token: [{ required: true, message: '请输入Token', trigger: 'blur' }],
  ranges: [{ required: true, message: '请选择范围', trigger: 'change' }],
  collection_frequency: [{ required: true, message: '请选择上报与采集频次', trigger: 'change' }]
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

// 存储小表单数据的数组
const subForms = ref([])

// 添加小表单的方法
const addSubForm = (key) => {
  const typeCount = subForms.value.filter((subForm) => subForm.data_name === key).length

  if (typeCount < 3) {
    subForms.value.push({
      data_name: key,
      auth_name: '',
      token: '',
      slot_name: typeCount + 1,
      host: ''
    })
  } else {
    ElMessage.warning(`${key} 类型最多只能添加三个`)
  }
}

// 减少小表单的方法
const removeSubForm = () => {
  if (subForms.value.length > 0) {
    subForms.value.pop()
  } else {
    ElMessage.warning('没有更多模块可以减少')
  }
}

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
