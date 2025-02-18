# 💸 richer-go (fork of rich-go)
An implementation of Discord's rich presence in Golang for Linux, macOS and Windows

This is a forked version. While developing [GeoGuessrRPC](https://github.com/RafaeloxMC/GeoGuessrRPC), I encountered a [bug](https://github.com/hugolgst/rich-go/issues/42) which has been reported for over a year and has not been fixed. This is the reason, why I forked this repository, so I can use it in my project without hoping that this bug might get fixed one day.

## 📲 Installation

Install `github.com/RafaeloxMC/richer-go`:

```
go get github.com/RafaeloxMC/richer-go
```

## 📕 Usage

First of all import richer-go
```golang
import "github.com/RafaeloxMC/richer-go/client"
```

Then login by sending the first handshake
```golang
err := client.Login("DISCORD_APPLICATION_ID")
if err != nil {
	panic(err)
}
```
You can find the application ID on your [Discord Developers Dashboard](https://discord.com/developers/applications). Select your application and copy the application ID on the page "General Information".

Finally set the Rich Presence activity.
```golang
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
	},
})

if err != nil {
	panic(err)
}
```
You can add the large and small images to your application on the [developer dashboard](https://discord.com/developers/applications). Select your application, click on "Rich Presence". You can then upload your images. The name you enter is important, it has to be the same in the code as on the dashboard.

To prevent timeouts you can logout the user at the end.
```golang
client.Logout()
```

You can find more details in the [example](https://github.com/RafaeloxMC/richer-go/blob/master/example/main.go)

## 🤝 Contributing

1. Fork the repository (https://github.com/RafaeloxMC/richer-go/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## 👨‍💻 Contributors

- [RafaeloxMC](https://github.com/RafaeloxMC) - forker, fork maintainer
- [hugolgst](https://github.com/hugolgst) - original creator, maintainer of original version
- [donovansolms](https://github.com/donovansolms) - contributor
- [heroslender](https://github.com/heroslender) - contributor
