# Go-Chi-Demo

Go Demo using Chi Router (Gitea)

### Setup
The following commands were used to create a PostgreSQL Instance and run the GO API Application;


``` bash
docker compose pull && docker compose up -d
docker compose run -p 8080:8080 --rm -it api /bin/bash
make run
```

### References
The following references were used;

- [https://www.youtube.com/watch?v=wpnN3RIRSxs&list=RDCMUCW5YeuERMmlnqo4oq8vwUpg&index=2](https://www.youtube.com/watch?v=wpnN3RIRSxs&list=RDCMUCW5YeuERMmlnqo4oq8vwUpg&index=2)
- [https://hub.docker.com/_/golang](https://hub.docker.com/_/golang)
- [https://www.youtube.com/watch?v=Yi9e0DAHY6Y&list=PLujhHB_uAFJws6Vv5q1KDoaQ4YcpS9UOm&index=35](https://www.youtube.com/watch?v=Yi9e0DAHY6Y&list=PLujhHB_uAFJws6Vv5q1KDoaQ4YcpS9UOm&index=35)
- [https://medium.com/@jitenderkmr/building-scalable-micro-services-with-go-golang-and-chi-framework-6db5f2f9ad28](https://medium.com/@jitenderkmr/building-scalable-micro-services-with-go-golang-and-chi-framework-6db5f2f9ad28)
