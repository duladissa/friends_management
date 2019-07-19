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
func (f *FriendDAO) Find(user *models.Friend) (map[string]bool, error) {
	db, session := f.database.GetSession()
	defer session.Close()

	collection := db.C(FriendsCollection)
	//Sanitizing the emails
	userID := strings.ToLower(user.Email)

	friends := make([]models.DBFriends, 0)
	condition1 := bson.M{"user_id": userID}
	collection.Find(condition1).All(&friends)

	beingAFriend := make([]models.DBFriends, 0)
	condition2 := bson.M{"friends": userID}
	collection.Find(condition2).All(&beingAFriend)

	results := make(map[string]bool, 0)
	if len(friends) == 0 && len(beingAFriend) == 0 {
		return results, errors.New("Not found any friends")
	}

	for _, friend := range friends {
		if !results[friend.Friends[0]] {
			results[friend.Friends[0]] = true
		}
	}

	for _, friend := range beingAFriend {
		if !results[friend.UserID] {
			results[friend.UserID] = true
		}
	}
	return results, nil
}
