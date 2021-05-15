const userRoute = [
  {
    path: '/regi',
    name: 'regi',
    component: () => import('@/views/Regi.vue'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/Login.vue'),
  },{
    path: '/profile',
    name: 'profile',
    meta:{
      auth:true,
    },
    component: () => import('@/views/Profile.vue'),
  },
]
export default userRoute