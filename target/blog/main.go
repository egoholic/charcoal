package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

type Association interface {
	Name() string
	Find(interface{}) interface{} // function that returns associated data
}
type Content struct {
	title string
	body  string
}

func NewContent(title, body string) *Content {
	return &Content{title, body}
}
func (c *Content) Title() string {
	return c.title
}
func (c *Content) Body() string {
	return c.body
}
func (c *Content) PK() string {
	return c.Title()
}

type Publisher struct {
	Account
}

type Publication struct {
	publishedAt time.Time
	content_key string
	content     *Content
	publishedBy *Publisher
}

func NewPublication(c *Content, p *Publisher) *Publication {
	return &Publication{time.Now(), c, p}
}

func (p *Publication) PK() string {
	return p.content.PK()
}

func (p *Publication) Content(find ContentFinder) *Content {
	if p.content != nil {
		return p.content
	}

	content = find(p)
}

type Account struct {
	email string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Can't connect to mongodb://localhost:27017")
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Can't ping the server!")
	} else {
		fmt.Println("Ping!")
	}

	blog := client.Database("blog")
	fmt.Printf("\n#####\nDatabase: %#v\n#####\n", *blog)

	publications := blog.Collection("publications")
	fmt.Printf("\n#####\nCollection: %#v\n#####\n", *publications)

	document := bson.D{{"title", "Title"},
		{"content", "Some content"}}
	result, err := publications.InsertOne(ctx, document)
	if err != nil {
		fmt.Printf("Can't insert a document! %s\n", err.Error())
	} else {
		fmt.Printf("InserOne() result: %#v\n", *result)
		fmt.Printf("ID: %s\n", result.InsertedID)
	}

	findOneResult := publications.FindOne(ctx, bson.D{{"_id", result.InsertedID}})

	fmt.Printf("FindOne() result: %#v\n", *findOneResult)
	publication := bson.D{}
	findOneResult.Decode(&publication)
	fmt.Printf("Publication: %#v\n", publication)
	fmt.Println("id:", publication[0].Value)
	fmt.Println("title:", publication[1].Value)
	fmt.Println("content:", publication[2].Value)

	docs := []interface{}{
		bson.D{{"title", "Title-2"}, {"content", "Content-2"}},
		bson.D{{"title", "Title-3"}, {"content", "Content-3"}},
		bson.D{{"title", "Title-4"}, {"content", "Content-4"}},
		bson.D{{"title", "Title-5"}, {"content", "Content-5"}},
		bson.D{{"title", "Title-6"}, {"content", "Content-6"}},
		bson.D{{"title", "Title-7"}, {"content", "Content-7"}},
		bson.D{{"title", "Title-8"}, {"content", "Content-8"}},
		bson.D{{"title", "Title-9"}, {"content", "Content-9"}},
		bson.D{{"title", "Title-10"}, {"content", "Content-10"}}}

	insertManyResult, err := publications.InsertMany(ctx, docs)
	if err != nil {
		fmt.Printf("Error with InsertMany(): %s\n", err.Error())
		return
	}
	fmt.Printf("insertManyResult: %#v\n", insertManyResult)

	cur, err := publications.Find(ctx, bson.D{{"title", bson.D{{"$in", bson.A{"Title", "Title-3", "Title-5"}}}}})

	if err != nil {
		fmt.Printf("Error with Find(): %s\n", err.Error())
		return
	}
	defer cur.Close(ctx)
	fmt.Printf("Cursor: %#v\n", cur)

	var d bson.D
	pubs := []bson.D{}
	i := 1
	for cur.Next(ctx) {
		cur.Decode(&d)
		fmt.Printf("#%d ---> %#v\n\n", i, d)
		pubs = append(pubs, d)
		i++
	}
	fmt.Println("###########")
	now := time.Now()
	dr1, _ := time.ParseDuration("1h30m30")
	dr2, _ := time.ParseDuration("50m15")
	dr3, _ := time.ParseDuration("1h11m")
	dr4, _ := time.ParseDuration("2h02m")
	dr5, _ := time.ParseDuration("03m33")

	docs = []interface{}{
		bson.D{{"title", "Some Title 1"},
			{"content", "Content-1"},
			{"publishedAt", now}},
		bson.D{{"title", "Some Title 2"},
			{"content", "Content-2"},
			{"publishedAt", now.Add(dr1)}},
		bson.D{{"title", "Some Title 3"},
			{"content", "Content-3"},
			{"publishedAt", now}},
		bson.D{{"title", "Some Title 4"},
			{"content", "Content-4"},
			{"publishedAt", now.Add(dr2)}},
		bson.D{{"title", "Some Title 5"},
			{"content", "Content-5"},
			{"publishedAt", now.Add(dr3)}},
		bson.D{{"title", "Some Title 6"},
			{"content", "Content-6"},
			{"publishedAt", now}},
		bson.D{{"title", "Some Title 7"},
			{"content", "Content-7"},
			{"publishedAt", now.Add(dr4)}},
		bson.D{{"title", "Some Title 8"},
			{"content", "Content-8"},
			{"publishedAt", now.Add(dr5)}},
		bson.D{{"title", "Some Title 9"},
			{"content", "Content-9"},
			{"publishedAt", now}},
	}

	insertManyResult, err = publications.InsertMany(ctx, docs)
	if err != nil {
		fmt.Printf("Error with InsertMany(): %s\n", err.Error())
		return
	}
	fmt.Printf("insertManyResult: %#v\n", insertManyResult)

	defaultDate, _ := time.Parse(time.RFC3339, "2000-11-23T00:00:01Z")
	updateManyResult, err := publications.UpdateMany(ctx, bson.D{{"publishedAt", bson.D{{"$exists", false}}}}, bson.D{{"$set", bson.D{{"publishedAt", defaultDate}}}})
	if err != nil {
		fmt.Printf("Error with UpdateMany(): %s\n", err.Error())
		return
	}
	fmt.Printf("updateManyResult: %#v\n", updateManyResult)

	pubs2 := []bson.D{}
	var d2 bson.D
	findOpts := &options.FindOptions{}
	findOpts = findOpts.SetLimit(5)
	findOpts = findOpts.SetSort(bson.D{{"publishedAt", -1}})
	cur, err = publications.Find(ctx, bson.D{}, findOpts)
	if err != nil {
		fmt.Printf("Error with Find(): %s\n", err.Error())
		return
	}
	defer cur.Close(ctx)
	i = 1
	for cur.Next(ctx) {
		cur.Decode(&d2)
		fmt.Printf("#%d ---> %#v\n\n", i, d2)
		pubs2 = append(pubs2, d2)
		i++
	}
	fmt.Println("Ok!")

}
