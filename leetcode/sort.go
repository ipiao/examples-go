package leetcode

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m <= 1 {
		return mat
	}

	var stx, sty int
	for stx = 0; stx < m; stx++ { // 左下
		sort(mat, stx, 0, m, n)
	}

	for sty = 1; sty < n; sty++ { // 右上
		sort(mat, 0, sty, m, n)
	}
	return mat
}

func sort(mat [][]int, stx, sty int, maxx, maxy int) {
	var i, j, i1, j1 int
	i = stx
	j = sty
	for i < maxx-1 && j < maxy-1 {
		i1 = i + 1
		j1 = j + 1
		for i1 < maxx && j1 < maxy {
			if mat[i1][j1] < mat[i][j] {
				mat[i][j], mat[i1][j1] = mat[i1][j1], mat[i][j]
			}
			i1++
			j1++
		}
		i++
		j++
	}
}
