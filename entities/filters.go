package entities

type FilterMeme struct {
	Name  string `param:"nameMeme"`
	Text  string `query:"text"`
	Text1 string `query:"text1"`
	Text2 string `query:"text2"`
	Text3 string `query:"text3"`
	Text4 string `query:"text4"`
	Url   string `query:"url"`
	Url1  string `query:"url1"`
	Url2  string `query:"url2"`
	Url3  string `query:"url3"`
	Url4  string `query:"url4"`
}
