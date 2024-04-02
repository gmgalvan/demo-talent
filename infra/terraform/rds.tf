
resource "aws_security_group" "postgres" {
  name_prefix = "postgres-"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.demo_security_group.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_db_instance" "postgres" {
  engine               = "postgres"
  engine_version       = "16.1"
  instance_class       = "db.t3.micro"
  identifier           = "demo-postgres-db"
  db_name              = "demoTalent"
  username             = "demouser"
  password             = var.db_password
  allocated_storage    = 20
  storage_type         = "gp2"
  publicly_accessible  = false
  skip_final_snapshot  = true
  vpc_security_group_ids = [aws_security_group.postgres.id]
}


output "postgres_endpoint" {
  value = aws_db_instance.postgres.endpoint
}