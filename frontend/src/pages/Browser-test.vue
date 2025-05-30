<template>
  <div>
    <Navbar />
  </div>

  <div class="flex min-h-screen bg-base-200 text-white">
    <aside class="w-64 bg-base-100 p-4 border-r border-base-content overflow-y-auto">
      <h2 class="text-xl font-bold mb-4">Search by tags</h2>

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

      <div class="mt-4" v-if="Object.keys(selectedTags).length">
        <h3 class="text-md font-semibold mb-2">Selected Tags:</h3>
        <div class="flex flex-wrap gap-2">
          <span
            v-for="(value, category) in selectedTags"
            :key="category"
            @click="removeTag(value)"
            class="bg-primary text-white rounded-full px-4 py-1 text-sm font-semibold cursor-pointer hover:bg-red-500"
          >
            {{ value }}
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
          :to="`/re/${recipe.id}`"
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
import { reactive, ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'
import Navbar from '@/components/Navbar.vue'
import { getAllRecipes } from '@/api/axios'

const route = useRoute()
const router = useRouter()

const categories = ref([])
const recipes = reactive([])
const filterQuery = ref('')
const dropdownVisible = ref(false)
const dropdownWrapper = ref(null)

const selectedTags = ref({})

const tagToCategory = computed(() => {
  const map = {}
  categories.value.forEach(cat => {
    cat.tags.forEach(tag => {
      map[tag] = cat.name.toLowerCase().replace(/\s/g, '_')
    })
  })
  return map
})

const allTags = computed(() =>
  categories.value.flatMap(cat => cat.tags)
)

const filteredTags = computed(() =>
  allTags.value.filter(tag =>
    tag.toLowerCase().startsWith(filterQuery.value.toLowerCase()) &&
    !Object.values(selectedTags.value).includes(tag)
  )
)

const filteredRecipes = computed(() => recipes.value)

function updateUrlWithTags() {
  router.replace({ query: { ...selectedTags.value } })
}

async function fetchRecipes() {
  try {
    console.log(route.query)
    const { data } = await getAllRecipes(selectedTags.value)
    recipes.value = data
  } catch (err) {
    console.error('Błąd pobierania przepisów:', err)
  }
}

function addTag(tag) {
  const category = tagToCategory.value[tag]
  if (!category) return

  selectedTags.value = {
    ...selectedTags.value,
    [category]: tag
  }

  updateUrlWithTags()
  filterQuery.value = ''
  dropdownVisible.value = false
  fetchRecipes()
}

function removeTag(tag) {
  const entry = Object.entries(selectedTags.value).find(([_, v]) => v === tag)
  if (entry) {
    const [category] = entry
    delete selectedTags.value[category]
    selectedTags.value = { ...selectedTags.value }
    updateUrlWithTags()
    fetchRecipes()
  }
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
    const tagsRes = await axios.get('/tags.json')
    categories.value = tagsRes.data

    selectedTags.value = { ...route.query }
    await fetchRecipes()
  } catch (err) {
    console.error('Błąd ładowania danych:', err)
  }

  document.addEventListener('click', handleClickOutside)
})
</script>