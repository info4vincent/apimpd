package main

// URI: https://api.kraken.com/0/public/Trades?pair=XETHZEUR&since=1439315345.846200

import (
	"flag"
	"fmt"
	"log"
	"time"

	mpd "github.com/fhs/gompd/mpd"
	gin "github.com/gin-gonic/gin"
)

func PlayMusic(host string) {
	// Connect to MPD server
	conn, err := mpd.Dial("tcp", host+":6600")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	conn.Play(-1)

	sleepTimer1 := time.NewTimer(30 * time.Minute)
	go func() {
		<-sleepTimer1.C
		fmt.Println("sleepTimer expired")
		StopMusic(host)
	}()
}

func StopMusic(host string) {
	// Connect to MPD server
	conn, err := mpd.Dial("tcp", host+":6600")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	conn.Stop()
}

func main() {
	mpdClientPtr := flag.String("mpd-client", "localhost", "Which mpd client to use?")
	flag.Parse()
	fmt.Println("mpd-client:", *mpdClientPtr)

	r := gin.Default()
	r.GET("/play", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "play started for 30 minutes",
		})
		PlayMusic(*mpdClientPtr)
	})
	r.GET("/stop", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "stop",
		})
		StopMusic(*mpdClientPtr)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
