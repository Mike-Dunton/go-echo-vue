<template>
  <div class="recipes container">
    <h2 class="subtitle is-3">
    Check out our new recipes
    </h2>
    <div class="columns is-multiline">
      <div v-for="recipe in recipes" :recipe="recipe" :key="recipe.id" class="column is-one-quarter">
        <router-link :to="'/recipe/' + recipe.id">
          <RecipeCard :recipe="recipe" />
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
  import RecipeCard from '@/components/RecipeCard';
  import { mapState } from 'vuex'
  export default {
    name: 'RecipesList',
    components : {
      RecipeCard
    },
    computed: mapState({
      recipes: state => state.recipes.all
    }),
    created () {
      this.$store.dispatch('recipes/getAllRecipes')
    }
  }
</script>
<style lang="scss" scoped>
  .recipes {
    margin-top: 100px;
  }
</style>