package dbops

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("TRUNCATE user")
	dbConn.Exec("TRUNCATE video")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil || pwd != "123" {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed!")
	}
}

var tempvid string

func TestVideosWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideos)
	t.Run("GetVideo", testGetVideos)
	t.Run("DelVideo", testDeleteVideos)
	t.Run("RegetVideo", testRegetVideos)
}

func testAddVideos(t *testing.T) {
	vi, err := AddNewVideos(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideos(t *testing.T) {
	_, err := GetVideos(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideo: %v", err)
	}
}

func testDeleteVideos(t *testing.T) {
	err := DeleteVideos(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideo: %v", err)
	}
}

func testRegetVideos(t *testing.T) {
	_, err := GetVideos(tempvid)
	if err != nil {
		t.Errorf("Error of RegetVideo: %v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid, aid, content := "123456", 1, "I like this video"
	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}

}

func testListComments(t *testing.T) {
	vid, form, to := "123456", 1677661739, 1682759339
	_, err := ListComments(vid, form, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
}
