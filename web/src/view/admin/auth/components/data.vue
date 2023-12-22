<template>
  <div>
    <warning-bar title="此功能仅用于创建角色和角色的many2many关系表，具体使用还须自己结合表实现业务，详情参考示例代码（客户示例）。此功能不建议使用，建议使用插件市场【组织管理功能（点击前往）】来管理资源权限。"
      href="https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=36" />
    <div class="sticky top-0.5 z-10 bg-white my-4">
      <el-button class="float-left" type="primary" @click="all">全选</el-button>
      <el-button class="float-left" type="primary" @click="self">本角色</el-button>
      <el-button class="float-left" type="primary" @click="selfAndChildren">本角色及子角色</el-button>
      <el-button class="float-right" type="primary" @click="authDataEnter">确 定</el-button>
    </div>
    <div class="clear-both pt-4">
      <el-checkbox-group v-model="dataAuthId" @change="selectAuthority">
        <el-checkbox v-for="(item, key) in auths" :key="key" :label="item">{{ item.authName }}</el-checkbox>
      </el-checkbox-group>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { auth_set_data } from '@/api/auth'
const props = defineProps({
  row: { default: function () { return {} }, type: Object },
  auths: { default: function () { return [] }, type: Array }
})
const auths = ref([])
const needConfirm = ref(false)
const dataAuthId = ref([])
// 平铺角色
const on_round_auth = (authsData) => {
  authsData && authsData.forEach(item => {
    auths.value.push({ authId: item.authId, authName: item.authName })
    if (item.children && item.children.length) { on_round_auth(item.children) }
  })
}
const init = () => {
  on_round_auth(props.auths)
  props.row.dataAuthId && props.row.dataAuthId.forEach(item => {
    const obj = auths.value && auths.value.filter(au => au.authId === item.authId) && auths.value.filter(au => au.authId === item.authId)[0]
    dataAuthId.value.push(obj)
  })
}
init()

// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authDataEnter()
}

const emit = defineEmits(['changeRow'])
const all = () => {
  dataAuthId.value = [...auths.value]
  emit('changeRow', 'dataAuthId', dataAuthId.value)
  needConfirm.value = true
}
const self = () => {
  dataAuthId.value = auths.value.filter(item => item.authId === props.row.authId)
  emit('changeRow', 'dataAuthId', dataAuthId.value)
  needConfirm.value = true
}
const selfAndChildren = () => {
  const arrBox = []
  getChildrenId(props.row, arrBox)
  dataAuthId.value = auths.value.filter(item => arrBox.indexOf(item.authId) > -1)
  emit('changeRow', 'dataAuthId', dataAuthId.value)
  needConfirm.value = true
}
const getChildrenId = (row, arrBox) => {
  arrBox.push(row.authId)
  row.children && row.children.forEach(item => {
    getChildrenId(item, arrBox)
  })
}
// 提交
const authDataEnter = async () => {
  const res = await auth_set_data(props.row)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '资源设置成功' })
  }
}

//   选择
const selectAuthority = () => {
  emit('changeRow', 'dataAuthId', dataAuthId.value)
  needConfirm.value = true
}

defineExpose({
  enterAndNext,
  needConfirm
})

</script>
