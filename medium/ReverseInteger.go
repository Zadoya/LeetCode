// условие задачи - 

func reverse(x int) int {
	var rev int32

    for x != 0 {
        dec := (int32(x) % 10)
		new := rev * 10 + dec
        if (new - dec) / 10 != rev {
            return 0
        }
		rev = new
        x /= 10
    }
    return int(rev)
}

/* same, but interesting solution

func reverse(x int) int {
    var rev int
    
    for x != 0 {
        rev = rev * 10 + x % 10
        x /= 10
        if rev > math.MaxInt32 || rev < math.MinInt32 {
            return 0
        }
    }
    
    return rev
}
*/