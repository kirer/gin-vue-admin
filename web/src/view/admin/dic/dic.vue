<template>
  <div>
    <warning-bar title="获取字典且缓存方法已在前端utils/dictionary 已经封装完成 不必自己书写 使用方法查看文件内注释" />
    <div class="dict-box flex gap-4">
      <div class="w-64 bg-white p-4">
        <div class="flex justify-between items-center">
          <span class="text font-bold">字典列表</span>
          <el-button type="primary" @click="on_create">新增</el-button>
        </div>
        <el-scrollbar class="mt-4" style="height: calc(100vh - 300px)">
          <div v-for="item in table_data" :key="item.ID"
            class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 hover:text-gray-800 group bg-gray-50"
            :class="select_id === item.ID && 'active'" @click="on_detail(item)">
            <span class="max-w-[160px] truncate">{{ item.name }}</span>
            <div>
              <el-icon class="group-hover:text-blue-500"
                :class="select_id === item.ID ? 'text-white-800' : 'text-blue-500'"
                @click.stop="on_update(item)">
                <Edit />
              </el-icon>
              <el-popover placement="top" width="160">
                <p>确定要删除吗？</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button type="primary" link @click="item.visible = false">取消</el-button>
                  <el-button type="primary" @click="on_delete(item)">确定</el-button>
                </div>
                <template #reference>
                  <el-icon class="ml-2 group-hover:text-red-500"
                    :class="select_id === item.ID ? 'text-white-800' : 'text-red-500'">
                    <Delete />
                  </el-icon>
                </template>
              </el-popover>
            </div>
          </div>
        </el-scrollbar>
      </div>
      <div class="flex-1 bg-white">
        <dic_detail :id="select_id" />
      </div>
    </div>
    <el-dialog v-model="dialog_visible" :before-close="on_dialog_cancel" :title="dialog_title">
      <el-form ref="dialog_form" :model="form_data" :rules="form_rules" label-width="110px">
        <el-form-item label="字典名（中）" prop="name">
          <el-input v-model="form_data.name" placeholder="请输入字典名（中）" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="字典名（英）" prop="type">
          <el-input v-model="form_data.type" placeholder="请输入字典名（英）" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="状态" prop="status" required>
          <el-switch v-model="form_data.status" active-text="开启" inactive-text="停用" />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input v-model="form_data.desc" placeholder="请输入描述" clearable :style="{ width: '100%' }" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="on_dialog_cancel">取 消</el-button>
          <el-button type="primary" @click="on_dialog_submit">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { Edit, Plus } from '@element-plus/icons-vue'
import dic_detail from './dic_detail.vue'
import { dic_create, dic_delete, dic_update, dic_get, dic_list, } from '@/api/dic'

const table_data = ref([])
const on_table_data = async () => {
  const res = await dic_list()
  if (res.code === 0) {
    table_data.value = res.data
  }
}
on_table_data()
//
const dialog_visible = ref(false)
const dialog_title = ref('新增')
const dialog_form = ref(null)
const form_data = ref({ name: null, type: null, status: true, desc: null, })
const form_rules = ref({
  name: [ { required: true, message: '请输入字典名（中）', trigger: 'blur', }, ],
  type: [ { required: true, message: '请输入字典名（英）', trigger: 'blur', }, ],
  desc: [ { required: true, message: '请输入描述', trigger: 'blur',  }, ],
})
const on_create = () => {
  dialog_title.value = '新增'
  dialog_visible.value = true
}
const on_update = async (row) => {
  const res = await dic_get({ ID: row.ID, status: row.status })
  console.log('on_update', res);
  if (res.code === 0) {
    form_data.value = res.data
    dialog_title.value = '变更'
    dialog_visible.value = true
  }
}
const on_dialog_submit = async () => {
  dialog_form.value.validate(async (valid) => {
    if (!valid) return
    let res
    if ('新增' == dialog_title.value) {
      res = await dic_create(form_data.value)
    }
    if ('变更' == dialog_title.value) {
      res = await dic_update(form_data.value)
    }
    if (res.code === 0) {
      ElMessage.success('操作成功')
      on_table_data()
    }
    on_dialog_cancel()
  })
}
const on_dialog_cancel = () => {
  dialog_visible.value = false
  form_data.value = { name: null, type: null, status: true, desc: null, }
}
const on_delete = async (row) => {
  row.visible = false
  const res = await dic_delete({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    on_table_data()
  }
}
//
const select_id = ref(1)
const on_detail = (row) => select_id.value = row.ID
</script>
<style>
.dict-box {
  height: calc(100vh - 240px);
}
.active {
  background-color: var(--el-color-primary) !important;
  color: #fff;
}</style>
