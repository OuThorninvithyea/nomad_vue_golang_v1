<script setup>
import { useSuggest } from '@/composable/useSuggest'

const { suggest, loading, error } = useSuggest()
</script>

<template>
  <div v-if="loading">Loading...</div>
  <div v-else-if="error">{{ error }}</div>
  <div v-else class="boxs" id="suggest">
    <div class="lable-suggest">{{ suggest.label }}</div>
    <div class="header-box">
      <div class="list-number">
        <span class="text-underline">{{ suggest.rank }}</span>
      </div>
      <div class="internet-speed">
        <div class="speed-wifi">
          <span>{{ suggest.speed }}</span>
          <span class="mbps">Mbps</span>
        </div>
      </div>
    </div>
    <div class="hero-box">
      <span class="country">{{ suggest.country }}</span>
      <span class="city">{{ suggest.city }}</span>
    </div>
    <div class="footer-box">
      <div class="weather">
        {{ suggest.weatherEmoji }}
        <div class="temp"><span>{{ suggest.temp }}°</span></div>
        {{ suggest.tempEmoji }}
        {{ suggest.aqiEmoji }}
        <div class="aqi">AQI <span>{{ suggest.aqi }}</span></div>
      </div>
      <div class="cost-living">{{ suggest.cost }}</div>
    </div>
  </div>
</template>
