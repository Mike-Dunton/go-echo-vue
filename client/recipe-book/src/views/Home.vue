<template>
<div class="home">
  <section class="hero is-dark">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">
          Welcome to the Recipe Book          
        </h1>
        <h2 class="subtitle">
          Check out recently added recipes below
        </h2>
      </div>
    </div>
  </section>
  <RecipeList v-if="_loaded" />
</div>
</template>
<script>
import RecipeList from '../components/RecipeList';
export default {
  name: 'home',
  computed: {
    _user() {
      return this.$auth.user() || {};
    },
    _loaded() {
      return this.$auth.ready() && this.$auth.check();
    }
  },
  components: {
    RecipeList
  }
}
</script>
<style lang="scss" scoped>
  .hero {    
    text-align: center;
    background-image: url('https://arc-anglerfish-washpost-prod-washpost.s3.amazonaws.com/public/V2R2L76BXQ6UBKO7GC3KMYBVWQ.jpg');
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    height: 375px;
  }
  .hero-body .title {
    text-shadow: 2px 2px 2px rgba(0, 0, 0, 0.6);
    padding: 10px 0 15px 0;
    font-size: 60px;
  }
  .subtitle {
    text-shadow:  2px 2px 2px rgba(0, 0, 0, 0.7);
    font-size: 30px;
  }
  .is-xl {
    font-size: 1.7rem;
  }
</style>