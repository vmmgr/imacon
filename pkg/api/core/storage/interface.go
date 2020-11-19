package storage

import "C"
import "github.com/jinzhu/gorm"

var Broadcast = make(chan FileTransfer)

type Search int

const (
	SearchID      = Search(C.SearchID)
	SearchGroupID = Search(C.SearchID)
	SearchType    = Search(C.SearchType)
	SearchAdmin   = Search(C.SearchAdmin)
)

type Update int

const (
	UpdateGroup = Update(C.SearchUUID)
	UpdatePath  = Update(C.SearchPath)
	UpdateAll   = Update(C.SearchID)
)

type Storage struct {
	gorm.Model
	GroupID   uint   `json:"group_id"`
	Type      uint   `json:"type"` //0: ISO 1:Image
	Path      string `json:"path"` //node側のパス
	CloudInit *bool  `json:"cloud_init"`
	MinCPU    uint   `json:"min_cpu"`
	MinMem    uint   `json:"min_mem"`
	OS        string `json:"os"`
	Admin     *bool  `json:"admin"`
	Lock      *bool  `json:"lock"` //削除保護
}

type Add struct {
	GroupID   uint   `json:"group_id"`
	Type      uint   `json:"type"` //0: ISO 1:Image
	FileName  string `json:"file_name"`
	MinCPU    uint   `json:"min_cpu"`
	MinMem    uint   `json:"min_mem"`
	OS        string `json:"os"`
	CloudInit bool   `json:"cloud_init"`
	Admin     bool   `json:"admin"`
}

type Convert struct {
	SrcFile string `json:"src_file"`
	SrcType string `json:"src_type"`
	DstFile string `json:"dst_file"`
	DstType string `json:"dst_type"`
}

type GenerateStorageXml struct {
	Storage       Storage
	Number        uint
	AddressNumber uint
}

type FileTransfer struct {
	URL         string
	CurrentSize int64
	AllSize     int64
}

type ResultOne struct {
	Status  bool    `json:"status"`
	Error   string  `json:"error"`
	Storage Storage `json:"storage"`
}

type ResultDatabase struct {
	Err     error
	Storage []Storage
}
