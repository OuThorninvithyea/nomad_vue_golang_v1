<script setup>
import { ref } from 'vue'
import { useSearchDropdown } from '@/composable/useSearchDropdown'

const { searchDropdown } = useSearchDropdown()
const placeholder = ref('Search or filter...')
</script>

<template>
  <div class="search-filter">
    <div class="search-dropdown">
      <div v-for="column in searchDropdown" :key="column.id" class="dropdown-search-column">
        <div v-for="item in column.items" :key="item.text" class="drop-down-info tooltip" :data-tooltip="item.tooltip" v-html="item.text"/>
      </div>
    </div>
    <input type="text" :placeholder="placeholder" class="search-btn-filter"
      @focus="placeholder = 'Type to search'"
      @blur="placeholder = 'Search or filter'">
    <button class="search-btn">
      <img src="/images/plus.svg" alt="plus-search">
    </button>
  </div>
</template>
