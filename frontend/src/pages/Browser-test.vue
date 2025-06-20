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
          v-if="dropdownVisible"
          class="absolute z-10 bg-base-100 border border-base-content mt-1 w-full rounded-box shadow max-h-96 overflow-y-auto"
        >
          <div
            v-for="category in filteredCategories"
            :key="category.name"
            class="border-b px-4 py-2"
          >
            <div
              @click="toggleCategory(category.name)"
              class="cursor-pointer font-semibold hover:text-primary flex justify-between items-center"
            >
              {{ category.name }}
              <span class="text-xs">{{ expandedCategory === category.name ? '▲' : '▼' }}</span>
            </div>

            <div
              v-if="expandedCategory === category.name"
              class="mt-2 max-h-40 overflow-y-auto pl-2 space-y-1"
            >
              <div
                v-for="tag in limitedTags(category.tags)"
                :key="tag"
                @click.stop="addTag(tag, category.name)"
                class="px-3 py-1 hover:bg-base-300 rounded cursor-pointer text-sm"
              >
                {{ tag }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-4 w-full break-words" v-if="Object.keys(selectedTags).length">
        <h3 class="text-md font-semibold mb-2">Selected Tags:</h3>
        <div class="flex flex-wrap gap-2 w-full">
          <span
            v-for="(tags, category) in selectedTags"
            :key="category"
          >
            <span
              v-for="tag in tags"
              :key="tag"
              @click="removeTag(category, tag)"
              class="bg-primary text-white rounded-full px-4 py-1 text-sm font-semibold cursor-pointer hover:bg-red-500 truncate max-w-full inline-block"
            >
              {{ tag }}
            </span>
          </span>
        </div>
      </div>
    </aside>

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
import { reactive, ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Navbar from '@/components/Navbar.vue'
import { getAllRecipes, getTags } from '@/api/axios'

const route = useRoute()
const router = useRouter()

const categories = ref([])
const recipes = reactive([])
const filterQuery = ref('')
const dropdownVisible = ref(false)
const dropdownWrapper = ref(null)
const selectedTags = ref({})

// Mapowanie tagów do kategorii
const tagToCategory = computed(() => {
  const map = {}
  categories.value.forEach(cat => {
    cat.tags.forEach(tag => {
      map[tag] = cat.name
    })
  })
  return map
})

// Filtrowanie kategorii pasujących do inputa
const filteredCategories = computed(() =>
  categories.value.filter(cat =>
    cat.tags.some(tag =>
      tag.toLowerCase().startsWith(filterQuery.value.toLowerCase()) &&
      !(selectedTags.value[cat.name] || []).includes(tag)
    )
  )
)

function limitedTags(tags) {
  return tags
    .filter(tag =>
      tag.toLowerCase().startsWith(filterQuery.value.toLowerCase()) &&
      !Object.values(selectedTags.value).flat().includes(tag)
    )
    .slice(0, 10)
}

const expandedCategory = ref(null)
function toggleCategory(name) {
  expandedCategory.value = expandedCategory.value === name ? null : name
}

const filteredRecipes = computed(() => recipes.value)

// Aktualizacja URL
function updateUrlWithTags() {
  const query = {}
  for (const [cat, tags] of Object.entries(selectedTags.value)) {
    query[cat] = tags
  }
  router.replace({ query })
}

// Pobranie przepisów z uwzględnieniem wielu tagów per kategoria
async function fetchRecipes() {
  try {
    const parsedTags = {}
    for (const [cat, val] of Object.entries(selectedTags.value)) {
      parsedTags[cat] = val
    }
    const { data } = await getAllRecipes(parsedTags)
    recipes.value = data
  } catch (err) {
    console.error('Błąd pobierania przepisów:', err)
  }
}

// Dodanie tagu
function addTag(tag, category) {
  if (!selectedTags.value[category]) {
    selectedTags.value[category] = []
  }
  if (!selectedTags.value[category].includes(tag)) {
    selectedTags.value[category].push(tag)
    updateUrlWithTags()
    filterQuery.value = ''
    dropdownVisible.value = false
    fetchRecipes()
  }
}

// Usunięcie tagu
function removeTag(category, tag) {
  selectedTags.value[category] = selectedTags.value[category].filter(t => t !== tag)
  if (selectedTags.value[category].length === 0) {
    delete selectedTags.value[category]
  }
  updateUrlWithTags()
  fetchRecipes()
}

// Ukrywanie dropdownu
function handleClickOutside(event) {
  setTimeout(() => {
    if (dropdownWrapper.value && !dropdownWrapper.value.contains(event.target)) {
      dropdownVisible.value = false
    }
  }, 0)
}

onMounted(async () => {
  try {
    const localTags = await getTags()
    categories.value = await localTags.data

    const parsed = {}
    for (const [cat, val] of Object.entries(route.query)) {
      parsed[cat] = Array.isArray(val) ? val : val.split(',')
    }
    selectedTags.value = parsed
    await fetchRecipes()
  } catch (err) {
    console.error('Błąd ładowania danych:', err)
  }

  document.addEventListener('click', handleClickOutside)
})
</script>