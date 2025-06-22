<template>
  <!-- Full-width footer bar -->
  <footer class="mt-auto w-full bg-[#060606] text-[#DEDBD7] text-xs border-t border-[#222]">
    <div class="max-w-screen-xl mx-auto flex flex-col sm:flex-row items-start sm:items-center justify-between px-8 py-4 w-full">
      
      <!-- Left side: Logo and copyright -->
      <div class="flex flex-col sm:flex-row items-start sm:items-center gap-2">
        <span class="text-base font-normal mr-2" style="font-family: 'Montserrat', 'Roboto', sans-serif;">
          MealsFinder
        </span>
        <div class="flex flex-col leading-tight">
          <span class="whitespace-nowrap">@ 2025 Meals Finder</span>
          <span class="whitespace-nowrap">All rights reserved</span>
        </div>
      </div>
      
      <!-- Right side: Links and photo credits -->
      <div class="flex flex-col sm:flex-row items-center gap-3 sm:gap-6 mt-2 sm:mt-0">
        <!-- Photo credits -->
        <span class="text-[11px] italic text-[#bdbdbd] flex items-center gap-1">
          <span class="text-white/50">📸</span>
          Photo sources:
          <a href="https://unsplash.com" target="_blank" rel="noopener" class="underline hover:text-white transition ml-1">Unsplash</a>
          and
          <a href="https://www.pexels.com" target="_blank" rel="noopener" class="underline hover:text-white transition ml-1">Pexels</a>
        </span>
        <!-- Separator -->
        <span class="hidden sm:inline mx-2 text-[#333]">|</span>
        <!-- Main navigation links -->
        <a href="#" class="hover:text-white transition font-semibold">About</a>
        <a href="#" class="hover:text-white transition font-semibold">Community</a>
        <a href="#" class="hover:text-white transition font-semibold">FAQ</a>
      </div>
    </div>
  </footer>
</template>

