export const _import_view = (viewName: string) => () => import('@/views/' + viewName + '.vue')
export const _import_com = (viewName: string) => () => import('@/components/' + viewName + '.vue')
