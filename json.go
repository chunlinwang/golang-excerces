package main

import (
	"encoding/json"
	"fmt"

	"github.com/liip/sheriff"
)

type person struct {
	Name   string `json:"name" groups:"read"`
	Sex    string `json:"per_sex" groups:"read"`
	Age    int    `json:"per_age" groups:"read"`
	Groups string `json:"groups,omitempty" groups:"else"`
}

func main() {

	// bolB, _ := json.Marshal(true)
	// fmt.Println(string(bolB))

	// intB, _ := json.Marshal(1)
	// fmt.Println(string(intB))

	// fltB, _ := json.Marshal(2.34)
	// fmt.Println(string(fltB))

	// strB, _ := json.Marshal("gopher")
	// fmt.Println(string(strB))

	// slcD := []string{"apple", "peach", "pear"}
	// slcB, _ := json.Marshal(slcD)
	// fmt.Println(string(slcB))

	// mapD := map[string]int{"apple": 5, "lettuce": 7}
	// mapB, _ := json.Marshal(mapD)
	// fmt.Println(string(mapB))

	// mapDs := []map[string]int{
	// 	{"apple": 5, "lettuce": 7},
	// 	{"apple": 2, "lettuce": 9},
	// }
	// mapBs, _ := json.Marshal(mapDs)
	// fmt.Println(string(mapBs))

	// serializer
	perD := person{
		Name:   "John",
		Sex:    "Male",
		Age:    23,
		Groups: "test",
	}
	perB, _ := json.Marshal(perD)
	fmt.Println(string(perB))

	// serializer with group.
	o := &sheriff.Options{
		Groups: []string{"read"},
	}

	data, err := sheriff.Marshal(o, perD)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	//deserializer
	perByte := `{"name":"John","per_sex":"Male","per_age":23}`
	perD2 := person{}

	if err := json.Unmarshal([]byte(perByte), &perD2); err != nil {
		panic(err)
	}

	fmt.Println(perD2)

	// byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// var dat map[string]interface{}

	// if err := json.Unmarshal(byt, &dat); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dat)

	// num := dat["num"].(float64)
	// fmt.Println(num)

	// strs := dat["strs"].([]interface{})
	// str1 := strs[0].(string)
	// fmt.Println(str1)

	// str := `{"page": 1, "fruits": ["apple", "peach"]}`
	// res := response2{}
	// json.Unmarshal([]byte(str), &res)
	// fmt.Println(res)
	// fmt.Println(res.Fruits[0])

	// enc := json.NewEncoder(os.Stdout)
	// d := map[string]int{"apple": 5, "lettuce": 7}
	// enc.Encode(d)
}
