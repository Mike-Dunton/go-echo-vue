<template>
  <div class="recipes container">
    <nav class="level">
      <!-- Left side -->
      <div class="level-left">
        <div class="level-item">
          <div class="field has-addons">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Find a Recipe">
              <span class="icon is-small is-left">
                <font-awesome-icon icon="utensils" />
              </span>
            </p>
            <p class="control">
              <button class="button">
                Search
              </button>
            </p>
          </div>
        </div>
      </div>

      <!-- Right side -->
      <div class="level-right">
        <div class="level-item">
          <div class="buttons has-addons">
            <button class="button">All</button>
            <button class="has-badge-rounded has-badge-danger has-badge-left has-badge-medium button is-link is-selected" v-bind:data-badge="recipeCount">Food</button>
            <button class="button">Drinks</button>
          </div>
        </div>
        <p class="level-item"><router-link to="/recipe/new" tag="button" class="button is-success">New Recipe</router-link></p>
      </div>
    </nav>
    <div class="columns is-multiline is-mobile">
      <div v-for="recipe in recipes" :recipe="recipe" :key="recipe.id" v-on:click="selectRecipe(recipe)" class="column is-one-quarter">
        <RecipeCard :recipe="recipe" />
      </div>
    </div>
  </div>
</template>
<script>
  import RecipeCard from '@/components/RecipeCard';
  import { mapState, mapActions } from 'vuex'
  export default {
    name: 'RecipesList',
    components : {
      RecipeCard
    },
    methods: mapActions("recipes", [
        'selectRecipe'
      ]),
    computed: mapState({
      recipes: state => state.recipes.all,
      recipeCount: state => state.recipes.all.length,
      selected: state => state.recipes.selected
    }),
    created () {
      this.$store.dispatch('recipes/getAllRecipes')
    }
  }
</script>
<style lang="scss" scoped>
  .recipes {
    margin-top: 50px;
  }
</style>