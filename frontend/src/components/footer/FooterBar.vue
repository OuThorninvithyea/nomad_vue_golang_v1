<script setup>
import { ref } from 'vue'

const visible = ref(true)
</script>

<template>
  <footer v-if="visible">
    <div class="footer">
      <div class="footer-content">
        <div class="footer-text">🌍 Join a community of remote workers living around the world</div>
        <a href="#">Join Nomads.com</a>
      </div>
      <div class="footer-close-btn" @click="visible = false">×</div>
    </div>
  </footer>
</template>