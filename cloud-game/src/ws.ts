import { ComponentInternalInstance } from 'vue'

export class IWebSocket {
	// timer tag
	wsHeartbeatTag: number = 0
	// root instance
	componentInstance: ComponentInternalInstance | null = null
	// hostname
	mHost: string | null = null

	constructor(currentInstance: ComponentInternalInstance | null, host: string | null) {
		// check instance
		if (currentInstance === null) {
			console.error("can not read component internal instance: null");
			return
		}

		if (host !== null) {
			this.mHost = host
		}
		this.componentInstance = currentInstance
	}

	// init ws and connect
	// currentInstance: ComponentInternalInstance | null
	initMsgWebsocket() {

		// connect manually
		// @callback onerror may be call
		if (this.mHost === null) {
			this.componentInstance?.appContext.config.globalProperties.$connect();
		} else {
			this.componentInstance?.appContext.config.globalProperties.$connect(this.mHost);
		}

		// 消息接收
		(this.componentInstance?.appContext.config.globalProperties.sockets).onmessage = (res: {
			data: string
		}) => {
			// msg case
			console.log('[ws] ' + res.data + 1)
		}

		// error handler
		(this.componentInstance?.appContext.config.globalProperties.sockets).onerror = () => {
			console.log('[ws] error found')
		}

		// connection open
		(this.componentInstance?.appContext.config.globalProperties.sockets).onopen = () => {
			console.log('[ws] open')
			this.WsHeartbeat(3 * 1000)
		}

		// connection close
		(this.componentInstance?.appContext.config.globalProperties.sockets).onclose = () => {
			console.log('[ws] close')
			console.log('[ws] clear interval ', this.wsHeartbeatTag);
			clearInterval(this.wsHeartbeatTag)
		}
	}

	WsHeartbeat(heardbeatInterval: number) {
		this.wsHeartbeatTag = setInterval(() => {
			if (
				this.componentInstance?.appContext.config.globalProperties.$socket
					.readyState === 1
			) {
				this.componentInstance?.appContext.config.globalProperties.$socket.send(
					'ping'
				)
			} else {
				console.log('[ws] socket not created')
			}
		}, heardbeatInterval)
	}
}
