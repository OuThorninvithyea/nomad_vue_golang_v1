<script setup>
import { useTodayPick } from '@/composable/useTodayPick'

const { todayPick, loading, error } = useTodayPick()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="today-pick">
    <div class="today-label">{{ todayPick.label }}</div>
    <h3>{{ todayPick.heading }}</h3>
    <p>{{ todayPick.views }}</p>
  </div>
</template>
