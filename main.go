package main

import (
	"fmt"
	cloudwatchtestimport "github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	test := cloudwatchtestimport.DeleteAlarmsInput{AlarmNames: nil}

	fmt.Println(test)
}
