package dto

/*
	{
	    "author": "Sarah Andersen",
	    "category": "webcomic",
	    "created_at": "2023-02-01T01:01:01+0700",
	    "deleted_at": "2023-02-01T01:01:01+0700",
	    "description": "Sarah's Scribbles is a webcomic by Sarah Andersen started in 2011.",
	    "id": "uuid",
	    "image_url": "",
	    "isbn": "9788833140940",
	    "language": "english",
	    "page": "150",
	    "price": "18.99",
	    "publish_date": "2021",
	    "publisher": "Tapas Media",
	    "rating": "4.5",
	    "title": "Sarah's Scribbles",
	    "updated_at": "2023-02-01T01:01:01+0700"
	}
*/
type BooksDTO struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	PublishDate string `json:"publish_date"`
	ImageUrl    string `json:"image_url"`
	Page        string `json:"page"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Language    string `json:"language"`
	CreatedAt   string `json:"created_at"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Isbn        string `json:"isbn"`
	Rating      string `json:"rating"`
	DeletedAt   string `json:"deleted_at"`
}
