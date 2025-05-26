<template>
  <div class="max-w-4xl mx-auto py-8 px-4 space-y-8 text-white">
    <div>
      <img :src="recipe.image" alt="recipe image" class="w-full h-64 object-cover rounded-xl mb-4" />
      <h1 class="text-3xl font-bold">{{ recipe.name }}</h1>
    </div>

    <div class="flex gap-6 text-gray-400 text-sm border-b pb-4">
      <span>Czas: <strong>{{ recipe.time }} min</strong></span>
      <span>Trudność: <strong>{{ recipe.difficulty }}</strong></span>
    </div>

    <div>
      <h2 class="text-xl font-semibold mb-2">Opis</h2>
      <p class="text-gray-300">{{ recipe.description }}</p>
    </div>

    <div>
      <h2 class="text-xl font-semibold mb-2">Składniki</h2>
      <ul class="list-disc list-inside">
        <li v-for="(item, index) in recipe.ingredients" :key="index">{{ item }}</li>
      </ul>
    </div>

    <div>
      <h2 class="text-xl font-semibold mb-2">Sposób przygotowania</h2>
      <ol class="list-decimal list-inside space-y-1">
        <li v-for="(step, index) in recipe.instructions" :key="index">{{ step }}</li>
      </ol>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const recipe = ref(null)

onMounted(async () => {
  const { data } = await axios.get('/recipes.json')
  recipe.value = data.find(r => r.id === route.params.id)
})
</script>