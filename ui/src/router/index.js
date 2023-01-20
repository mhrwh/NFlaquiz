import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '@/pages/TopPage'
// import EditBookmark from '@/component/EditBookmark'

const routes = [
    {
      path: '/',
      name: 'TopPage',
      component: TopPage
    },
    {
      path: '/quiz',
      name: 'QuizPage',
      component: () =>
        import("../pages/QuizPage.vue"),
    }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router