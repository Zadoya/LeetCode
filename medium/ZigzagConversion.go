// условие задачи - https://leetcode.com/problems/zigzag-conversion/description/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	new := make([]byte, 0, len(s))

	for i := 1; i <= numRows; i++ {
		for j, step, fwd := i - 1, 0, true; j < len(s); j, fwd = j + step, !fwd {
			if fwd {
				step = (numRows - i) * 2
			} else {
				step = (i - 1) * 2
			}
            if step != 0 {
                new = append(new, s[j])
            }
		}
	}
	return string(new)