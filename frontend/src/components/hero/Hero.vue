<script setup>
import { useHero } from '@/composable/useHero'

const {hero, loading, error} = useHero()
</script>

<template>
  <div v-if="loading" class="hero-loading">Loading...</div>
  <div v-else-if="error" class="hero-error">{{ error }}</div>
  <div v-else class="hero">
    <div class="background-video">
      <video autoplay muted loop playsinline>
        <source :src="hero.backgroundVideo.src" :type="hero.backgroundVideo.type" />
      </video>
    </div>
    <div class="hero-content">
      <!-- Left side -->
      <div class="award">
        <div class="award-texts" :style="{ backgroundImage: `url(${hero.left.laurelImage})` }">
          <p>{{ hero.left.awardHeader }}</p>
          <span class="star">
            <img v-for="s in hero.left.stars" :key="s" src="/images/star.svg" alt="star"/>
          </span>
          <p>{{ hero.left.awardFooter }}</p>
        </div>
        <h2>{{ hero.left.emoji }} {{ hero.left.mainTitle }}</h2>
        <h3>{{ hero.left.subTitle }}</h3>
        <div class="remote-traveler">
          <img v-for="photo in hero.left.remoteTravelers" :key="photo" :src="photo" alt="profile-photo"/>
        </div>
        <div class="benefits-list">
          <a v-for="benefit in hero.left.benefits" :key="benefit.href" :href="benefit.href" >{{ benefit.emoji }} <span class="text-underline">{{ benefit.text }}</span> {{ benefit.context }}</a>
        </div>
      </div>

      <!-- Right side -->
      <div class="hero-card-box">
        <span class="video-card-box">
          <img :src="hero.right.video.src" :alt="hero.right.video.alt" class="video-card-background"/>
          <img :src="hero.right.playButton.src":alt="hero.right.playButton.alt" class="play-btn"  />
        </span>
        <input type="text" :placeholder="`\u00A0\u00A0 ${hero.right.inputPlaceholder}`"/>
        <a :href="hero.right.ctaHref">{{ hero.right.ctaText }}</a>
        <p>{{ hero.right.loginText }}</p>
      </div>
    </div>
    <div class="curve-line">
      <svg viewBox="0 0 1440 60" preserveAspectRatio="none">
        <path d="M1440,21.2101911 L1440,120 L0,120 L0,21.2101911 C120,35.0700637 240,42 360,42 C480,42 600,35.0700637 720,21.2101911 C808.32779,12.416393 874.573633,6.87702029 918.737528,4.59207306 C972.491685,1.8109458 1026.24584,0.420382166 1080,0.420382166 C1200,0.420382166 1320,7.35031847 1440,21.2101911 Z"></path>
      </svg>
    </div>
  </div>
</template>
