package main

import (
	"database/sql"
	"log"
	_ "./pq"
	"fmt"
	//"time"
	//"strconv"

	"github.com/influxdata/influxdb/client/v2"
)

type DbInfo struct {
	id string
	person_id string
	timestamp string
	age string
	gender string
	attention string
	interest string
	happines string
	surprise string
	anger string
	disgust string
	fear string
	sadness string
	neutrall string

}

var DBObject DbInfo
var DbArray []DbInfo






func main() {
	//timeDayAgoForQuery := getTimeStamp() //lines 42-45 to sort data by the date you need. It uses timestamp convertion.
	//stringTimeDayAgoQuery := strconv.FormatInt(int64(timeDayAgoForQuery), 10)

	//fmt.Println(`querytime` + stringTimeDayAgoQuery)
	db, err := sql.Open("postgres", "postgres://")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT *FROM public.stat`) //stringTimeDayAgoQuery ) 
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&DBObject.id, &DBObject.person_id, &DBObject.timestamp,&DBObject.age, &DBObject.gender,&DBObject.attention,&DBObject.interest,&DBObject.happines,&DBObject.surprise,&DBObject.anger,&DBObject.disgust,&DBObject.fear,&DBObject.sadness,&DBObject.neutrall)
		if err != nil {
			log.Fatal(err)
		}
		InfluxData()
		DbArray = append(DbArray, DBObject)
	}
}
func InfluxData() {

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "cola",
		Precision: "s",
	})
	// Create a point and add to batch
	tags := map[string] string{

	}
	fields := map[string]interface{}{
		"timestamp": DBObject.timestamp,
		"id": DBObject.id,
		"person_id": DBObject.person_id,
		"age": DBObject.age,
		"gender":  DBObject.gender,
		"attention": DBObject.attention,
		"interest":  DBObject.interest,
		"happines":  DBObject.happines,
		"surprise": DBObject.surprise,
		"anger":  DBObject.anger,
		"disgust": DBObject.disgust,
		"fear": DBObject.fear,
		"sadness": DBObject.sadness,
		"neutrall":  DBObject.neutrall,
		}
	pt, err := client.NewPoint("hobbit", tags, fields)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)
	// Write the batch
	c.Write(bp)
}
/*func getTimeStamp() (queryTime int64){  //This Function converts date you need in timestamp.
	t := time.Now()
	dayAgo := int64(t.AddDate(0,0,1).Unix())
	return dayAgo
}*/

