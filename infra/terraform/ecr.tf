# Create an ECR repository
resource "aws_ecr_repository" "demo_app_repo" {
  name = "demo-app-repo"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = {
    Name = "demo-app-repo"
  }
}

# Output the repository URL
output "ecr_repository_url" {
  value = aws_ecr_repository.demo_app_repo.repository_url
}