package main

import (

 "net/http"

	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)


// BasicFeedPost will show you about
// how to post a feed to your timeline
func main() {


  http.HandleFunc("/", homeHandler)

  http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  // initalize facebook-go-sdk
  fb := facebook.New()

  // set your access token
  // NOTES: Please exchange with your access token
  fb.SetAccessToken(`EAAJl27dqpTwBAMJgujRfAU3d4HcqNsCxG1HfeMBGidS8sM3VTLQLKNJZCfRZC6l24CrO0fpPiLHJue7EJNjiBSaB4fMVsGYJad7sb0nDd9Rh0EvAh6KWVyGG3LwZAXrkLwm2zPi6kEstrrjPvvA1SLnkoSbPTgARL5bIoS0TT7oWgW3GWQg6ZCTErUMI9JkajPw6WwZAWyAZDZD`)

  // submit your feed
	data, err := fb.API(`/me/feed`).Messages(`ADVERTENCIA: Se ha detectado un incendio en la casa 1.1`).Post()
  if err != nil {
    panic(err)
  }

  fmt.Println(`
    ## SAMPLE - POST A FEED
  `)

  fmt.Println(data)

  // if you want to share a link you can use the `Link` method
  // e.g. fb.API(`/me/feed`).Messages(`coba yo`).Link(`http://madebyais.com`).Post()
}
