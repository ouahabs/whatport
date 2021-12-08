/*
Copyright © 2021 Abdelouahab

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"os"
	"log"
	"strings"
	"unicode"
)

type service struct {
	name string
	details []string
}
type port struct {
	num, protocol, source string
	ser service // considering using an anonymous struct, but i don't want any spaghetti code.
}

func sanity(args []string) bool {
	if len(os.Args[1:]) <= 0 {
		return false;
	} 
	return true;
}


func tableHeader() {
	th := "Port | Protocol | Service"
	line := "========================="
	fmt.Println(th)
	fmt.Println(line)
}

func usage() {
	fmt.Println("Usage:\n  whatport [port(s)] \t (Seperate ports with a space)")
	os.Exit(1)
}
func getArgs() []string {
	return os.Args[1:]
}

func SpaceStringsBuilder(str string) string {
    var b strings.Builder
    b.Grow(len(str))
    for _, ch := range str {
        if !unicode.IsSpace(ch) {
            b.WriteRune(ch)
        }
    }
    return b.String()
}

func getData(ports []string) []port {
	slices := make([]port, 0)
	for _, num := range ports {
		resp, err := soup.Get("https://www.speedguide.net/port.php?port=" + num)
		if err != nil {
			log.Fatalln(err)
		}
		html := soup.HTMLParse(resp)
		grid := html.Find("table", "class", "port").FindAll("td")
		

		p := port{}
		p.num = SpaceStringsBuilder(grid[0].Text())
		p.protocol = SpaceStringsBuilder(grid[1].Text())
		p.ser.name = SpaceStringsBuilder(grid[2].Text())
		slices = append(slices, p)
	}
	return slices
}

func scrapeJSON() {
	/*
		We want the json file to have this structure
		{
			"port":{
				"number": {{ port_number }}
				"protocol": {{ protocol }}
				"service": {
					"name": {{ service_name }}
					"details": {{ details }}
				}
				"source": {{ source }}
			}

		}
	*/
}

func prettify(l []port) {
	for i := 0; i < len(l); i++ {
		fmt.Printf("%s %s %s\n", l[i].num, l[i].protocol, l[i].ser.name)
	}
}

func main() {
	if !sanity(getArgs()) {
		usage()
	} else {
	tableHeader()
	args := getArgs()
	data := getData(args)
	// fmt.Printf("%s %s %s\n", data[0], data[1], data[2])
	prettify(data)
	}
}
