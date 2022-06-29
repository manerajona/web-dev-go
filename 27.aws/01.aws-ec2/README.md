# Deploying our session example

1. Build hello world

- GOOS=linux GOARCH=amd64 go build -o sessionapp

2. Copy your binary to the sever

- scp -i ~/.ssh/your-key.pem ./sessionapp ec2-user@your-dns:/home/ec2-user
- scp scp -r -i ~/.ssh/your-key.pem ./templates ec2-user@your-dns:/home/ec2-user

3. SSH into your server

- cd ~/.ssh
- ssh -i "your-key.pem" ec2-user@your-dns

4. Run your code

- sudo chmod 700 sessionapp
- sudo ./sessionapp &> go.out

5. go your-dns:80
