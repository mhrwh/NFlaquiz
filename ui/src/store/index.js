import { createStore} from 'vuex'
import createPersistedState from 'vuex-persistedstate'

const store = createStore({
  state: {
    auth: false,
    quizzes: [],
    mapMode: "correctAnswersRate",
  },
  mutations: {
    setAuth: (state, auth) => state.auth = auth,
    setQuizzes: (state, quizzes) => state.quizzes = quizzes,
    setMapMode: (state, mapMode) => state.mapMode = mapMode,
  },
  actions:{
    setAuth: (context, auth) => context.commit('setAuth', auth),
    setQuizzes: (context, quizzes) => context.commit('setQuizzes', quizzes),
    setMapMode: (context, mapMode) => context.commit('setMapMode', mapMode),
  },
  plugins: [createPersistedState(
    {
      key: 'NFlaquiz',
      
      storage: window.sessionStorage
    })]
})

export default store
