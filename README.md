# BATTLESHIP GAME BASIC GO API

```
go run cmd/simple_api/main.go
```


## Example endpoints


```
curl -H "Content-Type: application/json" -X POST -d '[[4, 8],[3, 2],[8, 4],[0,9]]' http://localhost:8080/game
```
```
 curl -X PUT http://localhost:8080/game?x=4&y=8
```
