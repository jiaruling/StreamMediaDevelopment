package limiter

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(concurrentConn int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: concurrentConn,
		bucket:         make(chan int, concurrentConn),
	}
}

func (c *ConnLimiter) GetConn() bool {
	if len(c.bucket) >= c.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	c.bucket <- 1
	return true
}

func (c *ConnLimiter) ReleaseConn() {
	<-c.bucket
	log.Printf("New connection coming")
}
