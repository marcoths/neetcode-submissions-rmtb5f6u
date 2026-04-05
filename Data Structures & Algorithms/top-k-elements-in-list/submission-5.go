
type Item struct {
	val int
	freq int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	h := &MinHeap{}

	for val, freq := range count {
		heap.Push(h, Item{val, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, k)

	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Item).val
	}
	return result

}
