<script setup>
import Logo from './Logo.vue'
import SearchFilter from './SearchFilter.vue'
import Compare from './Compare.vue'
import GridView from './GridView.vue'
import NomadScore from './NomadScore.vue'
import JoinNomad from './JoinNomad.vue'
</script>

<template>
  <div class="filters-bar">
    <div class="filters-left-side">
      <Logo />
      <slot />
      <SearchFilter />
    </div>
    <div class="filters-right-side">
      <Compare />
      <GridView />
      <NomadScore />
      <JoinNomad />
    </div>
  </div>
</template>