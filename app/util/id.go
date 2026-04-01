package main

import (
	"fmt"
	"strings"
)

func createSql(ids []int64) string {
	//if len(ids) == 0 {
	//	return []*core.Channel{}, nil
	//}

	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
	var channelQuery = fmt.Sprintf(`SELECT 
         id, name, code, channel_type, image, currency, category, 
         status, created_at_utc, updated_at_utc
       FROM %s
       WHERE id IN (%s)`, "channels", placeholders)
	return channelQuery
}

func main() {
	//id, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//uuidString := id.String()
	//fmt.Println(uuidString)

	//uuid.NewUUID()
	ret := createSql([]int64{})
	fmt.Println(ret)
}
