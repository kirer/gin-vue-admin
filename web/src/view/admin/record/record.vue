<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="search_info">
        <el-form-item label="请求方法">
          <el-input v-model="search_info.method" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="请求路径">
          <el-input v-model="search_info.path" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="结果状态码">
          <el-input v-model="search_info.status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="on_submit">查询</el-button>
          <el-button icon="refresh" @click="on_reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">

        <el-popover v-model="delete_visible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="delete_visible = false">取消</el-button>
            <el-button type="primary" @click="on_deletes">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!selections.length"
              @click="delete_visible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" :data="table_data" style="width: 100%" tooltip-effect="dark" row-key="ID" @selection-change="on_selection_changes">
        <el-table-column align="left" type="selection" width="55" />
        <el-table-column align="left" label="操作人" width="140">
          <template #default="scope">
            <div>{{ scope.row.user.userName }}({{ scope.row.user.nickName }})</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="状态码" prop="status" width="120">
          <template #default="scope">
            <div>
              <el-tag type="success">{{ scope.row.status }}</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="请求IP" prop="ip" width="120" />
        <el-table-column align="left" label="请求方法" prop="method" width="120" />
        <el-table-column align="left" label="请求路径" prop="path" width="240" />
        <el-table-column align="left" label="请求" prop="path" width="80">
          <template #default="scope">
            <div>
              <el-popover v-if="scope.row.body" placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ on_format_body(scope.row.body) }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;">
                    <warning />
                  </el-icon>
                </template>
              </el-popover>

              <span v-else>无</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="响应" prop="path" width="80">
          <template #default="scope">
            <div>
              <el-popover v-if="scope.row.resp" placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ on_format_body(scope.row.resp) }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;">
                    <warning />
                  </el-icon>
                </template>
              </el-popover>
              <span v-else>无</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="on_delete(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button icon="delete" type="primary" link @click="scope.row.visible = true">删除</el-button>
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
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { record_list, record_delete, record_deletes } from '@/api/record'
import { formatDate } from '@/utils/format'
//
const page = ref(1)
const total = ref(0)
const page_size = ref(10)
const table_data = ref([])
const search_info = ref({})
const on_reset = () => search_info.value = {}
const on_submit = () => {
  page.value = 1
  page_size.value = 10
  if (search_info.value.status === '') { search_info.value.status = null }
  on_table_data()
}
const on_size_changed = (val) => { page_size.value = val; on_table_data() }
const on_current_changed = (val) => { page.value = val; on_table_data() }
const on_table_data = async () => {
  const table = await record_list({ page: page.value, pageSize: page_size.value, ...search_info.value, })
  if (table.code === 0) {
    table_data.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    page_size.value = table.data.pageSize
  }
}
on_table_data()
const on_format_body = (value) => {
  try { return JSON.parse(value) } catch (err) { return value }
}
//
const delete_visible = ref(false)
const selections = ref([])
const on_selection_changes = (val) => selections.value = val
const on_delete = async (row) => {
  row.visible = false
  const res = await record_delete({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    if (table_data.value.length === 1 && page.value > 1) { page.value-- }
    on_table_data()
  }
}
const on_deletes = async () => {
  const ids = []
  selections.value && selections.value.forEach(item => ids.push(item.ID))
  const res = await record_deletes({ ids })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    if (table_data.value.length === ids.length && page.value > 1) { page.value--}
    delete_visible.value = false
    on_table_data()
  }
}
</script>
<style lang="scss">
.table-expand {
  padding-left: 60px;
  font-size: 0;

  label {
    width: 90px;
    color: #99a9bf;

    .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 50%;
    }
  }
}

.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
}

.popover-box::-webkit-scrollbar {
  display: none;
  /* Chrome Safari */
}
</style>
