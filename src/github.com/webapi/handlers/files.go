package handlers

import(
	"fmt"
	"net/http"
	//"strconv"
	//"runtime"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"os"
)

const (
	MongoURL = "mongodb://10.0.1.70:27017"
	DBName = "Music"
)

func check(err error){
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func FilesHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	//n := runtime.NumCPU()
	//t := strconv.Itoa(n)
	//fmt.Fprintf(w, t)
	if MongoURL == ""{
		fmt.Println("no connection string provided.")
		os.Exit(1)
	}
	if DBName == ""{
		fmt.Println("no dbname provided.")
		os.Exit(1)
	}
	sess, err := mgo.Dial(MongoURL)
	check(err)

	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})
	db := sess.DB(DBName)
	id := r.Form["id"][0]
	file, err := db.GridFS("fs").OpenId(id)
	check(err)
	size := file.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	check(err)
	_, err = w.Write(buffer)
	check(err)
	err = file.Close()
	check(err)
}
