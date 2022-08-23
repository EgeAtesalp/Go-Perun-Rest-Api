package hypercore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"restapidemo/model"
	"strconv"
	"time"
)

type users struct {
	User []user `json:"users"`
}

type user struct {
	Alias string `json:"alias"`
	IP    string `json:"address"`
	Port  string `json:"port"`
}

func GetUsers(url string) (map[string]model.UserSession, error) {
	u := users{}
	umap := make(map[string]model.UserSession)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return umap, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return umap, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return umap, err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return umap, err
	}

	for _, element := range u.User {
		port, err := strconv.Atoi(element.Port)
		if err != nil {
			fmt.Println("Atoi failed")
		}

		userSession := model.UserSession{
			Alias:   element.Alias,
			Ip_Addr: element.IP,
			Port:    int16(port),
		}

		umap[element.Alias] = userSession
	}

	return umap, nil
}

func StartUserSession(userSession model.UserSession) error {
	terminal := "gnome-terminal"
	arg1 := "--"
	arg2 := "node"
	arg3 := "hyper-server.js"
	arg4 := userSession.Alias
	arg5 := userSession.Ip_Addr
	arg6 := strconv.FormatInt(int64(userSession.Port), 10)

	cmd := exec.Command(terminal, arg1, arg2, arg3, arg4, arg5, arg6)
	cmd.Start()
	time.Sleep(time.Second)
	fmt.Println("Started hyper-server")
	return nil
}

func FindUserSession(alias string) (userSession model.UserSession, err error) {
	userMap, err := GetUsers("http://localhost:3002/db")
	if err != nil {
		fmt.Println("GetUsers() failed")
	}

	userSession = userMap[alias]
	return userSession, err
}

func main() {

	aliceSession := model.UserSession{
		Alias:   "alice",
		Ip_Addr: "127.0.0.1",
		Port:    5751,
	}

	err1 := StartUserSession(aliceSession)
	if err1 != nil {
		fmt.Println("cannot start user session,", err1)
	} else {
		println("Started Alice's Session")
	}

	sess, err2 := FindUserSession("alice")
	if err2 != nil {
		fmt.Println("cannot correctly find alice,", err2)
	}

	if sess == aliceSession {
		fmt.Println("emp1 annd emp2 are equal")
	} else {
		fmt.Println("emp1 annd emp2 are not equal")
	}

	// Last test only works if Bob has started their UserSession beforehand.
	Bobsess, err3 := FindUserSession("bob")
	if err3 != nil {
		fmt.Println("cannot correctly find bob,", err2)
	} else {
		fmt.Println("Found", Bobsess.Alias)
	}

}
