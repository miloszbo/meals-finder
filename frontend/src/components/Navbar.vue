<template>
  <header class="bg-[#0A1118] text-white text-base shadow-md">
    <div class="flex items-center justify-between px-6 py-4 h-20">
      <!-- Logo -->
      <router-link
        to="/"
        class="text-6xl font-bold hover:text-primary transition-colors duration-200"
        style="margin-top: 4px;"
      >
        MealsFinder
      </router-link>

      <!-- Ikony po prawej -->
      <div class="flex items-center gap-6 text-2xl relative">
        <!-- Przepisy -->
        <router-link to="/browser" title="Recipes">
          <i class="fa-solid fa-clipboard-list cursor-pointer hover:text-gray-300"></i>
        </router-link>

        <!-- Profil z dropdownem -->
        <div
          class="relative"
          @mouseenter="showDropdown = true"
          @mouseleave="showDropdown = false"
        >
          <!-- Ikona – przekierowuje -->
          <router-link
            to="/profile?tab=profile"
            title="Profile"
            class="cursor-pointer hover:text-gray-300"
          >
            <i class="fa-solid fa-user"></i>
          </router-link>

          <!-- Dropdown -->
          <div
            v-show="showDropdown"
            class="absolute top-full left-1/2 -translate-x-1/2 bg-base-100 text-white border border-base-content rounded shadow-md w-40 z-50"
          >
            <router-link
              to="/profile?tab=settings"
              class="block px-3 py-1.5 text-sm hover:bg-base-300"
            >
              Dane osobowe
            </router-link>
            <router-link
              to="/profile?tab=privacy"
              class="block px-3 py-1.5 text-sm hover:bg-base-300"
            >
              Prywatność
            </router-link>
            <router-link
              to="/profile?tab=security"
              class="block px-3 py-1.5 text-sm hover:bg-base-300"
            >
              Bezpieczeństwo
            </router-link>
            <router-link
              to="/profile?tab=sharing"
              class="block px-3 py-1.5 text-sm hover:bg-base-300"
            >
              Udostępnianie
            </router-link>
            <router-link
              to="/profile?tab=payments"
              class="block px-3 py-1.5 text-sm hover:bg-base-300"
            >
              Płatności
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue'

const showDropdown = ref(false)
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@700&family=Roboto:wght@700&display=swap');
</style>