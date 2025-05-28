<template>
  <div class="min-h-screen flex flex-col bg-[#181e26]">
    <div class="flex-1 flex flex-col">
      <router-view />
    </div>
    <Footer />
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
