package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"os"
)

func main() {
	// Read the instance URL from the environment variable
	instanceURL := os.Getenv("INSTANCE_URL")
	if instanceURL == "" {
		fmt.Println("❌ INSTANCE_URL environment variable is not set")
		return
	}

	client := &http.Client{}
	// Read synthetic data from JSON files
	expenseData, err := readJSONFile("./integration/smoke_tests/expense_data.json")
	if err != nil {
		fmt.Println("❌ Error reading expense data:", err)
		return
	}	
	testCreateExpense(client, instanceURL, expenseData)

	//test CreateBudget handler
	budgetData, err := readJSONFile("./integration/smoke_tests/budget_data.json")
	if err != nil {
		fmt.Println("❌ Error reading budget data:", err)
		return
	}
	testCreateBudget(client, instanceURL, budgetData)
}

func readJSONFile(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result map[string]interface{}
	err = json.NewDecoder(file).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func testHelloWorld(client *http.Client, instanceURL string) {
	url := fmt.Sprintf("%s/", instanceURL)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ HelloWorld test failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("❌ HelloWorld test failed: Unexpected status code", resp.StatusCode)
		return
	}

	fmt.Println("✅ HelloWorld test passed")
}

func testCreateExpense(client *http.Client, instanceURL string, data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("❌ Error marshalling JSON:", err)
		return
	}

	url := fmt.Sprintf("%s/expenses", instanceURL)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ CreateExpense test failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusBadRequest {
		fmt.Println("❌ CreateExpense test failed: Unexpected status code", resp.StatusCode)
		return
	}

	fmt.Println("✅ CreateExpense test passed")
}

// test budget handlers
// test CreateBudget handler
func testCreateBudget(client *http.Client, instanceURL string, data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("❌ Error marshalling JSON:", err)
		return
	}

	url := fmt.Sprintf("%s/budgets", instanceURL)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ CreateBudget test failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusBadRequest {
		fmt.Println("❌ CreateBudget test failed: Unexpected status code", resp.StatusCode)
		return
	}

	fmt.Println("✅ CreateBudget test passed")
}