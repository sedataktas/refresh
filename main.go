package main

func main() {
	quit := make(chan struct{})
	go watchCron()
	watch()
	<-quit
}
