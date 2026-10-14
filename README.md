# Sample Microservices Go project

Simple microservice project that showcases:

- Routing
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Hexagonal Architecture](<https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>)

## For testing

- GUI - Use [Bruno](https://www.usebruno.com/)
- CLI - curl

```bash
curl -i http://localhost:8000/greet
curl -i -H "Content-Type: application/xml" http://localhost:8000/customers
curl -i -H "Content-Type: application/json" http://localhost:8000/customers
```
