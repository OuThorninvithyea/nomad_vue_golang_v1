<script setup>
import { useNewMembers } from '@/composable/useNewMembers'

const { newMembers, loading, error } = useNewMembers()

function positionPopup(e) {
  const rect = e.currentTarget.getBoundingClientRect()
  const popup = e.currentTarget.querySelector('.new-members-popup-img')
  popup.style.left = rect.left + rect.width / 2 - 110 + 'px'
  popup.style.top = rect.bottom + 5 + 'px'
}
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="new-member">
    <div class="sidebar-hero-text"><h3>{{ newMembers.header }}</h3></div>
    <div class="new-members-container">
      <div v-for="member in newMembers.members" :key="member.id" class="new-members" @mouseenter="positionPopup">
        <div class="new-members-img">
          <img :src="member.thumb" alt="New member">
        </div>
        <div class="new-members-popup-img">
          <p>{{ newMembers.privacyText }}</p>
          <img :src="member.blur" alt="profile">
        </div>
      </div>
    </div>
  </div>
</template>
