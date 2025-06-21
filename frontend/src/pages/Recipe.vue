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
    <div class="p-4">
  <div class="flex flex-col md:flex-row md:items-start md:gap-12">
    <!-- Rating input -->
    <div class="flex-1">
      <p class="text-gray-400 mb-2">Oceń przepis poniżej:</p>
      <RecipeRating v-model="rating" />
      <p class="mt-2">Twoja ocena: <span class="font-semibold">{{ rating }}</span>/5</p>
    </div>

    <!-- Review breakdown -->
    <div class="flex-1 mt-6 md:mt-0">
      <h3 class="text-lg font-semibold mb-2">Oceny użytkowników:</h3>
      <ul class="space-y-1 text-sm text-gray-300">
        <li v-for="n in 5" :key="n" class="flex justify-between">
          <span>{{ n }}★</span>
          <span>{{ reviewStats[n] || 0 }} osób</span>
        </li>
      </ul>
    </div>
  </div>
</div>
  </div>

  <div v-else class="text-center text-white py-20">
    <h1 class="text-2xl font-semibold">Nie znaleziono przepisu</h1>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getRecipeById, submitRecipeRating, getRatings, getReview } from '@/api/axios'
import RecipeRating from '@/components/RecipeRating.vue'
import { track } from '@/services/analytics' //analytics gathering lib

const rating = ref(0)
const route = useRoute()
const recipe = ref(null)
const reviewStats = ref({})

onMounted(async () => {
  try {
    const { data } = await getRecipeById(route.params.id)
    const ratings = await getRatings(route.params.id)
    console.log(rating)
    const ratingData = await getReview(route.params.id)
    console.log(ratingData.data.review_score)
    rating.value = ratingData.data.review_score
    recipe.value = data
    track('recipe_opened', { recipe_id: route.params.id }) //analytics gathering event
    reviewStats.value = ratings.data
  } catch (err) {
    console.error('Błąd ładowania:', err)
  }
})

watch(rating, async (newRating) => {
  if (newRating > 0 && recipe.value?.id) {
    try {
      await submitRecipeRating({ stars: newRating, recipe_id: recipe.value.id })

      // Refresh review stats after submitting
      const { data } = await getRatings(route.params.id)
      reviewStats.value = data
    } catch (err) {
      console.error('Błąd podczas wysyłania oceny:', err)
    }
  }
})

</script>