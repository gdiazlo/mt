package main

type Digest []byte

func hash(data ...[]byte) []byte {
	var result byte
	for _, elem := range data {
		var sum byte
		for _, b := range elem {
			sum = sum ^ b
		}
		result = result ^ sum
	}

	return []byte{result}
}
