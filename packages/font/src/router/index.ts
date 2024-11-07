import { createRouter, createWebHistory } from 'vue-router';

export const constantRoutes = [
  {
    path: '/',
    component: () => import('./'),
  },
];

// 创建路由
const router = createRouter({
  history: createWebHistory(),
  routes: constantRoutes,
});

export default router;
