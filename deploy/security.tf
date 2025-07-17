resource "aws_security_group" "web_sg" {
  name        = "web_sg"
  description = "Permitir HTTP, SSH, Nginx"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Frontend Vue (serve -l 3000)
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"          # Todos los protocolos
    cidr_blocks = ["0.0.0.0/0"] # A cualquier destino
  }
}
