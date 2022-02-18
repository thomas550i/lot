package handlers

import (
	c "DataLotApi/class"
	orm "DataLotApi/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func Get_Show(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)

	db = orm.Db()
	repos := []c.Show{}
	rows, err := db.Query(`select 
	ID,
	Time,
	isactive
	 from shows`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.Show{}
		err = rows.Scan(
			&repo.ID,
			&repo.Time,
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

func Delete_Show(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Show{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	QRY := fmt.Sprintf(`DELETE FROM shows WHERE ID=%d`, p.ID)
	ckrows, ckerr := db.Query(QRY)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	} else {
		c.ResSuccess(w, nil, "Deleted Successfully")
	}
	defer ckrows.Close()
}

func Save_Show(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Show{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Users{}

	QRY := fmt.Sprintf(`Select Time from shows where Time='%s'`, p.Time)
	ckrows, ckerr := db.Query(QRY)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	}
	defer ckrows.Close()
	if ckrows.Next() {
		c.ResFail(w, nil, "Show already Exist!")
		return
	}
	QRY = fmt.Sprintf(`INSERT INTO shows (Time,IsActive) VALUES ('%s',%v);`, p.Time, p.IsActive)
	rows, err := db.Query(QRY)
	if err != nil {
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Saved Successfully")
}

func Update_Show(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Show{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Show{}

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

	QRY := fmt.Sprintf(`UPDATE shows SET  Time='%s',IsActive = %v WHERE ID = %d`, p.Time, p.IsActive, p.ID)
	rows, err := db.Query(QRY)
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Updated Successfully")
}
