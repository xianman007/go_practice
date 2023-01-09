package dao

func Removebucket(x string) string {
	var pos int = 0
	for ; pos < len(x); pos++ {
		if x[pos] == '(' {
			break
		}
	}
	// ix, _ := strconv.Atoi(string([]byte(x)[:pos]))
	return string([]byte(x)[:pos])
}
