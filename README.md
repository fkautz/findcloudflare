A tool to help you determine what keys to rotate due to cloudbleed.

```sh
go get github.com/fkautz/findcloudflare

# check against dns only
findcloudflare list.txt

# add a set of sites to check against
findcloudflare list.txt known-sites.txt
```


list.txt should be a list of urls
```
http://www.google.com
http://www.cloudflare.com
```


known-sites should be a list of domains. Such a list can be downloaded from https://github.com/pirate/sites-using-cloudflare
```
google.com
cloudflare.com
```

A warning from sites-using-cloudflare's repo:
---

`Theoretically sites not in this list can also be affected (because an affected site could have made an API request to a non-affected one), you should probably change all your important passwords.`

