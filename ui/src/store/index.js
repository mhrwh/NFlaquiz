import { createStore} from 'vuex'
import createPersistedState from 'vuex-persistedstate'

const store = createStore({
  state: {
    auth: false,
    quizDialog: false,
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth,
    setQuizDialog: (state, quizDialog) => state.quizDialog = quizDialog,
  },
  actions:{
    setAuth: (context, auth) => context.commit('setAuth', auth),
    setQuizDialog: (context, quizDialog) => context.commit('setQuizDialog', quizDialog),
  },
  plugins: [createPersistedState(
    {
      key: 'NFlaquiz',
      
      storage: window.sessionStorage
    })]
})

export default store