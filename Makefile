build:
	go build -o bin/main app/main.go

antlr:
	antlr4 -Dlanguage=Go -o app/parser Lang.g4

test:
	go run app/main.go tests/sample.c4
