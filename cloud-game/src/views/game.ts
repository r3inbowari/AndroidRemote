import { Router, useRouter } from 'vue-router'

export function exitGame(r: Router) {
  // 在这里处理退出程序
  r.replace({
    name: 'Home',
  })
}
