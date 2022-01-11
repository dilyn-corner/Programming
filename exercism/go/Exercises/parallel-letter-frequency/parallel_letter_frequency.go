/*
Package letter counts the frequency of letters in a given text.
It utilizes concurrency to speed this process up substantially.
*/
package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {

	// Create a buffered channel to alleviate some pressure between our
	// goroutines and our channel.
	freqChan := make(chan FreqMap, len(l))

	// Create a WaitGroup, which acts as a way of controlling concurrency.
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(l))

	// This anonymous function acts as a way to close our channel once all of
	// our goroutines are finished.
	go func() {
		waitGroup.Wait()
		close(freqChan)
	}()

	// Create a number of goroutines equal to the number of strings in l
	// When the goroutine has filled its channel, it decrements our wait count.
	for i := 0; i < len(l); i++ {
		y := l[i]
		go func(y string) {
			defer waitGroup.Done()
			freqChan <- Frequency(y)
		}(y)
	}

	// Create a blank map to store the results of our goroutines.
	finalMap := make(FreqMap)
	// Range over the contents of the channels we have painstakingly filled...
	for freqMap := range freqChan {
		// ... and add the contents to our final tally.
		for k, v := range freqMap {
			finalMap[k] += v
		}
	}
	return finalMap
}
