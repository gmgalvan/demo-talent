
resource "aws_security_group" "postgres" {
  name_prefix = "postgres-"

  ingress {
    from_port   = 5432
    to_port     = 5432
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

resource "aws_security_group_rule" "rds_ingress_demo" {
  type                     = "ingress"
  from_port                = 5432
  to_port                  = 5432
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.demo_security_group.id
  security_group_id        = aws_security_group.postgres.id
}


resource "aws_db_instance" "postgres" {
  engine              = "postgres"
  engine_version      = "16.1"
  instance_class      = "db.t3.micro"
  identifier          = "demo-postgres-db"
  db_name             = "demoTalent"
  username            = "demouser"
  password            = var.db_password
  allocated_storage   = 20
  storage_type        = "gp2"
  publicly_accessible = false
  skip_final_snapshot = true

  vpc_security_group_ids = [aws_security_group.postgres.id]
}


output "postgres_endpoint" {
  value = aws_db_instance.postgres.endpoint
}