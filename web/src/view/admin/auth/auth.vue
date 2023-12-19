<template>
  <div class="auth">
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addAuthority(0)">新增角色</el-button>
        <el-icon class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=8&vd_source=f2640257c21e3b547a790461ed94875e')">
          <VideoCameraFilled />
        </el-icon>
      </div>
      <el-table :data="table_data" :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        row-key="authId" style="width: 100%">
        <el-table-column label="角色ID" min-width="180" prop="authId" />
        <el-table-column align="left" label="角色名称" min-width="180" prop="authName" />
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button icon="setting" type="primary" link @click="opdendrawer(scope.row)">设置权限</el-button>
            <el-button icon="plus" type="primary" link @click="addAuthority(scope.row.authId)">新增子角色</el-button>
            <el-button icon="copy-document" type="primary" link @click="auth_copyFunc(scope.row)">拷贝</el-button>
            <el-button icon="edit" type="primary" link @click="editAuthority(scope.row)">编辑</el-button>
            <el-button icon="delete" type="primary" link @click="deleteAuth(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增角色弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="authForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader v-model="form.parentId" style="width:100%" :disabled="dialogType === 'add'"
            :options="AuthorityOption"
            :props="{ checkStrictly: true, label: 'authName', value: 'authId', disabled: 'disabled', emitPath: false }"
            :show-all-levels="false" filterable />
        </el-form-item>
        <el-form-item label="角色ID" prop="authId">
          <el-input v-model="form.authId" :disabled="dialogType === 'edit'" autocomplete="off" maxlength="15" />
        </el-form-item>
        <el-form-item label="角色姓名" prop="authName">
          <el-input v-model="form.authName" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer v-if="drawer" v-model="drawer" custom-class="auth-drawer" :with-header="false" title="角色配置">
      <el-tabs :before-leave="autoEnter" >
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="资源权限">
          <Datas ref="datas" :auths="table_data" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script setup>
import { auth_create, auth_delete, auth_update, auth_list, auth_copy } from '@/api/auth'
import Menus from '@/view/admin/auth/components/menu.vue'
import Apis from '@/view/admin/auth/components/api.vue'
import Datas from '@/view/admin/auth/components/data.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toDoc } from '@/utils/doc'
import { VideoCameraFilled } from '@element-plus/icons-vue'

defineOptions({ name: 'Authority' })
const mustUint = (rule, value, callback) => !/^[0-9]*[1-9][0-9]*$/.test(value) ? callback(new Error('请输入正整数')) : callback()

const AuthorityOption = ref([{ authId: 0, authName: '根角色' }])
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})

const dialogTitle = ref('新增角色')
const dialogFormVisible = ref(false)
const apiDialogFlag = ref(false)
const copyForm = ref({})

const form = ref({ authId: 0, authName: '', parentId: 0 })
const rules = ref({
  authId: [
    { required: true, message: '请输入角色ID', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
  ],
  authName: [
    { required: true, message: '请输入角色名', trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: '请选择父角色', trigger: 'blur' },
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const table_data = ref([])
const search_info = ref({})

// 查询
const on_table_data = async () => {
  const table = await auth_list({ page: page.value, pageSize: pageSize.value, ...search_info.value })
  if (table.code === 0) {
    table_data.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
on_table_data()
const changeRow = (key, value) => activeRow.value[key] = value
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
// 拷贝角色
const auth_copyFunc = (row) => {
  setOptions()
  dialogTitle.value = '拷贝角色'
  dialogType.value = 'copy'
  for (const k in form.value) {
    form.value[k] = row[k]
  }
  copyForm.value = row
  dialogFormVisible.value = true
}
const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}
// 删除角色
const deleteAuth = (row) => {
  ElMessageBox.confirm('此操作将永久删除该角色, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      const res = await auth_delete({ authId: row.authId })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!'
        })
        if (table_data.value.length === 1 && page.value > 1) {
          page.value--
        }
        on_table_data()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除'
      })
    })
}
// 初始化表单
const authForm = ref(null)
const initForm = () => {
  if (authForm.value) {
    authForm.value.resetFields()
  }
  form.value = {
    authId: 0,
    authName: '',
    parentId: 0
  }
}
// 关闭窗口
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
  apiDialogFlag.value = false
}
// 确定弹窗

const enterDialog = () => {
  authForm.value.validate(async valid => {
    if (valid) {
      form.value.authId = Number(form.value.authId)
      switch (dialogType.value) {
        case 'add':
          {
            const res = await auth_create(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功!'
              })
              on_table_data()
              closeDialog()
            }
          }
          break
        case 'edit':
          {
            const res = await auth_update(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功!'
              })
              on_table_data()
              closeDialog()
            }
          }
          break
        case 'copy': {
          const data = {
            auth: {
              authId: 0,
              authName: '',
              datauthId: [],
              parentId: 0
            },
            oldAuthorityId: 0
          }
          data.auth.authId = form.value.authId
          data.auth.authName = form.value.authName
          data.auth.parentId = form.value.parentId
          data.auth.dataAuthorityId = copyForm.value.dataAuthorityId
          data.oldAuthorityId = copyForm.value.authId
          const res = await auth_copy(data)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '复制成功！'
            })
            on_table_data()
          }
        }
      }

      initForm()
      dialogFormVisible.value = false
    }
  })
}
const setOptions = () => {
  AuthorityOption.value = [
    {
      authId: 0,
      authName: '根角色'
    }
  ]
  setAuthorityOptions(table_data.value, AuthorityOption.value, false)
}
const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
  form.value.authId = String(form.value.authId)
  AuthorityData &&
    AuthorityData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          authId: item.authId,
          authName: item.authName,
          disabled: disabled || item.authId === form.value.authId,
          children: []
        }
        setAuthorityOptions(
          item.children,
          option.children,
          disabled || item.authId === form.value.authId
        )
        optionsData.push(option)
      } else {
        const option = {
          authId: item.authId,
          authName: item.authName,
          disabled: disabled || item.authId === form.value.authId
        }
        optionsData.push(option)
      }
    })
}
// 增加角色
const addAuthority = (parentId) => {
  initForm()
  dialogTitle.value = '新增角色'
  dialogType.value = 'add'
  form.value.parentId = parentId
  setOptions()
  dialogFormVisible.value = true
}
// 编辑角色
const editAuthority = (row) => {
  setOptions()
  dialogTitle.value = '编辑角色'
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  dialogFormVisible.value = true
}

</script>

<style lang="scss">
.auth {
  .el-input-number {
    margin-left: 15px;
    span {
      display: none;
    }
  }
  .el-drawer__body {
    padding: 10px
  }
}

.tree-content {
  margin-top: 10px;
  height: calc(100vh - 158px);
  overflow: auto;
}
</style>
