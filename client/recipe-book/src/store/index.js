import Vue  from 'vue';
import Vuex from 'vuex';

import recipes from './modules/recipes';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
    modules: {
      recipes
    },
    strict: debug
});