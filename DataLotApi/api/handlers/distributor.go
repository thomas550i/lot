package handlers

import (
	c "DataLotApi/class"
	orm "DataLotApi/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Get_Distributor(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)

	db = orm.Db()
	repos := []c.Distributor{}
	rows, err := db.Query(`select 
	id,
	firsrname,
	lastname,
	address,
	city,
	email,
	distributorid,
	mobile,
	joined,
	isactive
	 from distributor`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.Distributor{}
		err = rows.Scan(
			&repo.ID,
			&repo.Firsrname,
			&repo.Lastname,
			&repo.Address,
			&repo.City,
			&repo.Email,
			&repo.DistributorID,
			&repo.Mobile,
			&repo.Joined,
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

func Delete_Distributor(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Distributor{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	QRY := fmt.Sprintf(`DELETE FROM Distributor WHERE ID=%d`, p.ID)
	ckrows, ckerr := db.Query(QRY)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	} else {
		c.ResSuccess(w, nil, "Deleted Successfully")
	}
	defer ckrows.Close()
}

func Save_Distributor(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Distributor{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Users{}

	QRY := fmt.Sprintf(`Select DistributorID from Distributor where DistributorID='%s'`, strings.ToUpper(p.DistributorID))
	ckrows, ckerr := db.Query(QRY)
	if ckerr != nil {
		c.ResFail(w, nil, ckerr.Error())
		return
	}
	defer ckrows.Close()
	if ckrows.Next() {
		c.ResFail(w, nil, "Distributor ID already Exist!")
		return
	}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	time := p.Joined.In(loc)

	y, m, d := time.Date()
	datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)
	QRY = fmt.Sprintf(`INSERT INTO Distributor (firsrname,lastname,address,city,email,distributorid,mobile,joined,IsActive) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s',%v);`, p.Firsrname, p.Lastname, p.Address, p.City, p.Email, strings.ToUpper(p.DistributorID), p.Mobile, datestring, p.IsActive)
	rows, err := db.Query(QRY)
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Saved Successfully")
}

func Update_Distributor(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.Distributor{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input")
		return
	}

	db = orm.Db()
	repos := []c.Distributor{}

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

	loc, _ := time.LoadLocation("Asia/Kolkata")
	time := p.Joined.In(loc)

	y, m, d := time.Date()
	datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)
	QRY := fmt.Sprintf(`UPDATE Distributor SET  firsrname='%s',
	lastname='%s',
	address='%s',
	city='%s',
	email='%s',
	distributorid='%s',	mobile='%s',
	joined='%s', IsActive = %v WHERE ID = %d`, p.Firsrname, p.Lastname, p.Address, p.City, p.Email, strings.ToUpper(p.DistributorID), p.Mobile, datestring, p.IsActive, p.ID)
	rows, err := db.Query(QRY)
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}
	defer rows.Close()

	c.ResSuccess(w, repos, "Updated Successfully")
}
