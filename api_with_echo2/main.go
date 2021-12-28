package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	users       []User
	courses     []Course
	instructors []Instructor
)

type server struct {
}

func init() {
	fmt.Println("Initializing")
	if err := readContent("./data/users.json", &users); err != nil {
		log.Fatal("Errors reading Users")
	}
	if err := readContent("./data/instructors.json", &instructors); err != nil {
		log.Fatal("Errors reading instructors")
	}
	if err := readContent("./data/courses.json", &courses); err != nil {
		log.Fatal("Errors reading courses")
	}
}

func readContent(file string, container interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	json.Unmarshal(b, container)
	return nil
}

func contains(container []string, entry string) bool {
	for _, content := range container {
		if content == entry {
			return true
		}
	}
	return false
}

func isEmpty(set []string) bool {
	if set == nil || len(set) == 0 {
		return true
	}
	return false
}
func intersect(setA []string, setB []string) (bool, []string) {
	var intersection []string
	if isEmpty(setA) || isEmpty(setB) {
		return false, nil
	}

	for _, entryA := range setA {
		if contains(setB, entryA) {
			intersection = append(intersection, entryA)
		}
	}

	return isEmpty(intersection), intersection
}

func getAllUsers(c echo.Context) error {
	interests := []string{}
	var _users []User
	if err := echo.QueryParamsBinder(c).Strings("interest", &interests).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid interests ")
	}

	for _, v := range users {
		if e, _ := intersect(v.Interests, interests); e == true {
			_users = append(_users, v)
		}
	}
	return c.JSON(http.StatusOK, _users)

}

func getUserById(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}
	var data *User

	for _, v := range users {
		if v.ID == id {
			data = &v
			break
		}
	}

	if data == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user with Id not found")
	}
	return c.JSON(http.StatusOK, data)

}

func getAllInstructors(c echo.Context) error {
	courses := []string{}
	var _instructors []Instructor

	if err := echo.QueryParamsBinder(c).Strings("course", &courses).BindError(); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid course passed in")
	}

	for _, v := range instructors {
		if e, _ := intersect(v.Expertise, courses); e == true {
			_instructors = append(_instructors, v)
		}
	}
	return c.JSON(http.StatusOK, _instructors)

}

func getInstructorsById(c echo.Context) error {
	id := -1

	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid Request")
	}

	for _, inst := range instructors {
		if inst.ID == id {
			return c.JSON(http.StatusOK, inst)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)

}

func addUser(c echo.Context) error {
	_user := new(User)

	if err := c.Bind(_user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	_u := User{
		ID:        _user.ID,
		Name:      _user.Name,
		Email:     _user.Email,
		Company:   _user.Company,
		Interests: _user.Interests,
	}
	users = append(users, _u)
	return c.JSON(http.StatusOK, _u)
}

func updateUser(c echo.Context) error {
	_id := -1
	_user := new(User)

	if err := echo.PathParamsBinder(c).Int("id", &_id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid User Id")
	}
	if err := c.Bind(_user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Body")
	}

	for _, user := range users {
		if user.ID == _id {
			user.Name = _user.Name
			user.Email = _user.Email
			user.Company = _user.Company
			user.Interests = _user.Interests
			return c.JSON(http.StatusOK, user)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "{}")
}

func addInstructor(c echo.Context) error {
	_instructor := new(Instructor)

	if err := c.Bind(_instructor); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "{}")
	}
	_instructor.ID = len(instructors) + 1

	instructors = append(instructors, *_instructor)
	return c.JSON(http.StatusOK, _instructor)
}

func updateInstructor(c echo.Context) error {
	_id := -1
	_instructor := new(Instructor)

	if err := echo.PathParamsBinder(c).Int("id", &_id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "{\"message\":\"Bad ID}")
	}

	if err := c.Bind(_instructor); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "{\"message\":\"Bad Body}")
	}

	for _, instructor := range instructors {

		if instructor.ID == _id {
			instructor.Company = _instructor.Company
			instructor.Email = _instructor.Email
			instructor.Expertise = _instructor.Expertise
			instructor.Name = _instructor.Name
			return c.JSON(http.StatusOK, instructor)
		}
	}
	return echo.ErrNotFound
}

// func updateUser(c echo.Context) error {

// }

// type User struct {
// 	Username string `json:username`
// 	Email    string `json:email`
// 	Age      int    `json:age`
// }

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Interests []string `json:"interests"`
}

type Instructor struct {
	ID        int      `json:"id"`
	Name      string   `json."name"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Expertise []string `json:"expertise"`
}

type Course struct {
	ID           int      `json:"id"`
	InstructorID int      `json:"instructor_id"`
	Name         string   `json:"name"`
	Topics       []string `json:"topics"`
	Attendees    []string `json:"attendees"`
}

// func (s *server) home(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello World")
// }

// func (s *server) getUser(c echo.Context) error {
// 	return c.JSON(http.StatusOK, s.user)
// }

// func (s *server) updateUser(c echo.Context) error {
// 	u := new(User)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}

// 	s.user = *u
// 	return c.JSON(http.StatusOK, u)

// }

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("==========Logging Middleware==============")
		log.Print(c.Request().URL)
		return next(c)
	}
}
func main() {
	courses := []string{"math", "science", "english"}
	courses3 := []string{"english", "chem", "science"}
	fmt.Println(courses)
	fmt.Println(isEmpty(courses))
	fmt.Printf("courses contain math:%v\n", contains(courses, "math"))
	fmt.Printf("Coourses contain chemistry:%v\n", contains(courses, "chrm"))
	fmt.Println(intersect(courses, courses3))

	e := echo.New()
	specialLogger := middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "time=${timerfc3339 method= ${method}, URI:${uri}, Status=${status}"})
	e.Use(Logger, specialLogger) // need this call to use the middlewares, can use multuple middlewares
	v1 := e.Group("/api/v1")
	v1.GET("/users", getAllUsers)
	v1.GET("/users/:id", getUserById)
	v1.POST("/users", addUser)
	v1.PUT("/user/:id", updateUser)

	v1.GET("/instructors", getAllInstructors)
	v1.GET("/instructor/:id", getInstructorsById)
	v1.POST("/instructors", addInstructor)
	v1.PUT("/instructors/:id", updateInstructor)

	e.Logger.Fatal(e.Start(":5000"))
	// s := &server{
	// 	user: User{
	// 		Username: "pranavsh",
	// 		Email:    "coolpranavsharma@gmail.com",
	// 		Age:      11,
	// 	},
	// }
	// e := echo.New()
	// e.GET("/", s.home)
	// e.GET("/user", s.getUser)
	// e.PUT("/user", s.updateUser)
	// e.Logger.Fatal(e.Start(":5000"))
}
