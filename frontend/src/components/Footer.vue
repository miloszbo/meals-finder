<template>
  <!-- GŁÓWNY WRAPPER FOOTERA: pełna szerokość, tło, kolor tekstu -->
  <footer class="w-full bg-[#111921] text-gray-400 text-sm">
    <!-- Kontener na treść: max szerokość, padding, flex na kolumny/wiersze -->
    <div class="max-w-screen-xl mx-auto flex flex-col md:flex-row justify-between items-center px-8 py-4">
      
      <!-- logo i prawa autorskie -->
      <div>
        <span class="font-bold text-white">MealsFinder</span>
        <span class="ml-2">&copy; 2025 Meals Finder</span>
        <span class="ml-2">All rights reserved</span>
      </div>
      
      <!-- PRAWA SEKCJA: linki do stron/info -->
      <div class="flex gap-6 mt-3 md:mt-0">
        <a href="#" class="hover:text-white transition">About</a>
        <a href="#" class="hover:text-white transition">Community</a>
        <a href="#" class="hover:text-white transition">FAQ</a>
      </div>
    </div>
  </footer>
</template>
