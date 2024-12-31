<template>
  <div class="agent-info-container">
    <el-card>
      <el-row :gutter="16">
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="agent_num.total_num">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  agent注册数
                  <el-tooltip effect="dark" content="已经注册的agent总数" placement="top">
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>合法 agents</span>
                <span class="green" v-if="agent_num.total_num > 0">
                  100%
                  <el-icon>
                    <CaretTop />
                  </el-icon>
                </span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="agent_num.alive_num">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  agent在线数
                  <el-tooltip effect="dark" content="目前在线的agent数量" placement="top">
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>agent在线率</span>
                <span class="red" v-if="agent_num.alive_num < agent_num.total_num">
                  {{ Math.floor((agent_num.alive_num / agent_num.total_num) * 100) }}%
                  <el-icon>
                    <CaretBottom />
                  </el-icon>
                </span>
              </div>
              <div class="footer-item">
                <span class="green" v-if="agent_num.alive_num == agent_num.total_num">
                  100%
                  <el-icon>
                    <CaretTop />
                  </el-icon>
                </span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="8">
          <div class="statistic-card">
            <el-statistic :value="agent_num.dead_num" title="New transactions today">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  agent离线数
                  <el-tooltip effect="dark" content="目前离线的agent数量" placement="top">
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>agent离线率</span>
                <span class="red" v-if="agent_num.alive_num < agent_num.total_num">
                  {{ Math.floor((agent_num.dead_num / agent_num.total_num) * 100) }}%
                  <el-icon>
                    <CaretTop />
                  </el-icon>
                </span>
                <span class="green" v-if="agent_num.alive_num == agent_num.total_num">
                  0%
                  <el-icon>
                    <CaretBottom />
                  </el-icon>
                </span>
              </div>
              <div class="footer-item">
                <!-- <el-icon :size="14">
                  <ArrowRight />
                </el-icon> -->
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
    </el-card>
    <div class="mt-4"></div>
    <el-card>
      <!-- 操作按钮区域 -->
      <div class="operation-area">
        <el-button type="primary" :disabled="!selectedData.length" @click="handleBatchOperation">
          批量下发
        </el-button>
        <el-button type="primary" @click="handleQuery"> 查询 </el-button>
        <el-select
          v-model="agent_option.type"
          placeholder="类型"
          style="width: 240px"
          :disabled="isDisabled"
        >
          <el-option
            v-for="item in options1"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select
          v-model="agent_option.platform"
          placeholder="平台"
          style="width: 240px"
          :disabled="isDisabled"
        >
          <el-option
            v-for="item in options2"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select
          v-model="agent_option.active"
          placeholder="状态"
          style="width: 240px"
          :disabled="isDisabled"
        >
          <el-option
            v-for="item in options3"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-input
          v-model="agent_option.ip"
          style="width: 240px"
          placeholder="请输入ip地址"
          @input="() => checkInput('ip')"
        />
        <el-input
          v-model="agent_option.uuid"
          style="width: 240px"
          placeholder="请输入机器uuid"
          @input="() => checkInput('uuid')"
        />
        <el-input
          v-model="agent_option.c_desc"
          style="width: 240px"
          placeholder="请输入配置id"
          @input="() => checkInput('c_desc')"
        />
        <el-input
          v-model="agent_option.c_desc2"
          style="width: 240px"
          placeholder="请输入更新失败的配置id"
          @input="() => checkInput('c_desc2')"
        />
      </div>
      <!-- 表格区域 -->
      <el-table
        v-loading="tableLoading"
        :data="tableData"
        border
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" :selectable="isRowSelectable" />
        <el-table-column prop="uuid" min-width="320">
          <template #header>
            <div style="display: inline-flex; align-items: center">
              UUID
              <el-tooltip effect="dark" content="悬浮查看资源占比" placement="top">
                <el-icon style="margin-left: 4px" :size="12">
                  <Warning />
                </el-icon>
              </el-tooltip>
            </div>
          </template>
          <template #default="{ row }">
            <el-popover
              v-if="row.active === 1"
              trigger="hover"
              placement="top"
              title="资源 (仅参考)"
              width="300"
              show-after="100"
            >
              <template #reference>
                <span>{{ row.uuid }}</span>
              </template>
              <div>
                <div class="resource-item">
                  <p class="resource-title">CPU使用率: {{ row.cpu_use }}</p>
                  <el-progress
                    :percentage="parseUsage(row.cpu_use)"
                    :status="
                      parseUsage(row.cpu_use) > 80
                        ? 'exception'
                        : parseUsage(row.cpu_use) > 60
                          ? 'warning'
                          : 'success'
                    "
                    :format="(percentage) => `${percentage}%`"
                  />
                </div>
                <div class="resource-item">
                  <p class="resource-title">内存使用率: {{ row.memory_use }}</p>
                  <el-progress
                    :percentage="parseUsage(row.memory_use)"
                    :status="
                      parseUsage(row.memory_use) > 80
                        ? 'exception'
                        : parseUsage(row.memory_use) > 60
                          ? 'warning'
                          : 'success'
                    "
                    :format="(percentage) => `${percentage}%`"
                  />
                </div>
                <template v-if="row.disk_use">
                  <div
                    v-for="(usage, drive) in parseDiskUsage(row.disk_use)"
                    :key="drive"
                    class="resource-item"
                  >
                    <p class="resource-title">磁盘 {{ drive }} 使用率: {{ usage }}%</p>
                    <el-progress
                      :percentage="usage"
                      :status="usage > 80 ? 'exception' : usage > 60 ? 'warning' : 'success'"
                      :format="(percentage) => `${percentage}%`"
                    />
                  </div>
                </template>
              </div>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column prop="active" label="活跃状态" min-width="100">
          <template #default="{ row }">
            <el-tag :type="row.active === 1 ? 'success' : 'danger'">
              {{ row.active === 1 ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" min-width="150">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="hostname" label="主机名" min-width="150" />
        <el-table-column prop="ipv4_first" label="IPv4" min-width="120" />
        <el-table-column prop="platform" label="平台" min-width="100" />
        <el-table-column prop="os" label="操作系统" min-width="100" />
        <el-table-column prop="kernel" label="内核版本" min-width="100" />
        <el-table-column prop="machine_type" label="机器类型" min-width="100" />
        <el-table-column prop="action_detail" label="配置详情" min-width="150" />
        <el-table-column prop="arch" label="arch型号" min-width="100" />
        <el-table-column prop="grpc_port" label="GRPC端口" min-width="100" />
        <el-table-column label="操作" min-width="185" fixed="right">
          <template #default="{ row }">
            <el-space>
              <el-button size="small" type="primary" @click="handleDetail(row)"> 数据 </el-button>
              <el-button size="small" type="primary" @click="handleDetail2(row)"> 容量 </el-button>
              <el-button size="small" type="danger" @click="deleteagent(row)"> 删除 </el-button>
            </el-space>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页区域 -->
      <div class="pagination-area">
        <el-pagination
          v-model:current-page="pagination.currentPage"
          v-model:page-size="pagination.pageSize"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <!-- 配置下发弹窗 -->
    <el-dialog v-model="dialogVisible_host" title="下发配置" width="50%">
      <ConfigForm v-model:visible="dialogVisible_host" :selected-hosts="selectedData" />
    </el-dialog>

    <!-- 详情弹窗 -->
    <el-dialog v-model="detailDialogVisible" :title="display_title_detail" width="80%">
      <div class="metadata-container">
        <vue-json-pretty
          :deep="3"
          :showSelectController="true"
          :highlightMouseoverNode="true"
          path="res"
          :data="detailData"
          :height="700"
          :virtual="true"
          :show-icon="true"
          :editable="false"
          :show-line-number="true"
        />
      </div>
      <el-space>
        <el-button size="small" type="primary" @click="changeDetail(detailData.uuid, 'stand1')">
          标准类型1(默认)
        </el-button>
        <el-button size="small" type="primary" @click="changeDetail(detailData.uuid, 'stand2')">
          标准类型2
        </el-button>
      </el-space>
    </el-dialog>

    <!-- 巡检详情弹窗 -->
    <el-dialog v-model="patrolVisible" :title="display_title_detail" width="80%">
      <div class="metadata-container">
        <vue-json-pretty
          :deep="3"
          :showSelectController="true"
          :highlightMouseoverNode="true"
          path="res"
          :data="patrolData"
          :height="700"
          :virtual="true"
          :show-icon="true"
          :editable="false"
          :show-line-number="true"
        />
      </div>
      <el-space />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
// 本人前端废物，代码写的很烂，不要学我，所以注释很多！
import { ref, reactive, onMounted, h } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getagentmetadata, delagent, getpatrolagent } from '@/api/login'
import ConfigForm from '@/views/Info/form.vue'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
import { fetchEventSource } from '@microsoft/fetch-event-source'
// 弹出
const dialogVisible_host = ref(false)
// 详情弹窗相关
const detailDialogVisible = ref(false)
const patrolVisible = ref(false)
const detailData = ref<{ uuid?: string }>({})
const patrolData = ref<{ uuid?: string }>({})
const display_title_detail = ref('')
import { CaretBottom, CaretTop, Warning } from '@element-plus/icons-vue'
const agent_option = reactive({
  type: '',
  platform: '',
  ip: '',
  uuid: '',
  active: '',
  c_desc: '',
  c_desc2: '',
  c_desc_f: ''
})

const options1 = [
  {
    value: 'physical',
    label: '物理机'
  },
  {
    value: 'virtual',
    label: '虚拟机'
  },
  {
    value: '',
    label: '所有'
  }
]

const options2 = [
  {
    value: 'linux',
    label: 'LINUX'
  },
  {
    value: 'windows',
    label: 'WINDOWS'
  },
  {
    value: '',
    label: '所有'
  }
]

const options3 = [
  {
    value: '1',
    label: '在线'
  },
  {
    value: '0',
    label: '离线'
  },
  {
    value: '',
    label: '所有'
  }
]

// agent 信息
interface AgentInfo {
  uuid: string
  net_ip: string
  hostname: string
  ipv4_first: string
  active: number
  status: string
  action_detail: string
  grpc_port: string
  platform: string
  machine_type: string
  os: string
  kernel: string
  arch: string
  disk_use: string
  cpu_use: string
  memory_use: string
  created_at: string
  updated_at: string
}
// 表格数据
const tableData = ref<AgentInfo[]>([])
const selectedData = ref<AgentInfo[]>([]) // 存储选中的完整数据
const selectedUUIDs = ref<string[]>([]) // 存储选中行的 UUID
const tableLoading = ref(false)

// 分页配置
const pagination = reactive({
  currentPage: 1,
  pageSize: 10,
  total: 0
})

// 统计数据
const agent_num = reactive({
  total_num: 0,
  dead_num: 0,
  alive_num: 0
})

// 元数据配置
const metadata = reactive({
  uuid: '',
  model_name: 'stand1'
})

// 获取状态标签类型
const getStatusType = (status: string) => {
  if (status.includes('异常')) return 'danger'
  if (status.includes('正常')) return 'success'
  return 'warning'
}

// 处理选择变化
const handleSelectionChange = (rows: AgentInfo[]) => {
  // ���贝选的数据
  selectedData.value = JSON.parse(JSON.stringify(rows))
  // 保存选中行的 UUID
  selectedUUIDs.value = rows.map((row) => row.uuid)
}

// 批量操作
const handleBatchOperation = () => {
  if (selectedData.value.length === 0) {
    ElMessage.warning('请先选择需要操作的主机')
    return
  }
  dialogVisible_host.value = true
}

// 查询
const handleQuery = () => {
  sseController.value.abort() // 查询时中断连接
  sseController.value = initSSE()
}

// 分页处理
const handleSizeChange = (val: number) => {
  pagination.pageSize = val
  sseController.value.abort() // 切换分页时中断连接
  sseController.value = initSSE()
}

const handleCurrentChange = (val: number) => {
  pagination.currentPage = val
  sseController.value.abort() // 切换分页时中断连接
  sseController.value = initSSE()
}

// 只有当active为1时才能勾选
const isRowSelectable = (row: AgentInfo) => {
  return row.active === 1
}

// 修改 initSSE 函数
const initSSE = () => {
  const controller = new AbortController()
  fetchEventSource(
    'http://127.0.0.1:8080/v1/info_sse?page=' +
      pagination.currentPage +
      '&pageSize=' +
      pagination.pageSize +
      '&ip=' +
      agent_option.ip +
      '&uuid=' +
      agent_option.uuid +
      '&platform=' +
      agent_option.platform +
      '&type=' +
      agent_option.type +
      '&active=' +
      agent_option.active +
      '&c_desc=' +
      agent_option.c_desc +
      '&c_desc_f=' +
      agent_option.c_desc2,
    {
      signal: controller.signal,
      headers: {
        Accept: 'text/event-stream',
        Authorization: '123456'
      },
      async onopen(response: { ok: any; status: number }) {
        if (response.ok && response.status === 200) {
          console.log('SSE 连接成功')
        } else if (response.status >= 400 && response.status < 500 && response.status !== 429) {
          throw new Error(`Failed to connect: ${response.status}`)
        }
      },
      onmessage(event) {
        try {
          if (!event.data) {
            console.warn('Empty event data received')
            return
          }
          const data = JSON.parse(event.data)
          updateTableData(data)
        } catch (error) {
          console.error('解析消息失败:', error)
        }
      },
      onclose() {
        console.log('SSE 连接关闭')
        // 可以在这里添加重连逻辑
        setTimeout(initSSE, 1000)
      },
      onerror(err) {
        console.error('SSE 错误:', err)
        // ElMessage.error('SSE 连接错误，请检查网络或服务器')
        controller.abort() // 出错时中断连接
      }
    }
  )
  return controller
}

const sseController = ref(initSSE())

const updateTableData = (rep) => {
  if (selectedUUIDs.value.length > 0) {
    console.log('存在勾选项时跳过更新')
    return
  }
  // notfiy(rep)
  if (!rep || !rep.data) {
    console.error('Invalid response data')
    ElMessage.error('更新表格数据失败：无效数据')
    return
  }

  tableLoading.value = true
  try {
    console.log(rep.data)
    tableData.value = rep.data.agentInfos || []
    pagination.total = rep.data.nums ?? pagination.total
    agent_num.total_num = rep.data.nums
    agent_num.alive_num = rep.data.anums
    agent_num.dead_num = rep.data.dnums
  } catch (error) {
    console.error('更新表格数据失败:', error)
    ElMessage.error('更新表格数据失败')
  } finally {
    tableLoading.value = false
  }
}

// const notfiy = (data) => {
//   data.data.agentInfos.forEach((newAgent) => {
//     const oldAgent = tableData.value.find((agent) => agent.uuid === newAgent.uuid)
//     if (oldAgent && oldAgent.active !== newAgent.active) {
//       ElNotification({
//         title: newAgent?.active === 1 ? '设备上线通知' : '设备离线通知',
//         message: `设备 ${newAgent?.hostname}(${newAgent?.net_ip}) ${newAgent?.active === 1 ? '已上线' : '已离线'}`,
//         type: newAgent?.active === 1 ? 'success' : 'warning',
//         duration: 3000,
//         position: 'top-right'
//       })
//     }
//   })
// }

// 获取表格数据（静态）
// const fetchTableData = async () => {
//   try {
//     tableLoading.value = true
//     const params = {
//       page: pagination.currentPage,
//       pageSize: pagination.pageSize
//     }
//     const res = await getagentinfo(params)
//     tableData.value = res.data.agentInfos
//     // 如果后端返回了总数，则更新
//     if (res.data.nums) {
//       pagination.total = res.data.nums
//     }
//   } catch (error) {
//     console.error('获取数据失败:', error)
//     ElMessage.error('获取数据失败，请重试')
//   } finally {
//     tableLoading.value = false
//   }
// }

// 删除agent
const deleteagent = async (row: AgentInfo) => {
  try {
    await ElMessageBox.confirm(
      h('div', null, [
        '确认删除agent "',
        row.hostname,
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
    await Delagent(row.uuid)
  } catch (error) {
    // 用户取消或发生错误
    if (error !== 'cancel') {
      ElMessage({
        type: 'error',
        message: '删除失败：' + error
      })
    }
  } finally {
    sseController.value.abort() // 切换分页时中断连接
    sseController.value = initSSE()
  }
}

// 删除配置
const Delagent = async (uuid) => {
  try {
    const params = {
      uuid: uuid
    }
    const res = await delagent(params)
    ElMessage.success(res.data)
  } catch (error) {
    console.error('agent删除失败:', error)
    ElMessage.error('agent删除失败，请重试')
  }
}

// 获取元数据
const fetchAgentMetadata = async (uuid, model_name = metadata.model_name) => {
  try {
    const params = {
      uuid: uuid,
      model_name: model_name
    }
    const res = await getagentmetadata(params)
    detailData.value = res.data
  } catch (error) {
    console.error('获取元数据失败:', error)
    ElMessage.error('获取元数据失败，请重试')
  } finally {
  }
}

const fetchAgentMetadata2 = async (uuid, model_name = metadata.model_name) => {
  try {
    const params = {
      uuid: uuid,
      model_name: model_name
    }
    const res = await getpatrolagent(params)
    patrolData.value = res.data
  } catch (error) {
    console.error('获取元数据失败:', error)
    ElMessage.error('获取元数据失败，请重试')
  } finally {
  }
}

const changeDetail = (uuid, model_name) => {
  fetchAgentMetadata(uuid, model_name)
  display_title_detail.value = '当前数据类型：' + model_name
  detailDialogVisible.value = true
}

const handleDetail = (row: AgentInfo) => {
  if (row.active == 0) {
    ElMessage.warning('设备不在线，无法获取实时数据')
    return
  }
  detailData.value = {}
  fetchAgentMetadata(row.uuid)
  display_title_detail.value = '当前数据模型：' + 'stand1'
  detailDialogVisible.value = true
}

// 巡检处理详情按钮点击
const handleDetail2 = (row: AgentInfo) => {
  if (row.active == 0) {
    ElMessage.warning('设备不在线，无法获取实时数据')
    return
  }
  patrolData.value = {}
  fetchAgentMetadata2(row.uuid)
  display_title_detail.value = '当前数据模型：' + 'stand1'
  patrolVisible.value = true
}

const isDisabled = ref(false)

const checkInput = (field: 'ip' | 'uuid' | 'c_desc' | 'c_desc2') => {
  switch (field) {
    case 'ip':
      if (agent_option.ip) {
        agent_option.c_desc2 = ''
        agent_option.uuid = ''
        agent_option.c_desc = ''
        agent_option.platform = ''
        agent_option.type = ''
        agent_option.active = ''
      }
      break
    case 'uuid':
      if (agent_option.uuid) {
        agent_option.c_desc2 = ''
        agent_option.c_desc = ''
        agent_option.ip = ''
        agent_option.platform = ''
        agent_option.type = ''
        agent_option.active = ''
      }
      break
    case 'c_desc':
      if (agent_option.c_desc) {
        agent_option.c_desc2 = ''
        agent_option.uuid = ''
        agent_option.ip = ''
        agent_option.platform = ''
        agent_option.type = ''
        agent_option.active = ''
      }
      break
    case 'c_desc2':
      if (agent_option.c_desc2) {
        agent_option.c_desc = ''
        agent_option.uuid = ''
        agent_option.ip = ''
        agent_option.platform = ''
        agent_option.type = ''
        agent_option.active = ''
      }
  }
  // 更新禁状态
  isDisabled.value = !!(agent_option.ip || agent_option.uuid || agent_option.c_desc)
}

// 解析使用率百分比
const parseUsage = (usage: string) => {
  if (!usage) return 0
  return parseInt(usage)
}

const parseDiskUsage = (diskUse: string) => {
  if (!diskUse) return {}
  try {
    // 直接使用对象，不需要 JSON.parse
    return Object.fromEntries(
      Object.entries(diskUse).map(([drive, usage]) => [drive, parseInt(usage.replace('%', ''))])
    )
  } catch (error) {
    console.error('解析磁盘使用率失败:', error)
    return {}
  }
}

// const formatDiskSummary = (diskUse: string) => {
//   if (!diskUse) return 'N/A'
//   try {
//     // 直接使用对象
//     return Object.entries(diskUse)
//       .map(([drive, usage]) => `${drive}: ${usage}`)
//       .join(', ')
//   } catch (error) {
//     console.error('格式化磁盘使用率失败:', error)
//     return 'N/A'
//   }
// }

onMounted(() => {
  sseController
})
</script>

<style scoped>
.agent-info-container {
  padding: 20px;
}

.operation-area {
  margin-bottom: 20px;
  display: flex;
  gap: 10px;
}

.pagination-area {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

.mt-4 {
  margin-top: 1rem;
}

.metadata-container {
  background-color: #f5f7fa;
  border: 1px solid #e4e7ed;
  border-radius: 4px;
  padding: 16px;
  max-height: 600px;
  overflow: hidden;
}

:global(h2#card-usage ~ .example .example-showcase) {
  background-color: var(--el-fill-color) !important;
}

.el-statistic {
  --el-statistic-content-font-size: 28px;
}

.statistic-card {
  height: 100%;
  padding: 20px;
  border-radius: 4px;
  background-color: var(--el-bg-color-overlay);
}

.statistic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  font-size: 12px;
  color: var(--el-text-color-regular);
  margin-top: 16px;
}

.statistic-footer .footer-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.statistic-footer .footer-item span:last-child {
  display: inline-flex;
  align-items: center;
  margin-left: 4px;
}

.green {
  color: var(--el-color-success);
}
.red {
  color: var(--el-color-error);
}

.resource-title {
  margin: 8px 0 4px;
  font-weight: bold;
}

.resource-item {
  margin-bottom: 12px;
}

.resource-item:last-child {
  margin-bottom: 0;
}
</style>
