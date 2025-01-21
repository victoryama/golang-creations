package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type DataHandler interface {
	ConvertJSONToMapInterface(jsonPath string) (MapInterface []map[string]interface{}, err error)
	ConvertTxtMaptoMapInterface(configMapName string) (MapInterface []map[string]interface{}, err error)
}

type DataIO struct{} //tipo concreto que implementa a interface

func (d DataIO) ConvertJSONToMapInterface(jsonPath string) (MapInterface []map[string]interface{}, err error) {

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

func (d DataIO) ConvertTxtMaptoMapInterface(configMapName string) (MapInterface []map[string]interface{}, err error) {

	return MapInterface, nil
}

func main() {
	var converter DataHandler = DataIO{}
	//Implementacao pela Interface
	MapInterface, err := converter.ConvertJSONToMapInterface("..\\configs\\conditionalsTest.json")
	if err != nil {
		fmt.Println("NOK", err)
	}
	fmt.Println("OK", MapInterface)

	//Implementacao normal
	// MapInterface, err := ConvertJSONToMapInterface("..\\configs\\conditionalsTest.json")
	// if err != nil {
	// 	fmt.Println("NOK", err)
	// }
	// fmt.Println("OK", MapInterface)

	// expressionStr := "cmapo1 < campo2"

	// expression, err := govaluate.NewEvaluableExpression(expressionStr)
	// if err != nil {
	// 	fmt.Println("Erro ao criar a expressao: ", err)
	// 	return
	// }

	// fmt.Println("OK: ", expression)
}
