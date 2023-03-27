package problems

import (
	"sort"
)

func web(usernames []string, times []int, websites []string) []string {
	ans := []string{}

	user2Webs := map[string]map[int]string{}
	size := len(usernames)
	for i := 0; i < size; i++ {
		username := usernames[i]
		time := times[i]
		website := websites[i]
		if _, ok := user2Webs[username]; !ok {
			user2Webs[username] = map[int]string{}
		}
		user2Webs[username][time] = website
	}
	routeTimes := map[string]int{}
	for uName := range user2Webs {
		time2Webs := user2Webs[uName]
		sortedWebs := sortMap(time2Webs)
		// fmt.Println(sortedWebs)
		total := len(sortedWebs)
		for i := 0; i < total-2; i++ {
			for j := 1; j < total-1; j++ {
				for k := 2; k < total; k++ {
					tmpRoute := sortedWebs[i] + "-" + sortedWebs[j] + "-" + sortedWebs[k]
					routeTimes[tmpRoute]++
				}
			}
		}
	}
	// fmt.Println(routeTimes)
	maxVisitedRoute := ""
	visitTime := 0
	for route := range routeTimes {
		if routeTimes[route] > visitTime {
			visitTime = routeTimes[route]
			maxVisitedRoute = route
		}
	}
	//fmt.Println(maxVisitedRoute)
	sPos := 0
	for i := 0; i < len(maxVisitedRoute); i++ {
		if maxVisitedRoute[i] == '-' || i == len(maxVisitedRoute)-1 {
			ans = append(ans, maxVisitedRoute[sPos:i])
			sPos = i + 1
		}
	}
	return ans
}

func sortMap(toMap map[int]string) []string {
	// fmt.Println(toMap)
	sortKeys := []int{}
	for k := range toMap {
		sortKeys = append(sortKeys, k)
	}
	sort.Ints(sortKeys)
	// fmt.Println(sortKeys)
	ans := []string{}
	for _, key := range sortKeys {
		ans = append(ans, toMap[key])
	}
	// fmt.Println(ans)
	return ans
}
