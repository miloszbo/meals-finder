<template>
  <div v-if="recipe" class="max-w-4xl mx-auto py-8 px-4 space-y-8 text-white">
    <div>
      <img :src="recipe.image" alt="recipe image" class="w-full h-64 object-cover rounded-xl mb-4" />
      <h1 class="text-3xl font-bold">{{ recipe.name }}</h1>
    </div>

    <div class="flex gap-6 text-gray-400 text-sm border-b pb-4">
      <span>Czas: <strong>{{ recipe.time }} min</strong></span>
      <span>Trudność:
        <strong class="text-yellow-400">
          {{ '★'.repeat(recipe.difficulty) }}{{ '☆'.repeat(5 - recipe.difficulty) }}
        </strong>
      </span>
    </div>

    <div>
      <h2 class="text-xl font-semibold mb-2">Składniki</h2>
      <ul class="list-disc list-inside">
        <li v-for="(ingredient, index) in recipe.ingredients.ingredients" :key="index">{{ ingredient.name }} {{ ingredient.amount }} {{ ingredient.unit }}</li>
      </ul>
    </div>

    <div>
      <h2 class="text-xl font-semibold mb-2">Sposób przygotowania</h2>
      <p class="text-sm text-gray-300">
            {{ recipe.recipe }}
      </p>
    </div>
  </div>

  <div v-else class="text-center text-white py-20">
    <h1 class="text-2xl font-semibold">Nie znaleziono przepisu</h1>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getRecipeById } from '@/api/axios'

const route = useRoute()
const recipe = ref(null)

onMounted(async () => {
  try {
    const { data } = await getRecipeById(route.params.id)
    console.log(data)
    recipe.value = data
  } catch (err) {
    console.error('Błąd wczytywania przepisu:', err)
  }
})
</script>