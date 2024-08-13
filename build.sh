GOOS=linux GOARCH=amd64 go build -o shodan-clone

tar -czf shodan-clone.tar.gz shodan-clone resources
rm shodan-clone
