<template>
  <div class="bg-white rounded-lg shadow-sm border border-slate-200 p-6">
    <h2 class="text-xl font-semibold text-slate-900 mb-6">添加新员工</h2>
    <form @submit.prevent="handleSubmit" class="space-y-6">
      <div class="grid grid-cols-2 gap-6">
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">工号</label>
          <input v-model="form.id" type="text" class="w-full px-3 py-2 border border-slate-300 rounded-md" placeholder="请输入工号" />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">姓名</label>
          <input v-model="form.name" type="text" class="w-full px-3 py-2 border border-slate-300 rounded-md" placeholder="请输入姓名" />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">部门</label>
          <select v-model="form.department" class="w-full px-3 py-2 border border-slate-300 rounded-md">
            <option value="">请选择部门</option>
            <option value="研发部">研发部</option>
            <option value="市场部">市场部</option>
            <option value="财务部">财务部</option>
          </select>
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">职位</label>
          <input v-model="form.position" type="text" class="w-full px-3 py-2 border border-slate-300 rounded-md" placeholder="请输入职位" />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">入职日期</label>
          <input v-model="form.joinDate" type="date" class="w-full px-3 py-2 border border-slate-300 rounded-md" />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-700">状态</label>
          <select v-model="form.status" class="w-full px-3 py-2 border border-slate-300 rounded-md">
            <option value="在职">在职</option>
            <option value="离职">离职</option>
          </select>
        </div>
      </div>
      <div class="flex justify-end space-x-4">
        <button type="button" @click="$emit('cancel')" class="px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100 border border-slate-300 rounded-md">
          取消
        </button>
        <button type="submit" class="px-4 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-md">
          提交
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { reactive, onMounted } from 'vue'

const props = defineProps({
  initialData: {
    type: Object,
    default: () => null
  },
  isEditing: {
    type: Boolean,
    default: false
  },
  existingIds: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['add-employee', 'cancel'])

const form = reactive({
  id: '',
  name: '',
  department: '',
  position: '',
  joinDate: '',
  status: '在职'
})

onMounted(() => {
  if (props.initialData) {
    Object.assign(form, props.initialData)
  }
})

const validateForm = () => {
  const errors = []
  if (!form.id) errors.push('请输入工号')
  if (!form.name) errors.push('请输入姓名')
  if (!form.department) errors.push('请选择部门')
  if (!form.position) errors.push('请输入职位')
  if (!form.joinDate) errors.push('请选择入职日期')
  
  if (!props.isEditing && props.existingIds.includes(form.id)) {
    errors.push('工号已存在')
  }

  return errors
}

const handleSubmit = () => {
  const errors = validateForm()
  if (errors.length > 0) {
    alert(errors.join('\n'))
    return
  }

  emit('add-employee', { ...form })
  
  if (!props.isEditing) {
    // 只在新增时重置表单
    form.id = ''
    form.name = ''
    form.department = ''
    form.position = ''
    form.joinDate = ''
    form.status = '在职'
  }
}
</script>