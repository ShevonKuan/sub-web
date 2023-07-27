import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        name: "SubConverter",
        component: () => import("../views/Subconverter.vue")
    },
    {
        path: "/controller",
        name: "SubController",
        component: () => import("../views/Subcontroller.vue"),
        meta: {requiresAuth: true}
    },
    {
        path: "/login",
        name: "Login",
        component: () => import("../views/Login.vue")
    },
];

const router = new VueRouter({
    mode: "hash",
    base: process.env.BASE_URL,
    routes
});

// 添加导航守卫
router.beforeEach((to, from, next) => {
    // 判断是否需要鉴权
    if (to.matched.some(record => record.meta.requiresAuth)) {
        // 判断是否已经登录
        if (!localStorage.getItem('token')) {
            // 未登录，跳转到登录页
            next('/login')
        } else {
            // 已登录，允许访问
            next()
        }
    } else {
        // 不需要鉴权，直接允许访问
        next()
    }
})

export default router;
