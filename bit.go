1. Count the bits set in an integer

func countSetBits(n int) int {
    count := 0
    for n > 0 {
        count += n & 1
        n >>= 1
    }
    return count
}

2. Reverse bits set in an integer

func reverseBits(n uint32) uint32 {
    var result uint32  // This variable will store the reversed bits
    for i := 0; i < 32; i++ {
        result <<= 1          // Left shift the result by 1
        result |= n & 1       // Set the least significant bit of result to the current bit of n
        n >>= 1               // Right shift n by 1
    }
    return result
}


3. Tell if 20th bit set is on

func main() {
    var num uint32
    fmt.Print("Enter a 32-bit integer: ")
    fmt.Scan(&num)

    is20thBitOn := (num >> 19) & 1 == 1

    if is20thBitOn {
        fmt.Printf("The 20th bit of %d is ON\n", num)
    } else {
        fmt.Printf("The 20th bit of %d is OFF\n", num)
    }
}


4. Flip bits

int32_t swapBits(int32_t n) {
    return ~n; // Inverts all bits
}


5. Swap odd and even bits

func swapOddEvenBits(x uint32) uint32 {
    even := x & 0xAAAAAAAA // Select even bits
    odd := x & 0x55555555  // Select odd bits
    even >>= 1             // Right shift even bits
    odd <<= 1              // Left shift odd bits
    return even | odd      // Combine
}

6. Find the single number ( Given an array of integers where every element appears twice except for one, find that single one )

func findSingle(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}

7. Check if the number is power of 2

func isPowerOfTwo(n int) bool {
    return n > 0 && (n&(n-1)) == 0
}

8. Question: Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

func missingNumber(nums []int) int {
    missing := len(nums)
    for i, num := range nums {
        missing ^= i ^ num
    }
    return missing
}

9. Question: Calculate the Hamming distance between two integers (the number of positions at which the corresponding bits are different).

func hammingDistance(x, y int) int {
    xor := x ^ y
    count := 0
    for xor != 0 {
        count += xor & 1
        xor >>= 1
    }
    return count
}
