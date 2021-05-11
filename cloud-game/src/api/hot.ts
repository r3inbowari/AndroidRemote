import { req, reqV } from './axios-wrap' // interface

const GET: string = 'get'
const POST: string = 'post'
const PUT: string = 'put'
const DEL: string = 'delete'

// get hot
export const getHot = function () {
  return req(GET, '/public/home/hot', {})
}

// del hot
export const delHot = function (id: string) {
  return req(DEL, '/public/home/hot/' + id, {})
}

// add hot
export const addHot = function (data: object) {
  return req(POST, '/public/home/hot', data)
}
