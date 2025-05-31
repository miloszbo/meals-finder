<template>
  <!-- Main section for Top of the Week recipes -->
  <section class="bg-[#0A1118] text-white px-8 py-10">
    <div class="max-w-screen-xl mx-auto">
      <!-- Section title and description -->
      <h2 class="text-4xl font-bold mb-2">Top of the Week</h2>
      <p class="text-gray-400 mb-8">This week's favorite dishes picked by our community</p>
      <!-- Responsive grid for recipe cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
        <div
          v-for="recipe in recipes"
          :key="recipe.title"
          class="bg-[#111921] rounded-xl overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-200"
        >
          <!-- Recipe image -->
          <img :src="recipe.image" :alt="recipe.title" class="w-full h-40 object-cover" />
          <div class="p-4">
            <!-- Recipe title -->
            <h3 class="text-lg font-semibold mb-1">{{ recipe.title }}</h3>
            <p class="text-sm text-gray-400 mb-4">{{ recipe.description }}</p>
            <!-- Button to view the recipe -->
            <a
              :href="recipe.link"
              class="inline-block px-4 py-2 bg-blue-600 text-white text-sm rounded hover:bg-blue-700 transition"
            >
              View Recipe
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
// Array of top recipes to display in the section
const recipes = [
  {
    title: 'Sesame Chicken',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: '/images/sesame-chicken.jpg',
    link: '#',
  },
  {
    title: 'Italian Pancake',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: '/images/italian-pancake.jpg',
    link: '#',
  },
  {
    title: 'Fruits Of the Forest Smoothie',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: '/images/smoothie.jpg',
    link: '#',
  },
  {
    title: 'Seafood ramen',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: '/images/seafood-ramen.jpg',
    link: '#',
  },
  {
    title: 'Beef Steak',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
    image: '/images/beef-steak.jpg',
    link: '#',
  },
]
</script>
