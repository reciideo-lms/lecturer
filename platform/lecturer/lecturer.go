package lecturer

type Lecturer struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
}

type Repo struct {
	Items []Lecturer
}

func New() *Repo {
	return &Repo{
		Items: []Lecturer{},
	}
}

func (r *Repo) Add(item Lecturer) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Lecturer {
	return r.Items
}
