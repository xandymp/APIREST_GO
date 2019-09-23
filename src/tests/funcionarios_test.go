package tests

import (
	"controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFuncionarioById(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/funcionario/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.ListFuncionario)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Erro no código de retorno: atual - %v | esperado - %v",
			status, http.StatusOK)
	}

	expected := `{
		"data": [
			{
				"ID": 1,
				"CreatedAt": "2019-09-22T01:41:37.154115-03:00",
				"UpdatedAt": "2019-09-22T01:41:37.154115-03:00",
				"DeletedAt": null,
				"Cargo": "Dev",
				"CPF": "36735109830",
				"Nome": "Alexandre Manoel",
				"UFNasc": "SP",
				"Salario": "6000",
				"Status": "ATIVO",
				"UserId": 1
			}
		],
		"message": "success",
		"status": true
	}`
	if rr.Body.String() != expected {
		t.Errorf("Erro no retorno: atual - %v | esperado - %v",
			rr.Body.String(), expected)
	}
}