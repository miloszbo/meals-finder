<template>
  <!-- Navbar -->
  <div>
    <Navbar />
  </div>

  <!-- Layout -->
  <div class="flex min-h-screen bg-base-200 text-white">
    <!-- Sidebar: Tag Selector -->
    <aside class="w-64 bg-base-100 p-4 border-r border-base-content overflow-y-auto">
      <h2 class="text-xl font-bold mb-4">Search by tags</h2>

      <!-- Input + Dropdown -->
      <div class="relative" ref="dropdownWrapper">
        <input
          type="text"
          v-model="filterQuery"
          placeholder="Type to search..."
          class="input input-sm input-bordered w-full mb-2"
          @focus="dropdownVisible = true"
        />

        <!-- Dropdown -->
        <div
          v-if="dropdownVisible && filteredTags.length > 0"
          class="absolute z-10 bg-base-100 border border-base-content mt-1 w-full max-h-48 overflow-y-auto rounded-box shadow"
        >
          <div
            v-for="tag in filteredTags"
            :key="tag"
            @click="addTag(tag)"
            class="px-4 py-2 hover:bg-base-300 cursor-pointer text-sm"
          >
            {{ tag }}
          </div>
        </div>
      </div>

      <!-- Wybrane tagi (z możliwością usunięcia) -->
      <div class="mt-4" v-if="selectedTags.length > 0">
        <h3 class="text-md font-semibold mb-2">Selected Tags:</h3>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(tag, index) in selectedTags"
            :key="index"
            @click="removeTag(tag)"
            class="bg-primary text-white rounded-full px-4 py-1 text-sm font-semibold cursor-pointer hover:bg-red-500 transition"
            title="Kliknij, aby usunąć"
          >
            {{ tag }}
          </span>
        </div>
      </div>
    </aside>

    <!-- Recipe grid -->
    <main class="flex-1 p-6 overflow-y-auto max-h-[calc(100vh-4rem)]">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <div
          v-for="(recipe, index) in recipes"
          :key="index"
          class="bg-base-100 p-4 rounded-box shadow hover:shadow-lg transition"
        >
          <h3 class="text-lg font-semibold mb-1">{{ recipe.title }}</h3>
          <p class="text-sm text-gray-400 mb-1">Cook: {{ recipe.time }}</p>
          <p class="text-sm">{{ recipe.description }}</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch, onBeforeUnmount } from 'vue'
import axios from 'axios'
import Navbar from '@/components/Navbar.vue'

// Dane z JSON-a
const categories = ref([])
const filterQuery = ref('')
const selectedTags = ref([])
const dropdownVisible = ref(false)

// Ref do obszaru dropdown + input
const dropdownWrapper = ref(null)

const recipes = ref([
  {
    title: 'Spicy Thai Curry',
    time: '30 min',
    description: 'A flavorful and spicy dish with coconut milk and veggies.'
  },
  {
    title: 'Quick Pasta',
    time: '15 min',
    description: 'Simple and fast pasta with garlic and cheese.'
  },
  {
    title: 'Vegan Buddha Bowl',
    time: '25 min',
    description: 'Healthy bowl with quinoa, chickpeas, and greens.'
  }
])

// Spłaszczone tagi z kategorii
const allTags = computed(() =>
  categories.value.flatMap(category => category.tags)
)

// Tylko tagi zaczynające się od wpisanych liter
const filteredTags = computed(() =>
  allTags.value
    .filter(tag =>
      tag.toLowerCase().startsWith(filterQuery.value.toLowerCase()) &&
      !selectedTags.value.includes(tag)
    )
    .slice(0, 10)
)

function addTag(tag) {
  selectedTags.value.push(tag)
  filterQuery.value = ''
  dropdownVisible.value = false
}

function removeTag(tag) {
  selectedTags.value = selectedTags.value.filter(t => t !== tag)
}

watch(filterQuery, (val) => {
  dropdownVisible.value = val !== ''
})

// Zamykanie dropdownu po kliknięciu poza nim
function handleClickOutside(event) {
  if (dropdownWrapper.value && !dropdownWrapper.value.contains(event.target)) {
    dropdownVisible.value = false
  }
}

onMounted(async () => {
  try {
    const res = await axios.get('/tags.json')
    categories.value = res.data
  } catch (err) {
    console.error('Błąd ładowania tagów:', err)
  }

  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>