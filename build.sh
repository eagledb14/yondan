GOOS=linux GOARCH=amd64 go build -o yondan

tar -czf yondan.tar.gz yondan resources
rm yondan
