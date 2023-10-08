package main

// written by Don Jackson for Walmart take home test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	RingSize int    `json:"ringSize"`
	FileName string `json:"fileName"`
}

type Data struct {
	Datetime  string `json:"datetime"`
	Value     string `json:"value"`
	Partition string `json:"partition"`
}

// RingBuffer is a simple ring buffer data structure for Data
type RingBuffer struct {
	buffer   []Data
	size     int
	readIdx  int
	writeIdx int
	count    int
}

func main() {
	config := getConfig()
	fmt.Println("Ring Size=", config.RingSize)
	fmt.Println("File Name=", config.FileName)

	rb := NewRingBuffer(config.RingSize)
	// Start a goroutine to continuously fetch data
	updates := make(chan string)
	go fetchData(rb, config.FileName, updates)
	for {
		// time.Sleep(1000 * time.Millisecond)
		data, found := rb.Pop()
		if found {
			rb.Process(data)
		}
	}
}

func fetchData(rb *RingBuffer, fileName string, updates chan string) {
	dataFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening data file:", err)
		return
	}
	defer dataFile.Close()
	decoder := json.NewDecoder(dataFile)

	for {
		var data Data

		if err := decoder.Decode(&data); err != nil {
			// Check for end of file
			if err == io.EOF {
				break // Reached the end of the file
			}

			fmt.Println("Error decoding data:", err)
			return
		}
		rb.Push(data)
		// time.Sleep(1000 * time.Millisecond)
	}

}

func getConfig() Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return Config{}
	}
	defer configFile.Close()

	// Decode the JSON data into a Config struct
	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return Config{}
	}

	// Access the constants
	return config
}

// NewRingBuffer creates a new ring buffer with the specified size
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer:   make([]Data, size),
		size:     size,
		readIdx:  0,
		writeIdx: 0,
		count:    0,
	}
}

// Process does something with the delement from the ring buffer - for now it just prints it
func (rb *RingBuffer) Process(data Data) {
	s := fmt.Sprintf("Process: Date/Time: %s - Value %s - Partition %s", data.Datetime, data.Value, data.Partition)
	fmt.Println(s)
}

// Push adds a Data element to the ring buffer, overwriting the oldest element if the buffer is full
func (rb *RingBuffer) Push(data Data) {
	fmt.Println("Add data: ", data)
	rb.buffer[rb.writeIdx] = data
	rb.writeIdx = (rb.writeIdx + 1) % rb.size

	if rb.count < rb.size {
		rb.count++
	} else {
		rb.readIdx = (rb.readIdx + 1) % rb.size // Overwrite the oldest element
	}
}

// Pop removes and returns the oldest Data element from the ring buffer
func (rb *RingBuffer) Pop() (Data, bool) {
	if rb.count == 0 {
		return Data{}, false // Buffer is empty
	}

	data := rb.buffer[rb.readIdx]
	rb.readIdx = (rb.readIdx + 1) % rb.size
	rb.count--

	return data, true
}

// Len returns the current number of elements in the ring buffer
func (rb *RingBuffer) Len() int {
	return rb.count
}
