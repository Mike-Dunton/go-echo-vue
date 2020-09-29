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
        <div class="button-block" >
          <button v-if="!$auth.check()" class="button is-xl is-dark">Sign Up to Add Your Own Recipes</button>
          <h3 v-if="$auth.check()" class="is-size-5 has-background-dark welcome">Welcome, {{ _user.name }}!</h3>
        </div>
      </div>
    </div>
  </section>
  <RecipeList />
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
      return this.$auth.ready();
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
  .button-block {
    text-align: center;
    margin-left: auto;
    margin-right: auto;
    width: 100%;
    position: absolute;
    bottom: -100px;
    .button {
      margin-right: 50px;
      padding-left: 50px;
      padding-right: 50px;
    }
    .welcome {
      width: 375px;
      padding: 9px;
      margin-left: auto;
      margin-right: auto;
    }
  }
  .is-xl {
    font-size: 1.7rem;
  }
</style>