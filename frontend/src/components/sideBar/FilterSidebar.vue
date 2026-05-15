<script setup>
import { useFilterSidebar } from '@/composable/useFilterSidebar'

const { filterSidebar, loading, error } = useFilterSidebar()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="filter-sidebar" style="display: flex;">
    <template v-for="section in filterSidebar" :key="section.header">
      <div v-if="section.header" class="header-filter">{{ section.header }}</div>
      <template v-if="section.rows">
        <template v-for="row in section.rows" :key="row.id || row.type || row.classes">
          <select v-if="row.type === 'select'" :id="row.id" :class="row.classes || 'row-filter'">
            <option value="">select your passport</option>
            <option v-for="opt in row.options" :key="opt" :value="opt">{{ opt }}</option>
          </select>
          <div v-else-if="row.type === 'divider'" class="plus-filter">{{ row.text }}</div>
          <div v-else :class="row.classes || 'row-filter'">
            <span v-for="item in row.items" :key="item.text" :class="item.classes" :data-tooltip="item.tooltip || undefined">{{ item.text }}</span>
          </div>
        </template>
      </template>
      <div v-if="section.grid" :class="section.grid.classes">
        <component v-for="item in section.grid.items" :key="item.text" :is="item.element || 'span'" :class="item.classes" :data-tooltip="item.tooltip || undefined">{{ item.text }}</component>
      </div>
    </template>
  </div>
</template>
