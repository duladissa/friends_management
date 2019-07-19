package dao

import (
	"errors"
	"strings"

	"github.com/duladissa/friends_management/database"
	"github.com/duladissa/friends_management/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	//UpdatesCollection ... Updates collection name
	UpdatesCollection = "updates"
)

//UpdatesDAO ... Data Access Object for Updates
type UpdatesDAO struct {
	database *database.Database
}

//NewUpdatesDAO ... Create a UpdatesDAO
func NewUpdatesDAO(db *database.Database) *UpdatesDAO {
	return &UpdatesDAO{database: db}
}

//Create ... Create CRUD operation
func (f *UpdatesDAO) Create(friends *models.SubscribeFriends) (bool, error) {
	db, session := f.database.GetSession()
	defer session.Close()

	collection := db.C(UpdatesCollection)
	//Sanitizing the emails
	requestor := strings.ToLower(friends.Requestor)
	target := strings.ToLower(friends.Target)

	existingUpdates := models.DBUpdates{}
	conditions := bson.M{"$and": []bson.M{bson.M{"user_id": requestor}, bson.M{"target": target}}}
	collection.Find(conditions).One(&existingUpdates)

	if len(existingUpdates.Target) > 0 {
		return false, errors.New("Already available updates")
	}
	existingUpdates.Target = append(existingUpdates.Target, target)
	//Updating the user friends
	updatedUpdates := models.DBUpdates{ID: bson.NewObjectId(), UserID: requestor, Target: existingUpdates.Target}
	err := collection.Insert(&updatedUpdates)
	if err != nil {
		return false, err
	}
	return true, nil
}
