<script setup>
import { useTraveling } from '@/composable/useTraveling'

const { traveling, loading, error } = useTraveling()

function positionPopup(e) {
  const rect = e.currentTarget.getBoundingClientRect()
  const popup = e.currentTarget.querySelector('.travler-popup-img')
  popup.style.left = rect.left + rect.width / 2 - 110 + 'px'
  popup.style.top = rect.bottom + 5 + 'px'
}
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="traveling">
    <div class="sidebar-hero-text"><h3>{{ traveling.header }}</h3></div>
    <div class="traveling">
      <div class="img-travler">
        <div v-for="member in traveling.members" :key="member.id" class="travler-images" @mouseenter="positionPopup">
          <img :src="member.thumb" alt="traveler">
          <div class="travler-popup-img">
            <p>{{ traveling.privacyText }}</p>
            <img :src="member.blur" alt="profile">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
