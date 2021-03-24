export const importView = (viewName: string) => () =>
  import('../views/' + viewName + '.vue')

export const importCom = (viewName: string) => () =>
  import('../components/' + viewName + '.vue')

// 打印环境
/**
 *
 */
export function printEnv() {
  if (import.meta.env.DEV) print('mode: development')
  else print('mode: production')
}

export function print(o: string) {
  console.log('🌸 %c' + o, 'color:pink')
}
