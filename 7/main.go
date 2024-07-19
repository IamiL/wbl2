package main

import (
	"fmt"
	"reflect"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	slCases := []reflect.SelectCase{}
	commonCh := make(chan interface{})
	for _, channel := range channels {
		slCases = append(slCases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(channel),
		})
	}
	go func() {
		chosen, recv, ok := reflect.Select(slCases)
		fmt.Println(chosen)
		fmt.Println(recv)
		fmt.Println(ok)
		close(commonCh)
	}()
	return commonCh
}

func MainChannels() {

	var sig = func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	// Ожидаем запись в канал, что основанная горутина не завершилась всех остальных
	<-or(
		sig(1*time.Second),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
