package voud

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	sisbf "beneficiofacil.gopkg.net/site"
	"beneficiofacil.gopkg.net/site/order"
	"beneficiofacil.gopkg.net/site/order/types"
	"beneficiofacil.gopkg.net/site/sqlclient"
	newdatastore "beneficiofacil.gopkg.net/site/vendor/google.golang.org/appengine/datastore"

	"appengine"
	"appengine/datastore"
)

func AttachmentReceiver(w http.ResponseWriter, r *http.Request) {

	c := sisbf.NewContext(r, w)
	ctx := c.NewContext()

	if r.Method != "POST" {
		c.Warningf("Método não implementado: %v", r.Method)
		http.Error(w, "Não implementado", http.StatusMethodNotAllowed)
		return
	}

	req := make(map[string]string, 0)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.Warningf("Erro ao decodificar requisição: %v", err)
		http.Error(w, "Requisição inválida (cod = 1.0)", http.StatusBadRequest)
		return
	}

	data, err := base64.StdEncoding.DecodeString(req["content"])
	if err != nil {
		c.Warningf("Erro ao decodificar conteúdo: %v", err)
		http.Error(w, "Requisição inválida (cod = 2.0)", http.StatusBadRequest)
		return
	}

	content := string(data)

	numeroContrato := strings.TrimSpace(content[6:16])
	valorTotal := strings.TrimLeft(content[56:68], "0")
	dataNF := fmt.Sprintf("%s-%s-%s 23:59:59", content[184:188], content[182:184], content[180:182])
	urlNF := strings.TrimSpace(content[307:607])
	urlBoleto := strings.TrimSpace(content[607:907])

	dataMax, err := time.Parse("2006-01-02 15:04:05", dataNF)
	if err != nil {
		c.Warningf("Erro ao converter data da NF: %v", err)
		http.Error(w, "Requisição inválida (cod = 3.0)", http.StatusBadRequest)
		return
	}

	dataMin := dataMax.AddDate(0, 0, -4)

	valor, err := strconv.Atoi(valorTotal)
	if err != nil {
		c.Warningf("Erro ao converter valor: %v", err)
		http.Error(w, "Requisição inválida (cod = 4.0)", http.StatusBadRequest)
		return
	}

	cnpj, err := sqlclient.GetOrgCNPJByCodeTicket(c, numeroContrato)
	if err != nil {
		c.Warningf("Erro ao obter CNPJ do cliente: %v", err)
		http.Error(w, "Requisição inválida (cod = 5.0)", http.StatusBadRequest)
		return
	}

	var search []types.Order
	query := newdatastore.NewQuery(types.OrderKind).Filter("Total =", valor).Filter("Type =", types.PAT).Filter("CreationDate >=", dataMin).Filter("CreationDate <=", dataMax).Filter("CNPJ =", cnpj).Filter("Status =", types.OrderStatusApproved).Filter("NFE.Code =", "")
	ids, err := query.GetAll(ctx, &search)
	if err != nil {
		c.Warningf("Erro ao consultar Order: %v", err)
		http.Error(w, "Requisição inválida (cod = 6.0)", http.StatusBadRequest)
		return
	}

	if len(search) == 0 {
		c.Warningf("Não foi encontrado Order: %#v", search)
		http.Error(w, "Requisição inválida (cod = 7.0)", http.StatusBadRequest)
		return
	}

	if len(search) > 1 {
		c.Warningf("Foram encontrados %d Orders: %v", len(search), search)
		http.Error(w, "Requisição inválida (cod = 8.0)", http.StatusBadRequest)
		return
	}

	orderID := ids[0].IntID()

	urlParamsNF := strings.Split(urlNF, "?")
	queryParametersNF := strings.Split(urlParamsNF[1], "&")

	var nf int64
	var code string
	for _, queryParameter := range queryParametersNF {

		params := strings.Split(queryParameter, "=")
		param := params[0]
		value := params[1]

		if param == "nf" {
			nf, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				c.Warningf("Erro ao converter nf: %v", err)
				http.Error(w, "Requisição inválida (cod = 9.0)", http.StatusBadRequest)
				return
			}
		}

		if param == "verificacao" {
			code = value
		}

	}

	err = datastore.RunInTransaction(c, func(c appengine.Context) error {
		ord, err := order.GetOrder(c, orderID)
		if err != nil {
			c.Warningf("Erro ao obter Order: %v", err)
			http.Error(w, "Requisição inválida (cod = 10.0)", http.StatusBadRequest)
			return err
		}

		ord.PaymentURL = urlBoleto
		ord.NFE.Number = nf
		ord.NFE.Code = code

		return order.Put(c, ord, nil)
	}, nil)

	if err = sqlclient.UpdateNF(ctx, orderID, nf, code); err != nil {
		c.Warningf("Erro ao atualizar NF: %v", err)
		http.Error(w, "Requisição inválida (cod = 11.0)", http.StatusBadRequest)
		return
	}

}
