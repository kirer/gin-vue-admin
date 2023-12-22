import service from '@/utils/request'
export const dic_create = (data) => service({ url: '/dic/create', method: 'post', data })
export const dic_delete = (data) => service({ url: '/dic/delete', method: 'delete', data })
export const dic_update = (data) => service({ url: '/dic/update', method: 'put', data })
export const dic_get = (params) => service({ url: '/dic/get', method: 'get', params })
export const dic_list = (params) => service({ url: '/dic/list', method: 'get', params })