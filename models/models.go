package models

// User is like a character in a game. Each user has a unique ID, a special name (username), and an email address.
// They also have a profile and a list of to-do lists.
type User struct {
	ID        uint       `gorm:"primaryKey"`      // This is like a unique badge number for each user.
	Username  string     `gorm:"unique;not null"` // This is the user's special name that no one else can have.
	Email     string     `gorm:"unique;not null"` // This is the user's email address, also unique.
	Profile   Profile    // This is the user's profile, like their personal page.
	TodoLists []TodoList // This is a list of to-do lists that belong to the user.
}

// Profile is like a personal page for the user. It has a unique ID, a bio, and a picture (avatar).
type Profile struct {
	ID        uint   `gorm:"primaryKey"` // This is a unique badge number for each profile.
	UserID    uint   `gorm:"unique"`     // This connects the profile to a specific user.
	Bio       string // This is a short description about the user.
	AvatarURL string // This is a link to the user's picture.
}

// TodoList is like a notebook where users write down tasks they want to do. Each list has a unique ID and a title.
type TodoList struct {
	ID     uint   `gorm:"primaryKey"` // This is a unique badge number for each to-do list.
	Title  string `gorm:"not null"`   // This is the name of the to-do list.
	UserID uint   // This connects the to-do list to a specific user.
	Tasks  []Task // This is a list of tasks in the to-do list.
}

// Task is like a job or activity that needs to be done. Each task has a unique ID, a title, and a description.
type Task struct {
	ID          uint   `gorm:"primaryKey"` // This is a unique badge number for each task.
	Title       string `gorm:"not null"`   // This is the name of the task.
	Description string // This is more information about what the task is.
	Completed   bool   // This tells us if the task is done (true) or not (false).
	TodoListID  uint   // This connects the task to a specific to-do list.
	Tags        []Tag  `gorm:"many2many:task_tags"` // These are labels that help categorize the task.
}

// Tag is like a label or sticker you can put on tasks to organize them. Each tag has a unique ID and a name.
type Tag struct {
	ID    uint   `gorm:"primaryKey"`          // This is a unique badge number for each tag.
	Name  string `gorm:"unique;not null"`     // This is the name of the tag, and it must be unique.
	Tasks []Task `gorm:"many2many:task_tags"` // These are the tasks that have this tag.
}
