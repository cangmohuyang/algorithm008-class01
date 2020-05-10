func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gen(n, 0, 0, "", &res)
	return res
}

func gen(total int, left int, right int, ins string, res *[]string) {
	//break
	if left >= total && right >= total {
		//fmt.Println(ins)
		*res = append(*res, ins)
		return
	}
	//process
	if left < total {
		//next
		gen(total, left+1, right, ins+"(", res)
	}
	if right < left && right < total {
		//next
		gen(total, left, right+1, ins+")", res)
	}
}