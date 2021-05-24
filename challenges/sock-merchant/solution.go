package sock_merchant

func SockMerchant(n int32, ar []int32) int32 {

	var pair int32
	mp := make(map[int32]int32)

	for i := range ar {
		if val, ok := mp[ar[i]]; ok {
			val++
			mp[ar[i]] = val
			continue
		}

		mp[ar[i]] = 1
	}

	for _, val := range mp {
		pair += val / 2
	}
	return pair
}
