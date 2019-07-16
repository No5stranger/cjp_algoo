package sort

func MergeSort(l []int) []int {
	var l1, l2 []int
	for i := 0; i < len(l); i++ {
		if i%2 == 0 {
			l2 = append(l2, l[i])
		} else {
			l1 = append(l1, l[i])
		}
	}
	insertSort(l1)
	insertSort(l2)
	ml := merge(l1, l2)
	return ml
}

func insertSort(l []int) {
	ll := len(l)
	for i := 1; i < ll; i++ {
		value := l[i]
		j := i - 1
		for ; j >= 0; j-- {
			if l[j] > value {
				l[j+1] = l[j]
			} else {
				break
			}
		}
		l[j+1] = value
	}
}

func merge(l1, l2 []int) []int {
	var tmp []int
	var i, j int = 0, 0
	len1 := len(l1)
	len2 := len(l2)
	for i < len1 && j < len2 {
		if l1[i] < l2[j] {
			tmp = append(tmp, l1[i])
			i++
		} else {
			tmp = append(tmp, l2[j])
			j++
		}
	}
	for i < len1 {
		tmp = append(tmp, l1[i])
		i++
	}
	for j < len2 {
		tmp = append(tmp, l2[j])
		j++
	}
	return tmp
}
