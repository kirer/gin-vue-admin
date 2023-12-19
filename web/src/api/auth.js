import service from '@/utils/request'
export const auth_create = (data) => service({ url: '/auth/create', method: 'post', data })
export const auth_delete = (data) => service({ url: '/auth/delete', method: 'delete', data })
export const auth_update = (data) => service({ url: '/auth/update', method: 'put', data })
export const auth_list = (data) => service({ url: '/auth/list', method: 'post', data })
export const auth_copy = (data) => service({ url: '/auth/copy', method: 'post', data })
export const auth_set_data = (data) => service({ url: '/auth/set_data', method: 'post', data })

