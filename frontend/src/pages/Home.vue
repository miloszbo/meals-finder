<template>
  <div class="main-page">
    <Navbar />
    <!-- Tu będzie dalsza część strony -->
    <div class="content p-6 text-white">
      <!-- przykładowa zawartość -->
      <p>Welcome to MealsFinder!</p>
    </div>
  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'

export default {
  name: 'HomePage',
  components: {
    Navbar
  }
}
</script>

<style scoped>
.main-page {
  background-color: #6785; 
}
</style>
