package main

import (
	"fmt"
)

func main() {

	// An exercise about Comparing Values in Operators
	visits := 15
	fmt.Println("First visit :", visits == 1)
	fmt.Println("First visit :", visits != 1)
	fmt.Println("Silver member :", visits >= 10 && visits < 21)
	fmt.Println("Gold member   :", visits > 20 && visits <= 30)
	fmt.Println("Platinum member :", visits > 30)
}

CREATE EVENT `call_partition_maintenance_all` ON SCHEDULE
EVERY 1 DAY 
STARTS TIMESTAMP(NOW() + INTERVAL 1 MINUTE)
ON COMPLETION PRESERVE ENABLE
DO CALL partition_maintenance_all('zabbix',60,400);