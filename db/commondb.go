/*
This file is under GNU AFFERO GENERAL PUBLIC LICENSE

Permissions of this strongest copyleft license are conditioned
on making available complete source code of licensed works and
modifications, which include larger works using a licensed work,
under the same license. Copyright and license notices must be preserved.
Contributors provide an express grant of patent rights.
When a modified version is used to provide a service over a network,
the complete source code of the modified version must be made available.

Edoardo Ottavianelli, https://edoardoottavianelli.it
*/

package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Website struct {
	Address   string `json:"address Str"`
	Body      string `json:"body Str"`
	Timestamp int64  `json:"timestamp Int"`
}

type User struct {
	Email string `json:"email Str"`
}

//GetDatabase returns the pointer to the database d(input).
func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)

	return database
}

//GetUsers returns the collection of users
func GetUsers(database *mongo.Database) *mongo.Collection {
	return database.Collection("users")
}

//GetWebsites returns the collection of wesbites
func GetWebsites(database *mongo.Database) *mongo.Collection {
	return database.Collection("websites")
}

//GetAllEmails returns all the emails entered
func GetAllEmails(users []User) []string {
	var results []string
	for _, value := range users {
		results = append(results, value.Email)
	}
	return results
}
