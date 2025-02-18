package main

import (
	"fmt"
	"time"

	"github.com/RafaeloxMC/richer-go/client"
)

func main() {
	err := client.Login("DISCORD_APP_ID")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "Playing",
		Details:    "RPC is made using richer-go",
		LargeImage: "largeimageid",
		LargeText:  "This is the large image",
		SmallImage: "smallimageid",
		SmallText:  "And this is the small image",
		Party: &client.Party{
			ID:         "-1",
			Players:    15,
			MaxPlayers: 24,
		},
		Timestamps: &client.Timestamps{
			Start: time.Now(),
		},	Start: &now,
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/hugolgst/rich-go",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 10)

	// Logout - this will stop showing the RPC on the users profile
	client.Logout()
}
