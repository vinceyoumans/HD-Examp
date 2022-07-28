package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"runtime"
)

func main() {
	fmt.Println("start here homedepot V03")
	fmt.Println("Hello there!  - using echo mux")

	fmt.Println("multi-stage build")

	osv := runtime.GOOS
	fmt.Printf("The OS is: %s\n", osv)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello there - from v03 - homedepot!")
	})
	//x001
	e.GET("/users/:id", getUser) // works

	//x02
	e.GET("/show", show)

	//x03
	e.POST("/users01", saveUser01)
	//curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:8080/save
	//curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:8080/saveuser01
	// => <b>Thank you! Joe Smith</b>

	e.POST("/users02", saveUser02)

	e.PUT("/users/:id", updateUser)

	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))

}

//=====================================================================================

// http://localhost:1323/users/Joe
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func updateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// curl -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:8080/save
// e.POST("/save", save)
func saveUser01(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveUser02(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

//type User struct {
//	Name  string `json:"name" xml:"name" form:"name" query:"name"`
//	Email string `json:"email" xml:"email" form:"email" query:"email"`
//}
//
//e.POST("/users", func(c echo.Context) error {
//	u := new(User)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	return c.JSON(http.StatusCreated, u)
//	// or
//	// return c.XML(http.StatusCreated, u)
//})
