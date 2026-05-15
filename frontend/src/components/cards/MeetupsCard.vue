<script setup>
import { useMeetups } from '@/composable/useMeetups'

const { meetups, loading, error } = useMeetups()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="sidebar-boxs" id="meetups">
    <div class="sidebar-hero-text"><h3>{{ meetups.header }}</h3></div>
    <div class="meetups">
      <div v-for="day in meetups.days" :key="day.schedule" class="meet-up-day meet-up-bottom-border">
        <div class="schedule-meetup">
          <div class="sechudal-with-lable">
            {{ day.schedule }} <span class="meet-up-lable">{{ day.rsvps }}</span>
          </div>
        </div>
        <div class="profile-meetup">
          <div class="two-profile">
            <img :src="day.cityImage" :alt="day.cityImageAlt">
            <img v-for="img in day.attendeeImages" :key="img" :src="img" alt="Profile Image">
          </div>
        </div>
      </div>
    </div>
    <a class="see-meeting" :href="meetups.footerHref">{{ meetups.footerText }}</a>
  </div>
</template>
