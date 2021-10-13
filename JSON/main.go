package main

import (
	"encoding/json"
	"fmt"
)

type Inventories struct {
	Id           int      `json:"inventory_ID"`
	Name         string   `json:"name"`
	Types        string   `json:"type"`
	Tags         []string `json:"tags"`
	Purchased_at int      `json:"purchased_at"`
	Placement    struct {
		Roomid int    `json:"room_id"`
		Name   string `json:"name"`
	}
}

func main() {
	var jsonString = `
	[
	  {
		"inventory_id": 9382,
		"name": "Brown Chair",
		"type": "furniture",
		"tags": [
		  "chair",
		  "furniture",
		  "brown"
		],
		"purchased_at": 1579190471,
		"placement": {
		  "room_id": 3,
		  "name": "Meeting Room"
		}
	  },
	  {
		"inventory_id": 9380,
		"name": "Big Desk",
		"type": "furniture",
		"tags": [
		  "desk",
		  "furniture",
		  "brown"
		],
		"purchased_at": 1579190642,
		"placement": {
		  "room_id": 3,
		  "name": "Meeting Room"
		}
	  },
	  {
		"inventory_id": 2932,
		"name": "LG Monitor 50 inch",
		"type": "electronic",
		"tags": [
		  "monitor"
		],
		"purchased_at": 1579017842,
		"placement": {
		  "room_id": 3,
		  "name": "Lavender"
		}
	  },
	  {
		"inventory_id": 232,
		"name": "Sharp Pendingin Ruangan 2PK",
		"type": "electronic",
		"tags": [
		  "ac"
		],
		"purchased_at": 1578931442,
		"placement": {
		  "room_id": 5,
		  "name": "Dhanapala"
		}
	  },
	  {
		"inventory_id": 9382,
		"name": "Alat Makan",
		"type": "tableware",
		"tags": [
		  "spoon",
		  "fork",
		  "tableware"
		],
		"purchased_at": 1578672242,
		"placement": {
		  "room_id": 10,
		  "name": "Rajawali"
		}
	  }
	]
	`

	var payment []Inventories
	json.Unmarshal([]byte(jsonString), &payment)

	for i := 0; i < len(payment); i++ {
		//meeting room
		if payment[i].Placement.Name == "Meeting Room" {
			fmt.Println(payment[i])
		}
	}
	fmt.Println("")
	for i := 0; i < len(payment); i++ {
		//electronic devices
		if payment[i].Types == "electronic" {
			fmt.Println(payment[i])
		}
	}
	fmt.Println("")
	for i := 0; i < len(payment); i++ {
		//furniture
		if payment[i].Types == "furniture" {
			fmt.Println(payment[i])
		}
	}
	fmt.Println("")
	//Time Purchase
	//cant compare time, so compare unix instead
	Target := 1579107600
	for i := 0; i < len(payment); i++ {

		if payment[i].Purchased_at < Target {
			fmt.Println(payment[i])
		}
	}
	fmt.Println("")
	for i := 0; i < len(payment); i++ {
		//Brown
		for j := 0; j < len(payment[i].Tags); j++ {
			if payment[i].Tags[j] == "brown" {
				fmt.Println(payment[i])
				continue
			}
		}
	}

}
