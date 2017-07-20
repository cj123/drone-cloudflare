# drone-cloudflare [![Build Status](https://drone-ci.ipsw.me/api/badges/cj123/drone-cloudflare/status.svg)](https://drone-ci.ipsw.me/cj123/drone-cloudflare)
a [drone.io](http://readme.drone.io/) deploy plugin to clear the [CloudFlare](https://cloudflare.com) cache for a domain.

### setup

follow drone's [official documentation on custom plugins](http://readme.drone.io/plugins/plugin-overview/) to install and use this plugin.

file: `.drone.yml`
```yml
pipeline:
  cloudflare:
    image: seejy/drone-cloudflare
    apikey: some-cloudflare-apikey
    email: john@smith.com
    domain: mydomain.com
```

## building

see `Makefile`
