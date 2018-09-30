package slices

type Slice []interface{}

func (s *Slice) Concat(slices ...Slice) Slice {
	slice := []interface{}(*s);
	for _, s := range slices {
		slice = append(slice, s...)
	}
	return slice
}

func (s *Slice) Includes(item interface{}) bool {
	slice := []interface{}(*s);
	for _, i := range slice {
			if i == item {
					return true
			}
	}
	return false
}
