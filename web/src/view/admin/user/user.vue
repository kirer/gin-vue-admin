<template>
  <div>
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">新增用户</el-button>
      </div>
      <el-table :data="table_data" row-key="ID">
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column align="left" label="用户名" min-width="150" prop="userName" />
        <el-table-column align="left" label="昵称" min-width="150" prop="nickName" />
        <el-table-column align="left" label="手机号" min-width="180" prop="phone" />
        <el-table-column align="left" label="邮箱" min-width="180" prop="email" />
        <el-table-column align="left" label="用户角色" min-width="200">
          <template #default="scope">
            <el-cascader v-model="scope.row.authIds" :options="auth_options" :show-all-levels="false" collapse-tags
              :props="{ multiple: true, checkStrictly: true, label: 'authName', value: 'authId', disabled: 'disabled', emitPath: false }"
              :clearable="false" @visible-change="(flag) => { changeAuthority(scope.row, flag, 0) }"
              @remove-tag="(removeAuth) => { changeAuthority(scope.row, false, removeAuth) }" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="启用" min-width="150">
          <template #default="scope">
            <el-switch v-model="scope.row.enable" inline-prompt :active-value="1" :inactive-value="2"
              @change="() => { switchEnable(scope.row) }" />
          </template>
        </el-table-column>

        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="on_delete(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete">删除</el-button>
              </template>
            </el-popover>
            <el-button type="primary" link icon="edit" @click="openEdit(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="magic-stick" @click="on_reset_password(scope.row)">重置密码</el-button>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="page_size" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="on_current_changed"
          @size-change="on_size_changed" />
      </div>
    </div>
    <el-dialog v-model="addUserDialog" title="用户" :show-close="false" :close-on-press-escape="false"
      :close-on-click-modal="false">
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
          <el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
            <el-input v-model="userInfo.userName" />
          </el-form-item>
          <el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
            <el-input v-model="userInfo.password" />
          </el-form-item>
          <el-form-item label="昵称" prop="nickName">
            <el-input v-model="userInfo.nickName" />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="userInfo.phone" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email" />
          </el-form-item>
          <el-form-item label="用户角色" prop="authId">
            <el-cascader v-model="userInfo.authIds" style="width:100%" :options="auth_options" :show-all-levels="false"
              :props="{ multiple: true, checkStrictly: true, label: 'authName', value: 'authId', disabled: 'disabled', emitPath: false }"
              :clearable="false" />
          </el-form-item>
          <el-form-item label="启用" prop="disabled">
            <el-switch v-model="userInfo.enable" inline-prompt :active-value="1" :inactive-value="2" />
          </el-form-item>
          <el-form-item label="头像" label-width="80px">
            <div style="display:inline-block" @click="openHeaderChange">
              <img v-if="userInfo.headerImg" alt="头像" class="header-img-box"
                :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http') ? path + userInfo.headerImg : userInfo.headerImg">
              <div v-else class="header-img-box">从媒体库选择</div>
            </div>
          </el-form-item>

        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddUserDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>


<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'
import { auth_list } from '@/api/auth'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'
import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
const path = ref(import.meta.env.VITE_BASE_API + '/')

// 
const page = ref(1)
const total = ref(0)
const page_size = ref(10)
const table_data = ref([])
const on_size_changed = (val) => { page_size.value = val; on_table_data() }
const on_current_changed = (val) => { page.value = val; on_table_data() }
const on_table_data = async () => {
  const table = await getUserList({ page: page.value, pageSize: page_size.value })
  if (table.code === 0) {
    table_data.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    page_size.value = table.data.pageSize
  }
}
// 
const auth_options = ref([])
const set_options = (auth_data) => { auth_options.value = []; set_auth_options(auth_data, auth_options.value) }
const set_auth_options = (auth_data, options_data) => {
  auth_data && auth_data.forEach(item => {
    if (item.children && item.children.length) {
      const option = { authId: item.authId, authName: item.authName, children: [] }
      set_auth_options(item.children, option.children)
      options_data.push(option)
    } else {
      const option = { authId: item.authId, authName: item.authName }
      options_data.push(option)
    }
  })
}
const init = async () => {
  on_table_data()
  const res = await auth_list({ page: 1, pageSize: 999 })
  set_options(res.data.list)
}
init()
// 
watch(() => table_data.value, () => {
  set_auth_ids()
})
const set_auth_ids = () => {
  table_data.value && table_data.value.forEach((user) => {
    user.authIds = user.auths && user.auths.map(i => {
      return i.authId
    })
  })
}
const on_reset_password = (row) => {
  ElMessageBox.confirm('是否将此用户密码重置为123456?', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning', }).then(async () => {
    const res = await resetPassword({ ID: row.ID, })
    res.code === 0 ? ElMessage({ type: 'success', message: res.msg, }) : ElMessage({ type: 'error', message: res.msg, })
  })
}
const on_delete = async (row) => {
  const res = await deleteUser({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await on_table_data()
  }
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authId: '',
  authIds: [],
  enable: 1,
})

const rules = ref({
  userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' },
  ],
  email: [
    { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur' },
  ],
  authId: [
    { required: true, message: '请选择用户角色', trigger: 'blur' }
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async () => {
  userInfo.value.authId = userInfo.value.authIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await on_table_data()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await on_table_data()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async (row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authIds: row.authIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '角色设置成功' })
  } else {
    if (!removeAuth) {
      row.authIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authIds = [removeAuth, ...row.authIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async (row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
    await on_table_data()
    userInfo.value.headerImg = ''
    userInfo.value.authIds = []
  }
}

</script>

<style lang="scss">
.header-img-box {
  @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
</style>
