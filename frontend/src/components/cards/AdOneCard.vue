<script setup>
import { useAdOne } from '@/composable/useAdOne'

const { adOne, loading, error } = useAdOne()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="ad-one">
    <div class="sidebar-img">
      <img :src="adOne.image" :alt="adOne.imageAlt">
    </div>
    <div class="sidebar-hero-text">
      <h3>{{ adOne.heading }}</h3>
      <a :href="adOne.ctaHref">{{ adOne.ctaText }}</a>
    </div>
    <div class="ad">{{ adOne.adLabel }}</div>
  </div>
</template>
