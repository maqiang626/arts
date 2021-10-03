func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if n, ok := hashTable[target-v]; ok {
			return []int{n, i}
		}
		hashTable[v] = i
	}
	return nil
}