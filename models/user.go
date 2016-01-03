package models

type User struct {
    *Models
    Id string `bson:"_id"`
    Name string `bson:"name"`
    Surname string `bson:"surname"`
    Role  string `bson:"role"`
    Score int `bson:"score"`
    Courses []struct {
        CourseId  string `bson:"courseId"`
        Progress  string `bson:"progress"`
        Score     string `bson:"score"`
    }
}
