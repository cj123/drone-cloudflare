# drone-cloudflare
a drone.io deploy plugin to clear the cloudflare cache for a domain

### setup

follow drone's [official documentation on custom plugins](http://readme.drone.io/devs/plugins/#custom-plugins:dce8ed91d073f65a191aa58c2338afcb) to whitelist the seejy/drone-cloudflare plugin

file: `.drone.yml`
```yml
deploy:
  cloudflare:
    image: seejy/drone-cloudflare
    apikey: some-cloudflare-apikey
    email: john@smith.com
    domain: mydomain.com
```

## developing

### building

see `Makefile`

### debugging

```
$ ./drone-cloudflare <<EOF
{
  "repo": {
    "clone_url": "git://github.com/drone/drone",
    "owner": "drone",
    "name": "drone",
    "full_name": "drone/drone"
  },
  "system": {
    "link_url": "https://beta.drone.io"
  },
  "build": {
    "number": 22,
    "status": "success",
    "started_at": 1421029603,
    "finished_at": 1421029813,
    "message": "Update the Readme",
    "author": "johnsmith",
    "author_email": "john.smith@gmail.com",
    "event": "push",
    "branch": "master",
    "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
    "ref": "refs/heads/master"
  },
  "workspace": {
    "root": "/drone/src",
    "path": "/drone/src/github.com/drone/drone"
  },
  "vargs": {
    "apikey": "<YOUR API KEY>",
    "email": "foo@bar.com",
    "domain": "somedomain.com"
  }
}
EOF
```
