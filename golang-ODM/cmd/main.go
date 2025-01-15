package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// type Conditionals struct {
// 	cond_1 bool `json:"condition_1"`
// 	cond_2 bool `json:"condition_2"`
// 	cond_3 bool `json:"condition_3"`
// 	cond_4 bool `json:"condition_4"`
// 	cond_5 bool `json:"condition_5"`
// }

func ConvertJSONToMapInterface(jsonPath string) (MapInterface []map[string]interface{}, err error) {

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("Erro ao ler Json", err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Erro ao ler os bytes do Json", err)
		return nil, err
	}

	err = json.Unmarshal(byteValue, &MapInterface)
	if err != nil {
		fmt.Println("Erro ao converter json", err)
		return nil, err
	}

	fmt.Println(len(MapInterface))
	fmt.Println(MapInterface)

	return MapInterface, nil

}

func main() {

	MapInterface, err := ConvertJSONToMapInterface("..\\configs\\conditionalsTest.json")
	if err != nil {
		fmt.Println("NOK")
	}
	fmt.Println("OK", MapInterface)

	// expressionStr := "cmapo1 < campo2"

	// expression, err := govaluate.NewEvaluableExpression(expressionStr)
	// if err != nil {
	// 	fmt.Println("Erro ao criar a expressao: ", err)
	// 	return
	// }

	// fmt.Println("OK: ", expression)
}
