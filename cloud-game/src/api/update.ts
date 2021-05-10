import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get update
export const getUpdate = function () {
  return req(GET, '/public/home/update', {})
}
