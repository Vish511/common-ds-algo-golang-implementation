package hashtable

type hashTable struct {
	size int
}

func (*hashTable) NewHashTable(size int) hashTable {
	if size == 0 {
		size = 53 //prime number as it helps in better distribution of values in the array
	}
	return hashTable{size: size}
}

// func (*hashTable) hash(key string) {
// 	var total int
// 	var
// }
