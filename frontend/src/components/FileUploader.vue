<template>
  <div class="flex flex-col gap-4">
    <input type="file" @change="onFileSelect" />
    <p v-if="fileName">📄 Wybrano: <strong>{{ fileName }}</strong></p>
    <button v-if="file" class="btn btn-primary w-fit" @click="uploadFile">Upload</button>
    <p v-if="status" :class="statusClass">{{ status }}</p>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import axios from 'axios'

const file = ref(null)
const fileName = ref('')
const status = ref('')

const statusClass = computed(() => {
  if (status.value.startsWith('✅')) return 'text-green-600'
  if (status.value.startsWith('❌')) return 'text-red-600'
  return ''
})

// 🔗 Wklej wygenerowany Signed URL
const signedUrl = 'https://storage.googleapis.com/project-storege/upload-1750171384462.jpg?X-Goog-Algorithm=GOOG4-RSA-SHA256&X-Goog-Credential=signed-url-service%40meals-finders.iam.gserviceaccount.com%2F20250617%2Fauto%2Fstorage%2Fgoog4_request&X-Goog-Date=20250617T144304Z&X-Goog-Expires=604800&X-Goog-SignedHeaders=content-type%3Bhost&X-Goog-Signature=5e7c9eac9760b07edf5c605e6ecba9b5a91f15dccff73d9e1890b5ec0e53d5ff05a8eb9217671203a9982018169cfb6d2240d9c25e881b3d9595b8484124bcfc030f17dbbe382e2c8955f2dba14b363080001e5b61170bfc0232a8b0e9a318e96c59b37a0ef8db1bee826ac20031d63109d1cd94240abf2a4ab980df836d387d5c417db8130741189928c82a38159d8a60e7c318f82b6b420f16c4b778d6d11e19df7a273cb47bf017bc03501792334961f5db5e1e7d4ead4def2c5c6d151295182cde7a6474de5efeb30d105c2ae94f1a8ae7b46ce826fd6b0b3a4ca748a1fd12de6da1c3f1f4c26a628bca47c0ac515826879c43624eaba1cc867372b828b7'

function onFileSelect(event) {
  file.value = event.target.files[0]
  fileName.value = file.value?.name || ''
}

async function uploadFile() {
  if (!file.value) return

  try {
    await axios.put(signedUrl, file.value, {
      headers: { 'Content-Type': 'application/octet-stream' },
    })
    status.value = '✅ Plik wysłany pomyślnie!'
  } catch (err) {
    console.error(err)
    status.value = '❌ Błąd uploadu: ' + err.message
  }
}
</script>
