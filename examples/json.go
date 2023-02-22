package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("gopher")
	fmt.Println(stringB)

	sliceA := []string{"apple", "peach", "pear"}
	sliceB, _ := json.Marshal(sliceA)
	fmt.Println(string(sliceB))

	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))

	res1A := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1A)
	fmt.Println(string(res1B))

	res2A := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2A)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
