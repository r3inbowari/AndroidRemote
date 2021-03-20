export const _import_view = viewName => () => import('@/views/' + viewName + '.vue')
export const _import_com = viewName => () => import('@/components/' + viewName + '.vue')