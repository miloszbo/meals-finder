<template>
  <div class="flex-1 flex flex-col">
    <Navbar />
    <TopOfTheWeek />
    <CategoryBrowser />
    <GotAFavoriteRecipe />
  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'
import TopOfTheWeek from '../components/TopOfTheWeek.vue'
import CategoryBrowser from '../components/CategoryBrowser.vue'
import GotAFavoriteRecipe from '../components/GotAFavoriteRecipe.vue'

export default {
  name: 'HomePage',
  components: {
    Navbar,
    TopOfTheWeek,
    CategoryBrowser,
    GotAFavoriteRecipe
  }
}
</script>
