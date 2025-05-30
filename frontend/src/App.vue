<template>
  <div class="min-h-screen flex flex-col bg-[#060606]">
    <div class="flex-1 flex flex-col">
      <router-view />
    </div>
    <!-- 
      <Footer />
      Footer is currently only used on the main page (Home.vue).
      Uncomment this if you want the footer to appear on every page.
    -->
  </div>
</template>

<script>
import Footer from './components/Footer.vue'
export default {
  name: 'App',
  components: {
    Footer
  }
}
</script>
