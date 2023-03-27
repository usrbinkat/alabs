package main

import (
	"fmt"
	"strings"
)

func main() {
	rackSkuId := "/subscriptions/a3eeb848-665a-4dbf-80a4-eb460930fb23/providers/Microsoft.NetworkCloud/rackSkus/VLab1_4_Aggregator_x70r3_9_sim"
	s := skuId(rackSkuId)
	fmt.Printf("%v", s)
}

func skuId(rackSkuId string) string {
	rackSku := strings.Split(rackSkuId, "/")[len(strings.Split(rackSkuId, "/"))-1]
	return rackSku
}
