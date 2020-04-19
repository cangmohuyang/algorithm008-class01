var m = make(map[int]int)
func climbStairs(n int) int {
	return climb_Stairs(n)
}

func climb_Stairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	i, j := 0, 0
	if v, ok := m[n-1]; ok {
		i = v
	} else {
		i = climb_Stairs(n - 1)
		m[n-1] = i
	}
	if v, ok := m[n-2]; ok {
		j = v
	} else {
		j = climb_Stairs(n - 2)
		m[n-2] = j
	}
	return i + j
}