<script setup>
import { useHero } from '@/composable/useHero'

const { hero } = useHero()
</script>

<template>
  <div class="join-filter buttons-filers-style">
    <a :href="hero.right?.ctaHref">{{ hero.right?.ctaText }}</a>
  </div>
</template>
