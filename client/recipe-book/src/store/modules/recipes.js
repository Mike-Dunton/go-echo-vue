import recipeService from '../../services/recipes'

// initial state
const state = () => ({
  selected: {},
  all: []
})

// getters
const getters = {}

// actions
const actions = {
  getAllRecipes ({ commit }) {
    recipeService.getRecipes(recipes => {
      commit('setRecipes', recipes)
    })
  },
  getRecipe ({ commit }, id) {
      recipeService.getRecipe(id, recipe => {
          commit('setRecipe', recipe)
      })
  }
}

// mutations
const mutations = {
  setRecipes (state, recipes) {
    state.all = recipes
  },
  setRecipe(state, recipe) {
      state.selected = recipe
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}