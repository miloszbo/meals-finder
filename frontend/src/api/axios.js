import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080",
    headers: {
        'Content-Type':'application/json'
    },
    withCredentials:true,
})

export const getAllRecipes = (tags) =>
  API.get('/recipes', tags ? { params: { tags } } : {})

export const getRecipeById = (id) =>
  API.get(`/recipes/${id}`)

export default api