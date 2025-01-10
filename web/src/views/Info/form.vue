<template>
  <div class="config-form">
    <!-- 已选主机信息 -->
    <div class="selected-hosts-section">
      <h4>已选主机</h4>
      <div class="selected-hosts">
        <el-tag v-for="host in selectedHosts" :key="host.uuid" class="host-tag" type="info">
          {{ host.hostname }} ({{ host.ipv4_first }})
        </el-tag>
      </div>
    </div>

    <!-- 配置选择表格 -->
    <div class="config-section">
      <h4>选择配置</h4>
      <el-table v-loading="tableLoading" :data="configList" border style="width: 100%">
        <el-table-column width="55" label="选择">
          <template #default="{ row }">
            <el-radio
              v-model="selectedConfigId"
              :label="row.id"
              @change="handleSelectionChange(row)"
            >
              <!-- 空标签用于仅显示单选框 -->
              <span></span>
            </el-radio>
          </template>
        </el-table-column>
        <el-table-column prop="id" label="ID" width="60" />
        <el-table-column prop="title" label="配置标题" min-width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === '生效中' ? 'success' : 'warning'">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-area">
        <el-pagination
          v-model:current-page="pagination.currentPage"
          v-model:page-size="pagination.pageSize"
          :page-sizes="[5, 10, 20, 50]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>

    <!-- 添加按钮区域 -->
    <div class="button-area">
      <el-button @click="handleCancel">取消</el-button>
      <el-button type="primary" @click="handleSubmit">确定</el-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { getagentconf, pushagentconfbyhost } from '@/api/login' // 假设这是你的API

interface Host {
  uuid: string
  hostname: string
  ipv4_first: string
}

interface ConfigItem {
  id: number
  title: string
  status: string
}

const props = defineProps({
  selectedHosts: {
    type: Array as PropType<Host[]>,
    required: true,
    default: () => []
  },
  visible: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:visible'])

const configList = ref<ConfigItem[]>([])
const tableLoading = ref(false)
const selectedConfigs = ref<ConfigItem>()
const selectedConfigId = ref<number>()

const pagination = reactive({
  currentPage: 1,
  pageSize: 5,
  total: 0
})

// 获取配置列表
const fetchConfigList = async () => {
  try {
    tableLoading.value = true
    const params = {
      page: pagination.currentPage,
      pageSize: pagination.pageSize
    }
    const res = await getagentconf(params)
    configList.value = res.data.configs
    pagination.total = res.data.nums || 0
  } catch (error) {
    console.error('获取配置列表失败:', error)
    ElMessage.error('获取配置列表失败')
  } finally {
    tableLoading.value = false
  }
}

// 处理配置选择
const handleSelectionChange = (row: ConfigItem) => {
  selectedConfigs.value = row
}

// 分页处理
const handleSizeChange = (val: number) => {
  pagination.pageSize = val
  fetchConfigList()
}

const handleCurrentChange = (val: number) => {
  pagination.currentPage = val
  fetchConfigList()
}

// 提交表单
const handleSubmit = async () => {
  if (selectedConfigs.value?.id == 0) {
    ElMessage.warning('请选择要下发的配置')
    return
  }

  try {
    const submitData = {
      config_id: selectedConfigs.value?.id,
      uuids: props.selectedHosts.map((host) => host.uuid)
    }

    await pushagentconfbyhost(submitData) // 假设这是你的提交API
    ElMessage.success('下发配置成功')
    handleCancel()
  } catch (error) {
    console.error('下发配置失败:', error)
    ElMessage.error('下发配置失败')
  }
}

// 取消
const handleCancel = () => {
  emit('update:visible', false)
}

onMounted(() => {
  fetchConfigList()
})

defineExpose({
  handleSubmit,
  handleCancel
})
</script>

<style scoped>
.config-form {
  padding: 10px;
}

.selected-hosts-section {
  margin-bottom: 20px;
}

.selected-hosts {
  margin-top: 10px;
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.host-tag {
  margin-right: 8px;
  margin-bottom: 8px;
}

.config-section {
  margin-bottom: 20px;
}

.pagination-area {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.button-area {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

h4 {
  margin: 0;
  padding: 10px 0;
  color: #606266;
}
</style>
