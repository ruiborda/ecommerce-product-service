package auth

// JwtPrivateClaims estructura que representa los claims personalizados en el token JWT
type JwtPrivateClaims struct {
	Email         string   `json:"email"`
	Roles         []string `json:"roles"`
	PermissionIds []int    `json:"permissionIds"`
}