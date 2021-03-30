import req from './axios-wrap' // 接口文件

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

/**
 * api 接口列表
 */

// get hello
export const HELLO = function (e: string) {
  return req(GET, '/hello', { qwe: e })
}

// get hello
export const HELLOQ = function (e: string) {
  return req(POST, '/reg', { mobile: e, password: e })
}
