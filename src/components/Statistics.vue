<template>
  <div class="space-y-6">
    <!-- 统计卡片 -->
    <div class="grid grid-cols-4 gap-6">
      <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
        <h3 class="text-sm font-medium text-slate-500">员工总数</h3>
        <p class="mt-2 text-3xl font-bold text-slate-900">{{ totalEmployees }}</p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
        <h3 class="text-sm font-medium text-slate-500">在职人数</h3>
        <p class="mt-2 text-3xl font-bold text-green-600">{{ activeEmployees }}</p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
        <h3 class="text-sm font-medium text-slate-500">离职人数</h3>
        <p class="mt-2 text-3xl font-bold text-red-600">{{ inactiveEmployees }}</p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
        <h3 class="text-sm font-medium text-slate-500">部门数量</h3>
        <p class="mt-2 text-3xl font-bold text-blue-600">{{ departmentCount }}</p>
      </div>
    </div>

    <!-- 部门分布 -->
    <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
      <h3 class="text-lg font-medium text-slate-900 mb-4">部门人员分布</h3>
      <div class="space-y-4">
        <div v-for="(count, dept) in departmentStats" :key="dept" class="flex items-center">
          <span class="w-24 text-sm text-slate-600">{{ dept }}</span>
          <div class="flex-1 h-4 bg-slate-100 rounded-full overflow-hidden">
            <div 
              class="h-full bg-blue-500 transition-all duration-500"
              :style="{ width: `${(count / totalEmployees) * 100}%` }"
            ></div>
          </div>
          <span class="ml-4 text-sm font-medium text-slate-700">{{ count }}人</span>
        </div>
      </div>
    </div>

    <!-- 入职日期折线图 -->
    <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
      <h3 class="text-lg font-medium text-slate-900 mb-4">入职日期分布</h3>
      <div class="h-[400px]">
        <v-chart class="w-full h-full" :option="joinDateChartOption" />
      </div>
    </div>

    <!-- 入职时间分布 -->
    <div class="bg-white p-6 rounded-lg shadow-sm border border-slate-200">
      <h3 class="text-lg font-medium text-slate-900 mb-4">入职时间分布</h3>
      <div class="space-y-4">
        <div v-for="(count, year) in joinYearStats" :key="year" class="flex items-center">
          <span class="w-24 text-sm text-slate-600">{{ year }}年</span>
          <div class="flex-1 h-4 bg-slate-100 rounded-full overflow-hidden">
            <div 
              class="h-full bg-green-500 transition-all duration-500"
              :style="{ width: `${(count / totalEmployees) * 100}%` }"
            ></div>
          </div>
          <span class="ml-4 text-sm font-medium text-slate-700">{{ count }}人</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { BarChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent } from 'echarts/components'
import VChart from 'vue-echarts'

// 注册必要的 ECharts 组件
use([
  CanvasRenderer,
  BarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent
])

const props = defineProps({
  employees: {
    type: Array,
    required: true
  }
})

// 基础统计
const totalEmployees = computed(() => props.employees.length)
const activeEmployees = computed(() => props.employees.filter(e => e.status === '在职').length)
const inactiveEmployees = computed(() => props.employees.filter(e => e.status === '离职').length)

// 部门统计
const departmentStats = computed(() => {
  return props.employees.reduce((acc, emp) => {
    acc[emp.department] = (acc[emp.department] || 0) + 1
    return acc
  }, {})
})

const departmentCount = computed(() => Object.keys(departmentStats.value).length)

// 入职年份统计
const joinYearStats = computed(() => {
  return props.employees.reduce((acc, emp) => {
    const year = emp.joinDate.split('-')[0]
    acc[year] = (acc[year] || 0) + 1
    return acc
  }, {})
})

// 计算入职日期分布数据
const joinDateStats = computed(() => {
  const stats = props.employees.reduce((acc, emp) => {
    const monthDate = emp.joinDate.substring(0, 7)
    acc[monthDate] = (acc[monthDate] || 0) + 1
    return acc
  }, {})
  
  const sortedDates = Object.keys(stats).sort()
  const counts = sortedDates.map(date => stats[date])
  
  // 计算平均值
  const average = counts.reduce((sum, count) => sum + count, 0) / counts.length
  
  return {
    dates: sortedDates.map(date => date + '月'),
    counts: counts,
    average: average
  }
})

// 生成柱状图配置
const joinDateChartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    formatter: '{b}<br />入职人数：{c}人<br />平均值：' + joinDateStats.value.average.toFixed(1) + '人'
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: joinDateStats.value.dates,
    axisLabel: {
      rotate: 30
    }
  },
  yAxis: {
    type: 'value',
    name: '入职人数',
    splitLine: {
      show: true,
      lineStyle: {
        type: 'dashed'
      }
    }
  },
  series: [{
    name: '入职人数',
    type: 'bar',
    data: joinDateStats.value.counts.map(value => ({
      value: value,
      itemStyle: {
        color: value > joinDateStats.value.average ? '#ef4444' : '#8b5cf6'
      }
    })),
    barWidth: '60%',
    markLine: {
      silent: true,
      symbol: ['none', 'none'],
      data: [{
        type: 'average',  // 改用 type: 'average' 而不是 yAxis
        name: '平均值',
        label: {
          show: true,
          position: 'start',
          formatter: '平均值: {c}人',
          color: '#475569',
          fontSize: 12,
          fontWeight: 'bold',
          backgroundColor: '#fff',
          padding: [4, 8]
        },
        lineStyle: {
          color: '#475569',
          type: 'dashed',
          width: 1.5,
          cap: 'round'
        }
      }]
    }
  }]
}))
</script>