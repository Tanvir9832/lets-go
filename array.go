package main

func Array() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	var arr2 = [5]int{10, 20, 30, 40, 50}

	var arr3 = [...]int{10, 20, 30, 40, 50}

	var arr4 = [5]string{3: "Chris", 4: "Ron"}

	var arr5 = [5]int{1: 10, 3: 30}

	var arrs []int

	for i := 0; i < 10; i++ {
		arrs = append(arrs, i)
	}

	for i := 0; i < len(arrs); i++ {
		println(arrs[i])
	}

}
