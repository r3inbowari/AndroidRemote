import axios from 'axios'
// post format
// import qs from 'qs'
// remove qs due to the backend server using gin

import { getToken } from '../utils'

// version param
const VERSION =
  import.meta.env['VITE_SERVER_API_VERSION'] === undefined
    ? ''
    : '/' + import.meta.env['VITE_SERVER_API_VERSION']

// from env base host setting
const BAER_URL =
  import.meta.env['VITE_AXIOS_BASE_URL'] === undefined
    ? 'http://192.168.5.67:8080/'
    : import.meta.env['VITE_AXIOS_BASE_URL']?.toString()

// axios instance create
var instance = axios.create({
  baseURL: BAER_URL, //接口统一域名
  timeout: 6000, //设置超时
  withCredentials: false,
})

// request interceptors
instance.interceptors.request.use(
  function (config) {
    let token = getToken();
    if (token) {
      config.headers["Authorization"] = "Bearer " + token;
    } else {
      console.log('token is not exist');
    }
    return config
  },
  function (error) {
    // do someting

    return Promise.reject(error)
  }
)

// response interceptors
instance.interceptors.response.use(
  function (response) {
    return response.data
  },
  function (error) {
    // do someting
    console.log('[axios] response reject')
    return Promise.reject(error)
  }
)

/**
 * export function req
 * @param {String} method  请求的方法：get、post、delete、put
 * @param {String} url     请求的url:
 * @param {Object} data    请求的参数
 * @returns {Promise}     返回一个promise对象
 */
export function req(method: string, url: string, data: object): Promise<any> {
  method = method.toLowerCase()
  if (method == 'post') {
    // return instance.post(url, qs.stringify(data))
    return instance.post(url, data)
  } else if (method == 'get') {
    return instance.get(url, { params: data })
  } else if (method == 'delete') {
    return instance.delete(url, { params: data })
  } else if (method == 'put') {
    return instance.put(url, data)
    // return instance.put(url, qs.stringify(data))
  } else {
    console.error('known method' + method)
    throw Error('known method')
  }
}

/**
 * export function req using version control
 * @param {String} method  请求的方法：get、post、delete、put
 * @param {String} url     请求的url:
 * @param {Object} data    请求的参数
 * @returns {Promise}     返回一个promise对象
 */
export function reqV(method: string, url: string, data: object): Promise<any> {
  url = VERSION + url
  method = method.toLowerCase()
  if (method == 'post') {
    // return instance.post(url, qs.stringify(data))
    return instance.post(url, data)
  } else if (method == 'get') {
    return instance.get(url, { params: data })
  } else if (method == 'delete') {
    return instance.delete(url, { params: data })
  } else if (method == 'put') {
    return instance.put(url, data)
    // return instance.put(url, qs.stringify(data))
  } else {
    console.error('known method' + method)
    throw Error('known method')
  }
}
