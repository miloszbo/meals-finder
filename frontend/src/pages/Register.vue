<template>
  <div class="flex min-h-screen">
    
    <!-- Left panel  -->
    <div class="hidden lg:block w-1/2 h-screen">
      <img
        src="/images/login_dish.png"
        alt="Delicious food"
        class="object-cover w-full h-full"
      />
    </div>
    <!-- Right panel (form) -->
    <div class="flex flex-col justify-center items-center w-full lg:w-1/2 h-screen bg-gradient-to-r from-[#060606] to-white to-900%">
      <div class="w-full max-w-md p-8 flex flex-col items-center justify-center">
        <router-link to="/home" class="text-4xl font-bold mb-8 text-white hover:text-primary transition-colors duration-200">
          MealsFinder
        </router-link>
        <form class="w-full flex flex-col items-center" @submit.prevent="registerUser">
          <!-- Step 1 -->
          <template v-if="step === 1">
            <div class="form-control mb-4 w-full">
              <input type="text" placeholder="Username" class="input input-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500" v-model="form.username" />
            </div>
            <div class="form-control mb-4 w-full">
              <input type="email" placeholder="Email" class="input input-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500" v-model="form.email" />
            </div>
            <div class="form-control mb-4 w-full">
              <input type="password" placeholder="Password" class="input input-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500" v-model="form.passwd" />
            </div>
            <div class="form-control mb-4 w-full">
              <input type="password" placeholder="Confirm Password" class="input input-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500" v-model="form.confirmpasswd" />
            </div>
            <div class="form-control mb-4 w-full flex justify-center">
              <button type="button" class="btn btn-primary w-full" @click="step = 2">Next</button>
            </div>
          </template>
          <!-- Step 2 -->
          <template v-else>
            <div class="form-control mb-4 w-full">
              <input type="number" placeholder="Phone Number" class="input input-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500" v-model="form.phoneNumber" />
            </div>
            <div class="form-control mb-4 w-full">
              <input type="number" placeholder="Age" class="input input-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500" v-model.number="form.age" min="0"/>
            </div>
            <div class="form-control mb-4 w-full">
              <select v-model="form.sex" class="select select-bordered w-full bg-[#060606] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500">
                <option disabled value="">Select gender</option>
                <option value="Mężczyzna">Male</option>
                <option value="Kobieta">Female</option>
                <option value="Nie chcę podawać">Prefer not to say</option>
              </select>
            </div>
            <div class="form-control mb-4 w-full flex flex-col gap-2">
              <button type="submit" class="btn btn-success w-full">Sign up</button>
              <button type="button" class="btn btn-outline w-full" @click="step = 1">← Back</button>
            </div>
          </template>
          <div class="text-center text-sm text-white w-full mt-2">
            <span>Have an account?</span>
            <router-link to="/Login" class="ml-1 text-[#466aff] hover:underline">Sign in</router-link>
          </div>
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
        alert('Fill all required fields.')
        return
      }
      if (this.form.passwd !== this.form.confirmpasswd) {
        alert('Passwords do not match')
        return
      }
      if (this.form.age > 100){
        alert("Incorrect age")
        return
      }
      if (this.form.phoneNumber.toString().length != 9){
        alert("Incorrect phone number")
        return
      }
      if (!this.form.email.includes("@")){
        alert("Incorrect email")
        return
      }
      try {
        const payload = {
          username: this.form.username,
          email: this.form.email,
          passwd: this.form.passwd,
          sex: this.form.sex,   
          age: this.form.age,
          phone_number: this.form.phoneNumber.toString()
        }
        const res = await api.post('/user/register', payload)
        console.log('Registered:', res.data)
        this.$track('user_registered', { username: this.form.username }) //analytical data
        this.$router.push('/login')
      } catch (err) {
        console.error('Registration error:', err)
        alert('Registration failed. Check your data.')
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
