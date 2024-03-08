package models

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max player"`
}

type Rooms struct {
	ID        int    `json:"id"`
	Room_name string `json:"room name"`
	Id_game   int    `json:"id games"`
}

type OutputRoomGetAll struct {
	ID        int    `json:"id"`
	Room_name string `json:"room name"`
}
type RoomGetAll struct {
	Room OutputRoomGetAll `json:"rooms"`
}

type RoomsResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []RoomGetAll `json:"data"`
}

type OutputParticipantGetAllDetail struct {
	ID         int    `json:"id"`
	Id_account string `json:"id_account"`
	Username   string `json:"username"`
}

type OutputRoomGetAllDetail struct {
	ID           int                           `json:"id"`
	Room_name    string                        `json:"room name"`
	Participants OutputParticipantGetAllDetail `json:"Participants"`
}

type RoomGetAllDetail struct {
	Room OutputRoomGetAllDetail `json:"rooms"`
}
type RoomsResponseGetAllDetail struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    []RoomGetAllDetail `json:"data"`
}

type Participants struct {
	ID         int `json:"id"`
	Id_room    int `json:"id room"`
	Id_account int `json:"id account"`
}

type ErrorResponse struct {
	Status int `json:"status"`
}

type SuccessResponse struct {
	Status int `json:"status"`
}
