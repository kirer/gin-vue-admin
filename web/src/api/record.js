import service from '@/utils/request'
export const record_delete = (data) => service({ url: '/record/delete', method: 'delete', data })
export const record_deletes = (data) => service({ url: '/record/deletes', method: 'delete', data })
export const record_list = (params) => service({ url: '/record/list', method: 'get', params })
