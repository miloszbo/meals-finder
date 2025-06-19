import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080",
    headers: {
        'Content-Type':'application/json'
    },
    withCredentials:true,
})

export const getAllRecipes = (params) =>
  api.get('/browser', { params })

export const getRecipeById = (id) =>
  api.get(`/re/${id}`)

export const verifyUser = () =>
  api.get('/verify') 

export const logoutUser = () =>
  api.get('/logout')

export const getTags = () =>
  api.get('/tags')

export const addRecipe = (formData) =>
  api.post('/recipes', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })

export default api