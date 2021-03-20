import { router } from './index'

// router.beforeEach((to, from, next) => {
//     if (to.meta.login) {
//         if (store.state.user.user.name) {
//             next()
//         }else{
//             Vue.prototype.$alert('请先登录!')
//                 .then( () => {
//                     store.state.user.isLogin = true
//                 })
//         }
//     }else{
//         if (to.meta.page) store.state.app.pageLoading = true
//         next() 
//     }
    
// })

// router.afterEach((to, from) => {
//     if (to.meta.page) store.state.app.pageLoading = false
//     document.title = to.name
// })

// export default router