package utils
import "reflect"
func StructToMap(item interface{}) map[string] interface{} {
	dic := make(map[string]interface{})
	s := reflect.ValueOf(item).Elem()
	T := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		dic[T.Field(i).Name] = f.Interface()
	}
	return dic
}