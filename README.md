# Sample GO Module based web app architecture

This is still under development.
But it is a Simple Web App which can use all basic need of a web application e.g. 
- Css/js from CDN 
- Template inherit 
- Module based architecture 

## Docker Hints
#### CDN Setup example:
`docker run -itd --name gowikicdn -p 5003:80 -v C:/Users/sabbi/Documents/CodeZone/GoYa/GoWebAppSample/public:/usr/share/nginx/html nginx`
#### Wiki WebServer example:
 - `docker run -itd --name gowiki -v C:/Users/sabbi/Documents/CodeZone/GoYa/GoWebAppSample:/go/app -p 8001:8001 golang`
 - `docker exec -it gowiki bash`
 - `cd /go/app && go build wiki.go && ./wiki`