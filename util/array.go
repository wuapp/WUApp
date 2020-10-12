package util

func Join(array []interface{},seprator string,prefix string, suffix string,quotes string) (v string)  {
	v = prefix
	l := len(array)
	for i,a := range array {
		v += ToString(a,quotes)
		if i == l - 1 {
			v += suffix
		} else {
			v += ","
		}
	}
	return v
}
