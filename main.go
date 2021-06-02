/* Super simple program for requesting 5 bnb from venus faucet.
Venus faucet limits requests for every 15 min so use cron for example every 16min
*/
package main

import (
"bytes"
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
)

func main() {

	wallet := "" //your BSC wallet address here

	values := map[string]string{"address": wallet, "amountType": "high", "asset": "bnb"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://testnetapi.venus.io/api/faucet", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}