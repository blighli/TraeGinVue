<template>
  <div class="w-64 bg-slate-800 text-white">
    <div class="p-6">
      <h2 class="text-xl font-bold text-slate-200">系统菜单</h2>
    </div>
    <nav class="space-y-1 px-3">
      <button
        @click="$emit('change-view', 'employees')"
        :class="[menuItemClass, currentView === 'employees' ? activeClass : inactiveClass]"
      >
        员工列表
      </button>
      <button
        @click="$emit('change-view', 'statistics')"
        :class="[menuItemClass, currentView === 'statistics' ? activeClass : inactiveClass]"
      >
        数据统计
      </button>
    </nav>
  </div>
</template>

<script setup>
const menuItemClass = 'w-full px-4 py-2 text-left rounded-md text-sm font-medium transition-colors'
const activeClass = 'bg-slate-700 text-white'
const inactiveClass = 'text-slate-300 hover:bg-slate-700 hover:text-white'

defineProps({
  currentView: {
    type: String,
    required: true
  }
})

defineEmits(['change-view'])
</script>