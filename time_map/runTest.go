package time_map

import "strconv"

func runTest(commands []string, payloads [][]string) []string {
	cLen := len(commands)
	timeMap := Constructor()
	result := []string{"null"}
	for idx := 1; idx < cLen; idx++ {
		command := commands[idx]
		payload := payloads[idx]
		timestamp, _ := strconv.Atoi(payload[len(payload)-1])
		switch command {
		case "set":
			timeMap.Set(payload[0], payload[1], timestamp)
			result = append(result, "null")
		case "get":
			value := timeMap.Get(payload[0], timestamp)
			result = append(result, value)
		}
	}
	return result
}
