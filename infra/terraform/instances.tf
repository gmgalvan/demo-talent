data "aws_vpc" "default" {
  default = true
}

data "aws_subnet" "default" {
  vpc_id            = data.aws_vpc.default.id
  availability_zone = "us-east-1a"
}

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

data "aws_key_pair" "demo_key_pair" {
  key_name = "demo-key-pair"
}

# Create an IAM policy for CloudWatch Logs
resource "aws_iam_policy" "cloudwatch_logs_policy" {
  name        = "cloudwatch-logs-policy"
  description = "IAM policy for CloudWatch Logs"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      }
    ]
  })
}

# Create an IAM role for the EC2 instance
resource "aws_iam_role" "demo_instance_role" {
  name = "demo-instance-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

# Attach the CloudWatch Logs policy to the IAM role
resource "aws_iam_role_policy_attachment" "cloudwatch_logs_policy_attachment" {
  policy_arn = aws_iam_policy.cloudwatch_logs_policy.arn
  role       = aws_iam_role.demo_instance_role.name
}

# Create an instance profile for the IAM role
resource "aws_iam_instance_profile" "demo_instance_profile" {
  name = "demo-instance-profile"
  role = aws_iam_role.demo_instance_role.name
}

resource "aws_instance" "demo_instance" {
  ami                    = "ami-0c4f7023847b90238"
  instance_type          = "t2.micro"
  key_name               = data.aws_key_pair.demo_key_pair.key_name
  subnet_id              = data.aws_subnet.default.id
  vpc_security_group_ids = [aws_security_group.demo_security_group.id]

  # Attach the instance profile to the EC2 instance
  iam_instance_profile = aws_iam_instance_profile.demo_instance_profile.name

  tags = {
    Name = "dev-demo-instance"
  }
}