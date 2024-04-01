# Demo project [backend]
## A Valid PR

For a PR (Pull Request) to be considered valid, it must fulfill the following criteria:

- Must contain functionality.
- Include unit tests.
- Provide Swagger documentation.
- If required, set up migrations.

## Execute with Docker Compose

```bash
docker-compose up --build
```

## Testing Endpoint
You can test various endpoints by using the curl command. Below are examples of how to test different operations:

- POST: To create a new expense:
```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "description": "Lunch expense",
    "amount": 15.50
}' http://localhost:8080/expenses
```

- GET: To retrieve an expense by its ID:
```bash
curl -X GET http://localhost:8080/expenses?id=<expense_id>
```

- PUT: To update an existing expense:
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "id": "<expense_id>",
    "description": "Updated expense description",
    "amount": 20.00
}' http://localhost:8080/expenses
```

- DELETE: To delete an expense by its ID:
```bash
curl -X DELETE http://localhost:8080/expenses?id=<expense_id>
```

## Documentation
To generate Swagger documentation for your API, use the following commands:

swag init
swagger generate spec -o ./swagger.json

## tests
- Create mocks with go mock
```bash
cd services
mockgen -package=mocks -destination=./mocks/mock_expense_service.go github.com/demo-talent/services ExpenseService
```
```bash
cd repository
mockgen -package=mocks -destination=./mocks/mock_expense_repository.go github.com/demo-talent/repository ExpenseRepositoryInterface
```
- Run tests
```bash
go test ./...
```

## Infra
- Terraform commands
```bash
terraform plan
terraform apply
```

- Modify hosts file on ansible: 
ec2-instance ansible_host=<public-ip-address>
- Run comamnd
```bash
ansible-playbook -i hosts docker-playbook.yml --private-key /path/to/private-key.pem
```

## Show Demo
## Implement filters by property on expenses (Reduced steps)
    0.- create new branch
    1.- Implement on repo the code for filtering
    2.- create the handler
    3.- Implement unit tests
    4.- Update swagger documents
    5.- creae migratiosn files if needed
    6.- Commit and push