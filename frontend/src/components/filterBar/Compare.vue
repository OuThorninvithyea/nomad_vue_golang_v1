<script setup>
import { ref } from 'vue'

const isActive = ref(false)
</script>

<template>
  <div class="tooltip compare-filter buttons-filers-style" data-tooltip="Toggle compare view" :style="isActive ? { backgroundColor: 'var(--color-primary)', color: 'var(--color-white-default)' } : {}"
    @click="isActive = !isActive">
    <div class="compare-label">{{ isActive ? 'Exit Compare' : 'Compare' }}</div>
  </div>
</template>