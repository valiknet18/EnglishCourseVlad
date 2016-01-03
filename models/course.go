package models

type Course struct {
    *Models
    Id string `bson:"_id"`
    Name  string  `bson:"name"`
    DependFrom  string `bson:"dependFrom"`
    Level   string `bson:"level"`
    Lessons []struct {
        Type string `bson:"type"`
        Title string `bson:"title"`
        Data  string `bson:"data"`
    }
    Tests []struct {
        Type string `bson:"type"`
        Question string `bson:"question"`
        CorrectAnswer string `bson:"question"`
        Answers []struct {
            Correct bool `bson:"correct"`
            Answer string `bson:"Answer"`
        }
    }
}
