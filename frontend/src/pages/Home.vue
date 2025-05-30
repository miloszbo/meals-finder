<template>
  <div class="min-h-screen flex flex-col bg-[#181e26]">
    <Navbar />
    <TopOfTheWeek />
    <CategoryBrowser />
    <GotAFavoriteRecipe />
    <Footer />
  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'
import TopOfTheWeek from '../components/TopOfTheWeek.vue'
import CategoryBrowser from '../components/CategoryBrowser.vue'
import GotAFavoriteRecipe from '../components/GotAFavoriteRecipe.vue'
import Footer from '../components/Footer.vue'

export default {
  name: 'HomePage',
  components: {
    Navbar,
    TopOfTheWeek,
    CategoryBrowser,
    GotAFavoriteRecipe,
    Footer
  }
}
</script>
