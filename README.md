# BATTLESHIP GAME BASIC GO API

```
$> go run cmd/battleship_api/main.go
```

Battlefield status after `PUT` method:

```
==== BATTLEFIELD STATUS =============================================
09 [4] 19 [0] 29 [0] 39 [0] 49 [0] 59 [0] 69 [0] 79 [0] 89 [0] 99 [0] 
08 [4] 18 [0] 28 [0] 38 [0] 48 [X] 58 [0] 68 [0] 78 [0] 88 [0] 98 [0] 
07 [4] 17 [0] 27 [0] 37 [0] 47 [1] 57 [0] 67 [0] 77 [0] 87 [0] 97 [0] 
06 [0] 16 [0] 26 [0] 36 [0] 46 [1] 56 [0] 66 [0] 76 [0] 86 [0] 96 [0] 
05 [0] 15 [0] 25 [0] 35 [0] 45 [0] 55 [0] 65 [0] 75 [0] 85 [0] 95 [0] 
04 [0] 14 [0] 24 [0] 34 [0] 44 [0] 54 [0] 64 [0] 74 [0] 84 [3] 94 [0] 
03 [0] 13 [0] 23 [0] 33 [0] 43 [0] 53 [0] 63 [0] 73 [0] 83 [3] 93 [0] 
02 [0] 12 [0] 22 [0] 32 [2] 42 [0] 52 [0] 62 [0] 72 [0] 82 [3] 92 [0] 
01 [0] 11 [0] 21 [0] 31 [2] 41 [0] 51 [0] 61 [0] 71 [0] 81 [0] 91 [0] 
00 [0] 10 [0] 20 [0] 30 [2] 40 [0] 50 [0] 60 [0] 70 [0] 80 [0] 90 [0] 
=====================================================================
```


## Example endpoints

```
$> curl -H "Content-Type: application/json" -X POST -d '[[4, 8],[3, 2],[8, 4],[0,9]]' http://localhost:8080/game
```
```
$> curl -X PUT http://localhost:8080/game?x=4&y=8
```
