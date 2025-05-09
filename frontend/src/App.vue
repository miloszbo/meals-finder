<template>
  <div>
    <NavPage v-if="currentPage === 'nav'" @go-to="setPage" />
    <HomePage v-if="currentPage === 'home'" />
    <LoginPage v-if="currentPage === 'login'" />
    <RegisterPage v-if="currentPage === 'register'" />
  </div>
</template>

<script>
import HomePage from "./pages/Home.vue"
import LoginPage from "./pages/Login.vue"
import RegisterPage from "./pages/Register.vue"
import NavPage from "./pages/Nav.vue"


export default {
  name: 'App',
  components: {
    HomePage,
    LoginPage,
    RegisterPage,
    NavPage
  },

  data() {
    return {
      currentPage: 'nav'
    }
  },
  methods: {
    setPage(page) {
      this.currentPage = page
    }
  }
}

</script>




<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
