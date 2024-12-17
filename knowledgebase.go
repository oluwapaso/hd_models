package models

import "database/sql"

type AddKnowledgebaseCatParams struct {
	Category_Name string
	Slug          string
	Tx            *sql.Tx
}

type CheckCateNameParams struct {
	Category_Name string
}

type CheckSlugParams struct {
	CheckBy string
	Slug    string
	Item_ID int
}

type ListKnowledgebaseCatsParams struct {
	Type      string
	Fields    string
	Page_Name string
	StartFrom int
	Limit     int
}

type LoadSingleCategoryParams struct {
	Slug   string
	Fields string
}

type UpdateKnowledgebaseCatParams struct {
	Category_Name string
	Category_ID   int
	Slug          string
	Tx            *sql.Tx
}

type DeleteKnowledgebaseCatParams struct {
	Category_ID int
	Tx          *sql.Tx
}

type AddTopicParams struct {
	Topic_Name  string
	Category_ID int
	Slug        string
	Tx          *sql.Tx
}

type ListTopicsParams struct {
	Type      string
	Fields    string
	Page_Name string
	StartFrom int
	Limit     int
}

type DeleteTopicParams struct {
	Topic_ID int
	Tx       *sql.Tx
}

type GetTopicParams struct {
	Slug      string
	Search_By string
	Topic_ID  int
	Fields    string
}

type UpdateTopicParams struct {
	Slug        string
	Title       string
	Topic_ID    int
	Category_ID int
	Excerpt     string
	Topic_Body  string
	Save_Type   string
	Tx          *sql.Tx
}

type ListCategoryTopicsParams struct {
	Category_ID int
	Fields      string
}
