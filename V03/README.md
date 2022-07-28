# HomeDepot demo v03


##  using a Echo Mux

go get github.com/labstack/echo/v4



docker build -t webhdv03 .

docker run -p 8080:8080 webhdv03


---


curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:1323/save
// => name:Joe Smith, email:joe@labstack.com


curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
// => <b>Thank you! Joe Smith</b>


curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
// => <b>Thank you! Joe Smith</b>


cd <project directory>
ls avatar.png
// => avatar.png




$ curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:8080/save







