// リスト7.5など
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`             // 規則1
	Id       string    `xml:"id,attr"`          // 規則2 XML要素postのid属性を対応付ける
	Content  string    `xml:"content"`          // 規則5 構造体タグを使って、構造体Post内部のフィールドContentに対応付ける
	Author   Author    `xml:"author"`           // 規則5 構造体タグを使って、構造体Post内部のフィールドAuthorに対応付ける
	Xml      string    `xml:",innerxml"`        // 規則4 XML要素post内の生のXMLを取得
	Comments []Comment `xml:"comments>comment"` // 規則6 構造体タグを使って、このフィールドをXML下位要素commentに対応付ける
}

// 属性idを取り込むため、Postとは別に構造体Authorを定義
type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` // 規則3 XML要素としてではなく，文字データとして対応付ける
}

type Comment struct {
	Id      string `xml:"id,attr"` // 規則2
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// post.xmlの入力をすベて読み込む（大きいファイルには適さない）
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println("##### 構造体Post #####")
	fmt.Println(post.XMLName.Local)
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author)
	fmt.Println(post.Xml)

	fmt.Println("\n##### 構造体Author #####")
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)

	fmt.Println("\n##### 構造体Comment #####")
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
	fmt.Println(post.Comments[1].Id)
	fmt.Println(post.Comments[1].Content)
	fmt.Println(post.Comments[1].Author)
}
