const STORAGE_KEY = 'employees'

const generateEmployees = () => {
  const departments = ['研发部', '市场部', '财务部', '人事部', '运营部']
  const positions = ['工程师', '经理', '总监', '专员', '主管']
  const statuses = ['在职', '离职']
  const names = [
    '张伟', '王芳', '李娜', '刘洋', '陈明', '杨华', '赵静', '周鹏', '吴婷', '郑阳',
    '黄河', '孙雪', '马超', '朱峰', '胡晓', '郭琳', '何磊', '高帆', '林涛', '徐梦',
    '梁晨', '韩雷', '叶航', '邓云', '曾波', '傅晴', '萧风', '凌霄', '柳依', '谭松',
    '姜山', '钱江', '崔星', '白露', '金鑫', '魏青', '秦明', '唐风', '范海', '方华',
    '石峰', '贾明', '夏雨', '彭云', '董静', '袁野', '潘阳', '江涛', '梅晓', '龙飞'
  ]
  
  return Array.from({ length: 100 }, (_, i) => ({
    id: `MS${String(i + 1).padStart(3, '0')}`,
    name: names[Math.floor(Math.random() * names.length)],
    department: departments[Math.floor(Math.random() * departments.length)],
    position: positions[Math.floor(Math.random() * positions.length)],
    joinDate: `2020-${String(Math.floor(Math.random() * 12) + 1).padStart(2, '0')}-${String(Math.floor(Math.random() * 28) + 1).padStart(2, '0')}`,
    status: statuses[Math.floor(Math.random() * statuses.length)]
  }))
}

const initialEmployees = generateEmployees()

export const employeeService = {
  getAll() {
    const data = localStorage.getItem(STORAGE_KEY)
    if (!data) {
      // 如果本地存储中没有数据，初始化并保存
      this.save(initialEmployees)
      return initialEmployees
    }
    return JSON.parse(data)
  },

  save(employees) {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(employees))
  },

  create(employee) {
    const employees = this.getAll()
    employees.push(employee)
    this.save(employees)
    return employee
  },

  update(id, updatedEmployee) {
    const employees = this.getAll()
    const index = employees.findIndex(e => e.id === id)
    if (index !== -1) {
      employees[index] = { ...employees[index], ...updatedEmployee }
      this.save(employees)
      return employees[index]
    }
    return null
  },

  delete(id) {
    const employees = this.getAll()
    const index = employees.findIndex(e => e.id === id)
    if (index !== -1) {
      employees.splice(index, 1)
      this.save(employees)
      return true
    }
    return false
  }
}