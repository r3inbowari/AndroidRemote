import req from './axios-wrap' // 接口文件

export const HELLO = function (e: string) {
  return req('get', '/hello', { qwe: e })
}
