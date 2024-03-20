package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	Name string
	Type string
	Path []string
}

func SetData(field *Field, obj map[string]any, data []any) error {
	if len(data) == 0 {
		return nil
	}
	current := obj
	field.InitFieldPath()
	for idx := 0; idx < len(field.Path)-1; idx++ {
		if value, ok := current[field.Path[idx]]; ok {
			if reflect.TypeOf(value).Kind() == reflect.Map {
				current = value.(map[string]any)
			} else {
				return fmt.Errorf("current path is not a map: %s", field.Path[idx])
			}
		} else {
			tmp := make(map[string]interface{})
			current[field.Path[idx]] = tmp
			current = tmp
		}
	}
	if len(data) == 1 {
		current[field.Path[len(field.Path)-1]] = data[0]
	} else {
		current[field.Path[len(field.Path)-1]] = data
	}
	return nil
}

func (field *Field) InitFieldPath() {
	if len(field.Path) == 0 {
		field.Path = strings.Split(field.Name, ".")
	}
}

func recursiveDelete(m map[string]interface{}, keyToDelete string) {
	for k, v := range m {
		if k == keyToDelete {
			delete(m, k)
		}
		if nestedMap, ok := v.(map[string]interface{}); ok {
			recursiveDelete(nestedMap, keyToDelete)
		}
	}
}

func deletemap() {
	a := `
		{
			"a": {
				"b": 1,
				"c": 2
			}
		}
	`

	var aaa map[string]any
	err := json.Unmarshal([]byte(a), &aaa)
	if err != nil {
		panic(err)
	}

	fmt.Println("1", aaa)

	recursiveDelete(aaa, "a.c")

	fmt.Println("2", aaa)
}

func pickData(origin map[string]any, removedFields []string) {
	for _, fieldName := range removedFields {
		deleteNestedField(origin, fieldName)
	}
}

func deleteNestedField(origin map[string]any, field string) {
	keys := strings.Split(field, ".")
	lastKey := keys[len(keys)-1]
	parent := origin

	for _, key := range keys[:len(keys)-1] {
		switch nested := parent[key].(type) {
		case map[string]any:
			parent = nested
		case []any:
			if len(nested) == 0 {
				return
			}

			switch nested[0].(type) {
			case map[string]any:
				for _, v := range nested {
					deleteNestedField(v.(map[string]any), lastKey)
				}

			case []any:
				return
			default:
				return
			}
		default:
			// 如果无法继续访问下一层，直接返回
			return
		}
	}

	delete(parent, lastKey)
}

func main() {
	// 初始的JSON字符串
	var origin map[string]any
	jsonData := `{
		"name": ["Alice", "Emma"], 
		"age": 14, 
		"slice": [
			{"x": 1, "y": 2}, 
			{"x": 3, "z": 4},
			{"x": 6, "z": 7}
		],
		"a": {
			"b": {
				"c": {
					"d": "yayayayaya",
					"e": 22222
				}
			}
		}
		}`
	json.Unmarshal([]byte(jsonData), &origin)

	pickData(origin, []string{"slice.z", "a.b.c.d", "slice.x", "name"})

	res, _ := json.Marshal(origin)

	fmt.Println(string(res))

	// field := &Field{
	// 	Name: "a.b.c.d",
	// 	Type: "keyword",
	// }
	// obj := map[string]any{}
	// SetData(field, obj, []any{"d1", "d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "d10", "d11"})
	// r, _ := json.Marshal(obj)
	// fmt.Println(string(r))

	// m, err := sonic.GetFromString(jsonData)
	// mm, err := m.Raw()

	// fmt.Println(mm, err)

	// n := m.GetByPath("slice", 0, "x")
	// fmt.Println(n)

	// nnn, err := n.Raw()

	// fmt.Println(nnn, err)

	// jsonObj := ast.NewObject([]ast.Pair{})
	// exist, err := jsonObj.Set("a.b.c", ast.NewAny("Hello World"))

	// fmt.Println("eweeeee", exist, err)

	// fmt.Println(jsonObj)
}
