package calculate

import "reflect"

func Add(X interface{}, Y interface{}) any {
	if reflect.TypeOf(X) != reflect.TypeOf(Y) {
		panic("The type should be the same.")
	}

	switch X.(type) {

	case int64:
		return addInt(X.(int64), Y.(int64))

	case float64:
		return addFloat(X.(float64), Y.(float64))

	case string:
		return addString(X.(string), Y.(string))

	default:
		panic("Unsupported type!")
	}
}

func addInt(x int64, y int64) int64 {
	return x + y
}

func addFloat(x float64, y float64) float64 {
	return x + y
}

func addString(x string, y string) string {
	return x + y
}
