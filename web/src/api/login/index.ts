import request from '@/axios'
import type {
  DelRecevier,
  DelRecevier_response,
  SelectMarkDownByStatus,
  SelectMarkDownByStatus_response,
  SelectReceivers_response,
  StepsFormType,
  StepsFormType_response,
  UserType,
  SelectMarkDownByStatus2Time_response,
  SelectMarkDownByStatus2Time,
  CreateRobot,
  CreateRobot_response,
  SelectRobot,
  SelectRobot_response,
  DelRobot,
  DelRobot_response,
  Updatemarkdowntemplate,
  Updatemarkdowntemplate_response,
  Selectmarkdowntemplate,
  Selectmarkdowntemplate_response,
  UpdateRobot,
  UpdateRobot_response
} from './types'

interface RoleParams {
  roleName: string
}

export const loginApi = (data: UserType): Promise<IResponse<UserType>> => {
  return request.post({
    url: '/v1/login',
    data,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN
    }
  })
}

export const loginOutApi = (): Promise<IResponse> => {
  return request.get({
    url: '/v1/loginOut',
    headers: {
      Authorization: import.meta.env.VITE_TOKEN
    }
  })
}

export const getAdminRoleApi = (
  params: RoleParams
): Promise<IResponse<AppCustomRouteRecordRaw[]>> => {
  return request.get({ url: '/mock/role/list', params })
}

export const getTestRoleApi = (params: RoleParams): Promise<IResponse<string[]>> => {
  return request.get({ url: '/mock/role/list2', params })
}

// 新增配置文件
export const addagentconf = (data: any): Promise<IResponse<any>> => {
  return request.post({
    url: '/v1/add',
    data,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN
    }
  })
}

// 查询配置文件
export const getagentconf = (params: any): Promise<IResponse<any>> => {
  return request.get({
    url: '/v1/get',
    params,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 下发配置文件
export const pushagentconf = (data: any): Promise<IResponse<any>> => {
  return request.post({
    url: '/v1/push',
    data,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN
    }
  })
}

// 删除配置文件
export const delagentconf = (params: any): Promise<IResponse<any>> => {
  return request.delete({
    url: '/v1/del',
    params,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 编辑配置文件
export const editagentconf = (data: any): Promise<IResponse<any>> => {
  return request.put({
    url: '/v1/edit',
    data,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 下发指定agent配置
export const pushagentconfbyhost = (data: any): Promise<IResponse<any>> => {
  return request.post({
    url: '/v1/push_host',
    data,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN
    }
  })
}

// 查询agent信息
export const getagentinfo = (params: any): Promise<IResponse<any>> => {
  return request.get({
    url: '/v1/info',
    params,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 查询agent元数据
export const getagentmetadata = (params: any): Promise<IResponse<any>> => {
  return request.get({
    url: '/v1/agent_id',
    params,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 删除agent
export const delagent = (params: any): Promise<IResponse<any>> => {
  return request.delete({
    url: '/v1/del_agent',
    params,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 巡检agent
export const getpatrolagent = (params: any): Promise<IResponse<any>> => {
  return request.get({
    url: '/v1/agent_id_patrol',
    params,
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 查询离线agent数量
export const getagentnumdead = (): Promise<IResponse<any>> => {
  return request.get({
    url: '/v1/get_dead',
    headers: {
      Authorization: import.meta.env.VITE_TOKEN,
      'Content-Type': 'application/json'
    }
  })
}

// 查询更新失败的agent数量
export const getagentconfigfail = (): Promise<IResponse<any>> => {
  return request.get({
    url: '/v1/get_cofail',
    headers: {
      Authorization: '123456',
      'Content-Type': 'application/json'
    }
  })
}