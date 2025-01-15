package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

func ConvertConfigMaptoMapInterface(configMapName string) (MapInterface []map[string]interface{}, err error) {

	return MapInterface, nil
}

func main() {

	// MapInterface, err := ConvertJSONToMapInterface("..\\configs\\conditionalsTest.json")
	// if err != nil {
	// 	fmt.Println("NOK")
	// }
	// fmt.Println("OK", MapInterface)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}

		time.Sleep(10 * time.Second)
	}

	//name: mongo-config
	// kubeconfig := "\\c\\Users\\victo\\.kube\\config"

	// expressionStr := "cmapo1 < campo2"

	// expression, err := govaluate.NewEvaluableExpression(expressionStr)
	// if err != nil {
	// 	fmt.Println("Erro ao criar a expressao: ", err)
	// 	return
	// }

	// fmt.Println("OK: ", expression)
}
