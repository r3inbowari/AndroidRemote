package queue

import bilicoin "CloudGameServer/utils"

type QUp struct {
	up    *EsQueue
	cnt   int
	enPos int
	dePos int
}

var QUP QUp

func InitQueueUP() {
	bilicoin.Info("[QUP] init qup")
	QUP.up = NewQueue(100)
	QUP.cnt = 0
	QUP.enPos = 0
	QUP.dePos = 0
}

// EnQueue
// return 出队编号
func (q *QUp) EnQueue(obj interface{}) int {
	q.up.Put(obj)
	q.enPos++
	return q.enPos
}

// DeQueue unsafe 呃 ...不管了
func (q *QUp) DeQueue() interface{} {
	val, _, _ := q.up.Get()
	q.dePos++
	return val
}

// CalcPrev 计算队列前方剩余数量
func (q *QUp) CalcPrev(i int) int {
	return i - q.dePos
}
