import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Recipes from '../views/Recipes.vue'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/recipes',
    name: 'Recipes',
    component: Recipes,
    meta: {
      auth: true
    }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/recipe/new',
    name: 'recipeNew',
    component: () => import('../views/RecipeNew.vue'),
    meta: {
      auth: true
    }
  },
  {
    path: '/recipe/:id',
    name: 'recipeSingle',
    component: () => import('../views/RecipeSingle.vue'),
    meta: {
      auth: true
    }
  },
  {
    path: '/login/:type',
    name: 'auth-login-social',
    component: Home
  }
]

Vue.router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default Vue.router;
