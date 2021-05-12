import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get ws/wss sessions
export const getSessions = function () {
  return req(GET, '/m/sessions', {})
}

// kill ws/wss sessions
export const killSession = function (stub: string) {
  return req(GET, '/m/kill/' + stub, {})
}
