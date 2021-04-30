package main

import "fmt"

type request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

func run(op binOp, req *request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *request) {
	for {
		req := <-service
		go run(op, req)
	}
}

func startServer(op binOp) chan *request {
	reqChan := make(chan *request)
	go server(op, reqChan)
	return reqChan
}

func main() {
	adder := startServer(func(a, b int) int {
		return a + b
	})
	const N = 100
	var reqs [N]request
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
	fmt.Println("done")
}
