import { createRouter, createWebHistory } from 'vue-router'
import Layout from './Layout.vue'
import HomeView from './views/HomeView.vue'
import LoginView from './views/LoginView.vue'
import NotFoundView from './views/NotFoundView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    {
      path: '/',
      component: Layout,
      
      children: [
        {
          path: '',
          name: 'home',
          component: HomeView
        }
      ]
    }, {
      path: '/login',
      name: 'login',
      component: LoginView
    }, {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: NotFoundView 
    }
  ]
})

export default router
