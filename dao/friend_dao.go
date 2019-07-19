package dao

import (
	"errors"
	"strings"

	"github.com/duladissa/friends_management/database"
	"github.com/duladissa/friends_management/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	//FriendsCollection ... Friends collection name
	FriendsCollection = "friends"
)

//FriendDAO ... Data Access Object for Friend
type FriendDAO struct {
	database *database.Database
}

//NewFriendDAO ... Create a FriendDAO
func NewFriendDAO(db *database.Database) *FriendDAO {
	return &FriendDAO{database: db}
}

//Create ... Create CRUD operation
func (f *FriendDAO) Create(friends *models.Friends) (bool, error) {
	db, session := f.database.GetSession()
	defer session.Close()

	collection := db.C(FriendsCollection)
	//Sanitizing the emails
	userID := strings.ToLower(friends.Friends[0])
	friend := strings.ToLower(friends.Friends[1])

	existingUser := models.DBFriends{}
	conditions := bson.M{"$and": []bson.M{bson.M{"user_id": userID}, bson.M{"friends": friend}}}
	collection.Find(conditions).One(&existingUser)

	if len(existingUser.Friends) > 0 {
		return false, errors.New("Already available friendship")
	}
	existingUser.Friends = append(existingUser.Friends, friend)
	//Updating the user friends
	updatedUser := models.DBFriends{ID: bson.NewObjectId(), UserID: userID, Friends: existingUser.Friends}
	err := collection.Insert(&updatedUser)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Find ... Find friends for a given user
func (f *FriendDAO) Find(user *models.Friend) {

}
