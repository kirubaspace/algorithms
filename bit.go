1. Count the bits set in an integer

int countSetBits(int n) {
    int count = 0;
    while (n) {
        count += n & 1;
        n >>= 1;
    }
    return count;
}

2. Reverse bits set in an integer

uint32_t reverseBits(uint32_t n) {
    uint32_t result = 0; // Initialize the result

    for (int i = 0; i < 32; i++) {
        result <<= 1;        // Left shift result by 1
        result |= n & 1;     // Set the least significant bit of result 
        n >>= 1;             // Right shift n by 1
    }

    return result;
}


3. Tell if 20th bit set is on

int main() {
    uint32_t num;
    printf("Enter a 32-bit integer: ");
    scanf("%u", &num);

    int is20thBitOn = (num >> 19) & 1 == 1;

    if (is20thBitOn) {
        printf("The 20th bit of %u is ON\n", num);
    } else {
        printf("The 20th bit of %u is OFF\n", num);
    }

    return 0;
}


4. Flip bits

int32_t swapBits(int32_t n) {
    return ~n; // Inverts all bits
}


5. Swap odd and even bits

uint32_t swapOddEvenBits(uint32_t x) {
    uint32_t even = x & 0xAAAAAAAA; // Select even bits
    uint32_t odd  = x & 0x55555555; // Select odd bits

    even >>= 1;  // Right shift even bits
    odd  <<= 1;  // Left shift odd bits

    return even | odd; // Combine
}

6. Find the single number ( Given an array of integers where every element appears twice except for one, find that single one )

int findSingle(int nums[], int numsSize) {
    int result = 0;
    for (int i = 0; i < numsSize; i++) {
        result ^= nums[i];
    }
    return result;
}

7. Check if the number is power of 2

bool isPowerOfTwo(int n) {
    return n > 0 && (n & (n - 1)) == 0;
}

8. Question: Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

int missingNumber(const std::vector<int>& nums) {
    int missing = nums.size();
    for (int i = 0; i < nums.size(); ++i) {
        missing ^= i ^ nums[i];
    }
    return missing;
}

9. Question: Calculate the Hamming distance between two integers (the number of positions at which the corresponding bits are different).

int hammingDistance(int x, int y) {
    int xorVal = x ^ y;
    int count = 0;
    while (xorVal != 0) {
        count += xorVal & 1;
        xorVal >>= 1;
    }
    return count;
}
