<template>
    <div class="navbar bg-base-100 shadow-md px-4">
      <!-- Lewa strona (logo) -->
      <div class="flex-1">
        <a class="btn btn-ghost text-xl">MojaAplikacja</a>
      </div>
  
      <!-- Prawa strona (duże ekrany) -->
      <div class="hidden lg:flex">
        <ul class="menu menu-horizontal px-1">
          <li><a>Home</a></li>
          <li><a>O nas</a></li>
          <li><a>Kontakt</a></li>
        </ul>
      </div>
  
      <!-- Hamburger + Dropdown (małe ekrany) -->
      <div class="dropdown dropdown-end lg:hidden">
        <div tabindex="0" role="button" class="btn btn-ghost btn-square">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
            viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round"
              stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </div>
        <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
          <li><a>Home</a></li>
          <li><a>O nas</a></li>
          <li><a>Kontakt</a></li>
        </ul>
      </div>
    </div>
  </template>
  
  <script setup>
  // Brak logiki — czysto struktura HTML
  </script>