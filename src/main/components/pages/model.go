package pages 
import (
	"lib"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"time"	
	"regexp"
	"strings"
	"github.com/SKAhack/go-shortid"
//	"log"
)

 
var FullDatabase *mgo.Database
var GsLang *lib.GsLang
var	GsLog *lib.GsLog
var	GsConfig *lib.GsConfig
var	GsAuthorize *lib.Authorize
var	GsAuthentication *lib.GsAuthentication

type Page struct {
	Id bson.ObjectId 		`json:"Id" bson:"_id,omitempty"`
	Title string			`json:"Title" bson:"Title"`
	Alias string			`json:"Alias" bson:"Alias"`
	Content string			`json:"Content" bson:"Content"`
	Lang string				`json:"Lang" bson:"Lang"`
	Tags []string 			`json:"Tags" bson:"Tags"`
	MetaTags []MetaTag		`json:"MetaTags" bson:"MetaTags"`
	CreatedAt time.Time 	`json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt time.Time 	`json:"UpdatedAt" bson:"UpdatedAt"`

}

type MetaTag struct{
	Name string `json:"Name" bson:"Name"`
	Content string `json:"Content" bson:"Content"`
}

func (p *Page) setEmpty(){
	p.Title = ""
	p.Alias = ""
	p.Content = ""
	p.Lang = ""
}
func (c *Page) pageById() (interface{}, error) {
	var result interface{}
	var err error

	if c.Id != ""{
		err = FullDatabase.C("pages").Find(bson.M{"Id": c.Id}).One(&result)
	}
	return result, err
}

func (c *Page) pageByAlias() (interface{}, error) {
	var result interface{}
	var err error

	if c.Alias != ""{
		err = FullDatabase.C("pages").Find(bson.M{"Alias": c.Alias}).One(&result)
	}
	return result, err
}


// Get pages to send to client for display
func (c *Page) getPages2Show() (interface{}, error) {
	var result []interface{}
	var err error

	err = FullDatabase.C("pages").Find(nil).Select(bson.M{"_id": 1,"Title": 1,"Alias": 1,"Lang": 1 }) .All(&result) 
	return result, err
}
// Get pages to send to client for display
func (c *Page) getPages2Manage() (interface{}, error) {
	var result []interface{}
	var err error

	err = FullDatabase.C("pages").Find(nil).All(&result) 
	return result, err
}


// Create new user
func (p *Page) createPage() (interface{}, error) {
	var result interface{}
    var err error

    // Handling page alias   -<
    if len(p.Alias)<=0 {
    	p.Alias = p.Title
    }
    p.Alias = Slug(p.Alias)
    result, err  = p.pageByAlias()
    if result != nil {
    	g := shortid.Generator()
    	g.SetSeed(1)
      	p.Alias =g.Generate()
    }
    // Handling page alias   ->

    
    p.CreatedAt = time.Now()

 	err = FullDatabase.C("pages").Insert(p)
    if err != nil {
        return "", err
    }
    GsLog.Info("Page created successfully: " + p.Alias)
    return p.Alias, err
}


func Slug(s string) string {
    var re = regexp.MustCompile("[^a-z0-9]+")
    return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

