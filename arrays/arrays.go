package arrays

func RemoveIntFromArray(numbers []int) []int {
	oddElements := 0

	for n := range numbers {
		if numbers[n]%2 != 0 {
			oddElements++
		}
	}

	arrayWithOddElements := make([]int, 0)

	for i := range numbers {
		number := numbers[i]

		if number%2 != 0 {
			arrayWithOddElements = append(arrayWithOddElements, number)
		}
	}

	return arrayWithOddElements
}

func MergeSortedArrays(a, b []int) []int {
	merged := make([]int, 0)
	lenA, lenB := len(a), len(b)
	var indexA, indexB = 0, 0

	for i := 0; i < lenA+lenB; i++ {
		if indexA >= lenA {
			merged = append(merged, b[indexB])
			indexB++
			continue
		} else if indexB >= lenB {
			merged = append(merged, a[indexA])
			indexA++
			continue
		}

		if a[indexA] < b[indexB] {
			merged = append(merged, a[indexA])
			indexA++
		} else {
			merged = append(merged, b[indexB])
			indexB++
		}
	}

	return merged
}
