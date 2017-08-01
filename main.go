package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

const url = "http://www.mydns.jp/login.html"

type Auth struct {
	ID       string
	Password string
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "useragent,ua",
					Value: "tickadns",
					Usage: "user-agent name",
				},
				cli.StringSliceFlag{
					Name:  "auth",
					Usage: "auth information list id:password",
				},
				cli.DurationFlag{
					Name:  "interval,i",
					Value: time.Minute * 60,
					Usage: "tick-time interval",
				},
			},
			Action: func(clc *cli.Context) error {
				hcl := http.DefaultClient
				dur := clc.Duration("interval")
				if dur < time.Minute*1 {
					return errors.New("too short interval")
				}
				auths := make([]Auth, 0)
				for _, v := range clc.StringSlice("auth") {
					vs := strings.Split(v, ":")
					if len(vs) != 2 {
						continue
					}
					a := Auth{
						ID:       strings.TrimSpace(vs[0]),
						Password: strings.TrimSpace(vs[1]),
					}
					auths = append(auths, a)
					fmt.Println(fmt.Sprintf("registerd: %s == %s ", a.ID, a.Password))
				}

				ticker := time.Tick(dur)
				for {
					select {
					case <-ticker:
						for i, auth := range auths {
							req, err := http.NewRequest("GET", url, nil)
							if err != nil {
								fmt.Println(err)
								continue
							}
							req.Header.Set("User-Agent", clc.String("useragent"))
							req.SetBasicAuth(auth.ID, auth.Password)
							res, err := hcl.Do(req)
							if err != nil {
								fmt.Println(err)
								continue
							}
							body, err := ioutil.ReadAll(res.Body)
							if err != nil {
								fmt.Println(err)
								continue
							}
							defer res.Body.Close()
							if strings.Contains(string(body), "notify OK") {
								fmt.Println(fmt.Sprintf("[%v] IP notify OK: %s", i, auth.ID))
								continue
							}
							fmt.Println(fmt.Sprintf("[%v] IP notify NG: %s", i, auth.ID))
						}
					}
				}
			},
		},
	}
	app.Run(os.Args)
}
