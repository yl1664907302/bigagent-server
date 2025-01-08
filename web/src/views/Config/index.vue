<template>
  <div class="agent-config-container">
    <el-card class="table-card">
      <!-- 操作按钮区域 -->
      <div class="operation-area">
        <el-button type="primary" @click="handleAdd" :icon="Plus"> 新增配置 </el-button>
      </div>

      <!-- 表格区域 -->
      <el-table
        v-loading="tableLoading"
        :data="tableData"
        border
        style="width: 100%"
        :header-cell-style="headerCellStyle"
        :cell-style="cellStyle"
        highlight-current-row
        stripe
      >
        <el-table-column prop="id" label="ID" min-width="60" />
        <el-table-column prop="title" label="配置标题" min-width="120" />
        <el-table-column prop="status" label="状态" min-width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === '有效' ? 'success' : 'danger'">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="times" label="次数" min-width="80" />
        <el-table-column prop="role_name" label="角色" min-width="100" />
        <el-table-column prop="ranges" label="范围" min-width="100" />
        <el-table-column label="接口信息" min-width="150">
          <template #default="{ row }">
            {{ row.host }}
          </template>
        </el-table-column>
        <el-table-column prop="auth_name" label="认证方式" min-width="100" />
        <el-table-column prop="slot_name" label="槽位" min-width="100" />
        <el-table-column prop="data_name" label="数据格式" min-width="100" />
        <el-table-column prop="collection_frequency" label="频次" min-width="100" />
        <el-table-column prop="created_at" label="创建时间" min-width="160">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="180" fixed="right">
          <template #default="{ row }">
            <el-space>
              <el-button
                size="small"
                type="primary"
                :loading="loadingStates[row.id]"
                @click="handleDo(row)"
              >
                下发(ALL)
              </el-button>
              <el-button size="small" type="warning" @click="handleEdit(row)"> 编辑 </el-button>
              <el-button size="small" type="danger" @click="handleDel(row)"> 删除 </el-button>
            </el-space>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页区域 -->
      <div class="pagination-area">
        <el-pagination
          v-model:current-page="pagination.currentPage"
          v-model:page-size="pagination.pageSize"
          :page-sizes="[5, 10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          background
        />
      </div>
    </el-card>
    <div>
      <!-- 弹窗表单 -->
      <el-dialog
        v-model="dialogVisible_add"
        :title="dialogTitle_add"
        width="50%"
        :close-on-click-modal="false"
        :destroy-on-close="true"
      >
        <ConfigForm @submit="handleSubmit" @reset="handleReset" />
      </el-dialog>
    </div>
    <div>
      <!-- 编辑表单 -->
      <el-dialog
        v-model="dialogVisible_edit"
        :title="dialogTitle_edit"
        width="50%"
        :close-on-click-modal="false"
        :destroy-on-close="true"
      >
        <ConfigForm_edit v-model="currentEditRow" @submit="handleSubmit" @reset="handleReset" />
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, h } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { delagentconf, getagentconf, pushagentconf } from '@/api/login'
import dayjs from 'dayjs'
import ConfigForm from './form.vue'
import ConfigForm_edit from './form_edit.vue'
import { Plus } from '@element-plus/icons-vue'
var dialogVisible_add = ref(false)
var dialogTitle_add = ref('新增agent配置')
var dialogVisible_edit = ref(false)
var dialogTitle_edit = ref('编辑agent配置')
interface TableItem {
  id: number
  title: string
  status: string
  times: number
  role_name: string
  details: string
  ranges: string
  auth_name: string
  data_name: string
  slot_name: string
  protocol: string
  host: string
  port: number
  path: string
  token: string
  collection_frequency: string
  created_at: string
  updated_at: string
}

// 添加 loading 状态变量
const loadingStates = ref<{ [key: number]: boolean }>({})

// 编辑框预填值
const currentEditRow = ref<TableItem | null>(null)

// 表格数据
const tableData = ref<TableItem[]>([])
const tableLoading = ref(false)

// 分页配置
const pagination = reactive({
  currentPage: 1,
  pageSize: 5,
  total: 0
})

// 格式化日期
const formatDate = (date: string) => {
  return dayjs(date).format('YYYY-MM-DD HH:mm:ss')
}

// 获取表格数据
const fetchTableData = async () => {
  try {
    tableLoading.value = true
    const params = {
      page: pagination.currentPage,
      pageSize: pagination.pageSize
    }
    const res = await getagentconf(params)
    tableData.value = res.data.configs
    // 如果后端返回了总数，则更新
    if (res.data.nums) {
      pagination.total = res.data.nums
    }
  } catch (error) {
    console.error('获取数据失败:', error)
    ElMessage.error('获取数据失败，请重试')
  } finally {
    tableLoading.value = false
  }
}

// 下发配置
const PushConfig = async (id) => {
  try {
    const params = {
      config_id: id
    }
    const res = await pushagentconf(params)
    ElMessage.success(res.data)
  } catch (error) {
    console.error('配置下发失败:', error)
    ElMessage.error('配置下发失败，请重试')
  }
}

// 删除配置
const DelConfig = async (id) => {
  try {
    const params = {
      config_id: id
    }
    const res = await delagentconf(params)
    ElMessage.success(res.data)
  } catch (error) {
    console.error('配置删除失败:', error)
    ElMessage.error('配置删除失败，请重试')
  }
}

// 分页处理
const handleSizeChange = (val: number) => {
  pagination.pageSize = val
  fetchTableData()
}

const handleCurrentChange = (val: number) => {
  pagination.currentPage = val
  fetchTableData()
}

// 新增表格操作
const handleAdd = () => {
  dialogVisible_add.value = true
}

// 编辑表格操作
const handleEdit = (row) => {
  currentEditRow.value = { ...row }
  dialogVisible_edit.value = true
}

// 添加表单处理函数
const handleSubmit = () => {
  dialogVisible_edit.value = false
  dialogVisible_add.value = false
  ElMessage.success('提交成功')
  fetchTableData()
}

const handleReset = () => {
  ElMessage.success('已重置')
}

const handleDo = async (row: TableItem) => {
  try {
    await ElMessageBox.confirm(
      h('div', null, [
        '确认下发配置 "',
        row.title,
        '" 吗？',
        h('span', { style: { color: 'red' } }, '范围是所有主机！')
      ]),
      {
        title: '下发确认',
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    await PushConfig(row.id)
  } catch (error) {
    // 用户取消或发生错误
    if (error !== 'cancel') {
      ElMessage({
        type: 'error',
        message: '下发失败：' + error
      })
    }
  } finally {
    await fetchTableData()
    // 清除 loading 状态
    loadingStates.value[row.id] = false
  }
}

const handleDel = async (row: TableItem) => {
  try {
    await ElMessageBox.confirm(
      h('div', null, [
        '确认删除配置 "',
        row.title,
        '" 吗？',
        h('span', { style: { color: 'red' } }, '删除后无法找回！')
      ]),
      {
        title: '删除确认',
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    // 用户点击确认后执行删除
    await DelConfig(row.id)
  } catch (error) {
    // 用户取消或发生错误
    if (error !== 'cancel') {
      ElMessage({
        type: 'error',
        message: '删除失败：' + error
      })
    }
  } finally {
    // 清除 loading 状态
    loadingStates.value[row.id] = false
    fetchTableData()
  }
}

// 添加表格样式配置
const headerCellStyle = {
  backgroundColor: '#f5f7fa',
  color: '#606266',
  fontWeight: 'bold',
  fontSize: '14px',
  height: '45px',
  padding: '8px'
}

const cellStyle = {
  fontSize: '14px',
  padding: '8px'
}

onMounted(() => {
  fetchTableData()
})
</script>

<style scoped>
.agent-config-container {
  padding: 12px;
}

.table-card {
  border-radius: 4px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
}

.operation-area {
  margin-bottom: 15px;
  display: flex;
  gap: 10px;
  padding: 0 5px;
}

.pagination-area {
  margin-top: 15px;
  display: flex;
  justify-content: flex-end;
  padding: 5px;
}

/* 美化滚动条 */
:deep(.el-table__body-wrapper::-webkit-scrollbar) {
  width: 8px;
  height: 8px;
}

:deep(.el-table__body-wrapper::-webkit-scrollbar-thumb) {
  background-color: #dcdfe6;
  border-radius: 4px;
}

:deep(.el-table__body-wrapper::-webkit-scrollbar-track) {
  background-color: #f5f7fa;
}

/* 表格hover效果增强 */
:deep(.el-table__row:hover) {
  background-color: #f5f7fa !important;
  transition: background-color 0.3s;
}

/* 表格边框美化 */
:deep(.el-table) {
  border-radius: 4px;
  overflow: hidden;
}

/* 按钮悬停效果 */
.el-button {
  transition: all 0.3s;
}

.el-button:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

:deep(.el-table .el-table__cell) {
  padding: 6px 0;
}

:deep(.el-table th.el-table__cell) {
  padding: 8px 0;
}
</style>
