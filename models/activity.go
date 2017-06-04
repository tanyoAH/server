package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Activity struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	VendorId    bson.ObjectId `json:"-" bson:"vendorId"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	// TODO - figure out tripadvisor integration
	TripAdvisorUrl   string          `json:"tripAdvisorUrl" bson:"trip_advisor_url"`
	Price            float64         `json:"price" bson:"price"`
	ThumbnailUrl     string          `json:"thumbnailUrl" bson:"thumbnail_url"`
	IsEvening        bool            `json:"isEvening" bson:"is_evening"`
	DurationHours    float64         `json:"durationHours" bson:"duration_hours"`
	TimePeriod       *TimePeriod     `json:"timePeriod" bson:"time_period"`
	Location         MgoXY           `json:"location" bson:"location"`
	Interests        []string        `json:"interests" bson:"interests"`
	CommittedTripIds []bson.ObjectId `json:"-" bson:"committed_trip_ids"`
	GroupChat        []ChatMessage   `json:"groupChat,omitempty" bson:"group_chat"`
	CreatedAt        time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt        time.Time       `json:"updatedAt" bson:"updated_at"`
}

func (activity *Activity) updateTS() {
	activity.UpdatedAt = time.Now()
}

func (activity *Activity) GetDetailedResponse(userId, tripId bson.ObjectId) (*DetailedActivityResponse, error) {
	// TODO
	err := activitiesC.FindId(activity.Id).One(activity)
	if err != nil {
		return nil, err
	}
	resp := DetailedActivityResponse{Activity: *activity}
	for _, tId := range resp.CommittedTripIds {
		if tId == tripId {
			resp.IsCommitted = true
			break
		}
	}
	if resp.IsCommitted {
		userIds, err := GetUserIdsFromTripIdsArray(resp.CommittedTripIds)
		if err != nil {
			return nil, err
		}
		resp.CommittedUsers, err = GetBasicUserResponsesFromIdArray(userIds)
		if err != nil {
			return nil, err
		}
		return &resp, nil
	}
	resp.GroupChat = nil
	return &resp, nil
}

func (activity *Activity) FindById() error {
	return activitiesC.FindId(activity.Id).One(activity)
}

func (activity *Activity) UpdateCommitted(tripId bson.ObjectId) error {
	err := activity.FindById()
	if err != nil {
		return err
	}
	activity.updateTS()
	activity.CommittedTripIds = append([]bson.ObjectId{tripId}, activity.CommittedTripIds...)
	err = activitiesC.UpdateId(activity.Id, bson.M{"$set": bson.M{"committed_trip_ids": activity.CommittedTripIds, "updated_at": activity.UpdatedAt}})
	if err != nil {
		return err
	}

	return nil
}

func (activity *Activity) AddChatMessage(basicUser BasicUserResponse, text string) error {
	err := activity.FindById()
	if err != nil {
		return err
	}
	activity.updateTS()
	activity.GroupChat = append([]ChatMessage{{User: basicUser, Text: text, CreatedAt: time.Now()}}, activity.GroupChat...)
	err = activitiesC.UpdateId(activity.Id, bson.M{"$set": bson.M{"group_chat": activity.GroupChat, "updated_at": activity.UpdatedAt}})
	if err != nil {
		return err
	}

	return nil
}

func (activity *Activity) Create() error {
	if !activity.Id.Valid() {
		activity.Id = bson.NewObjectId()
	}
	return activitiesC.Insert(activity)
}

func GetCommittedActivitiesForTrip(tripId bson.ObjectId) ([]BasicActivityResponse, error) {
	respArr := []BasicActivityResponse{}
	err := activitiesC.Find(bson.M{"committed_trip_ids": bson.M{"$elemMatch": bson.M{"$eq": tripId}}}).All(&respArr)
	if err != nil {
		return nil, err
	}
	return respArr, nil
}

func GetRecommendedActivitiesForTrip(user User, trip Trip) ([]BasicActivityResponse, error) {
	respArr := []BasicActivityResponse{}
	err := activitiesC.Find(bson.M{"interests": bson.M{"$elemMatch": bson.M{"$in": user.Interests}}}).All(&respArr)
	if err != nil {
		return nil, err
	}
	return respArr, nil
}
