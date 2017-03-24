package file

/**
 * uniqueStringSlice is a slice wrapper that prevents duplicate values in the slice
 */
type uniqueStringSlice struct {
	s []string
}

func (uss *uniqueStringSlice) add(val string) {
	for _, has := range uss.s {
		if has == val {
			return
		}
	}
	uss.s = append(uss.s, val)
}
func (uss *uniqueStringSlice) merge(vals []string) {
	for _, val := range vals {
		uss.add(val)
	}
}
func (uss *uniqueStringSlice) slice() []string {
	return uss.s
}

