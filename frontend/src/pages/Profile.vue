import { addUserTags } from "../api/axios"

<template>
  <div class="bg-base-200 text-white min-h-screen overflow-hidden flex flex-col">
    <div class="h-[80px] shrink-0">
      <Navbar />
    </div>

    <div class="flex-grow flex h-[calc(1080px-300px)]">
      <!-- LeftBar -->
      <aside class="w-64 bg-base-100 border-r border-base-content p-4 flex flex-col items-start">
        <div class="w-full space-y-1 mt-2">
          <button
            v-for="item in sidebarItems"
            :key="item.name"
            @click="tab = item.tab"
            :class="[
              'w-full flex items-center gap-2 px-4 py-2 rounded-full text-left transition',
              tab === item.tab ? 'bg-primary text-white' : 'hover:bg-base-300'
            ]"
          >
            <component :is="item.icon" class="w-5 h-5" />
            <span class="font-medium">{{ item.name }}</span>
          </button>
        </div>
      </aside>

      <main class="flex-1 p-6 space-y-6 overflow-hidden">
        <!-- Main -->
        <div v-if="tab === 'profile'" class="flex flex-col items-center justify-center h-full">
          <div class="avatar mb-4">
            <div class="w-36 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2">
              <img src="https://i.pravatar.cc/300" alt="User avatar" />
            </div>
          </div>
          <h2 class="text-2xl font-bold mb-2">Username</h2>
          <button class="btn btn-primary btn-sm">Edit profile</button>
        </div>

        <div v-else-if="tab === 'privacy'" class="max-w-xl space-y-4">
          <h2 class="text-xl font-bold">Zmiana danych</h2>
          <p class="text-gray-400">Tutaj możesz zarządzać swoimi danymi kontaktowymi i wiekiem.</p>

          <!-- Sub-tab selector -->
          <div class="flex border-b border-gray-600 mb-4">
            <button
              v-for="opt in ['E-mail','Imię i nazwisko','Telefon','Wiek']"
              :key="opt"
              @click="privacySub = opt"
              class="flex-1 text-center pb-2 transition"
              :class="privacySub === opt
                ? 'border-b-2 border-primary text-white'
                : 'text-gray-400 hover:text-white'"
            >
              {{ opt }}
            </button>
          </div>

          <!-- Email form -->
          <div v-if="privacySub === 'E-mail'" class="space-y-4">
            <label class="form-control">
              <span class="label-text text-white">Nowy adres e-mail</span>
              <input
                type="email"
                v-model="form.email"
                class="input input-bordered w-full"
                placeholder="jan.kowalski@example.com"
              />
            </label>
            <label class="form-control">
              <span class="label-text text-white">Potwierdź adres e-mail</span>
              <input
                type="email"
                v-model="form.emailConfirm"
                class="input input-bordered w-full"
                placeholder="Powtórz nowy e-mail"
              />
            </label>
            <button class="btn btn-success" @click="save('email')" :disabled="form.email !== form.emailConfirm">
              Zapisz e-mail
            </button>
          </div>

          <!-- Name form -->
          <div v-else-if="privacySub === 'Imię i nazwisko'" class="space-y-4">
            <label class="form-control">
              <span class="label-text text-white">Imię i nazwisko</span>
              <input
                type="text"
                v-model="form.fullName"
                class="input input-bordered w-full"
                placeholder="Wprowadź imię i nazwisko"
              />
            </label>
            <button class="btn btn-success" @click="save('fullName')">Zapisz imię i nazwisko</button>
          </div>

          <!-- Phone form -->
          <div v-else-if="privacySub === 'Telefon'" class="space-y-4">
            <label class="form-control">
              <span class="label-text text-white">Numer telefonu</span>
              <input
                type="tel"
                v-model="form.phone"
                class="input input-bordered w-full"
                placeholder="+48 123 456 789"
              />
            </label>
            <button class="btn btn-success" @click="save('phone')">Zapisz numer</button>
          </div>
          
          <!-- Birthdate form -->
          <div v-else-if="privacySub === 'Wiek'" class="space-y-4">
            <label class="form-control">
              <span class="label-text text-white">Data urodzenia</span>
              <input
                type="date"
                v-model="form.birthDate"
                class="input input-bordered w-full"
              />
            </label>
            <button class="btn btn-success" @click="save('birthDate')">Zapisz wiek</button>
          </div>
        </div>


        <div v-else-if="tab === 'sharing'" class="space-y-2">
          <h2 class="text-xl font-bold">Osoby i udostępnianie</h2>
          <p class="text-gray-400">Zarządzaj użytkownikami, którzy mogą przeglądać Twoje dane lub przepisy.</p>
        </div>
        <div v-else-if="tab === 'payments'" class="max-w-2xl space-y-4">
          <h2 class="text-xl font-bold">Płatności i subskrypcje</h2>
          <table class="table w-full text-white">
            <thead>
              <tr>
                <th>Plan</th>
                <th>Cena</th>
                <th>Funkcje</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Free</td>
                <td>0 zł</td>
                <td>Ograniczony dostęp do przepisów</td>
                <td><button class="btn btn-sm btn-disabled">Obecny</button></td>
              </tr>
              <tr>
                <td>Pro</td>
                <td>19,99 zł/mies.</td>
                <td>Nieograniczone przepisy, zapisane dania</td>
                <td><button class="btn btn-sm btn-primary">Wybierz</button></td>
              </tr>
            </tbody>
          </table>
        </div>
<!-- TEST -->
        <div v-else-if="tab === 'Tagi'">
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
      </div>
      </main>
    </div>
  </div>
</template>



<script setup>
import Navbar from '@/components/Navbar.vue'
import RecipeRating from '@/components/RecipeRating.vue'
import { reactive, ref, computed, onMounted, watch } from 'vue'
import { User, Settings, Lock, Users, CreditCard } from 'lucide-vue-next'
import { useRoute } from 'vue-router'
import { getTags, addUserTags, deleteUserTags, displayUserTags } from '@/api/axios'

const sidebarItems = [
  { name: 'Strona główna', tab: 'profile', icon: User },
  { name: 'Zmiana danych', tab: 'privacy', icon: Settings },
  { name: 'Osoby i udostępnianie', tab: 'sharing', icon: Users },
  { name: 'Płatności i subskrypcje', tab: 'payments', icon: CreditCard },
  { name: 'Tagi', tab: 'Tagi', icon: Users }  // <-- testowa zakładka
]

const selectedTags = ref({})
const tab = ref('profile')
const privacySub = ref('E-mail')
const form = ref({
  email: '',
  emailConfirm: '',
  fullName: '',
  phone: '',
  birthDate: ''
})
const rating = ref(0)
const route = useRoute()
const categories = ref([])
const recipes = reactive([])
const filterQuery = ref('')
const dropdownVisible = ref(false)
const dropdownWrapper = ref(null)

const expandedCategory = ref(null)

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

function toggleCategory(name) {
  expandedCategory.value = expandedCategory.value === name ? null : name
}

const filteredRecipes = computed(() => recipes.value)

async function addTag(tag,category) {
  if (!selectedTags.value[category]) {
    selectedTags.value[category] = []
  }
  if (!selectedTags.value[category].includes(tag)) {
    selectedTags.value[category].push(tag)
    filterQuery.value = ''
    dropdownVisible.value = false
    const payload = {
      name: tag,
      type: category
    }
    const res = await addUserTags(payload)
  }
}

async function removeTag(category, tag) {
  const tags = selectedTags.value[category]
  if (!Array.isArray(tags)) {
    console.warn(`Tags for category ${category} is not an array`, tags)
    return
  }
  selectedTags.value[category] = tags.filter(t => t !== tag)
  if (selectedTags.value[category].length === 0) {
    delete selectedTags.value[category]
  }
  await deleteUserTags(tag)
}

// Ukrywanie dropdownu
function handleClickOutside(event) {
  setTimeout(() => {
    if (dropdownWrapper.value && !dropdownWrapper.value.contains(event.target)) {
      dropdownVisible.value = false
    }
  }, 0)
}

function normalizeTags(data) {
  const result = {}
  data.forEach(item => {
    if (!result[item.category]) {
      result[item.category] = []
    }
    if (!result[item.category].includes(item.value)) {
      result[item.category].push(item.value)
    }
  })
  return result
}


onMounted(async () => {
  try {
    const tagsRes = await getTags()
    categories.value = tagsRes.data

    const res = await displayUserTags()
    selectedTags.value = normalizeTags(res.data)
  } catch (err) {
    console.error('Błąd ładowania danych:', err)
  }

  document.addEventListener('click', handleClickOutside)
  const tabQuery = route.query.tab
  if (tabQuery && typeof tabQuery === 'string') {
    tab.value = tabQuery
  }
})

function save(field) {
  console.log('Zapisuję', field, '->', form.value[field])
}


</script>