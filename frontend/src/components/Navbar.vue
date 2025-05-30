<style scoped>
  @import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@700&family=Roboto:wght@700&display=swap');
</style>

<script setup>
import UserMenu from './UserMenu.vue'
// List of recipe categories displayed in navbar
const categories = [
  "Europe",
  "Asian",
  "North American",
  "South American",
  "African",
  "Oceanian"
];
</script>

<template>
  <header class="bg-[#211C54] text-white text-base shadow-md">
    <!-- Top bar: Logo on the left, user/recipe icons on the right -->
    <div class="flex items-center justify-between px-6 py-4 h-20">
      <!-- Logo - links to homepage -->
      <router-link to="/" class="text-5xl font-bold hover:text-primary transition-colors duration-200">
        MealsFinder
      </router-link>

      <!-- Right side icons: Recipes and User menu -->
      <div class="flex items-center gap-4 text-lg">
        <!-- Recipes icon -->
        <router-link to="/browser" title="Recipes">
          <i class="fa-solid fa-clipboard-list cursor-pointer hover:text-gray-300"></i>
        </router-link>
        <!-- User dropdown menu (login, register, profile etc.) -->
        <UserMenu />
      </div>
    </div>

    <!-- Categories bar (centered navigation links for food categories) -->
    <nav class="bg-[#060606] border-t border-[#1f2a34] flex flex-wrap justify-center gap-6 px-4 py-2">
      <a
        v-for="cat in categories"
        :key="cat"
        :href="`/categories/${cat.toLowerCase().replace(/ /g, '-')}`"
        class="text-white"
        style="font-family: 'Roboto', sans-serif; font-weight: 700; font-size: 24px; line-height: 100%; text-align: center; vertical-align: middle;"
      >
        {{ cat }}
      </a>
    </nav>
  </header>
</template>
