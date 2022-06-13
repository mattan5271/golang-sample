package main

func PrintSlice[T any](s []T) {
	for _, v := range s {
		println(v)
	}
}

func main() {
	// controllers.StartEchoServer()

	PrintSlice[int]([]int{1, 2, 3})
	PrintSlice[string]([]string{"a", "b", "c"})
}
