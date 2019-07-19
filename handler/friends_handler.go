package handler

import (
	"context"

	"github.com/duladissa/friends_management/dao"
	"github.com/duladissa/friends_management/database"
	"github.com/duladissa/friends_management/models"
	"github.com/duladissa/friends_management/restapi/operations/friends"
	"github.com/go-openapi/runtime/middleware"
)

//FriendsAPI ... Friends API implementation
type FriendsAPI struct {
	// database *database.Database
	friend *dao.FriendDAO
}

//NewFriendsAPI ... Create New NewFriendsAPI
func NewFriendsAPI(db *database.Database) *FriendsAPI {
	return &FriendsAPI{
		// database: db,
		friend: dao.NewFriendDAO(db)}
}

// PostFriendsConnections is 1. As a user, I need an API to create a friend connection between two email addresses.
func (f *FriendsAPI) PostFriendsConnections(ctx context.Context, params friends.PostFriendsConnectionsParams) middleware.Responder {
	isAdded := false
	var err error
	var errorResponse models.ErrorResponse
	if len(params.Body.Friends) > 2 {
		errorResponse = models.ErrorResponse{
			Success: &isAdded,
			Message: "Bad input values",
			Type:    "exception",
		}
		return friends.NewPostFriendsConnectionsBadRequest().WithPayload(&errorResponse)
	}

	isAdded, err = f.friend.Create(params.Body)
	if err != nil {
		errorResponse = models.ErrorResponse{
			Success: &isAdded,
			Message: err.Error(),
			Type:    "exception",
		}
		return friends.NewPostFriendsConnectionsNotFound().WithPayload(&errorResponse)
	}

	success := models.SuccessResponse{
		Success: &isAdded,
	}
	return friends.NewPostFriendsConnectionsOK().WithPayload(&success)
}

// PostFriendsConnectionsList is 2. As a user, I need an API to retrieve the friends list for an email address.
func (f *FriendsAPI) PostFriendsConnectionsList(ctx context.Context, params friends.PostFriendsConnectionsListParams) middleware.Responder {
	response, err := f.friend.FindFriends(params.Body)
	isSuccess := false
	if err != nil {
		errorResponse := models.ErrorResponse{
			Success: &isSuccess,
			Message: err.Error(),
			Type:    "exception",
		}
		return friends.NewPostFriendsConnectionsListNotFound().WithPayload(&errorResponse)
	}
	isSuccess = true
	friendsNames := make([]string, 0)
	for k := range response {
		friendsNames = append(friendsNames, k)
	}
	friendsList := models.FriendsListResponse{
		Success: &isSuccess,
		Friends: friendsNames,
		Count:   int64(len(response)),
	}

	return friends.NewPostFriendsConnectionsListOK().WithPayload(&friendsList)
}

// PostFriendsCommonList is 3. As a user, I need an API to retrieve the common friends list between two email addresses.
func (f *FriendsAPI) PostFriendsCommonList(ctx context.Context, params friends.PostFriendsCommonListParams) middleware.Responder {
	response, err := f.friend.FindCommonFriends(params.Body)
	isSuccess := false
	if err != nil || len(response) == 0 {
		errorResponse := models.ErrorResponse{
			Success: &isSuccess,
			Message: "No common friends available",
			Type:    "exception",
		}
		return friends.NewPostFriendsCommonListNotFound().WithPayload(&errorResponse)
	}
	isSuccess = true
	friendsList := models.FriendsListResponse{
		Success: &isSuccess,
		Friends: response,
		Count:   int64(len(response)),
	}

	return friends.NewPostFriendsCommonListOK().WithPayload(&friendsList)
}

// PostFriendsUpdatesSubscribe is 4. As a user, I need an API to subscribe to updates from an email address.</br> Please note that subscribing to updates is NOT equivalent to adding a friend connection
func (f *FriendsAPI) PostFriendsUpdatesSubscribe(ctx context.Context, params friends.PostFriendsUpdatesSubscribeParams) middleware.Responder {
	return nil
}

// PostFriendsUpdatesBlock is 5. As a user, I need an API to block updates from an email address.</br> Suppose andy@example.com blocks john@example.com:</br> <ul> <li>if they are connected as friends, then andy will no longer receive notifications from john</li> <li>if they are not connected as friends, then no new friends connection can be added</li> </ul>
func (f *FriendsAPI) PostFriendsUpdatesBlock(ctx context.Context, params friends.PostFriendsUpdatesBlockParams) middleware.Responder {
	return nil
}

// PostFriendsUpdatesList is 6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.</br> Eligibility for receiving updates from i.e. john@example.com</br> <ul> <li>has not blocked updates from john@example.com, and</li> <li>at least one of the following:</li> <ul> <li>has a friend connection with john@example.com</li> <li>has subscribed to updates from john@example.com</li> <li>has been @mentioned in the update</li> </ul> </ul>
func (f *FriendsAPI) PostFriendsUpdatesList(ctx context.Context, params friends.PostFriendsUpdatesListParams) middleware.Responder {
	return nil
}
