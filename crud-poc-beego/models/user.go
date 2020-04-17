// package models

// import (
// 	"errors"
// 	"strconv"
// 	"time"
// )

// var (
// 	UserList map[string]*User
// )

// func init() {
// 	UserList = make(map[string]*User)
// 	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
// 	UserList["user_11111"] = &u
// }

// type User struct {
// 	Id       string
// 	Username string
// 	Password string
// 	Profile  Profile
// }

// type Profile struct {
// 	Gender  string
// 	Age     int
// 	Address string
// 	Email   string
// }

// func AddUser(u User) string {
// 	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.Id] = &u
// 	return u.Id
// }

// func GetUser(uid string) (u *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("User not exists")
// }

// func GetAllUsers() map[string]*User {
// 	return UserList
// }

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Username == username && u.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
package models

import (
	"context"
	users_db "crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/datasources/mongodb/userdb"
	"crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	//StatusActive user create has a default status of active
	StatusActive = "Active"
)

//User ...
type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName   string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	CreateDate string             `json:"create_date,omitempty" bson:"create_date,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
}

var (
	collection *mongo.Collection
)

func init() {
	collection = users_db.GetMongoInstance().Database("usersdb").Collection("users")
}

//Insert user ingto the database
func (user *User) Insert() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// id := fmt.Sprintf("%v", res.InsertedID)
	// user.ID = id[10 : len(id)-2]
	return nil
}

//GetUser get single user from users db
func (user *User) GetUser(id string) *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := collection.FindOne(ctx, User{ID: user.ID}).Decode(&user); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//GetAllUser ...
func (user *User) GetAllUser() ([]User, *errors.RestErr) {
	users := make([]User, 0)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, User{})
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

//UpdateUser ...
func (user *User) UpdateUser() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	updateBson := bson.M{}
	if user.FirstName != "" {
		updateBson["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		updateBson["last_name"] = user.LastName
	}
	if user.Email != "" {
		updateBson["email"] = user.Email
	}
	if user.Status != "" {
		updateBson["status"] = user.Status
	}
	update := bson.M{"$set": updateBson}
	result, err := collection.UpdateOne(ctx, User{ID: user.ID}, update)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	fmt.Println(result.ModifiedCount)
	return nil
}

//DeleteUser ...
func (user *User) DeleteUser() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.DeleteOne(ctx, User{ID: user.ID})
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
