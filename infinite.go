package goitertools

func Count(start,step int) <-chan int {
	chn := make(chan int)
	n := start
	go func() {
		for true {
			chn <- n
			n += step
		}
		close(chn)
	}()
	return chn
}

func Cycle[T any](iterable []T) <-chan T {
	chn := make(chan T)
	go func() {
		n := 0;
		for true {
			if n == len(iterable) {
				n = 0
			}
			chn <- iterable[n]
			n++
		}
		close(chn)
	}()
	return chn
}

func Repeat[T any](item T, n int) []T {
	a := []T{}
	for i := 0; i < n; i++ {
		a = append(a,item)
	}
	return a
}



