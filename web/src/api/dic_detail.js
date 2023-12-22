import service from '@/utils/request'
export const dic_detail_create = (data) => service({ url: '/dic_detail/create', method: 'post', data })
export const dic_detail_delete = (data) => service({ url: '/dic_detail/delete', method: 'delete', data })
export const dic_detail_update = (data) => service({ url: '/dic_detail/update', method: 'put', data })
export const dic_detail_get = (data) => service({ url: '/dic_detail/get', method: 'get', data })
export const dic_detail_list = (params) => service({ url: '/dic_detail/list', method: 'get', params })
