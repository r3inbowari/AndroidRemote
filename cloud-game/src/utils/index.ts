export const importView = (viewName: string) => () =>
  import('../views/' + viewName + '.vue')

export const importCom = (viewName: string) => () =>
  import('../components/' + viewName + '.vue')
