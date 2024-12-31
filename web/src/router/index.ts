import { createRouter, createWebHashHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import type { App } from 'vue'
import { Layout } from '@/utils/routerHelper'
import { useI18n } from '@/hooks/web/useI18n'

const { t } = useI18n()

export const constantRouterMap: AppRouteRecordRaw[] = [
  {
    path: '/',
    component: Layout,
    redirect: '/one',
    name: 'Root',
    meta: {
      hidden: true
    }
  },
  {
    path: '/redirect',
    component: Layout,
    name: 'Redirect',
    children: [
      {
        path: '/redirect/:path(.*)',
        name: 'Redirect1',
        component: () => import('@/views/Redirect/Redirect.vue'),
        meta: {}
      }
    ],
    meta: {
      hidden: true,
      noTagsView: true
    }
  },
  {
    path: '/login',
    component: () => import('@/views/Login/Login.vue'),
    name: 'Login',
    meta: {
      hidden: true,
      title: t('router.login'),
      noTagsView: true
    }
  },
  {
    path: '/404',
    component: () => import('@/views/Error/404.vue'),
    name: 'NoFind',
    meta: {
      hidden: true,
      title: '404',
      noTagsView: true
    }
  }
]

export const asyncRouterMap: AppRouteRecordRaw[] = [
  // {
  //   path: '/level',
  //   component: Layout,
  //   redirect: '/level/menu1/menu1-1/menu1-1-1',
  //   name: 'Level',
  //   meta: {
  //     title: t('router.level'),
  //     icon: 'carbon:skill-level-advanced'
  //   },
  //   children: [
  //     {
  //       path: 'menu1',
  //       name: 'Menu1',
  //       component: getParentLayout(),
  //       redirect: '/level/menu1/menu1-1/menu1-1-1',
  //       meta: {
  //         title: t('router.menu1')
  //       },
  //       children: [
  //         {
  //           path: 'menu1-1',
  //           name: 'Menu11',
  //           component: getParentLayout(),
  //           redirect: '/level/menu1/menu1-1/menu1-1-1',
  //           meta: {
  //             title: t('router.menu11'),
  //             alwaysShow: true
  //           },
  //           children: [
  //             {
  //               path: 'menu1-1-1',
  //               name: 'Menu111',
  //               component: () => import('@/views/Level/Menu111.vue'),
  //               meta: {
  //                 title: t('router.menu111')
  //               }
  //             }
  //           ]
  //         },
  //         {
  //           path: 'menu1-2',
  //           name: 'Menu12',
  //           component: () => import('@/views/Level/Menu12.vue'),
  //           meta: {
  //             title: t('router.menu12')
  //           }
  //         }
  //       ]
  //     },
  //     {
  //       path: 'menu2',
  //       name: 'Menu2',
  //       component: () => import('@/views/Level/Menu2.vue'),
  //       meta: {
  //         title: t('router.menu2')
  //       }
  //     }
  //   ]
  // },
  {
    path: '/one',
    component: Layout,
    redirect: '/one/one-1',
    name: 'One',
    meta: {
      title: t('router.one')
    },
    children: [
      {
        path: 'one-1',
        name: 'One-1',
        component: () => import('@/views/One/index.vue'),
        meta: {
          title: t('router.one')
        }
      }
    ]
  },
  {
    path: '/config',
    component: Layout,
    redirect: '/config/first',
    name: 'Config',
    meta: {
      title: t('router.base')
    },
    children: [
      {
        path: 'first',
        name: 'First',
        component: () => import('@/views/Config/index.vue'),
        meta: {
          title: t('router.config')
        }
      },
      {
        path: 'info',
        name: 'Info',
        component: () => import('@/views/Info/index.vue'),
        meta: {
          title: t('router.info')
        }
      }
    ]
  }
  // {
  //   path: '/info',
  //   component: Layout,
  //   redirect: '/info/first',
  //   name: 'Info',
  //   meta: {
  //     title: t('router.alertmanger')
  //   },
  //   children: [
  //     {
  //       path: 'first',
  //       name: 'First',
  //       component: () => import('@/views/Info/index.vue'),
  //       meta: {
  //         title: t('router.first')
  //       }
  //     }
  //   ]
  // }
]

const router = createRouter({
  history: createWebHashHistory(),
  strict: true,
  routes: constantRouterMap as RouteRecordRaw[],
  scrollBehavior: () => ({ left: 0, top: 0 })
})

export const resetRouter = (): void => {
  const resetWhiteNameList = ['Redirect', 'Login', 'NoFind', 'Root']
  router.getRoutes().forEach((route) => {
    const { name } = route
    if (name && !resetWhiteNameList.includes(name as string)) {
      router.hasRoute(name) && router.removeRoute(name)
    }
  })
}

export const setupRouter = (app: App<Element>) => {
  app.use(router)
}

export default router
