import { key } from '../store'
import { useStore } from 'vuex'

import { VueCookieNext } from 'vue-cookie-next'

import { ElNotification } from 'element-plus'

import { RendererNode, RendererElement, VNode } from 'vue'

export const importView = (viewName: string) => () =>
  import('../views/' + viewName + '.vue')

export const importCom = (viewName: string) => () =>
  import('../components/' + viewName + '.vue')

export function ops(
  type?: '' | 'success' | 'warning' | 'info' | 'error' | undefined,
  title?: string | undefined,
  message?:
    | string
    | VNode<
        RendererNode,
        RendererElement,
        {
          [key: string]: any
        }
      >
    | undefined,
  customClass?: string | undefined
) {
  ElNotification({
    title: title,
    message: message,
    type: type,
    customClass: customClass,
  })
}

export function formatDate() {
  //三目运算符
  const Dates = new Date()

  //年份
  const Year: number = Dates.getFullYear()

  //月份下标是0-11
  const Months: any =
    Dates.getMonth() + 1 < 10
      ? '0' + (Dates.getMonth() + 1)
      : Dates.getMonth() + 1

  //具体的天数
  const Day: any =
    Dates.getDate() < 10 ? '0' + Dates.getDate() : Dates.getDate()

  //小时
  const Hours =
    Dates.getHours() < 10 ? '0' + Dates.getHours() : Dates.getHours()

  //分钟
  const Minutes =
    Dates.getMinutes() < 10 ? '0' + Dates.getMinutes() : Dates.getMinutes()

  //秒
  const Seconds =
    Dates.getSeconds() < 10 ? '0' + Dates.getSeconds() : Dates.getSeconds()

  //返回数据格式
  return (
    Year +
    '-' +
    Months +
    '-' +
    Day +
    ' ' +
    Hours +
    ':' +
    Minutes +
    ':' +
    Seconds
  )
}

export function format(Dates: Date) {
  //年份
  const Year: number = Dates.getFullYear()

  //月份下标是0-11
  const Months: any =
    Dates.getMonth() + 1 < 10
      ? '0' + (Dates.getMonth() + 1)
      : Dates.getMonth() + 1

  //具体的天数
  const Day: any =
    Dates.getDate() < 10 ? '0' + Dates.getDate() : Dates.getDate()

  //小时
  const Hours =
    Dates.getHours() < 10 ? '0' + Dates.getHours() : Dates.getHours()

  //分钟
  const Minutes =
    Dates.getMinutes() < 10 ? '0' + Dates.getMinutes() : Dates.getMinutes()

  //秒
  const Seconds =
    Dates.getSeconds() < 10 ? '0' + Dates.getSeconds() : Dates.getSeconds()

  //返回数据格式
  return (
    Year +
    '-' +
    Months +
    '-' +
    Day +
    ' ' +
    Hours +
    ':' +
    Minutes +
    ':' +
    Seconds
  )
}

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

// not allowed
// export function getToken(): String {
//   const store = useStore(key)
//   return store.state.token
// }

export function getToken(): String {
  let t = VueCookieNext.getCookie('token')
  return t
}

// setTimeout 轮询
export function startTimer(duration: number, func: Function) {
  if (duration == -1) {
    return
  }

  duration = func()

  const timer = setTimeout(() => {
    startTimer(duration, func)
  }, duration)
}

export const keyCodeMap = new Map([
  [65, 29], // A
  [66, 30], // B
  [67, 31], // C
  [68, 32], // D
  [69, 33], // E
  [70, 34], // F
  [71, 35], // G
  [72, 36], // H
  [73, 37], // I
  [74, 38], // J
  [75, 39], // K
  [76, 40], // L
  [77, 41], // M
  [78, 42], // N
  [79, 43], // O
  [80, 44], // P
  [81, 45], // Q
  [82, 46], // R
  [83, 47], // S
  [84, 48], // T
  [85, 49], // U
  [86, 50], // V
  [87, 51], // W
  [88, 52], // X
  [89, 53], // Y
  [90, 54], // Z
  [48, 7], // 0
  [49, 8], // 1
  [50, 9], // 2
  [51, 10], // 3
  [52, 11], // 4
  [53, 12], // 5
  [54, 13], // 6
  [55, 14], // 7
  [56, 15], // 8
  [57, 16], // 9
  // [27, 1], // ESC 拦截

  [186, 74], // ;:
  [187, 70], // =+
  [188, 55], // ,<
  [189, 69], // -_
  [190, 56], // .>

  [191, 76], // /?
  [192, 68], // `~
  [219, 71], // [{
  [220, 73], // \|
  [221, 72], // ]}
  [222, 75], // "'
  [20, 115],
])
