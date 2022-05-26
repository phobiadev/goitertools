package goitertools

type Num interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64	
}

func AccumulateAdd[T Num](iterable []T, initial T) []T {
	total := initial
	a := []T{}
	for _, s := range iterable {
		total += s
		a = append(a,total)
	}
	return a
}

func Accumulate[T any](iterable []T, function func(T,T) T, initial T) []T {
	 total := initial
	 a := []T{}
	 for _, s := range iterable {
		 total = function(total,s)
		 a = append(a,total)
	 }
	 return a
}

func Chain[T any](iterables ...[]T) []T {
	chained := []T{}
	for _, iterable := range iterables {
		for _, s := range iterable {
			chained = append(chained,s)
		}
	}
	return chained
}

func ChainFromIterable[T any](iterables [][]T) []T {
	chained := []T{}
	for _, iterable := range iterables {
		for _, s := range iterable {
			chained = append(chained,s)
		}
	}
	return chained
}

func Compress[T any](data []T, selectors []int) []T {
	compressed := []T{}
	for i, s := range data {
		if selectors[i] == 1 {
			compressed = append(compressed, s)
		}
	}
	return compressed
}

func DropWhile[T any](predicate func(T) bool, iterable []T) []T {
	for i, s := range iterable {
		if !predicate(s) {
			return iterable[i:]
		}
	}
	return []T{}
}

func FilterFalse[T any](predicate func(T) bool, iterable []T) []T {
	filtered := []T{}
	for _, s := range iterable {
		if !predicate(s) {
			filtered = append(filtered,s)
		}
	}
	return filtered
}

func PairWise[T any](iterable []T) [][]T {
	pairs := [][]T{}
	for i := 0; i < len(iterable)-1; i++ {
		pairs = append(pairs,iterable[i:i+2])
	}
	return pairs
}

