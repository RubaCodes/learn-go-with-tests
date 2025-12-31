package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second,sleep: time.Sleep }
	Countdown(os.Stdout,sleeper)
}
// Interface to absolve the mocking need of time.Sleep
type Sleeper interface{
	Sleep()
}
// Helper struct needed to implement mocked the sleeper inteface
type DefaultSleeper struct{}
// Implements the default behavior for the Sleep() function
func (d DefaultSleeper) Sleep(){
	time.Sleep(1 * time.Second)
}
// Helper struct needed to implement the mocking sleeper inteface
type SpySleeper struct{
	Calls int
}
// Implements the mocking behavior of the Sleep() function
func (s *SpySleeper)Sleep(){
	s.Calls++
}

type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

// struct to help checking the correct order of execution of the Countdown function
type SpyCountdownOperations struct{
	Calls []string
}
// implements the Sleeper interface
func (spy * SpyCountdownOperations) Sleep(){
	spy.Calls = append(spy.Calls,sleep)
}
//implements the io.Writer interface
func (spy * SpyCountdownOperations) Write( b []byte)(n int, err error){
	spy.Calls = append(spy.Calls,write)
	return
}
const write = "write"
const sleep = "sleep"
const finalWord = "Go!"
const countdownStart = 3
func Countdown(out io.Writer,s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out,i)
		s.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

