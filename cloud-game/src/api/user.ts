import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// post reg
export const userRegister = function (mobile: string, password: string) {
  return reqV(POST, '/reg', { mobile: mobile, password: password })
}

// post login raw
export const userLogin = function (mobile: string, password: string) {
  return reqV(POST, '/login', { mobile: mobile, password: password })
}

// get user info
export const userInfo = function () {
  return reqV(GET, '/info', {})
}