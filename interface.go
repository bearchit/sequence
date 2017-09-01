package sequence

type ISequence interface {
	Contains(value int) bool

	Value(index int) int
	Index(value int) int

	ToArray(count int) []int
}
