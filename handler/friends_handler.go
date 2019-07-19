package handler

import (
	"context"

	"github.com/duladissa/friends_management/restapi/operations/friends"
	"github.com/go-openapi/runtime/middleware"
)

//FriendsAPI ... Friends API implementation
type FriendsAPI struct {
}

//NewFriendsAPI ... Create New NewFriendsAPI
func NewFriendsAPI() *FriendsAPI {
	return &FriendsAPI{}
}

// PostFriendsConnections is 1. As a user, I need an API to create a friend connection between two email addresses.
func (f *FriendsAPI) PostFriendsConnections(ctx context.Context, params friends.PostFriendsConnectionsParams) middleware.Responder {
	return nil
}

// PostFriendsConnectionsList is 2. As a user, I need an API to retrieve the friends list for an email address.
func (f *FriendsAPI) PostFriendsConnectionsList(ctx context.Context, params friends.PostFriendsConnectionsListParams) middleware.Responder {
	return nil
}

// PostFriendsCommonList is 3. As a user, I need an API to retrieve the common friends list between two email addresses.
func (f *FriendsAPI) PostFriendsCommonList(ctx context.Context, params friends.PostFriendsCommonListParams) middleware.Responder {
	return nil
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
