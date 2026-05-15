<script setup>
import { useTrustedCompany } from '@/composable/useTrustedCompany'

const { trustedCompany, loading, error } = useTrustedCompany()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="trusted-company">
    <p>{{ trustedCompany.header }}</p>
    <img v-for="company in trustedCompany.companies" :key="company.src" :src="company.src" :alt="company.alt"/>
  </div>
</template>
