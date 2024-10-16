### Write Dockerfile & Docker-compose file: Create 2 image

- Base from Golang image
- Base
- Golang app: 
    - use [Gin-gonic](https://github.com/gin-gonic/gin) to setup http server port 5000: method GET http://ip:5000/
    - each time call to this url: `countNumber += 1`
    - `countNumber` is save to mongoDB.
    - method GET http://ip:5000/count for checking `countNumber`.
- MongoDB: exposed at port 4000

** Note: push include your app!!!