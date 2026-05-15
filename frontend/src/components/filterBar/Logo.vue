<script setup>
import { useLogoDropdown } from '@/composable/useLogoDropdown'

const { logoDropdown } = useLogoDropdown()
</script>

<template>
  <div class="logo-toggle">
    <div class="logo" tabindex="0">
      <img src="/images/logo-trans.png" alt="nomad-logo">
      <span class="drop-down-toggle-filter"><img src="/images/drop-down.svg" alt=""></span>
    </div>
    <div class="logo-menu" id="logo-dropdown-menu">
      <div v-for="(section, sIdx) in logoDropdown" :key="sIdx" class="menu-section">
        <h4 v-if="section.title">{{ section.title }}</h4>
        <div :class="section.sectionClass || 'menu-grid'">
          <a v-for="(item, iIdx) in section.items" :key="iIdx" href="#">
            {{ item.emoji }} {{ item.text }}
          </a>
        </div>
      </div>
    </div>
  </div>
</template>
