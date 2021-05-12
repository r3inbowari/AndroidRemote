import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// heap
export const heap = function () {
  return req(GET, '/m/sys/heap', {})
}

// load
export const load = function () {
  return req(GET, '/m/sys/load', {})
}
