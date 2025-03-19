import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080/api'
})

export const employeeApi = {
  // 获取所有员工
  getEmployees() {
    return api.get('/employees')
  },

  // 创建员工
  createEmployee(employee) {
    return api.post('/employees', employee)
  },

  // 更新员工
  updateEmployee(id, employee) {
    return api.put(`/employees/${id}`, employee)
  },

  // 删除员工
  deleteEmployee(id) {
    return api.delete(`/employees/${id}`)
  }
}