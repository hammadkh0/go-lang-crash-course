//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//  each dashboard component

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

// * Create functions to calculate averages for each dashboard component
func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}
func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(c.temp); i++ {
		sum += int(c.temp[i])
	}
	return sum / len(c.temp)
}
func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

// * Using struct embedding, create a Dashboard structure that contains
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}
	//* Print out a 5-second average from each component using promoted
	//  methods on the Dashboard

	fmt.Printf("Average bandwidth usage: %v\n", dashboard.AverageBandwidth())
	fmt.Printf("Average CPU temperature: %v\n", dashboard.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dashboard.AverageMemoryUsage())
}
