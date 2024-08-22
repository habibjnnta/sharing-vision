import { createRouter, createWebHistory } from 'vue-router';
import Article from '../views/Article.vue';
import Preview from '../views/Preview.vue';
import PreviewDetail from '../views/PreviewDetail.vue';
import NotFoundPage from '../views/404.vue';

const routes = [
  {
    path: '/',
    name: 'Article',
    component: Article,
  },
  {
    path: '/preview',
    name: 'Preview Article',
    component: Preview,
  },
  {
    path: '/preview/:id',
    name: 'Preview Article Detail',
    component: PreviewDetail,
    props: true,
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'Not Found Page',
    component: NotFoundPage,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_BASE_URL),
  routes,
});

export default router;