package db

import (
	"context"

	"todo-app-dynamodb/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const TableName = "Todos"

func InsertTodo(todo models.Todo) error {
	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return err
	}
	_, err = DBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      item,
	})
	return err
}

func GetTodo(id string) (*models.Todo, error) {
	resp, err := DBClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}
	if resp.Item == nil {
		return nil, nil
	}

	var todo models.Todo
	err = attributevalue.UnmarshalMap(resp.Item, &todo)
	return &todo, err
}

func UpdateTodo(todo models.Todo) error {
	// Simple overwrite for now but we can also use UpdateTodoPartial
	return UpdateTodoPartial(todo)
}
func UpdateTodoPartial(todo models.Todo) error {
	// Build update expression dynamically
	update := expression.UpdateBuilder{}

	if todo.Title != "" {
		update = update.Set(expression.Name("title"), expression.Value(todo.Title))
	}
	update = update.Set(expression.Name("completed"), expression.Value(todo.Completed))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}

	_, err = DBClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: todo.ID},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})
	return err
}
func DeleteTodo(id string) error {
	_, err := DBClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
