<template>
  
  <div class="flex h-screen">

    <!-- Left panel -->
    <div class="w-1/2 hidden lg:block">
      <img
        src="//assets/login_register dish.png"
        alt="Delicious food"
        class="object-cover w-full h-full"
      />
    </div>

    <!-- Right panel -->
    <div class="w-full lg:w-1/2 flex items-center justify-center bg-base-200">
      <div class="w-full max-w-md p-8 flex flex-col items-center justify-center">

        
        <router-link to="/home" class="text-4xl font-bold mb-8 hover:text-primary transition-colors duration-200">
          MealsFinder
        </router-link>

        <form class="w-full flex flex-col items-center">
          <!-- Etap 1 -->
          <template v-if="step === 1">
            <div class="form-control mb-4 w-full">
              <input type="text" placeholder="Username" class="input input-bordered w-full" v-model="form.username" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="email" placeholder="Email" class="input input-bordered w-full" v-model="form.email" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="password" placeholder="Password" class="input input-bordered w-full" v-model="form.passwd" /> <!-- Passwd -->
            </div>

            <div class="form-control mb-4 w-full">
              <input type="password" placeholder="Confirm Password" class="input input-bordered w-full" v-model="form.confirmpasswd" />
            </div>

            <div class="form-control mb-4 w-full flex justify-center">
              <button type="button" class="btn btn-primary w-full" @click="step = 2">Next</button>
            </div>
          </template>

          <!-- Etap 2 -->
          <template v-else>
            <div class="form-control mb-4 w-full">
              <input type="number" placeholder="Phone Number" class="input input-bordered w-full" v-model="form.phoneNumber" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="number" placeholder="Age" class="input input-bordered w-full" v-model.number="form.age" min="0"/>
            </div>

            <div class="form-control mb-4 w-full">
              <select v-model="form.sex" class="select select-bordered w-full">
                <option disabled value="">Wybierz płeć</option>
                <option value="Mężczyzna">Mężczyzna</option>
                <option value="Kobieta">Kobieta</option>
                <option value="Nie chcę podawać">Nie chcę podawać</option>
              </select>
            </div>
            
            <div class="form-control mb-4 w-full">
              <button type="submit" class="btn btn-success w-full" @click.prevent="registerUser">Sign up</button>
              <button type="button" class="btn btn-outline w-full" @click="step = 1">← Back</button>

            </div>
          </template>

          

          <div class="text-center text-sm">
            <span>Have an account?</span>
            <router-link to="/Login" class="link link-primary ml-1">Sign in</router-link>
          </div>
          <!--
          <div class="text-center text-sm mt-2">
            <a class="link link-hover link-secondary font-semibold">Forgot password?</a>
          </div>
          -->
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/api/axios'

export default {
  name: 'RegisterPage',
  data() {
    return {
      step: 1,
      form: {
        username: '',
        email: '',
        passwd: '',
        confirmpasswd: '', 
        sex: '',
        age: null,
        phoneNumber: ''
      }
    }
  },
  methods: {
    async registerUser() {

     if (!this.form.username || !this.form.email || !this.form.passwd || !this.form.confirmpasswd || !this.form.age === "" || !this.form.sex || !this.form.phoneNumber){
     alert('Uzupełnij wszystkie wymagane pola.')
     return
     }

    if (this.form.passwd !== this.form.confirmpasswd) {
    alert('Hasła się nie zgadzają')
    return
     }

     if (this.form.age > 100){
      alert("Błędny wiek")
      return
     }

     if (this.form.phoneNumber.toString().length != 9){
      alert("Błedny numer telefonu")
      return
     }

     if (!this.form.email.includes("@")){
      alert("Błędny email")
      return
     }


      try {
        const payload = {
          username: this.form.username,
          email: this.form.email,
          passwd: this.form.passwd,
          sex: this.form.sex,   
          age: Number(this.form.age),
          phone_number: this.form.phoneNumber
        }

        const res = await api.post('/user/register', payload)
        console.log('Zarejestrowano:', res.data)
        alert.apply("Rejestracja zakończona sukcesem")
        this.$router.push('/login')
      } catch (err) {
        console.error('Błąd rejestracji:', err)
        alert('Rejestracja nie powiodła się. Sprawdź dane.')
      }
    }
  }
}
</script>

<style scoped>
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

</style>