package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
)

type Expense struct {
    ID           string  `json:"id,omitempty"` // omitempty for cases where ID isn't initially set
    Description  string  `json:"description"`
    Amount       float64 `json:"amount"`
    DateCreation int64   `json:"date_creation,omitempty"` // assuming this is auto-set by the server
}

func main() {
    instanceURL := os.Getenv("INSTANCE_URL")
    if instanceURL == "" {
        fmt.Println("❌ INSTANCE_URL environment variable is not set")
        return
    }

    client := &http.Client{}

    // Load test data from JSON file
    items, err := readJSONFile("./integration/smoke_tests/expense_data.json")
    if err != nil {
        fmt.Println("❌ Error reading expense data:", err)
        return
    }

    // Iterate through loaded test data to create expenses
    var createdExpenseIDs []string
    for _, item := range items {
        expenseID := testCreateExpense(client, instanceURL, item)
        if expenseID != "" {
            createdExpenseIDs = append(createdExpenseIDs, expenseID)
        }
    }

    // Clean up created test data and test delete operation
    for _, expenseID := range createdExpenseIDs {
		testDeleteExpense(client, instanceURL, expenseID)
    }
}

func readJSONFile(filename string) ([]Expense, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var items []Expense
    err = json.NewDecoder(file).Decode(&items)
    if err != nil {
        return nil, err
    }

    return items, nil
}

func testCreateExpense(client *http.Client, instanceURL string, expense Expense) string {
    jsonData, err := json.Marshal(expense)
    if err != nil {
        fmt.Println("❌ Error marshalling JSON:", err)
        return ""
    }

    url := fmt.Sprintf("%s/expenses", instanceURL)
    resp, err := makeHTTPRequest(client, "POST", url, jsonData)
    if err != nil || resp.StatusCode != http.StatusCreated {
        fmt.Printf("❌ CreateExpense test failed: %v, Status Code: %d\n", err, resp.StatusCode)
        return ""
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("❌ Error reading response body:", err)
        return ""
    }

    var responseExpense Expense
    if err := json.Unmarshal(responseBody, &responseExpense); err != nil {
        fmt.Println("❌ Error decoding response JSON:", err)
        return ""
    }

    if responseExpense.ID != "" {
        fmt.Println("✅ CreateExpense test passed with ID:", responseExpense.ID)
        return responseExpense.ID
    }

    return responseExpense.ID
}

func testDeleteExpense(client *http.Client, instanceURL, expenseID string) {
    url := fmt.Sprintf("%s/expenses?id=%s", instanceURL, expenseID)
    resp, err := makeHTTPRequest(client, "DELETE", url, nil)
    if err != nil || resp.StatusCode != http.StatusOK {
        fmt.Printf("❌ DeleteExpense test failed for ID %s: %v, Status Code: %d\n", expenseID, err, resp.StatusCode)
        return
    }

    fmt.Println("✅ DeleteExpense test passed for ID:", expenseID)
}

func makeHTTPRequest(client *http.Client, method, url string, body []byte) (*http.Response, error) {
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    // Note: defer statement removed from here as it should be the caller's responsibility to close the body

    return resp, nil
}
