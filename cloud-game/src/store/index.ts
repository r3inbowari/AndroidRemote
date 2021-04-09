// import { State } from '@vue/runtime-core'
// @ts-ignore
import { InjectionKey, App } from 'vue'

import { createStore, StoreOptions, Store } from 'vuex'

// import app from "../main"
import { getCurrentInstance } from 'vue'
import app from '../main'

// declare module '@vue/runtime-core' {
// 	// define your typings for the store state
// 	export interface State {
// 		count: number
// 	}

// 	interface ComponentCustomProperties {
// 		$store: Store<State>
// 	}
// }
export interface State {
  count: number
  avatar: String
  token: String
  // socket: {
  //   // 连接状态
  //   isConnected: Boolean,
  //   // 消息内容
  //   message: String,
  //   // 重新连接错误
  //   reconnectError: Boolean,
  //   // 心跳消息发送时间
  //   heartBeatInterval: number,
  //   // 心跳定时器
  //   heartBeatTimer: number
  // }
}

// define injection key
export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
  state: {
    count: 1,
    avatar: 'null',
    token: '',

    // ws socket
    // socket: {
    //   // 连接状态
    //   isConnected: false,
    //   // 消息内容
    //   message: "",
    //   // 重新连接错误
    //   reconnectError: false,
    //   // 心跳消息发送时间
    //   heartBeatInterval: 50000,
    //   // 心跳定时器
    //   heartBeatTimer: 0
    // }
  },
  mutations: {
    increment(state) {
      // mutate state
      state.count++
    },
    setToken(state, t) {
      
      state.token = t
      console.log('[store] setToken ' + state.token)
    },
    setAvatar(state, t) {
      state.avatar = t
      console.log('[store] setAvatar ' + state.avatar)
    },


    // 连接打开
    // SOCKET_ONOPEN(state, event) {
    //   // const app = getCurrentInstance()

    //   console.log('open', app);
      

    //  // app?.appContext.config.globalProperties.$socket = event.currentTarget;
    //   state.socket.isConnected = true;
    //   // 连接成功时启动定时发送心跳消息，避免被服务器断开连接
    //   state.socket.heartBeatTimer = setInterval(() => {
    //     const message = "心跳消息";
    //     state.socket.isConnected &&
    //     app?.appContext.config.globalProperties.$socket.sendObj({
    //         code: 200,
    //         msg: message
    //       });
    //   }, state.socket.heartBeatInterval);
    // },
    // // 连接关闭
    // SOCKET_ONCLOSE(state, event) {
    //   state.socket.isConnected = false;
    //   // 连接关闭时停掉心跳消息
    //   clearInterval(state.socket.heartBeatTimer);
    //   state.socket.heartBeatTimer = 0;
    //   console.log("连接已断开: " + new Date());
    //   console.log(event);
    // },
    // // 发生错误
    // SOCKET_ONERROR(state, event) {
    //   console.error(state, event);
    // },
    // // 收到服务端发送的消息
    // SOCKET_ONMESSAGE(state, message) {
    //   state.socket.message = message;
    // },
    // // 自动重连
    // SOCKET_RECONNECT(state, count) {
    //   console.info("消息系统重连中...", state, count);
    // },
    // // 重连错误
    // SOCKET_RECONNECT_ERROR(state) {
    //   state.socket.reconnectError = true;
    // }


  },
})
