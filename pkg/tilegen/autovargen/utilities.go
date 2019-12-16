package autovargen

// Counts amount of bits in a number (hardcoded for 32-bit numbers)
func countBits(n uint) uint {
    n = ((0xaaaaaaaa & n) >> 1) + (0x55555555 & n)
    n = ((0xcccccccc & n) >> 2) + (0x33333333 & n)
    n = ((0xf0f0f0f0 & n) >> 4) + (0x0f0f0f0f & n)
    n = ((0xff00ff00 & n) >> 8) + (0x00ff00ff & n)
    n = ((0xffff0000 & n) >> 16) + (0x0000ffff & n)
    return n
}
