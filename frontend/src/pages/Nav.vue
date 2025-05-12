<template>
  <div class="nav-page">
    <h1>Strona nawigacyjna</h1>
    <div class="buttons">
      <button @click="goTo('Home')">Home</button>
      <button @click="goTo('Login')">Login</button>
      <button @click="goTo('Register')">Register</button>
    </div>
  </div>
</template>
  
  <script>
  export default {
    name: 'NavPage',
    methods: {
      goTo(routeName) {
        this.$router.push({ name: routeName })
      }
    }
  }
  </script>
    
  <style scoped>
  .nav-page {
    width: 100vw;
    height: 100vh;
    background: #1e1e2f;
    color: white;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  
  .buttons {
    margin-top: 2rem;
    display: flex;
    gap: 1rem;
  }
  
  button {
    padding: 10px 20px;
    font-size: 1rem;
  }
  </style>