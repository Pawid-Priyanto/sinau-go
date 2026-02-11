package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	flatB, _ := json.Marshal(2.34)
	fmt.Println(string(flatB))

	strB, _ := json.Marshal("harrop")
	fmt.Println(string(strB))

	slcD := []string{"peach", "apple", "pear"}
	slcE, _ := json.Marshal(slcD)
	fmt.Println(string(slcE))

	mapD := map[string]int{"apple": 5, "lettuce": 9}
	mapE, _ := json.Marshal(mapD)
	fmt.Println(string(mapE))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "pineple", "grape"},
	}
	res1E, _ := json.Marshal(res1D)
	fmt.Println(string(res1E))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "pineple", "grape"},
	}
	res2E, _ := json.Marshal(res2D)
	fmt.Println(string(res2E))

	byt := []byte(`{"num":6.13,"strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat, "=== dat ===")

	num := dat["num"].(float64)

	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	fmt.Println(strs...)          // all strs
	fmt.Println(strs[0].(string)) // index 1
	fmt.Println(strs[1].(string)) // index 2

	str := `{"page": 1, "fruits": ["apple", "guava", "orange"]}`

	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "grape": 8}
	fmt.Println(enc.Encode(d), "tess")

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := Response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
