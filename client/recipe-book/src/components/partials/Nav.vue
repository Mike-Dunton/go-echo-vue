<template>
<nav class="navbar container" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <strong class="is-size-4">Recipe Book</strong>
    </a>
    <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbar"  @click="showNav = !showNav" :class="{ 'is-active': showNav }">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>
  <div id="navbar" class="navbar-menu" :class="{ 'is-active': showNav }" >
    <div class="navbar-start">
      <router-link to="/" class="navbar-item">Home</router-link>
      <router-link to="/about" class="navbar-item">About</router-link>
    </div>
    <div class="navbar-end" v-if="$auth.ready()">
      <div class="navbar-item">
        <div class="buttons"  v-if="!$auth.check()">
          <button class="button is-dark" @click="oauth2Default('google')">
            <strong>Sign In</strong>
          </button>
        </div>
        <div class="buttons"  v-if="$auth.check()">
          <button class="button is-dark" @click="logout()">
            <strong>Log Out</strong>
          </button>
        </div>
      </div>
      <div class="navbar-item  has-dropdown is-hoverable" v-if="$auth.check()">
        <a class="navbar-link">
          <figure class="image is-24x24">
            <img class="is-rounded" :src="this.$auth.user().picture">
          </figure>
        </a>
        <div class="navbar-dropdown">
          <a class="navbar-item">
            Profile
          </a>
          <a class="navbar-item">
            Lorem
          </a>
          <a class="navbar-item">
            Ipsum
          </a>
        </div>
      </div>
    </div>
  </div>
</nav>
</template>
<script>
export default {
    name: 'Nav',
    data: function() {
        return {
          showNav: false,
          form: {
            data: {},
            code: false,
            params: {
                state: {
                    remember: false,
                    staySignedIn: true,
                    fetchUser: true,
                }
            }
          }    
        }
    },
    watch: {
        '$route.params.type'() {
            this.reset();
        }
    },
    mounted() {
        this.reset();
        if (this.form.code) {
            this.oauth2Default(this.$route.params.type);
        }
    },
    methods: {
        reset() {
            var code  = this.$route.query.code;
            var type  = this.$route.params.type;
            var state = this.$route.query.state;
            delete this.form.url;
            delete this.form.state;
            this.form.data = {};
            this.form.code = code ? true : false;
            if (this.form.code) {
                this.form.url        = 'login/' + type;
                this.form.state      = state
                this.form.data.code  = code;
            }
        },  
        oauth2Default(type) {
            this.$auth.oauth2(type, this.form);
        },
        logout() {
          this.$auth.logout({
            makeRequest: false,
            redirect:'/'
          })
        }
    }
}
</script>
<style lang="scss" scoped>
  nav {
    margin-top: 25px;
    margin-bottom: 30px;
    a {
      font-weight: bold;
      color: #2c3e50;
      &.router-link-exact-active {
        color: #d88d00;
      }
    }  
  } 
</style>