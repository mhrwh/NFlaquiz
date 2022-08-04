import { createStore} from 'vuex'
import createPersistedState from 'vuex-persistedstate'

const store = createStore({
  state: {
    auth: false,
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth,
  },
  actions:{
    setAuth: (context, auth) => context.commit('setAuth', auth)
  },
  plugins: [createPersistedState(
    {
      key: 'NFlaquiz',
      
      storage: window.sessionStorage
    })]
})

export default store