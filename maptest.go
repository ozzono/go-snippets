package main

import "fmt"

func main() {
	for i := 0; i < 17; i++ {
		if acErr := acErrList(i); acErr != nil {
			fmt.Printf("acErr[\"errID\"] ------------> %d\n", acErr["errID"])
			fmt.Printf("acErr[\"errorCode\"] --------> %s\n", acErr["errorCode"])
			fmt.Printf("acErr[\"errorDescription\"] -> %s\n", acErr["errorDescription"])
			fmt.Printf("\n")
		}
		fmt.Printf("%d:ok\n", i)
	}
}

func acErrList(errID int) map[string]interface{} {
	switch errID {
	case 12:
		return map[string]interface{}{
			"errID":            errID,
			"errorCode":        "ERROR_CAPTCHA_UNSOLVABLE",
			"errorDescription": "Captcha could not be solved by 5 different workers"}
	case 16:
		return map[string]interface{}{
			"errID":            errID,
			"errorCode":        "ERROR_NO_SUCH_CAPCHA_ID",
			"errorDescription": "Task you are requesting does not exist in your current task list or has been expired. Tasks not found in active tasks"}
	default:
		return nil
	}
}
