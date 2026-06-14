package helper

import (
	"crypto/rand"
	"strconv"
)

func RandomNumbers(length int) (int, error) {

	const numbers = "0123456789"

	// allocate buffer
	buffer := make([]byte, length)

	// fill with random bytes
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}

	// map bytes to digits
	for i := 0; i < length; i++ {
		buffer[i] = numbers[int(buffer[i])%len(numbers)]
	}

	// convert to integer
	return strconv.Atoi(string(buffer))
}