<template>
  <div class="main-page">

    <!-- Tu możesz dodać treść strony głównej -->
    <h1>Witaj na stronie głównej!</h1>
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
  width: 100vw;
  height: 100vh;
  background: linear-gradient(to right, #4facfe, #00f2fe);
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  font-size: 2rem;
}
</style>