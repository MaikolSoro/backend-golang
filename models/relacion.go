package models

/*Relacion modelo, para obtener la relacion de un usuario con tro*/
type Relacion struct {
	UsuarioID  string `bson:"usuarioid" json:"usuarioId"`
	RelacionID string `bson:"relacionid" json:"relacionId"`
}
