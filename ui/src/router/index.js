import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '@/pages/TopPage'
import EditBookmark from '@/pages/EditBookmark'

const routes = [
  {
    path: '/',
    name: 'TopPage',
    component: TopPage
  },
  {
    path: '/editbookmark',
    name: 'EditBookmark',
    component: EditBookmark
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router