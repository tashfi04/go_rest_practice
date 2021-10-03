package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go_rest_test/config"
	"go_rest_test/helpers"
	"go_rest_test/models"
	"log"
	"net/http"
	"strconv"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	query, err := config.Db.Prepare("INSERT User SET user_name=?, user_type=?")
	helpers.Catch(err)
	defer query.Close()

	result, err := query.Exec(user.UserName, user.UserType)
	helpers.Catch(err)

	id, err := result.LastInsertId()
	helpers.Catch(err)
	helpers.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created", "id": strconv.FormatInt(id, 10)})
}

func UserList(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	query, err := config.Db.Query("SELECT * FROM User")
	helpers.Catch(err)
	defer query.Close()

	for query.Next() {
		var user models.User
		err := query.Scan(&user.Id, &user.UserName, &user.UserType)
		helpers.Catch(err)
		users = append(users, user)
	}

	helpers.RespondWithJSON(w, http.StatusOK, users)
}

func UserUpdate(w http.ResponseWriter, r *http.Request)  {

	var user models.User
	id := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	query, err := config.Db.Exec("UPDATE User SET user_name=?, user_type=? WHERE id=?", user.UserName, user.UserType, id)
	helpers.Catch(err)

	rowsAff, err := query.RowsAffected()
	helpers.Catch(err)
	helpers.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "successfully updated", "rows affected": strconv.FormatInt(rowsAff, 10)})
}

func UserDelete(w http.ResponseWriter, r *http.Request)  {

	id := chi.URLParam(r, "id")

	result, err := config.Db.Exec("DELETE FROM User WHERE id=?", id)
	helpers.Catch(err)

	rowsAff, err := result.RowsAffected()
	helpers.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "successfully deleted", "rows affected": strconv.FormatInt(rowsAff, 10)})
}

func UserDetails(w http.ResponseWriter, r *http.Request) {

	var user models.User
	id := chi.URLParam(r, "id")

	query := config.Db.QueryRow("SELECT * FROM User WHERE id=?", id)

	err := query.Scan(&user.Id, &user.UserName, &user.UserType)
	helpers.Catch(err)

	helpers.RespondWithJSON(w, http.StatusOK, user)
}