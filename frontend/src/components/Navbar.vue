<template>
  <header class="bg-[#0A1118] text-white text-base">
    <div class="flex items-center justify-between p-4">
      <!-- Logo -->
        <router-link to="/" class="text-5xl font-bold mb-8 hover:text-primary transition-colors duration-200">
          MealsFinder
        </router-link>

      <!-- Pasek wyszukiwania -->
      <div class="hidden md:flex justify-center mx-4" style="width: 1088px; height: 48px;">
        <div class="flex w-full bg-[#1a1f26] rounded-full overflow-hidden h-full">
          <input
            type="text"
            placeholder="Search..."
            class="w-full bg-transparent text-white px-4 outline-none"
          />
          <button class="px-4">
            <i class="fas fa-search"></i>
          </button>
        </div>
      </div>

      <!-- Ikony po prawej -->
      <div class="flex items-center gap-4 text-lg">
       <router-link to="/browser" title="Recipes">
      <i class="fa-solid fa-clipboard-list cursor-pointer hover:text-gray-300"></i>
       </router-link>

      
     <!--
     <a href="/favorites" title="Favorites">
     <i class="fa-solid fa-heart cursor-pointer hover:text-gray-300"></i>
      </a>
     -->
      <router-link to="/login" title="Profile">
        <i class="fa-solid fa-user cursor-pointer hover:text-gray-300"></i>
      </router-link>

      </div>
    </div>

    <!-- Kategorie -->
    <nav class="bg-[#0A1118] border-t border-[#1f2a34] flex flex-wrap justify-center gap-6 px-4 py-2">
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

<script setup>
const categories = [
  "Europe",
  "Asian",
  "North American",
  "South American",
  "African",
  "Oceanian"
];
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@700&family=Roboto:wght@700&display=swap');
</style>
