package base

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	n      = 100000
	values = makeRndValues(n)
	keys   = makeKeys(n)

	slice1   = prepSlice(1)
	slice5   = prepSlice(5)
	slice10  = prepSlice(10)
	slice25  = prepSlice(25)
	slice50  = prepSlice(50)
	slice100 = prepSlice(100)

	sliceKeyVal1   = prepSliceKeyVal(1)
	sliceKeyVal5   = prepSliceKeyVal(5)
	sliceKeyVal10  = prepSliceKeyVal(10)
	sliceKeyVal25  = prepSliceKeyVal(25)
	sliceKeyVal50  = prepSliceKeyVal(50)
	sliceKeyVal100 = prepSliceKeyVal(100)

	slicei1, sliceiv1     = prepSliceI(1)
	slicei5, sliceiv5     = prepSliceI(5)
	slicei10, sliceiv10   = prepSliceI(10)
	slicei25, sliceiv25   = prepSliceI(25)
	slicei50, sliceiv50   = prepSliceI(50)
	slicei100, sliceiv100 = prepSliceI(100)

	map1   = prepMap(1)
	map5   = prepMap(5)
	map10  = prepMap(10)
	map25  = prepMap(25)
	map50  = prepMap(50)
	map100 = prepMap(100)
)

func TestCap(t *testing.T) {
	var a []int
	for i := 0; i < 25; i++ {
		t.Log(i, cap(a))
		a = append(a, i)
	}
}

func TestCap2(t *testing.T) {
	a := make([]int, 0, 16)
	for i := 0; i < 25; i++ {
		t.Log(i, cap(a))
		a = append(a, i)
	}
}

func makeRndValues(n int) []string {
	r := []string{}
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%d", rand.Uint32()))
	}
	return r
}

func makeKeys(n int) []int {
	r := []int{}
	for i := 0; i < n; i++ {
		r = append(r, int(rand.Int31()))
	}
	return r
}

func BenchmarkSlice1(b *testing.B) {
	benchSlice(slice1, b)
}

func BenchmarkSlice5(b *testing.B) {
	benchSlice(slice5, b)
}

func BenchmarkSlice10(b *testing.B) {
	benchSlice(slice10, b)
}

func BenchmarkSlice25(b *testing.B) {
	benchSlice(slice25, b)
}

func BenchmarkSlice50(b *testing.B) {
	benchSlice(slice50, b)
}

func BenchmarkSlice100(b *testing.B) {
	benchSlice(slice100, b)
}

func BenchmarkSliceI1(b *testing.B) {
	benchSliceI(slicei1, sliceiv1, b)
}

func BenchmarkSliceI5(b *testing.B) {
	benchSliceI(slicei5, sliceiv5, b)
}

func BenchmarkSliceI10(b *testing.B) {
	benchSliceI(slicei10, sliceiv10, b)
}

func BenchmarkSliceI25(b *testing.B) {
	benchSliceI(slicei25, sliceiv25, b)
}

func BenchmarkSliceI50(b *testing.B) {
	benchSliceI(slicei50, sliceiv50, b)
}

func BenchmarkSliceI100(b *testing.B) {
	benchSliceI(slicei100, sliceiv100, b)
}

func BenchmarkSliceICall1(b *testing.B) {
	benchSliceICall(sliceKeyVal1, b)
}

func BenchmarkSliceICall5(b *testing.B) {
	benchSliceICall(sliceKeyVal5, b)
}

func BenchmarkSliceICall10(b *testing.B) {
	benchSliceICall(sliceKeyVal10, b)
}

func BenchmarkSliceICall25(b *testing.B) {
	benchSliceICall(sliceKeyVal25, b)
}

func BenchmarkSliceICall50(b *testing.B) {
	benchSliceICall(sliceKeyVal50, b)
}

func BenchmarkSliceICall100(b *testing.B) {
	benchSliceICall(sliceKeyVal100, b)
}

func findICall(s []KeyVal, key int) (KeyVal, bool) {
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v.Key() == key {
			//fmt.Println(key, i)
			return v, true
		}
	}

	return nil, false
}

func benchSlice(s []val, b *testing.B) {
	found := 0

	for i := 0; i < b.N; i++ {
		slen := len(s)
		key := keys[i%slen]
		for i := 0; i < slen; i++ {
			v := s[i]
			if v.k == key {
				found++
				break
			}
		}
	}
}

func benchSliceI(k []int, v []string, b *testing.B) {
	found := 0

	for i := 0; i < b.N; i++ {
		key := keys[i%len(v)]
		for i := 0; i < len(v); i++ {
			if k[i] == key {
				found++
				break
			}
		}
	}
}

func benchSliceICall(s []KeyVal, b *testing.B) {
	found := 0

	for i := 0; i < b.N; i++ {
		key := keys[i%len(s)]
		if _, ok := findICall(s, key); ok {
			found++
		}
	}

	//b.Log(found)
}

type val struct {
	k int
	v string
}

type KeyVal interface {
	Key() int
}

func (v val) Key() int {
	return v.k
}

func (v val) String() string {
	return v.v
}

func prepSlice(size int) []val {
	s := []val{}
	for i := 0; i < size; i++ {
		s = append(s, val{keys[i], values[i]})
	}
	return s
}

func prepSliceI(size int) ([]int, []string) {
	k := []int{}
	v := []string{}
	for i := 0; i < size; i++ {
		k = append(k, keys[i])
		v = append(v, values[i])
	}
	return k, v
}

func prepSliceKeyVal(size int) []KeyVal {
	s := []KeyVal{}
	for i := 0; i < size; i++ {
		s = append(s, val{keys[i], values[i]})
	}
	return s
}

func BenchmarkMap1(b *testing.B) {
	benchMap(map1, b)
}

func BenchmarkMap5(b *testing.B) {
	benchMap(map5, b)
}

func BenchmarkMap10(b *testing.B) {
	benchMap(map10, b)
}

func BenchmarkMap25(b *testing.B) {
	benchMap(map25, b)
}

func BenchmarkMap50(b *testing.B) {
	benchMap(map50, b)
}

func BenchmarkMap100(b *testing.B) {
	benchMap(map100, b)
}

func benchMap(m map[int]string, b *testing.B) {
	found := 0

	for i := 0; i < b.N; i++ {
		key := keys[i%len(m)]
		if _, ok := m[key]; ok {
			found++
		}
	}

	//b.Log(found)
}

func prepMap(n int) map[int]string {
	m := map[int]string{}
	for i := 0; i < n; i++ {
		m[keys[i]] = values[i]
	}
	return m
}
