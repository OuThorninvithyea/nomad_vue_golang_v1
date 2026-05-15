<script setup>
import { useChat } from '@/composable/useChat'

const { chat, loading, error } = useChat()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="chat">
    <div class="chat-box">
      <div class="chat-header">
        <div class="chat-title">{{ chat.header }}</div>
        <div class="chat-img"><img :src="chat.image" :alt="chat.imageAlt"></div>
      </div>
      <div class="chat-footer">
        <a :href="chat.ctaHref">{{ chat.ctaText }}</a>
      </div>
    </div>
  </div>
</template>
