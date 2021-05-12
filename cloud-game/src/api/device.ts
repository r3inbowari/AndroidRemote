import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get device count
export const getDeviceCount = function () {
  return req(GET, '/m/device/count', {})
}

// get device all
export const getDeviceAll = function () {
  return req(GET, '/m/device/all', {})
}
