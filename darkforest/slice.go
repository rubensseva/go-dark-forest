package darkforest

func remove[T comparable] (r T, s []T) []T {
	ns := make([]T, 0, len(s))
	for _, el := range s {
		if el == r {
			continue
		}
		ns = append(ns, el)
	}
	return ns
}
