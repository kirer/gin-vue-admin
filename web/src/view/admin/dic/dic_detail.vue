<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list justify-between">
        <span class="text font-bold">字典详细内容</span>
        <el-button type="primary" icon="plus" @click="on_create">新增字典项</el-button>
      </div>
      <el-table ref="multipleTable" :data="table_data" style="width: 100%" tooltip-effect="dark" row-key="ID">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="展示值" prop="label" />
        <el-table-column align="left" label="字典值" prop="value" />
        <el-table-column align="left" label="扩展值" prop="extend" />
        <el-table-column align="left" label="启用状态" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>

        <el-table-column align="left" label="排序标记" prop="sort" width="120" />

        <el-table-column align="left" label="操作" width="180">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="on_update(scope.row)">变更</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="on_delete(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="page_size" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="on_current_changed"
          @size-change="on_size_changed" />
      </div>
    </div>

    <el-dialog v-model="dialog_visible" :before-close="on_dialog_cancel" :title="type === 'create' ? '添加字典项' : '修改字典项'">
      <el-form ref="dialog_form" :model="form_data" :rules="form_rules" label-width="110px">
        <el-form-item label="展示值" prop="label">
          <el-input v-model="form_data.label" placeholder="请输入展示值" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input-number v-model.number="form_data.value" step-strictly :step="1" placeholder="请输入字典值" clearable
            :style="{ width: '100%' }" min="-2147483648" max="2147483647" />
        </el-form-item>
        <el-form-item label="扩展值" prop="extend">
          <el-input v-model="form_data.extend" placeholder="请输入扩展值" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="启用状态" prop="status" required>
          <el-switch v-model="form_data.status" active-text="开启" inactive-text="停用" />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input-number v-model.number="form_data.sort" placeholder="排序标记" />
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
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { dic_detail_create, dic_detail_delete, dic_detail_update, dic_detail_get, dic_detail_list } from '@/api/dic_detail' // 此处请自行替换地址
import { formatBoolean, formatDate } from '@/utils/format'

const props = defineProps({ id: { type: Number, default: 0 } })
//
const page = ref(1)
const total = ref(0)
const page_size = ref(10)
const table_data = ref([])
const on_size_changed = (val) => { page_size.value = val; on_table_data() }
const on_current_changed = (val) => { page.value = val; on_table_data() }
const on_table_data = async () => {
  const table = await dic_detail_list({ page: page.value, pageSize: page_size.value, dic_id: props.id })
  if (table.code === 0) {
    table_data.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    page_size.value = table.data.pageSize
  }
}
on_table_data()
watch(() => props.id, () => {
  on_table_data()
})
//
const type = ref('')
const dialog_visible = ref(false)
const dialog_title = ref('新增')
const dialog_form = ref(null)
const form_data = ref({ label: null, value: null, status: true, sort: null })
const form_rules = ref({
  label: [{ required: true, message: '请输入展示值', trigger: 'blur' }],
  value: [{ required: true, message: '请输入字典值', trigger: 'blur' }],
  sort: [{ required: true, message: '排序标记', trigger: 'blur' }]
})
const on_create = () => {
  dialog_title.value = '新增'
  dialog_visible.value = true
}
const on_update = async (row) => {
  const res = await dic_detail_get({ ID: row.ID })
  if (res.code === 0) {
    form_data.value = res.data.reSysDictionaryDetail
    dialog_title.value = '变更'
    dialog_visible.value = true
  }
}
const on_dialog_submit = async () => {
  dialog_form.value.validate(async valid => {
    form_data.value.sysDictionaryID = props.sysDictionaryID
    if (!valid) return
    let res
    if ('新增' == dialog_title.value) {
      res = await dic_detail_create(form_data.value)
    }
    if ('变更' == dialog_title.value) {
      res = await dic_detail_update(form_data.value)
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
  form_data.value = { label: null, value: null, status: true, sort: null }
}
const on_delete = async (row) => {
  row.visible = false
  const res = await dic_detail_delete({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    if (table_data.value.length === 1 && page.value > 1) {
      page.value--
    }
    on_table_data()
  }
}
</script>

<style></style>
