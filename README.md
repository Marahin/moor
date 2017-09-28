# Moor
  
[![pipeline status](https://git.3lab.re/marahin/moor/badges/master/pipeline.svg)](https://git.3lab.re/marahin/moor/commits/master)  

**M**edium D**oor** - **moor** is a WWW proxy that fetches data from given URL (bypassing CORS and `while(1)` anti-JSON-hijack trap). 

**It was primarily created for fetching data from Medium-based blogs**, where you do not have any kind of API, and instead you must retrieve data like Mediums' frontend does: from `%PATH%/?format=json`. There were two caveats though:  

* CORS did not allow requests from non-medium-dot-com domain,
* it begins with `])}while(1);</x>`

For that **moor** has come to life. 

## Use cases → How can I use it?

That's up to you. I use it on [my website](http://marahin.pl) to get my latest post title and link :-).

### Performance

For the maximum performance I suggest reverse proxy-ing to the application through [nginx]() or [varnish]() - basicly any reverse proxy that allows caching. 

In my case **nginx** helped me achieve following performance (sample requests):

| request | without cache | with cache                    |
|---------|---------------|-------------------------------|
| #1      | 2610ms        | 2140ms (cache initialization) |
| #2      | 1340ms        | 8.16ms                        |
| #3      | 1710ms        | 1.03ms                        |
| #4      | 872ms         | 0.59ms                        |
| #5      | 2330ms        | 0.89ms                        |
| #6      | 1080ms        | 0.93ms                        |

## Installation

This section covers only **moor** installation. **It is highly advised to put moor behind a cached reverse proxy to get maximum performance**. Sample **nginx** virtual host configuration file can be found [here](nginx.conf).

### Docker container

**Either**:

* Clone the repository and run `make docker-run`,  
* _or_ run `docker run -d -p 7999:7999 --name ${MOOR_DOCKER_NAME} -i moor-image` yourself. 

### Manual (build yourself!)

* clone the repository,

#### Without Docker

* `make build`
* run it with `./moor.bin`

#### With Docker

* `make install`
* make sure that you can see `moor` in `docker ps |grep moor` 

## Contributing

If you wish to contribute send a PR either on [GitLab (primary source tree)](http://git.3lab.re/marahin/moor)  or [on GitHub](http://github.com/marahin/moor).  
I should get notified by e-mail then, if I don't respond however, feel free to get in touch with me (you can learn how on [marahin.pl](http://marahin.pl)).