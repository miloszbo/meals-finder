<template>
  <div>
    <Navbar />
  </div>

  <div class="min-h-screen bg-base-200 text-base-content p-8">
    <div class="max-w-4xl mx-auto bg-base-100 p-8 rounded-2xl shadow-xl">
      <h2 class="text-3xl font-bold mb-6">Add a Recipe</h2>

      <!-- Upload obrazu -->
      <div class="form-control mb-4">
        <label class="label"><span class="label-text">Recipe Image</span></label>
        <input type="file" @change="handleImage" class="file-input file-input-bordered w-full" />
      </div>

      <!-- Nazwa przepisu -->
      <div class="form-control mb-4">
        <label class="label"><span class="label-text">Recipe Name</span></label>
        <input type="text" v-model="form.name" class="input input-bordered w-full" />
      </div>

      <!-- Dropdown tagów -->
      <div class="form-control mb-6">
        <label class="label"><span class="label-text">Tags</span></label>
        <div class="relative" ref="dropdownWrapper">
          <input
            type="text"
            v-model="filterQuery"
            placeholder="Search tag..."
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
        <!-- Wyświetlanie wybranych tagów -->
        <div class="mt-2 flex flex-wrap gap-2 w-full">
          <span
            v-for="(tags, category) in selectedTags"
            :key="category"
          >
            <span
              v-for="tag in tags"
              :key="tag"
              @click="removeTag(category, tag)"
              class="bg-primary text-white rounded-full px-4 py-1 text-sm font-semibold cursor-pointer hover:bg-red-500 mr-1"
            >
              {{ tag }}
            </span>
          </span>
        </div>
      </div>

      <!-- Składniki -->
      <div class="form-control mb-6">
        <label class="label"><span class="label-text">Ingredients</span></label>
        <div class="flex flex-wrap md:flex-nowrap gap-2 items-start relative" ref="ingredientDropdownWrapper">
          <div class="w-full relative">
            <input
              type="text"
              v-model="ingredientFilterQuery"
              placeholder="Search ingredient..."
              class="input input-sm input-bordered w-full"
              @focus="ingredientDropdownVisible = true"
            />
            <div
              v-if="ingredientDropdownVisible"
              class="absolute z-10 bg-base-100 border border-base-content mt-1 w-full rounded-box shadow max-h-60 overflow-y-auto"
            >
              <div
                v-for="item in filteredIngredients"
                :key="item"
                @click="selectIngredient(item)"
                class="px-3 py-2 hover:bg-base-300 cursor-pointer text-sm"
              >
                {{ item }}
              </div>
            </div>
          </div>

          <select v-model="ingredientInput.unit" class="select select-sm select-bordered">
            <option disabled value="">Unit</option>
            <option value="szt.">szt.</option>
            <option value="g">g</option>
            <option value="kg">kg</option>
            <option value="ml">ml</option>
            <option value="l">l</option>
          </select>

          <input
            v-model="ingredientInput.amount"
            type="number"
            placeholder="Amount"
            class="input input-sm input-bordered w-24"
          />

          <button @click="addIngredient" class="btn btn-sm btn-accent" type="button">Add</button>
        </div>

        <ul class="mt-4 space-y-2">
          <li
            v-for="(ingredient, index) in ingredients"
            :key="index"
            class="flex justify-between items-center bg-base-300 px-4 py-2 rounded"
          >
            <span>{{ ingredient.name }} – {{ ingredient.amount }} {{ ingredient.unit }}</span>
            <button @click="removeIngredient(index)" class="text-xl leading-none text-red-500 hover:text-red-700">&times;</button>
          </li>
        </ul>
      </div>

      <!-- Czas i trudność -->
      <div class="flex gap-4 mb-4">
        <div class="form-control w-1/2">
          <label class="label"><span class="label-text">Recipe Time</span></label>
          <input type="text" v-model="form.time" placeholder="e.g. 45 min" class="input input-bordered w-full" />
        </div>
        <div class="form-control w-1/2">
          <label class="label"><span class="label-text">Recipe Difficulty</span></label>
          <select v-model="form.difficulty" class="select select-bordered w-full">
            <option disabled value="">Select difficulty</option>
            <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
          </select>
        </div>
      </div>

      <!-- Opis przepisu -->
      <div class="form-control mb-4">
        <label class="label"><span class="label-text">Recipe Description</span></label>
        <textarea
          v-model="form.description"
          class="textarea textarea-bordered w-full resize-y min-h-[150px]"
          placeholder="Describe how to prepare the recipe..."
        ></textarea>
      </div>

      <button @click="submitRecipe" class="btn btn-primary w-full">Submit</button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import Navbar from '../components/Navbar.vue'
import { addRecipe } from '@/api/axios'

const form = ref({
  name: '',
  time: '',
  difficulty: '',
  description: '',
  tags: '',
  image: null,
})

const categories = ref([])
const selectedTags = ref({})
const filterQuery = ref('')
const dropdownVisible = ref(false)
const dropdownWrapper = ref(null)
const expandedCategory = ref(null)

const availableIngredients = ref([])
const ingredients = ref([])
const ingredientInput = ref({ name: '', unit: '', amount: '' })
const ingredientDropdownWrapper = ref(null)
const ingredientDropdownVisible = ref(false)
const ingredientFilterQuery = ref('')

const tagToCategory = computed(() => {
  const map = {}
  categories.value.forEach(cat => {
    cat.tags.forEach(tag => {
      map[tag] = cat.name
    })
  })
  return map
})

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

const filteredIngredients = computed(() =>
  availableIngredients.value
    .filter(item =>
      item.toLowerCase().includes(ingredientFilterQuery.value.toLowerCase()) &&
      !ingredients.value.some(i => i.name === item)
    )
    .slice(0, 10)
)

function toggleCategory(name) {
  expandedCategory.value = expandedCategory.value === name ? null : name
}

function addTag(tag, category) {
  if (!selectedTags.value[category]) {
    selectedTags.value[category] = []
  }
  if (!selectedTags.value[category].includes(tag)) {
    selectedTags.value[category].push(tag)
    dropdownVisible.value = false
    filterQuery.value = ''
  }
}

function removeTag(category, tag) {
  selectedTags.value[category] = selectedTags.value[category].filter(t => t !== tag)
  if (selectedTags.value[category].length === 0) {
    delete selectedTags.value[category]
  }
}

function selectIngredient(item) {
  ingredientInput.value.name = item
  ingredientFilterQuery.value = item
  ingredientDropdownVisible.value = false
}

function addIngredient() {
  const { name, unit, amount } = ingredientInput.value
  if (!name || !unit || !amount) {
    alert('Uzupełnij wszystkie pola składnika')
    return
  }
  ingredients.value.push({ ...ingredientInput.value })
  ingredientInput.value = { name: '', unit: '', amount: '' }
}

function removeIngredient(index) {
  ingredients.value.splice(index, 1)
}

function handleClickOutside(event) {
  setTimeout(() => {
    if (dropdownWrapper.value && !dropdownWrapper.value.contains(event.target)) {
      dropdownVisible.value = false
    }
  }, 0)
}

function handleIngredientClickOutside(event) {
  setTimeout(() => {
    if (ingredientDropdownWrapper.value && !ingredientDropdownWrapper.value.contains(event.target)) {
      ingredientDropdownVisible.value = false
    }
  }, 0)
}

function handleImage(e) {
  form.value.image = e.target.files[0]
}

onMounted(async () => {
  try {
    const tagsRes = await fetch('/tags.json')
    categories.value = await tagsRes.json()

    const ingredientsRes = await fetch('/ingredients.json')
    availableIngredients.value = await ingredientsRes.json()

    document.addEventListener('click', handleClickOutside)
    document.addEventListener('click', handleIngredientClickOutside)
  } catch (err) {
    console.error('Błąd ładowania danych:', err)
  }
})

const submitRecipe = async () => {
  const formData = new FormData()
  formData.append('name', form.value.name)
  formData.append('time', form.value.time)
  formData.append('difficulty', form.value.difficulty)
  formData.append('description', form.value.description)

  formData.append('tags', JSON.stringify(selectedTags.value))
  formData.append('ingredients', JSON.stringify(ingredients.value))

  if (form.value.image) {
    formData.append('image', form.value.image)
  }

  try {
    await addRecipe(formData)
    alert('Recipe submitted!')

    form.value = {
      name: '',
      time: '',
      difficulty: '',
      description: '',
      tags: '',
      image: null,
    }
    selectedTags.value = {}
    ingredients.value = []
  } catch (err) {
    console.error(err)
    alert('Error submitting recipe')
  }
}
</script>
