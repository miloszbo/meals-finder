<template>
  <!-- MAIN SECTION CONTAINER -->
  <section class="max-w-screen-xl mx-auto px-8 py-10">

    <!-- SECTION HEADER -->
    <h2 class="text-4xl font-bold mb-2 text-white">Browse by category</h2>

    <!-- Category GRID (responsive tiles) -->
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
      <!-- Single category card, rendered with v-for -->
      <div
        v-for="category in categories"
        :key="category.name"
        class="bg-[#060606] rounded-xl overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-200 flex flex-col"
      >
        <!-- Category image -->
        <img
          :src="category.image"
          :alt="category.name"
          class="w-full h-40 object-cover"
        />
        <!-- Category name below the image -->
        <div class="p-1 flex flex-col flex-1 items-center">
          <span class="text-lg font-semibold text-white">{{ category.name }}</span>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
// CATEGORIES ARRAY
const categories = [
  {
    name: "Breakfast",
    slug: "breakfast",
    image: "images/breakfast.jpg",
  },
  {
    name: "Dinner",
    slug: "dinner",
    image: "images/dinner.jpg",
  },
  {
    name: "Desserts",
    slug: "desserts",
    image: "images/desserts.jpg",
  },
  {
    name: "Vegan",
    slug: "vegan",
    image: "images/vegan.jpg",
  },
  {
    name: "Quick Meals",
    slug: "quick-meals",
    image: "images/quick-meals.jpg",
  },
]
</script>
