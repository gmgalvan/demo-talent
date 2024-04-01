package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"os"
)

func main() {
	// Read synthetic data from JSON files
	expenseData, err := readJSONFile("expense_data.json")
	if err != nil {
		fmt.Println("Error reading expense data:", err)
		return
	}

	// Initialize HTTP client
	client := &http.Client{}

	// Test handlers with synthetic data
	// Test HelloWorld handler
	testHelloWorld(client)

	// Test CreateExpense handler
	testCreateExpense(client, expenseData)
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

func testHelloWorld(client *http.Client) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HelloWorld test failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HelloWorld test failed: Unexpected status code", resp.StatusCode)
		return
	}

	fmt.Println("HelloWorld test passed")
}

func testCreateExpense(client *http.Client, data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, _ := http.NewRequest("POST", "/expenses", strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("CreateExpense test failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusBadRequest {
		fmt.Println("CreateExpense test failed: Unexpected status code", resp.StatusCode)
		return
	}

	fmt.Println("CreateExpense test passed")
}
