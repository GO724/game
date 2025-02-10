package game

import (
	"game/entity/item"
	"sort"
	"strings"
)

func Parce(c string) (command string, args []string) {
	for i, s := range strings.Split(c, " ") {
		if i == 0 {
			command = s
			continue
		}
		args = append(args, s)
	}
	return command, args
}

// func outputKeyMap(m map[string]struct{}) string {
// 	keys := make([]string, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys) // Сортировка ключей
// 	return strings.TrimSpace(strings.Join(keys, ", "))
// }

func outputExits(m map[string]*struct {
	Index      int
	InDoor     bool
	InDoorOpen bool
	InDoorKey  bool
}) string {
	keys := make([]string, len(m))
	for k, i := range m {
		keys[i.Index-1] = k
	}
	return strings.TrimSpace(strings.Join(keys, ", "))
}

func outputPlayerItems(m map[string]*item.Item) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Сортировка ключей
	return strings.Join(keys, ", ")
}

func outputItems(s []*item.Item) string {
	result := make([]string, 0, len(s))
	for _, v := range s {
		result = append(result, v.Name)
	}
	sort.Strings(result) // Сортировка ключей
	return strings.TrimSpace(strings.Join(result, ", "))
}
