<template>
  <!-- Section displaying the most popular recipes of the week -->
  <section class="max-w-screen-xl mx-auto px-8 py-10">
    <!-- Title of the section -->
    <h2 class="text-4xl font-bold mb-2 text-white">Top of the Week</h2>
    <!-- Section subtitle -->
    <p class="text-gray-400 mb-8">This week's favorite dishes picked by our community</p>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
      <!-- Render a card for each recipe -->
      <div
        v-for="recipe in recipes"
        :key="recipe.title"
        class="bg-[#060606] rounded-xl overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-200 flex flex-col"
      >
        <!-- Recipe image, title, description, and a button -->
        <img :src="recipe.image" :alt="recipe.title" class="w-full h-40 object-cover" />
        <div class="p-4 flex flex-col flex-1">
          <h3 class="text-lg font-semibold mb-1 text-white">{{ recipe.title }}</h3>
          <p class="text-sm text-gray-400 mb-4 flex-1">{{ recipe.description }}</p>
          <button
            class="w-full mt-4 py-2 text-base font-semibold rounded-full bg-[#605DFF] text-white hover:bg-[#4947c5] transition"
          >
            View Recipe
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
// Array of recipe objects
const recipes = [
  {
    title: 'Sesame Chicken',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: 'images/sesame-chicken.jpg',
  },
  {
    title: 'Italian Pancake',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: 'images/italian-pancake.jpg',
  },
  {
    title: 'Fruits Of the Forest Smoothie',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: 'images/smoothie.jpg',
  },
  {
    title: 'Seafood ramen',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: 'images/seafood-ramen.jpg',
  },
  {
    title: 'Beef Steak',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: 'images/beef-steak.jpg',
  },
]
</script>
