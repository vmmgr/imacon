package storage

import "github.com/jinzhu/gorm"

var Broadcast = make(chan FileTransfer)

const (
	SearchID      = 1
	SearchName    = 2
	SearchUUID    = 3
	SearchGroupID = 4
	SearchType    = 5
	SearchAdmin   = 6
	UpdateGroup   = 100
	UpdatePath    = 101
	UpdateAll     = 102
)

type Storage struct {
	gorm.Model
	GroupID   uint   `json:"group_id"` //0: All 1~: Only Group
	Type      uint   `json:"type"`     //0: ISO 1:Image
	Path      string `json:"path"`     //node側のパス
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	CloudInit *bool  `json:"cloud_init"` //cloud-init対応イメージであるか否か
	MinCPU    uint   `json:"min_cpu"`
	MinMem    uint   `json:"min_mem"`
	OS        string `json:"os"`
	Admin     *bool  `json:"admin"` //管理者専用イメージであるか否か
	Lock      *bool  `json:"lock"`  //削除保護
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

type Get struct {
	Path      string `json:"path"`
	CloudInit bool   `json:"cloudinit"`
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
