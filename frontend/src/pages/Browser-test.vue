<template>
  <div>
    <Navbar />
  </div>

  <div class="flex min-h-screen bg-base-200 text-white">
    <!-- Sidebar -->
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

      <!-- Wybrane tagi -->
      <div class="mt-4" v-if="selectedTags.length">
        <h3 class="text-md font-semibold mb-2">Selected Tags:</h3>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(tag, index) in selectedTags"
            :key="index"
            @click="removeTag(tag)"
            class="bg-primary text-white rounded-full px-4 py-1 text-sm font-semibold cursor-pointer hover:bg-red-500"
          >
            {{ tag }}
          </span>
        </div>
      </div>
    </aside>

    <!-- Recipe Grid -->
    <main class="flex-1 p-6 overflow-y-auto max-h-[calc(100vh-4rem)]">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <router-link
          v-for="recipe in filteredRecipes"
          :key="recipe.id"
          :to="`/recipe/${recipe.id}`"
          class="bg-base-100 p-4 rounded-box shadow hover:shadow-lg transition block"
        >
          <div class="text-sm font-bold mb-1">
            Trudność:
            <span class="text-yellow-400">
              {{ '★'.repeat(recipe.difficulty) }}{{ '☆'.repeat(5 - recipe.difficulty) }}
            </span>
          </div>
          <h3 class="text-lg font-semibold mb-1">{{ recipe.name }}</h3>
          <p class="text-sm text-gray-400 mb-1">Czas: {{ recipe.time }} min</p>
          <p class="text-sm text-gray-300 line-clamp-2">
            {{ recipe.description }}
          </p>
        </router-link>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import axios from 'axios'
import Navbar from '@/components/Navbar.vue'

const categories = ref([])
const recipes = ref([])
const selectedTags = ref([])
const filterQuery = ref('')
const dropdownVisible = ref(false)
const dropdownWrapper = ref(null)

const allTags = computed(() =>
  categories.value.flatMap(category => category.tags)
)

const filteredTags = computed(() =>
  allTags.value
    .filter(tag =>
      tag.toLowerCase().startsWith(filterQuery.value.toLowerCase()) &&
      !selectedTags.value.includes(tag)
    )
    .slice(0, 10)
)

const filteredRecipes = computed(() => {
  if (selectedTags.value.length === 0) return []
  return recipes.value.filter(recipe =>
    Array.isArray(recipe.tags) &&
    selectedTags.value.every(tag => recipe.tags.includes(tag))
  )
})

function addTag(tag) {
  if (!selectedTags.value.includes(tag)) {
    selectedTags.value.push(tag)
  }
  filterQuery.value = ''
  dropdownVisible.value = false
}

function removeTag(tag) {
  selectedTags.value = selectedTags.value.filter(t => t !== tag)
}

function handleClickOutside(event) {
  setTimeout(() => {
    if (dropdownWrapper.value && !dropdownWrapper.value.contains(event.target)) {
      dropdownVisible.value = false
    }
  }, 0)
}

onMounted(async () => {
  try {
    const [tagsRes, recipesRes] = await Promise.all([
      axios.get('/tags.json'),
      axios.get('/recipes.json')
    ])
    categories.value = tagsRes.data
    recipes.value = recipesRes.data
  } catch (err) {
    console.error('Błąd ładowania danych:', err)
  }

  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
/* Ograniczenie opisu do 2 linii */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
