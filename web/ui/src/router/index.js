import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from '../views/PoolList.vue';
import UserPool from '../views/PoolView';
import UserView from '../views/UserView';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'user-pools',
    component: Home,
  },
  {
    path: '/pool/:poolId',
    name: 'user-pool',
    component: UserPool,
  },
  {
    path: '/pool/:poolId/users/:username',
    name: 'user-view',
    component: UserView,
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
