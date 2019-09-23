package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	u "utils"
)

type Funcionario struct {
	gorm.Model
	Cargo string `json:"Cargo"`
	CPF string `json:"CPF"`
	Nome string `json:"Nome"`
	UFNasc string `json:"UFNasc"`
	Salario string `json:"Salario"`
	Status string `json:"Status"`
	UserId uint `json:"UserId"` // Usuário que realizou o cadastro do funcionário
}

func (funcionario *Funcionario) Validate() (map[string] interface{}, bool) {

	if funcionario.Cargo == "" {
		return u.Message(false, "Informe o cargo do funcionário"), false
	}

	if !ValidaCPF(funcionario.CPF)  {
		return u.Message(false, "CPF inválido " + funcionario.CPF), false
	}

	if funcionario.Nome == "" {
		return u.Message(false, "Informe o nome do funcionário"), false
	}

	if funcionario.UFNasc == "" {
		return u.Message(false, "Informe a UF de nascimento do funcionário"), false
	}

	if funcionario.Salario == "" {
		return u.Message(false, "Informe o salario do funcionário"), false
	}

	if funcionario.Status == "" {
		return u.Message(false, "Informe o status do funcionário"), false
	}

	if funcionario.UserId <= 0 {
		return u.Message(false, "Usuário não informado"), false
	}

	return u.Message(true, "success"), true
}

func (funcionario *Funcionario) Create() (map[string] interface{}) {

	if resp, ok := funcionario.Validate(); !ok {
		return resp
	}

	funcionario.Nome = strings.ReplaceAll(funcionario.Nome, "'", "")

	GetDB().Create(funcionario)

	resp := u.Message(true, "success")
	resp["funcionario"] = funcionario
	return resp
}

func (funcionario *Funcionario) Update(id int) (map[string] interface{}) {

	if resp, ok := funcionario.Validate(); !ok {
		return resp
	}

	funcionario.Nome = strings.ReplaceAll(funcionario.Nome, "'", "")

	funcionarioOld := Funcionario{
		Model:   gorm.Model{},
		Cargo:   "",
		CPF:     "",
		Nome:    "",
		UFNasc:  "",
		Salario: "",
		Status:  "",
		UserId:  0,
	}

	err := GetDB().Table("funcionarios").Where("id = ? AND deleted_at IS NULL", id).First(&funcionarioOld).Error
	if err != nil {
		return nil
	}

	funcionarioOld.Nome = funcionario.Nome
	funcionarioOld.Salario = funcionario.Salario
	funcionarioOld.UFNasc = funcionario.UFNasc
	funcionarioOld.Cargo = funcionario.Cargo
	funcionarioOld.CPF = funcionario.CPF
	funcionarioOld.Status = funcionario.Status
	funcionarioOld.UserId = funcionario.UserId

	GetDB().Save(funcionarioOld)

	resp := u.Message(true, "success")
	resp["funcionario"] = funcionarioOld
	return resp
}

func FetchFuncionario(id int) (*Funcionario) {

	funcionario := &Funcionario{}
	err := GetDB().Table("funcionarios").Where("id = ? AND deleted_at IS NULL", id).First(funcionario).Error
	if err != nil {
		return nil
	}
	return funcionario
}

func ListFuncionarioByArgs(
	nome string,
	cpf string,
	cargo string,
	salarioMin string,
	salarioMax string,
	status string) ([]*Funcionario) {

	nome = strings.ReplaceAll(nome, "'", "")
	cpf = strings.ReplaceAll(cpf, "'", "")
	cargo = strings.ReplaceAll(cargo, "'", "")
	status = strings.ReplaceAll(status, "'", "")
	salarioMin = strings.ReplaceAll(salarioMin, "'", "")
	salarioMax = strings.ReplaceAll(salarioMax, "'", "")

	if (status == "") {
		status = "%"
	}

	if (salarioMin == "") {
		salarioMin = "0"
	}

	if (salarioMax == "") {
		salarioMax = "999999999"
	}

	conditions := []string{"nome ILIKE ?"}
	params := []string{"%" + nome + "%"}

	conditions = append(conditions, "cpf ILIKE ?")
	params = append(params, "%" + cpf + "%")

	conditions = append(conditions, "cargo ILIKE ?")
	params = append(params, "%" + cargo + "%")

	conditions = append(conditions, "status ILIKE ?")
	params = append(params, status)

	conditions = append(conditions, "salario >= ?")
	params = append(params, salarioMin)

	conditions = append(conditions, "salario <= ?")
	params = append(params, salarioMax)

	conditions = append(conditions, "deleted_at IS NULL")

	where := strings.Join(conditions, " AND ")

	funcionarios := make([]*Funcionario, 0)
	err := GetDB().Table("funcionarios").Where(
		where,
		params[0],
		params[1],
		params[2],
		params[3],
		params[4],
		params[5]).Find(&funcionarios).Error

	if err != nil {
		return nil
	}

	return funcionarios
}

func DeleteFuncionario(cpf string) (msg string) {
	if !ValidaCPF(cpf) {
		return "CPF inválido"
	}

	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	err := GetDB().Table("funcionarios").Where("cpf = ?", cpf).Delete(Funcionario{}).Error
	if err != nil {
		return "Erro ao excluir"
	}

	return "Funcionário excluído com sucesso"
}

func ValidaCPF(cpf string) (bool) {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	if len(cpf) != 11 {
		return false
	}

	var eq bool
	var dig string

	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}

		if string(val) == dig {
			eq = true
			continue
		}

		eq = false
		break
	}

	if eq {
		return false
	}

	i := 10
	sum := 0

	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11

	if mod == 10 {
		mod = 0
	}

	digit1, _ := strconv.Atoi(string(cpf[9]))

	if mod != digit1 {
		return false
	}

	i = 11
	sum = 0

	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod = sum * 10
	mod = prod % 11

	if mod == 10 {
		mod = 0
	}

	digit2, _ := strconv.Atoi(string(cpf[10]))

	if mod != digit2 {
		return false
	}

	return true
}

