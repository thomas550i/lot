package handlers

import (
	c "DataLotApi/class"
	orm "DataLotApi/db"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

var db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type authuser struct {
		Username string
		Password string
	}
	var t authuser
	err := decoder.Decode(&t)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}
	log.Println(t.Username)

	db = orm.Db()
	type userinfo struct {
		Lastname  string
		Firstname string
		Username  string
		Gender    string
		Roalval   string
		SessionId string
	}
	repos := []userinfo{}
	rows, err := db.Query(`select 
	lastname,
	firstname,
	username,
	gender,
	roles.roalval
	 from users LEFT JOIN roles ON users.roalid = roles.roalid 
	 where username='` + t.Username + `' and password='` + t.Password + `' `)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := userinfo{}
		err = rows.Scan(
			&repo.Lastname,
			&repo.Firstname,
			&repo.Username,
			&repo.Gender,
			&repo.Roalval,
		)
		if err != nil {
			return
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	if len(repos) <= 0 {
		c.ResFail(w, nil, "Authentication failed. please check your username/password")
		return
	} else {
		repos[0].SessionId = setSession(repos[0].Username)
		c.ResSuccess(w, repos, "")
		return
	}

	// out, err := json.Marshal(repos)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, string(out))
}
func setSession(u string) string {

	return string(rand.Intn(200))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	//....
	// p := c.Users{}
	// c.GetPlayload(r, w, &p)
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&p)
	// if err != nil {
	// 	c.ResFail(w, nil, "Not a valid Input")
	// 	return
	// }

	db = orm.Db()
	repos := []c.Users{}
	rows, err := db.Query(`select 
	lastname,
	firstname,
	username,
	gender,
	roles.roalval,
	Password,
	users.isactive
	 from users LEFT JOIN roles ON users.roalid = roles.roalid`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.Users{}
		err = rows.Scan(
			&repo.Lastname,
			&repo.Firstname,
			&repo.Username,
			&repo.Gender,
			&repo.Roalval,
			&repo.Password,
			&repo.IsActive,
		)
		if err != nil {
			return
		}
		repos = append(repos, repo)
	}
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}

	c.ResSuccess(w, repos, "")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Users{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	ckrows, ckerr := db.Query(`DELETE FROM users WHERE username='` + p.Username + `'`)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	} else {
		c.ResSuccess(w, nil, "Deleted Successfully")
	}
	defer ckrows.Close()
}

func Saveuser(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Users{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Users{}

	ckrows, ckerr := db.Query(`Select username from users where username='` + p.Username + `'`)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	}
	defer ckrows.Close()
	if ckrows.Next() {
		c.ResFail(w, nil, "Username already Exist!")
		return
	}
	if p.IsActive == "true" {
		p.IsActive = "1"
	} else {
		p.IsActive = "0"
	}
	rows, err := db.Query(`INSERT INTO users (firstname, lastname, username, gender, roalid, Password,IsActive)
	VALUES ('` + p.Firstname + `', '` + p.Lastname + `', '` + p.Username + `', '` + p.Gender + `', '` + p.Roalval + `', '` + p.Password + `','` + p.IsActive + `');`)
	if err != nil {
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Saved Successfully")
}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Users{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Users{}

	// ckrows, ckerr := db.Query(`Select username from users where username='` + p.Username + `'`)
	// if ckerr != nil {
	// 	c.ResFail(w, nil, ckerr.Error())
	// 	return
	// }
	// defer ckrows.Close()
	// if ckrows.Next() {
	// 	c.ResFail(w, nil, "Username already Exist!")
	// 	return
	// }

	if p.IsActive == "true" {
		p.IsActive = "1"
	} else {
		p.IsActive = "0"
	}
	rows, err := db.Query(`UPDATE users
	SET  firstname='` + p.Firstname + `', lastname='` + p.Lastname + `', username='` + p.Username + `', gender='` + p.Gender + `', roalid='` + p.Roalval + `', Password='` + p.Password + `',IsActive = '` + p.IsActive + `'
	WHERE username = '` + p.Username + `';`)
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Updated Successfully")
}
