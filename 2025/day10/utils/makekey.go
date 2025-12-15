package utils
func MakeKey(index int, curr []int) uint64 {
    var key uint64
    key = uint64(index)
    for _, v := range curr {
        key = key<<4 | uint64(v) // assumes v <= 15
    }
    return key
}
