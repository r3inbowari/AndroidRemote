import axios from 'axios'
// post format
// import qs from 'qs'
// remove qs

var instance = axios.create({
  baseURL: import.meta.env['VITE_AXIOS_BASE_URL']?.toString(), //接口统一域名
  timeout: 6000, //设置超时
  withCredentials: false,
})

// request interceptors
instance.interceptors.request.use(
  function (config) {
    return config
  },
  function (error) {
    // 对请求错误做些什么

    return Promise.reject(error)
  }
)

// response interceptors
instance.interceptors.response.use(
  function (response) {
    return response.data
  },
  function (error) {
    // 对响应错误做点什么
    console.log('拦截器报错')
    return Promise.reject(error)
  }
)

/**
 * export default function
 * @param {String} method  请求的方法：get、post、delete、put
 * @param {String} url     请求的url:
 * @param {Object} data    请求的参数
 * @returns {Promise}     返回一个promise对象
 */
export default function (method: string, url: string, data: object) {
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
    return false
  }
}
