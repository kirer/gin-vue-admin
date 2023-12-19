import service from '@/utils/request'
export const casbin_update = (data) => service({ url: '/casbin/update', method: 'put', data })
export const casbin_get = (data) => service({ url: '/casbin/get', method: 'post', data })
