<template>

  <div>
    <Navbar />
  </div>
  <div class="min-h-screen bg-base-200 text-base-content p-8">
    <div class="max-w-4xl mx-auto bg-base-100 p-8 rounded-2xl shadow-xl">
      <h2 class="text-3xl font-bold mb-6">Add a Recipe</h2>

      <!-- Recipe Image Upload -->
      <div class="form-control mb-4">
        <label class="label">
          <span class="label-text">Recipe Image</span>
        </label>
        <input type="file" @change="handleImage" class="file-input file-input-bordered w-full" />
      </div>

      <!-- Recipe Name -->
      <div class="form-control mb-4">
        <label class="label">
          <span class="label-text">Recipe Name</span>
        </label>
        <input type="text" v-model="form.name" class="input input-bordered w-full" />
      </div>

      <!-- Time & Difficulty -->
      <div class="flex gap-4 mb-4">
        <div class="form-control w-1/2">
          <label class="label">
            <span class="label-text">Recipe Time</span>
          </label>
          <input type="text" v-model="form.time" placeholder="e.g. 45 min" class="input input-bordered w-full" />
        </div>

        <div class="form-control w-1/2">
          <label class="label">
            <span class="label-text">Recipe Difficulty</span>
          </label>
            <select v-model="form.difficulty" class="select select-bordered w-full">
            <option disabled value="">Select difficulty</option>
            <option v-for="n in 5" :key="n" :value="n"> {{ n }}</option>
            </select>
        </div>
      </div>

        <!-- Description -->
        <div class="form-control mb-4">
        <label class="label mb-2">
            <span class="label-text">Recipe Description</span>
        </label>
        <textarea
            v-model="form.description"
            class="textarea textarea-bordered w-full resize-y min-h-[150px]"
            placeholder="Describe how to prepare the recipe..."
        ></textarea>
        </div>

      <!-- Tags -->
      <div class="form-control mb-6">
        <label class="label">
          <span class="label-text">Tags</span>
        </label>
        <input type="text" v-model="form.tags" placeholder="e.g. vegan, italian" class="input input-bordered w-full" />
      </div>

      <!-- Submit Button -->
      <button @click="submitRecipe" class="btn btn-primary w-full">Submit</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Navbar from '../components/Navbar.vue'
import { addRecipe } from '@/api/axios' // dodaj import funkcji

const form = ref({
  name: '',
  time: '',
  difficulty: '',
  description: '',
  tags: '',
  image: null,
})

const handleImage = (e) => {
  form.value.image = e.target.files[0]
}

const submitRecipe = async () => {
  const formData = new FormData()
  formData.append('name', form.value.name)
  formData.append('time', form.value.time)
  formData.append('difficulty', form.value.difficulty)
  formData.append('description', form.value.description)
  formData.append('tags', form.value.tags)
  if (form.value.image) {
    formData.append('image', form.value.image)
  }

  try {
    await addRecipe(formData)
    alert('Recipe submitted!')
    
    form.value = {
      name: '',
      time: '',
      difficulty: '',
      description: '',
      tags: '',
      image: null,
    }
  } catch (err) {
    console.error(err)
    alert('Error submitting recipe')
  }
}
</script>