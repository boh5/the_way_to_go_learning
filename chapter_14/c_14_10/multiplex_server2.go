package main

import "fmt"

type Request struct {
	a, b   int
	replyc chan int
}

type BinOp func(a, b int) int

func Run(op BinOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func Server(op BinOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go Run(op, req)
		case <-quit:
			return
		}

	}
}

func StartServer(op BinOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go Server(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := StartServer(func(a, b int) int {
		return a + b
	})
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}

	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request", i, "is ok!")
		}
	}
	quit <- true
	fmt.Println("done")
}
