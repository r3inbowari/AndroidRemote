import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get backend version
export const getVersion = function () {
  return req(GET, '/public/version', {})
}
