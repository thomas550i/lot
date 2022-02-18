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

func Get_Findnumber(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)

	db = orm.Db()
	Approveddata := []c.DailyApproval{}

	type playload struct {
		Date         time.Time
		Showid       string
		CustomNumber string
	}
	p := playload{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	loc, _ := time.LoadLocation("Asia/Kolkata")
	time := p.Date.In(loc)

	y, m, d := time.Date()
	datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

	QRY := fmt.Sprintf(`select id,date,showid,distributorid from dailyapproval where showid='%s' and date='%s'`, p.Showid, datestring)

	rows, err := db.Query(QRY)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.DailyApproval{}
		err = rows.Scan(
			&repo.ID,
			&repo.Date,
			&repo.Showid,
			&repo.Distributorid,
		)
		if err != nil {
			return
		}
		Approveddata = append(Approveddata, repo)
	}
	err = rows.Err()
	if err != nil {
		c.ResFail(w, nil, err.Error())
		return
	}

	//getting all dataentry approved
	diids := []string{}
	Dailydata := []c.DailyData{}
	for _, id := range Approveddata {
		diids = append(diids, id.Distributorid)
	}
	QRY = fmt.Sprintf(`select * from dailydata where distributorid in ('%s') and showid='%s' and date='%s'`, strings.Join(diids, "', '"), p.Showid, datestring)

	rows, err = db.Query(QRY)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		_repo := c.DailyData{}
		err = rows.Scan(
			&_repo.ID,
			&_repo.Date,
			&_repo.Showid,
			&_repo.Distributorid,
			&_repo.A,
			&_repo.SetA,
			&_repo.B,
			&_repo.SetB,
			&_repo.C,
			&_repo.SetC,
			&_repo.AB,
			&_repo.SetAB,
			&_repo.AC,
			&_repo.SetAC,
			&_repo.BC,
			&_repo.SetBC,
			&_repo.ABCFULL,
			&_repo.SetABCFULL,
			&_repo.ABCHALF,
			&_repo.SetABCHALF,
			&_repo.BOXFULL,
			&_repo.SetBOXFULL,
			&_repo.BOXHALF,
			&_repo.SetBOXHALF,
			&_repo.Status,
		)
		if err != nil {
			return
		}
		Dailydata = append(Dailydata, _repo)
	}

	//calclulating number
	type sets struct {
		Setname   string
		Setnumber int64
		Setsales  int64
	}
	seta := []sets{}

	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "A" && maind.Setnumber == da.A {
				seta[idx].Setsales = seta[idx].Setsales + da.SetA
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "A"
			s.Setnumber = da.A
			s.Setsales = da.SetA
			seta = append(seta, s)
		}
	}

	//b
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "B" && maind.Setnumber == da.B {
				seta[idx].Setsales = seta[idx].Setsales + da.SetB
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "B"
			s.Setnumber = da.B
			s.Setsales = da.SetB
			seta = append(seta, s)
		}
	}

	//c
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "C" && maind.Setnumber == da.C {
				seta[idx].Setsales = seta[idx].Setsales + da.SetC
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "C"
			s.Setnumber = da.C
			s.Setsales = da.SetC
			seta = append(seta, s)
		}
	}

	//AB
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "AB" && maind.Setnumber == da.AB {
				seta[idx].Setsales = seta[idx].Setsales + da.SetAB
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "AB"
			s.Setnumber = da.AB
			s.Setsales = da.SetAB
			seta = append(seta, s)
		}
	}

	//BC
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "BC" && maind.Setnumber == da.BC {
				seta[idx].Setsales = seta[idx].Setsales + da.SetBC
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "BC"
			s.Setnumber = da.BC
			s.Setsales = da.SetBC
			seta = append(seta, s)
		}
	}

	//AC
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "AC" && maind.Setnumber == da.AC {
				seta[idx].Setsales = seta[idx].Setsales + da.SetAC
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "AC"
			s.Setnumber = da.AC
			s.Setsales = da.SetAC
			seta = append(seta, s)
		}
	}

	//ABCfull
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "ABCFULL" && maind.Setnumber == da.ABCFULL {
				seta[idx].Setsales = seta[idx].Setsales + da.SetABCFULL
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "ABCFULL"
			s.Setnumber = da.ABCFULL
			s.Setsales = da.SetABCFULL
			seta = append(seta, s)
		}
	}

	//ABCHALF
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "ABCHALF" && maind.Setnumber == da.ABCHALF {
				seta[idx].Setsales = seta[idx].Setsales + da.SetABCHALF
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "ABCHALF"
			s.Setnumber = da.ABCHALF
			s.Setsales = da.SetABCHALF
			seta = append(seta, s)
		}
	}

	//BOXFULL
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "BOXFULL" && maind.Setnumber == da.BOXFULL {
				seta[idx].Setsales = seta[idx].Setsales + da.SetBOXFULL
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "BOXFULL"
			s.Setnumber = da.BOXFULL
			s.Setsales = da.SetBOXFULL
			seta = append(seta, s)
		}
	}

	//BOXHALF
	for _, da := range Dailydata {
		found := false
		for idx, maind := range seta {
			if maind.Setname == "BOXHALF" && maind.Setnumber == da.BOXHALF {
				seta[idx].Setsales = seta[idx].Setsales + da.SetBOXHALF
				found = true
			}
		}
		if !found {
			s := sets{}
			s.Setname = "BOXHALF"
			s.Setnumber = da.BOXHALF
			s.Setsales = da.BOXHALF
			seta = append(seta, s)
		}
	}
	setnumbers := make(map[string]int64)
	for i := 1; i < 999; i++ {
		if i == 880 {
			fmt.Println("asdasd")
		}
		str := strconv.Itoa(i)
		digit1 := 0
		digit2 := 0
		digit3 := 0
		if len(str) == 1 {
			digit3 = i
		} else if len(str) == 2 {
			t := string(str[0])
			t2 := string(str[1])
			digit2, _ = strconv.Atoi(t)
			digit3, _ = strconv.Atoi(t2)
		} else if len(str) == 3 {
			t := string(str[0])
			t2 := string(str[1])
			t3 := string(str[2])
			digit1, _ = strconv.Atoi(t)
			digit2, _ = strconv.Atoi(t2)
			digit3, _ = strconv.Atoi(t3)
		}
		_total := 0

		for _, sd := range seta {
			if sd.Setname == "A" {
				if sd.Setnumber == int64(digit1) {
					_total = _total + (int(sd.Setsales) * 100)
				}
			}
			if sd.Setname == "B" {
				if sd.Setnumber == int64(digit2) {
					_total = _total + (int(sd.Setsales) * 100)
				}
			}
			if sd.Setname == "C" {
				if sd.Setnumber == int64(digit3) {
					_total = _total + (int(sd.Setsales) * 100)
				}
			}

			//
			if sd.Setname == "AB" {

				two := strconv.Itoa(digit1) + strconv.Itoa(digit2)
				twoint, _ := strconv.Atoi(two)
				if sd.Setnumber == int64(twoint) {
					_total = _total + (int(sd.Setsales) * 700)
				}
			}
			if sd.Setname == "BC" {
				two := strconv.Itoa(digit2) + strconv.Itoa(digit3)
				twoint, _ := strconv.Atoi(two)
				if sd.Setnumber == int64(twoint) {
					_total = _total + (int(sd.Setsales) * 700)
				}
			}
			if sd.Setname == "AC" {
				two := strconv.Itoa(digit1) + strconv.Itoa(digit3)
				twoint, _ := strconv.Atoi(two)
				if sd.Setnumber == int64(twoint) {
					_total = _total + (int(sd.Setsales) * 700)
				}
			}
			//abc full
			if sd.Setname == "ABCFULL" {
				two := strconv.Itoa(digit1) + strconv.Itoa(digit2) + strconv.Itoa(digit3)
				twoint, _ := strconv.Atoi(two)
				if sd.Setnumber == int64(twoint) {
					_total = _total + (int(sd.Setsales) * 24000)
				}
			}
			if sd.Setname == "ABCHALF" {
				two := strconv.Itoa(digit1) + strconv.Itoa(digit2) + strconv.Itoa(digit3)
				twoint, _ := strconv.Atoi(two)
				if sd.Setnumber == int64(twoint) {
					_total = _total + (int(sd.Setsales) * 12000)
				}
			}
		}
		secnumber := strconv.Itoa(digit1) + strconv.Itoa(digit2) + strconv.Itoa(digit3)
		setnumbers[secnumber] = int64(_total)

	}
	var lowprice int64
	var winningnumber string
	loopi := 0
	for in, k := range setnumbers {
		//fmt.Println(in, "====>", k)
		if loopi == 0 {
			lowprice = k
		}
		loopi++
		if k < lowprice {
			lowprice = k
			winningnumber = in
		}
	}
	//fmt.Println("winningnumber->", winningnumber)
	type result struct {
		RData         []sets
		Winningnumber string
		Lowprice      int64
	}
	if p.CustomNumber != "" {
		winningnumber = p.CustomNumber
		lowprice = setnumbers[p.CustomNumber]
	}
	newr := result{}
	newr.RData = seta
	newr.Winningnumber = winningnumber
	newr.Lowprice = lowprice
	c.ResSuccess(w, newr, "")

}

func Get_DailyData(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)

	db = orm.Db()
	repos := []c.DailyData{}

	type playload struct {
		Date          time.Time
		Distributorid string
		Showid        string
	}
	p := playload{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	loc, _ := time.LoadLocation("Asia/Kolkata")
	time := p.Date.In(loc)

	y, m, d := time.Date()
	datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

	QRY := fmt.Sprintf(`select * from dailydata where distributorid='%s' and showid='%s' and date='%s'`, p.Distributorid, p.Showid, datestring)

	rows, err := db.Query(QRY)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.DailyData{}
		err = rows.Scan(
			&repo.ID,
			&repo.Date,
			&repo.Showid,
			&repo.Distributorid,
			&repo.A,
			&repo.SetA,
			&repo.B,
			&repo.SetB,
			&repo.C,
			&repo.SetC,
			&repo.AB,
			&repo.SetAB,
			&repo.AC,
			&repo.SetAC,
			&repo.BC,
			&repo.SetBC,
			&repo.ABCFULL,
			&repo.SetABCFULL,
			&repo.ABCHALF,
			&repo.SetABCHALF,
			&repo.BOXFULL,
			&repo.SetBOXFULL,
			&repo.BOXHALF,
			&repo.SetBOXHALF,
			&repo.Status,
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

func Delete_DailyData(w http.ResponseWriter, r *http.Request) {
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

func Save_DailyData(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := []c.DailyData{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	db = orm.Db()
	repos := []c.DailyData{}

	for _, pl := range p {
		loc, _ := time.LoadLocation("Asia/Kolkata")
		time := pl.Date.In(loc)

		y, m, d := time.Date()
		datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)
		QRY := ""
		if pl.ID > 0 {
			QRY = fmt.Sprintf(`UPDATE DailyData set date='%s',showid='%s',distributorid='%s',a='%v',seta='%v',b='%v',setb='%v',c='%v',setc='%v',ab='%v',setab='%v',ac='%v',setac='%v',bc='%v',setbc='%v',abcfull='%v',setabcfull='%v',abchalf='%v',setabchalf='%v',boxfull='%v',setboxfull='%v',boxhalf='%v',setboxhalf='%v',status='%v' where id='%v';`,
				datestring, pl.Showid, pl.Distributorid, pl.A, pl.SetA, pl.B, pl.SetB, pl.C, pl.SetC, pl.AB, pl.SetAB, pl.AC, pl.SetAC, pl.BC, pl.SetBC, pl.ABCFULL, pl.SetABCFULL, pl.ABCHALF, pl.SetABCHALF, pl.BOXFULL, pl.SetBOXFULL, pl.BOXHALF, pl.SetBOXHALF, pl.Status, pl.ID)
		} else {
			QRY = fmt.Sprintf(`INSERT INTO DailyData (date,showid,distributorid,a,seta,b,setb,c,setc,ab,setab,ac,setac,bc,setbc,abcfull,setabcfull,abchalf,setabchalf,boxfull,setboxfull,boxhalf,setboxhalf,status) VALUES ('%s','%s','%s','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%s');`,
				datestring, pl.Showid, pl.Distributorid, pl.A, pl.SetA, pl.B, pl.SetB, pl.C, pl.SetC, pl.AB, pl.SetAB, pl.AC, pl.SetAC, pl.BC, pl.SetBC, pl.ABCFULL, pl.SetABCFULL, pl.ABCHALF, pl.SetABCHALF, pl.BOXFULL, pl.SetBOXFULL, pl.BOXHALF, pl.SetBOXHALF, pl.Status)
		}
		rows, _ := db.Query(QRY)

		defer rows.Close()
	}

	c.ResSuccess(w, repos, "Saved Successfully")
}

func Get_DailyApproval(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)

	db = orm.Db()
	repos := []c.DailyApproval{}

	p := c.DailyApproval{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// time := p.Date.In(loc)

	// y, m, d := time.Date()
	// datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

	QRY := ""
	if p.Appliedby != "" {
		QRY = fmt.Sprintf(`select dailyapproval.id,dailyapproval.date,shows.time,dailyapproval.appliedby,dailyapproval.approvedby,dailyapproval.status,dailyapproval.distributorid,shows.id from dailyapproval LEFT JOIN shows ON shows.id = dailyapproval.showid::INTEGER where  appliedby='%s' and status='%s'`, p.Appliedby, p.Status)
	} else if p.Approvedby != "" {
		QRY = fmt.Sprintf(`select dailyapproval.id,dailyapproval.date,shows.time,dailyapproval.appliedby,dailyapproval.approvedby,dailyapproval.status,dailyapproval.distributorid,shows.id from dailyapproval LEFT JOIN shows ON shows.id = dailyapproval.showid::INTEGER where  approvedby='%s' and status='%s'`, p.Approvedby, p.Status)
	} else {
		QRY = fmt.Sprintf(`select dailyapproval.id,dailyapproval.date,shows.time,dailyapproval.appliedby,dailyapproval.approvedby,dailyapproval.status,dailyapproval.distributorid,shows.id from dailyapproval LEFT JOIN shows ON shows.id = dailyapproval.showid::INTEGER  where   status='%s'`, p.Status)
	}

	rows, err := db.Query(QRY)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.DailyApproval{}
		err = rows.Scan(
			&repo.ID,
			&repo.Date,
			&repo.Showid,
			&repo.Appliedby,
			&repo.Approvedby,
			&repo.Status,
			&repo.Distributorid,
			&repo.Showidint,
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

func Get_DailyStatus(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)

	db = orm.Db()
	repos := []c.DailyApproval{}

	p := c.DailyApproval{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	loc, _ := time.LoadLocation("Asia/Kolkata")
	time := p.Date.In(loc)

	y, m, d := time.Date()
	datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

	QRY := fmt.Sprintf(`select * from dailyapproval where distributorid='%s' and showid='%s' and date='%s'`, p.Distributorid, p.Showid, datestring)

	rows, err := db.Query(QRY)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		repo := c.DailyApproval{}
		err = rows.Scan(
			&repo.ID,
			&repo.Date,
			&repo.Showid,
			&repo.Appliedby,
			&repo.Approvedby,
			&repo.Status,
			&repo.Distributorid,
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

func Daily_approveall(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.DailyApproval{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	db = orm.Db()
	repos := []c.DailyApproval{}

	QRY := fmt.Sprintf(`UPDATE dailyapproval SET status='%s',approvedby='%s' where id='%v';`,
		"Approved", p.Approvedby, p.ID)

	rows, errs := db.Query(QRY)

	if errs != nil {
		c.ResFail(w, nil, errs.Error())
		return
	}

	defer rows.Close()

	c.ResSuccess(w, repos, "Approved Successfully")
}

func Submitforapproval(w http.ResponseWriter, r *http.Request) {
	c.SetupResponse(&w, r)
	p := c.DailyApproval{}
	//c.GetPlayload(r, w, &p)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		c.ResFail(w, nil, "Not a valid Input"+err.Error())
		return
	}

	db = orm.Db()
	repos := []c.DailyApproval{}

	loc, _ := time.LoadLocation("Asia/Kolkata")
	time := p.Date.In(loc)

	y, m, d := time.Date()
	datestring := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

	QRY := fmt.Sprintf(`INSERT INTO dailyapproval (date,showid,distributorid,status,appliedby,approvedby) VALUES ('%s','%s','%s','%s','%s','');`,
		datestring, p.Showid, p.Distributorid, "User-Submitted", p.Appliedby)

	rows, _ := db.Query(QRY)

	defer rows.Close()

	c.ResSuccess(w, repos, "Submitted Successfully")
}

func Update_DailyData(w http.ResponseWriter, r *http.Request) {
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
