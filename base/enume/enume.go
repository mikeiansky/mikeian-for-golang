package main

import "fmt"

// 交易状态
type TransactionStatus string

const (
	TransactionCreated   TransactionStatus = "CREATED"
	TransactionActive    TransactionStatus = "ACTIVE"
	TransactionSuccess   TransactionStatus = "SUCCESS"
	TransactionFailed    TransactionStatus = "FAILED"
	TransactionCancelled TransactionStatus = "CANCELLED"
	TransactionExpired   TransactionStatus = "EXPIRED"
)

func main() {

	s1 := TransactionCreated
	fmt.Println(s1)

	s3 := TransactionStatus("CREATED")
	fmt.Println(s3)

	s4 := "CREATED"
	ts4 := TransactionStatus(s4)
	fmt.Println(ts4 == s3)
	fmt.Println(ts4 == s1)
	if "CREATED" == TransactionCreated {

	}

}
