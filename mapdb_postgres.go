package mapplz

import (
  _ "github.com/lib/pq"
  "database/sql"
  "encoding/json"
  "log"
  "bytes"
)

type PSQLDatabase struct {
  db *sql.DB
}

func NewPostGISDB() * PSQLDatabase {
  return &PSQLDatabase{}
}

func (psql *PSQLDatabase) Type() string {
  return "postgis"
}

func (psql *PSQLDatabase) SetDB(dbinfo interface{}) {
  psql.db = dbinfo.(*sql.DB)
}

func (psql *PSQLDatabase) Add(mip MapItem) {
  var shape_id int
  props_json, _ := json.Marshal(mip.Properties())
  props_str := string(props_json)
  wkt := mip.ToWKT()

  var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("props_str: %s", props_str)
  logger.Print("wkt: %s", wkt)

  psql.db.QueryRow("INSERT INTO mapplz (properties, geom) VALUES ('" + props_str + "', ST_GeomFromText('" + wkt + "')) RETURNING id'").Scan(&shape_id)
}

func (psql *PSQLDatabase) Query() []MapItem {
  var mitems []MapItem
  rows, err := psql.db.Query("SELECT id, ST_AsGeoJSON(geom) AS geo, properties FROM mapplz")

  var buf bytes.Buffer
  logger := log.New(&buf, "logger: ", log.Lshortfile)
  logger.Print("err: %s", err)

  for rows.Next() {
    var geo string
    rows.Scan(&geo)
    faker := NewMapPLZ()
    mip := faker.Add(geo)
    mitems = append(mitems, mip)
  }
  return mitems
}
