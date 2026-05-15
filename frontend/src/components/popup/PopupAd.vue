<script setup>
import { ref, computed, onMounted } from 'vue'
import { useHero } from '@/composable/useHero'

const { hero } = useHero()
const visible = ref(false)
const email = ref('')
const interests = ref('meetups')

const popupFeatures = [
  '283 meetups each year',
  '1,000+ city reviews',
  'Live chat with members',
]

const starIcons = computed(() => Array.from({ length: hero.value.left?.stars ?? 0 }))

onMounted(() => {
  setTimeout(() => { visible.value = true }, 1000)
})

function closeOnBackdrop(e) {
  if (e.target === e.currentTarget) visible.value = false
}

function closePopup() {
  visible.value = false
}

</script>

<template>
  <div
    v-if="visible"
    class="background-pop-up-ads"
    role="dialog"
    aria-modal="true"
    aria-labelledby="popup-ad-title"
    @click="closeOnBackdrop"
  >
    <div class="pop-up-ads" @click.stop>
      <button type="button" class="popup-close-btn" aria-label="Close popup" @click="closePopup">×</button>
      <div class="pop-up-ads-content">
        <section class="pop-up-sidebar">
          <div class="pop-up-header-container">
            <img :src="hero.right?.video?.src" :alt="hero.right?.video?.alt" />
          </div>
          <div class="pop-up-sidebar-header-content">
            <p class="pop-up-eyebrow">{{ hero.left?.awardHeader }}</p>
            <h2 id="popup-ad-title" class="pop-up-title">
              {{ hero.left?.emoji }} {{ hero.left?.mainTitle }}
            </h2>
            <p>
              {{ hero.left?.subTitle }}
              Join members who use Nomads.com to find the next city, community, and meetup.
            </p>
          </div>
          <div class="pop-up-sidebar-unlock-feature">
            <span v-for="feature in popupFeatures" :key="feature">{{ feature }}</span>
          </div>
          <div class="pop-up-rating" aria-label="Community rating">
            <div class="pop-up-rating-stars">
              <img v-for="(_, index) in starIcons" :key="index" src="/images/star.svg" alt="" aria-hidden="true" />
            </div>
            <p>Trusted by remote workers since {{ hero.left?.awardFooter?.replace('Since ', '') }}</p>
          </div>
          <div class="pop-up-sidebar-content">
            <p>Explore destinations, meet other nomads, and keep track of your travels in one place.</p>
            <p>Create an account to unlock community recommendations, city notes, and invite-only chats.</p>
          </div>
          <div class="fade-bottom"></div>
        </section>
        <aside class="hero-card-box popup">
          <span class="video-card-box">
            <img :src="hero.right?.video?.src" :alt="hero.right?.video?.alt" class="video-card-background" />
            <img :src="hero.right?.playButton?.src" :alt="hero.right?.playButton?.alt" class="play-btn" />
          </span>
          <input v-model="email" type="email" placeholder="  Type your email..." autocomplete="email" />
          <select v-model="interests" aria-label="Choose your main interest">
            <option value="meetups">I want local meetups</option>
            <option value="cities">I want city rankings</option>
            <option value="chat">I want member chat access</option>
          </select>
          <a :href="hero.right?.ctaHref">{{ hero.right?.ctaText }}</a>
          <p>{{ hero.right?.loginText }}</p>
        </aside>
      </div>
    </div>
  </div>
</template>
