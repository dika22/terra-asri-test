package main

import "fmt"

func main()  {
	text := "test🍔"
	fmt.Println(decryptText(text))
}

func decryptText(text string) string {
	checkIsHaveEmo(text)
	return ""
}

func checkIsHaveEmo(text string)  {
	emoticons(text)
}

func emoticons(text string) map[string]string {
	// Food:  = burger,  = fries, 🍕= pizza
	emo := map[string]string{}

	emo["burger"] = "🍟"
	emo["fries"]  = "🍟"

	return emo
}

// func regex()string{
// 	return ""
// }