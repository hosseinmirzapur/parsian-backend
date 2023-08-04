package models

type Admin struct {
	BaseModel
	Username   string `gorm:"type:string;size:20;not null;unique"`
	Name       string `gorm:"type:string;size:50;null"`
	Password   string
	AdminRoles *[]AdminRole
}

type Role struct {
	BaseModel
	Name            string `gorm:"type:string;size:20;not null;unique"`
	AdminRoles      *[]AdminRole
	RolePermissions *[]RolePermission
}

type Permission struct {
	Name            string `gorm:"type:string;size:20;not null;unique"`
	RolePermissions *[]RolePermission
}

type AdminRole struct {
	BaseModel
	Admin   Admin `gorm:"foreignkey:AdminId;constraint:onDelete:CASCADE"`
	Role    Role  `gorm:"foreignkey:RoleId;constraint:onDelete:CASCADE"`
	AdminId uint
	RoleId  uint
}

type RolePermission struct {
	Role         Role       `gorm:"foreignkey:RoleId;constraint:onDelete:CASCADE"`
	Permission   Permission `gorm:"foreignkey:PermissionId;constraint:onDelete:CASCADE"`
	RoleId       uint
	PermissionId uint
}
