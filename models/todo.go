package models

type Todo struct {
	ID        string `json:"id" dynamodbav:"id"` // Partition Key
	Title     string `json:"title" dynamodbav:"title"`
	Completed bool   `json:"completed" dynamodbav:"completed"`
}
