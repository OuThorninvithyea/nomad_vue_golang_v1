<script setup>
import { useAdTwo } from '@/composable/useAdTwo'

const { adTwo, loading, error } = useAdTwo()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="ad-two">
    <div class="sidebar-img">
      <img :src="adTwo.image" :alt="adTwo.imageAlt">
    </div>
    <div class="sidebar-hero-text">
      <h3>{{ adTwo.heading }}</h3>
      <a :href="adTwo.ctaHref">{{ adTwo.ctaText }}</a>
    </div>
    <div class="ad">{{ adTwo.adLabel }}</div>
  </div>
</template>
