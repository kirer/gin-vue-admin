<template>
  <div>
    <div class="sticky top-0.5 z-10 bg-white">
      <el-input v-model="filterText" class="w-3/5" placeholder="筛选" />
      <el-button class="float-right" type="primary" @click="on_update">确 定</el-button>
    </div>
    <div class="tree-content">
      <el-scrollbar>
        <el-tree ref="apiTree" :data="tree_data" :default-checked-keys="tree_ids" :props="default_props" highlight-current
          node-key="onlyId" show-checkbox :filter-node-method="on_filter" @check="on_change" />
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
import { api_all } from '@/api/api'
import { casbin_update, casbin_get } from '@/api/casbin'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({ row: { default: () => { }, type: Object } })
const default_props = ref({ children: 'children', label: 'desc' })
const filterText = ref('')
const tree_data = ref([])
const tree_ids = ref([])
const activeUserId = ref('')
const init = async () => {
  var res = await api_all()
  const apis = res.data.apis
  tree_data.value = on_build_tree(apis)
  res = await casbin_get({ authId: props.row.authId })
  activeUserId.value = props.row.authId
  tree_ids.value = []
  res.data.paths && res.data.paths.forEach(item => { tree_ids.value.push('p:' + item.path + 'm:' + item.method) })
}
init()
// 创建api树方法
const on_build_tree = (apis) => {
  const api = {}
  apis && apis.forEach(item => {
    item.onlyId = 'p:' + item.path + 'm:' + item.method
    if (Object.prototype.hasOwnProperty.call(api, item.group)) {
      api[item.group].push(item)
    } else {
      Object.assign(api, { [item.group]: [item] })
    }
  })
  const tree = []
  for (const key in api) { tree.push({ ID: key, desc: key + '组', children: api[key] }) }
  return tree
}

const needConfirm = ref(false)
const on_change = () => { needConfirm.value = true }
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  on_update()
}

// 关联关系确定
const apiTree = ref(null)
const on_update = async () => {
  const checkArr = apiTree.value.getCheckedNodes(true)
  var casbinInfos = []
  checkArr && checkArr.forEach(item => {
    var casbinInfo = {
      path: item.path,
      method: item.method
    }
    casbinInfos.push(casbinInfo)
  })
  const res = await casbin_update({
    authId: activeUserId.value,
    casbinInfos
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'api设置成功' })
  }
}

defineExpose({ needConfirm, enterAndNext })

const on_filter = (value, data) => {
  if (!value) return true
  return data.desc.indexOf(value) !== -1
}
watch(filterText, (val) => { apiTree.value.filter(val) })

</script>

