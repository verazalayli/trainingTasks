package task1_parseLog

import (
	"fmt"
	"os"
	"strings"
)

type UsersWithoutLogin struct {
	name   string
	logOut bool
}

func task1() {
	getData("server.log")
}

func getData(s string) {
	data, err := os.ReadFile(s)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	var logins int
	var logouts int
	var slice map[string]bool
	logins, logouts, slice = checkData(data)
	for k, v := range slice {
		fmt.Println(k, v)
	}
	fmt.Println("logins: ", logins)
	fmt.Println("logouts: ", logouts)

}

func checkData(b []byte) (int, int, map[string]bool) {
	str := string(b)
	words := strings.Split(str, "\n")
	users := make(map[string]bool)
	var userId string
	var loginCount int
	var logoutCount int
	for _, word := range words {
		startInd := strings.LastIndex(word, "user_id=")
		if startInd == -1 {
			continue // user_id не найден
		}
		endInd := strings.IndexAny(word[startInd:], " \n")
		userId = word[startInd : startInd+endInd]
		if strings.Contains(word, "login") {
			loginCount++
			users[userId] = false
		}
		if strings.Contains(word, "logout") {
			logoutCount++
			users[userId] = true
		}
	}
	return loginCount, logoutCount, users
}
