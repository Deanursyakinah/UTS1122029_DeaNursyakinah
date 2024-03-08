package controller

import (
	m "UTS_1122029/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT r.id, r.room_name FROM rooms r"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var roomDetail m.RoomGetAll
	var roomsDetails []m.RoomGetAll
	for rows.Next() {
		if err := rows.Scan(
			&roomDetail.Room.ID, &roomDetail.Room.Room_name); err != nil {
			log.Println("Error :", err)
			print(err.Error())
			return
		} else {
			roomsDetails = append(roomsDetails, roomDetail)
		}
	}
	var response m.RoomsResponse
	w.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Data = roomsDetails
	json.NewEncoder(w).Encode(response)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT r.id, r.room_name, p.id, p.id_account, a.username FROM rooms r INNER JOIN participants p ON r.id = p.id_room INNER JOIN account a ON p.id_account = a.id"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	var roomDetail m.RoomGetAllDetail
	var roomsDetails []m.RoomGetAllDetail
	for rows.Next() {
		if err := rows.Scan(
			&roomDetail.Room.ID, &roomDetail.Room.Room_name, &roomDetail.Room.Participants.ID, &roomDetail.Room.Participants.Id_account, &roomDetail.Room.Participants.Username); err != nil {
			log.Println("Error :", err)
			print(err.Error())
			return
		} else {
			roomsDetails = append(roomsDetails, roomDetail)
		}
	}
	var response m.RoomsResponseGetAllDetail
	w.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Message = "Success"
	response.Data = roomsDetails
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
		return
	}

	roomName := r.Form.Get("room_name")
	gameID := r.Form.Get("id_game")

	if roomName == "" || gameID == "" {
		SendErrorResponse(w, http.StatusBadRequest)
		return
	}

	idGame, err := strconv.Atoi(gameID)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest)
		return
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM games WHERE id = ?", idGame).Scan(&count)
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}

	if count == 0 {
		SendErrorResponse(w, 404)
		return
	}

	var maxPlayers int
	err = db.QueryRow("SELECT max_player FROM games WHERE id = ?", idGame).Scan(&maxPlayers)
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}

	var participants int
	err = db.QueryRow("SELECT COUNT(*) FROM rooms WHERE id_game = ?", idGame).Scan(&participants)
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}

	if participants >= maxPlayers {
		SendErrorResponse(w, 500)
		return
	}

	query := "INSERT INTO rooms (room_name, id_game) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(roomName, idGame)
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}
	if rowsAffected == 0 {
		SendErrorResponse(w, 500)
		return
	}
	SendSuccesResponse(w, 200)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	ID := mux.Vars(r)["id"]

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM participants WHERE id = ?", ID).Scan(&count)
	if err != nil {
		SendErrorResponse(w, 500)
		return
	}
	if count == 0 {
		SendErrorResponse(w, 404)
		return
	}

	query := "DELETE FROM participants WHERE id = ?"

	_, err = db.Exec(query, ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendSuccesResponse(w, 200)
}
