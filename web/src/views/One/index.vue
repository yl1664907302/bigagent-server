<template>
  <div class="dashboard-container">
    <el-row :gutter="24">
      <el-col :span="12">
        <el-card class="event-card">
          <div class="card-header">
            <h2><i class="el-icon-warning-outline"></i> agent事件</h2>
          </div>
          <el-list class="event-list">
            <el-list-item class="event-item" v-if="agentEvents.dnum > 0" @click="handleClick">
              <div class="status-indicator"></div>
              仍有<span class="number">{{ agentEvents.dnum }}</span
              >个agent实例离线未处理
            </el-list-item>
            <el-list-item class="event-item" v-if="agentConfigFail.fnum > 0" @click="handleClick">
              <div class="status-indicator"></div>
              配置ID <span class="number2">{{ agentConfigFail.id }}</span
              >, 仍有<span class="number">{{ agentConfigFail.fnum }}</span
              >个agent配置更新失败
            </el-list-item>
          </el-list>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { getagentnumdead, getagentconfigfail } from '@/api/login'
import { onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const agentEvents = reactive({
  dnum: 0
})

const agentConfigFail = reactive({
  id: 0,
  fnum: 0
})

const handleClick = () => {
  router.push('/config/info')
}

const handleClick2 = () => {
  router.push({
    path: '/config/info',
    query: {
      page: 1,
      pageSize: 10,
      c_desc_f: agentConfigFail.id
    }
  })
}

const handleGetAgentNumDead = async () => {
  try {
    const res = await getagentnumdead()
    agentEvents.dnum = res.data
  } catch (error) {
    console.error('获取agent数量失败:', error)
  }
}

const handleGetAgentConfigFail = async () => {
  try {
    const res = await getagentconfigfail()
    agentConfigFail.fnum = res.data.fnum
    agentConfigFail.id = res.data.id
  } catch (error) {
    console.error('获取agent数量失败:', error)
  }
}

onMounted(() => {
  handleGetAgentNumDead()
  handleGetAgentConfigFail()
})
</script>

<style scoped>
.dashboard-container {
  padding: 20px;
}

.event-card {
  border-radius: 20px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  background: linear-gradient(to bottom right, #ffffff, #f8f9fa);
}

.card-header {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

h2 {
  color: #2c3e50;
  font-size: 1.5rem;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.event-list {
  padding: 10px 0;
}

.event-item {
  display: flex;
  align-items: center;
  padding: 15px;
  border-radius: 12px;
  background-color: rgba(0, 0, 0, 0.02);
  margin-bottom: 10px;
  transition: all 0.2s ease;
  cursor: pointer;
}

.event-item:hover {
  background-color: rgba(0, 0, 0, 0.04);
  transform: translateX(5px);
}

.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #ff4d4f;
  margin-right: 12px;
  animation: pulse 2s infinite;
}

.number {
  color: #ff4d4f;
  font-weight: bold;
  margin: 0 4px;
}

.number2 {
  color: #4a65e0;
  font-weight: bold;
  margin: 0 4px;
}

@keyframes pulse {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.3);
    opacity: 0.7;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}
</style>
