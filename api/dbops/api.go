package dbops

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/yolomc/my-video-server/api/defs"
)

func AddUserCredential(userName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (username,pwd) values (?,?);")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(userName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func GetUserCredential(userName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where username = ? ;")
	if err != nil {
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(userName).Scan(&pwd)
	//这里 NoRows 表示没有数据，会作为 error 返回，所以这里排除掉这种情况
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(userName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where username=? and pwd=? ;")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(userName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {

	vid := uuid.New().String()

	ctime := time.Now().Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("insert into videos (id, author_id,name,display_ctime) values (?,?,?,?);")
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	defer stmtIns.Close()
	vInfo := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return vInfo, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from videos where id=? ;")

	var (
		aid   int
		name  string
		ctime string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var vInfo *defs.VideoInfo
	if err != sql.ErrNoRows {
		vInfo = &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	}

	defer stmtOut.Close()
	return vInfo, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from videos where id=? ;")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComment(vid string, aid int, content string) error {
	id := uuid.New().String()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("insert into comments (id, video_id, author_id,content) values (?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func GetCommentsList(vid string, form, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare("select comments.id, users.username, comments.content from comments inner join users on comments.author_id=users.id where comments.video_id=? and comments.create_time>from_unixtime(?) and comments.create_time<=from_unixtime(?);")
	var cms []*defs.Comment

	rows, err := stmtOut.Query(vid, form, to)
	if err != nil {
		return cms, err
	}

	for rows.Next() {
		var id, username, content string
		if err := rows.Scan(&id, &username, &content); err != nil {
			return cms, err
		}
		cms = append(cms, &defs.Comment{
			Id:      id,
			Author:  username,
			Content: content,
		})
	}

	defer stmtOut.Close()
	return cms, nil
}
