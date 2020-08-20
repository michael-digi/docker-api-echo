run:
	go run test.go

get:
	curl -H "Content-Type: application/json" \
	-H "x-api-key: thisisanapikey" \
	http://localhost:8000/protected/containers

getTest:
	curl -H "Content-Type: application/json" \
	http://localhost:8000/testing
