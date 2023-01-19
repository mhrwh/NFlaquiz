import { createStore} from 'vuex'
import createPersistedState from 'vuex-persistedstate'

const store = createStore({
  state: {
    auth: false,
    quizzes: [],
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth,
    setQuizzes: (state, quizzes) => state.quizzes = quizzes,
  },
  actions:{
    setAuth: (context, auth) => context.commit('setAuth', auth),
    setQuizzes: (context, quizzes) => context.commit('setQuizzes', quizzes),
  },
  plugins: [createPersistedState(
    {
      key: 'NFlaquiz',
      
      storage: window.sessionStorage
    })]
})

export default store