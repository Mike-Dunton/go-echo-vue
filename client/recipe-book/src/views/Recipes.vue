<template>
<div class="home">
  <section class="hero is-dark" v-bind:style="heroStyles">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">
          {{ selected.name }}          
        </h1>
        <h2 class="subtitle">
          {{ selected.description}}
        </h2>
      </div>
    </div>
  </section>
  <RecipeList v-if="_loaded" />
</div>
</template>
<script>
import RecipeList from '../components/RecipeList';
import { mapState } from 'vuex'
export default {
  name: 'recipes',
  computed: {
    _user() {
      return this.$auth.user() || {};
    },
    _loaded() {
      return this.$auth.ready() && this.$auth.check();
    },
    ...mapState({
      recipes: state => state.recipes.all,
      selected: state => state.recipes.selected,
      heroStyles: state => {
        if (!state.recipes.selected || !state.recipes.selected.featuredImage) {
          return {    
            'text-align': 'center',
            'background-color': 'hsl(0, 0%, 96%)',
            'background-size': 'cover',
            'background-position': 'center',
            'background-repeat': 'no-repeat',
            'height':' 375px'
          }
        } else {
          return {    
           'text-align': 'center',
            'background-image': `url('${state.recipes.selected.featuredImage}')`,
            'background-size': 'cover',
            'background-position': 'center',
            'background-repeat': 'no-repeat',
            'height':' 375px'
          }
        }
      }
    })
  },
  components: {
    RecipeList
  }
}
</script>
<style lang="scss" scoped>
  // .hero {    
  //   text-align: center;
  //   background-image: var(--url);
  //   background-size: cover;
  //   background-position: center;
  //   background-repeat: no-repeat;
  //   height: 375px;
  // }
  .hero-body .title {
    text-shadow: 2px 2px 2px rgba(0, 0, 0, 0.6);
  }
  .subtitle {
    text-shadow:  2px 2px 2px rgba(0, 0, 0, 0.7);
  }
  .is-xl {
    font-size: 1.7rem;
  }
</style>