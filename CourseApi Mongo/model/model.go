package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	CourseId    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseName  string             `json:"coursename,omitempty"`
	CoursePrice int                `json:"price,omitempty"`
	Author      *Author            `json:"author,omitempty"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}
