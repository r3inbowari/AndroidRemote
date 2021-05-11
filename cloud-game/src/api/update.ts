import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get update
export const getUpdate = function () {
  return req(GET, '/public/home/update', {})
}

// del update
export const delUpdate = function (id: string) {
  return req(DEL, '/public/home/update/' + id, {})
}

// add update
export const addUpdate = function (data: object) {
  return req(POST, '/public/home/update', data)
}
