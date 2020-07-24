package dbops

import (
	"fmt"
	"testing"
	"time"
)

var tempVid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate videos")
	dbConn.Exec("truncate comments")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUser)
	t.Run("Get", TestGetUser)
	t.Run("Delete", TestDeleteUser)
	t.Run("Reget", TestRegetUser)
}

func TestAddUser(t *testing.T) {
	err := AddUserCredential("zhangsan", "12345")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}

}

func TestGetUser(t *testing.T) {
	pwd, err := GetUserCredential("aaa")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("aaa", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func TestRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("aaa")
	if err != nil {
		t.Errorf("Error of TestRegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", TestAddUser)
	t.Run("AddVideo", TestAddVideoInfo)
	t.Run("GetVideo", TestGetVideoInfo)
	t.Run("DelVideo", TestDeleteVideoInfo)
	t.Run("RegetVideo", TestRegetVideoInfo)
}

func TestAddVideoInfo(t *testing.T) {
	v, err := AddNewVideo(1, "vvv")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempVid = v.Id
}

func TestGetVideoInfo(t *testing.T) {
	v, err := GetVideoInfo(tempVid)
	if err != nil || v == nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}
func TestDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}
func TestRegetVideoInfo(t *testing.T) {
	v, err := GetVideoInfo(tempVid)
	if err != nil || v != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	// t.Run("PrepareUser", TestAddUser)
	// t.Run("PrepareVideo", TestAddVideoInfo)
	t.Run("AddNewComment", TestAddNewComment)
	t.Run("GetCommentsList", TestGetCommentsList)
}

func TestAddNewComment(t *testing.T) {
	err := AddNewComment("123", 1, "good video")
	if err != nil {
		t.Errorf("Error of AddNewComment: %v", err)
	}
}
func TestGetCommentsList(t *testing.T) {
	vid := "123"
	from := 151476800
	to := int(time.Now().Unix())

	cms, err := GetCommentsList(vid, from, to)
	if err != nil {
		t.Errorf("Error of GetCommentsList: %v", err)
	}

	for i, v := range cms {
		fmt.Printf("comment: %d, %v \n", i, v)
	}
}
