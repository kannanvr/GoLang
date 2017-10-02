package addMap

import "fmt"

var data = make(map[string]struct{})

func AddKeyonlyInMap(key string) bool {
	//Store Data in the Map
	_, ok := data[key]
	if ok != true {
		data[key] = struct{}{}
		return true
	} else {

		fmt.Println("Already added the Data to the Map")
		return false
	}

}

func DeleteValueOnlyMap(key string) bool{

	_, ok := data[key]
	if ok == true {

		delete(data, key)

	} else {

		fmt.Println("Could not delete the Data from Map")
		return false

	}
return true
}
