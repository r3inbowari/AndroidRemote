import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get detail
export const getDetail = function (aid: string) {
  return req(GET, '/public/detail/' + aid, {})
}
