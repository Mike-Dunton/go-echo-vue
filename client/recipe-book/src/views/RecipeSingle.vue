<template>
  <div class="recipe-single">
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            {{ recipe.name }}
          </h1>
          <h2 class="subtitle ">
            {{recipe.category}}
          </h2>
        </div>
      </div>
    </section>
    <section class="recipe-content">
      <div class="container">
        <p class="is-size-4 description">{{recipe.description}}</p>
        <div class="recipe-images columns is-multiline has-text-centered">
          <div v-for="image in recipe.images" :key="image.id" class="column is-one-third">
            <img :src="image" :alt="recipe.name">
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
  import { mapState } from 'vuex'
  export default {
    name: 'RecipeSingle',
    computed: mapState({
      recipe: state => state.recipes.selected
    }),
    created () {
      const ID = Number(this.$route.params.id);
      this.$store.dispatch('recipes/getRecipe', ID)
    }
  }
</script>
<style lang="scss" scoped>
  .recipe-single {
    margin-top: 30px;
  }
  .hero {
    margin-bottom: 70px;
  }
  .recipe-images {
    margin-top: 50px;
  }
  .description {
    margin-bottom: 30px;
  }
</style>