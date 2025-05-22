<template>
  <div class="bg-black text-white p-6">
    <h2 class="text-2xl font-semibold mb-4 text-center">Browse by category</h2>
    <div class="flex gap-4 justify-center flex-wrap">
      <RouterLink
        v-for="category in categories"
        :key="category.slug"
        :to="`/categories/${category.slug}`"
        class="flex-shrink-0 w-40 hover:scale-105 transition-transform duration-200"
      >
        <div class="rounded-md overflow-hidden mb-2">
          <img
            :src="category.image"
            :alt="category.name"
            class="object-cover w-full h-[100px] rounded-md"
          />
        </div>
        <p class="text-center text-sm">{{ category.name }}</p>
      </RouterLink>
    </div>
  </div>
</template>

<script setup>
const categories = [
  {
    name: "Breakfast",
    slug: "breakfast",
    image: "/images/breakfast.jpg",
  },
  {
    name: "Dinner",
    slug: "dinner",
    image: "/images/dinner.jpg",
  },
  {
    name: "Desserts",
    slug: "desserts",
    image: "/images/desserts.jpg",
  },
  {
    name: "Vegan",
    slug: "vegan",
    image: "/images/vegan.jpg",
  },
  {
    name: "Quick Meals",
    slug: "quick-meals",
    image: "/images/quick-meals.jpg",
  },
]
</script>
