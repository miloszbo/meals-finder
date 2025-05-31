<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { verifyUser, logoutUser } from '@/api/axios'

const isLoggedIn = ref(false)
const dropdownOpen = ref(false)
const router = useRouter()

onMounted(async () => {
  try {
    await verifyUser()
    isLoggedIn.value = true
  } catch {
    isLoggedIn.value = false
  }
})

const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value
}

const logout = async () => {
  try {
    await logoutUser()
    isLoggedIn.value = false
    router.push('/login')
  } catch (err) {
    console.error('Błąd przy wylogowaniu:', err)
  }
}
</script>

<template>
  <header class="bg-[#211C54] text-white text-base shadow-md">
    <div class="flex items-center justify-between px-6 py-4 h-20">
      <router-link to="/" class="text-5xl font-bold hover:text-primary transition-colors duration-200">
        MealsFinder
      </router-link>

      <!-- 🔽 Ikonki po prawej -->
      <div class="flex items-center gap-4 text-lg relative">
        <!-- Recipes link -->
        <router-link to="/browser" title="Recipes">
          <i class="fa-solid fa-clipboard-list cursor-pointer hover:text-gray-300"></i>
        </router-link>

        <!-- User icon + dropdown -->
        <div class="relative">
          <i
            class="fa-solid fa-user cursor-pointer hover:text-gray-300"
            @click="toggleDropdown"
          ></i>
          <div
            v-if="dropdownOpen"
            class="absolute right-0 mt-2 w-40 bg-base-100 text-white shadow-lg rounded z-50"
          >
            <div v-if="isLoggedIn">
              <router-link to="/profile" class="block px-4 py-2 hover:bg-base-300">Profil</router-link>
              <button @click="logout" class="w-full text-left px-4 py-2 hover:bg-red-600">Wyloguj</button>
            </div>
            <div v-else>
              <router-link to="/login" class="block px-4 py-2 hover:bg-base-300">Zaloguj się</router-link>
              <router-link to="/register" class="block px-4 py-2 hover:bg-base-300">Zarejestruj się</router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <nav class="bg-[#060606] border-t border-[#1f2a34] flex flex-wrap justify-center gap-6 px-4 py-2">
      <a
        v-for="cat in categories"
        :key="cat"
        :href="`/categories/${cat.toLowerCase().replace(/ /g, '-')}`"
        class="text-white"
        style="font-family: 'Roboto', sans-serif; font-weight: 700; font-size: 24px;"
      >
        {{ cat }}
      </a>
    </nav>
  </header>
</template>