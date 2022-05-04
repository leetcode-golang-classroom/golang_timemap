package time_map

import "log"

type Node struct {
	timestamp int
	value     string
}
type Values *[]Node
type TimeMap struct {
	Stores map[string]Values
}

func Constructor() TimeMap {
	return TimeMap{
		Stores: make(map[string]Values),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	stores, ok := this.Stores[key]
	// find insert position
	if !ok {
		this.Stores[key] = &[]Node{{value: value, timestamp: timestamp}}
		return
	}
	pos := FindInsertPosition(stores, timestamp)
	right := append([]Node{}, ((*stores)[pos+1:])...)
	left := append((*stores)[0:pos+1], Node{value: value, timestamp: timestamp})
	*stores = append(left, right...)
	this.Stores[key] = stores
	log.Printf("stores=%v, pos=%v, left=%v, right=%v", stores, pos, left, right)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	stores, ok := this.Stores[key]
	if !ok {
		return ""
	}
	pos := FindInsertPosition(stores, timestamp)
	log.Printf("stores=%v, pos=%v", stores, pos)
	if (*stores)[pos].timestamp > timestamp {
		return ""
	}
	return (*stores)[pos].value
}

func FindInsertPosition(stores *[]Node, target int) int {
	sLen := len(*stores)
	array := *stores
	left := 0
	right := sLen - 1
	for left < right {
		mid := (left + right) / 2
		if array[mid].timestamp > target {
			right = mid - 1
		}
		if array[mid].timestamp < target {
			left = mid + 1
		}
		if array[mid].timestamp == target {
			return mid
		}
	}
	if left > 0 && array[left].timestamp > target {
		left = left - 1
	}
	return left
}
