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

type element struct {
	port, protocol, service string
	details string
}

func sanity(args []string) bool {
	if len(os.Args[1:]) <= 0 {
		return false;
	} 
	return true;
}


func tableHeader() {
	th := "Port | Protocol | Service"
	line := "=========================="
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

func getData(ports []string) []element {
	slices := make([]element, 0)
	for _, port := range ports {
		resp, err := soup.Get("https://www.speedguide.net/port.php?port=" + port)
		if err != nil {
			log.Fatalln(err)
		}
		html := soup.HTMLParse(resp)
		grid := html.Find("table", "class", "port").FindAll("td")
		
		p := element{}
		p.port = SpaceStringsBuilder(grid[0].Text())
		p.protocol = SpaceStringsBuilder(grid[1].Text())
		p.service = SpaceStringsBuilder(grid[2].Text())
		slices = append(slices, p)
	}
	return slices
}

func prettify(l []element) {
	for i := 0; i < len(l); i++ {
		fmt.Printf("%s %s %s\n", l[i].port, l[i].protocol, l[i].service)
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
