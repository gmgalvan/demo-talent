# Get the default subnet
data "aws_subnet" "default" {
  vpc_id            = data.aws_vpc.default.id
  availability_zone = "us-east-1a"  
}

# Create a security group
resource "aws_security_group" "demo_security_group" {
  name        = "demo-security-group"
  description = "Allow inbound traffic on port 8080 and SSH"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Get the existing key pair
data "aws_key_pair" "demo_key_pair" {
  key_name = "demo-key-pair"
}


# Create an EC2 instance
resource "aws_instance" "demo_instance" {
  ami           = "ami-0c4f7023847b90238"
  instance_type = "t2.micro"
  key_name      = data.aws_key_pair.demo_key_pair.key_name
  subnet_id     = data.aws_subnet.default.id
  vpc_security_group_ids = [aws_security_group.demo_security_group.id]

  tags = {
    Name = "dev-demo-instance"
  }
}


# Create a security group for the prod instance
resource "aws_security_group" "prod_security_group" {
  name        = "prod-security-group"
  description = "Allow inbound traffic on port 80"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create an EC2 instance for prod
resource "aws_instance" "prod_instance" {
  ami           = "ami-0c4f7023847b90238"
  instance_type = "t2.micro"
  key_name      = data.aws_key_pair.demo_key_pair.key_name
  subnet_id     = data.aws_subnet.default.id
  vpc_security_group_ids = [aws_security_group.prod_security_group.id]

  tags = {
    Name = "prod-demo-instance"
  }
}