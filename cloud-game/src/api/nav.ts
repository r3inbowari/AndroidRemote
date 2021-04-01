import { req, reqV } from './axios-wrap' // interface

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
