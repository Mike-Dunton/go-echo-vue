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
  },
  selectRecipe (context, selectedRecipe) {
    context.commit('setRecipe', { update: selectedRecipe})
  }
}

// mutations
const mutations = {
  setRecipes (state, recipes) {
    state.all = recipes
  },
  setRecipe(state, selectedRecipe) {
    console.log("we mutate", selectedRecipe)
    if (selectedRecipe == null || selectedRecipe.update === undefined) {
      state.selected = selectedRecipe
    } else {
       var update = selectedRecipe.update
       state.selected = {...state.selected, ...update}
    }
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}