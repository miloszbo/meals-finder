<template>
  <!-- GŁÓWNY KONTENER SEKCJI -->
  <section class="max-w-screen-xl mx-auto px-8 py-10">

    <!-- Nagłówek sekcji -->
    <h2 class="text-4xl font-bold mb-2 text-white">Browse by category</h2>

    <!-- GRID kategorii (kafelki, responsywnie) -->
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
      <!-- Jeden kafelek kategorii, powielany v-for -->
      <div
        v-for="category in categories"
        :key="category.name"
        class="bg-[#111921] rounded-xl overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-200 flex flex-col"
      >
        <!-- Obrazek kategorii -->
        <img
          :src="category.image"
          :alt="category.name"
          class="w-full h-40 object-cover"
        />
        <!-- Nazwa kategorii pod obrazkiem -->
        <div class="p-4 flex flex-col flex-1 items-center">
          <span class="text-lg font-semibold text-white">{{ category.name }}</span>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
// TABLICA KATEGORII 
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
