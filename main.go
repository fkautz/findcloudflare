package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {

	reader, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	known_hosts := make(map[string]bool)
	if len(os.Args) == 3 {
		log.Println("Loading known hosts list...")
		full_list, err := os.Open(os.Args[2])
		defer full_list.Close()
		if err != nil {
			log.Fatalln(err)
		}
		scanner := bufio.NewScanner(full_list)
		for scanner.Scan() {
			hostname := strings.TrimSpace(scanner.Text())
			if hostname != "" {
				known_hosts[hostname] = true
			}
		}
		log.Println("loaded")
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		parsedUrl, err := url.Parse(scanner.Text())
		if err != nil {
			log.Println("Could not parse:", err)
			continue
		}
		hostname := parsedUrl.Hostname()
		if strings.TrimSpace(hostname) == "" {
			continue
		}

		split_hostname := strings.Split(hostname, ".")

		found := false
		for len(split_hostname) > 0 {
			joined_hostname := strings.Join(split_hostname, ".")
			if _, ok := known_hosts[joined_hostname]; ok == true {
				fmt.Println("Found (cache):\t", joined_hostname)
				found = true
				break
			}
			split_hostname = split_hostname[1:]
		}
		if found == true {
			continue
		}

		cname, err := net.LookupCNAME(hostname)
		if err != nil {
			log.Println("Unable to look up cname:", hostname)
			continue
		}
		if strings.Contains(cname, "cloudflare") {
			fmt.Println("Found (dns):\t", hostname)
		}
	}
}
