# MapPLZ-Go

[MapPLZ](http://mapplz.com) is a framework to make mapping quick and easy in
your favorite language.

<img src="https://raw.githubusercontent.com/mapmeld/mapplz-go/master/logo.jpg" width="140"/>

## Getting started

MapPLZ consumes many many types of geodata. It can process data for a script or dump
it into a database.

Go does not support method overloading, so you need to use Add, Add2, or Add3 depending
on the number of parameters.

You can also call specific Add functions. Function names are based on their parameters
separated by an underscore, so send (lat, lng) to Add_Lat_Lng, send (lng, lat)
to Add_Lng_Lat, and send an array [lat, lng] to Add_LatLng.

Adding some data:

```
mapstore := mapplz.NewMapPLZ()

// add points
mapstore.Add2(40, -70)
mapstore.Add2(-70, 40)
mapstore.Add( []float64{ 40, -70 } )

// add lines
mapstore.Add_LatLngPath([][]float64{{40, -70}, {50, 20}})

// add polygons
mapstore.Add_LatLngPoly([][]float64{{40, -70}, {50, 20}, {40, 40}, {40, -70}})

// GeoJSON strings (points, lines, and polygons)
mapstore.Add(`{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [-70, 40] }}`)

// add properties
pt := mapstore.Add(`{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [-70, 40] }, "properties": { "color": "#0f0" }}`)
pt.Properties()["color"]
```

Each feature is added to the mapstore and returned as a MapItem

```
pt := mapstore.Add2(40, -70)
line := mapstore.Add_LatLngPath_Json([][]float64{{40, -70}, {50, 20}}, `{ "color": "red" }`)

len(mapstore.MapItems) == 2
// export all with mapstore.ToGeoJson()

pt.Lat() == 40
pt.ToGeoJson() == `{ "type": "Feature", "geometry": { "type": "Point", "coordinates": [-70, 40] }}`

line.Type() == "line"
line.Path() == [[[40, -70], [50, 20]]]
line.Properties()["color"]
```

## Packages

* <a href="https://github.com/kellydunn/golang-geo">golang-geo</a> from Kelly Dunn (MIT license)
* <a href="https://github.com/kpawlik/geojson">geojson</a> from Kris Pawlik (MIT license)

## License

Free BSD License
