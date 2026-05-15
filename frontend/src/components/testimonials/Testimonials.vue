<script setup>
import { useHoverCard } from '@/composable/useHoverCard'

const { hoverCards, loading, error } = useHoverCard()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="hover-card">
    <div v-for="card in hoverCards" :key="card.imageAlt" class="hover-card-box" :class="card.animationClass">
      <p>"{{ card.quote }}"</p>
      <img :src="card.image" :alt="card.imageAlt">
    </div>
  </div>
</template>
