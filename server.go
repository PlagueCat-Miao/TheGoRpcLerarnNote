package main
import "github.com/koding/kite"

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.Port = 6000
	k.Run()
}