<script setup>
import { ref } from 'vue'
import Hero from '@/components/hero/Hero.vue'
import Testimonials from "@/components/testimonials/Testimonials.vue"
import FilterBar from "@/components/filterBar/FilterBar.vue"
import FilterSidebar from "@/components/sideBar/FilterSidebar.vue"
import CityCard from "@/components/cards/CityCard.vue"
import AdOneCard from "@/components/cards/AdOneCard.vue"
import AdTwoCard from "@/components/cards/AdTwoCard.vue"
import MeetupsCard from "@/components/cards/MeetupsCard.vue"
import TodayPickCard from "@/components/cards/TodayPickCard.vue"
import TravelingCard from "@/components/cards/TravelingCard.vue"
import NewMembersCard from "@/components/cards/NewMembersCard.vue"
import SuggestCard from "@/components/cards/SuggestCard.vue"
import ChatCard from "@/components/cards/ChatCard.vue"
import TrustedCompany from "@/components/trustedCompanies/TrustedCompany.vue";

const sidebarOpen = ref(false)
</script>

<template>
  <main>
    <Hero v-if="!sidebarOpen" />
    <TrustedCompany />
    <div class="page-filter-sidebar">
      <FilterSidebar v-if="sidebarOpen" />
      <div class="page-main-content">
        <Testimonials />
        <div class="main-container">
          <FilterBar>
            <div class="filter-toggle" @click="sidebarOpen = !sidebarOpen">{{ sidebarOpen ? 'Close filter' : 'Filters' }}</div>
          </FilterBar>
          <div class="main-content" :class="{ 'filter-open': sidebarOpen }">
            <div class="main-boxs" id="main-boxs">
              <CityCard />
              <AdOneCard />
              <MeetupsCard />
              <TodayPickCard />
              <TravelingCard />
              <NewMembersCard />
              <SuggestCard />
              <AdTwoCard />
              <ChatCard />
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
