package lib

const MAX_QUEUE_LEN = 100
const MAX_QUEUE_COUNT = 50

type Queue struct {
	List []*Package
}

var Pool = make(map[string]*Queue)

func CreateQueue(name string) *Queue {
	queue := &Queue{}
	Pool[name] = queue

	return queue
}

func (q *Queue) Push(data *Package) bool {
	q.List = append(q.List, data)

	return true
}

func (q *Queue) Count() int {
	return len(q.List)
}

func (q *Queue) Pull() *Package {
	return q.List[0]
}

func (q *Queue) Check() []*Package {
	return q.List
}

func DeleteQueue() {

}
