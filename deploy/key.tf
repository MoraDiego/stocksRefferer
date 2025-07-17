resource "aws_key_pair" "clave_ec2" {
  key_name   = "clave-ec2"
  public_key = file("~/.ssh/clave-ec2.pub")
}
