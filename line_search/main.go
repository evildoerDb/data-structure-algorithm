package main

func search(data []int, target int) int {

	for i := 0; i < len(data); i++ {
		if target == data[i] {
			return i
		}
	}

	return -1
}

func main() {

}
