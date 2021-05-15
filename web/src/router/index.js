import Vue from 'vue';
import VueRouter from 'vue-router';
import userRoute from './module/user';
import store from '@/store'
Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
...userRoute
];

const router = new VueRouter({
  routes,
});
router.beforeEach((to,from,next)=>{
  if(to.matched.some(res=>res.meta.auth)){//判断是否需要登录
      if (store.state.userModule.token) {
          next();
      }else{
          next({
              path:"/login",
              query:{
                  redirect:to.fullPath
              }
          });
      }

  }else{
      next()
  }
});
export default router;
