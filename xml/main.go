package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Post struct {
	XMLName   xml.Name  `xml:"post"`      //①XML要素名自体を取得
	Xml       string    `xml:",innerxml"` //②生の(未処理の)XMLを取得
	Id        string    `xml:"id,attr"`   //③XML要素の属性を取得
	Content   string    `xml:"content"`
	Developer Developer `xml:"developer"`
	Comments  []Comment `xml:"comments>comment"` //④ネストされた下位要素を直接取得したい場合は構造タグ「a>b」を用いる
}

type Developer struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` //⑤XML要素の文字データを取得
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Review  string `xml:"review"`
}

func main() {
	pwd, err := os.Getwd()
	xmlFile, err := os.Open(path.Join(pwd, "xml", "post_review.xml"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)

	fmt.Println(post.XMLName.Local)
	fmt.Println(post.Xml)
	fmt.Println("postID: " + post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Developer.Id + ": " + post.Developer.Name)
	fmt.Println(post.Comments[0].Id + ": " + post.Comments[0].Content + ", Review: " + post.Comments[0].Review)
	fmt.Println(post.Comments[1].Id + ": " + post.Comments[1].Content + ", Review: " + post.Comments[1].Review)
}
