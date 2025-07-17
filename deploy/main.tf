provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "prod" {
  ami           = "ami-0c101f26f147fa7fd"  # Amazon Linux 2 (us-east-1)
  instance_type = "t2.micro"              # Gratuito
  key_name      = "clave-ec2"           # Lo creamos más abajo

  user_data = file("${path.module}/startup.sh")          # Script para instalar dependencias

  tags = {
    Name = "stocksRefferer"
  }

  provisioner "file" {
    source      = "${path.module}/.env"
    destination = "/home/ec2-user/.env"
  }

  provisioner "remote-exec" {
    inline = ["echo '.env copiado'"]
  }

  connection {
    type        = "ssh"
    user        = "ec2-user"
    private_key = file("~/.ssh/clave-ec2")
    host        = self.public_ip
  }
  vpc_security_group_ids = [aws_security_group.web_sg.id]
}

