<template>
  <div class="relative flex items-center">
    <!-- Avatar Icon -->
    <button
  @click="open = !open"
  class="flex items-center justify-center focus:outline-none hover:text-blue-400 transition"
  aria-label="User menu"
>
  <!-- Klasyczny user SVG bez tła -->
  <svg xmlns="http://www.w3.org/2000/svg" class="w-7 h-7" fill="currentColor" viewBox="0 0 24 24">
    <path d="M12 12c2.7 0 4.5-1.8 4.5-4.5S14.7 3 12 3 7.5 4.8 7.5 7.5 9.3 12 12 12zm0 1.5c-2.4 0-7.5 1.2-7.5 3.75V21h15v-3.75C19.5 14.7 14.4 13.5 12 13.5z"/>
  </svg>
</button>


    <!-- Dropdown menu -->
    <transition name="fade">
      <div
        v-if="open"
        class="absolute right-0 top-12 w-56 bg-[#111921] text-white rounded-xl shadow-lg py-2 z-50 border border-gray-700"
      >
        <router-link
          to="/login"
          class="block px-6 py-2 text-base hover:bg-[#1a2533] transition"
          @click="open = false"
        >Log in</router-link>
        <router-link
          to="/register"
          class="block px-6 py-2 text-base hover:bg-[#1a2533] transition"
          @click="open = false"
        >Register</router-link>
        <router-link
          to="/favorites"
          class="block px-6 py-2 text-base hover:bg-[#1a2533] transition"
          @click="open = false"
        >Favorite Recipes</router-link>
        <router-link
          to="/profile"
          class="block px-6 py-2 text-base hover:bg-[#1a2533] transition"
          @click="open = false"
        >Profile</router-link>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'

const open = ref(false)
function handleClick(event) {
  const menu = document.querySelector('.relative.flex.items-center')
  if (open.value && menu && !menu.contains(event.target)) {
    open.value = false
  }
}
onMounted(() => window.addEventListener('mousedown', handleClick))
onBeforeUnmount(() => window.removeEventListener('mousedown', handleClick))
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity .15s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>
