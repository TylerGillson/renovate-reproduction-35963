package main

import "github.com/onsi/ginkgo/v2"

func main() {
	ginkgo.It("should not produce a warning", func() {
		ginkgo.By("doing nothing")
	})
}
