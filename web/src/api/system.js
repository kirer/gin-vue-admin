import service from '@/utils/request'
export const system_get_config = () => service({ url: '/system/get_config', method: 'post' })
export const system_set_config = (data) => service({ url: '/system/set_config', method: 'post', data })
export const system_get_info = () => service({ url: '/system/get_info', method: 'post', donNotShowLoading: true })
