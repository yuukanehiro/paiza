package main

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
