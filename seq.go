// Zero-Based sequence
package sequence

type seq struct {
	a, b int
}

func NewZbSeq(a, b int) *seq {
	return &seq{a, b}
}

func (s *seq) Contains(value int) bool {
	return (value-s.a)%s.b == 0
}

func (s *seq) Value(index int) int {
	return s.a + s.b*index
}

func (s *seq) Index(value int) int {
	return (value - s.a) / s.b
}

func (s *seq) ToArray(count int) []int {
	arr := []int{}
	i := 0
	for {
		v := s.Value(i)
		if v >= 0 {
			arr = append(arr, v)
		}

		if len(arr) >= count {
			break
		}

		i++
	}

	return arr
}
