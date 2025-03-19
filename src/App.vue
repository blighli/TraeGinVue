<template>
  <div class="min-h-screen bg-background flex">
    <Sidebar :current-view="currentView" @change-view="currentView = $event" />
    <div class="flex-1">
      <header class="bg-white shadow-sm border-b border-slate-200">
        <div class="max-w-7xl mx-auto py-6 px-4">
          <h1 class="text-3xl font-bold text-slate-900">
            人事档案管理系统
          </h1>
        </div>
      </header>
      <main class="max-w-7xl mx-auto py-8 px-4">
        <div v-if="loading" class="text-center">
          加载中...
        </div>
        <div v-else-if="error" class="text-red-600 text-center">
          {{ error }}
        </div>
        <template v-else>
          <!-- 搜索栏和操作按钮 -->
          <div v-if="currentView === 'employees'" class="mb-6 flex items-center justify-between">
            <div class="flex space-x-4">
              <input 
                v-model="searchQuery" 
                type="text" 
                placeholder="搜索员工..." 
                class="px-4 py-2 border border-slate-300 rounded-md w-64"
              />
              <select 
                v-model="departmentFilter" 
                class="px-4 py-2 border border-slate-300 rounded-md"
              >
                <option value="">所有部门</option>
                <option v-for="dept in departments" :key="dept" :value="dept">
                  {{ dept }}
                </option>
              </select>
            </div>
            <button 
              @click="showAddDialog = true"
              class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
            >
              添加员工
            </button>
          </div>

          <!-- 数据统计页面 -->
          <Statistics 
            v-if="currentView === 'statistics'"
            :employees="employees"
          />

          <!-- 员工列表页面 -->
          <div v-else-if="currentView === 'employees'" class="bg-white rounded-lg shadow-sm border border-slate-200">
            <table class="w-full">
              <thead>
                <tr class="bg-gradient-to-r from-blue-50 to-indigo-50">
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">工号</th>
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">姓名</th>
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">部门</th>
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">职位</th>
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">入职日期</th>
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">状态</th>
                  <th class="px-6 py-4 text-left text-sm font-semibold text-blue-900">操作</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-blue-50">
                <tr v-for="employee in paginatedEmployees" 
                    :key="employee.id" 
                    class="hover:bg-blue-50 transition-colors duration-150 ease-in-out">
                  <td class="px-6 py-4 text-sm font-medium text-blue-900">{{ employee.id }}</td>
                  <td class="px-6 py-4 text-sm font-medium text-blue-900">{{ employee.name }}</td>
                  <td class="px-6 py-4 text-sm text-blue-700">{{ employee.department }}</td>
                  <td class="px-6 py-4 text-sm text-blue-700">{{ employee.position }}</td>
                  <td class="px-6 py-4 text-sm text-blue-700">{{ employee.joinDate }}</td>
                  <td class="px-6 py-4">
                    <span :class="[
                      'px-3 py-1 text-sm font-medium rounded-full',
                      employee.status === '在职' 
                        ? 'bg-green-100 text-green-700 ring-1 ring-green-700/10' 
                        : 'bg-red-100 text-red-700 ring-1 ring-red-700/10'
                    ]">
                      {{ employee.status }}
                    </span>
                  </td>
                  <td class="px-6 py-4">
                    <div class="flex items-center space-x-2">
                      <button 
                        @click="editEmployee(employee)"
                        class="text-blue-600 hover:text-blue-800"
                      >
                        编辑
                      </button>
                      <button 
                        @click="deleteEmployee(employee)"
                        class="text-red-600 hover:text-red-800"
                      >
                        删除
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
            
            <!-- 分页控件 -->
            <div class="flex items-center justify-between px-6 py-4 border-t border-slate-200">
              <div class="text-sm text-slate-700">
                显示 {{ startIndex + 1 }} - {{ Math.min(endIndex, filteredEmployees.length) }} 条，共 {{ filteredEmployees.length }} 条
              </div>
              <div class="flex space-x-2">
                <button 
                  @click="currentPage--" 
                  :disabled="currentPage === 1"
                  class="px-3 py-1 text-sm border rounded-md"
                  :class="currentPage === 1 ? 'text-slate-400 border-slate-200' : 'text-slate-700 border-slate-300 hover:bg-slate-50'"
                >
                  上一页
                </button>
                <button 
                  v-for="page in totalPages" 
                  :key="page"
                  @click="currentPage = page"
                  class="px-3 py-1 text-sm border rounded-md"
                  :class="currentPage === page ? 'bg-blue-50 text-blue-600 border-blue-200' : 'text-slate-700 border-slate-300 hover:bg-slate-50'"
                >
                  {{ page }}
                </button>
                <button 
                  @click="currentPage++" 
                  :disabled="currentPage === totalPages"
                  class="px-3 py-1 text-sm border rounded-md"
                  :class="currentPage === totalPages ? 'text-slate-400 border-slate-200' : 'text-slate-700 border-slate-300 hover:bg-slate-50'"
                >
                  下一页
                </button>
              </div>
            </div>
          </div>

          <!-- 添加员工对话框 -->
          <div v-if="showAddDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
            <div class="bg-white rounded-lg p-6 w-[600px]">
              <h3 class="text-lg font-semibold mb-4">添加新员工</h3>
              <AddEmployee 
                @add-employee="handleAddEmployee"
                @cancel="showAddDialog = false"
                :existing-ids="employees.map(e => e.id)"
              />
            </div>
          </div>

          <!-- 编辑对话框 -->
          <div v-if="showEditDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
            <div class="bg-white rounded-lg p-6 w-[600px]">
              <h3 class="text-lg font-semibold mb-4">编辑员工信息</h3>
              <AddEmployee 
                :initial-data="editingEmployee"
                :is-editing="true"
                @add-employee="handleEditEmployee"
                @cancel="showEditDialog = false"
              />
            </div>
          </div>
        </template>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { employeeApi } from './api/employee'
import Statistics from './components/Statistics.vue'
import Sidebar from './components/Sidebar.vue'
import AddEmployee from './components/AddEmployee.vue'

// 视图状态
const currentView = ref('employees')
const searchQuery = ref('')
const departmentFilter = ref('')
const currentPage = ref(1)
const pageSize = 10

// 弹窗状态
const showAddDialog = ref(false)
const showEditDialog = ref(false)
const editingEmployee = ref(null)

// 部门列表
const departments = computed(() => {
  return [...new Set(employees.value.map(e => e.department))]
})

// 过滤后的员工列表
const filteredEmployees = computed(() => {
  return employees.value.filter(emp => {
    const matchesSearch = emp.name.includes(searchQuery.value) || 
                         emp.id.includes(searchQuery.value)
    const matchesDepartment = !departmentFilter.value || 
                             emp.department === departmentFilter.value
    return matchesSearch && matchesDepartment
  })
})

// 分页数据
const totalPages = computed(() => Math.ceil(filteredEmployees.value.length / pageSize))
const startIndex = computed(() => (currentPage.value - 1) * pageSize)
const endIndex = computed(() => startIndex.value + pageSize)
const paginatedEmployees = computed(() => {
  return filteredEmployees.value.slice(startIndex.value, endIndex.value)
})

// 员工操作处理函数
const handleAddEmployee = async (employee) => {
  await createEmployee(employee)
  showAddDialog.value = false
}

const handleEditEmployee = async (employee) => {
  await updateEmployee(employee.id, employee)
  showEditDialog.value = false
  editingEmployee.value = null
}

const editEmployee = (employee) => {
  editingEmployee.value = { ...employee }
  showEditDialog.value = true
}

const employees = ref([])
const loading = ref(false)
const error = ref(null)

// 获取员工数据
const fetchEmployees = async () => {
  loading.value = true
  try {
    const response = await employeeApi.getEmployees()
    employees.value = response.data
  } catch (err) {
    error.value = '获取数据失败'
    console.error(err)
  } finally {
    loading.value = false
  }
}

// 创建员工
const createEmployee = async (employee) => {
  try {
    await employeeApi.createEmployee(employee)
    await fetchEmployees()
  } catch (err) {
    console.error(err)
  }
}

// 更新员工
const updateEmployee = async (id, employee) => {
  try {
    await employeeApi.updateEmployee(id, employee)
    await fetchEmployees()
  } catch (err) {
    console.error(err)
  }
}

// 删除员工
const deleteEmployee = async (id) => {
  try {
    await employeeApi.deleteEmployee(id)
    await fetchEmployees()
  } catch (err) {
    console.error(err)
  }
}

// 组件挂载时获取数据
onMounted(fetchEmployees)
</script>

