package categories

type CategoryPublic struct {
	//ID is nanoid not primary key
	ID          string `db:"public_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ParentID    string `db:"parent_id"`
	CreatedBy   string `db:"users.public_id"`
}

// type ParentCategoryPublic struct {
// 	//ID is nanoid not primary key
// 	ID          string `db:"public_id"`
// 	Name        string `db:"name"`
// 	Description string `db:"description"`
// 	CreatedBy   string `db:"users.public_id"`
// }

// used for db queries for all fields of the table
type Category struct {
	ID          string `db:"id"`
	PublicID    string `db:"public_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ParentID    string `db:"parent_id"`
	CreatedBy   string `db:"created_by"`
}

type NewCategoryRequest struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
	ParentID    string `json:"ParentID"`
	CreatedBy   string `json:"CreatedBy" binding:"required"`
}

type EditCategoryRequest struct {
	ID          string `json:"ID" binding:"required"`
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
}
