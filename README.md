# golang_timeMap

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the `TimeMap` class:

- `TimeMap()` Initializes the object of the data structure.
- `void set(String key, String value, int timestamp)` Stores the key `key` with the value `value` at the given time `timestamp`.
- `String get(String key, int timestamp)` Returns a value such that `set` was called previously, with `timestamp_prev <= timestamp`. If there are multiple such values, it returns the value associated with the largest `timestamp_prev`. If there are no values, it returns `""`.

## Examples

**Example 1:**

```
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"

```

**Constraints:**

- `1 <= key.length, value.length <= 100`
- `key` and `value` consist of lowercase English letters and digits.
- `1 <= timestamp <= $10^7$`
- All the timestamps `timestamp` of `set` are strictly increasing.
- At most $`2*10^5`$ calls will be made to `set` and `get`.

## 解析

1. 題目需要設計一個資料結構讓 key  對應到多個 {value, timestamp} 值
2. 並且需要把每個 {value, timestamp} 根據 timestamp 大小做 increasing 放置
3. 每次要放入與存取都要先搜尋該 key 的位置 所以是 O(logN)

## 程式碼

```go
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
}

func (this *TimeMap) Get(key string, timestamp int) string {
	stores, ok := this.Stores[key]
	if !ok {
		return ""
	}
	pos := FindInsertPosition(stores, timestamp)
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
```

## 困難點

1. 必須設計一個資料結構來保存 key 與 value 以及 timestamp的對應
2. 根據 timestamp 做搜詢必須要能夠使用 binary search

## Solve Point

- [x]  Understand what problem would to solve
- [x]  Design a Data structure to solve this problem
- [x]  Analysis Complexity