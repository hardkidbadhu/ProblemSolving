package sock_merchant

import (
	"log"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	t.Run("TestSockMerchantWithInput1", TestSockMerchantWithInput1)
	t.Run("TestSockMerchantWithInput2", TestSockMerchantWithInput2)
}

func TestSockMerchantWithInput1(t *testing.T) {
	pairs := SockMerchant(9, []int32{10,20, 20, 10, 10, 30, 50, 10, 20})
	if pairs != 3 {
		log.Println("expected pair is 3")
		t.Fail()
	}
}

func TestSockMerchantWithInput2(t *testing.T) {
	pairs := SockMerchant(5, []int32{10, 10, 10, 10, 10})
	if pairs != 2 {
		log.Println("expected pair is 3")
		t.Fail()
	}
}
