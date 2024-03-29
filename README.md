Whatport is an open source tool that scrapes port information from [SpeedGuide's Port Database](https://www.speedguide.net/ports.php)

## Usage
```
whatport [port(s)] (Seperate ports with a space)
```

## Installation
(Uses Go 1.17)
This package uses the web scraper [soup](https://github.com/anaskhan96/soup), so you will want that package installed beforehand.
```shell
$ git clone https://github.com/ouahabs/whatport.git
$ go build -o whatport
```

## To-do
* Add options.
	* -d/--detailed, detailed output.
	* -n, number of potential services per port.
	* -c found CVEs
	* --offline for offline mode (generates json data for offline use)
* Play around with spf13/Cobra for cli commands.
* trojan filtering.
* Add port seperation with commas.


## Side Notes
This tool serves as training for my Golang skills, this is my first ever official open source Go package, so please feel free to open an issue or a pull request for features to be added and whatnot.