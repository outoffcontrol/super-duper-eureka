package main

import (
	"fmt"
  "gopkg.in/ahmdrz/goinsta.v2"
  "github.com/stianeikeland/go-rpio"
  "os"
  //"io/ioutil"
  //"encoding/json"
  "time"
)


func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {  

  if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

  insta := goinsta.New(os.Getenv("LOGIN"), os.Getenv("PASSWORD"))
    err1 := insta.Login()
    check(err1)

  insta.Export(".goinsta")

  

  insta.Inbox.Sync()
  temp := insta.Inbox.Conversations[0].Items[0].Text


  //out, err := json.Marshal(temp)

  //err1 := ioutil.WriteFile(".dat1", out, 0644)
  //check(err1)

  fmt.Printf("Start! %s", temp)
  pin := rpio.Pin(24)
  pin.Mode(rpio.Output)
  pinNew := rpio.Pin(23)
  pinNew.Mode(rpio.Output)
  pin.Low()
  pinNew.Low()

  for true {

		pin.High()

    insta.Inbox.Sync()
    msg := insta.Inbox.Conversations[0].Items[0].Text

    fmt.Printf("!!!!!!!!!!!!!! %s", msg)

    if msg != temp {
      fmt.Printf("New Msg!")
	  pinNew.High()
      break 
    }

    time.Sleep(time.Second * 1)
	pin.Low()
	time.Sleep(time.Second * 1)

  }


//fmt.Printf("%+v\n", temp)

}