#!/bin/bash
# ==== 1. Actualizar el sistema e instalar herramientas ====
yum update -y

yum install -y git

wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz

rm -rf /usr/local/go

tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

export HOME=/home/ec2-user
export GOCACHE=/tmp/go-cache
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

cd /home/ec2-user/stocksRefferer/backend
go build -o backend
./backend &

source ~/.bashrc

yum install -y nginx

curl -fsSL https://rpm.nodesource.com/setup_22.x | sudo bash -
sudo yum install -y nsolid

# ==== 2. Clonar repositorio ====
cd /home/ec2-user/
git clone https://github.com/MoraDiego/stocksRefferer.git
mv /home/ec2-user/.env /home/ec2-user/stocksRefferer/api/.env
cd stocksRefferer

# ==== 3. Backend Go ====
sudo chmod -R 757 api
cd api
go version
go build -buildvcs=false -o backend
./backend &

# ==== 4. Frontend Vue ====
cd ../frontend

npm install
npm run build

sudo tee /etc/nginx/nginx.conf > /dev/null <<EOF
user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen 80;
        server_name _;

        root /home/ec2-user/stocksRefferer/frontend/dist;
        index index.html;

        location / {
            try_files \$uri \$uri/ /index.html;
        }

        location /api/ {
            proxy_pass http://localhost:8080/;
            proxy_http_version 1.1;
            proxy_set_header Upgrade \$http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host \$host;
            proxy_cache_bypass \$http_upgrade;
        }
    }
}
EOF
sudo chmod -R 755 /home
systemctl enable nginx
systemctl restart nginx