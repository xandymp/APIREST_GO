package controllers

import (
	"encoding/csv"
	"encoding/json"
	"models"
	"net/http"
	"os"
	"strconv"
	u "utils"
)

var SaveFuncionario = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint)
	funcionario := &models.Funcionario{}

	err := json.NewDecoder(r.Body).Decode(funcionario)
	if err != nil {
		u.Respond(w, u.Message(false, "Erro com o Decoder"))
		return
	}

	funcionario.UserId = user

	params := r.URL.Query()
	id, erro := strconv.Atoi(params.Get("id"))
	if (erro == nil && id > 0) {
		resp := funcionario.Update(id)
		u.Respond(w, resp)
		return
	}

	resp := funcionario.Create()
	u.Respond(w, resp)
}

var ListFuncionario = func(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	id, erro := strconv.Atoi(params.Get("id"))
	nome := params.Get("nome")
	cpf := params.Get("cpf")
	cargo := params.Get("cargo")
	salarioMin := params.Get("salarioMin")
	salarioMax := params.Get("salarioMax")
	status := params.Get("status")


	if (erro == nil && id > 0) {
		data := models.FetchFuncionario(id)
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
		return
	}

	data := models.ListFuncionarioByArgs(nome, cpf, cargo, salarioMin, salarioMax, status)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteFuncionario = func(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	cpf := params.Get("cpf")

	data := models.DeleteFuncionario(cpf)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
	return
}

var ImportFuncionario = func(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("Funcionarios.csv")
	resp := u.Message(false, "Erro ao abrir o arquivo")

	if err != nil {
		u.Respond(w, resp)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	csvData, err := reader.ReadAll()

	if err != nil {
		u.Respond(w, resp)
		return
	}

	for _, linha := range csvData {

		var funcionario models.Funcionario

		if linha[0] != "Cargo" {
			funcionario.Cargo = linha[0]
		}
		if linha[1] != "Cpf" {
			funcionario.CPF = linha[1]
		}
		if linha[2] != "Nome" {
			funcionario.Nome = linha[2]
		}
		if linha[3] != "UfNasc" {
			funcionario.UFNasc = linha[3]
		}
		if linha[4] != "Salario" {
			funcionario.Salario = linha[4]
		}
		if linha[5] != "Status" {
			funcionario.Status = linha[5]
		}

		funcionario.UserId = 1

		if (funcionario.Cargo != "" && funcionario.CPF != "" && funcionario.Nome != "" && funcionario.UFNasc != "" && funcionario.Salario != "" && funcionario.Status != "") {
			resp = funcionario.Create()
			if resp["status"] == false {
				u.Respond(w, resp)
			}
		}
	}

	resp = u.Message(true, "Importação finalizada")
	u.Respond(w, resp)
}